package main

import (
	"fmt"
	"gopp"
	"log"
	"os"
	"time"
	"tox-homeserver/thspbs"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

type RoomListItem struct {
	*Ui_ContactItemView

	cticon *qtgui.QIcon
	subws  []qtwidgets.QWidget_ITF

	pressed  bool
	hovered  bool
	isgroup  bool
	frndInfo *thspbs.FriendInfo
	grpInfo  *thspbs.GroupInfo

	unreaded int
}

func NewRoomListItem() *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.init()
	return this
}

func NewRoomListItem2(info *thspbs.FriendInfo) *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.init()
	this.frndInfo = info
	return this
}

func NewRoomListItem3(info *thspbs.GroupInfo) *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.init()
	this.isgroup = true
	this.grpInfo = info
	return this
}

func (this *RoomListItem) init() {
	labs := []*qtwidgets.QLabel{this.Label_2, this.Label_3, this.Label_4, this.Label_5, this.Label_6}
	for _, lab := range labs {
		lab.SetText("")
		lab.SetAttribute(qtcore.Qt__WA_TranslucentBackground, false)
		lab.SetMouseTracking(true)
		this.subws = append(this.subws, lab)
	}
	this.ToolButton.SetText("")
	this.ToolButton.SetMouseTracking(true)
	this.ToolButton_2.SetMouseTracking(true)

	w := this.ContactItemView
	w.SetMouseTracking(true)

	onMousePress := func(event *qtgui.QMouseEvent) {
		// log.Println(event)
		for _, room := range ctitmdl {
			if room != this {
				room.SetPressState(false)
			}
		}
		this.SetPressState(true)
	}
	onMouseRelease := func(event *qtgui.QMouseEvent) {
		// log.Println(event)
	}
	onMouseMove := func(event *qtgui.QMouseEvent) {
		if true {
			return
		}
		// log.Println(event)
		if !this.hovered {
			this.hovered = true
			for _, room := range ctitmdl {
				if room != this {
					room.OnHover(false)
				}
			}
			if !this.pressed {
				this.OnHover(true)
			}
		}
	}
	_ = onMouseMove
	onMouseLeave := func(event *qtcore.QEvent) {
		this.OnHover(false)
	}
	onMouseEnter := func(event *qtcore.QEvent) {
		this.OnHover(true)
	}

	w.InheritMousePressEvent(onMousePress)
	w.InheritMouseReleaseEvent(onMouseRelease)
	// w.InheritMouseMoveEvent(onMouseMove)
	w.InheritLeaveEvent(onMouseLeave)
	w.InheritEnterEvent(onMouseEnter)

	for _, lab := range labs {
		lab.InheritMousePressEvent(onMousePress)
		lab.InheritMouseReleaseEvent(onMouseRelease)
		// lab.InheritMouseMoveEvent(onMouseMove)
	}
}

func (this *RoomListItem) SetContactInfo(info interface{}) {
	trtxt := gopp.StrSuf4ui

	switch ct := info.(type) {
	case *thspbs.FriendInfo:
		this.frndInfo = ct
		name := gopp.IfElseStr(ct.GetName() == "", ct.GetPubkey()[:7], ct.GetName())
		nametip := gopp.IfElseStr(ct.GetName() == "", ct.GetPubkey()[:17], ct.GetName())
		this.Label_2.SetText(trtxt(name, 26))
		this.Label_2.SetToolTip(nametip)
		this.Label_4.SetText(trtxt(ct.GetStmsg(), 36))
		this.Label_4.SetToolTip(ct.GetStmsg())
		avataricon := fmt.Sprintf("%s/.config/tox/avatars/%s.png", os.Getenv("HOME"), ct.GetPubkey())
		if gopp.FileExist(avataricon) {
			this.cticon = qtgui.NewQIcon_2(avataricon)
			this.ToolButton_2.SetIcon(this.cticon)
		}
	case *thspbs.GroupInfo:
		this.grpInfo = ct
		this.isgroup = true
		this.Label_2.SetText(trtxt(ct.GetTitle(), 26))
		this.Label_2.SetToolTip(ct.GetTitle())
		this.Label_4.SetHidden(true)
		this.QWidget_PTR().SetFixedHeight(this.QWidget_PTR().Height() - 20)
		this.cticon = qtgui.NewQIcon_2(":/icons/groupgray.png")
		this.ToolButton_2.SetIcon(this.cticon)
	default:
		log.Fatalln("wtf")
	}

}

func (this *RoomListItem) GetName() string {
	return gopp.IfElseStr(this.isgroup, this.grpInfo.GetTitle(), this.frndInfo.GetName())
}

func (this *RoomListItem) GetId() string {
	if this.isgroup {
		log.Println(this.grpInfo.GetGroupId(), this.grpInfo.Title)
	}
	return gopp.IfElseStr(this.isgroup, this.grpInfo.GetGroupId(), this.frndInfo.GetPubkey())
}
func (this *RoomListItem) SetLastMsg(msg string) {
	this.Label_3.SetText(gopp.StrSuf4ui(msg, 36))
	this.Label_3.SetToolTip(msg)
	tm := time.Now()
	this.Label_6.SetText(fmt.Sprintf("%2d:%2d", tm.Hour(), tm.Minute()))
	this.Label_6.SetToolTip(gopp.TimeToFmt1(tm))
}

func (this *RoomListItem) SetPressState(pressed bool) {
	changed := this.pressed != pressed
	log.Println("changed:", changed, "pressed:", pressed, this.GetName())
	if changed {
		this.pressed = pressed
		this.SetBgColor(gopp.IfElseStr(pressed, "selected", "default"))
	}
}

func (this *RoomListItem) OnHover(hover bool) {
	this.hovered = hover
	if !this.pressed {
		this.SetBgColor(gopp.IfElseStr(hover, "hover", "default"))
	}
}

func (this *RoomListItem) SetBgColor(p string) {
	log.Println("set color:", p)
	css := ""
	switch p {
	case "selected":
		css = "background: #38A3D8;"
	case "hover":
		css = "background: #c8c8c8;"
	case "default":
		css = "background: #FFFFFF;"
	default:
		log.Println("wtf", p)
	}
	if css != "" {
		this.QWidget_PTR().SetStyleSheet(css)
		for _, w := range this.subws {
			if false {
				w.QWidget_PTR().SetStyleSheet(css)
			}
		}
	}
}
