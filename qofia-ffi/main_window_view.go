package main

import (
	"gopp"
	"log"
	"runtime"
	"strings"
	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

// TODO split main_window.go's some code here

func (this *MainWindow) initMainWindow() {
	this.initMainWindowUi()
	this.initMainWindowQml()
	this.initMainWindowSignals1() // TODO merge two
	this.initMainWindowSignals()
	this.switchUiStack(uictx.uiw.StackedWidget.CurrentIndex())
	// this.Widget.SetStyleSheet(GetBg(_HEADER_BG))

	this.initMainWindowEvents()
}

func (this *MainWindow) initMainWindowUi() {
	this.setConnStatus(false)
	this.LineEdit_5.SetVisible(false)
	this.LineEdit_6.SetVisible(false)
	SetScrollContentTrackerSize(this.ScrollArea)
}

func (this *MainWindow) initMainWindowSignals1() {

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

}

func (this *MainWindow) initMainWindowSignals() {
	uiw := uictx.uiw

	qtrt.Connect(uiw.ToolButton_19, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		ShowToast("hehehh哈哈eehhe", 1)
	})
	qtrt.Connect(uiw.ToolButton_20, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		// testRunOnAndroidThread()
		KeepScreenOn(checked)
	})
	qtrt.Connect(uiw.ToolButton_21, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		setAppStyleSheet()
	})

	qtrt.Connect(uiw.ToolButton_23, "clicked(bool)", func(checked bool) {
		log.Println(checked, uictx.msgwin.item == nil)
		if uictx.msgwin.item != nil {
			go func() {
				hisfet.PullPrevHistoryByRoomItem(uictx.msgwin.item)
			}()
		}
	})

	qtrt.Connect(uiw.ToolButton_24, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		uiw.ScrollAreaWidgetContents_2.SetSizePolicy_1(qtwidgets.QSizePolicy__Fixed, qtwidgets.QSizePolicy__Fixed)
	})
	qtrt.Connect(uiw.ToolButton_25, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		uiw.ScrollAreaWidgetContents_2.SetSizePolicy_1(qtwidgets.QSizePolicy__Maximum, qtwidgets.QSizePolicy__Fixed)
	})
	qtrt.Connect(uiw.ToolButton_26, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		uiw.ScrollAreaWidgetContents_2.SetSizePolicy_1(qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Fixed)
	})
	qtrt.Connect(uiw.ToolButton_27, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		uiw.ScrollAreaWidgetContents_2.SetSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Fixed)
	})
	qtrt.Connect(uiw.ToolButton_28, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		uiw.ScrollAreaWidgetContents_2.SetSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Fixed)
	})
	qtrt.Connect(uiw.ToolButton_29, "clicked(bool)", func(checked bool) {
		log.Println(checked)
		uiw.ScrollAreaWidgetContents_2.SetSizePolicy_1(qtwidgets.QSizePolicy__MinimumExpanding, qtwidgets.QSizePolicy__Fixed)
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
	// switch ui
	qtrt.Connect(uiw.ComboBox, "currentTextChanged(const QString&)", func(text string) {
	})

	//
	qtrt.Connect(uiw.ToolButton_4, "clicked(bool)", func(bool) {
		this.switchUiStack(UIST_ADD_FRIEND)
		cb := uictx.qtapp.Clipboard()
		log.Println(cb.Text__())
		toxid := cb.Text__()
		if toxid != "" {
			this.LineEdit_4.SetText(toxid)
		}
		addmsg := qtcore.NewQString_5(this.TextEdit.PlaceholderText()).Arg_11_(appctx.GetLigTox().SelfGetName())
		this.TextEdit.SetPlainText(addmsg)

		/*
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
		*/
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

	uiw.Label_2.InheritMousePressEvent(func(ev *qtgui.QMouseEvent) {
		uiw.Label_2.SetVisible(false)
		uiw.LineEdit_5.SetText(uiw.Label_2.ToolTip())
		uiw.LineEdit_5.SetVisible(true)
	})
	uiw.Label_3.InheritMousePressEvent(func(ev *qtgui.QMouseEvent) {
		uiw.Label_3.SetVisible(false)
		uiw.LineEdit_6.SetText(uiw.Label_3.ToolTip())
		uiw.LineEdit_6.SetVisible(true)
	})
	qtrt.Connect(uiw.LineEdit_5, "editingFinished()", func() {
		if !uiw.LineEdit_5.IsVisible() {
			return // it's lost focus event
		}
		txt := uiw.LineEdit_5.Text()
		uiw.LineEdit_5.SetVisible(false)
		uiw.Label_2.SetVisible(true)

		txt = strings.TrimSpace(txt)
		if txt != "" && txt != uiw.Label_3.ToolTip() {
			uiw.Label_2.SetText(gopp.StrSuf4ui(txt, thscom.UiNameLen))
			uiw.Label_2.SetToolTip(txt)
			SetQLabelElideText(uiw.Label_2, txt)
			vtcli.SelfSetName(txt)
		}
	})
	qtrt.Connect(uiw.LineEdit_6, "editingFinished()", func() {
		if !uiw.LineEdit_6.IsVisible() {
			return // it's lost focus event
		}
		txt := uiw.LineEdit_6.Text()
		uiw.LineEdit_6.SetVisible(false)
		uiw.Label_3.SetVisible(true)

		txt = strings.TrimSpace(txt)
		if txt != "" && txt != uiw.Label_3.ToolTip() {
			uiw.Label_3.SetText(gopp.StrSuf4ui(txt, thscom.UiStmsgLen))
			uiw.Label_3.SetToolTip(txt)
			SetQLabelElideText(uiw.Label_3, txt)
			vtcli.SelfSetStatusMessage(txt)
		}
	})

}

func (this *MainWindow) initMainWindowQml() {
	qw := uictx.uiw.QuickWidget
	// qw.Engine().AddImportPath(":/qmlsys")
	qw.Engine().AddImportPath(":/qmlapp")
	qw.SetSource(qtcore.NewQUrl_1("qrc:/qmlapp/area.qml", 0))
	proot := qw.RootObject()
	gopp.NilPrint(proot, "qml root object nil")
}

func (this *MainWindow) initMainWindowEvents() {
	// can not capture back button for all of these methods
	// in C++, it's override keyPressEvent(), works fine. but why here not work???
	// good, it's captured by centralWidget
	// capwgt := this.MainWindow
	capwgt := this.Centralwidget
	// must accept when event.Key() == Qt__Key_Back, or the app is crash, not exit.
	// but in qt.go, it's seems has some problem for handle this case, always crash.
	capwgt.InheritKeyPressEvent(func(event *qtgui.QKeyEvent) {
		log.Println(event.Key(), event.Text())
		switch event.Key() {
		case qtcore.Qt__Key_Back:
			log.Println("[[Back button]]")
			quit := this.onAppBackButton(true)
			this.quitClean(quit)
			event.SetAccepted(!quit)
		case qtcore.Qt__Key_Menu:
			log.Println("[[Menu button]]")
			event.Ignore() // default don't touch it.
		case qtcore.Qt__Key_TopMenu:
			log.Println("[[Top menu button]]")
			event.Ignore() // default don't touch it.
		default:
			event.Ignore() // default don't touch it.
		}
	})

	firstShow := true
	capwgt.InheritShowEvent(func(event *qtgui.QShowEvent) {
		event.Ignore()
		log.Println(event.Spontaneous())
		if firstShow {
			firstShow = false
			// do after show
			runOnUiThread(func() { this.initFirstShow() })
		}
	})
}
