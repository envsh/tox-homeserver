package main

import (
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtwidgets"
)

func dispatchEvent(jso *simplejson.Json) {
	uiw, ctitmdl := uictx.uiw, uictx.ctitmdl
	// listwp1 := Ui_MainWindow_Get_listWidget(uiw)
	// listw1 := widgets.NewQListWidgetFromPointer(listwp1)

	evtName := jso.Get("name").MustString()
	switch evtName {
	case "SelfConnectionStatus": // {"name":"SelfConnectionStatus","args":["2"],"margs":["CONNECTION_UDP"]}
		status := gopp.MustUint32(jso.Get("args").GetIndex(0).MustString())
		uictx.mw.setConnStatus(status > 0)
	case "FriendRequest":
		///
		// pubkey := jso.Get("args").GetIndex(0).MustString()
		// _, err := appctx.store.AddFriend(pubkey, 0, "", "")
		// gopp.ErrPrint(err, jso.Get("args"))

	case "FriendMessage":
		// jso.Get("args").GetIndex(0).MustString()
		msg := jso.Get("args").GetIndex(1).MustString()
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _, _ = msg, fname, pubkey

		itext := fmt.Sprintf("%s: %s", fname, msg)
		uiw.ListWidget.AddItem(itext)
		uiw.ListWidget.ScrollToBottom()

		item := uictx.iteman.Get(pubkey)
		if item == nil {
			log.Println("wtf", fname, pubkey, msg)
		} else {
			msgo := NewMessageForFriend(jso)
			item.AddMessage(msgo, false)
		}

		///
		// _, err := appctx.store.AddFriendMessage(msg, pubkey)
		// gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _ = fname, pubkey
		st := gopp.MustInt(jso.Get("args").GetIndex(1).MustString())

		item := uictx.iteman.Get(pubkey)
		if item != nil {
			item.setConnStatus(int32(st))
			if item.GetName() != fname && fname != "" {
				item.UpdateName(fname)
			}
		} else {
			log.Println("item not found:", fname, pubkey)
		}

	case "FriendName":
		fname := jso.Get("args").GetIndex(1).MustString()
		pubkey := jso.Get("margs").GetIndex(0).MustString()
		_, _ = fname, pubkey
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			item.UpdateName(fname)
		} else {
			log.Println("item not found:", fname, pubkey)
		}

	case "FriendStatusMessage":
		statusText := jso.Get("args").GetIndex(1).MustString()
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _ = fname, pubkey
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			item.UpdateStatusMessage(statusText)
		} else {
			log.Println("item not found:", fname, pubkey)
		}

	case "FriendStatus":
		status := gopp.MustInt(jso.Get("args").GetIndex(1).MustString())
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _ = fname, pubkey
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			item.setUserStatus(status)
		} else {
			log.Println("item not found:", fname, pubkey)
		}

	case "ConferenceInvite":
		groupNumber := jso.Get("margs").GetIndex(2).MustString()
		cookie := jso.Get("args").GetIndex(2).MustString()
		groupId := thscli.ConferenceCookieToIdentifier(cookie)
		log.Println(groupId)
		_ = groupNumber

		item := uictx.iteman.Get(groupId)
		if item == nil {
			item = NewRoomListItem()
			item.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
				uictx.mw.onRoomContextMenu(item, w, pos)
			}
			item.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}
			uictx.iteman.addRoomItem(item)
			grpInfo := &thspbs.GroupInfo{}
			grpInfo.GroupId = groupId
			grpInfo.Gnum = gopp.MustUint32(groupNumber)
			grpInfo.Title = fmt.Sprintf("Group #%s", groupNumber)
			item.SetContactInfo(grpInfo)
			log.Println("New group contact item:", groupNumber, grpInfo.Title, groupId)
		} else {
			log.Println("Reuse group contact item:", groupNumber, item.grpInfo.Title, groupId)
			if gopp.MustUint32(groupNumber) != item.grpInfo.Gnum {
				log.Println("GroupNumber changed, update it.", item.grpInfo.Gnum, groupNumber)
				item.grpInfo.Gnum = gopp.MustUint32(groupNumber)
			}
		}

		///
		// _, err := appctx.store.AddGroup(groupId, ctis.cnum, ctis.ctname)
		// gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupNumber := jso.Get("args").GetIndex(1).MustString()
		groupTitle := jso.Get("args").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(0).MustString()
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}
		_ = groupTitle

		item := uictx.iteman.Get(groupId)
		if item != nil {
			item.UpdateName(groupTitle)
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
			item.SetContactInfo(grpInfo)
			log.Println("New group contact item:", groupNumber, groupId, groupTitle)
		}
	case "ConferencePeerName":
		gnum := gopp.MustUint32(jso.Get("args").GetIndex(0).MustString())
		pnum := gopp.MustUint32(jso.Get("args").GetIndex(1).MustString())
		groupId := jso.Get("margs").GetIndex(3).MustString()
		pname := jso.Get("margs").GetIndex(0).MustString()
		ppubkey := jso.Get("margs").GetIndex(1).MustString()
		vtcli.Binfo.UpdatePeerInfo(gnum, groupId, ppubkey, pname, pnum)
		peeros := vtcli.Binfo.GetGroupMembers(gnum)
		item := uictx.iteman.Get(groupId)
		if item != nil {
			if item.peerCount != len(peeros) {
				// item.SetPeerCount(len(peeros))
			}
		}
	case "ConferencePeerListChange":
		groupId := jso.Get("margs").GetIndex(1).MustString()
		peerCount := gopp.MustInt(jso.Get("margs").GetIndex(2).MustString())
		item := uictx.iteman.Get(groupId)
		if item != nil {
			if item.peerCount != peerCount {
				item.SetPeerCount(peerCount)
			}
		}
		// update deleted ones
		gnum := gopp.MustUint32(jso.Get("args").GetIndex(0).MustString())
		deletedPeerPubkeysjs := jso.Get("margs").GetIndex(4).MustString()
		deletedPeerPubkeys := []string{}
		err := json.Unmarshal([]byte(deletedPeerPubkeysjs), deletedPeerPubkeys)
		gopp.ErrPrint(err, deletedPeerPubkeysjs)
		for _, pubkey := range deletedPeerPubkeys {
			vtcli.Binfo.DeletePeerInfo(gnum, groupId, pubkey)
		}
	case "ConferenceNameListChange": // depcreated
		groupTitle := jso.Get("margs").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(3).MustString()
		log.Println(groupId)
		if thscli.ConferenceIdIsEmpty(groupId) {
			log.Println("empty")
			break
		}
		_ = groupTitle

	case "ConferenceMessage":
		groupId := jso.Get("margs").GetIndex(3).MustString()
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}

		message := jso.Get("args").GetIndex(3).MustString()
		peerName := jso.Get("margs").GetIndex(0).MustString()
		groupTitle := jso.Get("margs").GetIndex(2).MustString()

		// raw message show area
		itext := fmt.Sprintf("%s@%s: %s", peerName, groupTitle, message)
		uiw.ListWidget.AddItem(itext)
		uiw.ListWidget.ScrollToBottom()
		// log.Println("item:", itext)

		ccstate.curpos = uiw.ScrollArea_2.VerticalScrollBar().Value()
		ccstate.maxpos = uiw.ScrollArea_2.VerticalScrollBar().Maximum()

		for _, room := range ctitmdl {
			// log.Println(room.GetName(), ",", groupTitle, ",", room.GetId(), ",", groupId)
			if room.GetId() == groupId && room.GetName() == groupTitle {
				room.AddMessage(NewMessageForGroup(jso), false)
				break
			}
		}

	case "FriendSendMessage":
		itext := jso.Get("args").GetIndex(1).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		eventId := gopp.MustInt64(jso.Get("margs").GetIndex(2).MustString())

		found := false
		for _, room := range ctitmdl {
			if room.GetId() == pubkey {
				msgo := NewMessageForMeFromJson(itext, eventId)
				room.AddMessage(msgo, false)
				found = true
				break
			}
		}
		log.Println(found, pubkey, itext)

	case "ConferenceSendMessage":
		itext := jso.Get("args").GetIndex(2).MustString()
		groupTitle := jso.Get("margs").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(3).MustString()
		eventId := gopp.MustInt64(jso.Get("margs").GetIndex(4).MustString())

		found := false
		for _, room := range ctitmdl {
			if room.GetId() == groupId && room.GetName() == groupTitle {
				msgo := NewMessageForMeFromJson(itext, eventId)
				room.AddMessage(msgo, false)
				found = true
				break
			}
		}
		log.Println(found, groupId, itext)

	default:
		log.Println(jso)
	}
}

