package server

import (
	"errors"
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"tox-homeserver/store"
	"tox-homeserver/thscom"
	"tox-homeserver/thspbs"

	tox "github.com/TokTok/go-toxcore-c"
	"github.com/dustin/go-humanize"
	"github.com/envsh/go-toxcore/xtox"
	"gopkg.in/h2non/filetype.v1"
)

func (this *ToxVM) setupEventsForFile() {
	this.setupEventsForFileFromTox()
	this.setupEventsForFileFromHttp()
}

// auto recv and save to toxhsfs
// if bigger than max, drop now
func (this *ToxVM) setupEventsForFileFromTox() {
	t := this.t
	t.CallbackFileRecvControlAdd(func(_ *tox.Tox, friendNumber uint32, fileNumber uint32,
		control int, userData interface{}) {
		frndname, _ := t.FriendGetName(friendNumber)
		log.Println("Recv ctrl,", frndname, control, fileNumber)
		switch control {
		case tox.FILE_CONTROL_CANCEL:
			if fio, ok := fileSendState[fileNumber]; ok {
				if fio.donefn != nil {
					fio.donefn(errors.New("canceled"))
				}
			}
			delete(fileRecvState, fileNumber)
			delete(fileSendState, fileNumber)
		case tox.FILE_CONTROL_PAUSE:
			// timeout???
		case tox.FILE_CONTROL_RESUME:
			//
		}
	}, nil)

	t.CallbackFileRecvAdd(func(_ *tox.Tox, friendNumber uint32, fileNumber uint32, kind uint32, fileSize uint64,
		fileName string, userData interface{}) {
		// frndpk, _ := t.FriendGetPublicKey(friendNumber)
		frndname, _ := t.FriendGetName(friendNumber)
		if fileSize > uint64(thscom.MaxAutoRecvFileSize) {
			log.Println("File too large, now, max:", fileSize, thscom.MaxAutoRecvFileSize)
			t.FileControl(friendNumber, fileNumber, tox.FILE_CONTROL_CANCEL)
			return
		}
		log.Println("Recv file begin...", kind, fileName, humanize.Bytes(fileSize), "from", frndname)
		fio := NewFileInfo(fileSize)
		fio.fname = fileName
		fio.fsize = fileSize
		fio.fkind = kind
		fio.fnum = fileNumber
		fio.frndnum = friendNumber
		fileRecvState[fileNumber] = fio
		_, err := t.FileControl(friendNumber, fileNumber, tox.FILE_CONTROL_RESUME)
		gopp.ErrPrint(err)
	}, nil)

	chkrecvdone := func(pos uint64, data []byte, size uint64) bool {
		if pos > size {
			return true
		}
		if len(data) == 0 {
			return true
		}
		return false
	}

	t.CallbackFileRecvChunkAdd(func(_ *tox.Tox, friendNumber uint32, fileNumber uint32, position uint64,
		data []byte, userData interface{}) {
		if fio, ok := fileRecvState[fileNumber]; ok {
			if chkrecvdone(position, data, fio.fsize) {
				// push file message
				this.onFileRecvDone(fio)
				delete(fileRecvState, fileNumber)
				return
			}

			fio.WriteAt(position, data)
		} else {
			log.Println("Not found recv file info:", fileNumber, gopp.Retn(t.FriendGetName(fio.frndnum)))
		}
	}, nil)

	chksenddone := func(pos uint64, length int) bool {
		if length == 0 {
			return true
		}
		return false
	}
	// sendfile???
	t.CallbackFileChunkRequestAdd(func(_ *tox.Tox, friend_number uint32, file_number uint32, position uint64,
		length int, user_data interface{}) {
		frndpk, _ := t.FriendGetPublicKey(friend_number)
		frndname, _ := t.FriendGetName(friend_number)

		if chksenddone(position, length) {
			log.Println("Send file done.", frndname, frndpk)
			fio := fileSendState[file_number]
			if fio != nil && fio.donefn != nil {
				fio.donefn(nil)
			}
			delete(fileSendState, file_number)
			return
		}

		fio := fileSendState[file_number]
		high := gopp.IfElseInt(position+uint64(length) > uint64(len(fio.fdata)), len(fio.fdata), int(position+uint64(length)))
		chunk := fio.fdata[position:high]
		_, err := t.FileSendChunk(friend_number, file_number, position, chunk)
		gopp.ErrPrint(err, file_number, position, length, len(chunk))
	}, nil)
}

