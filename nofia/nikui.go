package main

/*
#cgo CFLAGS: -I/opt/nim/lib
#cgo LDFLAGS: -L. -lnikui -ldl -lX11 -lXft -lXrender

#include "nikui.h"
*/
import "C"
import (
	"log"
	"runtime"
	"time"
	"unsafe"
)

var gnenv unsafe.Pointer

func nimMain() { C.NimMain() }
func newNimenv() unsafe.Pointer {
	ne := C.newNimenv()
	log.Println(ne)
	gnenv = ne
	return unsafe.Pointer(ne)
}

func newNkwindow()  { C.newNkwindow(gnenv) }
func NkwindowOpen() { C.NkwindowOpen(gnenv) }

func startuimain() {
	go func() {
		runtime.LockOSThread()
		nimMain()
		time.Sleep(50 * time.Millisecond)
		newNimenv()
		newNkwindow()
		NkwindowOpen()
	}()
}

////
