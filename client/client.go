package client

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"

	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"github.com/nats-io/go-nats"
	"google.golang.org/grpc"
)

// friend callback type
type cb_friend_request_ftype func(this *LigTox, pubkey string, message string, userData interface{})
type cb_friend_message_ftype func(this *LigTox, friendNumber uint32, message string, userData interface{})
type cb_friend_name_ftype func(this *LigTox, friendNumber uint32, newName string, userData interface{})
type cb_friend_status_message_ftype func(this *LigTox, friendNumber uint32, newStatus string, userData interface{})
type cb_friend_status_ftype func(this *LigTox, friendNumber uint32, status int, userData interface{})
type cb_friend_connection_status_ftype func(this *LigTox, friendNumber uint32, status int, userData interface{})
type cb_friend_typing_ftype func(this *LigTox, friendNumber uint32, isTyping uint8, userData interface{})
type cb_friend_read_receipt_ftype func(this *LigTox, friendNumber uint32, receipt uint32, userData interface{})
type cb_friend_lossy_packet_ftype func(this *LigTox, friendNumber uint32, data string, userData interface{})
type cb_friend_lossless_packet_ftype func(this *LigTox, friendNumber uint32, data string, userData interface{})

// self callback type
type cb_self_connection_status_ftype func(this *LigTox, status int, userData interface{})

// file callback type
type cb_file_recv_control_ftype func(this *LigTox, friendNumber uint32, fileNumber uint32,
	control int, userData interface{})
type cb_file_recv_ftype func(this *LigTox, friendNumber uint32, fileNumber uint32, kind uint32, fileSize uint64,
	fileName string, userData interface{})
type cb_file_recv_chunk_ftype func(this *LigTox, friendNumber uint32, fileNumber uint32, position uint64,
	data []byte, userData interface{})
type cb_file_chunk_request_ftype func(this *LigTox, friend_number uint32, file_number uint32, position uint64,
	length int, user_data interface{})

type cb_baseinfo_ftype func(this *LigTox, bi *thspbs.BaseInfo, user_data interface{})

type LigTox struct {
	ToxId    string
	Status   int
	Stmsg    string
	Binfo    *thspbs.BaseInfo
	bemsgs   [][]byte
	bemsgsmu sync.RWMutex
	OnNewMsg func()

	rpcli    *grpc.ClientConn
	ntscli   *nats.Conn
	wsc4push *websocket.Conn
	wsc4rpc  *websocket.Conn
	usewstp  bool // use websocket transport for replace grpc+nats

	// some callbacks, should be private. &fn => ud
	cb_friend_requests           map[unsafe.Pointer]interface{}
	cb_friend_messages           map[unsafe.Pointer]interface{}
	cb_friend_names              map[unsafe.Pointer]interface{}
	cb_friend_status_messages    map[unsafe.Pointer]interface{}
	cb_friend_statuss            map[unsafe.Pointer]interface{}
	cb_friend_connection_statuss map[unsafe.Pointer]interface{}
	cb_friend_typings            map[unsafe.Pointer]interface{}
	cb_friend_read_receipts      map[unsafe.Pointer]interface{}
	cb_friend_lossy_packets      map[unsafe.Pointer]interface{}
	cb_friend_lossless_packets   map[unsafe.Pointer]interface{}
	cb_self_connection_statuss   map[unsafe.Pointer]interface{}

	cb_conference_invites          map[unsafe.Pointer]interface{}
	cb_conference_messages         map[unsafe.Pointer]interface{}
	cb_conference_actions          map[unsafe.Pointer]interface{}
	cb_conference_titles           map[unsafe.Pointer]interface{}
	cb_conference_namelist_changes map[unsafe.Pointer]interface{}

	cb_file_recv_controls  map[unsafe.Pointer]interface{}
	cb_file_recvs          map[unsafe.Pointer]interface{}
	cb_file_recv_chunks    map[unsafe.Pointer]interface{}
	cb_file_chunk_requests map[unsafe.Pointer]interface{}

	cb_baseinfos map[unsafe.Pointer]interface{}

	cb_iterate_data              interface{}
	cb_conference_message_setted bool
}

