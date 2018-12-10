package main

import (
	"log"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

type EmojiPanel struct {
	*Ui_EmojiPanel

	OnEmojiSelected func(string, string)
	catwins         []*EmojiCategory
}

func NewEmojiPanel() *EmojiPanel {
	this := &EmojiPanel{}

	this.Ui_EmojiPanel = NewUi_EmojiPanel2()

	w := this.QWidget_PTR()
	w.SetAttribute(qtcore.Qt__WA_ShowWithoutActivating, true)
	w.SetWindowFlags(qtcore.Qt__Tool | qtcore.Qt__FramelessWindowHint | qtcore.Qt__NoDropShadowWindowHint)

	SetScrollContentTrackerSize(this.ScrollArea)
	vsbar := this.ScrollArea.VerticalScrollBar()
	vsbar.SetSingleStep(vsbar.SingleStep() * 2)
	vsbar.SetPageStep(vsbar.PageStep() * 2)
	qtwidgets.QScroller_GrabGesture(this.ScrollArea, qtwidgets.QScroller__LeftMouseButtonGesture)

	innames, emojivecs := EmojiProvider.ToVec()
	shownames := EmojiProvider.ShowNames()
	for i := 0; i < len(innames); i++ {
		tmpi := i
		catwin := NewEmojiCategory(shownames[i], emojivecs[i])
		this.catwins = append(this.catwins, catwin)
		catwin.OnEmojiSelected = func(emoji string, shtname string) {
			log.Println("emoji:", tmpi, innames[tmpi], emoji, shtname)
			if this.OnEmojiSelected != nil {
				this.OnEmojiSelected(emoji, shtname)
			}
		}
		this.VerticalLayout_3.AddWidgetp(catwin.QWidget_PTR())
	}

	catbtnsz := 20
	catbtns := []*qtwidgets.QToolButton{
		this.ToolButton_33, this.ToolButton_34, this.ToolButton_35, this.ToolButton_36,
		this.ToolButton_37, this.ToolButton_38, this.ToolButton_39, this.ToolButton_40}
	for i, btn := range catbtns {
		btn.SetIconSize(qtcore.NewQSize1(catbtnsz, catbtnsz))
		btn.SetToolTip(shownames[i])
		tmpi := i
		qtrt.Connect(btn, "clicked(bool)", func(bool) {
			log.Println("clicked", tmpi, innames[tmpi], shownames[tmpi])
			this.showCategory(this.catwins[tmpi])
		})
	}

	// nothing happended???
	w.InheritFocusOutEvent(func(event *qtgui.QFocusEvent) {
		log.Println("hehehhe")
	})
	// no focus event found
	// w.InheritEvent(func(event *qtcore.QEvent) bool {
	//	log.Println(event.Type())
	//	return false
	// })
	return this
}

func (this *EmojiPanel) clear() {
	// signals
	// members
}

func (this *EmojiPanel) showCategory(category *EmojiCategory) {
	this.ScrollArea.EnsureWidgetVisiblep(category.Label)
}
