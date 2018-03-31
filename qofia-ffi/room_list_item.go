package main

import (
	"fmt"
	"gopp"
	"log"
	"os"
	"strings"
	"time"

	"go-purple/msgflt-prpl/bridges"
	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
	"mvdan.cc/xurls"
)

type Message struct {
	Msg  string
	Peer string
	Time time.Time

	Me        bool
	MsgUi     string
	PeerUi    string
	TimeUi    string
	LastMsgUi string
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

	this.refmtmsg()
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

	this.refmtmsg()
	return this
}

func (this *Message) refmtmsg() {
	this.LastMsgUi = this.Msg

	refmtmsgfns := []func(){this.refmtmsgRUser, this.refmtmsgLink}
	for _, fn := range refmtmsgfns {
		fn()
	}
}
func (this *Message) refmtmsgRUser() {
	if this.Me {
		this.PeerUi, this.MsgUi = this.Peer, this.Msg
	} else {
		newPeer, newMsg, _ := bridges.ExtractRealUser(this.Peer, this.Msg)
		this.PeerUi = newPeer
		this.MsgUi = newMsg
		this.LastMsgUi = newMsg
	}
}
func (this *Message) refmtmsgLink() {
	urls := xurls.Strict().FindAllString(this.MsgUi, -1)
	s := this.MsgUi
	for _, u := range urls {
		s = strings.Replace(s, u, fmt.Sprintf(`<a href="%s">%s</a>`, u, u), -1)
	}
	this.MsgUi = s
}

////////////////
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

func (this *RoomListMan) onClicked(item *RoomListItem, pos *qtcore.QPoint) {
	uictx.mw.switchUiStack(UIST_MESSAGEUI)

	uictx.msgwin.SetRoom(item)
}
func (this *RoomListMan) onLongTouched(item *RoomListItem, gpos *qtcore.QPoint) {
	item.OnContextMenu2(gpos)
}