func NewLigTox() *LigTox {
	this := &LigTox{}
	this.bemsgs = make([][]byte, 0)
	this.initCbmap()

	// this.initConnections()
	log.Println("connecting grpc:", thscom.GrpcAddr)
	rpcli, err := grpc.Dial(thscom.GrpcAddr, grpc.WithInsecure())
	gopp.ErrPrint(err, rpcli)
	this.rpcli = rpcli

	uo, err := url.Parse(thscom.GnatsAddr)
	gopp.ErrPrint(err, uo)
	ntscli, err := nats.Connect(thscom.GnatsAddr)
	gopp.ErrPrint(err, thscom.GnatsAddr)
	if err == nats.ErrNoServers {
		// try local
		ntscli, err = nats.Connect(thscom.GnatsAddrlo)
	}
	this.ntscli = ntscli
	log.Println("gnats connected:", ntscli.IsConnected(), thscom.GnatsAddr)
	ntscli.Subscribe(thscom.CBEventBusName, this.onBackendEventNats)

	this.wsc4rpc, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/toxhsrpc", thscom.WSAddr), nil)
	gopp.ErrPrint(err, thscom.WSAddr)
	this.wsc4push, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/toxhspush", thscom.WSAddr), nil)
	gopp.ErrPrint(err, thscom.WSAddr)
	go this.serveBackendEventWS()
	// TODO reconnect

	return this
}

func (this *LigTox) initCbmap() {
	this.cb_friend_requests = make(map[unsafe.Pointer]interface{})
	this.cb_friend_messages = make(map[unsafe.Pointer]interface{})
	this.cb_friend_names = make(map[unsafe.Pointer]interface{})
	this.cb_friend_status_messages = make(map[unsafe.Pointer]interface{})
	this.cb_friend_statuss = make(map[unsafe.Pointer]interface{})
	this.cb_friend_connection_statuss = make(map[unsafe.Pointer]interface{})
	this.cb_friend_typings = make(map[unsafe.Pointer]interface{})
	this.cb_friend_read_receipts = make(map[unsafe.Pointer]interface{})
	this.cb_friend_lossy_packets = make(map[unsafe.Pointer]interface{})
	this.cb_friend_lossless_packets = make(map[unsafe.Pointer]interface{})
	this.cb_self_connection_statuss = make(map[unsafe.Pointer]interface{})

	this.cb_conference_invites = make(map[unsafe.Pointer]interface{})
	this.cb_conference_messages = make(map[unsafe.Pointer]interface{})
	this.cb_conference_actions = make(map[unsafe.Pointer]interface{})
	this.cb_conference_titles = make(map[unsafe.Pointer]interface{})
	this.cb_conference_namelist_changes = make(map[unsafe.Pointer]interface{})

	this.cb_file_recv_controls = make(map[unsafe.Pointer]interface{})
	this.cb_file_recvs = make(map[unsafe.Pointer]interface{})
	this.cb_file_recv_chunks = make(map[unsafe.Pointer]interface{})
	this.cb_file_chunk_requests = make(map[unsafe.Pointer]interface{})

	this.cb_baseinfos = make(map[unsafe.Pointer]interface{})

}

func (this *LigTox) onBackendEventNats(msg *nats.Msg) {
	log.Println("nats:", msg.Subject, string(msg.Data))
	jso, err := simplejson.NewJson(msg.Data)
	gopp.ErrPrint(err)

	this.onBackendEventDeduped(jso, msg.Data)
}