func (this *ToxVM) setupEventsForFileFromHttp() {
	locfso := store.GetFS()
	locfso.OnFileUploaded(func(md5str string, frndpk string, userCodeStr string) {
		log.Println("New file uplaoded:", md5str, frndpk, userCodeStr)
		t := this.t
		_, err := t.FriendByPublicKey(frndpk)
		if err == nil {
			this.onFriendFileUploaded(md5str, frndpk, userCodeStr)
		} else {
			_, found := xtox.ConferenceGetByIdentifier(t, frndpk)
			if found {
				this.onGroupFileUploaded(md5str, frndpk, userCodeStr)
			} else {
				log.Println("Not found dest:", frndpk)
			}
		}
	})
}

func (this *ToxVM) onFriendFileUploaded(md5str string, frndpk string, userCodeStr string) {
	locfso := store.GetFS()

	data, err := locfso.ReadFile(md5str)
	gopp.ErrPrint(err, md5str)
	oname, err := locfso.GetOrigName(md5str)
	gopp.ErrPrint(err, oname)
	log.Println(md5str, oname, len(data))

	frndnum, err := this.t.FriendByPublicKey(frndpk)
	gopp.ErrPrint(err)
	// frndname, _ := this.t.FriendGetName(frndnum)
	userCode := gopp.MustInt64(userCodeStr)
	selfpk := this.t.SelfGetPublicKey()

	evto := &thspbs.Event{}
	evto.EventName = "FriendSendMessage"
	evto.Args = gopp.ToStrs(frndnum, msgContentFromFileData(data, md5str, oname))
	evto.UserCode = userCode

	// save
	msgo, err := appctx.st.AddFriendMessage(evto.Args[1], frndpk, selfpk, 0, userCode)
	gopp.ErrPrint(err, md5str, oname)
	msgty, mimety, _ := msgTypeFromFileData(data, false)
	evto.Margs = gopp.ToStrs(0, 0, frndpk, msgty, mimety)
	evto.EventId = msgo.EventId

	// publish to other client
	this.pubmsg(evto)

	this.startSendFileData(frndnum, oname, data, func(err error) {
		log.Println("Send file to friend:", err)
		// set sent ok
		appctx.st.SetMessageSent(msgo.Id)
		// publish the last sent state
		evto.EventName = "FriendSendMessageResp"
		evto.Margs[1] = gopp.ToStr(1)
		this.pubmsg(evto)
	})
}

