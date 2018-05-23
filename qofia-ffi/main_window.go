package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"runtime"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"

	simplejson "github.com/bitly/go-simplejson"

	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"
)

var appctx *thscli.AppContext
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
	gtreco  *RoomGestureReco

	// 用于go的线程与qt主线程通知触发
	mech *Notifier
	// metmer  *qtcore.QTimer
	themeNo int
}

func NewUiContext() *UiContext {
	this := &UiContext{}
	this.ctitmdl = []*RoomListItem{}
	this.msgwin = NewMessageListWin()
	this.iteman = NewRoomListMan()
	this.gtreco = NewRoomGestureReco()

	// this.metmer = qtcore.NewQTimer(nil)
	this.themeNo = _STL_SYSTEM
	return this
}

func (this *UiContext) init() *UiContext {
	this.gtreco.OnClick = this.iteman.onClicked
	this.gtreco.OnLongTouch = this.iteman.onLongTouched
	return this
}

type MainWindow struct {
	*Ui_MainWindow

	roomCtxMenu        *qtwidgets.QMenu
	rcactOpen          *qtwidgets.QAction
	rcact1             *qtwidgets.QAction
	rcact2             *qtwidgets.QAction
	rcact3             *qtwidgets.QAction
	rcactAddGroup      *qtwidgets.QAction
	rcactAddFriend     *qtwidgets.QAction
	rcactInviteFriend  *qtwidgets.QAction
	rcact4             *qtwidgets.QAction
	curRoomCtxMenuItem *RoomListItem

	RoomChatState

	sticon *qtgui.QIcon
}

func NewMainWindow() *MainWindow {
	qtrt.SetDebugFFICall(false)
	this := &MainWindow{}
	this.Ui_MainWindow = NewUi_MainWindow2()
	uictx.uiw = this.Ui_MainWindow

	this.init()
	return this
}

func (this *MainWindow) init() {
	this.initMainWin()
	this.initRoomChat()
	this.initInivteFriend()
}

func (this *MainWindow) initMainUi() {
	this.setConnStatus(false)
}

func (this *MainWindow) initMainWin() {
	this.initMainUi()
	this.initQml()
	this.connectSignals()
	this.switchUiStack(uictx.uiw.StackedWidget.CurrentIndex())
	// this.Widget.SetStyleSheet(GetBg(_HEADER_BG))

	this.roomCtxMenu = qtwidgets.NewQMenu(nil)
	this.rcactOpen = this.roomCtxMenu.AddAction("Open Chat")
	this.rcact1 = this.roomCtxMenu.AddAction("Leave Group")
	this.rcact2 = this.roomCtxMenu.AddAction("Remove Friend")
	this.rcact3 = this.roomCtxMenu.AddAction("View Info")
	this.rcactAddGroup = this.roomCtxMenu.AddAction("Create Group")
	this.rcactAddFriend = this.roomCtxMenu.AddAction("Add Friends")
	this.rcactInviteFriend = this.roomCtxMenu.AddAction("Invite Friends")
	this.rcact4 = this.roomCtxMenu.AddAction("PlaceHolder3")

	if runtime.GOOS == "android" {
		sz := qtcore.NewQSize_1(32, 32)
		this.ToolButton_11.SetIconSize(sz)
		this.ToolButton_12.SetIconSize(sz)
		this.ToolButton_19.SetIconSize(sz)
		this.ToolButton_20.SetIconSize(sz)
		this.ToolButton_21.SetIconSize(sz)
		this.ToolButton_4.SetIconSize(sz)
		this.ToolButton_5.SetIconSize(sz)
		this.ToolButton_6.SetIconSize(sz)
		this.ToolButton_7.SetIconSize(sz)
	}

	qtrt.Connect(this.rcactOpen, "triggered(bool)", func(checked bool) {
		this.onRoomContextTriggered(this.curRoomCtxMenuItem, checked, this.rcactOpen)
	})

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

	//
	log.Println("Has scroller:", qtwidgets.QScroller_HasScroller(uictx.uiw.ScrollArea),
		qtwidgets.QScroller_HasScroller(uictx.uiw.ScrollArea_2))
	uictx.uiw.ScrollArea.GrabGesture(qtcore.Qt__SwipeGesture, 0)
	uictx.uiw.ScrollArea.GrabGesture(qtcore.Qt__PanGesture, 0)
	uictx.uiw.ScrollArea.GrabGesture(qtcore.Qt__PinchGesture, 0)
	qtwidgets.QScroller_GrabGesture(uictx.uiw.ScrollArea, qtwidgets.QScroller__LeftMouseButtonGesture)
	qtwidgets.QScroller_GrabGesture(uictx.uiw.ScrollArea_2, qtwidgets.QScroller__LeftMouseButtonGesture)
	log.Println("Has scroller:", qtwidgets.QScroller_HasScroller(uictx.uiw.ScrollArea),
		qtwidgets.QScroller_HasScroller(uictx.uiw.ScrollArea_2))

	/*
		uictx.uiw.ScrollArea.InheritEvent(func(event *qtcore.QEvent) bool {
			log.Println(event.Type())
			// return false
			return uictx.uiw.ScrollArea.Event(event)
		})
	*/

	go this.initAppBackend()

}

