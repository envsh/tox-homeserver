package main

import (
	"fmt"
	"gopp"
	"log"
	"strings"
	"time"

	thscom "tox-homeserver/common"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func (this *MainWindow) initMessageListWin() {
	this.initMessageListUi()
	this.initMessageListSignals()
	this.initMessageListEvents()
}

func (this *MainWindow) initMessageListUi() {
}

func (this *MainWindow) initMessageListSignals() {
}

func (this *MainWindow) initMessageListEvents() {
	le := this.LineEdit_2
	le.InheritKeyPressEvent(func(arg0 *qtgui.QKeyEvent) {
		arg0.Ignore()
		// Ctrl+C
		log.Println(arg0.Matches(qtgui.QKeySequence__Paste))
		if arg0.Matches(qtgui.QKeySequence__Paste) {
			this.checkClipboardImage()
		}
		le.KeyPressEvent(arg0)
	})
	le.InheritMousePressEvent(func(arg0 *qtgui.QMouseEvent) {
		// middle button
		if arg0.Button() == qtcore.Qt__MiddleButton {
			this.checkClipboardImage()
			arg0.Accept()
		} else {
			arg0.Ignore()
			le.MousePressEvent(arg0)
		}
	})
}

func (this *MainWindow) checkClipboardImage() {
	cb := uictx.qtapp.Clipboard()
	md := cb.MimeData__()
	fmtlst := md.Formats()
	fmtlstx := qtcore.NewQStringListxFromPointer(fmtlst.GetCthis())
	fmtnames := fmtlstx.ConvertToSlice()
	log.Println(fmtnames)
	if md.HasFormat("application/x-qt-image") {
		imgv := md.ImageData()
		log.Println(imgv.Type(), imgv.TypeName())
		// since qtcore.QVariant has no ToImage wrapper
		ba := md.Data("application/x-qt-image")
		log.Println(ba.Length())
		rawData := ba.Data_fix()
		this.sendFileData([]byte(rawData))
	}
}

// 聊天输入框显示规则，默认显示一行，当输入超过一行时自动扩展输入框高度。最多显示3-5行，然后出现滚动条。
// 使用Enter发送消息，使用Shift+Entry换行。
// 在回退的时候，也需要减小高度，直到默认高度
func (this *MainWindow) checkExpandLineEditHeight() {

}

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
		cnt := len(item.grpInfo.Members)
		mw.Label_6.SetText(fmt.Sprintf("%d users in chat", cnt))
	} else {
		mw.Label_6.SetVisible(false)
		mw.Label_7.SetVisible(true)
		mw.Label_7.SetText(gopp.StrSuf4ui(item.frndInfo.GetStmsg(), 26))
		mw.Label_7.SetToolTip(item.frndInfo.GetStmsg())
		// SetQLabelElideText(mw.Label_7, item.frndInfo.GetStmsg()) // TODO
	}
	if item.cticon != nil {
		mw.Label_4.SetPixmap(item.cticon.Pixmap_1_(32, 32))
	}
	enableFriend := !item.isgroup && !thscom.IsFixedSpecialContact(item.GetNum())
	mw.updateRoomOpMenu(enableFriend)
	mw.LabelMsgCount2.SetText(fmt.Sprintf("%3d", item.totalCount))
	mw.LabelMsgCount.SetText(fmt.Sprintf("%3d", item.totalCount))

	this.ReloadMessages(oldItem)
}

func (this *MessageListWin) ReloadMessages(oldItem *RoomListItem) {
	item := this.item

	vlo8 := uictx.uiw.VerticalLayout_3

	if oldItem != nil {
		btime := time.Now()
		// i > 0 leave the QSpacerItem there // not need QSpacerItem anymore
		for i := vlo8.Count() - 1; i >= 0; i-- {
			itemv := vlo8.TakeAt(i)
			itemv.Widget().SetVisible(false)
		}
		log.Println("Clean done, used:", time.Since(btime), vlo8.Count(), item.GetName())
	}

	btime := time.Now()
	for _, msgiw := range item.msgitmdl {
		vlo8.Layout().AddWidget(msgiw.QWidget_PTR())
		msgiw.QWidget_PTR().SetVisible(true)
	}
	log.Println("Add to msg list win, used:", time.Since(btime), len(item.msgitmdl), item.GetName())
	// TODO too slow, 500ms+
}

func (this *MessageListWin) SetIconForItem(item *RoomListItem) {
	if item != this.item {
		return
	}
	mw := uictx.mw
	mw.Label_4.SetPixmap(item.cticon.Pixmap_1_(32, 32))
}

func (this *MessageListWin) ClearAll() {
	this.ClearItemInfos()
	this.ClearMessages()
}
func (this *MessageListWin) ClearItemInfos() {
	item := this.item
	if item == nil {
		return
	}
	mw := uictx.mw
	mw.Label_5.Clear()
	mw.Label_5.SetToolTip("")
	mw.Label_6.Clear()
	mw.Label_6.SetToolTip("")
	mw.Label_7.Clear()
	mw.Label_7.SetToolTip("")
	mw.LabelMsgCount.Clear()
	mw.LabelMsgCount.SetToolTip("")
	mw.LabelMsgCount2.Clear()
	mw.LabelMsgCount2.SetToolTip("")
	mw.Label_4.Clear()
	mw.Label_4.SetToolTip("")
}
func (this *MessageListWin) ClearMessages() {
	item := this.item
	if item == nil {
		return
	}

	vlo8 := uictx.uiw.VerticalLayout_3

	btime := time.Now()
	elemcnt := vlo8.Count()
	// i > 0 leave the QSpacerItem there // not need QSpacerItem anymore
	for i := elemcnt - 1; i >= 0; i-- {
		itemv := vlo8.TakeAt(i)
		_ = itemv
		// itemv.Widget().SetVisible(t)
		// itemv.Widget().DeleteLater()
		// qtwidgets.DeleteQWidget(itemv.Widget())
	}
	log.Println("Clean done, used:", time.Since(btime), vlo8.Count(), elemcnt, item.GetName())
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

func tr(string, string) {}

func (this *MessageListWin) SetPeerCount(n int) {
	mw := uictx.mw
	mw.Label_6.SetText(fmt.Sprintf("%d users in chat", n))
	// QObject::tr("ccc", "dummy123")
}