func (this *ToxVM) onGroupFileUploaded(md5str string, frndpk string, userCodeStr string) {
	locfso := store.GetFS()

	data, err := locfso.ReadFile(md5str)
	gopp.ErrPrint(err, md5str)
	oname, err := locfso.GetOrigName(md5str)
	gopp.ErrPrint(err, oname)
	log.Println(md5str, oname, len(data))

	frndnum, found := xtox.ConferenceGetByIdentifier(this.t, frndpk)
	gopp.FalsePrint(found, frndpk)
	// frndname, _ := this.t.FriendGetName(frndnum)
	userCode := gopp.MustInt64(userCodeStr)
	selfpk := this.t.SelfGetPublicKey()

	go func() {
		// put to external file store
		extfso := store.GetExtFS()
		urlval, err := extfso.PutFileByData(data, oname)
		gopp.ErrPrint(err, oname)
		log.Println("Send file to external file store:", err)

		evto := &thspbs.Event{}
		evto.EventName = "ConferenceSendMessage"
		// 好像url不需要存储啊。只发送出去就好了
		// evto.Args = gopp.ToStrs(frndnum, 0, msgContentFromFileDataUrl(data, md5str, oname, urlval))
		evto.Args = gopp.ToStrs(frndnum, 0, msgContentFromFileData(data, md5str, oname))
		evto.UserCode = userCode

		// save
		msgo, err := appctx.st.AddGroupMessage(evto.Args[2], "0", frndpk, selfpk, 0, userCode)
		gopp.ErrPrint(err, md5str, oname)
		msgty, mimety, mimeext := msgTypeFromFileData(data, false)
		evto.EventId = msgo.EventId
		peerName := this.t.SelfGetName()
		peerPubkey := selfpk
		title := ""
		groupId := frndpk
		evto.Margs = gopp.ToStrs(peerName, peerPubkey, title, groupId, msgo.EventId, "0",
			msgty, mimety)
		//evt.Margs = []string{peerName, peerPubkey, title, groupId, gopp.ToStr(msgo.EventId), "1",
		//	thscom.MSGTYPE_TEXT, "text/plain"}

		// send message/url to group
		// db store full FileInfoLine.String(), but send message don't
		// this.t.ConferenceSendMessage(frndnum, 0, evto.Args[1])
		b58str := thscom.Base58EncodeFromHex(md5str)
		urlmsg := gopp.IfElseStr(mimeext != "" && !strings.HasSuffix(urlval, "."+mimeext),
			fmt.Sprintf("%s.%s?%s", urlval, mimeext, b58str), urlval+"?"+b58str)
		this.t.ConferenceSendMessage(frndnum, 0, urlmsg)
		log.Println("|", urlmsg, "|")

		// publish to other client
		this.pubmsg(evto)

		// set sent ok
		appctx.st.SetMessageSent(msgo.Id)
		// publish the last sent state
		evto.EventName = "ConferenceSendMessageResp"
		evto.Margs = gopp.ToStrs(0, "1", groupId)
		this.pubmsg(evto)

		// update mime
		evto.EventName = "ConferenceMessageReload"
		this.pubmsg(evto)
	}()
}

func (this *ToxVM) testSendFile(friendNumber uint32, msg string) bool {
	frndname, _ := this.t.FriendGetName(friendNumber)
	if frndname == "envoy" {
		data, _ := ioutil.ReadFile("/home/me/Figure_1.png")
		this.startSendFileData(friendNumber, "Figure_1.png", data, nil)
		return true
	}
	return false
}

func (this *ToxVM) startSendFileData(friendNumber uint32, fileName string, data []byte, donefn func(error)) {
	this.startSendFile(friendNumber, tox.FILE_KIND_DATA, fileName, data, donefn)
}

func (this *ToxVM) startSendAvatar(friendNumber uint32, fileName string, data []byte, donefn func(error)) {
	this.startSendFile(friendNumber, tox.FILE_KIND_AVATAR, fileName, data, donefn)
}
func (this *ToxVM) startSendFile(friendNumber uint32, kind uint32, fileName string, data []byte, donefn func(error)) {
	t := this.t
	fileId := gopp.Md5AsStr(data)
	log.Println(friendNumber, kind, len(data), fileId, fileName)
	fnum, err := t.FileSend(friendNumber, kind, uint64(len(data)), fileId, fileName)
	gopp.ErrPrint(err, fileName)
	if err != nil {
		// TODO move to offline file
		return
	}
	fio := NewFileInfo(uint64(len(data)))
	fio.fnum = fnum
	fio.frndnum = friendNumber
	fio.fsize = uint64(len(data))
	fio.fname = fileName
	fio.fdata = data
	fio.md5str = fileId
	fio.fkind = tox.FILE_KIND_DATA
	fio.donefn = donefn
	fileSendState[fnum] = fio
}

func (this *ToxVM) onFileRecvDone(fio *FileInfo) {
	if fio.fkind == 1 {
		this.onFileRecvDoneAvatar(fio)
	} else {
		this.onFileRecvDoneFile(fio)
	}
}

func (this *ToxVM) onFileRecvDoneFile(fio *FileInfo) {
	t := this.t
	frndpk, _ := t.FriendGetPublicKey(fio.frndnum)
	frndname, _ := t.FriendGetName(fio.frndnum)

	log.Println("Recv file done:", fio.fname, humanize.Bytes(fio.fsize), "from", frndname)
	md5str, err := store.GetFS().SaveFile(fio.fdata, fio.fname)
	gopp.ErrPrint(err, fio, md5str)

	evto := NewEventFromFileInfo(fio, frndname, frndpk, -1)
	msg := evto.Args[1]
	msgo, err := appctx.st.AddFriendMessage(msg, frndpk, frndpk, 0, 0)
	gopp.ErrPrint(err)
	appctx.st.SetMessageSent(msgo.Id)
	evto.EventId = msgo.EventId
	evto.Margs[2] = gopp.ToStr(msgo.EventId)

	this.pubmsg(evto)
}

