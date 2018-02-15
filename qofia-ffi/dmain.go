package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"

	simplejson "github.com/bitly/go-simplejson"

	thscli "tox-homeserver/client"
	"tox-homeserver/gofia"
)

var appctx *gofia.AppContext
var vtcli *thscli.LigTox

var uiw *Ui_MainWindow
var msgitmdl = []*Ui_MessageItemView{}

// var dyslot *qtrt.QDynSlotObject
var mech *qtcore.QBuffer

func DumpCallers(pcs []uintptr) {
	log.Println("DumpCallers...", len(pcs))
	for idx, pc := range pcs {
		pcfn := runtime.FuncForPC(pc)
		file, line := pcfn.FileLine(pc)
		log.Println(idx, pcfn.Name(), file, line)
	}
	if len(pcs) > 0 {
		log.Println()
	}
}

func finalizerFilter(ov reflect.Value) bool {
	parts := strings.Split(ov.Type().String(), ".")
	clsname := parts[len(parts)-1]
	callers := qtrt.GetCtorAllocStack(clsname)
	_ = callers

	insure := false
	switch ov.Type().String() {
	// case "*qtcore.QString":
	// case "*qtcore.QSize":
	// case "*qtwidgets.QSpacerItem": // crash
	//	insure = true
	//	DumpCallers(callers)
	default:
		insure = true
	}
	if insure {
		log.Println(ov.Type().String(), ov)
	}
	return insure
}

type contentAreaState struct {
	isBottom bool
	curpos   int
	maxpos   int
}

var ccstate = &contentAreaState{isBottom: true}

func main() {
	qtrt.SetFinalizerObjectFilter(finalizerFilter)

	// Create application
	if runtime.GOOS == "android" {
		os.Setenv("QT_AUTO_SCREEN_SCALE_FACTOR ", "1.5")
		qtcore.QCoreApplication_SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true)
	}
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	if false {
		app.SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true) // for android
	}

	setStyleSheet := func() {
		bcc, err := ioutil.ReadFile("../qofia/app.css")
		gopp.ErrPrint(err)
		if true {
			fp := qtcore.NewQFile_1(":/app.css")
			fp.Open(qtcore.QIODevice__ReadOnly)
			bcc = []byte(qtcore.NewQIODeviceFromPointer(fp.GetCthis()).ReadAll().Data())
		}
		app.SetStyleSheet(string(bcc))
	}
	// setStyleSheet()

	// Create main window
	window := qtwidgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize_1(200, 200)

	var winitf qtwidgets.QWidget_ITF
	winitf = window
	log.Println(winitf, winitf.QWidget_PTR())

	uiw = NewUi_MainWindow()
	log.Println(uiw)
	uiw.SetupUi(window)
	// uiw.ListWidget_2.AddItem(qtcore.NewQString_5("aaaaaaaaaaa"))
	// Show the window
	window.Show()

	tb9 := uiw.ToolButton_9
	// dyslot = qtrt.NewQDynSlotObject("abc", 123)
	qtrt.Connect(tb9, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		setStyleSheet()
	})
	stkw := uiw.StackedWidget
	qtrt.Connect(uiw.ToolButton_11, "clicked(bool)", func(checked bool) {
		cidx := stkw.CurrentIndex()
		if cidx > 0 {
			stkw.SetCurrentIndex(cidx - 1)
		}
	})
	qtrt.Connect(uiw.ToolButton_12, "clicked(bool)", func(checked bool) {
		cidx := stkw.CurrentIndex()
		if cidx < stkw.Count()-1 {
			stkw.SetCurrentIndex(cidx + 1)
		}
	})

	{
		qw := uiw.QuickWidget
		// qw.Engine().AddImportPath(":/qmlsys")
		qw.Engine().AddImportPath(":/qmlapp")
		qw.SetSource(qtcore.NewQUrl_1("qrc:/qmlapp/area.qml", 0))
		proot := qw.RootObject()
		gopp.NilPrint(proot, "qml root object nil")
	}

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

	mech = qtcore.NewQBuffer(nil)
	mech.Open(qtcore.QIODevice__ReadWrite)
	qtrt.Connect(mech, "readyRead()", func() {
		log.Println("hehehehhee")
		mech.ReadAll()
		tryReadEvent()
	})
	go initAppBackend()

	if false {
		tmer := qtcore.NewQTimer(nil)
		tmer.Start(500)
		qtrt.Connect(tmer, "timeout()", func() { tryReadEvent() }) // 去掉这个定时器会节省CPU
	}

	vlo10 := uiw.VerticalLayout_10
	_ = vlo10
	for i := 0; i < 30; i++ {
		itext := fmt.Sprintf("hehe %d", i)
		ctivw := qtwidgets.NewQPushButton_1(itext, nil)
		vlo10.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(ctivw.GetCthis()))
	}

	// Execute app
	app.Exec()
}

