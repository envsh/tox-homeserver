package main

import (
	"sort"
	"tox-homeserver/thspbs"

	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

var stateGroupMemberList = &struct {
	tblItems []*qtwidgets.QTableWidgetItem
}{}

func (this *MainWindow) initGroupMemberList() {
	this.initGroupMemberListUi()
	this.initGroupMemberListSignals()
}

func (this *MainWindow) initGroupMemberListUi() {
	this.PushButton_9.SetVisible(false)
	this.TableWidget_2.SetColumnCount(3) // 0:icon, 1: name, 2:pubkey
	this.TableWidget_2.SetColumnWidth(0, 35)
	this.TableWidget_2.SetColumnWidth(1, 130)
	this.TableWidget_2.SetColumnWidth(2, 180)
	this.TableWidget_2.SetSelectionBehavior(qtwidgets.QAbstractItemView__SelectRows)
}

func (this *MainWindow) initGroupMemberListSignals() {
	qtrt.Connect(this.PushButton_8, "clicked(bool)", func(bool) {
		this.switchUiStack(UIST_MESSAGEUI)
		// clear?
		this.ClearGroupMemberList()
	})
	qtrt.Connect(this.PushButton_9, "clicked(bool)", func(bool) {})
}

func (this *MainWindow) LoadGroupMemberList(pubkey string) {
	// clear?
	this.ClearGroupMemberList()

	vtcli := appctx.GetLigTox()
	binfo := vtcli.Binfo
	peers := binfo.GetGroupMembersByPubkey(pubkey)
	grpids := []string{}
	for grpid, _ := range peers {
		grpids = append(grpids, grpid)
	}
	sort.Strings(grpids)
	for _, grpid := range grpids {
		peero := peers[grpid]
		this.AddGroupMember(peero)
	}
}

func (this *MainWindow) AddGroupMember(peero *thspbs.MemberInfo) {
	tabwgt := this.TableWidget_2

	tabwgt.InsertRow(0)
	cell := qtwidgets.NewQTableWidgetItem_1(peero.Name, 0)
	tabwgt.SetItem(0, 1, cell)
	this.tblItems = append(this.tblItems, cell)

	cell = qtwidgets.NewQTableWidgetItem_1(peero.Pubkey, 0)
	tabwgt.SetItem(0, 2, cell)
	ste := stateGroupMemberList
	ste.tblItems = append(ste.tblItems, cell)
}

func (this *MainWindow) ClearGroupMemberList() {
	tblwgt := this.TableWidget_2

	tblwgt.Clear()
	for i := tblwgt.RowCount() - 1; i >= 0; i-- {
		tblwgt.RemoveRow(i)
	}
	stateGroupMemberList.tblItems = nil
}
