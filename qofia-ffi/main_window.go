package main

// really is main_window_ctrl.go

import (
	"encoding/json"
	"fmt"
	"gopp"
	"io/ioutil"
	"log"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"

	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
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
	this.msgwin.InitMessageListGesture()
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

	this.initStartup()
	// this.initFirstShow() // TODO really first show
	return this
}

func (this *MainWindow) initStartup() {
	log.Println("Init startup ui...")
	this.initMainWindow()
	this.initHeaderFooter()
	this.initProfile()
	this.initRoomChat()
	this.initInivteFriend()
	this.initAddFriend()
	this.initGroupMemberList()
	this.initRoomFile()
	log.Println("Init startup ui done.")
}

func (this *MainWindow) initFirstShow() {
	log.Println("Init first show ui...")
	this.loadStorage()

	this.initProfileStorage()
	this.initRoomFileStorage()
	// this.initOtherStorage()
	go _CheckIntentMessage()
	log.Println("Init first show ui done.")
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

var create_room_dlg *Ui_Dialog
var add_friend_dlg *Ui_AddFriendDialog

const (
	UIST_QMLMCTRL = iota + 1
	UIST_QMLORIGIN
	UIST_SETTINGS
	UIST_LOGINUI
	UIST_CONTACTUI
	UIST_MESSAGEUI
	UIST_VIDEOUI
	UIST_PICKCALLUI // TODO video
	UIST_ADD_GROUP
	UIST_ADD_FRIEND
	UIST_INVITE_FRIEND
	UIST_MEMBERS
	UIST_TESTUI
	UIST_LOGUI
	UIST_ABOUTUI
)

// push
func (this *MainWindow) switchUiStack(x int) {
	_HeaderFooterState.viewStack.Push(x)
	uictx.uiw.ComboBox.SetCurrentIndex(x)
	uictx.uiw.StackedWidget.SetCurrentIndex(x)
}

// pop
func (this *MainWindow) switchUiStackPop(x int) {
	// _HeaderFooterState.viewStack.Pop()
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
		_, err := vtcli.FriendDelete(item.frndInfo.GetFnum())
		gopp.ErrPrint(err, item.frndInfo.GetFnum())
		if err == nil {
			uictx.iteman.Delete(item)
		}
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
	if len(itext) > thscom.MaxMessageLen {
		ShowToast("Message too long", 1)
		return
	}
	item := uictx.msgwin.item
	if item != nil && len(itext) > 0 {
		this.sendMessageImpl(item, itext, item.isgroup, item.GetNum())
		uiw.LineEdit_2.Clear()
	} else {
		log.Println("not send:", len(itext), item)
	}
}

func (this *MainWindow) sendMessageImpl(item *RoomListItem, itext string, isgroup bool, ctnum uint32) *Message {
	userCode := thscli.NextUserCode(devInfo.Uuid)
	msgo := NewMessageForMe(itext)
	msgo.UserCode = userCode
	item.AddMessage(msgo, false)

	if isgroup {
		vtcli.ConferenceSendMessage(ctnum, 0, itext, userCode)
	} else {
		vtcli.FriendSendMessage(ctnum, itext, userCode)
	}
	return msgo
}

var baseInfoGot bool = false
var contactQueue = make(chan interface{}, 1234)
var uifnQueue = make(chan func(), 1234)
var intentQueue = make(chan *thspbs.Event, 123) // for android intent message

func runOnUiThread(fn func()) {
	uifnQueue <- fn
	uictx.mech.Trigger()
}

// recv left and return
func tryReadEvent() {
	tryReadUifnEvent()

	if !baseInfoGot {
		log.Println("baseInfoGot is not set, not need other works.")
		return
	}

	tryReadContactEvent()
	tryReadMessageEvent()
	tryRecvIntentMessageEvent()
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
		if uictx.iteman.Get(ctv.GetId()) != nil {
			log.Println("Already in list:", ctv.GetName(), ctv.GetId())
			continue
		}
		uictx.iteman.addRoomItem(ctv)

		ctv.OnConextMenu = func(w *qtwidgets.QWidget, pos *qtcore.QPoint) {
			uictx.mw.onRoomContextMenu(ctv, w, pos)
		}
		ctv.timeline = thscli.TimeLine{NextBatch: vtcli.Binfo.NextBatch, PrevBatch: vtcli.Binfo.NextBatch - 1}
		ctv.SetContactInfo(contactx)

		log.Println("add contact", gopp.IfElseStr(ctv.isgroup, "group", "friend"), "...", len(uictx.ctitmdl))
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
			evto := &thspbs.Event{}
			err := json.Unmarshal(bcc, evto)
			gopp.ErrPrint(err)
			if err == nil {
				dispatchEvent(evto)
				dispatchEventResp(evto)
			}
		}
	}
}

func tryRecvIntentMessageEvent() {
	for len(intentQueue) > 0 {
		evto := <-intentQueue
		dispatchOtherEvent(evto)
	}
}
