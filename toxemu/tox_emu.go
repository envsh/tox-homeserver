package main

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static void array_item_set(uint32_t *lst, int n, uint32_t v) { lst[n] = v; }

static void on_friend_message(void* cbfn, uint32_t friendNumber, int mtype, uint8_t* message, size_t length) {
    ((void (*)(void*, uint32_t, int, uint8_t*, size_t, void*))(cbfn))
    (NULL, friendNumber, mtype, message, length, NULL);
}

*/
import "C"
import (
	"encoding/hex"
	"gopp"
	"io/ioutil"
	"log"
	"unsafe"

	thscli "tox-homeserver/client"
)

//export tox_version_major
func tox_version_major() C.uint32_t {
	return 0
}

type ToxOption struct {
}

//export tox_options_new
func tox_options_new() unsafe.Pointer {
	opt := &ToxOption{}
	return unsafe.Pointer(opt)
}

type Tox struct{}
type ToxCallbacks struct {
	cbfn_self_connection_status unsafe.Pointer
	cbfn_friend_message         unsafe.Pointer
}

var lt *thscli.LigTox

func setupCallbacks() {
	lt.CallbackSelfConnectionStatusAdd(func(this *thscli.LigTox, status int, userData interface{}) {
		log.Println(status)
	}, nil)

	lt.CallbackFriendConnectionStatusAdd(func(this *thscli.LigTox, friendNumber uint32, status int, userData interface{}) {
		log.Println(friendNumber, status)
	}, nil)

	lt.CallbackFriendMessageAdd(func(this *thscli.LigTox, friendNumber uint32, message string, userData interface{}) {
		log.Println(friendNumber, message)
		if dmtcb.cbfn_friend_message != nil {
			log.Println("Invoking cbfn...", friendNumber, message)
			msgp := (*C.uint8_t)((unsafe.Pointer)(&[]byte(message)[0]))
			C.on_friend_message(dmtcb.cbfn_friend_message, C.uint32_t(friendNumber), C.int(0), msgp, C.size_t(len(message)))
		}
	}, nil)
}

var dmt *Tox
var dmtcb = &ToxCallbacks{}

//export tox_new
func tox_new() unsafe.Pointer {
	t := &Tox{}
	dmt = t
	lt = thscli.NewLigTox()
	setupCallbacks()
	lt.GetBaseInfo()
	return unsafe.Pointer(t)
}

//export tox_kill
func tox_kill(t unsafe.Pointer) {
	log.Println(t)
}

//export tox_self_get_name_size
func tox_self_get_name_size(unsafe.Pointer) C.size_t {
	return C.size_t(lt.SelfGetNameSize())
}

//export tox_self_get_name
func tox_self_get_name(t unsafe.Pointer, name *C.char) {
	s := lt.SelfGetName()
	log.Println(s)
	n := []byte(s)
	C.strcpy(name, (*C.char)((unsafe.Pointer)(&n[0])))
}

//export tox_self_get_status_message_size
func tox_self_get_status_message_size(unsafe.Pointer) C.size_t {
	return C.size_t(lt.SelfGetStatusMessageSize())
}

//export tox_self_get_status_message
func tox_self_get_status_message(t unsafe.Pointer, stmsg *C.char) {
	s, _ := lt.SelfGetStatusMessage()
	log.Println(s)
	n := []byte(s)
	C.strcpy(stmsg, (*C.char)((unsafe.Pointer)(&n[0])))
}

//export tox_self_get_address
func tox_self_get_address(t unsafe.Pointer, addr *C.uint8_t) {
	s := lt.SelfGetAddress()
	log.Println(s)
	n, err := hex.DecodeString(s)
	gopp.ErrPrint(err)
	C.memcpy(unsafe.Pointer(addr), (unsafe.Pointer)(&n[0]), C.size_t(len(s)/2))
}

//export tox_self_get_friend_list_size
func tox_self_get_friend_list_size(unsafe.Pointer) C.size_t {
	return C.size_t(lt.SelfGetFriendListSize())
}

//export tox_self_get_friend_list
func tox_self_get_friend_list(t unsafe.Pointer, lst *C.uint32_t) {
	fns := lt.SelfGetFriendList()
	for i, fn := range fns {
		// lst[i] = fn
		C.array_item_set(lst, C.int(i), C.uint32_t(fn))
	}
}

//export tox_self_get_public_key
func tox_self_get_public_key(t unsafe.Pointer, pk *C.uint8_t) {
	s := lt.SelfGetPublicKey()
	log.Println(s)
	b, err := hex.DecodeString(s)
	gopp.ErrPrint(err)
	C.memcpy(unsafe.Pointer(pk), unsafe.Pointer(&b[0]), 32)
}

//export tox_iterate
func tox_iterate(unsafe.Pointer) {
}

//export tox_iteration_interval
func tox_iteration_interval(unsafe.Pointer) {
}

//export tox_self_get_connection_status
func tox_self_get_connection_status(t unsafe.Pointer) C.int {
	return C.int(lt.SelfGetConnectionStatus())
}

//export tox_self_set_status_message
func tox_self_set_status_message(unsafe.Pointer) {}

