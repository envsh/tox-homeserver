package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"

	simplejson "github.com/bitly/go-simplejson"

	thscli "tox-homeserver/client"
	"tox-homeserver/gofia"
	"tox-homeserver/thspbs"
)

var appctx *gofia.AppContext
var vtcli *thscli.LigTox

// ui context
var uictx = NewUiContext()

type UiContext struct {
	qtapp   *qtwidgets.QApplication
	uiw     *Ui_MainWindow
	mw      *MainWindow
	ctitmdl []*RoomListItem
	msgwin  *MessageListWin
	iteman  *RoomListMan

	// 用于go的线程与qt主线程通知触发
	mech    *qtcore.QBuffer
	themeNo int
}

func NewUiContext() *UiContext {
	this := &UiContext{}
	this.ctitmdl = []*RoomListItem{}
	this.msgwin = NewMessageListWin()
	this.iteman = NewRoomListMan()

	this.mech = qtcore.NewQBuffer(nil)
	this.themeNo = _STL_SYSTEM
	return this
}

type MainWindow struct {
	*Ui_MainWindow

	roomCtxMenu        *qtwidgets.QMenu
	rcact1             *qtwidgets.QAction
	rcact2             *qtwidgets.QAction
	rcact3             *qtwidgets.QAction
	rcact4             *qtwidgets.QAction
	curRoomCtxMenuItem *RoomListItem
}

func NewMainWindow() *MainWindow {
	this := &MainWindow{}
	this.Ui_MainWindow = NewUi_MainWindow2()
	uictx.uiw = this.Ui_MainWindow

	this.init()
	return this
}

func (this *MainWindow) init() {
	this.initQml()
	this.connectSignals()
	this.switchUiStack(uictx.uiw.StackedWidget.CurrentIndex())
	this.Widget.SetStyleSheet(GetBg(_HEADER_BG))

	this.roomCtxMenu = qtwidgets.NewQMenu(nil)
	this.rcact1 = this.roomCtxMenu.AddAction("Leave Group")
	this.rcact2 = this.roomCtxMenu.AddAction("Remove Friend")
	this.rcact3 = this.roomCtxMenu.AddAction("View Info")
	this.rcact4 = this.roomCtxMenu.AddAction("PlaceHolder3")
	qtrt.Connect(this.rcact1, "triggered(bool)", func(checked bool) {
		this.onRoomContextTriggered(this.curRoomCtxMenuItem, checked, this.rcact1)
	})
	qtrt.Connect(this.rcact2, "triggered(bool)", func(checked bool) {
		this.onRoomContextTriggered(this.curRoomCtxMenuItem, checked, this.rcact2)
	})
	qtrt.Connect(this.rcact3, "triggered(bool)", func(checked bool) {
		this.onRoomContextTriggered(this.curRoomCtxMenuItem, checked, this.rcact3)
	})
	qtrt.Connect(this.rcact4, "triggered(bool)", func(checked bool) {
		this.onRoomContextTriggered(this.curRoomCtxMenuItem, checked, this.rcact4)
	})

	go initAppBackend()
}

func setAppStyleSheet() {
	bcc, err := []byte{}, error(nil)
	if gopp.IsAndroid() { // simple test
		bcc, err = ioutil.ReadFile("/sdcard/app.css")
	} else {
		bcc, err = ioutil.ReadFile("./app.css")
	}
	gopp.ErrPrint(err)
	if err != nil {
		fp := qtcore.NewQFile_1(":/app.css")
		fp.Open(qtcore.QIODevice__ReadOnly)
		bcc = []byte(qtcore.NewQIODeviceFromPointer(fp.GetCthis()).ReadAll().Data())
	}
	uictx.qtapp.SetStyleSheet(string(bcc))
}

func (this *MainWindow) connectSignals() {
	uiw := uictx.uiw

	qtrt.Connect(uiw.ToolButton_21, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		setAppStyleSheet()
	})
	stkw := uiw.StackedWidget
	qtrt.Connect(uiw.ToolButton_11, "clicked(bool)", func(checked bool) {
		cidx := stkw.CurrentIndex()
		if cidx > 0 {
			this.switchUiStack(cidx - 1)
		}
	})
	qtrt.Connect(uiw.ToolButton_12, "clicked(bool)", func(checked bool) {
		cidx := stkw.CurrentIndex()
		if cidx < stkw.Count()-1 {
			this.switchUiStack(cidx + 1)
		}
	})

	qtrt.Connect(uiw.ScrollArea_2.VerticalScrollBar(), "rangeChanged(int,int)", func(min int, max int) {
		log.Println(min, max)
		curpos := uiw.ScrollArea_2.VerticalScrollBar().Value()
		if ccstate.isBottom && curpos < max {
			uiw.ScrollArea_2.VerticalScrollBar().SetValue(max)
		}
		ccstate.maxpos = max
	})
	qtrt.Connect(uiw.ScrollArea_2.VerticalScrollBar(), "valueChanged(int)", func(value int) {
		ccstate.curpos = value
		maxval := ccstate.maxpos
		ccstate.isBottom = gopp.IfElse(value == maxval, true, false).(bool)
	})

	mech := uictx.mech
	mech.Open(qtcore.QIODevice__ReadWrite)
	qtrt.Connect(mech, "readyRead()", func() {
		log.Println("hehehehhee")
		mech.ReadAll()
		tryReadEvent()
	})
}

