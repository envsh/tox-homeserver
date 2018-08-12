package main

import (
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"

	"tox-homeserver/gomain2c"
)

// for android, but other OS still ok
func init() { gomain2c.Set(main) }

func main() {
	log.Println("Enter main...")

	if runtime.GOOS == "linux" {
		SetupGops()
	}
	// Create application
	if runtime.GOOS == "android" {
		os.Setenv("QT_AUTO_SCREEN_SCALE_FACTOR ", "1.5")
		qtcore.QCoreApplication_SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true)
	} else {
		// qtrt.SetFinalizerObjectFilter(finalizerFilter)
	}
	qtrt.SetDebugFFICall(false)
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	setupFinalizerOnUi()
	if false {
		app.SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true) // for android
	}
	uictx.qtapp = app
	qtgui.QFontDatabase_AddApplicationFont("resource/fzlt.ttf")
	qtgui.QFontDatabase_AddApplicationFont("resource/emojione-android.ttf") // "Emoji One"

	setAppStyleSheet()
	// Create main window
	uictx.mw = NewMainWindow()
	uictx.uiw.MainWindow.Show()
	setAppStyleSheetTheme(0)
	uictx.init()

	go func() {
		time.Sleep(2 * time.Second)
		// app.Exit(0)
	}()
	// Execute app
	app.Exec()
}

func DumpCallers(pcs []uintptr) {
	log.Println("DumpCallers...", len(pcs))
	for idx, pc := range pcs {
		pcfn := runtime.FuncForPC(pc)
		file, line := pcfn.FileLine(pc)
		log.Println(idx, pcfn.Name(), file, line)
	}
	if len(pcs) > 0 {
		log.Println()
	}
}

func finalizerFilter(ov reflect.Value) bool {
	parts := strings.Split(ov.Type().String(), ".")
	clsname := parts[len(parts)-1]
	callers := qtrt.GetCtorAllocStack(clsname)
	_ = callers

	insure := false
	switch ov.Type().String() {
	// case "*qtcore.QString":
	// case "*qtcore.QSize":
	// case "*qtwidgets.QSpacerItem": // crash
	//	insure = true
	//	DumpCallers(callers)
	default:
		insure = true
	}
	if insure {
		log.Println(ov.Type().String(), ov)
	}
	return insure
}

var finalMech *Notifier

func setupFinalizerOnUi() {
	// usually finalizer run on go's seperate thread, cause crash often. so fix it.
	uiqc := make(chan func(), 1)
	mech := NewNotifier(func() {
		for len(uiqc) > 0 {
			remfn := <-uiqc
			remfn()
		}
	})

	qtrt.FinalProxyFn = func(f func()) {
		uiqc <- f
		mech.Trigger()
		// f()
	}
	finalMech = mech
}