// TODO 两个同时接收导致重复消息
// should block
func (this *LigTox) serveBackendEventWS() {
	for {
		c := this.wsc4push
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("wsrecv: %s\n", message)

		jso, err := simplejson.NewJson(message)
		gopp.ErrPrint(err)
		if rdatao, ok := jso.CheckGet("data"); ok {
			rmessage, _ := rdatao.Encode()
			this.onBackendEventDeduped(rdatao, rmessage)
		} else if _, ok := jso.CheckGet("name"); ok {
			this.onBackendEventDeduped(jso, message)
		} else {
			log.Println("Unknown packet:", string(message))
		}
	}
	log.Println("done")
}

var evtmd5smu sync.Mutex
var evtmd5s = make(map[string]time.Time)

func (this *LigTox) onBackendEventDeduped(jso *simplejson.Json, data []byte) {
	data, err := jso.Encode()
	gopp.ErrPrint(err)
	md5b := md5.New().Sum(data)
	isdup := false
	evtmd5smu.Lock()
	if tm, ok := evtmd5s[string(md5b)]; ok {
		if time.Now().Sub(tm).Seconds() < 10 {
			// ok dup message
			isdup = true
		}
	} else {
		evtmd5s[string(md5b)] = time.Now()
	}
	for s, tm := range evtmd5s {
		if time.Now().Sub(tm).Seconds() > 30 {
			delete(evtmd5s, s)
		}
	}
	evtmd5smu.Unlock()
	if !isdup {
		this.onBackendEvent(jso, data)
	}
}

func (this *LigTox) onBackendEvent(jso *simplejson.Json, data []byte) {

	defer func() {
		this.bemsgsmu.Lock()
		this.bemsgs = append(this.bemsgs, data)
		if len(this.bemsgs) > 500 {
			log.Println("queue too large.", len(this.bemsgs))
			this.bemsgs = this.bemsgs[len(this.bemsgs)-500:]
		}
		this.bemsgsmu.Unlock()
		if this.OnNewMsg != nil {
			this.OnNewMsg()
		}
	}()

	argso := jso.Get("args")
	evtName := jso.Get("name").MustString()
	switch evtName {
	case "FriendConnectionStatus":
		fnum := gopp.MustUint32(argso.GetIndex(0).MustString())
		st := gopp.MustInt(argso.GetIndex(1).MustString())
		this.callbackFriendConnectionStatus(fnum, st)
	case "FriendMessage":
		fnum := gopp.MustUint32(argso.GetIndex(0).MustString())
		log.Println(fnum)
		this.callbackFriendMessage(fnum, 0, argso.GetIndex(1).MustString())
	case "ConferenceInvite":
		fnum := gopp.MustUint32(argso.GetIndex(0).MustString())
		itype := gopp.MustInt(argso.GetIndex(1).MustString())
		cookie := argso.GetIndex(2).MustString()
		this.callbackConferenceInvite(fnum, itype, cookie)
	case "ConferenceMessage":
		gnum := gopp.MustUint32(argso.GetIndex(0).MustString())
		pnum := gopp.MustUint32(argso.GetIndex(1).MustString())
		mtype := gopp.MustInt(argso.GetIndex(2).MustString())
		msg := argso.GetIndex(3).MustString()
		this.callbackConferenceMessage(gnum, pnum, mtype, msg)
	case "ConferenceTitle":
		gnum := gopp.MustUint32(argso.GetIndex(0).MustString())
		pnum := gopp.MustUint32(argso.GetIndex(1).MustString())
		title := argso.GetIndex(2).MustString()
		this.callbackConferenceTitle(gnum, pnum, title)
	case "ConferenceNameListChange":
		gnum := gopp.MustUint32(argso.GetIndex(0).MustString())
		pnum := gopp.MustUint32(argso.GetIndex(1).MustString())
		change := gopp.MustInt(argso.GetIndex(2).MustString())
		this.callbackConferenceNameListChange(gnum, pnum, change)
	}
}

