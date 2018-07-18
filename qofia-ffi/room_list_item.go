package main

import (
	"fmt"
	"gopp"
	"log"
	"runtime"
	"strings"
	"time"

	"go-purple/msgflt-prpl/bridges"
	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
	store "tox-homeserver/store"
	"tox-homeserver/thspbs"

	humanize "github.com/dustin/go-humanize"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
	"github.com/levigross/grequests"
	"mvdan.cc/xurls"

	"github.com/spf13/afero"
)

type Message struct {
	Msg      string
	PeerName string
	Time     time.Time
	EventId  int64

	Me         bool
	MsgUi      string
	PeerNameUi string
	TimeUi     string
	LastMsgUi  string
	Sent       bool
	UserCode   int64
}

func NewMessageForGroup(evto *thspbs.Event) *Message {
	groupId := evto.Margs[3]
	if thscli.ConferenceIdIsEmpty(groupId) {
		// break
	}

	message := evto.Args[3]
	peerName := evto.Margs[0]
	groupTitle := evto.Margs[2]
	_ = groupTitle
	eventId := gopp.MustInt64(evto.Margs[4])

	this := &Message{}
	this.Msg = message
	this.PeerName = peerName
	this.Time = time.Now()
	this.EventId = eventId

	this.refmtmsg()
	return this
}

func NewMessageForFriend(evto *thspbs.Event) *Message {
	msg := evto.Args[1]
	fname := evto.Margs[0]
	pubkey := evto.Margs[1]
	_, _, _ = msg, fname, pubkey
	eventId := gopp.MustInt64(evto.Margs[2])

	this := &Message{}
	this.Msg = msg
	this.PeerName = fname
	this.Time = time.Now()
	this.EventId = eventId

	this.refmtmsg()
	return this
}

func NewMessageForMe(itext string) *Message {
	msgo := &Message{}
	msgo.Msg = itext
	msgo.PeerName = vtcli.SelfGetName()
	msgo.Time = time.Now()
	msgo.Me = true
	// msgo.UserCode = thscli.NextUserCode(devInfo.Uuid)

	msgo.refmtmsg()
	return msgo
}

func NewMessageForMeFromJson(itext string, eventId int64) *Message {
	msgo := NewMessageForMe(itext)
	msgo.EventId = eventId
	return msgo
}

func (this *Message) refmtmsg() {
	this.LastMsgUi = this.Msg
	this.resetTimezone()

	refmtmsgfns := []func(){this.refmtmsgRUser, this.refmtmsgLink}
	for _, fn := range refmtmsgfns {
		fn()
	}
}
func (this *Message) refmtmsgRUser() {
	if this.Me {
		this.PeerNameUi, this.MsgUi = this.PeerName, this.Msg
	} else {
		newPeerName, newMsg, _ := bridges.ExtractRealUser(this.PeerName, this.Msg)
		this.PeerNameUi = newPeerName
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
func (this *Message) resetTimezone() {
	if runtime.GOOS == "android" {
		// this.Time = this.Time.Add(8 * time.Hour)
	}
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
	uictx.uiw.VerticalLayout_9.InsertWidget(0, item.QWidget_PTR(), 0, 0)
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

	uictx.uiw.VerticalLayout_9.RemoveWidget(item.QWidget_PTR())
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

	floatUnreadCountLabel *qtwidgets.QLabel

	msgitmdl []*MessageItem
	msgos    []*Message

	pressed  bool
	hovered  bool
	isgroup  bool
	frndInfo *thspbs.FriendInfo
	grpInfo  *thspbs.GroupInfo

	unreadedCount int
	totalCount    int
	timeline      thscli.TimeLine

	WaitSyncStoreTimeLineCount int
	LastMsgEventId             int64
}

func NewRoomListItem() *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.initUis()
	this.initEvents()
	return this
}

func NewRoomListItem2(info *thspbs.FriendInfo) *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.initUis()
	this.SetContactInfo(info)
	this.initEvents()
	return this
}

func NewRoomListItem3(info *thspbs.GroupInfo) *RoomListItem {
	this := &RoomListItem{}
	this.Ui_ContactItemView = NewUi_ContactItemView2()
	this.initUis()
	this.SetContactInfo(info)
	this.initEvents()
	return this
}

