package main

import (
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"tox-homeserver/common"
	"tox-homeserver/thspbs"

	tox "github.com/kitech/go-toxcore"
	"github.com/kitech/go-toxcore/xtox"
)

type ToxVM struct {
	t *tox.Tox
}

var tvmCtx = xtox.NewToxContext("toxhs.tsbin", "toxhs0", "i'm toxhs0")

func newToxVM() *ToxVM {
	this := &ToxVM{}
	this.t = xtox.New(tvmCtx)
	log.Println(this.t == nil)
	xtox.SetAutoBotFeatures(this.t, xtox.FOTA_ADD_NET_HELP_BOTS|
		xtox.FOTA_ACCEPT_FRIEND_REQUEST|xtox.FOTA_ACCEPT_GROUP_INVITE)
	this.setupCallbacks()
	err := xtox.Connect(this.t)
	gopp.ErrPrint(err)
	return this
}

func (this *ToxVM) setupCallbacks() {
	t := this.t

	t.CallbackSelfConnectionStatusAdd(func(_ *tox.Tox, status int, userData interface{}) {
		if status == tox.CONNECTION_NONE {
		} else {
		}
		log.Println(status, tox.ConnStatusString(status))
		evt := thspbs.Event{}
		evt.Name = "SelfConnectionStatus"
		evt.Args = []string{fmt.Sprintf("%d", status)}

		this.pubmsg(&evt)
	}, nil)

	t.CallbackFriendRequestAdd(func(_ *tox.Tox, pubkey string, message string, userData interface{}) {
		evt := thspbs.Event{}
		evt.Name = "FriendRequest"
		evt.Args = []string{pubkey, message}

		this.pubmsg(&evt)
	}, nil)

	t.CallbackFriendMessageAdd(func(_ *tox.Tox, friendNumber uint32, message string, userData interface{}) {
		evt := thspbs.Event{}
		evt.Name = "FriendMessage"
		evt.Args = []string{fmt.Sprintf("%d", friendNumber), message}
		pubkey, err := t.FriendGetPublicKey(friendNumber)
		gopp.ErrPrint(err)
		fname, err := t.FriendGetName(friendNumber)
		gopp.ErrPrint(err)
		evt.Margs = []string{fname, pubkey}
		this.pubmsg(&evt)
	}, nil)
	/*
	   type cb_friend_request_ftype func(this *Tox, pubkey string, message string, userData interface{})
	   type cb_friend_message_ftype func(this *Tox, friendNumber uint32, message string, userData interface{})
	   type cb_friend_name_ftype func(this *Tox, friendNumber uint32, newName string, userData interface{})
	   type cb_friend_status_message_ftype func(this *Tox, friendNumber uint32, newStatus string, userData interface{})
	   type cb_friend_status_ftype func(this *Tox, friendNumber uint32, status int, userData interface{})
	   type cb_friend_connection_status_ftype func(this *Tox, friendNumber uint32, status int, userData interface{})
	   type cb_friend_typing_ftype func(this *Tox, friendNumber uint32, isTyping uint8, userData interface{})
	   type cb_friend_read_receipt_ftype func(this *Tox, friendNumber uint32, receipt uint32, userData interface{})
	   type cb_friend_lossy_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})
	   type cb_friend_lossless_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})

	   // self callback type
	   type cb_self_connection_status_ftype func(this *Tox, status int, userData interface{})

	   // file callback type
	   type cb_file_recv_control_ftype func(this *Tox, friendNumber uint32, fileNumber uint32,
	   	control int, userData interface{})
	   type cb_file_recv_ftype func(this *Tox, friendNumber uint32, fileNumber uint32, kind uint32, fileSize uint64,
	   	fileName string, userData interface{})
	   type cb_file_recv_chunk_ftype func(this *Tox, friendNumber uint32, fileNumber uint32, position uint64,
	   	data []byte, userData interface{})
	   type cb_file_chunk_request_ftype func(this *Tox, friend_number uint32, file_number uint32, position uint64,
	   	length int, user_data interface{})

	*/
}

func (this *ToxVM) pubmsg(evt *thspbs.Event) error {
	bcc, err := json.Marshal(evt)
	gopp.ErrPrint(err)
	err = appctx.rpcs.nc.Publish(common.CBEventBusName, bcc)
	gopp.ErrPrint(err)
	// TODO reconnect
	return err
}