func (this *LigTox) rmtCall(args *thspbs.Event) (*thspbs.Event, error) {
	if this.usewstp {
		data, err := json.Marshal(args)
		gopp.ErrPrint(err)
		err = this.wsc4rpc.WriteMessage(websocket.TextMessage, data)
		gopp.ErrPrint(err)
		mt, rdata, err := this.wsc4rpc.ReadMessage()
		gopp.ErrPrint(err, mt)
		rsp := &thspbs.Event{}
		err = json.Unmarshal(rdata, rsp)
		gopp.ErrPrint(err)
		return rsp, err
	} else {
		cli := thspbs.NewToxhsClient(this.rpcli)
		rsp, err := cli.RmtCall(context.Background(), args)
		return rsp, err
	}
}

func (this *LigTox) GetBaseInfo() {
	if this.usewstp {
		in := &thspbs.Event{}
		in.Name = "GetBaseInfo"
		rsp, err := this.rmtCall(in)
		gopp.ErrPrint(err)

		binfo := &thspbs.BaseInfo{}
		err = json.Unmarshal([]byte(rsp.Args[0]), binfo)
		gopp.ErrPrint(err)
		this.ParseBaseInfo(binfo)
	} else {
		cli := thspbs.NewToxhsClient(this.rpcli)
		in := &thspbs.EmptyReq{}
		bi, err := cli.GetBaseInfo(context.Background(), in)
		gopp.ErrPrint(err)
		this.ParseBaseInfo(bi)
	}
}

func (this *LigTox) ParseBaseInfo(bi *thspbs.BaseInfo) {
	this.Binfo = bi
	appctx.persistBaseInfo(bi)
	this.callbackBaseInfo(bi)
}

func (this *LigTox) GetNextBackenEvent() []byte {
	this.bemsgsmu.Lock()
	defer this.bemsgsmu.Unlock()

	if len(this.bemsgs) > 0 {
		ret := this.bemsgs[0]
		this.bemsgs = this.bemsgs[1:]
		return ret
	}
	return nil
}

func (this *LigTox) putcbevts(cbfn func()) { cbfn() }

///
func (this *LigTox) callbackBaseInfo(bi *thspbs.BaseInfo) {
	for cbfni, ud := range this.cb_baseinfos {
		cbfn := *(*cb_baseinfo_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, bi, ud) })
	}
}

func (this *LigTox) CallbackBaseInfo(cbfn cb_baseinfo_ftype, userData interface{}) {
	this.CallbackBaseInfoAdd(cbfn, userData)
}
func (this *LigTox) CallbackBaseInfoAdd(cbfn cb_baseinfo_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_baseinfos[cbfnp]; ok {
		return
	}
	this.cb_baseinfos[cbfnp] = userData
}

///
func (this *LigTox) callbackFriendRequest(pubkey, message string) {

	for cbfni, ud := range this.cb_friend_requests {
		cbfn := *(*cb_friend_request_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, pubkey, message, ud) })
	}
}

func (this *LigTox) CallbackFriendRequest(cbfn cb_friend_request_ftype, userData interface{}) {
	this.CallbackFriendRequestAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendRequestAdd(cbfn cb_friend_request_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_requests[cbfnp]; ok {
		return
	}
	this.cb_friend_requests[cbfnp] = userData

}

func (this *LigTox) callbackFriendMessage(a0 uint32, mtype int, message string) {

	for cbfni, ud := range this.cb_friend_messages {
		cbfn := *(*cb_friend_message_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), message, ud) })
	}
}

func (this *LigTox) CallbackFriendMessage(cbfn cb_friend_message_ftype, userData interface{}) {
	this.CallbackFriendMessageAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendMessageAdd(cbfn cb_friend_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_messages[cbfnp]; ok {
		return
	}
	this.cb_friend_messages[cbfnp] = userData

}

func (this *LigTox) callbackFriendName(a0 uint32, name string) {

	for cbfni, ud := range this.cb_friend_names {
		cbfn := *(*cb_friend_name_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), name, ud) })
	}
}

