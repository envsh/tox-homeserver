package server

import (
	"context"
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"math"
	"strings"
	"time"

	"tox-homeserver/thscom"
	"tox-homeserver/thspbs"

	tox "github.com/TokTok/go-toxcore-c"
	"github.com/envsh/go-toxcore/xtox"
)

type ToxVM struct {
	t   *tox.Tox
	tav *tox.ToxAV

	autobotFeatures  int
	groupPeerPubkeys map[string][]string // groupId = > peerPubkey
	groupPeerNames   map[string]string   // peerPubkey => name
}

var tvmCtx = xtox.NewToxContext("toxhs.tsbin", "toxhs0", "i'm toxhs0")

func newToxVM() *ToxVM {
	this := &ToxVM{}
	this.groupPeerPubkeys = make(map[string][]string)
	this.groupPeerNames = make(map[string]string)

	this.t = xtox.New(tvmCtx)
	gopp.NilFatal("tox is nil")
	tav, err := tox.NewToxAV(this.t)
	gopp.ErrPrint(err)
	this.tav = tav
	this.autobotFeatures = xtox.FOTA_ADD_NET_HELP_BOTS | xtox.FOTA_REMOVE_ONLY_ME_ALL |
		xtox.FOTA_ACCEPT_FRIEND_REQUEST // | xtox.FOTA_ACCEPT_GROUP_INVITE
	xtox.SetAutoBotFeatures(this.t, this.autobotFeatures)
	this.setupEvents()

	err = xtox.Connect(this.t)
	gopp.ErrPrint(err)

	this.exportFriendInfoToDb()
	return this
}

func (this *ToxVM) exportFriendInfoToDb() {
	for i := uint32(0); i < math.MaxUint32; i++ {
		if !this.t.FriendExists(i) {
			break
		}
		name, err := this.t.FriendGetName(i)
		stmsg, err := this.t.FriendGetStatusMessage(i)
		pubkey, err := this.t.FriendGetPublicKey(i)
		gopp.ErrPrint(err)
		_, err = appctx.st.AddFriend(pubkey, i, name, stmsg)
		gopp.ErrPrint(err)
	}
}

func (this *ToxVM) setupEvents() {
	this.setupEventsForMessage()
	this.setupEventsForFile()
	this.setupEventsForAV()
}

