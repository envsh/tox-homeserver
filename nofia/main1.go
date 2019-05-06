package main

import (
	"flag"
	"gopp"

	"os/exec"

	nk "mkuse/nuklear"
	thscli "tox-homeserver/client"

	"github.com/atotto/clipboard"
	xdgopen "github.com/skratchdot/open-golang/open"
)

type Render interface {
	render() func(*nk.Context)
}

func main() {
	flag.Parse()

	startuimain()
	select {}

	app := nk.NewApp()
	uictx.app = app
	mdl := thscli.NewDataModel(func() { app.Repaint() })
	uictx.mdl = mdl

	appctx = thscli.NewAppContext()
	go runbkdproc()

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

type UiService struct {
}

func (this *UiContext) XdgOpen(u string) {
	err := xdgopen.Run(u)
	gopp.ErrPrint(err)

	if false {
		xdgopen, err := exec.LookPath("xdg-open")
		gopp.ErrPrint(err)
		if err == nil {
			err = exec.Command(xdgopen, u).Run()
			gopp.ErrPrint(err, u)
		}
	}
}

func (this *UiContext) Copy2Clipboard(txt string) {
	err := clipboard.WriteAll(txt)
	gopp.ErrPrint(err)
}

func (this *UiContext) ReadClipboard() string {
	s, err := clipboard.ReadAll()
	gopp.ErrPrint(err)
	return s
}
