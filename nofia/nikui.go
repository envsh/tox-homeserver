package main

/*
#cgo CFLAGS: -I/opt/nim/lib -I/usr/include/freetype2
#cgo LDFLAGS: -L. -lnikui -ldl -lX11 -lXft -lXrender -lm

#include "nikui.h"
*/
import "C"
import (
	"gopp"
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

		// must call nim func in this thread
		for {
			evtdat := <-nimthch
			dispatchEventNim1(evtdat)
		}
	}()
}

var nimthch = make(chan []byte)

func dispatchEventNim(evtdat []byte) {
	evtdat1 := gopp.BytesDup(evtdat)
	nimthch <- evtdat1
}
func dispatchEventNim1(evtdat []byte) {
	evtdat = append(evtdat, 0)
	p := (*C.char)(unsafe.Pointer(&evtdat[0]))
	// p := (unsafe.Pointer(&evtdat[0]))
	C.dispatchEventNim(p)
}

func dispatchEventRespNim1(evtdat []byte) {
	return
	evtdat = append(evtdat, 0)
	p := (*C.char)(unsafe.Pointer(&evtdat[0]))
	// p := (unsafe.Pointer(&evtdat[0]))
	C.dispatchEventRespNim(p)
}

////
