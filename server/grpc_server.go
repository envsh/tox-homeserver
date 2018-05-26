package server

import (
	"context"
	"encoding/json"
	"gopp"
	"log"
	"net"
	"sync"

	// tox "github.com/envsh/go-toxcore"

	"github.com/nats-io/go-nats"

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

	connsmu   sync.Mutex
	grpcConns map[thspbs.Toxhs_PollCallbackServer]int
}

func newGrpcServer() *GrpcServer {
	this := &GrpcServer{}
	this.grpcConns = make(map[thspbs.Toxhs_PollCallbackServer]int)

	// TODO 压缩支持
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

	// TODO tls支持
	log.Println("Connecting gnatsd:", common.GnatsAddrlo)
	nc, err := nats.Connect(common.GnatsAddrlo)
	gopp.ErrPrint(err)
	this.nc = nc

	this.register()
	err = this.srv.Serve(this.lsner)
	gopp.ErrPrint(err)
}

func (this *GrpcServer) register() {
	dyngrpc.RegisterService(demofn1, "thsdemo", "pasv")
}

func (this *GrpcServer) checkOrReconnNats(err error) {
	if err == nats.ErrConnectionClosed {
		log.Println("Reconnecting...")
		nc, err2 := nats.Connect(common.GnatsAddr)
		gopp.ErrPrint(err2)
		if err2 == nil {
			this.nc = nc
		}
	}
}

// TODO
func (this *GrpcServer) Close() {

}

type GrpcService struct {
}

func (this *GrpcService) GetBaseInfo(ctx context.Context, req *thspbs.EmptyReq) (*thspbs.BaseInfo, error) {
	log.Println(req, appctx.tvm.t.SelfGetAddress())
	out, err := packBaseInfo(appctx.tvm.t)
	gopp.ErrPrint(err)

	common.BytesRecved(len(req.String()))
	common.BytesSent(len(out.String()))
	return out, nil
}

// 自己的消息做多终端同步转发
func (this *GrpcService) RmtCall(ctx context.Context, req *thspbs.Event) (*thspbs.Event, error) {
	return RmtCallHandlers(ctx, req)
}

func (this *GrpcService) Ping(ctx context.Context, req *thspbs.EmptyReq) (*thspbs.EmptyReq, error) {
	out := &thspbs.EmptyReq{}
	common.BytesRecved(len(req.String()))
	common.BytesSent(len(out.String()))
	return out, nil
}

func (this *GrpcService) PollCallback(req *thspbs.EmptyReq, stm thspbs.Toxhs_PollCallbackServer) error {
	log.Println("New grpc stream poll connect.", len(appctx.rpcs.grpcConns))
	appctx.rpcs.connsmu.Lock()
	appctx.rpcs.grpcConns[stm] = 1
	appctx.rpcs.connsmu.Unlock()
	defer func() {
		appctx.rpcs.connsmu.Lock()
		delete(appctx.rpcs.grpcConns, stm)
		appctx.rpcs.connsmu.Unlock()
	}()

	for {
		select {}
	}
	// return nil
}

func demofn1() {

}

///
func pubmsgall(ctx context.Context, evt *thspbs.Event) error {
	var err error
	err = pubmsg2nats(ctx, evt)
	if err == nil {
		err = pubmsg2ws(ctx, evt)
	}
	{
		err := pubmsg2grpc(ctx, evt)
		gopp.ErrPrint(err, ctx)
	}
	return err
}

func pubmsg2nats(ctx context.Context, evt *thspbs.Event) error {
	bcc, err := json.Marshal(evt)
	gopp.ErrPrint(err)
	err = appctx.rpcs.nc.Publish(common.CBEventBusName, bcc)
	gopp.ErrPrint(err)
	// reconnect
	if err != nil {
		appctx.rpcs.checkOrReconnNats(err)
		err = appctx.rpcs.nc.Publish(common.CBEventBusName, bcc)
		gopp.ErrPrint(err)
	}
	if err == nil {
		// log.Println("pubmsg ok", len(bcc))
	}
	common.BytesSent(len(bcc))
	return err
}

func pubmsg2ws(ctx context.Context, evt *thspbs.Event) error {
	return appctx.wssrv.pushevt(evt)
}

func pubmsg2grpc(ctx context.Context, evt *thspbs.Event) error {
	var errtop error
	var stms []thspbs.Toxhs_PollCallbackServer
	appctx.rpcs.connsmu.Lock()
	for stm, _ := range appctx.rpcs.grpcConns {
		stms = append(stms, stm)
	}
	appctx.rpcs.connsmu.Unlock()

	for _, stm := range stms {
		err := stm.Send(evt)
		gopp.ErrPrint(err)
		if err != nil {
			errtop = err
		}
	}
	return errtop
}