func (this *MainWindow) initQml() {
	qw := uictx.uiw.QuickWidget
	// qw.Engine().AddImportPath(":/qmlsys")
	qw.Engine().AddImportPath(":/qmlapp")
	qw.SetSource(qtcore.NewQUrl_1("qrc:/qmlapp/area.qml", 0))
	proot := qw.RootObject()
	gopp.NilPrint(proot, "qml root object nil")
}

func (this *MainWindow) switchUiStack(x int) {
	uictx.uiw.ComboBox.SetCurrentIndex(x)
	uictx.uiw.StackedWidget.SetCurrentIndex(x)
}

func (this *MainWindow) onRoomContextMenu(item *RoomListItem, w *qtwidgets.QWidget, pos *qtcore.QPoint) {
	this.curRoomCtxMenuItem = item
	if item.isgroup {
		this.rcact1.SetVisible(true)
		this.rcact2.SetVisible(false)
	} else {
		this.rcact1.SetVisible(false)
		this.rcact2.SetVisible(true)
	}
	this.roomCtxMenu.Popup(pos, nil)
}

func (this *MainWindow) onRoomContextTriggered(item *RoomListItem, checked bool, act *qtwidgets.QAction) {
	log.Println(item, checked, act.Text(), item.GetName(), item.GetId())
	if act == this.rcact1 {
		log.Println(item.grpInfo.GetGnum())
		vtcli.ConferenceDelete(item.grpInfo.GetGnum())
		uictx.iteman.Delete(item)
	} else if act == this.rcact2 {

	} else if act == this.rcact3 {

	}
}

func initAppBackend() {
	mech, uiw := uictx.mech, uictx.uiw

	gofia.AppOnCreate()
	appctx = gofia.GetAppCtx()
	vtcli = appctx.GetLigTox()
	vtcli.OnNewMsg = func() { mech.Write_1("5") }

	for {
		time.Sleep(500 * time.Millisecond)
		if vtcli.SelfGetAddress() != "" {
			break
		}
	}
	log.Println(vtcli.SelfGetAddress())

	uiw.Label_2.SetText(vtcli.SelfGetName())
	stmsg, _ := vtcli.SelfGetStatusMessage()
	uiw.Label_3.SetText(stmsg)
	uiw.ToolButton_17.SetToolTip(vtcli.SelfGetAddress())

	listw := uiw.ListWidget_2

	for fn, frnd := range vtcli.Binfo.Friends {
		itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
		listw.AddItem(itext)
		contactQueue <- frnd
	}

	for gn, grp := range vtcli.Binfo.Groups {
		itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
		listw.AddItem(itext)
		contactQueue <- grp
	}

	log.Println("get base info done.")
	baseInfoGot = true
	mech.Write_1("z")
	select {}
}

var baseInfoGot bool = false
var contactQueue = make(chan interface{}, 1234)

func tryReadEvent() {

	if !baseInfoGot {
		return
	}

	// 这个return不会节省cpu???
	if true {
		// return
	}

	tryReadContactEvent()
	tryReadMessageEvent()
}

func tryReadContactEvent() {

	for len(contactQueue) > 0 {
		contactx := <-contactQueue
		ctv := NewRoomListItem()
		ctv.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
			uictx.mw.onRoomContextMenu(ctv, w, pos)
		}

		uictx.iteman.addRoomItem(ctv)
		ctv.SetContactInfo(contactx)
		log.Println("add contact...", len(uictx.ctitmdl))
		if len(uictx.ctitmdl) == 1 {
			// ctv.SetPressState(true)
		}
	}
}

func tryReadMessageEvent() {
	for {
		bcc := vtcli.GetNextBackenEvent()
		if bcc == nil {
			break
		} else {
			jso, err := simplejson.NewJson(bcc)
			gopp.ErrPrint(err, jso)
			if err == nil {
				dispatchEvent(jso)
			}
		}
	}
}

