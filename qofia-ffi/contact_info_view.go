package main

import (
	"log"

	"github.com/kitech/qt.go/qtrt"
)

func (this *MainWindow) initContactInfoPage() {
	this.initContactInfoUi()
	this.initContactInfoSignals()
	this.initContactInfoEvents()
}

func (this *MainWindow) initContactInfoUi() {

}

func (this *MainWindow) initContactInfoSignals() {
	qtrt.Connect(this.PushButton_10, "clicked(bool)", func(checked bool) {
		this.switchUiStackPopBack()
	})
	qtrt.Connect(this.PushButton_11, "clicked(bool)", func(checked bool) {
		this.switchUiStackPopBack()
	})

	qtrt.Connect(this.CheckBox_4, "clicked(bool)", func(checked bool) {
		log.Println(checked)
	})
	qtrt.Connect(this.RadioButton_5, "clicked(bool)", func(checked bool) {
		log.Println(checked)
	})
	qtrt.Connect(this.RadioButton_6, "clicked(bool)", func(checked bool) {
		log.Println(checked)
	})
	qtrt.Connect(this.RadioButton_7, "clicked(bool)", func(checked bool) {
		log.Println(checked)
	})
	qtrt.Connect(this.CheckBox_5, "clicked(bool)", func(checked bool) {
		log.Println(checked)
	})
	qtrt.Connect(this.PushButton_12, "clicked(bool)", func(checked bool) {
		log.Println(checked)
	})
}

func (this *MainWindow) initContactInfoEvents() {

}

func (this *MainWindow) fillConactInfo(item *RoomListItem) {
	this.Label_32.SetText(item.GetName())
	this.LineEdit_9.SetText(item.GetId())
	if item.cticon != nil {
		this.ToolButton_31.SetIcon(item.cticon)
	}
}

func (this *MainWindow) fillConactSetting(item *RoomListItem) {

}
