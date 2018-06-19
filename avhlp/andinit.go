//+build android

package avhlp

/*
#include <jni.h>

// extern JavaVM* current_vm; // see golang.org/x/mobile/internal/mobileinit
*/
import "C"
import (
	"log"
	"unsafe"

	"golang.org/x/mobile/exp/audio/al"
)

func SetCurrentVM(jvm, actx unsafe.Pointer) {
	jvmpp := al.GetCurrentVMAddr()
	actxpp := al.GetCurrentCtxAddr()

	*(**C.JavaVM)(jvmpp) = (*C.JavaVM)(jvm)
	*(*C.jobject)(actxpp) = (C.jobject)(actx)
	log.Println("Set mobileinit.current_vm to:", jvm, actx)
	// actx is like 0x765, strange...
}