// don't save as a message
func (this *ToxVM) onFileRecvDoneAvatar(fio *FileInfo) {
	t := this.t
	frndpk, _ := t.FriendGetPublicKey(fio.frndnum)
	frndname, _ := t.FriendGetName(fio.frndnum)

	log.Println("Recv avatar done:", fio.fname, humanize.Bytes(fio.fsize), "from", frndname)
	err := store.GetFS().SaveFileWithName(frndpk, fio.fdata)
	gopp.ErrPrint(err, fio, frndpk)

	evto := NewEventFromFileInfo(fio, frndname, frndpk, -1)
	this.pubmsg(evto)
}

func msgContentFromFileData(data []byte, md5str string, origName string) string {
	ftyo, err := filetype.Match(data)
	gopp.ErrPrint(err)
	return thscom.NewFileInfoLine(int64(len(data)), ftyo.MIME.Value,
		md5str, ftyo.Extension, origName).String()
}
func msgContentFromFileDataUrl(data []byte, md5str string, origName string, urlval string) string {
	ftyo, err := filetype.Match(data)
	gopp.ErrPrint(err)
	return thscom.NewFileInfoLineUrl(int64(len(data)), ftyo.MIME.Value,
		md5str, ftyo.Extension, origName, urlval).String()
}
func NewEventFromFileInfo(fio *FileInfo, frndname, frndpk string, EventId int64) *thspbs.Event {
	md5str := fio.Md5Sum()
	msg := msgContentFromFileData(fio.fdata, md5str, fio.fname)

	evto := &thspbs.Event{}
	evto.EventName = "FriendMessage"
	evto.Args = []string{fmt.Sprintf("%d", fio.frndnum), msg}

	msgty, mimety, _ := msgTypeFromFileData(fio.fdata, fio.fkind == 1)
	// for file, the sent is already 1. because we save it and then here
	evto.Margs = []string{frndname, frndpk, gopp.ToStr(EventId), "1", msgty, mimety}
	evto.EventId = EventId
	return evto
}

func msgTypeFromFileData(data []byte, avatar bool) (msgty, mimety, ext string) {
	ftyo, err := filetype.Match(data)
	gopp.ErrPrint(err)
	mimety = ftyo.MIME.Value
	ext = ftyo.Extension

	msgty = thscom.MSGTYPE_FILE
	if avatar {
		msgty = thscom.MSGTYPE_AVATAR
	} else if filetype.IsImage(data) {
		msgty = thscom.MSGTYPE_IMAGE
	} else if filetype.IsAudio(data) {
		msgty = thscom.MSGTYPE_AUDIO
	} else if filetype.IsVideo(data) {
		msgty = thscom.MSGTYPE_VIDEO
	}
	return
}

func (this *FileInfo) WriteAt(pos uint64, data []byte) int {
	// log.Println(pos, len(data), this.fname, len(this.fdata))
	n := copy(this.fdata[pos:pos+uint64(len(data))], data)
	gopp.Assertf(n == len(data), "Want save %d, but %d", len(data), n)
	return n
}

// TODO save in memory is big cost
func NewFileInfo(size uint64) *FileInfo {
	this := &FileInfo{}
	this.fdata = make([]byte, size)
	this.btime = time.Now()
	return this
}

func (this *FileInfo) Md5Sum() string {
	if this.md5str == "" {
		this.md5str = gopp.Md5AsStr(this.fdata)
	}
	return this.md5str
}

var fileRecvState = map[uint32]*FileInfo{}
var fileSendState = map[uint32]*FileInfo{}
var tfrs = fileRecvState
var tfss = fileSendState

type FileInfo struct {
	fnum    uint32
	frndnum uint32
	fname   string
	fsize   uint64
	fkind   uint32
	fdata   []byte
	md5str  string
	btime   time.Time
	done    bool
	donefn  func(error)
}