func (this *LigTox) CallbackFriendName(cbfn cb_friend_name_ftype, userData interface{}) {
	this.CallbackFriendNameAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendNameAdd(cbfn cb_friend_name_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_names[cbfnp]; ok {
		return
	}
	this.cb_friend_names[cbfnp] = userData

}

func (this *LigTox) callbackFriendStatusMessage(a0 uint32, stmsg string) {

	for cbfni, ud := range this.cb_friend_status_messages {
		statusText := stmsg
		cbfn := *(*cb_friend_status_message_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), statusText, ud) })
	}
}

func (this *LigTox) CallbackFriendStatusMessage(cbfn cb_friend_status_message_ftype, userData interface{}) {
	this.CallbackFriendStatusMessageAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendStatusMessageAdd(cbfn cb_friend_status_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_status_messages[cbfnp]; ok {
		return
	}
	this.cb_friend_status_messages[cbfnp] = userData

}

func (this *LigTox) callbackFriendStatus(a0 uint32, a1 int) {

	for cbfni, ud := range this.cb_friend_statuss {
		cbfn := *(*cb_friend_status_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), int(a1), ud) })
	}
}

func (this *LigTox) CallbackFriendStatus(cbfn cb_friend_status_ftype, userData interface{}) {
	this.CallbackFriendStatusAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendStatusAdd(cbfn cb_friend_status_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_statuss[cbfnp]; ok {
		return
	}
	this.cb_friend_statuss[cbfnp] = userData
}

func (this *LigTox) callbackFriendConnectionStatus(a0 uint32, a1 int) {

	for cbfni, ud := range this.cb_friend_connection_statuss {
		cbfn := *(*cb_friend_connection_status_ftype)((unsafe.Pointer)(cbfni))
		this.putcbevts(func() { cbfn(this, uint32(a0), int(a1), ud) })
	}
}

func (this *LigTox) CallbackFriendConnectionStatus(cbfn cb_friend_connection_status_ftype, userData interface{}) {
	this.CallbackFriendConnectionStatusAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendConnectionStatusAdd(cbfn cb_friend_connection_status_ftype, userData interface{}) {
	cbfnp := unsafe.Pointer(&cbfn)
	if _, ok := this.cb_friend_connection_statuss[cbfnp]; ok {
		return
	}
	this.cb_friend_connection_statuss[cbfnp] = userData

}

func (this *LigTox) callbackFriendTyping(a0 uint32, a1 uint8) {

	for cbfni, ud := range this.cb_friend_typings {
		cbfn := *(*cb_friend_typing_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), uint8(a1), ud) })
	}
}

func (this *LigTox) CallbackFriendTyping(cbfn cb_friend_typing_ftype, userData interface{}) {
	this.CallbackFriendTypingAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendTypingAdd(cbfn cb_friend_typing_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_typings[cbfnp]; ok {
		return
	}
	this.cb_friend_typings[cbfnp] = userData

}

func (this *LigTox) callbackFriendReadReceipt(a0 uint32, a1 uint32) {

	for cbfni, ud := range this.cb_friend_read_receipts {
		cbfn := *(*cb_friend_read_receipt_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), uint32(a1), ud) })
	}
}

func (this *LigTox) CallbackFriendReadReceipt(cbfn cb_friend_read_receipt_ftype, userData interface{}) {
	this.CallbackFriendReadReceiptAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendReadReceiptAdd(cbfn cb_friend_read_receipt_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_read_receipts[cbfnp]; ok {
		return
	}
	this.cb_friend_read_receipts[cbfnp] = userData

}

func (this *LigTox) callbackFriendLossyPacket(a0 uint32, msg string) {

	for cbfni, ud := range this.cb_friend_lossy_packets {
		cbfn := *(*cb_friend_lossy_packet_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), msg, ud) })
	}
}

