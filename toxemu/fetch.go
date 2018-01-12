package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"encoding/json"
	"gopp"
	"log"
	"unsafe"
)

//export thc_get_base_info
func thc_get_base_info(l *C.char) *C.char {
	if lt == nil {
		log.Println("New instance:")
		tox_new()
	}
	lt.GetBaseInfo()
	bcc, err := json.Marshal(lt.Binfo)
	gopp.ErrPrint(err)
	// log.Println(len(bcc), string(bcc))
	// *l = (C.int)(len(bcc))
	C.memcpy(unsafe.Pointer(l), unsafe.Pointer(&bcc[0]), C.size_t(len(bcc)))
	return l
}

//export thc_poll_event
func thc_poll_event(b *C.char) C.int {
	if e := lt.GetNextBackenEvent(); e != nil {
		C.memcpy(unsafe.Pointer(b), unsafe.Pointer(&e[0]), C.size_t(len(e)))
		return 1
	}
	return 0
}
