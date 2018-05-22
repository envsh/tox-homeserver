package client

import (
	"context"
	"gopp"
	"log"
	"net/http"
	"sync"
	"time"
	thscom "tox-homeserver/common"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/go-xorm/xorm"
	"github.com/kitech/godsts/maps/hashmap"
	"github.com/nats-io/nats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type AppContext struct {
	nc    *nats.Conn
	rpcli *grpc.ClientConn
	vtcli *LigTox
	logFn func(s string)

	store *Storage

	chatFormStates *hashmap.Map // id => chat form state datas
	contactStates  *hashmap.Map // id => contact state datas

	// logState *LogState
}

var appctx *AppContext
var appctxOnce sync.Once

func GetAppCtx() *AppContext { return appctx }

func AppOnCreate() {
	appctxOnce.Do(func() {
		// printBuildInfo(true)
		log.Println("Start pprof server: *:8089")
		go func() { gopp.ErrPrint(http.ListenAndServe(":8089", nil)) }()

		// 初始化顺序: server => memory => disk => network
		appctx = &AppContext{}
		// appctx.logState = newLogState()
		appctx.vtcli = NewLigTox()
		appctx.contactStates = hashmap.New()
		appctx.chatFormStates = hashmap.New()

		appctx.store = store.NewStorage()
		if appctx.store.DeviceEmpty() {
			err := appctx.store.AddDevice()
			gopp.ErrPrint(err)
		}
		dv := appctx.store.GetDevice()
		if dv != nil {
			log.Println("my device:", dv.Uuid)
		} else {
			log.Println("my device not exist: wtf")
		}

		log.Println("connecting gnats:", thscom.GnatsAddr)
		nc, err := nats.Connect(thscom.GnatsAddr)
		gopp.ErrPrint(err)
		appctx.nc = nc

		log.Println("connecting grpc:", thscom.GrpcAddr)
		rpcli, err := grpc.Dial(thscom.GrpcAddr, grpc.WithInsecure())
		gopp.ErrPrint(err, rpcli)
		appctx.rpcli = rpcli
		go appctx.keepConn()
		time.Sleep(1 * time.Millisecond)

		//ping
		cc := appctx.rpcli
		thsc := thspbs.NewToxhsClient(cc)
		in := &thspbs.EmptyReq{}
		_, err = thsc.Ping(context.Background(), in)
		gopp.ErrPrint(err)

		go func() {

			appctx.getBaseInfo()
			go appctx.pollGrpc()
			go appctx.pollNats()
		}()
	})
}

