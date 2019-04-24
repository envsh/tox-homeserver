package main

import (
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"strings"

	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"
)

func dispatchEvent(evto *thspbs.Event) {

	switch evto.Name {
	case "SelfConnectionStatus": // {"Name":"SelfConnectionStatus","Args":["2"],"Margs":["CONNECTION_UDP"]}
		status := gopp.MustUint32(evto.Args[0])
		// uictx.mw.setConnStatus(status > 0)
		uictx.minfov.sttxt = fmt.Sprintf("%d", status)
	case "FriendRequest":
		///
		// pubkey := jso.Get("Args").GetIndex(0).MustString()
		// _, err := appctx.store.AddFriend(pubkey, 0, "", "")
		// gopp.ErrPrint(err, jso.Get("Args"))

	case "FriendMessage":
		// jso.Get("Args").GetIndex(0).MustString()
		msg := evto.Args[1]
		fname := evto.Margs[0]
		pubkey := evto.Margs[1]
		_, _, _ = msg, fname, pubkey
		uictx.chatform.newmsg(pubkey, msg)

		/*
			itext := fmt.Sprintf("%s: %s", fname, msg)
			uiw.ListWidget.AddItem(itext)
			uiw.ListWidget.ScrollToBottom()

			item := uictx.iteman.Get(pubkey)
			if item == nil {
				log.Println("wtf friend item not found:", fname, pubkey, msg)
			} else {
				msgty := evto.Margs[4]
				msgo := NewMessageForFriend(evto)
				if msgty == thscom.MSGTYPE_AVATAR {
					item.SetAvatarForMessage(msgo, pubkey)
				} else {
					item.AddMessage(msgo, false)
				}
			}
		*/

	case "FriendConnectionStatus":
		fname := evto.Margs[0]
		pubkey := evto.Margs[1]
		_, _ = fname, pubkey
		st := gopp.MustInt(evto.Args[1])
		_ = st
		/*
			item := uictx.iteman.Get(pubkey)
			if item != nil {
				item.setConnStatus(int32(st))
				if item.GetName() != fname && fname != "" {
					item.UpdateName(fname)
				}
			} else {
				log.Println("item not found:", fname, pubkey)
			}
		*/

	case "FriendName":
		fname := evto.Args[1]
		pubkey := evto.Margs[0]
		_, _ = fname, pubkey
		/*
			item := uictx.iteman.Get(pubkey)
			if item != nil {
				item.UpdateName(fname)
			} else {
				log.Println("item not found:", fname, pubkey)
			}
		*/

	case "FriendStatusMessage":
		statusText := evto.Args[1]
		fname := evto.Margs[0]
		pubkey := evto.Margs[1]
		_, _, _ = fname, pubkey, statusText
		/*
			item := uictx.iteman.Get(pubkey)
			if item != nil {
				item.UpdateStatusMessage(statusText)
			} else {
				log.Println("item not found:", fname, pubkey)
			}
		*/

	case "FriendStatus":
		status := gopp.MustInt(evto.Args[1])
		fname := evto.Margs[0]
		pubkey := evto.Margs[1]
		_, _, _ = fname, pubkey, status
		/*
			item := uictx.iteman.Get(pubkey)
			if item != nil {
				item.setUserStatus(status)
			} else {
				log.Println("item not found:", fname, pubkey)
			}
		*/

	case "ConferenceInvite":
		groupNumber := evto.Margs[2]
		cookie := evto.Args[2]
		groupId := thscli.ConferenceCookieToIdentifier(cookie)
		log.Println(groupId)
		_ = groupNumber

		/*
			item := uictx.iteman.Get(groupId)
			if item == nil {
				item = NewRoomListItem()
				item.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
					uictx.mw.onRoomContextMenu(item, w, pos)
				}
				item.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}
				uictx.iteman.addRoomItem(item)
				gopp.Assert(!vtcli.Binfo.HasGroup(gopp.MustUint32(groupNumber)),
					"group already exist in BaseInfo", groupNumber)
				vtcli.Binfo.AddGroup(gopp.MustUint32(groupNumber), groupId)
				grpInfo := vtcli.Binfo.GetGroup(gopp.MustUint32(groupNumber))
				grpInfo.Title = fmt.Sprintf("Group #%s.%s", groupNumber, groupId[:5])
				item.SetContactInfo(grpInfo)
				log.Println("New group contact item:", groupNumber, grpInfo.Title, groupId)
				grpInfo.AddPeerInfo(evto.Margs[1], evto.Margs[0], gopp.MustUint32(evto.Args[1]))
			} else {
				log.Println("Reuse group contact item:", groupNumber, item.grpInfo.Title, groupId)
				if gopp.MustUint32(groupNumber) != item.grpInfo.Gnum {
					log.Println("GroupNumber changed, update it.", item.grpInfo.Gnum, groupNumber)
					item.grpInfo.Gnum = gopp.MustUint32(groupNumber)
				}
			}
		*/

		///
		// _, err := appctx.store.AddGroup(groupId, ctis.cnum, ctis.ctname)
		// gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupNumber := evto.Args[1]
		groupTitle := evto.Args[2]
		groupStmsg := evto.Args[3]
		groupId := evto.Margs[0]
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}
		_ = groupTitle
		_ = groupStmsg
		_ = groupNumber

		/*
			item := uictx.iteman.Get(groupId)
			if item != nil {
				item.UpdateName(groupTitle)
				item.UpdateStatusMessage(groupStmsg)
				log.Println("Reuse item and update group contact title:", groupNumber, groupId, groupTitle)
			} else {
				item = NewRoomListItem()
				item.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
					uictx.mw.onRoomContextMenu(item, w, pos)
				}
				item.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}
				uictx.iteman.addRoomItem(item)
				grpInfo := &thspbs.GroupInfo{}
				grpInfo.GroupId = groupId
				grpInfo.Gnum = gopp.MustUint32(groupNumber)
				grpInfo.Title = groupTitle
				grpInfo.Stmsg = groupStmsg
				item.SetContactInfo(grpInfo)
				log.Println("New group contact item:", groupNumber, groupId, groupTitle)
			}
		*/
	case "ConferencePeerName":
		gnum := gopp.MustUint32(evto.Args[0])
		pnum := gopp.MustUint32(evto.Args[1])
		groupId := evto.Margs[3]
		peerName := evto.Margs[0]
		peerPubkey := evto.Margs[1]
		oldPeerCount := len(vtcli.Binfo.GetGroupMembers(gnum))
		vtcli.Binfo.UpdatePeerInfo(gnum, groupId, peerPubkey, peerName, pnum)
		_ = oldPeerCount
		/*
			item := uictx.iteman.Get(groupId)
			if item != nil {
				newPeerCount := len(item.grpInfo.Members)
				if newPeerCount == oldPeerCount+1 {
					item.SetPeerCount(newPeerCount)
				}
			}
		*/
	case "ConferencePeerListChange":
		groupId := evto.Margs[1]
		/*
			item := uictx.iteman.Get(groupId)
			if item != nil {
			}
		*/
		// update deleted ones
		gnum := gopp.MustUint32(evto.Args[0])
		deletedPeerPubkeysjs := evto.Margs[4]
		deletedPeerPubkeys := []string{}
		err := json.Unmarshal([]byte(deletedPeerPubkeysjs), &deletedPeerPubkeys)
		gopp.ErrPrint(err, deletedPeerPubkeysjs)
		for _, pubkey := range deletedPeerPubkeys {
			vtcli.Binfo.DeletePeerInfo(gnum, groupId, pubkey)
		}
		addedPeerPubkeysjs := evto.Margs[3]
		addedPeerPubkeys := []string{}
		err = json.Unmarshal([]byte(addedPeerPubkeysjs), &addedPeerPubkeys)
		gopp.ErrPrint(err, addedPeerPubkeysjs)
		for _, pubkey := range addedPeerPubkeys { // not know peer num, so can not update peer info here
			// vtcli.Binfo.UpdatePeerInfo(gnum, groupId, pubkey, "", 0)
			_ = pubkey
		}
		if len(deletedPeerPubkeys) > 0 || len(addedPeerPubkeys) > 0 {
			// item.SetPeerCount(len(item.grpInfo.Members))
		}
	case "ConferenceNameListChange": // depcreated
		groupTitle := evto.Margs[2]
		groupId := evto.Margs[3]
		log.Println(groupId)
		if thscli.ConferenceIdIsEmpty(groupId) {
			log.Println("empty")
			break
		}
		_ = groupTitle

	case "ConferenceMessage":
		groupId := evto.Margs[3]
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}

		message := evto.Args[3]
		peerName := evto.Margs[0]
		groupTitle := evto.Margs[2]
		_, _, _ = message, peerName, groupTitle
		uictx.chatform.newmsg(groupId, peerName+": "+message)

		// raw message show area
		// itext := fmt.Sprintf("%s@%s: %s", peerName, groupTitle, message)
		// uiw.ListWidget.AddItem(itext)
		// uiw.ListWidget.ScrollToBottom()
		// log.Println("item:", itext)

		// ccstate.curpos = uiw.ScrollArea_2.VerticalScrollBar().Value()
		// ccstate.maxpos = uiw.ScrollArea_2.VerticalScrollBar().Maximum()

		// for _, room := range ctitmdl {
		// log.Println(room.GetName(), ",", groupTitle, ",", room.GetId(), ",", groupId)
		//	if room.GetId() == groupId /* && room.GetName() == groupTitle*/ {
		// room.AddMessage(NewMessageForGroup(evto), false)
		//		break
		//	}
		// }

	case "FriendSendMessage":
		itext := evto.Args[1]
		pubkey := evto.Margs[1]
		eventId := gopp.MustInt64(evto.Margs[2])
		_, _, _ = itext, pubkey, eventId

		/*
			found := false
			for _, room := range ctitmdl {
				if room.GetId() == pubkey {
					msgo := NewMessageForMeFromJson(itext, eventId)
					msgo.UserCode = evto.UserCode
					// 有可能是自己同步回来的消息，所以要确定是添加还是更新
					if room.FindMessageByUserCode(evto.UserCode) != nil {
						room.UpdateMessageState(msgo)
					} else {
						room.AddMessage(msgo, false)
					}
					found = true
					break
				}
			}
		*/
		// log.Println(found, gopp.IfElseStr(found, "", "not"), "found", pubkey, itext)

	case "ConferenceSendMessage":
		itext := evto.Args[2]
		// groupTitle := evto.Margs[2]
		groupId := evto.Margs[3]
		eventId := gopp.MustInt64(evto.Margs[4])
		gopp.G_USED(itext, groupId, eventId)

		/*
				found := false
				for _, room := range ctitmdl {
					if room.GetId() == groupId {
						msgo := NewMessageForMeFromJson(itext, eventId)
						msgo.UserCode = evto.UserCode
						log.Println(msgo)
						// 有可能是自己同步回来的消息，所以要确定是添加还是更新
						if room.FindMessageByUserCode(evto.UserCode) != nil {
							log.Println("Updateing...")
							room.UpdateMessageState(msgo)
						} else {
							log.Println("Adding...")
							room.AddMessage(msgo, false)
						}
						found = true
						break
					}
			}
		*/
		//log.Println(found, groupId, itext)

	case "Call": // always from friend
		/*
			avm := AVMan()
			uargs := evto.Uargs
			onNewAudioFrame := func(aframe []byte, sampleCount uint32, channels uint8, samplingRate uint32) {
				// send aframe to friend
				_ = uargs
				vtcli.AudioSendFrame(uargs.FriendNumber, aframe, sampleCount, channels, samplingRate)
			}
			onNewVideoFrame := func(vframe []byte, width, height uint16) {
				vtcli.VideoSendFrame(uargs.FriendNumber, vframe, width, height)
			}
			avm.NewSession(uargs.FriendPubkey, uargs.AudioEnabled == 1, uargs.VideoEnabled == 1,
				onNewAudioFrame, onNewVideoFrame)
		*/
	case "CallState":
		// log.Println(evto.Name, evto.Uargs.CallState, CallStateString(evto.Uargs.CallState))
		/*
			uargs := evto.Uargs
			if uargs.CallState == 2  || uargs.CallState == 1  {
				AVMan().RemoveSession(evto.Uargs.FriendPubkey, evto.Uargs.FriendName)
			}
		*/
	case "AudioReceiveFrame":
		/*
			pcm := evto.Uargs.Pcm
			evto.Uargs.Pcm = nil
			if false {
				log.Println(evto.Uargs, len(pcm))
			}
			AVMan().PutAudioFrame(evto.Uargs.FriendPubkey, pcm)
		*/
	case "VideoReceiveFrame":
		/*
			vdfrm := evto.Uargs.VideoFrame
			evto.Uargs.VideoFrame = nil
			if false {
				log.Println(evto.Uargs, len(vdfrm))
			}
			AVMan().PutVideoFrame(evto.Uargs.FriendPubkey, vdfrm, int(evto.Uargs.Width), int(evto.Uargs.Height))
		*/
	case "ConferenceAudioReceiveFrame":
		/*
			uargs := evto.Uargs
			pcm := uargs.Pcm
			uargs.Pcm = nil
			if !AVMan().HasSession(uargs.GroupIdentity) {
				// delete session when delete group
				// TODO or maybe create when join
				AVMan().NewSession(uargs.GroupIdentity, true, false,
					func(aframe []byte, sampleCount uint32, channels uint8, samplingRate uint32) {
						vtcli.GroupSendAudio(uargs.GroupNumber, aframe, uint(sampleCount), channels, samplingRate)
					}, nil)
			}
			AVMan().PutAudioFrame(uargs.GroupIdentity, pcm)
		*/
	default:
		if !strings.HasSuffix(evto.Name, "Resp") &&
			!strings.HasSuffix(evto.Name, "Reload") {
			log.Printf("Unimpled: %+v\n", evto)
		}
	}
}

