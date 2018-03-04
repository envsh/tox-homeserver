package main

import (
	"log"
	"runtime"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
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
			log.Println("clicked")
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
