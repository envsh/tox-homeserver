package main

import (
	"gopp"
	"log"
	"math"
	"tox-homeserver/thspbs"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

type RoomChatState struct {
	roomOpMenu             *qtwidgets.QMenu
	actDisableNotification *qtwidgets.QAction
	actShowMembers         *qtwidgets.QAction
	actInviteFrient        *qtwidgets.QAction
	actLeaveRoom           *qtwidgets.QAction
	tblItems               []*qtwidgets.QTableWidgetItem
}

func (this *MainWindow) initRoomChat() {
	this.initRoomChatState()
	this.initRoomChatUi()
	this.initRoomChatSignals()
	this.initRoomChatEvents()
}

func (this *MainWindow) initRoomChatState() {}

func (this *MainWindow) initRoomChatUi() {
	this.TableWidget.InsertColumn(0)
	this.TableWidget.InsertColumn(1)
	this.TableWidget.InsertColumn(2)
	this.TableWidget.SetColumnCount(3)
	this.roomOpMenu = qtwidgets.NewQMenu(nil)
	this.actDisableNotification = this.roomOpMenu.AddAction("Disable Notification")
	this.actShowMembers = this.roomOpMenu.AddAction("Members")
	this.actInviteFrient = this.roomOpMenu.AddAction("Invite Friend")
	this.actLeaveRoom = this.roomOpMenu.AddAction("Leave Room")
	this.ToolButton_22.SetMenu(this.roomOpMenu)
	this.ScrollArea_2.SetWidgetResizable(true)
}

func (this *MainWindow) initRoomChatSignals() {
	qtrt.Connect(this.ToolButton_22, "clicked()", func() {
		log.Println("heheh")
		this.ToolButton_22.ShowMenu()
	})

	qtrt.Connect(this.actShowMembers, "triggered()", func() {
		log.Println("hehhe")
		this.switchUiStack(UIST_MEMBERS)
		this.LoadGroupMemberList(uictx.msgwin.item.GetId())
	})
	qtrt.Connect(this.actInviteFrient, "triggered()", func() {
		log.Println("hehhe")
		this.switchUiStack(UIST_INVITE_FRIEND)
		this.updateInviteFriendPage()
	})
	qtrt.Connect(this.actLeaveRoom, "triggered()", func() {
		log.Println("hehhe")
	})
}

func (this *MainWindow) initRoomChatEvents() {
	// for long content on QLabel, this will truncate can not wrap part
	SetScrollContentTrackerSize(this.ScrollArea_2)
	this.LineEdit_2.InheritDragEnterEvent(func(arg0 *qtgui.QDragEnterEvent) {
		arg0.AcceptProposedAction()
	})
	this.LineEdit_2.InheritDropEvent(func(arg0 *qtgui.QDropEvent) {
		mmdt := arg0.MimeData()
		lst := mmdt.Formats()
		fmts := qtcore.NewQStringListxFromPointer(lst.GetCthis())
		// [application/x-qt-image text/uri-list text/plain text/html image/png image/bmp image/bw image/cur image/eps image/epsf image/epsi image/icns image/ico image/jp2 image/jpeg image/jpg image/pbm BITMAP image/pcx image/pgm image/pic image/ppm image/rgb image/rgba image/sgi image/tga image/tif image/tiff image/wbmp image/webp image/xbm image/xpm]
		mimes := fmts.ConvertToSlice()
		log.Println(mimes)
		urls := mmdt.Urls()
		urlsx := qtcore.NewQUrlListxFromPointer(urls.GetCthis_())
		accepted := false
		if mmdt.HasFormat("text/uri-list") {
			for idx, urlo := range urlsx.ConvertToSlice() {
				log.Println(idx, urlo.IsLocalFile(), urlo.Scheme(), urlo.Port__(), urlo.ToLocalFile())
				if urlo.IsLocalFile() {
					this.sendFile(urlo.ToLocalFile())
					accepted = true
				} else {
					// log.Println(urlo.ToPercentEncoding__("123??")) // crash
				}
			}
		}

		if !accepted && (mmdt.HasFormat("text/plain") || mmdt.HasFormat("text/html")) {
			txt := mmdt.Text()
			gopp.FalsePrint(txt != "", "Why no text value?")
			if txt != "" {
				this.LineEdit_2.SetText(txt)
				accepted = true
			}
		} else if !accepted && mmdt.HasFormat("application/x-qt-image") {
			vimg := mmdt.ImageData()
			log.Println(vimg.ToByteArray().Length(), "Can not process this format.")
		}
		arg0.Ignore()
	})
}

func (this *MainWindow) updateInviteFriendPage() {
	item := uictx.msgwin.item
	log.Println(item.GetName())

	for i := this.TableWidget.RowCount() - 1; i >= 0; i-- {
		this.TableWidget.RemoveRow(i)
	}
	this.tblItems = nil

	for fn, fo := range appctx.GetLigTox().Binfo.GetFriends() {
		log.Println(fn, fo.Name)
		this.TableWidget.InsertRow(0)
		cell := qtwidgets.NewQTableWidgetItem_1(fo.Name, 0)
		log.Println(cell)
		this.TableWidget.SetItem(0, 0, cell)
		this.tblItems = append(this.tblItems, cell)

		cell = qtwidgets.NewQTableWidgetItem__()
		cell.SetCheckState(qtcore.Qt__Unchecked)
		this.TableWidget.SetItem(0, 1, cell)

		cell = qtwidgets.NewQTableWidgetItem_1(fo.GetPubkey(), 0)
		this.TableWidget.SetItem(0, 2, cell)
		this.tblItems = append(this.tblItems, cell)
	}
}

//
func (this *MainWindow) initInivteFriend() {
	this.initInivteFriendUi()
	this.initInivteFriendSignals()
}

func (this *MainWindow) initInivteFriendUi() {

}

func (this *MainWindow) initInivteFriendSignals() {
	qtrt.Connect(this.PushButton, "clicked(bool)", func(bool) {
		this.switchUiStack(UIST_MESSAGEUI)
	})
	qtrt.Connect(this.PushButton_2, "clicked(bool)", func(bool) {
		// log.Println("hehhe")
		selectedCount := 0
		rc := this.TableWidget.RowCount()
		for i := 0; i < rc; i++ {
			cell := this.TableWidget.Item(i, 1)
			// log.Println(cell.CheckState())
			selectedCount += gopp.IfElseInt(cell.CheckState() > 0, 1, 0)
			if cell.CheckState() > 0 {
				cell = this.TableWidget.Item(i, 2)
				// log.Println(cell.Text())
				this.inviteFriendByPubkey(cell.Text())
			}
		}
		if selectedCount > 0 {
			this.switchUiStack(UIST_MESSAGEUI)
		}
	})
}

func (this *MainWindow) inviteFriendByPubkeys(pubkeys []string) {
	for _, pubkey := range pubkeys {
		this.inviteFriendByPubkey(pubkey)
	}
}

func (this *MainWindow) inviteFriendByPubkey(pubkey string) {
	grpnum := this.getCurrentRoom()
	frndnum := uint32(0)
	frndnum, err := appctx.GetLigTox().FriendByPublicKey(pubkey)
	gopp.ErrPrint(err, pubkey)
	log.Printf("Inviting %d to group: %d\n", frndnum, grpnum)
	if true {
		appctx.GetLigTox().ConferenceInvite(grpnum, frndnum)
	}
}

func (this *MainWindow) getCurrentRoom() uint32 {
	item := uictx.msgwin.item
	return item.grpInfo.GetGnum()
}

//
func (this *MainWindow) initAddFriend() {
	this.initAddFriendUi()
	this.initAddFriendSignals()
}

func (this *MainWindow) initAddFriendUi() {

}

func (this *MainWindow) initAddFriendSignals() {
	qtrt.Connect(this.PushButton_3, "clicked(bool)", func(bool) {
		this.switchUiStack(UIST_CONTACTUI)
	})

	qtrt.Connect(this.PushButton_4, "clicked(bool)", func(bool) {
		toxid := this.LineEdit_4.Text()
		addmsg := this.TextEdit.ToPlainText()
		phmsg := this.TextEdit.PlaceholderText()
		frndno, err := this.addFriendByToxId(toxid, addmsg, phmsg)
		gopp.ErrPrint(err, toxid, addmsg)
		if err != nil {
			log.Println("faild")
		} else {
			frndo := &thspbs.FriendInfo{}
			frndo.Fnum = frndno
			frndo.Pubkey = toxid[:64]
			frndo.Name = toxid
			contactQueue <- frndo
			uictx.mech.Trigger()
			this.switchUiStack(UIST_CONTACTUI)
		}
	})
}

func (this *MainWindow) addFriendByToxId(toxid, addmsg string, phmsg string) (uint32, error) {
	if true {
		t := appctx.GetLigTox()
		if addmsg == phmsg {
			addmsg = qtcore.NewQString_5(phmsg).Arg_11_(t.SelfGetName())
		}
		frndno, err := t.FriendAdd(toxid, phmsg)
		return frndno, err
	}
	return math.MaxUint32, nil
}

func (this *MainWindow) updateRoomOpMenu(enableFriend bool) {

	this.actShowMembers.SetEnabled(!enableFriend)
	this.actInviteFrient.SetEnabled(!enableFriend)
	this.actLeaveRoom.SetEnabled(!enableFriend)

	this.ToolButton_14.SetEnabled(enableFriend)
}
