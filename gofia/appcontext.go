package gofia

import (
	"context"
	"gopp"
	"log"
	"sync"
	"time"

	"gomatcha.io/matcha/view"

	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kitech/godsts/maps/hashmap"
	"github.com/nats-io/nats"
	"google.golang.org/grpc"
)

type AppContext struct {
	nc    *nats.Conn
	rpcli *grpc.ClientConn
	vtcli *thscli.LigTox
	logFn func(s string)

	contacts  []*ContactItem
	contactsv []view.View
	mvst      *mainViewState

	mainV view.View
	currV view.View

	// friend: pubkey => *ChatFormView,
	// invited group: cookie => *ChatFormView
	// ours group: 无法直接获取自己创建的群组的cookie
	cfs *hashmap.Map // chat form views
	cts *hashmap.Map // contact item views
}

var appctx *AppContext
var appctxOnce sync.Once

func AppOnCreate() {
	appctxOnce.Do(func() {
		printBuildInfo(true)
		appctx = &AppContext{}
		appctx.vtcli = thscli.NewLigTox()
		appctx.cfs = hashmap.New()
		appctx.cts = hashmap.New()
		appctx.contacts = make([]*ContactItem, 0)
		appctx.contactsv = make([]view.View, 0)
		appctx.mvst = &mainViewState{}

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
					this.dispatchEvent(jso)
				}
			case <-stopC:
				return
			}
		}
	}
}

func (this *AppContext) dispatchEvent(jso *simplejson.Json) {
	evtName := jso.Get("name").MustString()
	switch evtName {
	case "SelfConnectionStatus":
	case "FriendRequest":
	case "FriendMessage":
		// jso.Get("args").GetIndex(0).MustString()
		msg := jso.Get("args").GetIndex(1).MustString()
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		cfx, found := this.cfs.Get(pubkey)
		if !found {
			log.Println("wtf, chat form view not found:", fname, pubkey)
		} else {
			cf := cfx.(*ChatFormView)
			msgo := &ContactMessage{}
			msgo.msg = msg
			msgo.tm = time.Now()
			cf.cfst.msgs.Add(msgo)
			if this.currV == cf {
				cf.Signal()
			}
		}
	default:
	}
}

///
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
