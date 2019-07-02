package main

import (
	"unsafe"

	thscli "tox-homeserver/client"
)

type daemon struct {
	appctx *thscli.AppContext
	vtcli  *thscli.LigTox
	uictx  unsafe.Pointer
	mdl    *thscli.DataModel

	baseInfoGot bool
	msgnotifych chan bool
}

var dm *daemon // = daemon{}

func newDaemon() *daemon {
	dm := &daemon{}
	dm.appctx = thscli.NewAppContext()
	dm.msgnotifych = make(chan bool, 32)
	dm.mdl = thscli.NewDataModel(512, func() {})

	go dm.pollmsg()
	return dm
}
