package main

import (
	"log"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

type PullRefreshGestureReco struct {
	qapp *qtwidgets.QApplication
	srlo *qtwidgets.QScroller
	srla *qtwidgets.QScrollArea
	srlw *qtwidgets.QWidget

	pulling       bool // after pull start and before release pull
	backing       bool // after release pull and before backing scoller done
	pullStartY    int
	pullStopY     int
	pullDeltaY    int
	scrollLastY   int
	mousePressing bool

	OnPullRefresh  func()
	OnPullBackDone func() // 控制下拉复位完成
}

func NewPullRefreshGestureReco(area *qtwidgets.QScrollArea, w *qtwidgets.QWidget, qapp *qtwidgets.QApplication) *PullRefreshGestureReco {
	this := new(PullRefreshGestureReco)
	this.srla = area
	this.srlw = w
	this.qapp = qapp

	area.InheritEvent(this._Event)
	w.InheritMousePressEvent(this._MousePressEvent)
	w.InheritMouseReleaseEvent(this._MouseReleaseEvent)
	w.InheritMouseMoveEvent(this._MouseMoveEvent)
	return this
}

func (this *PullRefreshGestureReco) vertScrollBarVisible() bool {
	n := this.srla.VerticalScrollBar().Maximum()
	return n > 0
}

// 如果有滚动条，则使用gesture检测。如果没有滚动条，则需要使用纯keypress检测。
func (this *PullRefreshGestureReco) _Event(arg0 *qtcore.QEvent) bool {
	sbvis := this.vertScrollBarVisible()
	if sbvis {
		return this._EventHasScroll(arg0)
	} else {
		return this._EventNoScroll(arg0)
	}
}
func (this *PullRefreshGestureReco) _EventNoScroll(arg0 *qtcore.QEvent) bool {
	srla := this.srla
	arg0.Ignore()
	srla.Event(arg0)

	// log.Println("hehehhe", arg0.Type(), rand.Int())
	evty := arg0.Type()

	if evty == qtcore.QEvent__MouseButtonPress {
		if this.mousePressing == false {
			this.mousePressing = true
			mevt := qtgui.NewQMouseEventFromPointer(arg0.GetCthis())
			this.pullStartY = mevt.Pos().Y()
		}

	} else if evty == qtcore.QEvent__MouseButtonRelease {
		if this.mousePressing == true {
			this.mousePressing = false
			mevt := qtgui.NewQMouseEventFromPointer(arg0.GetCthis())
			cury := mevt.Pos().Y()
			offsety := cury - this.pullStartY
			if offsety > 40 {
				log.Println("release mouse, maybe should trigger refresh action")
				time.AfterFunc(1500*time.Millisecond, func() {
					log.Println("maybe should trigger back done action")
				})
			}
		}
	} else if evty == qtcore.QEvent__MouseMove {

	}

	return true
}
func (this *PullRefreshGestureReco) _EventHasScroll(arg0 *qtcore.QEvent) bool {
	srla := this.srla
	srlw := this.srlw
	vsbar := srla.VerticalScrollBar()
	arg0.Ignore()
	srla.Event(arg0)

	// log.Println("hehehhe", arg0.Type(), rand.Int())
	evty := arg0.Type()

	// mouse tracking
	mousePressingOld := this.mousePressing
	if evty == 202 || evty == 198 || evty == 205 {
		mousebtns := this.qapp.MouseButtons()
		if mousebtns == qtcore.Qt__LeftButton {
			if !this.mousePressing {
				this.mousePressing = true
			}
		} else {
			if this.mousePressing {
				this.mousePressing = false
			}
		}
	}

	// log.Println("hehehe", arg0.Type(), rand.Int(), this.mousePressing, mousePressingOld)
	if evty == 202 {
		if vsbar.Value() == 0 {
			this.pulling = true
			pos2 := srlw.Pos()
			pos2 = srlw.ParentWidget().MapToGlobal(pos2)
			if false {
				log.Println("guesture start", pos2.X(), pos2.Y())
			}
			this.pullStartY = pos2.Y()
			this.scrollLastY = pos2.Y()
		}
	} else if evty == 198 {
		pos2 := srlw.Pos()
		pos2 = srlw.ParentWidget().MapToGlobal(pos2)
		if false {
			log.Println("guesture in", pos2.X(), pos2.Y())
		}
		this.pullDeltaY = pos2.Y() - this.pullStartY
		// log.Println("offsety:", this.pullDeltaY, this.pullStartY, pos2.Y(), vsbar.Value())
		if mousePressingOld == true && this.mousePressing == false {
			log.Println("release mouse, maybe should trigger refresh action")
		}
	} else if evty == 205 {
		pos2 := srlw.Pos()
		pos2 = srlw.ParentWidget().MapToGlobal(pos2)

		// log.Println("scroll...", evty, this.pullStartY, pos2.Y()-this.scrollLastY, this.mousePressing, mousePressingOld)
		this.scrollLastY = pos2.Y()
		if this.pulling && this.scrollLastY == this.pullStartY {
			this.pulling = false // also reset other state
			log.Println("maybe should trigger back done action")
		}
	}

	return false
}

// seems press/release/move event swallow by gesture
func (this *PullRefreshGestureReco) _MousePressEvent(arg0 *qtgui.QMouseEvent) {
	sbvis := this.vertScrollBarVisible()
	log.Println("hehhee", sbvis)
	if sbvis {
		this._MousePressEventHasScroll(arg0)
	} else {
		this._MousePressEventNoScroll(arg0)
	}
}
func (this *PullRefreshGestureReco) _MousePressEventNoScroll(arg0 *qtgui.QMouseEvent) {
}
func (this *PullRefreshGestureReco) _MousePressEventHasScroll(arg0 *qtgui.QMouseEvent) {
	log.Println("hehhee")
	if !this.mousePressing {
		this.mousePressing = true
		this.pullStartY = arg0.GlobalY()
	}
}

func (this *PullRefreshGestureReco) _MouseReleaseEvent(arg0 *qtgui.QMouseEvent) {
	sbvis := this.vertScrollBarVisible()
	if sbvis {
		this._MouseReleaseEventHasScroll(arg0)
	} else {
		this._MouseReleaseEventNoScroll(arg0)
	}
}
func (this *PullRefreshGestureReco) _MouseReleaseEventNoScroll(arg0 *qtgui.QMouseEvent) {
}
func (this *PullRefreshGestureReco) _MouseReleaseEventHasScroll(arg0 *qtgui.QMouseEvent) {
	// TODO check left mouse
	if this.mousePressing {
		this.mousePressing = false
		log.Println(this.pulling)
		if this.pulling {
			log.Println("start refresh. maybe should set backing mode")
		} else if this.backing {

		}
	}
}

func (this *PullRefreshGestureReco) _MouseMoveEvent(event *qtgui.QMouseEvent) {
	sbvis := this.vertScrollBarVisible()
	log.Println("hehehhe", sbvis)
	if sbvis {
		this._MouseMoveEventHasScroll(event)
	} else {
		this._MouseMoveEventNoScroll(event)
	}
}
func (this *PullRefreshGestureReco) _MouseMoveEventNoScroll(event *qtgui.QMouseEvent) {
}
func (this *PullRefreshGestureReco) _MouseMoveEventHasScroll(event *qtgui.QMouseEvent) {
	if this.mousePressing {

	}
}

func (this *PullRefreshGestureReco) ReleaseReco() {

}
