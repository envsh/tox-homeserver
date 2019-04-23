package main

import (
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"time"

	thscli "tox-homeserver/client"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"
)

var baseInfoGot bool = false

func condWait(timeoutms int, f func() bool) {
	for {
		time.Sleep(time.Duration(timeoutms) * time.Millisecond)
		if f() {
			break
		}
	}
}

func runbkdproc() {
	appctx.OpenStrorage()
	loadDataForProfileUi()
	initAppBackend()
}

var devInfo *store.Device

func loadDataForProfileUi() {
	st := appctx.GetStorage()
	devInfo = st.GetDevice()

	setting, err := st.GetSetting(store.SK_HOMESERVER_URL)
	gopp.ErrPrint(err)
	if setting != nil {
		// this.ComboBox_6.SetCurrentText(setting.Value)
		log.Println(setting.Value)
	}

	setting, err = st.GetSetting(store.SK_LAST_LOGINED)
	gopp.ErrPrint(err, store.SK_LAST_LOGINED)
	if setting != nil {
		if setting.Value == "true" {
			// uictx.uiw.PushButton_7.SetDisabled(true)
			time.AfterFunc(60*time.Millisecond, func() {
				// runOnUiThread(func() { this.login(false) })
			})
		}
	} else {
		st.SetSetting(store.SK_LAST_LOGINED, "false")
	}
}

// should block
func initAppBackend() {
	// TODO maybe do not read/write ui in goroutine
	srvurl := "s2.natfrp.org:22080"
	err := thscli.AppConnect(srvurl)
	gopp.ErrPrint(err, srvurl)
	if err != nil {
		/*
			runOnUiThread(func() {
				this.Label_24.SetText("connect error:" + err.Error())
				this.PushButton_7.SetEnabled(true)
			})
		*/
		return
	}
	vtcli = appctx.GetLigTox()
	vtcli.OnNewMsg = func() {
		log.Println("new message event")
		// mech.Trigger()
		for {
			bcc := vtcli.GetNextBackenEvent()
			if bcc == nil {
				break
			}

			evto := &thspbs.Event{}
			err := json.Unmarshal(bcc, evto)
			gopp.ErrPrint(err)
			if err == nil {
				dispatchEvent(evto)
				// dispatchEventResp(evto)
			}

			// uictx.chatform.newmsg(which uint32, msg string)
		}

	}

	condWait(50, func() bool { return vtcli.SelfGetAddress() != "" })
	log.Println("My ToxID:", vtcli.SelfGetAddress())
	// runOnUiThread(func() { this.switchUiStack(UIST_CONTACTUI) })
	uictx.minfov.id = vtcli.SelfGetAddress()
	uictx.minfov.name = vtcli.SelfGetName()
	uictx.minfov.sttxt = fmt.Sprintf("%d", vtcli.SelfGetConnectionStatus())
	uictx.minfov.stmsg, _ = vtcli.SelfGetStatusMessage()

	// uiw.Label_2.SetText(gopp.StrSuf4ui(vtcli.SelfGetName(), thscom.UiNameLen))
	// uiw.Label_2.SetToolTip(vtcli.SelfGetName())
	// runOnUiThread(func() { SetQLabelElideText(uiw.Label_2, vtcli.SelfGetName(), "") })
	log.Println("My Name:", vtcli.SelfGetName())
	stmsg, _ := vtcli.SelfGetStatusMessage()
	_ = stmsg
	// uiw.Label_3.SetText(gopp.StrSuf4ui(stmsg, thscom.UiStmsgLen))
	// uiw.Label_3.SetToolTip(stmsg)
	// runOnUiThread(func() { SetQLabelElideText(uiw.Label_3, stmsg, "") })
	// uiw.ToolButton_17.SetToolTip(vtcli.SelfGetAddress())

	uictx.ctview.setFriendInfos(vtcli.Binfo.Friends)
	uictx.ctview.setGroupInfos(vtcli.Binfo.Groups)
	/*
		listw := uiw.ListWidget_2

		for fn, frnd := range vtcli.Binfo.Friends {
			itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
			_ = itext
			listw.AddItem(itext)
			contactQueue <- frnd
			contactQueue <- nil
		}

		for gn, grp := range vtcli.Binfo.Groups {
			itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
			_ = itext
			listw.AddItem(itext)
			contactQueue <- grp
			contactQueue <- nil
		}

		uifnQueue <- func() { this.setConnStatus(gopp.IfElse(vtcli.Binfo.ConnStatus > 0, true, false).(bool)) }
		uifnQueue <- func() { this.setCoreVersion(vtcli.Binfo.ToxVersion) }
	*/
	log.Println("get base info done.")
	baseInfoGot = true

	//	mech.Trigger()

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