func (this *LigTox) CallbackFriendLossyPacket(cbfn cb_friend_lossy_packet_ftype, userData interface{}) {
	this.CallbackFriendLossyPacketAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendLossyPacketAdd(cbfn cb_friend_lossy_packet_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_lossy_packets[cbfnp]; ok {
		return
	}
	this.cb_friend_lossy_packets[cbfnp] = userData

}

func (this *LigTox) callbackFriendLosslessPacket(a0 uint32, msg string) {

	for cbfni, ud := range this.cb_friend_lossless_packets {
		cbfn := *(*cb_friend_lossless_packet_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), msg, ud) })
	}
}

func (this *LigTox) CallbackFriendLosslessPacket(cbfn cb_friend_lossless_packet_ftype, userData interface{}) {
	this.CallbackFriendLosslessPacketAdd(cbfn, userData)
}
func (this *LigTox) CallbackFriendLosslessPacketAdd(cbfn cb_friend_lossless_packet_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_lossless_packets[cbfnp]; ok {
		return
	}
	this.cb_friend_lossless_packets[cbfnp] = userData

}

func (this *LigTox) callbackSelfConnectionStatus(status int) {

	for cbfni, ud := range this.cb_self_connection_statuss {
		cbfn := *(*cb_self_connection_status_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, int(status), ud) })
	}
}

func (this *LigTox) CallbackSelfConnectionStatus(cbfn cb_self_connection_status_ftype, userData interface{}) {
	this.CallbackSelfConnectionStatusAdd(cbfn, userData)
}
func (this *LigTox) CallbackSelfConnectionStatusAdd(cbfn cb_self_connection_status_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_self_connection_statuss[cbfnp]; ok {
		return
	}
	this.cb_self_connection_statuss[cbfnp] = userData

}

// 包内部函数
func (this *LigTox) callbackFileRecvControl(friendNumber uint32, fileNumber uint32,
	control int) {

	for cbfni, ud := range this.cb_file_recv_controls {
		cbfn := *(*cb_file_recv_control_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(friendNumber), uint32(fileNumber), int(control), ud) })
	}
}

func (this *LigTox) CallbackFileRecvControl(cbfn cb_file_recv_control_ftype, userData interface{}) {
	this.CallbackFileRecvControlAdd(cbfn, userData)
}
func (this *LigTox) CallbackFileRecvControlAdd(cbfn cb_file_recv_control_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_recv_controls[cbfnp]; ok {
		return
	}
	this.cb_file_recv_controls[cbfnp] = userData

}

func (this *LigTox) callbackFileRecv(friendNumber uint32, fileNumber uint32, kind uint32,
	fileSize uint64, fileName string) {

	for cbfni, ud := range this.cb_file_recvs {
		cbfn := *(*cb_file_recv_ftype)(cbfni)
		fileName_ := fileName
		this.putcbevts(func() {
			cbfn(this, uint32(friendNumber), uint32(fileNumber), uint32(kind),
				uint64(fileSize), fileName_, ud)
		})
	}
}

func (this *LigTox) CallbackFileRecv(cbfn cb_file_recv_ftype, userData interface{}) {
	this.CallbackFileRecvAdd(cbfn, userData)
}
func (this *LigTox) CallbackFileRecvAdd(cbfn cb_file_recv_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_recvs[cbfnp]; ok {
		return
	}
	this.cb_file_recvs[cbfnp] = userData

}

func (this *LigTox) callbackFileRecvChunk(friendNumber uint32, fileNumber uint32,
	position uint64, data []byte) {

	for cbfni, ud := range this.cb_file_recv_chunks {
		cbfn := *(*cb_file_recv_chunk_ftype)(cbfni)
		data_ := data
		this.putcbevts(func() { cbfn(this, uint32(friendNumber), uint32(fileNumber), uint64(position), data_, ud) })
	}
}