func initAppBackend() {
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

	lab := uiw.Label_2
	lab.SetText(vtcli.SelfGetName())
	lab = uiw.Label_3
	stmsg, _ := vtcli.SelfGetStatusMessage()
	lab.SetText(stmsg)

	listw := uiw.ListWidget_2

	for fn, frnd := range vtcli.Binfo.Friends {
		itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
		listw.AddItem(itext)
	}

	for gn, grp := range vtcli.Binfo.Groups {
		itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
		listw.AddItem(itext)
	}

	baseInfoGot = true
	select {}
}

var baseInfoGot bool = false

// var lastMsgIvw *widgets.QWidget

func tryReadEvent() {

	if !baseInfoGot {
		return
	}

	// 这个return不会节省cpu???
	if true {
		// return
	}

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

		// cfx, found := this.cfvs.Get(pubkey)
		/*
			cfsx, found := this.chatFormStates.Get(pubkey)
			if !found {
				log.Println("wtf, chat form view not found:", fname, pubkey)
			} else {
				cfs := cfsx.(*ChatFormState)
				msgo := &ContactMessage{}
				msgo.msg = msg
				msgo.tm = time.Now()
				cfs.msgs.Add(msgo)
				// if this.currV != nil && this.currV.(*ChatFormView).cfst == cfs {
				//	this.currV.(*ChatFormView).Signal()
				// }
				if appctx.app.Child != nil && appctx.app.Child.(*ChatFormView).cfst == cfs {
					appctx.app.Child.(*ChatFormView).Signal()
				}
				// InterBackRelay.Signal()
			}
		*/

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

		/*
			valuex, found := appctx.contactStates.Get(groupId)
			var ctis *ContactItemState
			if !found {
				ctis = newContactItemState()
				appctx.contactStates.Put(groupId, ctis)
				log.Println("new group contact:", groupId)
			} else {
				ctis = valuex.(*ContactItemState)
			}
			ctis.group = true
			ctis.cnum = uint32(gopp.MustInt(groupNumber))
			ctis.ctid = groupId

			if appctx.app.Child == nil {
				InterBackRelay.Signal()
			}
		*/
		///
		// _, err := appctx.store.AddGroup(groupId, ctis.cnum, ctis.ctname)
		// gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupTitle := jso.Get("args").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(0).MustString()
		log.Println(groupId)
		if thscli.ConferenceIdIsEmpty(groupId) {
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

		vlo8 := uiw.VerticalLayout_8
		msgivw := qtwidgets.NewQWidget(nil, 0)
		msgivp := NewUi_MessageItemView()
		msgivp.SetupUi(msgivw)
		vlo8.Layout().AddWidget(msgivw)
		msgitmdl = append(msgitmdl, msgivp)

		ccstate.curpos = uiw.ScrollArea_2.VerticalScrollBar().Value()
		ccstate.maxpos = uiw.ScrollArea_2.VerticalScrollBar().Maximum()
		msgivp.Label_5.SetText(itext)
		msgivp.Label_3.SetText(fmt.Sprintf("%s@%s", peerName, groupTitle))
		msgivp.Label_4.SetText(gopp.TimeToFmt1(time.Now()))

		qtrt.Connect(msgivp.ToolButton, "clicked(bool)", func(bool) {
		})

		/*
			valuex, found := appctx.contactStates.Get(groupId)
			var ctis *ContactItemState
			if !found {
				log.Println("group contact not found:", groupId)
			} else {
				ctis = valuex.(*ContactItemState)
			}

			if ctis != nil {
				msgo := &ContactMessage{}
				msgo.msg = message
				msgo.tm = time.Now()
				ctis.msgs.Add(msgo)

				this.signalProperView(ctis, false)
			}
		*/

		//
		// peerPubkey := jso.Get("margs").GetIndex(1).MustString()
		// _, err := appctx.store.AddGroupMessage(message, "0", groupId, peerPubkey)
		// gopp.ErrPrint(err)

	default:
	}
}
