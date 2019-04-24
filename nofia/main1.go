package main

import (
	"flag"
	nk "mkuse/nuklear"
	thscli "tox-homeserver/client"
)

type Render interface {
	render() func(*nk.Context)
}

func main() {
	flag.Parse()

	appctx = thscli.NewAppContext()
	go runbkdproc()

	app := nk.NewApp()
	uictx.app = app

	app.Exec(
		uictx.icop.render(),

		uictx.minfov.render(),
		uictx.ctview.render(),
		uictx.myactv.render(),
		uictx.fiview.render(),
		uictx.chatform.render(),
		uictx.sendv.render(),
		uictx.mblstv.render(),
		uictx.setfmv.render(),
	)
}

var appctx *thscli.AppContext
var vtcli *thscli.LigTox

// ui context
var uictx = NewUiContext()

type UiContext struct {
	app      *nk.NkApp
	mdl      *thscli.DataModel
	minfov   *MyinfoView
	ctview   *ContectView
	myactv   *MyactionView
	fiview   *FriendInfoView
	chatform *ChatForm
	sendv    *SendForm
	mblstv   *MemberListForm
	setfmv   *SettingForm

	icop *IconPool //

	themeNo int
}

func NewUiContext() *UiContext {
	this := &UiContext{}

	this.mdl = thscli.NewDataModel()

	this.minfov = &MyinfoView{}
	this.ctview = NewcontactView()
	this.myactv = &MyactionView{}
	this.fiview = &FriendInfoView{}
	this.chatform = NewChatForm()
	this.sendv = NewSendForm()
	this.mblstv = NewMemberListForm()
	this.setfmv = NewSettingForm()

	this.icop = &IconPool{}

	// this.themeNo = _STL_SYSTEM
	return this
}

func (this *UiContext) init() *UiContext {
	return this
}

func (this *UiContext) runOnUiThread() {

}
