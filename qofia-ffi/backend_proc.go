package main

import (
	"fmt"
	"gopp"
	"log"
	"time"
	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
)

// should block
func (this *MainWindow) initAppBackend() {
	mech, uiw := uictx.mech, uictx.uiw

	// TODO maybe do not read/write ui in goroutine
	srvurl := uiw.ComboBox_6.CurrentText()
	err := thscli.AppConnect(srvurl)
	gopp.ErrPrint(err, srvurl)
	if err != nil {
		runOnUiThread(func() { this.Label_24.SetText("connect error:" + err.Error()) })
		return
	}
	vtcli = appctx.GetLigTox()
	vtcli.OnNewMsg = func() { mech.Trigger() }

	condWait(50, func() bool { return vtcli.SelfGetAddress() != "" })
	log.Println("My ToxID:", vtcli.SelfGetAddress())
	runOnUiThread(func() { this.switchUiStack(UIST_CONTACTUI) })

	uiw.Label_2.SetText(gopp.StrSuf4ui(vtcli.SelfGetName(), thscom.UiNameLen))
	uiw.Label_2.SetToolTip(vtcli.SelfGetName())
	runOnUiThread(func() { SetQLabelElideText(uiw.Label_2, vtcli.SelfGetName(), "") })
	stmsg, _ := vtcli.SelfGetStatusMessage()
	uiw.Label_3.SetText(gopp.StrSuf4ui(stmsg, thscom.UiStmsgLen))
	uiw.Label_3.SetToolTip(stmsg)
	runOnUiThread(func() { SetQLabelElideText(uiw.Label_3, stmsg, "") })
	uiw.ToolButton_17.SetToolTip(vtcli.SelfGetAddress())

	listw := uiw.ListWidget_2

	for fn, frnd := range vtcli.Binfo.Friends {
		itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
		_ = itext
		listw.AddItem(itext)
		contactQueue <- frnd
	}

	for gn, grp := range vtcli.Binfo.Groups {
		itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
		_ = itext
		listw.AddItem(itext)
		contactQueue <- grp
	}

	uifnQueue <- func() { this.setConnStatus(gopp.IfElse(vtcli.Binfo.ConnStatus > 0, true, false).(bool)) }
	uifnQueue <- func() { this.setCoreVersion(vtcli.Binfo.ToxVersion) }

	log.Println("get base info done.")
	baseInfoGot = true

	mech.Trigger()

	// 加载每个房间的最新消息, force schedue, or contact maybe not show in ui
	go func() {
		btime := time.Now()
		log.Println("Waiting contacts show on UI...") // about 31ms with 7 contacts
		condWait(10, func() bool { return len(contactQueue) == 0 })
		log.Println("Show base contacts on UI done.", time.Since(btime))
		pullAllRoomsLatestMessages()
	}()
	select {}
}