/////////////////
type RoomListItem struct {
	*Ui_ContactItemView

	OnConextMenu func(w *qtwidgets.QWidget, pos *qtcore.QPoint)

	cticon *qtgui.QIcon
	sticon *qtgui.QIcon
	subws  []qtwidgets.QWidget_ITF
	menu   *qtwidgets.QMenu

	msgitmdl []*Ui_MessageItemView
	msgos    []*Message

	pressed  bool
	hovered  bool
	isgroup  bool
	frndInfo *thspbs.FriendInfo
	grpInfo  *thspbs.GroupInfo

	unreadedCount int
	totalCount    int
	peerCount     int
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

	labs := []*qtwidgets.QLabel{this.Label_2, this.Label_3, this.Label_4, this.Label_5, this.LabelLastMsgTime}
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
		uictx.gtreco.onMousePress(this, event)
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
		uictx.gtreco.onMouseRelease(this, event)
	}
	onMouseMove := func(event *qtgui.QMouseEvent) {
		uictx.gtreco.onMouseMove(this, event)
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

func (this *RoomListItem) OnContextMenu2(gpos *qtcore.QPoint) {
	w := this.ContactItemView
	if this.OnConextMenu != nil {
		this.OnConextMenu(w, gpos)
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
		if ct.GetConnStatus() == 0 {
			this.sticon = qtgui.NewQIcon_2(":/icons/offline_30.png")
			this.ToolButton.SetIcon(this.sticon)
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
		this.peerCount = len(ct.Members)
		if this.peerCount > 0 {
		}
	default:
		log.Fatalln("wtf")
	}
	this.ToolButton_2.SetToolTip(this.GetName() + "." + gopp.SubStr(this.GetId(), 7))
}

func (this *RoomListItem) AddMessage(msgo *Message) {
	this.msgos = append(this.msgos, msgo)
	msgiw := NewUi_MessageItemView2()
	this.msgitmdl = append(this.msgitmdl, msgiw)

	lastSame := func(peerUi string) bool {
		if len(this.msgos) >= 2 {
			if this.msgos[len(this.msgos)-2].PeerUi == peerUi {
				return true
			}
		}
		return false
	}

	showMeIcon := msgo.Me // 是否显示自己的icon。根据是否是自己的消息
	showName := !lastSame(msgo.PeerUi)
	showPeerIcon := !lastSame(msgo.PeerUi) && !msgo.Me // 是否显示对方的icon。根据前一条消息判断

	msgiw.Label_5.SetText(msgo.MsgUi)
	msgiw.Label_3.SetText(fmt.Sprintf("%s", msgo.PeerUi))
	msgiw.LabelMsgTime.SetText(Time2Today(msgo.Time))
	msgiw.LabelMsgTime.SetToolTip(gopp.TimeToFmt1(msgo.Time))
	msgiw.ToolButton_3.SetVisible(showMeIcon)
	msgiw.ToolButton_2.SetVisible(showPeerIcon)
	msgiw.Label_3.SetVisible(showName)
	msgiw.ToolButton.SetVisible(false)

	if uictx.msgwin.item == this {
		vlo8 := uictx.uiw.VerticalLayout_8
		vlo8.Layout().AddWidget(msgiw.QWidget_PTR())
	}
	this.SetLastMsg(fmt.Sprintf("%s: %s", gopp.StrSuf4ui(msgo.PeerUi, 8, 1), msgo.LastMsgUi))

	this.totalCount += 1
	if uictx.msgwin.item == this {
		uictx.uiw.LabelMsgCount2.SetText(fmt.Sprintf("%3d", this.totalCount))
		uictx.uiw.LabelMsgCount.SetText(fmt.Sprintf("%3d", this.totalCount))
	}
	this.unreadedCount += 1
	this.ToolButton.SetText(fmt.Sprintf("%d", this.unreadedCount))
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
func (this *RoomListItem) UpdateReaded() {
	mw := uictx.mw
	curval := mw.ScrollArea_2.VerticalScrollBar().Value()
	minval := mw.ScrollArea_2.VerticalScrollBar().Minimum()
	maxval := mw.ScrollArea_2.VerticalScrollBar().Maximum()
	log.Println(this.unreadedCount, this.pressed, curval, minval, maxval)

	if this.unreadedCount > 0 && this.pressed {
		if curval == maxval || maxval == -1 {
			this.unreadedCount = 0
			this.ToolButton.SetText("")
		}
	}
}

func (this *RoomListItem) SetLastMsg(msg string) {
	cmsg := msg
	this.Label_3.SetText(gopp.StrSuf4ui(cmsg, 36))
	this.Label_3.SetToolTip(cmsg)
	tm := time.Now()
	this.LabelLastMsgTime.SetText(Time2TodayMinute(tm))
	this.LabelLastMsgTime.SetToolTip(gopp.TimeToFmt1(tm))
}

func (this *RoomListItem) SetPressState(pressed bool) {
	changed := this.pressed != pressed
	log.Println("changed:", changed, "pressed:", pressed, this.GetName())
	if changed {
		this.pressed = pressed
		this.SetBgColor(gopp.IfElseStr(pressed, "selected", "default"))
	}

	// uictx.mw.switchUiStack(4)
	if changed {
		// uictx.msgwin.SetRoom(this)
	}
	if pressed {
		this.UpdateReaded()
	}
}

func (this *RoomListItem) IsPressStateChanged(pressed bool) bool {
	return this.pressed != pressed
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
	po := qtcore.NewQVariant_15(p)
	this.ContactItemView.SetProperty("istate", po)
	syl := this.ContactItemView.Style()
	syl.Unpolish(this.ContactItemView)
	syl.Polish(this.ContactItemView)
	if css != "" {
		if false {
			this.QWidget_PTR().SetStyleSheet(css)
		}
		for _, w := range this.subws {
			if false {
				w.QWidget_PTR().SetStyleSheet(css)
			}
		}
	}
}

func (this *RoomListItem) setConnStatus(st int32) {
	if st > 0 {
		this.sticon = qtgui.NewQIcon_2(":/icons/online_30.png")
		this.ToolButton.SetIcon(this.sticon)
	} else {
		this.sticon = qtgui.NewQIcon_2(":/icons/offline_30.png")
		this.ToolButton.SetIcon(this.sticon)
	}
}
