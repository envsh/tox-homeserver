package server

import (
	"context"
	"encoding/json"
	"gopp"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"

	// tox "github.com/envsh/go-toxcore"

	// "atapi/dorpc/dyngrpc"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"

	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"
)

type GrpcServer struct {
	srv   *grpc.Server
	lsner net.Listener
	svc   *GrpcService

	connsmu   sync.Mutex
	grpcConns map[thspbs.Toxhs_PollCallbackServer]uint64
	grpcUuids map[string]uint64 // uuid => uint64

	grpcConns2 map[string]struct {
		stm   thspbs.Toxhs_PollCallbackServer
		conno uint64
	}
}

func newGrpcServer() *GrpcServer {
	this := &GrpcServer{}
	this.grpcConns = make(map[thspbs.Toxhs_PollCallbackServer]uint64)
	this.grpcUuids = make(map[string]uint64)

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

	this.register()
	err = this.srv.Serve(this.lsner)
	gopp.ErrPrint(err)
}

func (this *GrpcServer) register() {
	// dyngrpc.RegisterService(demofn1, "thsdemo", "pasv")
}

// TODO
func (this *GrpcServer) Close() {

}

type GrpcService struct {
}

func (this *GrpcService) GetBaseInfo(ctx context.Context, req *thspbs.Event) (*thspbs.BaseInfo, error) {
	log.Println(appctx.tvm.t.SelfGetAddress(), req)
	log.Printf("ctx=%+v\n", ctx)
	out, err := packBaseInfo(appctx.tvm.t)
	gopp.ErrPrint(err)

	thscom.BytesRecved(len(req.String()))
	thscom.BytesSent(len(out.String()))
	return out, nil
}

// 自己的消息做多终端同步转发
func (this *GrpcService) RmtCall(ctx context.Context, req *thspbs.Event) (*thspbs.Event, error) {
	return RmtCallHandlers(ctx, req)
}

func (this *GrpcService) Ping(ctx context.Context, req *thspbs.EmptyReq) (*thspbs.EmptyReq, error) {
	out := &thspbs.EmptyReq{}
	thscom.BytesRecved(len(req.String()))
	thscom.BytesSent(len(out.String()))
	return out, nil
}

var grpcStreamConnNo uint64 = 0

func (this *GrpcService) PollCallback(req *thspbs.Event, stm thspbs.Toxhs_PollCallbackServer) error {
	nowt := time.Now()
	conno := atomic.AddUint64(&grpcStreamConnNo, 1)
	log.Println("New grpc stream poll connect.", len(appctx.rpcs.grpcConns), conno, req.DeviceUuid)
	appctx.rpcs.connsmu.Lock()
	if oldconno, ok := appctx.rpcs.grpcUuids[req.DeviceUuid]; ok {
		log.Println("already connected device:", req.DeviceUuid, oldconno)
		for stm, tconno := range appctx.rpcs.grpcConns {
			if tconno == oldconno {
				delete(appctx.rpcs.grpcConns, stm)
				// stm.Close() // ???
				break
			}
		}
		delete(appctx.rpcs.grpcUuids, req.DeviceUuid)
	}
	appctx.rpcs.grpcConns[stm] = conno
	appctx.rpcs.grpcUuids[req.DeviceUuid] = conno
	appctx.rpcs.connsmu.Unlock()
	defer func() {
		appctx.rpcs.connsmu.Lock()
		delete(appctx.rpcs.grpcConns, stm)
		if oldconno, ok := appctx.rpcs.grpcUuids[req.DeviceUuid]; ok && oldconno == conno {
			delete(appctx.rpcs.grpcUuids, req.DeviceUuid)
		}
		appctx.rpcs.connsmu.Unlock()
	}()

	select {
	case <-stm.Context().Done():
		break
	}
	log.Println("A stream done:", len(appctx.rpcs.grpcConns), conno, req.DeviceUuid, time.Since(nowt))
	return nil
}

func demofn1() {

}

// fill full args, so client can not crash
func fillmsgall(evto *thspbs.Event) {
	for len(evto.Args) < 10 {
		evto.Args = append(evto.Args, "")
	}
	for len(evto.Margs) < 10 {
		evto.Margs = append(evto.Margs, "")
	}
}

///
func pubmsgall(ctx context.Context, evt *thspbs.Event) error {
	fillmsgall(evt)

	var err error
	err = pubmsg2ws(ctx, evt)
	{
		err := pubmsg2grpc(ctx, evt)
		gopp.ErrPrint(err, ctx)
	}
	return err
}

func pubmsg2ws(ctx context.Context, evt *thspbs.Event) error {
	return appctx.wssrv.pushevt(evt)
}

func pubmsg2grpc(ctx context.Context, evt *thspbs.Event) error {
	var errtop error
	var stms = map[uint64]thspbs.Toxhs_PollCallbackServer{}
	var toconnid uint64
	appctx.rpcs.connsmu.Lock()
	for stm, connid := range appctx.rpcs.grpcConns {
		stms[connid] = stm
	}
	if evt.DeviceUuid != "" {
		toconnid, _ = appctx.rpcs.grpcUuids[evt.DeviceUuid]
	}
	appctx.rpcs.connsmu.Unlock()

	// specific the connection to push this event, or push to all connections
	for connid, stm := range stms {
		log.Println(connid, toconnid, evt.DeviceUuid, evt.Name)
		if evt.DeviceUuid == "" || (evt.DeviceUuid != "" && connid == toconnid) {
			err := stm.Send(evt)
			gopp.ErrPrint(err)
			if err != nil {
				errtop = err
			}
			jcc, _ := json.Marshal(evt)
			log.Println(err == nil, connid, toconnid, evt.DeviceUuid, evt.Name, len(jcc))
		}
	}
	return errtop
}

func pubmsg2grpcbyconn(ctx context.Context, evt *thspbs.Event, connid uint64) error {
	return nil
}
