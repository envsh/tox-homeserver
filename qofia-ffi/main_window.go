package main

// really is main_window_ctrl.go

import (
	"encoding/json"
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"strings"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtmultimedia"
	"github.com/kitech/qt.go/qtwidgets"

	"tox-homeserver/avhlp"
	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	"github.com/kitech/qt.go/qtandroidextras"
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
	// qtrt.SetDebugFFICall(false)
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
	this.initMessageListWin()
	this.initInivteFriend()
	this.initAddFriend()
	this.initGroupMemberList()
	this.initContactInfoPage()
	this.initRoomFile()
	this.initVideoCall()
	this.initAboutPage()
	log.Println("Init startup ui done.")
}

func (this *MainWindow) initFirstShow() {
	log.Println("Init first show ui...")
	// FindProperFontFile(this.LineEdit.QWidget_PTR())
	PrepareFont()
	this.loadStorage()

	this.initProfileStorage()
	this.initRoomFileStorage()
	// this.initOtherStorage()
	go _CheckIntentMessage()
	log.Println("Init first show ui done.")

	// tstvcap = avhlp.NewVideoRecorder2Auto(func(d []byte, w uint16, h uint16) {
	//	log.Println(len(d), w, h)
	// })
	if gopp.IsAndroid() {
		jvm := qtandroidextras.QAndroidJniEnvironment_JavaVM()
		actx := qtandroidextras.AndroidContext().Object()
		avhlp.SetCurrentVM(jvm, actx)
	}
	cami := qtmultimedia.QCameraInfo_DefaultCamera()
	log.Println(cami.DeviceName(), cami.IsNull())
	camis := qtmultimedia.QCameraInfo_AvailableCameras(qtmultimedia.QCamera__BackFace)
	log.Println(camis.Count_1())
}

var tstvcap *avhlp.VideoRecorder2Auto

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
	UIST_QMLMCTRL = iota + 0
	UIST_QMLORIGIN
	UIST_SETTINGS
	UIST_LOGINUI
	UIST_CONTACTUI
	UIST_MESSAGEUI
	UIST_VIDEOUI
	// UIST_PICKCALLUI // TODO video
	UIST_ADD_GROUP
	UIST_ADD_FRIEND
	UIST_INVITE_FRIEND
	UIST_MEMBERS
	UIST_CONTACT_INFO
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

// pop
func (this *MainWindow) switchUiStackPopBack() {
	if _, ok := _HeaderFooterState.viewStack.Pop(); ok {
		valx, _ := _HeaderFooterState.viewStack.Peek()
		this.switchUiStackPop(valx.(int))
	}
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
	log.Println(checked, act.Text(), item.GetName(), item.GetId(), item)
	if act == this.rcactOpen {
		uictx.mw.switchUiStack(UIST_MESSAGEUI)
		uictx.msgwin.SetRoom(item)
	} else if act == this.rcact1 {
		log.Println("quiting group...", item.grpInfo.GetGnum(), item.GetName(), item.GetId())
		AVMan().RemoveSession(item.GetId(), item.GetName())
		vtcli.ConferenceDelete(item.grpInfo.GetGnum())
		uictx.iteman.Delete(item)
		vtcli.Binfo.RemoveGroup(item.grpInfo.GetGnum())
	} else if act == this.rcact2 {
		_, err := vtcli.FriendDelete(item.frndInfo.GetFnum())
		gopp.ErrPrint(err, item.frndInfo.GetFnum())
		if err == nil {
			uictx.iteman.Delete(item)
		}
	} else if act == this.rcact3 {
		this.switchUiStack(UIST_CONTACT_INFO)
		this.fillConactInfo(item)
		this.fillConactSetting(item)
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
	itext := uiw.TextEdit_3.Document().ToPlainText()
	itext = strings.Trim(itext, " ")
	if len(itext) > thscom.MaxMessageLen {
		ShowToast("Message too long", 1)
		return
	}
	if len(itext) == 0 {
		ShowToast("Empty message", 1)
		return
	}
	item := uictx.msgwin.item
	if item != nil {
		this.sendMessageImpl(item, itext, item.isgroup, item.GetNum())
		uiw.TextEdit_3.Clear()
	} else {
		log.Println("Message not send:", len(itext), item == nil)
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
	alldone := false
	// procn := 0
	for !alldone {
		done1 := tryReadUifnEvent()

		if !baseInfoGot {
			log.Println("baseInfoGot is not set, not need other works.")
			return
		}

		done2 := tryReadContactEvent()
		done3 := tryReadMessageEvent()
		done4 := tryRecvIntentMessageEvent()

		alldone = done1 && done2 && done3 && done4
	}
}

// return done
func tryReadUifnEvent() bool {
	procn := 0
	for len(uifnQueue) > 0 && procn < 20 {
		uifn := <-uifnQueue
		uifn()
		procn++
	}
	return len(uifnQueue) == 0
}

func tryReadContactEvent() bool {
	procn := 0
	for len(contactQueue) > 0 && procn < 20 {
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
		procn++
	}
	return len(contactQueue) == 0
}

func tryReadMessageEvent() bool {
	procn := 0
	for procn < 20 {
		bcc := vtcli.GetNextBackenEvent()
		if bcc == nil {
			return true
		} else {
			evto := &thspbs.Event{}
			err := json.Unmarshal(bcc, evto)
			gopp.ErrPrint(err)
			if err == nil {
				dispatchEvent(evto)
				dispatchEventResp(evto)
			}
		}
		procn++
	}
	return false
}

func tryRecvIntentMessageEvent() bool {
	procn := 0
	for len(intentQueue) > 0 && procn < 20 {
		evto := <-intentQueue
		dispatchOtherEvent(evto)
		procn++
	}
	return len(intentQueue) == 0
}
