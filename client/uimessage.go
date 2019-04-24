package client

import (
	"fmt"

	"gopp"
	"runtime"
	"strings"
	"time"

	"mvdan.cc/xurls"

	"go-purple/msgflt-prpl/bridges"
	"tox-homeserver/thspbs"
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
	if ConferenceIdIsEmpty(groupId) {
		// break
	}

	message := evto.Args[3]
	peerName := evto.Margs[0]
	groupTitle := evto.Margs[2]
	peerId := evto.Margs[1]
	_ = groupTitle
	eventId := gopp.MustInt64(evto.Margs[4])

	this := &Message{}
	this.Msg = message
	this.PeerName = peerName
	this.Time = time.Now()
	this.EventId = eventId

	if peerName == "" {
		this.PeerName = peerId[:8]
	}
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
	var vtcli = appctx.vtcli
	msgo.PeerName = vtcli.SelfGetName()
	msgo.Time = time.Now()
	msgo.Me = true
	// msgo.UserCode = NextUserCode(devInfo.Uuid)

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
