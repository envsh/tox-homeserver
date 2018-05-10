package server

import (
	"context"
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"tox-homeserver/common"
	"tox-homeserver/thspbs"

	tox "github.com/TokTok/go-toxcore-c"
	"github.com/envsh/go-toxcore/xtox"
)

//

func packBaseInfo(t *tox.Tox) (*thspbs.BaseInfo, error) {

	out := &thspbs.BaseInfo{}
	out.Id = t.SelfGetAddress()
	out.Name = t.SelfGetName()
	out.Stmsg, _ = t.SelfGetStatusMessage()
	out.Status = uint32(t.SelfGetStatus())
	out.ConnStatus = int32(t.SelfGetConnectionStatus())
	out.Friends = make(map[uint32]*thspbs.FriendInfo)
	out.Groups = make(map[uint32]*thspbs.GroupInfo)

	fns := t.SelfGetFriendList()
	for _, fn := range fns {
		pubkey, _ := t.FriendGetPublicKey(fn)
		fname, _ := t.FriendGetName(fn)
		stmsg, _ := t.FriendGetStatusMessage(fn)
		fstatus, _ := t.FriendGetConnectionStatus(fn)

		fi := &thspbs.FriendInfo{}
		fi.Pubkey = pubkey
		fi.Fnum = fn
		fi.Name = fname
		fi.Stmsg = stmsg
		fi.Status = uint32(fstatus)
		fi.ConnStatus = int32(fstatus)

		out.Friends[fn] = fi
	}

	gns := t.ConferenceGetChatlist()
	for _, gn := range gns {
		title, err := t.ConferenceGetTitle(gn)
		gopp.ErrPrint(err, title)

		mtype, err := t.ConferenceGetType(gn)
		gopp.ErrPrint(err, mtype)

		groupId, _ := xtox.ConferenceGetIdentifier(t, gn)

		gi := &thspbs.GroupInfo{}
		gi.Members = make(map[uint32]*thspbs.MemberInfo)
		gi.Gnum = gn
		gi.GroupId = groupId
		gi.Title = title
		gi.Ours = !xtox.IsInvitedGroup(t, gn)
		gi.Mtype = uint32(mtype)

		pcnt := t.ConferencePeerCount(gn)
		for i := uint32(0); i < pcnt; i++ {
			pname, err := t.ConferencePeerGetName(gn, i)
			gopp.ErrPrint(err, pname)
			if pname == "" {
				pname = common.DefaultUserName
			}
			ppubkey, err := t.ConferencePeerGetPublicKey(gn, i)
			gopp.ErrPrint(err, ppubkey)

			mi := &thspbs.MemberInfo{}
			mi.Pnum = i
			mi.Pubkey = ppubkey
			mi.Name = pname

			gi.Members[i] = mi
		}

		out.Groups[gn] = gi
	}

	return out, nil
}

// TODO 自己的消息做多终端同步转发
func RmtCallHandler(ctx context.Context, req *thspbs.Event) (*thspbs.Event, error) {
	log.Println(req.Id, req.Name, req.Args, req.Margs)
	out := &thspbs.Event{}

	var err error
	t := appctx.tvm.t
	switch req.Name {
	case "GetBaseInfo":
		binfo, err := packBaseInfo(t)
		gopp.ErrPrint(err)
		bdata, err := json.Marshal(binfo)
		gopp.ErrPrint(err)
		out.Args = []string{string(bdata)}
	case "FriendSendMessage": // "friendNumber", "msg"
		fnum := gopp.MustInt(req.Args[0])
		wn, err := t.FriendSendMessage(uint32(fnum), req.Args[1])
		gopp.ErrPrint(err)
		pubkey := t.SelfGetPublicKey()
		msgid, err := appctx.st.AddFriendMessage(req.Args[1], pubkey)
		gopp.ErrPrint(err, msgid)
		out.Mid = msgid
		out.Args = append(out.Args, fmt.Sprintf("%d", wn))

		// groups
	case "ConferenceNew": // args:name,returns:
		rname := req.Args[0]
		gn, err := t.ConferenceNew()
		gopp.ErrPrint(err, rname)
		out.Mid = int64(gn)
		_, err = t.ConferenceSetTitle(gn, rname)
		gopp.ErrPrint(err, gn, rname)
		groupId, _ := t.ConferenceGetIdentifier(gn)
		out.Args = append(out.Args, groupId)

		_, err = appctx.st.AddGroup(groupId, gn, rname)
		gopp.ErrPrint(err, gn, rname, groupId)

	case "ConferenceDelete": // "groupNumber"
		gnum := gopp.MustUint32(req.Args[0])
		_, err = t.ConferenceDelete(gnum)
		gopp.ErrPrint(err, req.Args)

	case "ConferenceSendMessage": // "groupNumber","mtype","msg"
		gnum := gopp.MustInt(req.Args[0])
		mtype := gopp.MustInt(req.Args[1])
		_, err = t.ConferenceSendMessage(uint32(gnum), mtype, req.Args[2])
		gopp.ErrPrint(err)
		if err != nil {
			out.Ecode = -1
			out.Emsg = err.Error()
		}
		cookie, _ := xtox.ConferenceGetCookie(t, uint32(gnum))
		pubkey := t.SelfGetPublicKey()
		msgid, err := appctx.st.AddGroupMessage(req.Args[2], "0", cookie, pubkey)
		gopp.ErrPrint(err, msgid)
		out.Mid = msgid
	case "ConferenceJoin": // friendNumber, cookie
		fnum := gopp.MustUint32(req.Args[0])
		cookie := req.Args[1]
		gn, err := t.ConferenceJoin(fnum, cookie)
		gopp.ErrPrint(err, fnum, len(cookie))
		out.Mid = int64(gn)

	case "ConferencePeerCount": // groupNumber
		gnum := gopp.MustUint32(req.Args[0])
		cnt := t.ConferencePeerCount(gnum)
		out.Mid = int64(cnt)
	case "ConferencePeerGetName": // groupNumber, peerNumber
		gnum := gopp.MustUint32(req.Args[0])
		pnum := gopp.MustUint32(req.Args[1])
		pname, err := t.ConferencePeerGetName(gnum, pnum)
		gopp.ErrPrint(err, req.Args)
		out.Args = append(out.Args, pname)
	case "ConferenceInvite": // groupNumber, friendNumber
		gnum := gopp.MustUint32(req.Args[0])
		fnum := gopp.MustUint32(req.Args[1])
		_, err := t.ConferenceInvite(fnum, gnum)
		gopp.ErrPrint(err, req.Args)
		// TODO
		// case "GetHistory":
	default:
		log.Println("unimpled:", req.Name, req.Args)
		out.Ecode = -1
		out.Emsg = fmt.Sprintf("Unimpled: %s", req.Name)
	}

	common.BytesRecved(len(req.String()))
	common.BytesSent(len(out.String()))
	return out, nil
}