func (this *AppContext) GetLigTox() *LigTox         { return this.vtcli }
func (this *AppContext) GetStorage() *store.Storage { return this.store }

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
		///
		pubkey := jso.Get("args").GetIndex(0).MustString()
		_, err := appctx.store.AddFriend(pubkey, 0, "", "")
		gopp.ErrPrint(err, jso.Get("args"))

	case "FriendMessage":
		// jso.Get("args").GetIndex(0).MustString()
		msg := jso.Get("args").GetIndex(1).MustString()
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()

		cfsx, found := this.chatFormStates.Get(pubkey)
		_ = cfsx
		if !found {
			log.Println("wtf, chat form view not found:", fname, pubkey)
		} else {
			cfs := cfsx.(*ChatFormState)
			_ = cfs
			// msgo := &ContactMessage{}
			// msgo.msg = msg
			// msgo.tm = time.Now()
			// cfs.msgs.Add(msgo)
			// if this.currV != nil && this.currV.(*ChatFormView).cfst == cfs {
			//	this.currV.(*ChatFormView).Signal()
			// }
			// if appctx.app.Child != nil && appctx.app.Child.(*ChatFormView).cfst == cfs {
			//appctx.app.Child.(*ChatFormView).Signal()
			//}
			// InterBackRelay.Signal()
		}

		///
		eventId := int64(gopp.MustInt(jso.Get("margs").GetIndex(2).MustString()))
		_, err := appctx.store.AddFriendMessage(msg, pubkey, eventId)
		gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		fname := jso.Get("margs").GetIndex(0).MustString()
		pubkey := jso.Get("margs").GetIndex(1).MustString()
		// cfx, found := this.cfvs.Get(pubkey)
		cfsx, found := this.chatFormStates.Get(pubkey)
		_ = cfsx
		if !found {
			log.Println("wtf, chat form view not found:", fname, pubkey)
		} else {
			cfs := cfsx.(*ChatFormState)
			cfs.status = uint32(gopp.MustInt(jso.Get("args").GetIndex(1).MustString()))
			// this.signalProperView(cfs, true)
		}

	case "ConferenceInvite":
		groupNumber := jso.Get("margs").GetIndex(2).MustString()
		_ = groupNumber
		cookie := jso.Get("args").GetIndex(2).MustString()
		groupId := ConferenceCookieToIdentifier(cookie)
		log.Println(groupId)
		valuex, found := appctx.contactStates.Get(groupId)
		_, _ = valuex, found
		var ctis *ContactItemState
		if !found {
			ctis = newContactItemState()
			appctx.contactStates.Put(groupId, ctis)
			log.Println("new group contact:", groupId)
		} else {
			ctis = valuex.(*ContactItemState)
		}
		ctis.group = true
		ctis.cnum = uint32(gopp.MustInt(groupNumber))
		ctis.ctid = groupId

		// if appctx.app != nil && appctx.app.Child == nil {
		//	InterBackRelay.Signal()
		//}

		///
		_, err := appctx.store.AddGroup(groupId, uint32(gopp.MustInt(groupNumber)), ctis.ctname)
		gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupTitle := jso.Get("args").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(0).MustString()
		log.Println(groupId, groupTitle)
		if ConferenceIdIsEmpty(groupId) {
			break
		}
		valuex, found := appctx.contactStates.Get(groupId)
		_, _ = valuex, found
		var ctis *ContactItemState
		if !found {
			ctis = newContactItemState()
			ctis.group = true
			appctx.contactStates.Put(groupId, ctis)
			log.Println("new group contact:", groupId)
		} else {
			ctis = valuex.(*ContactItemState)
		}
		if groupTitle != "" && groupTitle != ctis.ctname {
			ctis.ctname = groupTitle
			//	if appctx.app != nil && appctx.app.Child == nil {
			//		InterBackRelay.Signal()
			//	}
		}
	case "ConferenceNameListChange":
		groupTitle := jso.Get("margs").GetIndex(2).MustString()
		groupId := jso.Get("margs").GetIndex(3).MustString()
		log.Println(groupId, groupTitle)
		if ConferenceIdIsEmpty(groupId) {
			log.Println("empty")
			break
		}
		valuex, found := appctx.contactStates.Get(groupId)
		_, _ = valuex, found
		var ctis *ContactItemState
		if !found {
			ctis = newContactItemState()
			ctis.group = true
			appctx.contactStates.Put(groupId, ctis)
			log.Println("new group contact:", groupId)
		} else {
			ctis = valuex.(*ContactItemState)
		}
		if groupTitle != "" && groupTitle != ctis.ctname {
			ctis.ctname = groupTitle
			//	if appctx.app != nil && appctx.app.Child == nil {
			//		InterBackRelay.Signal()
			//	}
		}

		///
		peerPubkey := jso.Get("margs").GetIndex(1).MustString()
		_, err := appctx.store.AddPeer(peerPubkey, 0, "")
		gopp.ErrPrint(err)

	case "ConferenceMessage":
		groupId := jso.Get("margs").GetIndex(3).MustString()
		log.Println(groupId)
		if ConferenceIdIsEmpty(groupId) {
			break
		}
		valuex, found := appctx.contactStates.Get(groupId)
		_, _ = valuex, found
		var ctis *ContactItemState
		if !found {
			log.Println("group contact not found:", groupId)
		} else {
			ctis = valuex.(*ContactItemState)
		}

		message := jso.Get("args").GetIndex(3).MustString()
		if ctis != nil {
			//	msgo := &ContactMessage{}
			//	msgo.msg = message
			//	msgo.tm = time.Now()
			//	ctis.msgs.Add(msgo)

			// this.signalProperView(ctis, false)
			/*
				if appctx.app.Child != nil && appctx.app.Child.(*ChatFormView).cfst == ctis {
					appctx.app.Child.(*ChatFormView).Signal()
				}
			*/
		}

		//
		peerName := jso.Get("margs").GetIndex(0).MustString()
		peerPubkey := jso.Get("margs").GetIndex(1).MustString()
		_, err := appctx.store.GetContactByPubkey(peerPubkey)
		if err == xorm.ErrNotExist {
			peerNum := gopp.MustUint32(jso.Get("args").GetIndex(1).MustString())
			appctx.store.AddPeer(peerPubkey, peerNum, peerName)
		}
		eventId := int64(gopp.MustInt(jso.Get("margs").GetIndex(4).MustString()))
		_, err = appctx.store.AddGroupMessage(message, "0", groupId, peerPubkey, eventId)
		gopp.ErrPrint(err)

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
	gopp.ErrPrint(err, info) // rpc error: code = Internal desc = grpc: failed to unmarshal the received message proto: bad wiretype for field thspbs.GroupInfo.Ours: got wiretype 2, want 0 <nil>
	log.Println(info, len(info.Friends))

	this.vtcli.ParseBaseInfo(info)
	log.Println("herehehe")
}

func (this *AppContext) persistBaseInfo(bi *thspbs.BaseInfo) {
	for _, frndo := range bi.Friends {
		appctx.store.AddFriend(frndo.Pubkey, frndo.Fnum, frndo.Name, frndo.Stmsg)
	}
	for _, grpo := range bi.Groups {
		for _, peero := range grpo.GetMembers() {
			appctx.store.AddPeer(peero.Pubkey, peero.Pnum, peero.Name)
		}
	}
	for _, grpo := range bi.Groups {
		appctx.store.AddGroup(grpo.GroupId, grpo.Gnum, grpo.Title)
	}
}

func (this *AppContext) doCall() {
	cc := this.rpcli
	thsc := thspbs.NewToxhsClient(cc)
	_ = thsc
}

// should block
func (this *AppContext) keepConn() {
	for false {
		ok := this.rpcli.WaitForStateChange(context.Background(), connectivity.Idle)
		log.Println(ok, this.rpcli.GetState().String())
	}
}
