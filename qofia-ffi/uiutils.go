package main

import (
	"gopp"
	"log"

	"github.com/kitech/qt.go/qtcore"
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

func SetQLabelElideText(lab *qtwidgets.QLabel, txt string, skipTooltip ...bool) {
	// font := lab.PaintEngine().Painter().Font()
	font := lab.Font()
	rect := lab.Rect()
	sz1 := lab.Size()
	sz2 := lab.SizeHint()
	gm := lab.Geometry()

	elwidth := int(gopp.MaxU32([]uint32{uint32(rect.Width()), uint32(sz2.Width())}))
	elwidth = gopp.IfElseInt(elwidth > 500, rect.Width(), elwidth)
	elwidth = rect.Width()
	elwidth = gopp.IfElseInt(elwidth < 150, sz2.Width(), elwidth)
	elwidth = gopp.IfElseInt(elwidth > 150, elwidth-10, elwidth)
	if false {
		log.Println(rect.Width(), sz1.Width(), sz2.Width(), gm.Width(), lab.ObjectName(), txt)
	}

	fm := qtgui.NewQFontMetrics(font)
	etxt := fm.ElidedText__(txt, qtcore.Qt__ElideRight, elwidth)
	if false {
		log.Println(len(txt), len(etxt))
	}

	lab.SetText(etxt)
	if len(skipTooltip) == 0 {
		lab.SetToolTip(txt)
	}
}

func SetQWidgetDropable(w qtwidgets.QWidget_ITF, dropable bool) {
	w.QWidget_PTR().InheritDragEnterEvent(func(event *qtgui.QDragEnterEvent) {
		if dropable {
			event.AcceptProposedAction()
		}
	})
}
