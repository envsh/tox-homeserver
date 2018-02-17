package main

type contentAreaState struct {
	isBottom bool
	curpos   int
	maxpos   int
}

var ccstate = &contentAreaState{isBottom: true}

// for message content list scroll area
type MessageListItem struct {
	*Ui_MessageItemView
	uiw *Ui_MessageItemView
}

func NewMessageListItem() *MessageListItem {
	this := &MessageListItem{}
	this.Ui_MessageItemView = NewUi_MessageItemView2()
	this.uiw = this.Ui_MessageItemView
	return this
}

func (this *MessageListItem) initUiBase() {
	// this.uiw.Label_4.SetProperty(name string, value qtcore.QVariant_ITF)
}