func setAppStyleSheet() {
	bcc, err := []byte{}, error(nil)
	if gopp.IsAndroid() { // simple test
		bcc, err = ioutil.ReadFile("/sdcard/apptst.css")
	} else {
		bcc, err = ioutil.ReadFile("./theme/apptst.css")
	}
	gopp.ErrPrint(err)
	if err != nil {
		fp := qtcore.NewQFile_1(":/theme/apptst.css")
		fp.Open(qtcore.QIODevice__ReadOnly)
		bcc = []byte(qtcore.NewQIODeviceFromPointer(fp.GetCthis()).ReadAll().Data())
		qtcore.NewQIODeviceFromPointer(fp.GetCthis()).Close()
	}
	uictx.qtapp.SetStyleSheet(string(bcc))
}

func setAppStyleSheetTheme(index int) {

	fp := qtcore.NewQFile_1(fmt.Sprintf(":/theme/%s.css", styleSheets[index]))
	fp.Open(qtcore.QIODevice__ReadOnly)
	bcc := []byte(qtcore.NewQIODeviceFromPointer(fp.GetCthis()).ReadAll().Data())
	qtcore.NewQIODeviceFromPointer(fp.GetCthis()).Close()

	uictx.qtapp.SetStyleSheet(string(bcc))
}

func (this *MainWindow) connectSignals() {
	uiw := uictx.uiw

	qtrt.Connect(uiw.ToolButton_19, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		ShowToast("hehehh哈哈eehhe", 1)
	})
	qtrt.Connect(uiw.ToolButton_20, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		// testRunOnAndroidThread()
		KeepScreenOn(true)
	})
	qtrt.Connect(uiw.ToolButton_21, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		setAppStyleSheet()
	})

	qtrt.Connect(uiw.ToolButton_23, "clicked(bool)", func(checked bool) {
		log.Println(checked, uictx.msgwin.item == nil)
		if uictx.msgwin.item != nil {
			go func() {
				hisfet.PullPrevHistoryByRoom(uictx.msgwin.item)
			}()
		}
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

	uictx.mech = NewNotifier(func() { tryReadEvent() })

	// send message button
	qtrt.Connect(uiw.ToolButton_18, "clicked(bool)", func(checked bool) {
		this.sendMessage()
	})
	qtrt.Connect(uiw.LineEdit_2, "returnPressed()", this.sendMessage)

	// switch theme
	qtrt.Connect(uiw.ComboBox_2, "currentIndexChanged(int)", func(index int) {
		setAppStyleSheetTheme(index)
	})

	//
	qtrt.Connect(uiw.ToolButton_4, "clicked(bool)", func(bool) {
		if add_friend_dlg == nil {
			add_friend_dlg = NewUi_AddFriendDialog2()
			qtrt.ConnectSlot(add_friend_dlg.ButtonBox, "accepted()", add_friend_dlg.AddFriendDialog, "accept()")
			qtrt.ConnectSlot(add_friend_dlg.ButtonBox, "rejected()", add_friend_dlg.AddFriendDialog, "reject()")
			setAppStyleSheet()
		} else {
			add_friend_dlg.LineEdit.Clear()
		}

		log.Println(add_friend_dlg.ButtonBox.GetCthis(), add_friend_dlg.AddFriendDialog.GetCthis())
		r := add_friend_dlg.AddFriendDialog.Exec()
		log.Println(r, qtwidgets.QDialog__Accepted, qtwidgets.QDialog__Rejected)
		if r == qtwidgets.QDialog__Accepted {
			log.Println(add_friend_dlg.LineEdit.Text())
		}
	})
	qtrt.Connect(uiw.ToolButton_5, "clicked(bool)", func(bool) {
		if create_room_dlg == nil {
			create_room_dlg = NewUi_Dialog2()
			qtrt.ConnectSlot(create_room_dlg.ButtonBox, "accepted()", create_room_dlg.Dialog, "accept()")
			qtrt.ConnectSlot(create_room_dlg.ButtonBox, "rejected()", create_room_dlg.Dialog, "reject()")
			setAppStyleSheet()
		} else {
			create_room_dlg.LineEdit.Clear()
		}

		log.Println(create_room_dlg.ButtonBox.GetCthis(), create_room_dlg.Dialog.GetCthis())
		r := create_room_dlg.Dialog.Exec()
		log.Println(r, qtwidgets.QDialog__Accepted, qtwidgets.QDialog__Rejected)
		if r == qtwidgets.QDialog__Accepted {
			name := create_room_dlg.LineEdit.Text()
			log.Println(create_room_dlg.LineEdit.Text())
			go func() {
				vtcli := appctx.GetLigTox()
				gn, id, err := vtcli.ConferenceNew(name)
				gopp.ErrPrint(err, name)
				log.Println("Group created:", gn, name)
				grp := &thspbs.GroupInfo{}
				grp.Gnum = gn
				grp.Mtype = 0
				grp.Title = name
				grp.GroupId = id
				contactQueue <- grp
				uictx.mech.Trigger()
			}()
		}
	})
	qtrt.Connect(uiw.ToolButton_7, "clicked(bool)", func(bool) { this.switchUiStack(UIST_SETTINGS) })
}

var create_room_dlg *Ui_Dialog
var add_friend_dlg *Ui_AddFriendDialog

func (this *MainWindow) initQml() {
	qw := uictx.uiw.QuickWidget
	// qw.Engine().AddImportPath(":/qmlsys")
	qw.Engine().AddImportPath(":/qmlapp")
	qw.SetSource(qtcore.NewQUrl_1("qrc:/qmlapp/area.qml", 0))
	proot := qw.RootObject()
	gopp.NilPrint(proot, "qml root object nil")
}

const (
	UIST_QMLMCTRL      = 0
	UIST_QMLORIGIN     = 1
	UIST_SETTINGS      = 2
	UIST_MAINUI        = 3
	UIST_MESSAGEUI     = 4
	UIST_ADD_GROUp     = 5
	UIST_ADD_FRIEND    = 6
	UIST_INVITE_FRIEND = 7
	UIST_TESTUI        = 8
	UIST_LOGUI         = 9
)

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
	if act == this.rcactOpen {
		uictx.mw.switchUiStack(UIST_MESSAGEUI)
		uictx.msgwin.SetRoom(item)
	} else if act == this.rcact1 {
		log.Println(item.grpInfo.GetGnum())
		vtcli.ConferenceDelete(item.grpInfo.GetGnum())
		uictx.iteman.Delete(item)
	} else if act == this.rcact2 {

	} else if act == this.rcact3 {

	}
}