func (this *ToxVM) setupEventsForMessage() {
	t := this.t

	t.CallbackSelfConnectionStatusAdd(func(_ *tox.Tox, status int, userData interface{}) {
		if status == tox.CONNECTION_NONE {
		} else {
		}
		log.Println(status, tox.ConnStatusString(status))
		evt := thspbs.Event{}
		evt.EventName = "SelfConnectionStatus"
		evt.Args = []string{fmt.Sprintf("%d", status)}
		evt.Margs = []string{tox.ConnStatusString(status)}

		this.pubmsg(&evt)
	}, nil)

	t.CallbackFriendRequestAdd(func(_ *tox.Tox, pubkey string, message string, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "FriendRequest"
		evt.Args = []string{pubkey, message}

		ctid, err := appctx.st.AddFriend(pubkey, 0, "", "")
		gopp.ErrPrint(err)
		evt.Margs = append(evt.Margs, fmt.Sprintf("%d", ctid))

		this.pubmsg(&evt)
	}, nil)

	t.CallbackFriendMessageAdd(func(_ *tox.Tox, friendNumber uint32, message string, userData interface{}) {
		// if this.testSendFile(friendNumber, message) {
		//	return
		// }
		evt := thspbs.Event{}
		evt.EventName = "FriendMessage"
		evt.Args = []string{fmt.Sprintf("%d", friendNumber), message}
		friendpk, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		fname, err := t.FriendGetName(friendNumber)
		gopp.ErrPrint(err)

		selfpk := t.SelfGetPublicKey()
		msgo, err := appctx.st.AddFriendMessage(message, friendpk, selfpk, 0, 0)
		gopp.ErrPrint(err)
		if err == nil {
			err := appctx.st.SetMessageSent(msgo.Id)
			gopp.ErrPrint(err, msgo.Id)
		}

		evt.Margs = []string{fname, friendpk, fmt.Sprintf("%d", msgo.EventId), "1", thscom.MSGTYPE_TEXT, "text/plain"}
		this.pubmsg(&evt)
	}, nil)

	t.CallbackFriendConnectionStatusAdd(func(_ *tox.Tox, friendNumber uint32, status int, userData interface{}) {
		evt := &thspbs.Event{}
		evt.EventName = "FriendConnectionStatus"
		evt.Args = []string{fmt.Sprintf("%d", friendNumber), fmt.Sprintf("%d", status)}
		pubkey, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		fname, err := t.FriendGetName(friendNumber)
		gopp.ErrPrint(err)
		evt.Margs = []string{fname, pubkey, tox.ConnStatusString(status)}
		this.pubmsg(evt)

		// offline message
		if status > 0 {
			time.AfterFunc(20*time.Millisecond, func() {
				ofm := OffMsgMan()
				offmsgs := ofm.GetByPubkey(pubkey)
				if len(offmsgs) == 0 {
					return
				}
				sentCount := 0
				for idx, offmsg := range offmsgs {
					time.Sleep(time.Duration(idx+1) * 345 * time.Millisecond)
					_, err := t.FriendSendMessage(friendNumber, offmsg.Content)
					gopp.ErrPrint(err)
					if err == nil {
						sentCount += 1
						ofm.DeleteMessage(pubkey, offmsg.Id)
						appctx.st.SetMessageSent(offmsg.Id)
						// construct a FriendSendMessageResp
						evt := &thspbs.Event{}
						evt.EventName = "FriendSendMessageResp"
						evt.Args = gopp.ToStrs(friendNumber)
						evt.Margs = gopp.ToStrs(0, 1, pubkey)
						evt.EventId, evt.UserCode = offmsg.EventId, offmsg.UserCode
						this.pubmsg(evt)
					}
				}
				log.Printf("Send unsents: %d/%d to %s\n", sentCount, len(offmsgs), fname)
			})
		}
	}, nil)

	t.CallbackFriendNameAdd(func(_ *tox.Tox, friendNumber uint32, fname string, userData interface{}) {
		evt := &thspbs.Event{}
		evt.EventName = "FriendName"
		evt.Args = gopp.ToStrs(friendNumber, fname)
		pubkey, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		evt.Margs = []string{pubkey}
		this.pubmsg(evt)
	}, nil)

	t.CallbackFriendStatusMessageAdd(func(_ *tox.Tox, friendNumber uint32, statusText string, userData interface{}) {
		evt := &thspbs.Event{}
		evt.EventName = "FriendStatusMessage"
		evt.Args = gopp.ToStrs(friendNumber, statusText)
		pubkey, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		fname, err := t.FriendGetName(friendNumber)
		gopp.ErrPrint(err)
		evt.Margs = []string{fname, pubkey}
		this.pubmsg(evt)
	}, nil)

	t.CallbackFriendStatusAdd(func(_ *tox.Tox, friendNumber uint32, status int, ud interface{}) {
		evt := &thspbs.Event{}
		evt.EventName = "FriendStatus"
		evt.Args = gopp.ToStrs(friendNumber, status)
		pubkey, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		fname, err := t.FriendGetName(friendNumber)
		gopp.ErrPrint(err)
		evt.Margs = []string{fname, pubkey}
		this.pubmsg(evt)
	}, nil)
	/*
	   type cb_friend_request_ftype func(this *Tox, pubkey string, message string, userData interface{})
	   type cb_friend_message_ftype func(this *Tox, friendNumber uint32, message string, userData interface{})
	   type cb_friend_name_ftype func(this *Tox, friendNumber uint32, newName string, userData interface{})
	   type cb_friend_status_message_ftype func(this *Tox, friendNumber uint32, newStatus string, userData interface{})
	   type cb_friend_status_ftype func(this *Tox, friendNumber uint32, status int, userData interface{})
	   type cb_friend_connection_status_ftype func(this *Tox, friendNumber uint32, status int, userData interface{})
	   type cb_friend_typing_ftype func(this *Tox, friendNumber uint32, isTyping uint8, userData interface{})
	   type cb_friend_read_receipt_ftype func(this *Tox, friendNumber uint32, receipt uint32, userData interface{})
	   type cb_friend_lossy_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})
	   type cb_friend_lossless_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})

	   // self callback type
	   type cb_self_connection_status_ftype func(this *Tox, status int, userData interface{})

	   // file callback type
	   type cb_file_recv_control_ftype func(this *Tox, friendNumber uint32, fileNumber uint32,
	   	control int, userData interface{})
	   type cb_file_recv_ftype func(this *Tox, friendNumber uint32, fileNumber uint32, kind uint32, fileSize uint64,
	   	fileName string, userData interface{})
	   type cb_file_recv_chunk_ftype func(this *Tox, friendNumber uint32, fileNumber uint32, position uint64,
	   	data []byte, userData interface{})
	   type cb_file_chunk_request_ftype func(this *Tox, friend_number uint32, file_number uint32, position uint64,
	   	length int, user_data interface{})

	*/

	t.CallbackConferenceInviteAdd(func(_ *tox.Tox, friendNumber uint32, itype uint8, cookie string, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "ConferenceInvite"
		evt.Args = []string{fmt.Sprintf("%d", friendNumber), fmt.Sprintf("%d", itype), cookie}

		pubkey, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		fname, err := t.FriendGetName(friendNumber)
		gopp.ErrPrint(err)
		evt.Margs = []string{fname, pubkey}

		gn, err := xtox.ConferenceJoin(t, friendNumber, itype, cookie, this.onGroupAudioFrame)
		gopp.ErrPrint(err, fname, gn, len(cookie))
		if err != nil {
			if false {
				time.Sleep(300 * time.Millisecond)
			}
			// join group chat failed: 4
			if gn_, found := xtox.ConferenceGetByCookie(t, cookie); found {
				gn = gn_
				log.Println("Maybe already joined, correct to:", gn)
			} else {
				log.Println("why not found:", cookie)
			}
		}
		evt.Margs = append(evt.Margs, fmt.Sprintf("%d", gn))

		cookie2, _ := xtox.ConferenceGetIdentifier_s(t, gn)
		log.Println(cookie2 == cookie, cookie, cookie2)
		gopp.Assert(!xtox.ConferenceIdIsEmpty(cookie2), fname)

		ctid, err := appctx.st.AddGroup(cookie2, gn, "")
		gopp.ErrPrint(err)
		evt.Margs = append(evt.Margs, fmt.Sprintf("%d", ctid))
		// TODO if title not none, append to Margs

		this.pubmsg(&evt)
	}, nil)

	t.CallbackConferenceMessageAdd(func(_ *tox.Tox, groupNumber uint32, peerNumber uint32, message string, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "ConferenceMessage"
		evt.Args = []string{fmt.Sprintf("%d", groupNumber), fmt.Sprintf("%d", peerNumber), "0", message}

		peerPubkey, err := t.ConferencePeerGetPublicKey(groupNumber, peerNumber)
		gopp.ErrPrint(err)
		peerName, err := t.ConferencePeerGetName(groupNumber, peerNumber)
		gopp.ErrPrint(err)

		title, err := t.ConferenceGetTitle(groupNumber)
		gopp.ErrPrint(err)

		groupId, _ := xtox.ConferenceGetIdentifier(t, groupNumber)
		if xtox.ConferenceIdIsEmpty(groupId) {
			groupId, _ = t.ConferenceGetIdentifier(groupNumber)
		}

		isme := t.SelfGetPublicKey() == peerPubkey
		if isme {
			// dont store, dont dispatch. because already processed by rpc handler.
			return
		}

		// here only process peer that not me's message
		msgo, err := appctx.st.AddGroupMessage(message, "0", groupId, peerPubkey, 0, 0)
		gopp.ErrPrint(err)
		if err == nil {
			err := appctx.st.SetMessageSent(msgo.Id)
			gopp.ErrPrint(err, msgo.Id)
			msgo.Sent = 1
		}
		evt.EventId = msgo.EventId

		evt.Margs = []string{peerName, peerPubkey, title, groupId, gopp.ToStr(msgo.EventId), "1",
			thscom.MSGTYPE_TEXT, "text/plain"}

		if isme {
			// dont dispatch. because already processed by rpc handler.
		} else {
			this.pubmsg(&evt)
		}

		// if url, try to see if it image or file.
		// and then try download to local and extract file info, and update client
		go detectGroupMessageMime(&evt, message)
	}, nil)

	t.CallbackConferenceActionAdd(func(_ *tox.Tox, groupNumber uint32, peerNumber uint32, message string, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "ConferenceMessage"
		evt.Args = []string{fmt.Sprintf("%d", groupNumber), fmt.Sprintf("%d", peerNumber), "1", message}

		peerPubkey, err := t.ConferencePeerGetPublicKey(groupNumber, peerNumber)
		gopp.ErrPrint(err)
		peerName, err := t.ConferencePeerGetName(groupNumber, peerNumber)
		gopp.ErrPrint(err)

		title, err := t.ConferenceGetTitle(groupNumber)
		gopp.ErrPrint(err)

		groupId, _ := xtox.ConferenceGetIdentifier(t, groupNumber)
		if xtox.ConferenceIdIsEmpty(groupId) {
			groupId, _ = t.ConferenceGetIdentifier(groupNumber)
		}

		// here only process peer that not me's message
		msgo, err := appctx.st.AddGroupMessage(message, "0", groupId, peerPubkey, 0, 0)
		gopp.ErrPrint(err)
		if err == nil {
			err := appctx.st.SetMessageSent(msgo.Id)
			gopp.ErrPrint(err, msgo.Id)
			msgo.Sent = 1
		}
		evt.EventId = msgo.EventId

		evt.Margs = []string{peerName, peerPubkey, title, groupId, gopp.ToStr(msgo.EventId), "1",
			thscom.MSGTYPE_TEXT, "text/plain"}

		this.pubmsg(&evt)
	}, nil)

	t.CallbackConferencePeerNameAdd(func(_ *tox.Tox, groupNumber uint32, peerNumber uint32, name string, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "ConferencePeerName"
		evt.Args = []string{fmt.Sprintf("%d", groupNumber), fmt.Sprintf("%d", peerNumber), name}

		peerPubkey, err := t.ConferencePeerGetPublicKey(groupNumber, peerNumber)
		gopp.ErrPrint(err)
		peerName, err := t.ConferencePeerGetName(groupNumber, peerNumber)
		gopp.ErrPrint(err)

		title, err := t.ConferenceGetTitle(groupNumber)
		gopp.ErrPrint(err)

		groupId, _ := xtox.ConferenceGetIdentifier(t, groupNumber)
		if xtox.ConferenceIdIsEmpty(groupId) {
			groupId, _ = t.ConferenceGetIdentifier(groupNumber)
		}

		oldPeerName := this.groupPeerNames[peerPubkey]
		this.groupPeerNames[peerPubkey] = peerName
		ctid, err := appctx.st.AddPeerOrUpdateName(peerPubkey, groupNumber, peerName)
		gopp.ErrPrint(err, peerNumber, name, peerPubkey)

		evt.Margs = []string{peerName, peerPubkey, title, groupId, fmt.Sprintf("%d", ctid), oldPeerName}
		this.pubmsg(&evt)
	}, nil)
	// detect which peer added/deleted here and directly tell client, make client lighter.
	t.CallbackConferencePeerListChangedAdd(func(_ *tox.Tox, groupNumber uint32, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "ConferencePeerListChange"
		evt.Args = []string{fmt.Sprintf("%d", groupNumber)}

		title, err := t.ConferenceGetTitle(groupNumber)
		gopp.ErrPrint(err)

		groupId, _ := xtox.ConferenceGetIdentifier(t, groupNumber)
		if xtox.ConferenceIdIsEmpty(groupId) {
			groupId, _ = t.ConferenceGetIdentifier(groupNumber)
		}

		oldPeerPubkeys := this.groupPeerPubkeys[groupId]
		newPeerPubkeys := t.ConferenceGetPeerPubkeys(groupNumber)
		added, deleted := DiffSliceAsString(oldPeerPubkeys, newPeerPubkeys)

		if len(added) > 0 || len(deleted) > 0 { // omit empty event
			this.groupPeerPubkeys[groupId] = newPeerPubkeys

			addedjs, _ := json.Marshal(added)
			deletedjs, _ := json.Marshal(deleted)
			evt.Margs = []string{title, groupId, gopp.ToStr(len(newPeerPubkeys)), string(addedjs), string(deletedjs)}
			this.pubmsg(&evt)
		}
	}, nil)

	/*
		t.CallbackConferenceNameListChangeAdd(func(_ *tox.Tox, groupNumber uint32, peerNumber uint32, change uint8, userData interface{}) {
			evt := thspbs.Event{}
			evt.EventName = "ConferenceNameListChange"
			evt.Args = []string{fmt.Sprintf("%d", groupNumber),
				fmt.Sprintf("%d", peerNumber), fmt.Sprintf("%d", change)}

			peerPubkey, err := t.ConferencePeerGetPublicKey(groupNumber, peerNumber)
			gopp.ErrPrint(err)
			peerName, err := t.ConferencePeerGetName(groupNumber, peerNumber)
			gopp.ErrPrint(err)

			title, err := t.ConferenceGetTitle(groupNumber)
			gopp.ErrPrint(err)

			groupId, _ := xtox.ConferenceGetIdentifier(t, groupNumber)
			if xtox.ConferenceIdIsEmpty(groupId) {
				groupId, _ = t.ConferenceGetIdentifier(groupNumber)
			}

			ctid, err := appctx.st.AddPeer(peerPubkey, groupNumber)
			gopp.ErrPrint(err)
			evt.Margs = []string{peerName, peerPubkey, title, groupId, fmt.Sprintf("%d", ctid)}

			this.pubmsg(&evt)
		}, nil)
	*/

	t.CallbackConferenceTitleAdd(func(_ *tox.Tox, groupNumber uint32, peerNumber uint32, title string, userData interface{}) {
		evt := thspbs.Event{}
		evt.EventName = "ConferenceTitle"
		stmsg := "" // fake group status message
		if pos := strings.Index(title, thscom.GroupTitleSep); pos > 0 {
			stmsg = title[pos+len(thscom.GroupTitleSep):]
			title = title[:pos]
		}
		evt.Args = []string{fmt.Sprintf("%d", groupNumber),
			fmt.Sprintf("%d", peerNumber), title, stmsg}

		groupId, _ := xtox.ConferenceGetIdentifier(t, groupNumber)
		if xtox.ConferenceIdIsEmpty(groupId) {
			groupId, _ = t.ConferenceGetIdentifier(groupNumber)
		}
		peerPubkey, _ := xtox.ConferencePeerGetPubkey(t, groupNumber, peerNumber)
		peerName, _ := xtox.ConferencePeerGetName(t, groupNumber, peerNumber)
		if peerName == "" || peerPubkey == "" {
			log.Println("peer not found:", groupNumber, peerNumber, title, peerName, peerPubkey)
		}

		ctid, err := appctx.st.AddGroup(groupId, groupNumber, title)
		gopp.ErrPrint(err, ctid)
		if err != nil {
			_, err := appctx.st.UpdateGroup(groupId, groupNumber, title)
			gopp.ErrPrint(err)
		}

		evt.Margs = []string{groupId, peerName, peerPubkey}
		this.pubmsg(&evt)
	}, nil)

	/*
		// conference callback type
		type cb_conference_invite_ftype func(this *Tox, friendNumber uint32, itype uint8, data []byte, userData interface{})
		type cb_conference_message_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, message string, userData interface{})

		type cb_conference_action_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, action string, userData interface{})
		type cb_conference_title_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, title string, userData interface{})
		type cb_conference_namelist_change_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, change uint8, userData interface{})
	*/
}

func (this *ToxVM) pubmsg(evt *thspbs.Event) error {
	log.Println("pubing msg...", evt.EventName)
	appctx.brker.Pubmsg(context.Background(), evt)
	return pubmsgall(nil, evt)
}
