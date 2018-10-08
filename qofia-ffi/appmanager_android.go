package main

import (
	"log"
	"strings"
	"time"

	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	"github.com/kitech/qt.go/qtandroidextras"
	"github.com/kitech/qt.go/qtrt"
)

func hello() {
	log.Println(qtandroidextras.QAndroidBinder__Normal)
}

func testRunOnAndroidThread() {
	qtandroidextras.RunOnAndroidThread_(func() {
		log.Println("this is run on android thread.")
	})
	qtandroidextras.RunOnAndroidThread_(func() { _ExceptionCheckAndClear() })
}

func ShowToast(message string, duration uint64) {
	qtandroidextras.ShowToast(message, duration)
}

func KeepScreenOn(on bool) {
	// maybe hang here
	// qtandroidextras.RunOnAndroidThread(func() { _KeepScreenOn(on) })
	// 好像是qtandroidextras.RunOnAndroidThread再调回来的问题，假设不调回来可能行
	// qtandroidextras.RunOnAndroidThread(func() { _KeepScreenOn2(on) })
	// _KeepScreenOn2(true)
	qtandroidextras.KeepScreenOn(on)
}

func _KeepScreenOn(on bool) {
	activity := qtandroidextras.AndroidActivity()
	if activity.IsValid() {
		window := activity.CallObjectMethod("getWindow", "()Landroid/view/Window;")
		if window.IsValid() {
			FLAG_KEEP_SCREEN_ON := 0x00000080 // 128
			if on {
				window.CallMethod("addFlags", "(I)V", FLAG_KEEP_SCREEN_ON)
			} else {
				window.CallMethod("clearFlags", "(I)V", FLAG_KEEP_SCREEN_ON)
			}
		}
	}

	_ExceptionCheckAndClear()
}

func _ExceptionCheckAndClear() {
	jenv := qtandroidextras.NewQAndroidJniEnvironment()
	if jenv.JNIEnv().ExceptionCheck() {
		log.Println("have JNIEnv exception, clear...")
		jenv.JNIEnv().ExceptionClear()
	}
}

// 直接调用完整封装的也会hang住，所以应该不是qt.go封装的方法的问题。
// hang住的原因是CPU使用率高，这是为啥？
// 25001  2  50% S    29 2209720K 144988K  fg u0_a196  org.qtproject.example.golem
// 对应的qt代码
/*
   QtAndroid::runOnAndroidThread([on]{
           QAndroidJniObject activity = QtAndroid::androidActivity();
           if (activity.isValid()) {
               QAndroidJniObject window =
                   activity.callObjectMethod("getWindow", "()Landroid/view/Window;");

               if (window.isValid()) {
                   const int FLAG_KEEP_SCREEN_ON = 128;
                   if (on) {
                       window.callMethod<void>("addFlags", "(I)V", FLAG_KEEP_SCREEN_ON);
                   } else {
                       window.callMethod<void>("clearFlags", "(I)V", FLAG_KEEP_SCREEN_ON);
                   }
               }
           }
           QAndroidJniEnvironment env;
           if (env->ExceptionCheck()) {
               env->ExceptionClear();
           }
       });
*/
func _KeepScreenOn2(on bool) {
	ion := qtrt.IfElseInt(on, 1, 0)
	rv, err := qtrt.InvokeQtFunc6("C_KeepScreenOn", qtrt.FFI_TYPE_POINTER, ion)
	qtrt.ErrPrint(err, rv)
}

// for Android Intent
// should block
func _CheckIntentMessage() {
	className := "im/fixlan/golem/PendingIntents"
	for {
		time.Sleep(3 * time.Second)

		cnt := qtandroidextras.QAndroidJniObject__callStaticMethod_6(
			className, "GetPendingCount", "()I")
		if cnt > 0 {
			log.Println("PendingIntents count:", cnt)
		}
		if cnt > 0 {
			jstro := qtandroidextras.QAndroidJniObject__callStaticObjectMethod(
				className, "GetPendingData", "()Ljava/lang/String;")
			str := jstro.ToString()
			log.Println("PendingIntents data:", len(str), str)

			if strings.HasPrefix(str, thscom.GroupTitleSep) {
				str = str[len(thscom.GroupTitleSep):]
			}
			parts := strings.Split(str, thscom.GroupTitleSep)
			log.Println(len(parts), parts)
			for idx := 0; idx < len(parts)/2; idx++ {
				evto := &thspbs.Event{}
				evto.Name = "IntentMessage"
				evto.Args = []string{parts[idx*2], parts[idx*2+1]}
				intentQueue <- evto
				uictx.mech.Trigger()
				log.Println("Populate a ", evto.Name)
			}
		}
	}
	log.Println("done.")
}
