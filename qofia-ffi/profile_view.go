package main

// profile and login ui

import (
	"fmt"
	"gopp"
	"log"
	"time"

	thscli "tox-homeserver/client"
	"tox-homeserver/store"

	"github.com/kitech/qt.go/qtrt"
)

func (this *MainWindow) initProfile() {
	this.initProfileUi()
	this.initProfileSignals()
	this.initProfileEvents()
}

func (this *MainWindow) initProfileUi() {
}

func (this *MainWindow) initProfileSignals() {
	uiw := uictx.uiw
	qtrt.Connect(uiw.PushButton_7, "clicked(bool)", this.login)
	qtrt.Connect(uiw.ToolButton_32, "clicked(bool)", this.logout)

	// test
}

func (this *MainWindow) initProfileEvents() {
}

func (this *MainWindow) initProfileStorage() {
	this.loadDataForProfileUi()
}

func (this *MainWindow) loadStorage() {
	appctx = thscli.NewAppContext()
	appctx.OpenStrorage()
	log.Println("Open storage done")
}

func (this *MainWindow) login(bool) {
	this.Label_24.SetText("...")
	this.saveLastLoginHost()
	go this.initAppBackend()
}
func (this *MainWindow) loginDone(ok bool) {
	st := appctx.GetStorage()
	st.SetSetting(store.SK_LAST_LOGINED, fmt.Sprintf("%v", ok))
	runOnUiThread(func() { uictx.uiw.PushButton_7.SetEnabled(true) })
}
func (this *MainWindow) logout(bool) {
	runOnUiThread(func() {
		uictx.uiw.PushButton_7.SetEnabled(true)
		uictx.uiw.ToolButton_32.SetEnabled(false)
	})
	if baseInfoGot == false || appctx == nil || vtcli == nil {
		// return
	}
	st := appctx.GetStorage()
	st.SetSetting(store.SK_LAST_LOGINED, fmt.Sprintf("%v", false))
	runOnUiThread(func() {
		// TODO cleanup...
		baseInfoGot = false
		// client close
		thscli.AppDestroy()
		vtcli = nil
		appctx = nil

		// ui clear
		uictx.iteman.ClearRoomList()

		// switch login page
	})
}

var devInfo *store.Device

func (this *MainWindow) loadDataForProfileUi() {
	st := appctx.GetStorage()
	devInfo = st.GetDevice()

	setting, err := st.GetSetting(store.SK_HOMESERVER_URL)
	gopp.ErrPrint(err)
	if setting != nil {
		this.ComboBox_6.SetCurrentText(setting.Value)
	}

	setting, err = st.GetSetting(store.SK_LAST_LOGINED)
	gopp.ErrPrint(err, store.SK_LAST_LOGINED)
	if setting != nil {
		if setting.Value == "true" {
			uictx.uiw.PushButton_7.SetDisabled(true)
			time.AfterFunc(60*time.Millisecond, func() {
				runOnUiThread(func() { this.login(false) })
			})
		}
	} else {
		st.SetSetting(store.SK_LAST_LOGINED, "false")
	}
}

func (this *MainWindow) saveLastLoginHost() {
	uiw := uictx.uiw
	srvurl := uiw.ComboBox_6.CurrentText()
	st := appctx.GetStorage()
	log.Println(st)
	st.SetSetting(store.SK_HOMESERVER_URL, srvurl)
}
