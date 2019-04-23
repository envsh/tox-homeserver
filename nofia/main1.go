package main

import (
	nk "mkuse/nuklear"
	thscli "tox-homeserver/client"
)

type Render interface {
	render() func(*nk.Context)
}

func main() {
	appctx = thscli.NewAppContext()
	go runbkdproc()

	app := nk.NewApp()
	uictx.app = app

	app.Exec(
		uictx.minfov.render(),
		uictx.ctview.render(),
		uictx.myactv.render(),
		uictx.fiview.render(),
		uictx.chatform.render(),
		uictx.sendv.render(),
	)
}

var appctx *thscli.AppContext
var vtcli *thscli.LigTox

// ui context
var uictx = NewUiContext()

type UiContext struct {
	app      *nk.NkApp
	minfov   *MyinfoView
	ctview   *ContectView
	myactv   *MyactionView
	fiview   *FriendInfoView
	chatform *ChatForm
	sendv    *SendForm

	themeNo int
}

func NewUiContext() *UiContext {
	this := &UiContext{}

	this.minfov = &MyinfoView{}
	this.ctview = NewcontactView()
	this.myactv = &MyactionView{}
	this.fiview = &FriendInfoView{}
	this.chatform = NewChatForm()
	this.sendv = NewSendForm()

	// this.themeNo = _STL_SYSTEM
	return this
}

func (this *UiContext) init() *UiContext {
	return this
}

func (this *UiContext) runOnUiThread() {

}
