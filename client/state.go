package client

import "github.com/kitech/godsts/lists/arraylist"

type ContactItemState struct {
	group  bool
	cnum   uint32
	ctid   string
	ctname string
	status uint32
	stmsg  string
	avatar string

	// *ContactMessage
	msgs *arraylist.List
	// Text           *text.Text
	// Responder      *keyboard.Responder
	// ScrollPosition *view.ScrollPosition
}
type ChatFormState = ContactItemState

func newContactItemState() *ContactItemState {
	this := &ContactItemState{}
	this.msgs = arraylist.New()
	// this.Text = text.New("")
	// this.Responder = &keyboard.Responder{}
	// this.ScrollPosition = &view.ScrollPosition{}
	return this
}
