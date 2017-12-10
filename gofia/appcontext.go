package gofia

import (
	"context"
	"gopp"
	"log"
	"sync"

	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/nats-io/nats"
	"google.golang.org/grpc"
)

type AppContext struct {
	nc    *nats.Conn
	rpcli *grpc.ClientConn
	vtcli *thscli.LigTox
	logFn func(s string)
}

var appctx *AppContext
var appctxOnce sync.Once

func AppOnCreate() {
	appctxOnce.Do(func() {
		printBuildInfo(true)
		appctx = &AppContext{}
		appctx.vtcli = thscli.NewLigTox()

		log.Println("connecting gnats:", thscom.GnatsAddr)
		nc, err := nats.Connect(thscom.GnatsAddr)
		gopp.ErrPrint(err)
		appctx.nc = nc

		log.Println("connecting grpc:", thscom.GrpcAddr)
		rpcli, err := grpc.Dial(thscom.GrpcAddr, grpc.WithInsecure())
		gopp.ErrPrint(err, rpcli)
		appctx.rpcli = rpcli

		go func() {
			appctx.getBaseInfo()
			go appctx.pollGrpc()
			go appctx.pollNats()
		}()
	})
}

func (this *AppContext) pollNats() {

	for {
		stopC := make(chan struct{}, 0)
		ch := make(chan *nats.Msg, 16)
		subh, err := this.nc.ChanSubscribe(thscom.CBEventBusName, ch)
		gopp.ErrPrint(err, subh)
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					log.Println("msg chan err, conn lost?")
				} else {
					log.Println(string(m.Data))
					jso, err := simplejson.NewJson(m.Data)
					gopp.ErrPrint(err, jso)
					if this.logFn != nil {
						this.logFn(string(m.Data))
					}
				}
			case <-stopC:
				return
			}
		}
	}
}

func (this *AppContext) pollGrpc() {

}

func (this *AppContext) getBaseInfo() {
	cc := this.rpcli
	thsc := thspbs.NewToxhsClient(cc)
	in := &thspbs.EmptyReq{}
	info, err := thsc.GetBaseInfo(context.Background(), in)
	gopp.ErrPrint(err, info)
	log.Println(info, len(info.Friends))

	this.vtcli.ParseBaseInfo(info)
	log.Println("herehehe")
}
