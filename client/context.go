package client

import (
	"gopp"
	"log"
	"net/http"
	"sync"
	"time"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"

	"github.com/go-xorm/xorm"
)

type AppContext struct {
	srvtp Transport
	vtcli *LigTox
	logFn func(s string)

	store *Storage

	// logState *LogState
}

var appctx *AppContext
var appctxOnce sync.Once

func NewAppContext() *AppContext {
	appctx = &AppContext{}

	_AppOnCreate()
	return appctx
}

func (this *AppContext) OpenStrorage() {
	appctx.store = store.NewStorage()
	if appctx.store.DeviceEmpty() {
		err := appctx.store.AddDevice()
		gopp.ErrPrint(err)
	}
}

// func GetAppCtx() *AppContext { return appctx }

func AppConnect(srvurl string) error {
	// 初始化顺序: server => memory => disk => network
	// appctx.logState = newLogState()
	var err error
	vtcli := NewLigTox(srvurl)
	err = vtcli.Connect()
	gopp.ErrPrint(err, srvurl)
	if err != nil {
		return err
	}
	vtcli.start()
	appctx.vtcli = vtcli

	dv := appctx.store.GetDevice()
	if dv != nil {
		log.Println("my device:", dv.Uuid)
	} else {
		log.Println("my device not exist: wtf")
	}

	appctx.srvtp = appctx.vtcli.srvtp
	appctx.srvtp.OnData(func(evto *thspbs.Event, data []byte) { appctx.dispatchEvent(evto) })

	time.Sleep(1 * time.Millisecond)
	go appctx.getAndPersistBaseInfo()

	return nil
}

func _AppOnCreate() {
	appctxOnce.Do(func() {
		// printBuildInfo(true)
		log.Println("Start pprof server: *:8089")
		go func() { gopp.ErrPrint(http.ListenAndServe(":8089", nil)) }()
	})
}

func (this *AppContext) GetLigTox() *LigTox         { return this.vtcli }
func (this *AppContext) GetStorage() *store.Storage { return this.store }

// 只用做消息存储
func (this *AppContext) dispatchEvent(evto *thspbs.Event) {
	switch evto.Name {
	case "SelfConnectionStatus":
	case "FriendRequest":
		///
		// pubkey := jso.Get("Args").GetIndex(0).MustString()
		pubkey := evto.Args[0]
		_, err := appctx.store.AddFriend(pubkey, 0, "", "")
		gopp.ErrPrint(err, evto.Args)

	case "FriendMessage":
		// jso.Get("Args").GetIndex(0).MustString()
		// msg := jso.Get("Args").GetIndex(1).MustString()
		msg := evto.Args[1]
		// fname := jso.Get("Margs").GetIndex(0).MustString()
		fname := evto.Margs[0]
		// pubkey := jso.Get("Margs").GetIndex(1).MustString()
		pubkey := evto.Margs[1]
		_ = fname
		///
		// eventId := int64(gopp.MustInt(jso.Get("Margs").GetIndex(2).MustString()))
		eventId := gopp.MustInt64(evto.Margs[2])
		_, err := appctx.store.AddFriendMessage(msg, pubkey, eventId)
		gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		// fname := jso.Get("Margs").GetIndex(0).MustString()
		fname := evto.Margs[0]
		// pubkey := jso.Get("Margs").GetIndex(1).MustString()
		pubkey := evto.Margs[1]
		_, _ = fname, pubkey

	case "ConferenceInvite":
		// groupNumber := jso.Get("Margs").GetIndex(2).MustString()
		groupNumber := evto.Margs[2]
		_ = groupNumber
		// cookie := jso.Get("Args").GetIndex(2).MustString()
		cookie := evto.Args[2]
		groupId := ConferenceCookieToIdentifier(cookie)

		///
		_, err := appctx.store.AddGroup(groupId, uint32(gopp.MustInt(groupNumber)), "")
		gopp.ErrPrint(err)

	case "ConferenceTitle":
		// groupTitle := jso.Get("Args").GetIndex(2).MustString()
		groupTitle := evto.Args[2]
		// groupId := jso.Get("Margs").GetIndex(0).MustString()
		groupId := evto.Margs[0]
		if ConferenceIdIsEmpty(groupId) {
			break
		}
		_, err := appctx.store.SetGroup(groupId, 0, groupTitle)
		gopp.ErrPrint(err, groupTitle)

	case "ConferencePeerName":
		// TODO
		// peerNumber := gopp.MustUint32(jso.Get("Args").GetIndex(1).MustString())
		peerNumber := gopp.MustUint32(evto.Args[1])
		// peerName := jso.Get("Margs").GetIndex(1).MustString()
		peerName := evto.Margs[1]
		// peerPubkey := jso.Get("Margs").GetIndex(1).MustString()
		peerPubkey := evto.Margs[2] // TODO two same???
		_, err := appctx.store.AddPeer(peerPubkey, 0, "")
		gopp.ErrPrint(err)
		if store.IsUniqueConstraintErr(err) {
			_, err = appctx.store.UpdatePeer(peerPubkey, peerNumber, peerName)
			gopp.ErrPrint(err)
		}
	case "ConferencePeerListChange":
		// TODO
	case "ConferenceNameListChange": // TODO depcreated
		// groupTitle := jso.Get("Margs").GetIndex(2).MustString()
		groupTitle := evto.Margs[2]
		// groupId := jso.Get("Margs").GetIndex(3).MustString()
		groupId := evto.Margs[3]
		if ConferenceIdIsEmpty(groupId) {
			log.Panicln("not possible", evto.Args)
		}
		_ = groupTitle
		///
		// peerPubkey := jso.Get("Margs").GetIndex(1).MustString()
		peerPubkey := evto.Margs[1]
		_, err := appctx.store.AddPeer(peerPubkey, 0, "")
		gopp.ErrPrint(err)

	case "ConferenceMessage":
		// groupId := jso.Get("Margs").GetIndex(3).MustString()
		groupId := evto.Margs[3]
		if ConferenceIdIsEmpty(groupId) {
			log.Panicln("not possible", evto.Args)
		}

		// message := jso.Get("Args").GetIndex(3).MustString()
		message := evto.Args[3]
		// peerName := jso.Get("Margs").GetIndex(0).MustString()
		peerName := evto.Margs[0]
		// peerPubkey := jso.Get("Margs").GetIndex(1).MustString()
		peerPubkey := evto.Margs[1]
		_, err := appctx.store.GetContactByPubkey(peerPubkey)
		if err == xorm.ErrNotExist {
			// peerNum := gopp.MustUint32(jso.Get("Args").GetIndex(1).MustString())
			peerNum := gopp.MustUint32(evto.Args[1])
			appctx.store.AddPeer(peerPubkey, peerNum, peerName)
		}
		// eventId := int64(gopp.MustInt(jso.Get("Margs").GetIndex(4).MustString()))
		eventId := gopp.MustInt64(evto.Margs[4])
		_, err = appctx.store.AddGroupMessage(message, "0", groupId, peerPubkey, eventId)
		gopp.ErrPrint(err, evto)

	default:
	}
}

func (this *AppContext) getAndPersistBaseInfo() {
	this.vtcli.GetBaseInfo()
	this.persistBaseInfo(this.vtcli.Binfo)
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