func (this *LigTox) CallbackFileRecvChunk(cbfn cb_file_recv_chunk_ftype, userData interface{}) {
	this.CallbackFileRecvChunkAdd(cbfn, userData)
}
func (this *LigTox) CallbackFileRecvChunkAdd(cbfn cb_file_recv_chunk_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_recv_chunks[cbfnp]; ok {
		return
	}
	this.cb_file_recv_chunks[cbfnp] = userData

}

func (this *LigTox) callbackFileChunkRequest(friendNumber uint32, fileNumber uint32,
	position uint64, length int) {

	for cbfni, ud := range this.cb_file_chunk_requests {
		cbfn := *(*cb_file_chunk_request_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(friendNumber), uint32(fileNumber), uint64(position), int(length), ud) })
	}
}

func (this *LigTox) CallbackFileChunkRequest(cbfn cb_file_chunk_request_ftype, userData interface{}) {
	this.CallbackFileChunkRequestAdd(cbfn, userData)
}
func (this *LigTox) CallbackFileChunkRequestAdd(cbfn cb_file_chunk_request_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_chunk_requests[cbfnp]; ok {
		return
	}
	this.cb_file_chunk_requests[cbfnp] = userData

}

///// directly methods
func (this *LigTox) GetSavedataSize() int32 {
	return int32(0)
}

func (this *LigTox) GetSavedata() []byte {
	return nil
}

/*
 * @param pubkey hex 64B length
 */
func (this *LigTox) Bootstrap(addr string, port uint16, pubkey string) (bool, error) {
	return true, nil
}

func (this *LigTox) SelfGetAddress() string {
	return this.Binfo.GetId()
}

func (this *LigTox) SelfGetConnectionStatus() int {
	return int(this.Binfo.GetConnStatus())
}
func (this *LigTox) SelfSetName(name string) error {
	return nil
}

func (this *LigTox) SelfGetName() string {
	return this.Binfo.GetName()
}

func (this *LigTox) SelfGetNameSize() int {
	return len(this.Binfo.GetName())
}

func (this *LigTox) SelfSetStatusMessage(status string) (bool, error) {
	return true, nil
}

func (this *LigTox) SelfSetStatus(status uint8) {
}

func (this *LigTox) SelfGetStatusMessageSize() int {
	return len(this.Binfo.GetStmsg())
}

func (this *LigTox) SelfGetStatusMessage() (string, error) {
	return this.Binfo.GetStmsg(), nil
}

func (this *LigTox) SelfGetStatus() int {
	return int(this.Binfo.GetStatus())
}

func (this *LigTox) FriendAdd(friendId string, message string) (uint32, error) {
	return uint32(0), nil
}

func (this *LigTox) FriendAddNorequest(friendId string) (uint32, error) {
	return uint32(0), nil
}

func (this *LigTox) FriendByPublicKey(pubkey string) (uint32, error) {
	frnds := this.Binfo.GetFriends()
	for frndnum, frndo := range frnds {
		if frndo.Pubkey == pubkey {
			return frndnum, nil
		}
	}
	return uint32(0), nil
}

func (this *LigTox) FriendGetPublicKey(friendNumber uint32) (string, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return frndo.Pubkey, nil
	}
	return "", nil
}

func (this *LigTox) FriendDelete(friendNumber uint32) (bool, error) {
	return true, nil
}

func (this *LigTox) FriendGetConnectionStatus(friendNumber uint32) (int, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return int(frndo.ConnStatus), nil
	}
	return int(0), nil
}

func (this *LigTox) FriendExists(friendNumber uint32) bool {
	return false
}

func (this *LigTox) FriendSendMessage(friendNumber uint32, message string) (uint32, error) {
	fname := this.getMethodName()
	args := thspbs.Event{}
	args.Name = fname
	args.Args = []string{fmt.Sprintf("%d", friendNumber), message}

	// cli := thspbs.NewToxhsClient(this.rpcli)
	// rsp, err := cli.RmtCall(context.Background(), &args)
	rsp, err := this.rmtCall(&args)

	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	wn := gopp.MustUint32(rsp.Args[0])
	return wn, nil
}