func (this *RoomListItem) initUis() {
	if !gopp.IsAndroid() {
		this.ToolButton.SetIconSize(qtcore.NewQSize_1(12, 12))
	}
	this.floatUnreadCountLabel = this.floatTextOverWidget(this.ToolButton)
	// this.Ui_ContactItemView.ContactItemView.SetMinimumHeight(20 * 2)
}

func (this *RoomListItem) initEvents() {
	labs := []*qtwidgets.QLabel{this.Label_2, this.Label_3, this.Label_4, this.LabelLastMsgTime}
	for _, lab := range labs {
		lab.Clear()
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
		SetQLabelElideText(this.Label_2, name)
		SetQLabelElideText(this.Label_4, ct.Stmsg)

		avataricon := store.GetFSC().GetFilePath(ct.GetPubkey())
		if gopp.FileExist(avataricon) {
			this.cticon = qtgui.NewQIcon_2(avataricon)
			this.ToolButton_2.SetIcon(this.cticon)
		} else {
			this.cticon = GetIdentIcon(ct.Pubkey)
			this.ToolButton_2.SetIcon(this.cticon)
		}
		if ct.GetConnStatus() == 0 {
			this.sticon = qtgui.NewQIcon_2(":/icons/offline_30.png")
			this.ToolButton.SetIcon(this.sticon)
		}
		if ct.Fnum == thscom.FileHelperFnum {
			this.cticon = qtgui.NewQIcon_2(":/icons/filehelper.png")
			this.ToolButton_2.SetIcon(this.cticon)
		}
	case *thspbs.GroupInfo:
		this.grpInfo = ct
		this.isgroup = true
		this.Label_4.SetHidden(true)
		this.Label_2.SetText(trtxt(ct.GetTitle(), 26))
		this.Label_2.SetToolTip(ct.GetTitle())
		SetQLabelElideText(this.Label_2, ct.Title)

		// this maybe call multiple times, so -20 -20 then, the item is 0 height.
		// this.QWidget_PTR().SetFixedHeight(this.QWidget_PTR().Height() - 20)
		if false {
			this.cticon = qtgui.NewQIcon_2(":/icons/groupgray.png")
		} else {
			this.cticon = GetInitAvatar(gopp.IfElseStr(ct.Title == "", ct.GroupId, ct.Title))
		}
		this.ToolButton_2.SetIcon(this.cticon)
		this.setConnStatus(int32(thscli.CONN_STATUS_UDP))
	default:
		log.Fatalln("wtf")
	}
	this.ToolButton_2.SetToolTip(this.GetName() + "." + gopp.SubStr(this.GetId(), 7))
}

func (this *RoomListItem) AddMessage(msgo *Message, prev bool) {
	// check in list
	for _, msgoe := range this.msgos {
		if msgoe.EventId == msgo.EventId && msgo.EventId != 0 {
			log.Printf("msg already in list: %d, %+v\n", msgo.EventId, msgo)
			return
		}
	}

	if prev {
		this.msgos = append([]*Message{msgo}, this.msgos...)
		msgiw := NewMessageItem()
		msgiw.Sent = msgo.Sent
		msgiw.UserCode = msgo.UserCode
		this.msgitmdl = append([]*MessageItem{msgiw}, this.msgitmdl...)
		this.AddMessageImpl(msgo, msgiw, prev)
		this.UpdateMessageMimeContent(msgo, msgiw)
	} else {
		this.msgos = append(this.msgos, msgo)
		msgiw := NewMessageItem()
		msgiw.Sent = msgo.Sent
		msgiw.UserCode = msgo.UserCode
		this.msgitmdl = append(this.msgitmdl, msgiw)
		this.AddMessageImpl(msgo, msgiw, prev)
		this.UpdateMessageMimeContent(msgo, msgiw)
		// test and update storage's sync info
		if msgo.EventId >= this.timeline.NextBatch {
			this.timeline.NextBatch = msgo.EventId + 1
			this.WaitSyncStoreTimeLineCount += 1
			if this.WaitSyncStoreTimeLineCount >= 1 /*common.PullPageSize*/ {
				this.WaitSyncStoreTimeLineCount = 0
				go hisfet.RefreshPrevStorageTimeLine(&this.timeline, this.GetId(), this.GetName())
			}
		}
	}
}

