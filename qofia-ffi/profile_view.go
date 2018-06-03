package main

// profile and login ui

import (
	"gopp"
	"log"

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
	qtrt.Connect(uiw.PushButton_7, "clicked(bool)", func(bool) {
		this.Label_24.SetText("...")
		this.saveLastLoginHost()
		go this.initAppBackend()
	})
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

func (this *MainWindow) loadDataForProfileUi() {
	st := appctx.GetStorage()
	setting, err := st.GetSetting(store.SK_HOMESERVER_URL)
	gopp.ErrPrint(err)
	if setting != nil {
		this.ComboBox_6.SetCurrentText(setting.Value)
	}
}

func (this *MainWindow) saveLastLoginHost() {
	uiw := uictx.uiw
	srvurl := uiw.ComboBox_6.CurrentText()
	appctx.GetStorage().SetSetting(store.SK_HOMESERVER_URL, srvurl)
}