func dispatchEventResp(jso *simplejson.Json) {
	// uiw, ctitmdl := uictx.uiw, uictx.ctitmdl
	// listwp1 := Ui_MainWindow_Get_listWidget(uiw)
	// listw1 := widgets.NewQListWidgetFromPointer(listwp1)

	evtName := jso.Get("name").MustString()
	switch evtName {
	case "FriendAddResp":
		fnum := gopp.MustUint32(jso.Get("args").GetIndex(0).MustString())
		toxid := jso.Get("margs").GetIndex(0).MustString()
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
		fnum := gopp.MustUint32(jso.Get("args").GetIndex(0).MustString())
		toxid := jso.Get("margs").GetIndex(0).MustString()
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
		pubkey := jso.Get("margs").GetIndex(0).MustString()
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			uictx.iteman.Delete(item)
		}
	case "ConferenceNewResp":
		pubkey := jso.Get("args").GetIndex(0).MustString()
		title := jso.Get("margs").GetIndex(0).MustString()
		gnum := gopp.MustUint32(jso.Get("args").GetIndex(1).MustString())
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
		pubkey := jso.Get("args").GetIndex(0).MustString()
		item := uictx.iteman.Get(pubkey)
		if item != nil {
			uictx.iteman.Delete(item)
		}
	case "SelfSetNameResp":
		name := jso.Get("args").GetIndex(0).MustString()
		uictx.uiw.Label_2.SetText(name)
	case "SelfSetStatusMessageResp":
		stmsg := jso.Get("args").GetIndex(0).MustString()
		uictx.uiw.Label_3.SetText(stmsg)
	default:
		log.Println(jso)
	}
}