func (this *RoomListItem) AddMessageImpl(msgo *Message, msgiw *MessageItem, prev bool) {

	showMeIcon := msgo.Me // 是否显示自己的icon。根据是否是自己的消息
	showName := true
	showPeerIcon := !showMeIcon

	msgiw.Label_5.SetText(msgo.MsgUi)
	msgiw.LabelUserName4MessageItem.SetText(fmt.Sprintf("%s", msgo.PeerNameUi))
	msgiw.LabelMsgTime.SetText(Time2Today(msgo.Time))
	msgiw.LabelMsgTime.SetToolTip(gopp.TimeToFmt1(msgo.Time))
	msgiw.ToolButton_3.SetVisible(showMeIcon)
	msgiw.ToolButton_2.SetVisible(showPeerIcon)
	msgiw.LabelUserName4MessageItem.SetVisible(showName)
	msgiw.ToolButton.SetVisible(false)
	if msgo.Me && !msgo.Sent {
		msgiw.LabelSendState.SetPixmap(qtgui.NewQPixmap_3_(":/icons/MessageListSending@2x.png"))
	}
	if msgo.Me {
		msgiw.LabelSendState.SetToolTip(gopp.ToStr(gopp.ToStrs(msgo.Sent, msgo.UserCode)))
	} else /*!msgo.Me*/ {
		// msgiw.LabelSendState.SetVisible(false)
	}

	if uictx.msgwin.item == this {
		vlo3 := uictx.uiw.VerticalLayout_3
		if prev {
			vlo3.InsertWidget__(0, msgiw.QWidget_PTR())
		} else {
			vlo3.Layout().AddWidget(msgiw.QWidget_PTR())
		}
	}
	this.SetLastMsg(fmt.Sprintf("%s: %s",
		gopp.StrSuf4ui(msgo.PeerNameUi, 9, 1), msgo.LastMsgUi), msgo.Time, msgo.EventId)

	this.totalCount += 1
	if uictx.msgwin.item == this {
		uictx.uiw.LabelMsgCount2.SetText(fmt.Sprintf("%3d", this.totalCount))
		uictx.uiw.LabelMsgCount.SetText(fmt.Sprintf("%3d", this.totalCount))
	}
	this.unreadedCount += 1
	this.ToolButton.SetText(fmt.Sprintf("%d", this.unreadedCount))
	// this.floatUnreadCountLabel.SetText(fmt.Sprintf("%d", this.unreadedCount))
}

func (this *RoomListItem) UpdateMessageMimeContent(msgo *Message, msgiw *MessageItem) {
	if !msgo.IsFile() {
		return
	}

	fil := msgo.GetFileInfoLine()
	gopp.NilPrint(fil, msgo.Msg)
	if fil == nil {
		return
	}

	locfname := store.GetFSC().GetFilePath(fil.Md5str)
	rmturl := thscli.HttpFsUrlFor(fil.Md5str)

	reloadMsgItem := func(txt string) { msgiw.Label_5.SetText(txt) }

	locdir := store.GetFSC().GetDir()
	if ok, _ := afero.Exists(afero.NewOsFs(), locfname); ok {
		richtxt := Msg2FileText(fil, locdir)
		log.Println(msgo.Msg, richtxt)
		reloadMsgItem(richtxt)
	} else {
		richtxt := fmt.Sprintf("Loading... %s: %s", fil.Mime, humanize.Bytes(uint64(fil.Length)))
		log.Println(msgo.Msg, richtxt)
		reloadMsgItem(richtxt)

		go func() {
			time.Sleep(3 * time.Second)
			ro := &grequests.RequestOptions{}
			resp, err := grequests.Get(rmturl, ro)
			gopp.ErrPrint(err, rmturl)
			err = resp.DownloadToFile(locfname)
			gopp.ErrPrint(err, rmturl)

			runOnUiThread(func() { reloadMsgItem("Switching...") })
			time.Sleep(3 * time.Second)
			richtxt := Msg2FileText(fil, locdir)
			log.Println(msgo.Msg, richtxt)
			runOnUiThread(func() { reloadMsgItem(richtxt) })

		}()
	}
}

func (this *RoomListItem) UpdateMessageState(msgo *Message) {
	for idx := len(this.msgos) - 1; idx >= 0; idx-- {
		msgo_ := this.msgos[idx]
		if msgo_.UserCode == msgo.UserCode {
			msgo_.EventId = msgo.EventId
			msgo_.Sent = msgo.Sent
			break
		}
	}
	for idx := len(this.msgitmdl) - 1; idx >= 0; idx-- {
		msgitm := this.msgitmdl[idx]
		if msgitm.UserCode == msgo.UserCode {
			if !msgitm.Sent && msgo.Sent {
				msgitm.Sent = msgo.Sent
				msgitm.LabelSendState.Clear()
				msgitm.LabelSendState.SetToolTip(gopp.ToStr(gopp.ToStrs(msgo.Sent, msgo.UserCode)))
			}

			break
		}
	}
}

