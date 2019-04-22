package main

import (
	nk "mkuse/nuklear"
)

type Render interface {
	render() func(*nk.Context)
}

func main() {
	app := nk.NewApp()
	minfov := &MyinfoView{}
	ctview := &ContectView{}
	myactv := &MyactionView{}
	fiview := &FriendInfoView{}
	chatform := NewChatForm()
	sendv := NewSendForm()

	app.Exec(
		minfov.render(),
		ctview.render(),
		myactv.render(),
		fiview.render(),
		chatform.render(),
		sendv.render(),
	)
}