/*
func dispatchEventResp(evto *thspbs.Event) {
	// uiw, ctitmdl := uictx.uiw, uictx.ctitmdl
	// listwp1 := Ui_MainWindow_Get_listWidget(uiw)
	// listw1 := widgets.NewQListWidgetFromPointer(listwp1)

	switch evto.Name {
	case "FriendAddResp":
		fnum := gopp.MustUint32(evto.Args[0])
		toxid := evto.Margs[0]
		pubkey := toxid[:64]
		item := uictx.iteman.Get(pubkey)
		if item == nil {
			frndo := &thspbs.FriendInfo{}
			frndo.Fnum = fnum
			frndo.Pubkey = pubkey
			frndo.Name = pubkey
			contactQueue <- frndo
			uictx.mech.Trigger()
		}
	case "FriendAddNorequestResp":
		fnum := gopp.MustUint32(evto.Args[0])
		toxid := evto.Margs[0]
		pubkey := toxid[:64]
		item := uictx.iteman.Get(pubkey)
		if item == nil {
			frndo := &thspbs.FriendInfo{}
			frndo.Fnum = fnum
			frndo.Pubkey = pubkey
			frndo.Name = pubkey
			contactQueue <- frndo
			uictx.mech.Trigger()
		}
	case "FriendDeleteResp":
		pubkey := evto.Margs[0]
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			uictx.iteman.Delete(item)
		}
	case "ConferenceNewResp":
		pubkey := evto.Args[0]
		title := evto.Margs[0]
		gnum := gopp.MustUint32(evto.Args[1])
		item := uictx.iteman.Get(pubkey)
		if item == nil {
			grpo := &thspbs.GroupInfo{}
			grpo.Gnum = gnum
			grpo.GroupId = pubkey
			grpo.Title = title
			contactQueue <- grpo
			uictx.mech.Trigger()
		}
	case "ConferenceDeleteResp":
		pubkey := evto.Args[0]
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			uictx.iteman.Delete(item)
		}
	case "SelfSetNameResp":
		name := evto.Args[0]
		uictx.uiw.Label_2.SetText(name)
	case "SelfSetStatusMessageResp":
		stmsg := evto.Args[0]
		uictx.uiw.Label_3.SetText(stmsg)
	case "FriendSendMessageResp":
		sent := evto.Margs[1] == "1"
		roompk := evto.Margs[2]

		roomo := uictx.iteman.Get(roompk)
		if roomo != nil {
			msgo := NewMessageForMeFromJson("", evto.EventId)
			msgo.Sent = sent
			msgo.UserCode = evto.UserCode
			roomo.UpdateMessageState(msgo) // 必定已经存在
		}
	case "ConferenceSendMessageResp":
		sent := evto.Margs[1] == "1"
		roompk := evto.Margs[2]

		roomo := uictx.iteman.Get(roompk)
		if roomo != nil {
			msgo := NewMessageForMeFromJson("", evto.EventId)
			msgo.Sent = sent
			msgo.UserCode = evto.UserCode
			roomo.UpdateMessageState(msgo) // 必定已经存在
		} else {
			log.Println("room not found:", sent, roompk, evto)
		}
	case "ConferenceMessageReload":
		groupId := evto.Margs[3]
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}

		// message := evto.Args[3]
		// peerName := evto.Margs[0]
		// groupTitle := evto.Margs[2]

		room := uictx.iteman.Get(groupId)
		if room != nil {
			msgo := NewMessageForGroup(evto)
			msgiw := room.FindMessageViewByEventId(msgo.EventId)
			gopp.NilPrint(msgiw, "View not found:", msgo.EventId, msgo.UserCode)
			room.UpdateMessageMimeContent(msgo, msgiw)
		}
	case "AudioReceiveFrame": // do nothing
	case "VideoReceiveFrame": // do nothing
	default:
		if strings.HasSuffix(evto.Name, "Resp") ||
			strings.HasSuffix(evto.Name, "Reload") {
			log.Printf("Unimpled: %+v\n", evto)
		}
	}
}
*/

// intent
/*
func dispatchOtherEvent(evto *thspbs.Event) {
	log.Println(evto)

	switch evto.Name {
	case "IntentMessage":
		mtype := evto.Args[0]
		mcontent := evto.Args[1]

		item := uictx.iteman.Get(thscom.FileHelperPk)
		gopp.NilPrint(item, "Why FileHelper item nil?")
		uictx.msgwin.SetRoom(item)
		uictx.mw.switchUiStack(UIST_MESSAGEUI)
		uictx.mw.sendMessageImpl(item, mtype+":"+mcontent, false, thscom.FileHelperFnum)
	default:
		// log.Printf("Unimpled: %+v\n", evto)
	}
}
*/