func (this *RoomListItem) ClearAvatar(frndpk string) {
	this.cticon = GetIdentIcon(frndpk)
	this.ToolButton_2.SetIcon(this.cticon)
	uictx.msgwin.SetIconForItem(this)
}

func (this *RoomListItem) SetAvatar(idico *qtgui.QIcon) {
	this.cticon = idico
	this.ToolButton_2.SetIcon(this.cticon)
	uictx.msgwin.SetIconForItem(this)
}

func (this *RoomListItem) SetAvatarForId(frndpk string) {
	locfname := store.GetFSC().GetFilePath(frndpk)
	idico := qtgui.NewQIcon_2(locfname)
	this.SetAvatar(idico)
}

func (this *RoomListItem) SetAvatarForMessage(msgo *Message, frndpk string) {
	fil := msgo.GetFileInfoLine()
	gopp.NilPrint(fil, msgo.Msg)
	if fil == nil {
		return
	}

	locfname := store.GetFSC().GetFilePath(frndpk)
	rmturl := thscli.HttpFsUrlFor(frndpk)

	setFriendIcon := func(thefname string) {
		icon := qtgui.NewQIcon_2(thefname)
		if icon != nil && !icon.IsNull() {
			this.SetAvatar(icon)
		} else {
			log.Println("Friend icon not supported:", locfname)
		}
	}
	if fil.Length == 0 { // clear avatar
		this.ClearAvatar(frndpk)
		return
	}
	go func() {
		ro := &grequests.RequestOptions{}
		resp, err := grequests.Get(rmturl, ro)
		gopp.ErrPrint(err, rmturl)
		err = resp.DownloadToFile(locfname)
		gopp.ErrPrint(err, rmturl)

		runOnUiThread(func() { setFriendIcon(locfname) })
	}()
}

func (this *RoomListItem) FindMessageByUserCode(userCode int64) *Message {
	for idx := len(this.msgos) - 1; idx >= 0; idx-- {
		msgo_ := this.msgos[idx]
		if msgo_.UserCode == userCode {
			return msgo_
		}
	}
	return nil
}

func (this *RoomListItem) FindMessageViewByEventId(eventId int64) *MessageItem {
	for idx := len(this.msgos) - 1; idx >= 0; idx-- {
		msgo_ := this.msgos[idx]
		if msgo_.EventId == eventId {
			return this.msgitmdl[idx]
		}
	}
	return nil
}

// TODO 计算是否省略掉显示与上一条相同的用户名
func (this *RoomListItem) AddMessageHiddenCloseSameUser(prev bool) {
	// prev is true, compare [0], [1]
	// prev is false, compare [len-2], [len-1]
	if len(this.msgos) < 2 {
		return
	}

	var m0, m1 *Message
	if prev {
		m0 = this.msgos[0]
		m1 = this.msgos[1]
	} else {
		m0 = this.msgos[len(this.msgos)-2]
		m1 = this.msgos[len(this.msgos)-1]
	}

	if m0.PeerNameUi == m1.PeerNameUi {
		// can not get Ui_MessageItemView
	}
}

func (this *RoomListItem) GetName() string {
	return gopp.IfElseStr(this.isgroup, this.grpInfo.GetTitle(), this.frndInfo.GetName())
}

func (this *RoomListItem) GetId() string {
	if this.isgroup {
		// log.Println(this.grpInfo.GetGroupId(), this.grpInfo.Title)
	}
	return gopp.IfElseStr(this.isgroup, this.grpInfo.GetGroupId(), this.frndInfo.GetPubkey())
}

func (this *RoomListItem) GetNum() uint32 {
	return uint32(gopp.IfElseInt(this.isgroup, int(this.grpInfo.GetGnum()), int(this.frndInfo.GetFnum())))
}

