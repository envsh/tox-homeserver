package main

import (
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"time"

	thscli "tox-homeserver/client"
	_ "tox-homeserver/client/grpctp"
	"tox-homeserver/thspbs"
)

// should block
func (dm *daemon) initAppBackend() {
	// mech, uiw := uictx.mech, uictx.uiw

	// TODO maybe do not read/write ui in goroutine
	// srvurl := uiw.ComboBox_6.CurrentText()
	srvurl := "127.0.0.1:2080"
	err := thscli.AppConnect(srvurl)
	gopp.ErrPrint(err, srvurl)
	if err != nil {
		// runOnUiThread(func() {
		// 	this.Label_24.SetText("connect error:" + err.Error())
		// 	this.PushButton_7.SetEnabled(true)
		// })
		return
	}
	dm.vtcli = dm.appctx.GetLigTox()
	dm.vtcli.OnNewMsg = func() {
		// mech.Trigger()
		log.Println("newinmsg")
		select {
		case dm.msgnotifych <- true:
		default:
		}
	}

	condWait(50, func() bool { return dm.vtcli.SelfGetAddress() != "" })
	log.Println("My ToxID:", dm.vtcli.SelfGetAddress())
	// runOnUiThread(func() { this.switchUiStack(UIST_CONTACTUI) })

	// uiw.Label_2.SetText(gopp.StrSuf4ui(vtcli.SelfGetName(), thscom.UiNameLen))
	// uiw.Label_2.SetToolTip(vtcli.SelfGetName())
	// runOnUiThread(func() { SetQLabelElideText(uiw.Label_2, vtcli.SelfGetName(), "") })
	stmsg, _ := dm.vtcli.SelfGetStatusMessage()
	log.Println(stmsg)
	// uiw.Label_3.SetText(gopp.StrSuf4ui(stmsg, thscom.UiStmsgLen))
	// uiw.Label_3.SetToolTip(stmsg)
	// runOnUiThread(func() { SetQLabelElideText(uiw.Label_3, stmsg, "") })
	// uiw.ToolButton_17.SetToolTip(vtcli.SelfGetAddress())

	// listw := uiw.ListWidget_2

	// for fn, frnd := range vtcli.Binfo.Friends {
	// 	itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
	// 	_ = itext
	// 	listw.AddItem(itext)
	// 	contactQueue <- frnd
	// 	contactQueue <- nil
	// }

	// for gn, grp := range vtcli.Binfo.Groups {
	// 	itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
	// 	_ = itext
	// 	listw.AddItem(itext)
	// 	contactQueue <- grp
	// 	contactQueue <- nil
	// }

	// uifnQueue <- func() { this.setConnStatus(gopp.IfElse(vtcli.Binfo.ConnStatus > 0, true, false).(bool)) }
	// uifnQueue <- func() { this.setCoreVersion(vtcli.Binfo.ToxVersion) }

	log.Println("get base info done.")
	dm.baseInfoGot = true

	//
	fwdevt2c := func(name string, args ...interface{}) {
		evt := &thspbs.Event{}
		evt.EventName = name
		for _, arg := range args {
			evt.Args = append(evt.Args, fmt.Sprintf("%v", arg))
		}
		bcc, err := json.Marshal(evt)
		gopp.ErrPrint(err, name)
		dispatchEvent2c(string(bcc))
	}
	fwdevt2c("SelfInfo", dm.vtcli.SelfGetAddress(), dm.vtcli.SelfGetName(), stmsg)
	for fn, frnd := range dm.vtcli.Binfo.Friends {
		fwdevt2c("AddFriendItem", fn, frnd.Pubkey, frnd.Name, frnd.Stmsg)
	}
	for gn, grp := range dm.vtcli.Binfo.Groups {
		fwdevt2c("AddGroupItem", gn, grp.GroupId, grp.Title, grp.Stmsg)
	}

	// mech.Trigger()

	// 加载每个房间的最新消息, force schedue, or contact maybe not show in ui
	/*go*/
	func() {
		// this.loginDone(true)
		btime := time.Now()
		log.Println("Waiting contacts show on UI...") // about 31ms with 7 contacts
		// TODO 这种方式不太准确
		// condWait(10, func() bool { return len(contactQueue) == 0 })
		time.Sleep(10 * time.Millisecond)
		log.Println("Show base contacts on UI done.", time.Since(btime))
		// pullAllRoomsLatestMessages()
	}()
	// select {}
}

// should block
func (dm *daemon) pollmsg() {
	for {
		<-dm.msgnotifych

		bcc := dm.vtcli.GetNextBackenEvent()
		if bcc == nil {
			continue
		}

		evto := &thspbs.Event{}
		err := json.Unmarshal(bcc, evto)
		gopp.ErrPrint(err)
		if err == nil {
			dispatchEvent(evto)
			// dispatchEventResp(evto)
		}
		dispatchEvent2c(string(bcc))
	}
}
