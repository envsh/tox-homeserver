package client

import (
	"errors"
	"fmt"
	"gopp"
	"log"
	"strings"
	"tox-homeserver/thspbs"
	"unsafe"
)

// conference callback type
type cb_conference_invite_ftype func(this *LigTox, friendNumber uint32, itype int, cookie string, userData interface{})
type cb_conference_message_ftype func(this *LigTox, groupNumber uint32, peerNumber uint32, mtype int, message string, userData interface{})

// type cb_conference_action_ftype func(this *LigTox, groupNumber uint32, peerNumber uint32, action string, userData interface{})
type cb_conference_title_ftype func(this *LigTox, groupNumber uint32, peerNumber uint32, title string, userData interface{})
type cb_conference_namelist_change_ftype func(this *LigTox, groupNumber uint32, peerNumber uint32, change int, userData interface{})

// tox_callback_conference_***

func (this *LigTox) callbackConferenceInvite(a0 uint32, a1 int, cookie string) {
	for cbfni, ud := range this.cb_conference_invites {
		cbfn := *(*cb_conference_invite_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, a0, a1, cookie, ud) })
	}
}

func (this *LigTox) CallbackConferenceInvite(cbfn cb_conference_invite_ftype, userData interface{}) {
	this.CallbackConferenceInviteAdd(cbfn, userData)
}
func (this *LigTox) CallbackConferenceInviteAdd(cbfn cb_conference_invite_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_invites[cbfnp]; ok {
		return
	}
	this.cb_conference_invites[cbfnp] = userData

}

func (this *LigTox) callbackConferenceMessage(a0 uint32, a1 uint32, mtype int, msg string) {
	for cbfni, ud := range this.cb_conference_messages {
		cbfn := *(*cb_conference_message_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, uint32(a0), uint32(a1), mtype, msg, ud) })
	}
}

func (this *LigTox) CallbackConferenceMessage(cbfn cb_conference_message_ftype, userData interface{}) {
	this.CallbackConferenceMessageAdd(cbfn, userData)
}
func (this *LigTox) CallbackConferenceMessageAdd(cbfn cb_conference_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_messages[cbfnp]; ok {
		return
	}
	this.cb_conference_messages[cbfnp] = userData
}

func (this *LigTox) callbackConferenceTitle(a0 uint32, a1 uint32, title string) {
	for cbfni, ud := range this.cb_conference_titles {
		cbfn := *(*cb_conference_title_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, a0, a1, title, ud) })
	}
}

func (this *LigTox) CallbackConferenceTitle(cbfn cb_conference_title_ftype, userData interface{}) {
	this.CallbackConferenceTitleAdd(cbfn, userData)
}
func (this *LigTox) CallbackConferenceTitleAdd(cbfn cb_conference_title_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_titles[cbfnp]; ok {
		return
	}
	this.cb_conference_titles[cbfnp] = userData
}

func (this *LigTox) callbackConferenceNameListChange(a0 uint32, a1 uint32, change int) {
	for cbfni, ud := range this.cb_conference_namelist_changes {
		cbfn := *(*cb_conference_namelist_change_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, a0, a1, change, ud) })
	}
}

func (this *LigTox) CallbackConferenceNameListChange(cbfn cb_conference_namelist_change_ftype, userData interface{}) {
	this.CallbackConferenceNameListChangeAdd(cbfn, userData)
}
func (this *LigTox) CallbackConferenceNameListChangeAdd(cbfn cb_conference_namelist_change_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_namelist_changes[cbfnp]; ok {
		return
	}
	this.cb_conference_namelist_changes[cbfnp] = userData
}

//////
func (this *LigTox) ConferenceSendMessage(groupNumber uint32, mtype int, msg string, userCode int64) error {
	fname := this.getMethodName()
	args := &thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%d", groupNumber), fmt.Sprintf("%d", mtype), msg}
	args.UserCode = userCode

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return errors.New(rsp.ErrMsg)
	}
	return err
}

func (this *LigTox) ConferenceJoin(friendNumber uint32, cookie string) (uint32, error) {
	fname := this.getMethodName()
	args := &thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%d", friendNumber), cookie}

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return 0, errors.New(rsp.ErrMsg)
	}
	return uint32(rsp.EventId), nil
}

func (this *LigTox) ConferenceNew(name string) (uint32, string, error) {
	fname := this.getMethodName()
	args := &thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%s", name)}

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return 0, "", errors.New(rsp.ErrMsg)
	}
	return uint32(rsp.EventId), rsp.Args[0], nil
}

func (this *LigTox) ConferenceDelete(groupNumber uint32) (uint32, error) {
	fname := this.getMethodName()
	args := &thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%d", groupNumber)}

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return 0, errors.New(rsp.ErrMsg)
	}
	return uint32(rsp.EventId), nil
}

func (this *LigTox) ConferencePeerCount(groupNumber uint32) (uint32, error) {
	fname := this.getMethodName()
	args := &thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%d", groupNumber)}

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return 0, errors.New(rsp.ErrMsg)
	}
	return uint32(rsp.EventId), nil
}

func (this *LigTox) ConferencePeerGetName(groupNumber uint32, peerNumber uint32) (string, error) {
	fname := this.getMethodName()
	args := &thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%d", groupNumber), fmt.Sprintf("%d", peerNumber)}

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return "", errors.New(rsp.ErrMsg)
	}
	return rsp.Args[0], nil
}

func (this *LigTox) ConferenceInvite(groupNumber uint32, peerNumber uint32) error {
	fname := this.getMethodName()
	args := thspbs.Event{}
	args.EventName = fname
	args.Args = []string{fmt.Sprintf("%d", groupNumber), fmt.Sprintf("%d", peerNumber)}

	// cli := thspbs.NewToxhsClient(this.rpcli)
	// rsp, err := cli.RmtCall(context.Background(), &args)
	rsp, err := this.rmtCall(&args)

	gopp.ErrPrint(err, rsp)
	log.Println(rsp)
	if rsp.ErrCode != 0 {
		return errors.New(rsp.ErrMsg)
	}
	return nil
}

// utils
/////
func ConferenceCookieToIdentifier(cookie string) string {
	if len(cookie) >= 6 {
		return cookie[6:]
	}
	return ""
}

func ConferenceIdIsEmpty(groupId string) bool {
	return groupId == "" || strings.Replace(groupId, "0", "", -1) == ""
}