func dispatchEvent(jso *simplejson.Json) {
	uiw, ctitmdl := uictx.uiw, uictx.ctitmdl
	// listwp1 := Ui_MainWindow_Get_listWidget(uiw)
	// listw1 := widgets.NewQListWidgetFromPointer(listwp1)

	evtName := jso.Get("name").MustString()
	switch evtName {
	case "SelfConnectionStatus":
	case "FriendRequest":
		///
		// pubkey := jso.Get("args").GetIndex(0).MustString()
		// _, err := appctx.store.AddFriend(pubkey, 0, "", "")
		// gopp.ErrPrint(err, jso.Get("args"))

	case "FriendMessage":
		// jso.Get("args").GetIndex(0).MustString()
		msg := jso.Get("args").GetIndex(1).MustString()
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _, _ = msg, fname, pubkey

		itext := fmt.Sprintf("%s: %s", fname, msg)
		uiw.ListWidget.AddItem(itext)
		uiw.ListWidget.ScrollToBottom()

		item := uictx.iteman.Get(pubkey)
		if item == nil {
			log.Println("wtf", fname, pubkey, msg)
		} else {
			msgo := NewMessageForFriend(jso)
			item.AddMessage(msgo)
		}

		///
		// _, err := appctx.store.AddFriendMessage(msg, pubkey)
		// gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _ = fname, pubkey

		// cfx, found := this.cfvs.Get(pubkey)
		/*
			cfsx, found := this.chatFormStates.Get(pubkey)
			if !found {
				log.Println("wtf, chat form view not found:", fname, pubkey)
			} else {
				cfs := cfsx.(*ChatFormState)
				cfs.status = uint32(gopp.MustInt(jso.Get("args").GetIndex(1).MustString()))
				this.signalProperView(cfs, true)
			}
		*/

	case "ConferenceInvite":
		groupNumber := jso.Get("margs").GetIndex(2).MustString()
		cookie := jso.Get("args").GetIndex(2).MustString()
		groupId := thscli.ConferenceCookieToIdentifier(cookie)
		log.Println(groupId)
		_ = groupNumber

		item := uictx.iteman.Get(groupId)
		if item == nil {
			item = NewRoomListItem()
			item.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
				uictx.mw.onRoomContextMenu(item, w, pos)
			}
			uictx.iteman.addRoomItem(item)
			grpInfo := &thspbs.GroupInfo{}
			grpInfo.GroupId = groupId
			grpInfo.Gnum = gopp.MustUint32(groupNumber)
			grpInfo.Title = fmt.Sprintf("Group #%s", groupNumber)
			item.SetContactInfo(grpInfo)
			log.Println("new group contact:", groupNumber, grpInfo.Title, groupId)
		}

		///
		// _, err := appctx.store.AddGroup(groupId, ctis.cnum, ctis.ctname)
		// gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupNumber := jso.Get("args").GetIndex(1).MustString()
		groupTitle := jso.Get("args").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(0).MustString()
		log.Println(groupId)
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}
		_ = groupTitle

		item := uictx.iteman.Get(groupId)
		if item != nil {
			item.UpdateName(groupTitle)
			log.Println("update group contact title:", groupNumber, groupId, groupTitle)
		} else {
			item = NewRoomListItem()
			item.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
				uictx.mw.onRoomContextMenu(item, w, pos)
			}
			uictx.iteman.addRoomItem(item)
			grpInfo := &thspbs.GroupInfo{}
			grpInfo.GroupId = groupId
			grpInfo.Gnum = gopp.MustUint32(groupNumber)
			grpInfo.Title = groupTitle
			item.SetContactInfo(grpInfo)
			log.Println("new group contact:", groupNumber, groupId, groupTitle)
		}

	case "ConferenceNameListChange":
		groupTitle := jso.Get("margs").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(3).MustString()
		log.Println(groupId)
		if thscli.ConferenceIdIsEmpty(groupId) {
			log.Println("empty")
			break
		}
		_ = groupTitle

		/*
			valuex, found := appctx.contactStates.Get(groupId)
			var ctis *ContactItemState
			if !found {
				ctis = newContactItemState()
				ctis.group = true
				appctx.contactStates.Put(groupId, ctis)
				log.Println("new group contact:", groupId)
			} else {
				ctis = valuex.(*ContactItemState)
			}
			if groupTitle != "" && groupTitle != ctis.ctname {
				ctis.ctname = groupTitle
				if appctx.app.Child == nil {
					InterBackRelay.Signal()
				}
			}
		*/

		///
		// peerPubkey := jso.Get("margs").GetIndex(1).MustString()
		// _, err := appctx.store.AddPeer(peerPubkey, 0)
		// gopp.ErrPrint(err)

	case "ConferenceMessage":
		groupId := jso.Get("margs").GetIndex(3).MustString()
		log.Println(groupId)
		if thscli.ConferenceIdIsEmpty(groupId) {
			break
		}

		message := jso.Get("args").GetIndex(3).MustString()
		peerName := jso.Get("margs").GetIndex(0).MustString()
		groupTitle := jso.Get("margs").GetIndex(2).MustString()

		itext := fmt.Sprintf("%s@%s: %s", peerName, groupTitle, message)
		uiw.ListWidget.AddItem(itext)
		uiw.ListWidget.ScrollToBottom()
		log.Println("item:", itext)

		ccstate.curpos = uiw.ScrollArea_2.VerticalScrollBar().Value()
		ccstate.maxpos = uiw.ScrollArea_2.VerticalScrollBar().Maximum()

		for _, room := range ctitmdl {
			log.Println(room.GetName(), ",", groupTitle, ",", room.GetId(), ",", groupId)
			if room.GetId() == groupId && room.GetName() == groupTitle {
				room.AddMessage(NewMessageForGroup(jso))
				break
			}
		}

	default:
	}
}
