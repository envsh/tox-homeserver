package main

import (
	"fmt"
	"gopp"
	"log"
	"time"
)

type MessageListWin struct {
	item *RoomListItem
}

func NewMessageListWin() *MessageListWin {
	this := &MessageListWin{}
	return this
}

func (this *MessageListWin) SetRoom(item *RoomListItem) {
	oldItem := this.item
	this.item = item

	mw := uictx.mw
	mw.Label_5.SetText(gopp.StrSuf4ui(item.GetName(), 36))
	if item.isgroup {
		mw.Label_6.SetVisible(true)
		mw.Label_7.SetVisible(false)
		mw.Label_6.SetText(fmt.Sprintf("%d users in chat", 0))
	} else {
		mw.Label_6.SetVisible(false)
		mw.Label_7.SetVisible(true)
		mw.Label_7.SetText(gopp.StrSuf4ui(item.frndInfo.GetStmsg(), 36))
	}

	this.ReloadMessages(oldItem)
}

func (this *MessageListWin) ReloadMessages(oldItem *RoomListItem) {
	item := this.item

	btime := time.Now()
	vlo8 := uictx.uiw.VerticalLayout_8
	log.Println("clean msg list win:", vlo8.Count())
	if oldItem != nil {
		log.Println("clean msg list win:", vlo8.Count(), len(oldItem.msgitmdl))
		// i > 0 leave the QSpacerItem there
		for i := vlo8.Count() - 1; i > 0; i-- {
			itemv := vlo8.TakeAt(i)
			itemv.Widget().SetVisible(false)
		}
	}
	log.Println(time.Now().Sub(btime))
	log.Println("add msg list win:", len(item.msgitmdl), item.GetName())
	for _, msgiw := range item.msgitmdl {
		vlo8.Layout().AddWidget(msgiw.QWidget_PTR())
		msgiw.QWidget_PTR().SetVisible(true)
	}
	log.Println(time.Now().Sub(btime))
	// TODO too slow, 500ms+
}
