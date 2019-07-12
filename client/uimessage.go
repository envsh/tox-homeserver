package client

import (
	"fmt"
	"gopp"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unicode"

	"golang.org/x/net/html"
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

	Links []string
	Index int64
}

func Isnumtype(ty reflect.Type) bool {
	switch ty.Kind() {
	case reflect.Uintptr, reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64,
		reflect.Bool, reflect.Int32, reflect.Uint32, reflect.Int16, reflect.Uint16,
		reflect.Float32, reflect.Float64, reflect.Int8, reflect.Uint8:
		return true
	}
	return false
}

func Calcmemlen(v interface{}) int {
	refval := reflect.ValueOf(v)
	refty := refval.Type()

	switch refty.Kind() {
	case reflect.Ptr:
		return Calcmemlen(refval.Elem().Interface())
	case reflect.Struct:
		len1 := int(refty.Size())
		for i := 0; i < refty.NumField(); i++ {
			fldname := refty.Field(i).Name
			fldty := refty.Field(i).Type
			if Isnumtype(fldty) {
			} else {
				if unicode.IsUpper(rune(fldname[0])) {
					len1 += Calcmemlen(refval.Field(i).Interface())
				} else {
					log.Printf("unexport %s.%s %s\n", refty.Name(), fldname, fldty.String())
				}
			}
		}
		return len1
	case reflect.Slice:
		len1 := 0
		for i := 0; i < refval.Len(); i++ {
			len1 += Calcmemlen(refval.Index(i).Interface())
		}
		return len1
	case reflect.Map:
		len1 := 0
		for _, kval := range refval.MapKeys() {
			len1 += Calcmemlen(kval.Interface())
			len1 += Calcmemlen(refval.MapIndex(kval).Interface())
		}
		return len1
	case reflect.Chan:
		len1 := 0
		len1 = refval.Cap() * Calcmemlen(refval.Elem().Interface())
		return len1
	case reflect.UnsafePointer:
		log.Println("oh raw UnsafePointer")
		return 0
	case reflect.String:
		return refval.Len()
	default:
		if Isnumtype(refty) {
			return refty.Bits() / 8
		} else {
			log.Println("todo", refty.Kind())
		}
	}
	return 0
}

func calcmemlen_test() {
	m := &Message{}
	m.Msg = ""
	m.Links = []string{"abc", "efg"}
	rv := Calcmemlen(m)
	log.Println(rv)
	log.Println(Calcmemlen(567.890))
}

// for debug memory
func (this *Message) Memlen() int {
	return Calcmemlen(this)
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
	this.EventId = eventId
	this.Time = time.Now()

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
	this.TimeUi = Time2Today(this.Time)

	// refmtmsgfns := []func(){this.refmtmsgRUser, this.refmtmsgLink}
	refmtmsgfns := []func(){this.refmtmsgRUser, this.findmsgLinks}
	for _, fn := range refmtmsgfns {
		fn()
	}
}
func (this *Message) refmtmsgRUser() {
	if this.Me {
		this.PeerNameUi, this.MsgUi = this.PeerName, this.Msg
	} else {
		// newPeerName, newMsg, _ := bridges.ExtractRealUser(this.PeerName, this.Msg)
		newPeerName, newMsg, _ := bridges.ExtractRealUserMD(this.PeerName, this.Msg)
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
func (this *Message) findmsgLinks() {
	urls := xurls.Strict().FindAllString(this.MsgUi, -1)
	this.Links = urls
}

// TODO
func (this *Message) trimHTMLTags() {
	var trimfn func(n *html.Node)
	trimfn = func(n *html.Node) {
		// if note is script tag
		if n.Type == html.ElementNode && n.Data == "script" {
			n.Parent.RemoveChild(n)
			return // script tag is gone...
		}
		// traverse DOM
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			trimfn(c)
		}
	}
}
func (this *Message) resetTimezone() {
	if runtime.GOOS == "android" {
		// this.Time = this.Time.Add(8 * time.Hour)
	}
}
