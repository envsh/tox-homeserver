package main

import (
	"gopp"
	"log"
	"runtime"
	"strings"
	"time"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

var _HeaderFooterState = &struct {
	viewStack *lls.Stack
	quitTimes int
	ctrlMenu  *qtwidgets.QMenu
	acts      []*qtwidgets.QAction
}{
	lls.New(), 0, nil, nil,
}

func (this *MainWindow) initHeaderFooter() {
	this.initHeaderFooterState()
	this.initHeaderFooterUi()
	this.initHeaderFooterSignals()
	this.initHeaderFooterEvents()
}

func (this *MainWindow) initHeaderFooterState() {
	hfs := _HeaderFooterState
	hfs.ctrlMenu = qtwidgets.NewQMenu__()

	var act *qtwidgets.QAction
	act = hfs.ctrlMenu.AddAction("Load css for test")
	act.SetToolTip("Load css for test.(./theme/apptst.css)")
	qtrt.Connect(act, "triggered(bool)", func(bool) { setAppStyleSheet() })
	hfs.acts = append(hfs.acts, act)

	act = hfs.ctrlMenu.AddAction("Show toast for test")
	qtrt.Connect(act, "triggered(bool)", func(bool) { ShowToast("hehehh哈哈eehhe", 1) })
	hfs.acts = append(hfs.acts, act)

	act = hfs.ctrlMenu.AddAction("Keep screen on")
	act.SetToolTip("Keep screen on. (Android only)")
	act.SetCheckable(true)
	qtrt.Connect(act, "triggered(bool)", func(checked bool) { KeepScreenOn(checked) })
	hfs.acts = append(hfs.acts, act)

	act = hfs.ctrlMenu.AddAction("Placeholder1")
	act.SetCheckable(true)
	hfs.acts = append(hfs.acts, act)

	act = hfs.ctrlMenu.AddSeparator()
	hfs.acts = append(hfs.acts, act)

	act = hfs.ctrlMenu.AddAction("Run Go GC")
	qtrt.Connect(act, "triggered(bool)", func(bool) { runtime.GC() })
	hfs.acts = append(hfs.acts, act)
}
func (this *MainWindow) initHeaderFooterUi() {}

func (this *MainWindow) initHeaderFooterSignals() {
	qtrt.Connect(this.ToolButton_33, "clicked(bool)", func(bool) {
		quit := this.onAppBackButton(true)
		this.quitClean(quit)
	})
	qtrt.Connect(this.ToolButton_19, "clicked(bool)", func(bool) {
		pos := this.ToolButton_19.Pos()
		pos.SetY(pos.Y() + this.ToolButton_19.Height())
		pos = this.MainWindow.MapToGlobal(pos)
		_HeaderFooterState.ctrlMenu.Popup__(pos)
	})
}

func (this *MainWindow) initHeaderFooterEvents() {
	hfs := _HeaderFooterState
	hfs.ctrlMenu.InheritEvent(func(arg0 *qtcore.QEvent) bool {
		hlpevt := qtgui.NewQHelpEventFromPointer(arg0.GetCthis())
		if arg0.Type() == qtcore.QEvent__ToolTip && hfs.ctrlMenu.ActiveAction() != nil {
			qtwidgets.QToolTip_ShowText(hlpevt.GlobalPos(), hfs.ctrlMenu.ActiveAction().ToolTip(), nil)
		} else {
			qtwidgets.QToolTip_HideText()
		}
		hlpevt.SetCthis(nil) // avoid double free?
		return hfs.ctrlMenu.Event(arg0)
	})
}

func (this *MainWindow) quitClean(quit bool) {
	if !quit {
		return
	}
	// do something clean...
	if gopp.IsAndroid() { // how elegant exit?
	} else {
		uictx.qtapp.Quit()
	}
}

// return true for real quit
func (this *MainWindow) onAppBackButton(bool) bool {
	stack := _HeaderFooterState.viewStack
	log.Println("Current stack:", stack.Size(), strings.Replace(stack.String(), "\n", " ", -1))

	initNum := 2 // 1: loginui, 2: mainui
	switch {
	case stack.Size() < initNum: // do nothing
		return true
	case stack.Size() == initNum:
		if _HeaderFooterState.quitTimes == 0 {
			_HeaderFooterState.quitTimes += 1
			// show toast here, "tap again to quit"
			hintxt := "Click back button again in 1 sec to quit."
			ShowToast(hintxt, 1)
			log.Println(hintxt)
			time.AfterFunc(1*time.Second, func() { _HeaderFooterState.quitTimes = 0 })
		} else {
			// real quit
			log.Println("Double click back button in 1 sec, should quit.")
			// uictx.qtapp.Quit()
			return true
		}
	case stack.Size() > initNum:

		uinox, ok := stack.Pop()
		if ok {
			uinox, _ = stack.Peek()
			this.switchUiStackPop(uinox.(int))
		}
	}
	return false
}

func (this *MainWindow) onAppHomeButton(bool) {

}

func (this *MainWindow) onAppSwitchButton(bool) {

}
