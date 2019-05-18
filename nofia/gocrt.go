package main

/*
#cgo LDFLAGS: -ldl
#cgo CFLAGS: -D_GNU_SOURCE

#include <stdio.h>
#include <dlfcn.h>

static void* cnimcallfnptr = 0;

// why cannot
static void dlsyms() {
    void* sym = dlsym(RTLD_DEFAULT, "cnimcall");
    printf("111cnimcallfnptr=%p\n", sym);
}
static void cnimcallset(void*ptr) {cnimcallfnptr = ptr;}
static void cnimcall(void*fn, void*args) {
    ((void(*)(void*, void*))(cnimcallfnptr))(fn, args);
}
*/
import "C"
import (
	"log"
	"math/rand"
	"sync"
	"time"
	"unsafe"

	"github.com/alangpierce/go-forceexport"
)

func init() {
	//C.dlsyms()
}

//export goinit
func goinit(cnimcallfnptr unsafe.Pointer) {
	C.cnimcallset(cnimcallfnptr)
}

//export gomain
func gomain() {
	// C.dlsyms()
	go func() {
		for {
			time.Sleep(5 * time.Hour)
		}
	}()
	main()
}
func main() { select {} }

var cgocall func(fnptr unsafe.Pointer, args unsafe.Pointer) int32

func init() {
	rand.Seed(time.Now().Unix())
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.SetFlags(log.Flags() ^ log.Ldate)
	// log.SetFlags(log.Flags() ^ log.Ltime)

	err := forceexport.GetFunc(&cgocall, "runtime.cgocall")
	if err != nil {
		log.Println(err)
	}

	if rand.Uint64() == 1 && rand.Uint64() == 2 {
		gogo(nil, nil)
	}
}

func cnimcall(fn unsafe.Pointer, args unsafe.Pointer) {
	C.cnimcall(fn, args)
}

//export gogo
func gogo(fnptr unsafe.Pointer, args unsafe.Pointer) {
	go func() {
		log.Println(fnptr, args)
		cnimcall(fnptr, args)
	}()
}

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

type _type struct {
	a int
}
type sudog struct {
	a int
}
type waitq struct {
	first *sudog
	last  *sudog
}

type mutex struct {
	key uintptr
}

var gorefs sync.Map // pointer => interface

func chan2pointer(c chan unsafe.Pointer) unsafe.Pointer {
	cp := *(**hchan)(unsafe.Pointer(&c))
	return unsafe.Pointer(cp)
}
func pointer2chan(p unsafe.Pointer) chan unsafe.Pointer {
	a := *(*chan unsafe.Pointer)(unsafe.Pointer(&p))
	return a
}

//export gochannew
func gochannew(n int) unsafe.Pointer {
	c := make(chan unsafe.Pointer, n)
	p := chan2pointer(c)
	gorefs.Store(p, c)
	return p
}

//export gochanfree
func gochanfree(c unsafe.Pointer) {
	gorefs.Delete(c)
}

//export gochansend
func gochansend(pc unsafe.Pointer, v unsafe.Pointer) {
	c := pointer2chan(pc)
	c <- v
}

//export gochanrecv
func gochanrecv(pc unsafe.Pointer) unsafe.Pointer {
	c := pointer2chan(pc)
	v := <-c
	return v
}