func (this *RoomListItem) UpdateName(name string) {
	if this.isgroup {
		if this.grpInfo.Title != name {
			this.grpInfo.Title = name
			this.SetContactInfo(this.grpInfo)
			// this.Label_2.SetText(gopp.StrSuf4ui(name, 26))
			// this.Label_2.SetToolTip(name)
			// this.ToolButton_2.SetToolTip(name + "." + this.GetId()[:7])
		}
	} else {
		if this.frndInfo.Name != name {
			this.frndInfo.Name = name
			this.SetContactInfo(this.frndInfo)
			// this.Label_2.SetText(gopp.StrSuf4ui(name, 26))
			// this.Label_2.SetToolTip(name)
			// this.ToolButton_2.SetToolTip(name + "." + this.GetId()[:7])
		}
	}
}
func (this *RoomListItem) UpdateStatusMessage(statusText string) {
	if !this.isgroup {
		if this.frndInfo.Stmsg != statusText {
			this.frndInfo.Stmsg = statusText
			this.SetContactInfo(this.frndInfo)
		}
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

// TODO how custom setting this
func init() {
	if runtime.GOOS == "android" {
		secondsEastOfUTC := int((8 * time.Hour).Seconds())
		cqzone := time.FixedZone("Chongqing", secondsEastOfUTC)
		time.Local = cqzone
	}
}

// 两类时间，server time, client time
func (this *RoomListItem) SetLastMsg(msg string, tm time.Time, eventId int64) {
	if this.LastMsgEventId > eventId {
		return
	}

	this.LastMsgEventId = eventId
	refmter := func(s string) string {
		// s = gopp.StrSuf4ui(s, 36)
		return strings.Replace(s, "\n", " ", -1)
	}
	cmsg := refmter(msg)
	this.Label_3.SetText(cmsg)
	this.Label_3.SetToolTip(msg)
	SetQLabelElideText(this.Label_3, cmsg)
	this.LabelLastMsgTime.SetText(Time2TodayMinute(tm))
	this.LabelLastMsgTime.SetToolTip(gopp.TimeToFmt1(tm))
}

func (this *RoomListItem) SetPressState(pressed bool) {
	changed := this.pressed != pressed
	// log.Println("changed:", changed, "pressed:", pressed, this.GetName())
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
	// log.Println("set color:", p, css)
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
	if !this.isgroup {
		this.frndInfo.ConnStatus = st
	}
	iconNames := map[string]map[int]string{
		"friend": {thscli.CONN_STATUS_NONE: ":/icons/offline_30.png",
			thscli.CONN_STATUS_TCP: ":/icons/online_30.png",
			thscli.CONN_STATUS_UDP: ":/icons/online_30.png"},
		"group": {0: ":/icons/dot_groupchat.png"},
		// "group": {0: ":/icons/online_30.png"},
	}
	if this.isgroup {
		if false {
			// android not run svg well now???
			pxm := qtgui.NewQPixmap_3_(iconNames["group"][0]).Scaled__(9, 9)
			this.sticon = qtgui.NewQIcon_1(pxm)
		} else { // backup method
			this.sticon = qtgui.NewQIcon_2(iconNames["group"][0])
		}
	} else {
		this.sticon = qtgui.NewQIcon_2(iconNames["friend"][int(st)])
	}
	this.ToolButton.SetIcon(this.sticon)
}

func (this *RoomListItem) setUserStatus(st int) {
	clricon := qtgui.NewQIcon_2(":/icons/offline_30.png")
	switch st {
	case thscli.USER_STATUS_NONE:
		if !this.isgroup {
			this.setConnStatus(this.frndInfo.ConnStatus)
		}
	case thscli.USER_STATUS_AWAY:
		this.sticon = qtgui.NewQIcon_2(":/icons/dot_away.png")
		this.ToolButton.SetIcon(clricon)
		this.ToolButton.SetIcon(this.sticon)
	case thscli.USER_STATUS_BUSY:
		this.sticon = qtgui.NewQIcon_2(":/icons/dot_busy.png")
		this.ToolButton.SetIcon(clricon)
		this.ToolButton.SetIcon(this.sticon)
	}
}

func (this *RoomListItem) floatTextOverWidget(w qtwidgets.QWidget_ITF) *qtwidgets.QLabel {
	lo := qtwidgets.NewQVBoxLayout_1(w)
	lo.SetContentsMargins(0, 0, 0, 0)
	lo.AddStretch__()
	lab := qtwidgets.NewQLabel__()
	lo.AddWidget(lab, 0, qtcore.Qt__AlignCenter)
	return lab
}

func (this *RoomListItem) SetPeerCount(n int) {
	if uictx.msgwin.item == this {
		uictx.msgwin.SetPeerCount(n)
	}
}
