package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"time"

	"qt.go/cffiqt"
	"qt.go/qtcore"
	"qt.go/qtwidgets"

	thscli "tox-homeserver/client"
	"tox-homeserver/gofia"

	simplejson "github.com/bitly/go-simplejson"
)

var appctx *gofia.AppContext
var vtcli *thscli.LigTox

var uiw *Ui_MainWindow
var dyslot *ffiqt.QDynSlotObject

func main() {
	// Create application
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	setStyleSheet := func() {
		bcc, err := ioutil.ReadFile("../qofia/app.css")
		gopp.ErrPrint(err)
		app.SetStyleSheet(qtcore.NewQString_5(string(bcc)))
	}
	// setStyleSheet()

	// Create main window
	window := qtwidgets.NewQMainWindow(nil, 0)
	// window.SetWindowTitle("Hello World Example")
	// window.SetMinimumSize2(200, 200)

	uiw = NewUi_MainWindow()
	log.Println(uiw)
	uiw.SetupUi(window)
	uiw.ListWidget_2.AddItem(qtcore.NewQString_5("aaaaaaaaaaa"))

	// Show the window
	window.Show()

	tb9 := uiw.ToolButton_9
	dyslot = ffiqt.NewQDynSlotObject("abc", 123)
	dyslot.Connect(tb9, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		setStyleSheet()
	})
	/*
		a2 := widgets.NewQScrollAreaFromPointer(Ui_MainWindow_Get_scrollArea_2(uiw))
		toval := a2.VerticalScrollBar().Maximum() + 80
		log.Println(a2.VerticalScrollBar().Value(), toval)
		a2.VerticalScrollBar().ConnectRangeChanged(func(min int, max int) {
			log.Println(min, max)
		})

	*/

	go initAppBackend()

	tmer := qtcore.NewQTimer(nil)
	tmer.Start(500)
	dyslot.Connect(tmer, "timeout()", func() { tryReadEvent() })

	vlo10 := uiw.VerticalLayout_10
	_ = vlo10
	for i := 0; i < 30; i++ {
		itext := fmt.Sprintf("hehe %d", i)
		ctivw := qtwidgets.NewQPushButton_1(qtcore.NewQString_5(itext), nil)
		vlo10.AddWidget(qtwidgets.NewQWidgetFromPointer(ctivw.GetCthis()))
	}

	// Execute app
	app.Exec()
}

