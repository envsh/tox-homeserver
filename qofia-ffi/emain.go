package main

import (
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {

	// Create application
	if runtime.GOOS == "android" {
		os.Setenv("QT_AUTO_SCREEN_SCALE_FACTOR ", "1.5")
		qtcore.QCoreApplication_SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true)
	} else {
		// qtrt.SetFinalizerObjectFilter(finalizerFilter)
	}
	// qtrt.SetDebugFFICall(true)
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	if false {
		app.SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true) // for android
	}
	uictx.qtapp = app

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