func (this *LigTox) FriendSendAction(friendNumber uint32, action string) (uint32, error) {
	return uint32(0), nil
}

func (this *LigTox) FriendGetName(friendNumber uint32) (string, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return frndo.Name, nil
	}
	return string(""), nil
}

func (this *LigTox) FriendGetNameSize(friendNumber uint32) (int, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return len(frndo.Name), nil
	}
	return int(0), nil
}

func (this *LigTox) FriendGetStatusMessageSize(friendNumber uint32) (int, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return len(frndo.Stmsg), nil
	}
	return int(0), nil
}

func (this *LigTox) FriendGetStatusMessage(friendNumber uint32) (string, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return frndo.Stmsg, nil
	}
	return string(""), nil
}

func (this *LigTox) FriendGetStatus(friendNumber uint32) (int, error) {
	frnds := this.Binfo.GetFriends()
	if frndo, ok := frnds[friendNumber]; ok {
		return int(frndo.Status), nil
	}
	return int(0), nil
}

func (this *LigTox) FriendGetLastOnline(friendNumber uint32) (uint64, error) {
	return uint64(0), nil
}

func (this *LigTox) SelfSetTyping(friendNumber uint32, typing bool) (bool, error) {
	return true, nil
}

func (this *LigTox) FriendGetTyping(friendNumber uint32) (bool, error) {
	return true, nil
}

func (this *LigTox) SelfGetFriendListSize() uint32 {
	return uint32(len(this.Binfo.GetFriends()))
}

func (this *LigTox) SelfGetFriendList() []uint32 {
	fns := []uint32{}
	for _, fo := range this.Binfo.GetFriends() {
		fns = append(fns, fo.GetFnum())
	}
	return fns
}

// tox_callback_***

func (this *LigTox) SelfGetNospam() uint32 {
	return uint32(0)
}

func (this *LigTox) SelfSetNospam(nospam uint32) {
}

func (this *LigTox) SelfGetPublicKey() string {
	return this.Binfo.GetId()[:64]
}

func (this *LigTox) SelfGetSecretKey() string {
	return strings.ToUpper("")
}

// tox_lossy_***

func (this *LigTox) FriendSendLossyPacket(friendNumber uint32, data string) error {
	return nil
}

func (this *LigTox) FriendSendLosslessPacket(friendNumber uint32, data string) error {
	return nil
}

// tox_callback_avatar_**

func (this *LigTox) Hash(data string, datalen uint32) (string, bool, error) {
	return string(""), true, nil
}

// tox_callback_file_***
func (this *LigTox) FileControl(friendNumber uint32, fileNumber uint32, control int) (bool, error) {
	return true, nil
}

func (this *LigTox) FileSend(friendNumber uint32, kind uint32, fileSize uint64, fileId string, fileName string) (uint32, error) {
	return uint32(0), nil
}

func (this *LigTox) FileSendChunk(friendNumber uint32, fileNumber uint32, position uint64, data []byte) (bool, error) {
	return true, nil
}

func (this *LigTox) FileSeek(friendNumber uint32, fileNumber uint32, position uint64) (bool, error) {
	return true, nil
}

func (this *LigTox) FileGetFileId(friendNumber uint32, fileNumber uint32) (string, error) {
	return "", nil
}

// boostrap, see upper
func (this *LigTox) AddTcpRelay(addr string, port uint16, pubkey string) (bool, error) {
	return true, nil
}

func (this *LigTox) IsConnected() int {
	return int(0)
}

func (this *LigTox) getMethodName() string {
	pc, _, _, _ := runtime.Caller(1)
	fno := runtime.FuncForPC(pc)
	parts := strings.Split(fno.Name(), ".")
	fname := parts[2]
	return fname
}
