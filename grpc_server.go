package main

import (
	"context"
	"gopp"
	"log"
	"net"

	"github.com/kitech/go-toxcore/xtox"
	"github.com/nats-io/nats"

	"atapi/dorpc/dyngrpc"

	"google.golang.org/grpc"

	"tox-homeserver/common"
	"tox-homeserver/thspbs"
)

type GrpcServer struct {
	srv   *grpc.Server
	lsner net.Listener
	nc    *nats.Conn
	svc   *GrpcService
}

func newGrpcServer() *GrpcServer {
	this := &GrpcServer{}

	// TODO 压缩支持
	// TODO tls支持
	nc, _ := nats.Connect(common.GnatsAddr)
	this.nc = nc
	this.srv = grpc.NewServer()

	this.svc = &GrpcService{}
	thspbs.RegisterToxhsServer(this.srv, this.svc)

	return this
}

func (this *GrpcServer) run() {
	lsner, err := net.Listen("tcp", ":2080")
	gopp.ErrPrint(err)
	this.lsner = lsner
	log.Println("listen on:", lsner.Addr())

	this.register()
	err = this.srv.Serve(this.lsner)
	gopp.ErrPrint(err)
}

func (this *GrpcServer) register() {
	dyngrpc.RegisterService(demofn1, "thsdemo", "pasv")
}

type GrpcService struct {
}

func (this *GrpcService) GetBaseInfo(ctx context.Context, req *thspbs.EmptyReq) (*thspbs.BaseInfo, error) {
	log.Println(req, appctx.tvm.t.SelfGetAddress())
	out := &thspbs.BaseInfo{}
	t := appctx.tvm.t
	out.Id = t.SelfGetAddress()
	out.Name = t.SelfGetName()
	out.Stmsg, _ = t.SelfGetStatusMessage()
	out.Status = uint32(t.SelfGetConnectionStatus())
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

		out.Friends[fn] = fi
	}

	gns := t.ConferenceGetChatlist()
	for _, gn := range gns {
		title, err := t.ConferenceGetTitle(gn)
		gopp.ErrPrint(err, title)

		mtype, err := t.ConferenceGetType(gn)
		gopp.ErrPrint(err, mtype)

		gi := &thspbs.GroupInfo{}
		gi.Members = make(map[uint32]*thspbs.MemberInfo)
		gi.Gnum = gn
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

func (this *GrpcService) RmtCall(ctx context.Context, req *thspbs.Event) (*thspbs.Event, error) {
	log.Println(req.Id, req.Name, req.Args, req.Margs)
	out := thspbs.Event{}

	t := appctx.tvm.t
	switch req.Name {
	case "FriendSendMessage":
		fnum := gopp.MustInt(req.Args[0])
		_, err := t.FriendSendMessage(uint32(fnum), req.Args[1])
		gopp.ErrPrint(err)
	}
	return &out, nil
}

func (this *GrpcService) PollCallback(req *thspbs.EmptyReq, stm thspbs.Toxhs_PollCallbackServer) error {
	return nil
}

func demofn1() {

}

///