func initAppBackend() {
	gofia.AppOnCreate()
	appctx = gofia.GetAppCtx()
	vtcli = appctx.GetLigTox()

	for {
		time.Sleep(500 * time.Millisecond)
		if vtcli.SelfGetAddress() != "" {
			break
		}
	}
	log.Println(vtcli.SelfGetAddress())

	lab := uiw.Label_2
	lab.SetText(qtcore.NewQString_5(vtcli.SelfGetName()))
	lab = uiw.Label_3
	stmsg, _ := vtcli.SelfGetStatusMessage()
	lab.SetText(qtcore.NewQString_5(stmsg))

	listw := uiw.ListWidget_2

	for fn, frnd := range vtcli.Binfo.Friends {
		itext := fmt.Sprintf("%d-%s", fn, frnd.GetName())
		listw.AddItem(qtcore.NewQString_5(itext))
	}

	for gn, grp := range vtcli.Binfo.Groups {
		itext := fmt.Sprintf("%d-%s", gn, grp.GetTitle())
		listw.AddItem(qtcore.NewQString_5(itext))
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
		uiw.ListWidget.AddItem(qtcore.NewQString_5(itext))
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
		uiw.ListWidget.AddItem(qtcore.NewQString_5(itext))
		uiw.ListWidget.ScrollToBottom()
		log.Println("item:", itext)

		vlo8 := uiw.VerticalLayout_8
		msgivw := qtwidgets.NewQWidget(nil, 0)
		msgivp := NewUi_MessageItemView()
		msgivp.SetupUi(msgivw)
		vlo8.AddWidget(msgivw)

		tbrw := msgivp.TextBrowser
		tbrw.SetText(qtcore.NewQString_5(itext))
		msgivp.Label_3.SetText(qtcore.NewQString_5(fmt.Sprintf("%s@%s", peerName, groupTitle)))
		msgivp.Label_4.SetText(qtcore.NewQString_5(time.Now().String()))

		dyslot.Connect(msgivp.ToolButton, "clicked(bool)", func(bool) {
			log.Println(tbrw)
			log.Println(tbrw.GetCthis())
			log.Println(tbrw.Size())
			log.Println(tbrw.Size().GetCthis())
			log.Println(tbrw.Size().Height(), tbrw.Size().Width())
			log.Println(tbrw.Size().Rheight(), tbrw.Size().Rwidth())
			log.Println(tbrw.Document().Size().IsValid())
			log.Println(tbrw.Document().Size().IsNull())
			log.Println(tbrw.Document().Size().IsEmpty())
			szo := tbrw.Document().Size()
			log.Println(szo.Width(), szo.Height())
			log.Println(szo.Rwidth(), szo.Rheight())
			log.Println(tbrw.Document().Size().Rheight())
			log.Println(tbrw.Document().Size().Height())
			// tbrw.SetFixedHeight(int(tbrw.Document().Size().Rheight()))
		})

		/*
			// not called
			tbrw.ConnectSizeHint(func() *core.QSize {
				log.Println(tbrw.Document().Size().Rheight())
				log.Println(tbrw.Size().Rheight())
				return core.NewQSize2(int(tbrw.Document().Size().Rwidth()), int(tbrw.Document().Size().Rheight()))
			})

			// 这个只在有滚动条的时候有效
			tbrw.VerticalScrollBar().ConnectRangeChanged(func(min int, max int) {
				log.Println(min, max, tbrw.Viewport().Size().Rheight())
				log.Print(tbrw.Document().Size().Rheight())
				// tbrw.SetFixedHeight(int(tbrw.Document().Size().Rheight()) + max)
				// tbrw.Document().AdjustSize() // 依旧有滚动条
			})
			tbrw.HorizontalScrollBar().ConnectRangeChanged(func(min int, max int) {
				log.Println(min, max, tbrw.Viewport().Size().Rheight())
				log.Print(tbrw.Document().Size().Rheight())
			})
			// not called
			msgivw.ConnectResizeEvent(func(event *gui.QResizeEvent) {
				oldWidth := event.OldSize().Rwidth()
				newWidth := event.Size().Rwidth()
				log.Println(oldWidth, "=>", newWidth, tbrw.Document().Size().Rheight())
				fixnum := 2
				if tbrw.Size().Rheight() != int(tbrw.Document().Size().Rheight())+fixnum {
					tbrw.SetFixedHeight(int(tbrw.Document().Size().Rheight()) + fixnum)

				}
			})
			// not called
			tbrw.Viewport().ConnectResizeEvent(func(event *gui.QResizeEvent) {
				oldWidth := event.OldSize().Rwidth()
				newWidth := event.Size().Rwidth()
				log.Println(oldWidth, "=>", newWidth, tbrw.Document().Size().Rheight())
			})
			tbrw.Document().ConnectBlockCountChanged(func(newBlockCount int) {
				log.Println(newBlockCount)
			})
			tbrw.Document().ConnectDocumentLayoutChanged(func() {
				log.Println("111")
			})
			tbrw.Document().ConnectContentsChanged(func() {
				log.Print(tbrw.Document().Size().Rheight())
				log.Print(tbrw.Size().Rheight())
			})

			a2 := widgets.NewQScrollAreaFromPointer(Ui_MainWindow_Get_scrollArea_2(uiw))
			toval := a2.VerticalScrollBar().Maximum() + int(tbrw.Document().Size().Rheight()) + 80
			log.Println(a2.VerticalScrollBar().Value(), toval)
			if a2.VerticalScrollBar().Value() != toval {
				// a2.EnsureWidgetVisible(msgivw, 0, 0)
				a2.VerticalScrollBar().SetValue(toval)
			}
			if lastMsgIvw != nil {
				lastMsgIvw.DisconnectResizeEvent()
				lastMsgIvw = msgivw
			}
			a2.Viewport().ConnectResizeEvent(func(event *gui.QResizeEvent) {
				oldWidth := event.OldSize().Rwidth()
				newWidth := event.Size().Rwidth()
				log.Println(oldWidth, "=>", newWidth, tbrw.Document().Size().Rheight())
			})
		*/
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
