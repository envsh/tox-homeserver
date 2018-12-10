package main

import (
	"log"
	"unsafe"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

var _vcs = _VideoCallState
var _VideoCallState = &struct {
	contact string
	imgch   chan *ImgData
}{
	"", make(chan *ImgData, 64),
}

func (this *MainWindow) initVideoCall() {
	this.initVideoCallUi()
	this.initVideoCallSignals()
	this.initVideoCallEvents()
}

func (this *MainWindow) initVideoCallUi() {

}

func (this *MainWindow) initVideoCallSignals() {
	qtrt.Connect(this.ToolButton_20, "clicked(bool)", func(checked bool) {
		// on enable mic
	})

	qtrt.Connect(this.ToolButton_30, "clicked(bool)", func(checked bool) {
		// on mute mixer
	})

	qtrt.Connect(this.ToolButton_21, "clicked(bool)", func(bool) {
		// on hangup
	})
}

func (this *MainWindow) initVideoCallEvents() {
	w := this.Widget_3
	w.InheritPaintEvent(func(event *qtgui.QPaintEvent) {
		p := NewQPainter(w)
		// p.FillRect_2(p.Viewport(), qtgui.NewQBrush_1(0))
		for len(_vcs.imgch) > 0 {
			imgd := <-_vcs.imgch
			if imgd == nil {
				p.FillRect2(p.Viewport(), qtgui.NewQBrush1(0))
				break
			}
			idptr := unsafe.Pointer(&imgd.data[0])
			width, height := imgd.width, imgd.height
			// imgo := qtgui.NewQImage_5_(idptr, width,height, d.ls, qtgui.QImage__Format_RGB32) // why black?
			// imgo := qtgui.NewQImage_3_(idptr, width, height, qtgui.QImage__Format_RGB32) // why ok?
			// OK format: QImage__Format_RGB888
			// strange format: QImage__Format_RGB666
			imgo := qtgui.NewQImage3p(idptr, width, height, qtgui.QImage__Format_RGB888) // why ok?
			p.DrawImage5(p.Viewport(), imgo)
		}
		qtgui.DeleteQPainter(p)
	})

}

func (this *MainWindow) initVideoCallStorage() {

}

//
type VideoPlayer struct {
	w       *qtwidgets.QWidget
	stopped bool
}

func NewVideoPlayer() *VideoPlayer {
	this := &VideoPlayer{}
	this.w = uictx.mw.Widget_3
	return this
}

func (this *VideoPlayer) Stop() {
	// just clean
	this.stopped = true
	for len(_vcs.imgch) > 0 {
		<-_vcs.imgch
	}
	// clear screen?
	_vcs.imgch <- nil
	runOnUiThread(func() { this.w.Update() })
}

type ImgData struct {
	data   []byte
	width  int
	height int
}

// TODO decoder frame
func (this *VideoPlayer) PutFrame(vframe []byte, width, height int) {
	if this.stopped {
		return
	}
	// log.Println("Convert video frame to QImage...", len(vframe), width, height, len(_vcs.imgch))
	// convert to QImage
	if len(_vcs.imgch) == cap(_vcs.imgch) {
		log.Println("chan is full,", len(vframe), width, height, len(_vcs.imgch))
		return
	}
	_vcs.imgch <- &ImgData{vframe, width, height}
	runOnUiThread(func() { this.w.Update() })
}
