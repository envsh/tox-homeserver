package main

/*
#cgo LDFLAGS: -L ../qofiaui -lqofiaui
// -lQt5Widgets -lQt5Gui -lQt5Core -lstdc++

#include <stdlib.h>

// #include "../qofiaui/qofiaui.h"
extern void qofiaui_main(void* ctx);
extern void qofiaui_dmcommand(char*cmdmsg);

extern void onui_command(char* cmd);
extern char* onui_loadmsg(char* uid, int maxcnt);
*/
import "C"
import (
	"encoding/json"
	"gopp"
	"log"
	"strings"
	"unsafe"
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

const uicmdsep = "/"

//export onui_command
func onui_command(cmdc *C.char) {
	cmd := C.GoString(cmdc)
	onui_command_go(cmd)
}
func onui_command_go(cmdmsg string) {
	log.Println(cmdmsg)
	args := strings.Split(cmdmsg, uicmdsep)
	switch args[0] {
	case "login":
		go dm.initAppBackend()
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
		imsgos = append(imsgos, imsgo)
	}
	bcc, err := json.Marshal(imsgos)
	gopp.ErrPrint(err, uid, maxcnt)
	return string(bcc)
}

func dispatchEvent2c(evtmsg string) {
	evtmsgc := C.CString(evtmsg)
	defer C.free(unsafe.Pointer(evtmsgc))

	C.qofiaui_dmcommand(evtmsgc)
}
