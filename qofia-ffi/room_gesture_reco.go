package main

import (
	"gopp"
	"log"
	"runtime"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

type RoomGestureReco struct {
	touchStart    time.Time
	touchStartEvt *qtgui.QMouseEvent
	touchEnd      time.Time
	touchEndEvt   *qtgui.QMouseEvent

	OnClick       func(item *RoomListItem, pos *qtcore.QPoint)
	OnLongTouch   func(item *RoomListItem, pos *qtcore.QPoint)
	OnDoubleClick func(item *RoomListItem, pos *qtcore.QPoint)
}

func NewRoomGestureReco() *RoomGestureReco {
	this := &RoomGestureReco{}
	return this
}

func (this *RoomGestureReco) onMousePress(item *RoomListItem, event *qtgui.QMouseEvent) {
	// log.Println(event)
	if event.Button() == qtcore.Qt__LeftButton {
		this.touchStart = time.Now()
		this.touchStartEvt = event
		if runtime.GOOS == "linux" {
			// think as clicked
			// log.Println("clicked")
			if this.OnClick != nil {
				//	this.OnClick()
			}
		}
	}
}

func (this *RoomGestureReco) onMouseRelease(item *RoomListItem, event *qtgui.QMouseEvent) {
	// log.Println(event)
	now := time.Now()
	if event.Button() == qtcore.Qt__LeftButton {
		// if runtime.GOOS == "android" {
		if now.Sub(this.touchStart).Seconds() < 0.5 {
			// reco as clicked
			if this.OnClick != nil {
				this.OnClick(item, event.Pos())
			}
		}
		// this.touchStart = time.Unix(0, 0)
		// }

		if now.Sub(this.touchStart).Seconds() > 2. {
			// reco as long touch
			log.Println("long touch")
			if this.OnLongTouch != nil {
				this.OnLongTouch(item, event.GlobalPos())
			}
		}
		this.touchStart = time.Unix(0, 0)
	}
}

func (this *RoomGestureReco) onMouseMove(item *RoomListItem, event *qtgui.QMouseEvent) {
}

type MessageListGesture struct {
	w           qtwidgets.QWidget_ITF
	OnLongTouch func(*qtcore.QPointF)
}

func NewMessageListGesture(w qtwidgets.QWidget_ITF) *MessageListGesture {
	this := &MessageListGesture{}
	this.w = w

	w.QWidget_PTR().GrabGesturep(qtcore.Qt__TapAndHoldGesture)
	w.QWidget_PTR().InheritEvent(this.OnEvent)
	return this
}

func (this *MessageListGesture) OnEvent(event *qtcore.QEvent) bool {
	if this.OnLongTouch == nil {
		return this.w.QWidget_PTR().Event(event)
	}

	if event.Type() == qtcore.QEvent__Gesture {
		gevt := qtwidgets.NewQGestureEventFromPointer(event.GetCthis())
		gt := gevt.Gesture(qtcore.Qt__TapAndHoldGesture)
		if gt.GetCthis() != nil {
			log.Println("Got Qt__TapAndHoldGesture:",
				gopp.IfElseStr(gt.State() == qtcore.Qt__GestureFinished, "Finished", "Start"))
			if gt.State() == qtcore.Qt__GestureFinished {
				if this.OnLongTouch != nil {
					this.OnLongTouch(gt.HotSpot())
					event.Accept()
					return true
				}
			}
		}
	}
	return this.w.QWidget_PTR().Event(event)
}