func (this *MainWindow) setConnStatus(on bool) {
	if on {
		this.sticon = qtgui.NewQIcon_2(":/icons/online_30.png")
		this.ToolButton.SetIcon(this.sticon)
		this.ToolButton.SetToolTip("online")
	} else {
		this.sticon = qtgui.NewQIcon_2(":/icons/offline_30.png")
		this.ToolButton.SetIcon(this.sticon)
		this.ToolButton.SetToolTip("offline")
	}
}

func (this *MainWindow) sendMessage() {
	uiw := uictx.uiw
	itext := uiw.LineEdit_2.Text()
	item := uictx.msgwin.item
	if item != nil && len(itext) > 0 {
		if item.isgroup {
			vtcli.ConferenceSendMessage(item.grpInfo.Gnum, 0, itext)
		} else {
			vtcli.FriendSendMessage(item.frndInfo.Fnum, itext)
		}
		uiw.LineEdit_2.Clear()
		msgo := NewMessageForMe(itext)
		log.Println(msgo)
		item.AddMessage(msgo, false)
	} else {
		log.Println("not send:", len(itext), item)
	}
}

// should block
func (this *MainWindow) initAppBackend() {
	mech, uiw := uictx.mech, uictx.uiw

	thscli.AppOnCreate()
	appctx = thscli.GetAppCtx()
	vtcli = appctx.GetLigTox()
	vtcli.OnNewMsg = func() { mech.Trigger() }

	for {
		if vtcli.SelfGetAddress() != "" {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	log.Println(vtcli.SelfGetAddress())

	uiw.Label_2.SetText(vtcli.SelfGetName())
	stmsg, _ := vtcli.SelfGetStatusMessage()
	uiw.Label_3.SetText(stmsg)
	uiw.ToolButton_17.SetToolTip(vtcli.SelfGetAddress())

	listw := uiw.ListWidget_2

	for fn, frnd := range vtcli.Binfo.Friends {
		itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
		_ = itext
		listw.AddItem(itext)
		contactQueue <- frnd
	}

	for gn, grp := range vtcli.Binfo.Groups {
		itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
		_ = itext
		listw.AddItem(itext)
		contactQueue <- grp
	}

	uifnQueue <- func() { this.setConnStatus(gopp.IfElse(vtcli.Binfo.ConnStatus > 0, true, false).(bool)) }

	log.Println("get base info done.")
	baseInfoGot = true

	mech.Trigger()

	// 加载每个房间的最新消息, force schedue, or contact maybe not show in ui
	go func() {
		btime := time.Now()
		log.Println("Waiting contacts show on UI...") // about 31ms with 7 contacts
		for len(contactQueue) > 0 {
			time.Sleep(10 * time.Millisecond)
		}
		log.Println("Show base contacts on UI done.", time.Since(btime))
		pullAllRoomsLatestMessages()
	}()
	select {}
}

var baseInfoGot bool = false
var contactQueue = make(chan interface{}, 1234)
var uifnQueue = make(chan func(), 1234)

func runOnUiThread(fn func()) { uifnQueue <- fn }

func tryReadEvent() {

	if !baseInfoGot {
		log.Println("wtf")
		return
	}

	// 这个return不会节省cpu???
	if true {
		// return
	}

	tryReadUifnEvent()
	tryReadContactEvent()
	tryReadMessageEvent()
}

func tryReadUifnEvent() {
	for len(uifnQueue) > 0 {
		uifn := <-uifnQueue
		uifn()
	}
}

func tryReadContactEvent() {

	for len(contactQueue) > 0 {
		contactx := <-contactQueue
		ctv := NewRoomListItem()
		ctv.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
			uictx.mw.onRoomContextMenu(ctv, w, pos)
		}
		ctv.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}

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
	case "SelfConnectionStatus": // {"name":"SelfConnectionStatus","args":["2"],"margs":["CONNECTION_UDP"]}
		status := gopp.MustUint32(jso.Get("args").GetIndex(0).MustString())
		uictx.mw.setConnStatus(status > 0)
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
			item.AddMessage(msgo, false)
		}

		///
		// _, err := appctx.store.AddFriendMessage(msg, pubkey)
		// gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		_, _ = fname, pubkey
		st := gopp.MustInt(jso.Get("args").GetIndex(1).MustString())

		item := uictx.iteman.Get(pubkey)
		if item != nil {
			item.setConnStatus(int32(st))
		} else {
			log.Println("item not found:", fname, pubkey)
		}

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
			item.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}
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
			item.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}
			uictx.iteman.addRoomItem(item)
			grpInfo := &thspbs.GroupInfo{}
			grpInfo.GroupId = groupId
			grpInfo.Gnum = gopp.MustUint32(groupNumber)
			grpInfo.Title = groupTitle
			item.SetContactInfo(grpInfo)
			log.Println("new group contact:", groupNumber, groupId, groupTitle)
		}
	case "ConferencePeerName":
		log.Println("TODO", jso)
	case "ConferencePeerListChange":
		log.Println("TODO", jso)
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

		// raw message show area
		itext := fmt.Sprintf("%s@%s: %s", peerName, groupTitle, message)
		uiw.ListWidget.AddItem(itext)
		uiw.ListWidget.ScrollToBottom()
		// log.Println("item:", itext)

		ccstate.curpos = uiw.ScrollArea_2.VerticalScrollBar().Value()
		ccstate.maxpos = uiw.ScrollArea_2.VerticalScrollBar().Maximum()

		for _, room := range ctitmdl {
			// log.Println(room.GetName(), ",", groupTitle, ",", room.GetId(), ",", groupId)
			if room.GetId() == groupId && room.GetName() == groupTitle {
				room.AddMessage(NewMessageForGroup(jso), false)
				break
			}
		}

	case "FriendSendMessage":
		pubkey := jso.Get("args").GetIndex(2).MustString()
		itext := jso.Get("args").GetIndex(1).MustString()

		found := false
		for _, room := range ctitmdl {
			if room.GetId() == pubkey {
				msgo := NewMessageForMe(itext)
				room.AddMessage(msgo, false)
				found = true
				break
			}
		}
		log.Println(found, pubkey, itext)

	case "ConferenceSendMessage":
		groupId := jso.Get("args").GetIndex(3).MustString()
		itext := jso.Get("args").GetIndex(2).MustString()
		groupTitle := jso.Get("args").GetIndex(3).MustString()

		found := false
		for _, room := range ctitmdl {
			if room.GetId() == groupId && room.GetName() == groupTitle {
				msgo := NewMessageForMe(itext)
				room.AddMessage(msgo, false)
				found = true
				break
			}
		}
		log.Println(found, groupId, itext)

	default:
		log.Println(jso)
	}
}
