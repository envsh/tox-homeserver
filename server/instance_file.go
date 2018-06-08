package server

import (
	"errors"
	"fmt"
	"gopp"
	"hash"
	"io/ioutil"
	"log"
	"time"

	thscom "tox-homeserver/common"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"

	tox "github.com/TokTok/go-toxcore-c"
	"github.com/dustin/go-humanize"
	"gopkg.in/h2non/filetype.v1"
)

func (this *ToxVM) setupCallbacksForFile() {
	this.setupCallbacksForFileFromTox()
	this.setupCallbacksForFileFromHttp()
}

// auto recv and save to toxhsfs
// if bigger than max, drop now
func (this *ToxVM) setupCallbacksForFileFromTox() {
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
		gopp.ErrPrint(err)
	}, nil)
}

func (this *ToxVM) setupCallbacksForFileFromHttp() {
	fso := store.GetFS()
	fso.OnFileUploaded(func(md5str string, frndpk string) {
		log.Println("New file uplaoded:", md5str)
		data, err := fso.ReadFile(md5str)
		gopp.ErrPrint(err, md5str)
		oname, err := fso.GetOrigName(md5str)
		gopp.ErrPrint(err, oname)
		log.Println(md5str, oname, len(data))

		frndnum, err := this.t.FriendByPublicKey(frndpk)
		gopp.ErrPrint(err)

		this.startSendFileData(frndnum, oname, data, func(err error) {
			log.Println("send file:", err)
		})
	})
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

func (this *ToxVM) startSendAvatar(friendNumber uint32) {

}
func (this *ToxVM) startSendFile(friendNumber uint32, kind uint32, fileName string, data []byte, donefn func(error)) {
	t := this.t
	fileId := gopp.Md5AsStr(data)
	fnum, err := t.FileSend(friendNumber, kind, uint64(len(data)), fileId, fileName)
	gopp.ErrPrint(err, fileName)
	if err != nil {
		return
	}
	fio := NewFileInfo(uint64(len(data)))
	fio.fnum = fnum
	fio.frndnum = friendNumber
	fio.fsize = uint64(len(data))
	fio.fname = fileName
	fio.fdata = data
	fio.fkind = tox.FILE_KIND_DATA
	fio.donefn = donefn
	fileSendState[fnum] = fio
}

func (this *ToxVM) onFileRecvDone(fio *FileInfo) {
	t := this.t
	frndpk, _ := t.FriendGetPublicKey(fio.frndnum)
	frndname, _ := t.FriendGetName(fio.frndnum)

	log.Println("Recv file done:", fio.fname, humanize.Bytes(fio.fsize), "from", frndname)
	md5str, err := store.GetFS().SaveFile(fio.fdata, fio.fname)
	gopp.ErrPrint(err, fio, md5str)

	ftyo, err := filetype.Match(fio.fdata)
	gopp.ErrPrint(err)
	msg := fmt.Sprintf("file;%s;%d;%s.%s", ftyo.MIME.Value, fio.fsize, md5str, ftyo.Extension)
	msgo, err := appctx.st.AddFriendMessage(msg, frndpk, frndpk, 0, 0)
	gopp.ErrPrint(err)
	appctx.st.SetMessageSent(msgo.Id)

	evto := &thspbs.Event{}
	evto.Name = "FriendMessage"
	evto.Args = []string{fmt.Sprintf("%d", fio.frndnum), msg}
	msgty := thscom.MSGTYPE_FILE
	if filetype.IsImage(fio.fdata) {
		msgty = thscom.MSGTYPE_IMAGE
	} else if filetype.IsAudio(fio.fdata) {
		msgty = thscom.MSGTYPE_AUDIO
	} else if filetype.IsVideo(fio.fdata) {
		msgty = thscom.MSGTYPE_VIDEO
	}
	evto.Margs = []string{frndname, frndpk, gopp.ToStr(msgo.EventId), "1", msgty, ftyo.MIME.Value}
	evto.EventId = msgo.EventId
	this.pubmsg(evto)
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
	md5h    *hash.Hash
	btime   time.Time
	done    bool
	donefn  func(error)
}
