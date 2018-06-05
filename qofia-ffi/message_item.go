package main

// Ui_MessageItemView's wrapper
type MessageItem struct {
	*Ui_MessageItemView

	Sent     bool
	UserCode int64
}

func NewMessageItem() *MessageItem {
	this := &MessageItem{}
	this.Ui_MessageItemView = NewUi_MessageItemView2()

	return this
}
