package main

import (
	"fmt"
	"gopp"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

type Message struct {
	Msg  string
	Peer string
	Time time.Time
}

func NewMessageForGroup(jso *simplejson.Json) *Message {
	groupId := jso.Get("margs").GetIndex(3).MustString()
	log.Println(groupId)
	if thscli.ConferenceIdIsEmpty(groupId) {
		// break
	}

	message := jso.Get("args").GetIndex(3).MustString()
	peerName := jso.Get("margs").GetIndex(0).MustString()
	groupTitle := jso.Get("margs").GetIndex(2).MustString()
	_ = groupTitle

	this := &Message{}
	this.Msg = message
	this.Peer = peerName
	this.Time = time.Now()

	return this
}

func NewMessageForFriend(jso *simplejson.Json) *Message {
	msg := jso.Get("args").GetIndex(1).MustString()
	fname := jso.Get("margs").GetIndex(0).MustString()
	pubkey := jso.Get("margs").GetIndex(1).MustString()
	_, _, _ = msg, fname, pubkey

	this := &Message{}
	this.Msg = msg
	this.Peer = fname
	this.Time = time.Now()

	return this
}

type RoomListMan struct{}

func NewRoomListMan() *RoomListMan { return &RoomListMan{} }
func (this *RoomListMan) Get(id string) *RoomListItem {
	for _, item := range uictx.ctitmdl {
		if item.GetId() == id {
			return item
		}
	}
	return nil
}
func (this *RoomListMan) addRoomItem(item *RoomListItem) {
	uictx.ctitmdl = append(uictx.ctitmdl, item)
	uictx.uiw.VerticalLayout_10.InsertWidget(0, item.QWidget_PTR(), 0, 0)
}

func (this *RoomListMan) Delete(item *RoomListItem) {
	for i := 0; i < len(uictx.ctitmdl); i++ {
		tmpi := uictx.ctitmdl[i]
		if tmpi == item {
			for j := i + 1; j < len(uictx.ctitmdl); j++ {
				uictx.ctitmdl[j-1] = uictx.ctitmdl[j]
			}
			uictx.ctitmdl = uictx.ctitmdl[:len(uictx.ctitmdl)-1]
			break
		}
	}

	uictx.uiw.VerticalLayout_10.RemoveWidget(item.QWidget_PTR())
	item.QWidget_PTR().SetVisible(false)
	// TODO really destroy
}

type RoomListItem struct {
	*Ui_ContactItemView

	OnConextMenu func(w *qtwidgets.QWidget, pos *qtcore.QPoint)

	cticon *qtgui.QIcon
	subws  []qtwidgets.QWidget_ITF
	menu   *qtwidgets.QMenu

	msgitmdl []*Ui_MessageItemView
	msgos    []*Message

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
	this.SetContactInfo(info)
	this.init()
	return this
}

func NewRoomListItem3(info *thspbs.GroupInfo) *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.SetContactInfo(info)
	this.init()
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
		if event.Button() == qtcore.Qt__LeftButton {
			for _, room := range uictx.ctitmdl {
				if room != this {
					room.SetPressState(false)
				}
			}
			this.SetPressState(true)
		}
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
			for _, room := range uictx.ctitmdl {
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

	w.InheritContextMenuEvent(func(event *qtgui.QContextMenuEvent) {
		gpos := event.GlobalPos()
		log.Println(event.Type(), gpos.X(), gpos.Y())
		if this.OnConextMenu != nil {
			this.OnConextMenu(w, gpos)
		}
	})
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
		log.Println(ct.GetTitle(), ct.Title, trtxt(ct.GetTitle(), 26))
		this.grpInfo = ct
		this.isgroup = true
		this.Label_2.SetText(trtxt(ct.GetTitle(), 26))
		log.Println(this.Label_2.Text())
		this.Label_2.SetToolTip(ct.GetTitle())
		this.Label_4.SetHidden(true)
		this.QWidget_PTR().SetFixedHeight(this.QWidget_PTR().Height() - 20)
		this.cticon = qtgui.NewQIcon_2(":/icons/groupgray.png")
		this.ToolButton_2.SetIcon(this.cticon)
	default:
		log.Fatalln("wtf")
	}
	this.ToolButton_2.SetToolTip(this.GetName() + "." + this.GetId()[:7])
}

func (this *RoomListItem) AddMessage(msgo *Message) {
	this.msgos = append(this.msgos, msgo)
	msgiw := NewUi_MessageItemView2()
	this.msgitmdl = append(this.msgitmdl, msgiw)

	msgiw.Label_5.SetText(msgo.Msg)
	msgiw.Label_3.SetText(fmt.Sprintf("%s@%s", msgo.Peer, this.GetName()))
	msgiw.Label_4.SetText(gopp.TimeToFmt1(msgo.Time))

	if uictx.msgwin.item == this {
		vlo8 := uictx.uiw.VerticalLayout_8
		vlo8.Layout().AddWidget(msgiw.QWidget_PTR())
	}
	this.SetLastMsg(fmt.Sprintf("%s: %s", gopp.StrSuf4ui(msgo.Peer, 8, 1), msgo.Msg))
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

func (this *RoomListItem) UpdateName(name string) {
	if this.isgroup {
		if this.grpInfo.Title != name {
			this.grpInfo.Title = name
			this.Label_2.SetText(gopp.StrSuf4ui(name, 26))
			this.Label_2.SetToolTip(name)
			this.ToolButton_2.SetToolTip(name + "." + this.GetId()[:7])
		}
	} else {
	}
}

func (this *RoomListItem) SetLastMsg(msg string) {
	cmsg := msg
	if strings.HasPrefix(msg, "toxync:") {
		reg := regexp.MustCompile(`toxync: \[(.+)@([^\[]+)\] (.+)`)
		if reg.MatchString(msg) {
			mats := reg.FindAllStringSubmatch(msg, -1)
			if len(mats) > 0 {
				cmsg = fmt.Sprintf("%s@%s: %s", mats[0][1], mats[0][2], mats[0][3])
			}
		}
	}
	this.Label_3.SetText(gopp.StrSuf4ui(cmsg, 36))
	this.Label_3.SetToolTip(cmsg)
	tm := time.Now()
	this.Label_6.SetText(fmt.Sprintf("%02d:%02d", tm.Hour(), tm.Minute()))
	this.Label_6.SetToolTip(gopp.TimeToFmt1(tm))
}

func (this *RoomListItem) SetPressState(pressed bool) {
	changed := this.pressed != pressed
	log.Println("changed:", changed, "pressed:", pressed, this.GetName())
	if changed {
		this.pressed = pressed
		this.SetBgColor(gopp.IfElseStr(pressed, "selected", "default"))
	}

	uictx.mw.switchUiStack(4)
	if changed {
		uictx.msgwin.SetRoom(this)
	}
}

func (this *RoomListItem) OnHover(hover bool) {
	this.hovered = hover
	if !this.pressed {
		this.SetBgColor(gopp.IfElseStr(hover, "hover", "default"))
	}
}

func (this *RoomListItem) SetBgColor(p string) {
	css := ""
	switch p {
	case "selected":
		css = GetBg(_ROOM_ITEM_BG_SELECTED)
	case "hover":
		css = GetBg(_ROOM_ITEM_BG_HOVER)
	case "default":
		css = GetBg(_BACKGROUND)
	default:
		log.Println("wtf", p)
	}
	log.Println("set color:", p, css)
	if css != "" {
		this.QWidget_PTR().SetStyleSheet(css)
		for _, w := range this.subws {
			if false {
				w.QWidget_PTR().SetStyleSheet(css)
			}
		}
	}
}
