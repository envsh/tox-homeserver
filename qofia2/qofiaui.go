package main

/*
#cgo LDFLAGS: -L ../qofiaui -lqofiaui
// -lQt5Widgets -lQt5Gui -lQt5Core -lstdc++

#include <stdlib.h>

// #include "../qofiaui/qofiaui.h"
// called by go
extern void qofiaui_main(void* ctx);
extern void qofiaui_daemoncmd(char*cmdmsg);

// called by cpp
extern void onui_command(char* cmd);
extern char* onui_loadmsg(char* uid, int maxcnt);
*/
import "C"
import (
	"encoding/json"
	"gopp"
	"gopp/cgopp"
	"log"
	"math/rand"
	"strings"
	"unsafe"

	thscli "tox-homeserver/client"
)

func qofiaui_main() {
	type uicontextc struct {
		onui_command unsafe.Pointer
		onui_loadmsg unsafe.Pointer
	}
	var uictx uicontextc
	uictx.onui_command = (unsafe.Pointer)(C.onui_command)
	uictx.onui_loadmsg = (unsafe.Pointer)(C.onui_loadmsg)

	C.qofiaui_main(unsafe.Pointer(&uictx))
}

const uicmdsep = "|"

//export onui_command
func onui_command(cmdc *C.char) {
	cmd := C.GoString(cmdc)
	onui_command_go(cmd)
}
func onui_command_go(cmdmsg string) {
	args := strings.Split(cmdmsg, uicmdsep)
	log.Println(cmdmsg, args)
	switch args[0] {
	case "login":
		go dm.initAppBackend(args[1])
	case "sendmsg":
		uid := args[1]
		msg := args[2]
		frndnum, err := dm.vtcli.FriendByPublicKey(uid)
		grpo := dm.vtcli.Binfo.GetGroupById(uid)
		if grpo != nil {
			dm.vtcli.ConferenceSendMessage(grpo.Gnum, 0, msg, 0)
		} else if err == nil {
			dm.vtcli.FriendSendMessage(frndnum, msg, 0)
		}
	}
}

//export onui_loadmsg
func onui_loadmsg(uid *C.char, maxcnt C.int) *C.char {
	s := onui_loadmsg_go(C.GoString(uid), int(maxcnt))
	ret := C.CString(s)
	// defer C.free(unsafe.Pointer(ret))
	return ret
}
func onui_loadmsg_go(uid string, maxcnt int) string {
	msgos := dm.mdl.GetNewestMsgs(uid, maxcnt)
	var imsgos []interface{}
	for _, msgo := range msgos {
		var imsgo []string
		imsgo = append(imsgo, msgo.MsgUi)
		imsgo = append(imsgo, msgo.PeerNameUi)
		imsgo = append(imsgo, msgo.TimeUi)
		imsgo = append(imsgo, msgo.RefmtToHtml())
		imsgo = msgargsfill(imsgo)
		imsgos = append(imsgos, imsgo)
	}
	// for len(imsgos) < maxcnt {
	// 	imsgos = append(imsgos, testmsg())
	// }
	bcc, err := json.Marshal(imsgos)
	gopp.ErrPrint(err, uid, maxcnt)
	return string(bcc)
}
func msgo2cfmt(msgo *thscli.Message, uid string) []string {
	var imsgo []string
	imsgo = append(imsgo, msgo.MsgUi)
	imsgo = append(imsgo, msgo.PeerNameUi)
	imsgo = append(imsgo, msgo.TimeUi)
	imsgo = append(imsgo, msgo.RefmtToHtml())
	// imsgo = append(imsgo, uid)
	imsgo = msgargsfill(imsgo)
	return imsgo
}

// 填满固定个数
func msgargsfill(msgo []string) []string {
	const num = 10
	for i := 0; i < num-len(msgo); i++ {
		msgo = append(msgo, "")
	}
	return msgo
}
func testmsg() []string {
	var imsgo []string
	imsgo = append(imsgo, "hehe 消息oiefjaweiofwefwefifae")
	imsgo = append(imsgo, "hehe sender xyz")
	imsgo = append(imsgo, "ab:ef:egg.iwafwef")
	return imsgo
}

func dispatchEvent2c(evtmsg string) {
	evtmsgc := C.CString(evtmsg)
	defer func() {
		C.free(unsafe.Pointer(evtmsgc))
		if rand.Uint32()%5 == 0 {
			cgopp.MallocTrim()
		}
	}()

	C.qofiaui_daemoncmd(evtmsgc)
}
