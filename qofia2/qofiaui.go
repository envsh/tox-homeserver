package main

/*
#cgo LDFLAGS: -L ../qofiaui -lqofiaui

#include <stdlib.h>

// #include "../qofiaui/qofiaui.h"
extern void qofiaui_main(void* ctx);
extern void qofiaui_dmcommand(char*cmdmsg);

extern void onui_command(char* cmd);
*/
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

func qofiaui_main() {
	type uicontextc struct {
		onui_command unsafe.Pointer
	}
	var uictx uicontextc
	uictx.onui_command = (unsafe.Pointer)(C.onui_command)

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

func dispatchEvent2c(evtmsg string) {
	evtmsgc := C.CString(evtmsg)
	defer C.free(unsafe.Pointer(evtmsgc))

	C.qofiaui_dmcommand(evtmsgc)
}