//export tox_self_set_status
func tox_self_set_status(unsafe.Pointer) {}

//export tox_friend_get_public_key
func tox_friend_get_public_key(t unsafe.Pointer, friend_number C.uint32_t, public_key *C.uint8_t, cerr unsafe.Pointer) {
	pubkey, err := lt.FriendGetPublicKey(uint32(friend_number))
	gopp.ErrPrint(err)
	log.Println(friend_number, pubkey)

	b, err := hex.DecodeString(pubkey)
	gopp.ErrPrint(err)
	C.memcpy(unsafe.Pointer(public_key), unsafe.Pointer(&b[0]), 32)
}

//export tox_friend_get_name_size
func tox_friend_get_name_size(t unsafe.Pointer, friend_number C.uint32_t) C.size_t {
	name, err := lt.FriendGetName(uint32(friend_number))
	gopp.ErrPrint(err)
	log.Println(name)
	return C.size_t(len(name))
}

//export tox_friend_get_name
func tox_friend_get_name(t unsafe.Pointer, friend_number C.uint32_t, name *C.uint8_t) {
	name_, err := lt.FriendGetName(uint32(friend_number))
	gopp.ErrPrint(err)
	log.Println(friend_number, name_)

	if len(name_) > 0 {
		n := []byte(name_)
		C.memcpy(unsafe.Pointer(name), unsafe.Pointer(&n[0]), C.size_t(len(name_)))
	}
}

//export tox_friend_get_status_message_size
func tox_friend_get_status_message_size(t unsafe.Pointer, friend_number C.uint32_t) C.size_t {
	sz, err := lt.FriendGetStatusMessageSize(uint32(friend_number))
	gopp.ErrPrint(err)
	return C.size_t(sz)
}

//export tox_friend_get_status_message
func tox_friend_get_status_message(t unsafe.Pointer, friend_number C.uint32_t, stmsg *C.uint8_t) {
	stmsg_, err := lt.FriendGetStatusMessage(uint32(friend_number))
	gopp.ErrPrint(err)

	if len(stmsg_) > 0 {
		n := []byte(stmsg_)
		C.memcpy(unsafe.Pointer(stmsg), unsafe.Pointer(&n[0]), C.size_t(len(stmsg_)))
	}
}

//export tox_group_message_send
func tox_group_message_send(unsafe.Pointer) {}

//export tox_get_savedata_size
func tox_get_savedata_size(unsafe.Pointer) C.size_t {
	bcc, err := ioutil.ReadFile("./tox_save.tox")
	gopp.ErrPrint(err)
	return C.size_t(len(bcc))
}

//export tox_get_savedata
func tox_get_savedata(t unsafe.Pointer, d *C.uint8_t) {
	bcc, err := ioutil.ReadFile("./tox_save.tox")
	gopp.ErrPrint(err)
	C.memcpy(unsafe.Pointer(d), unsafe.Pointer(&bcc[0]), C.size_t(len(bcc)))
}

//export tox_callback_self_connection_status
func tox_callback_self_connection_status(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)

}

//export tox_callback_friend_status
func tox_callback_friend_status(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_message
func tox_callback_friend_message(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
	dmtcb.cbfn_friend_message = cfn
}

//export tox_callback_friend_connection_status
func tox_callback_friend_connection_status(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_name
func tox_callback_friend_name(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_status_message
func tox_callback_friend_status_message(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_request
func tox_callback_friend_request(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_typing
func tox_callback_friend_typing(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_read_receipt
func tox_callback_friend_read_receipt(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_file_recv
func tox_callback_file_recv(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_file_recv_chunk
func tox_callback_file_recv_chunk(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_file_chunk_request
func tox_callback_file_chunk_request(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_file_recv_control
func tox_callback_file_recv_control(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_lossless_packet
func tox_callback_friend_lossless_packet(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_friend_lossy_packet
func tox_callback_friend_lossy_packet(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_group_invite
func tox_callback_group_invite(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_group_message
func tox_callback_group_message(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_group_action
func tox_callback_group_action(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_group_title
func tox_callback_group_title(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export tox_callback_group_namelist_change
func tox_callback_group_namelist_change(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

///////////////////////
type ToxAV struct{}

var dmtav *ToxAV

//export toxav_new
func toxav_new(unsafe.Pointer) unsafe.Pointer {
	v := &ToxAV{}
	dmtav = v
	return unsafe.Pointer(v)
}

//export toxav_kill
func toxav_kill(unsafe.Pointer) {

}

//export toxav_iterate
func toxav_iterate(unsafe.Pointer) {

}

//export toxav_iteration_interval
func toxav_iteration_interval(unsafe.Pointer) C.int {
	return 0
}

//export toxav_callback_call_state
func toxav_callback_call_state(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {}

//export toxav_callback_call
func toxav_callback_call(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export toxav_callback_audio_receive_frame
func toxav_callback_audio_receive_frame(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}

//export toxav_callback_video_receive_frame
func toxav_callback_video_receive_frame(t unsafe.Pointer, cfn unsafe.Pointer, ud unsafe.Pointer) {
	log.Println(t, cfn)
}
