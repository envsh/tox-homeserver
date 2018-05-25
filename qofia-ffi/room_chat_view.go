package main

import (
	"gopp"
	"log"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

type RoomChatState struct {
	roomOpMenu             *qtwidgets.QMenu
	actDisableNotification *qtwidgets.QAction
	actInviteFrient        *qtwidgets.QAction
	actLeaveRoom           *qtwidgets.QAction
	tblItems               []*qtwidgets.QTableWidgetItem
}

func (this *MainWindow) initRoomChat() {
	this.initRoomChatUi()
	this.initRoomChatSignals()
}

func (this *MainWindow) initRoomChatUi() {
	this.TableWidget.InsertColumn(0)
	this.TableWidget.InsertColumn(1)
	this.TableWidget.InsertColumn(2)
	this.roomOpMenu = qtwidgets.NewQMenu(nil)
	this.actDisableNotification = this.roomOpMenu.AddAction("Disable Notification")
	this.actInviteFrient = this.roomOpMenu.AddAction("Invite Friend")
	this.actLeaveRoom = this.roomOpMenu.AddAction("Leave Room")
	this.ToolButton_22.SetMenu(this.roomOpMenu)
}

func (this *MainWindow) initRoomChatSignals() {
	qtrt.Connect(this.ToolButton_22, "clicked()", func() {
		log.Println("heheh")
		this.ToolButton_22.ShowMenu()
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
		this.switchUiStack(UIST_MAINUI)
	})

	qtrt.Connect(this.PushButton_4, "clicked(bool)", func(bool) {
		toxid := this.LineEdit_4.Text()
		addmsg := this.TextEdit.ToPlainText()
		phmsg := this.TextEdit.PlaceholderText()
		err := this.addFriendByToxId(toxid, addmsg, phmsg)
		gopp.ErrPrint(err, toxid, addmsg)
		if err != nil {
			log.Println("faild")
		} else {
			this.switchUiStack(UIST_MAINUI)
		}
	})
}

func (this *MainWindow) addFriendByToxId(toxid, addmsg string, phmsg string) error {
	if true {
		t := appctx.GetLigTox()
		if addmsg == phmsg {
			addmsg = qtcore.NewQString_5(phmsg).Arg_11_(t.SelfGetName())
		}
		_, err := t.FriendAdd(toxid, phmsg)
		return err
	}
	return nil
}
