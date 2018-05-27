package main

import (
	"fmt"
	"gopp"
	"log"
	"strings"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

// for message list page usage
type MessageListWin struct {
	item  *RoomListItem
	gesto *MessageListGesture

	selinfo struct {
		has     bool
		start   int
		text    string
		objname string
		objwgt  *qtwidgets.QWidget
	}

	ctxmenu    *qtwidgets.QMenu
	actcopy    *qtwidgets.QAction
	actcopyone *qtwidgets.QAction
	actselall  *qtwidgets.QAction
	actselone  *qtwidgets.QAction
	actsep1    *qtwidgets.QAction
	actclear   *qtwidgets.QAction
	actquote   *qtwidgets.QAction
}

func NewMessageListWin() *MessageListWin {
	this := &MessageListWin{}
	return this
}

func (this *MessageListWin) SetRoom(item *RoomListItem) {
	if item == this.item {
		return
	}
	oldItem := this.item
	this.item = item

	mw := uictx.mw
	mw.Label_5.SetText(gopp.StrSuf4ui(item.GetName(), 32))
	if item.isgroup {
		mw.Label_6.SetVisible(true)
		mw.Label_7.SetVisible(false)
		mw.Label_6.SetText(fmt.Sprintf("%d users in chat", item.peerCount))
	} else {
		mw.Label_6.SetVisible(false)
		mw.Label_7.SetVisible(true)
		mw.Label_7.SetText(gopp.StrSuf4ui(item.frndInfo.GetStmsg(), 32))
	}
	mw.LabelMsgCount2.SetText(fmt.Sprintf("%3d", item.totalCount))
	mw.LabelMsgCount.SetText(fmt.Sprintf("%3d", item.totalCount))

	this.ReloadMessages(oldItem)
}

func (this *MessageListWin) ReloadMessages(oldItem *RoomListItem) {
	item := this.item

	btime := time.Now()
	vlo8 := uictx.uiw.VerticalLayout_3
	log.Println("clean msg list win:", vlo8.Count())
	if oldItem != nil {
		log.Println("clean msg list win:", vlo8.Count(), len(oldItem.msgitmdl))
		// i > 0 leave the QSpacerItem there // not need QSpacerItem anymore
		for i := vlo8.Count() - 1; i >= 0; i-- {
			itemv := vlo8.TakeAt(i)
			itemv.Widget().SetVisible(false)
		}
	}
	log.Println(time.Now().Sub(btime))
	log.Println("add msg list win:", len(item.msgitmdl), item.GetName())
	for _, msgiw := range item.msgitmdl {
		vlo8.Layout().AddWidget(msgiw.QWidget_PTR())
		msgiw.QWidget_PTR().SetVisible(true)
	}
	log.Println(time.Now().Sub(btime))
	// TODO too slow, 500ms+
}

///////////
func (this *MessageListWin) InitMessageListGesture() {
	this.InitContextMenu()
	w := uictx.mw.ScrollAreaWidgetContents_2
	// w.SetAttribute__(qtcore.Qt__WA_AcceptTouchEvents)
	this.gesto = NewMessageListGesture(w)
	this.gesto.OnLongTouch = this.OnSCWLongTouch
}

func (this *MessageListWin) InitContextMenu() {
	this.ctxmenu = qtwidgets.NewQMenu__()
	this.actcopy = this.ctxmenu.AddAction("&Copy Text")
	this.actcopyone = this.ctxmenu.AddAction("Copy &Message")
	this.actselall = this.ctxmenu.AddAction("Select &All")
	this.actselone = this.ctxmenu.AddAction("Select &One")
	this.actsep1 = this.ctxmenu.AddSeparator()
	this.actclear = this.ctxmenu.AddAction("C&lear")
	this.actquote = this.ctxmenu.AddAction("&Quote Selected")

	qtrt.Connect(this.actcopy, "triggered(bool)", this.ProcessActionCopy)
	qtrt.Connect(this.actselall, "triggered(bool)", this.ProcessActionSelectAll)
	qtrt.Connect(this.actclear, "triggered(bool)", this.ProcessActionClear)
	qtrt.Connect(this.actquote, "triggered(bool)", this.ProcessActionQuote)
}

func (this *MessageListWin) ProcessActionCopy() {
	uictx.qtapp.Clipboard().SetText__(this.selinfo.text)
}

func (this *MessageListWin) ProcessActionCopyOne() {
	// nick name + time + message
	if strings.HasPrefix(this.selinfo.objname, "QLabel") {
		// find parent, that's should be MessageItemView
	}
	if this.selinfo.objname == "MessageItemView" {
	}
	// uictx.qtapp.Clipboard().SetText__(this.selinfo.text)
}

func (this *MessageListWin) ProcessActionSelectAll() {
	// all list
}

func (this *MessageListWin) ProcessActionClear() {
	// all list
}

func (this *MessageListWin) ProcessActionQuote() {
	uictx.mw.LineEdit_2.SetText(fmt.Sprintf("> %s\n", this.selinfo.text))
}

// when
func (this *MessageListWin) ClearSelectInfo() {

}

// SCW = scroll content widget
// pos is global pos
func (this *MessageListWin) OnSCWLongTouch(pos *qtcore.QPointF) {
	ctw := uictx.mw.ScrollAreaWidgetContents_2
	mypos := ctw.MapFromGlobal(pos.ToPoint()) // ctw's cordinate pos
	chw := ctw.ChildAt_1(mypos)

	if chw != nil {
		log.Println(chw.ObjectName(), chw.MetaObject().ClassName(), mypos.X(), mypos.Y())
		if chw.MetaObject().ClassName() == "QLabel" {
			chl := qtwidgets.NewQLabelFromPointer(chw.GetCthis())
			this.selinfo.has = chl.HasSelectedText()
			this.selinfo.text = chl.SelectedText()
			this.selinfo.start = chl.SelectionStart()
			this.selinfo.objname = chw.ObjectName()
			this.selinfo.objwgt = chw
			this.ShowSCWContextMenu(pos.ToPoint()) // use global pos for show menu
		}
	} else {
		gopp.NilPrint(chw.GetCthis(), "Child at point is nil:", pos.X(), pos.Y())
	}
}

func (this *MessageListWin) ShowSCWContextMenu(pos *qtcore.QPoint) {
	// ctw := uictx.mw.ScrollAreaWidgetContents_2
	// gpos := ctw.MapToGlobal(pos)
	log.Printf("%#v\n", this.selinfo)
	this.actcopy.SetEnabled(this.selinfo.has)
	this.actquote.SetEnabled(this.selinfo.has)
	gpos := pos
	this.ctxmenu.Popup__(gpos)
}
