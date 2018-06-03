package main

import (
	"gopp"
	"log"
	"strings"
	"time"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/kitech/qt.go/qtrt"
)

var _HeaderFooterState = struct {
	viewStack *lls.Stack
	quitTimes int
}{
	lls.New(), 0,
}

func (this *MainWindow) initHeaderFooter() {
	this.initHeaderFooterUi()
	this.initHeaderFooterSignals()
	this.initHeaderFooterEvents()
}

func (this *MainWindow) initHeaderFooterUi() {}

func (this *MainWindow) initHeaderFooterSignals() {
	qtrt.Connect(this.PushButton_13, "clicked(bool)", func(bool) {
		quit := this.onAppBackButton(true)
		this.quitClean(quit)
	})
	qtrt.Connect(this.PushButton_14, "clicked(bool)", this.onAppHomeButton)
	qtrt.Connect(this.PushButton_15, "clicked(bool)", this.onAppSwitchButton)
}

func (this *MainWindow) initHeaderFooterEvents() {}

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

	initNum := 2
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
			this.switchUiStackPop(uinox.(int))
		}
	}
	return false
}

func (this *MainWindow) onAppHomeButton(bool) {

}

func (this *MainWindow) onAppSwitchButton(bool) {

}
