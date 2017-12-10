package client

import (
	"tox-homeserver/thspbs"
	"unsafe"
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
	ToxId  string
	Status int
	Stmsg  string
	Binfo  *thspbs.BaseInfo

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
	this.initCbmap()

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

func (this *LigTox) ParseBaseInfo(bi *thspbs.BaseInfo) {
	this.Binfo = bi
	this.callbackBaseInfo(bi)
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
