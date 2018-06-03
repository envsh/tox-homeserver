package main

import (
	"log"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func Repolish(w qtwidgets.QWidget_ITF) {
	syl := w.QWidget_PTR().Style()
	syl.Unpolish(w)
	syl.Polish(w)
}

func SetScrollContentTrackerSize(sa *qtwidgets.QScrollArea) {
	wgt := sa.Widget()
	sa.InheritResizeEvent(func(arg0 *qtgui.QResizeEvent) {
		osz := arg0.OldSize()
		nsz := arg0.Size()
		if false {
			log.Println(osz.Width(), osz.Height(), nsz.Width(), nsz.Height())
		}
		if osz.Width() != nsz.Width() {
			wgt.SetMaximumWidth(nsz.Width())
		}
		// this.ScrollArea_2.ResizeEvent(arg0)
		arg0.Ignore() // I ignore, you handle it. replace explict call parent's
	})
}
