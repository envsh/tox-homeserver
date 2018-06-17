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
	devo  *store.Device
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
	appctx.devo = appctx.store.GetDevice()
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
		msg := evto.Args[1]
		fname := evto.Margs[0]
		friendpk := evto.Margs[1]
		_ = fname
		///
		selfpk := this.vtcli.SelfGetPublicKey()
		eventId := gopp.MustInt64(evto.Margs[2])
		_, err := appctx.store.AddFriendMessage(msg, friendpk, selfpk, eventId, evto.UserCode)
		gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		fname := evto.Margs[0]
		pubkey := evto.Margs[1]
		_, _ = fname, pubkey

	case "ConferenceInvite":
		groupNumber := evto.Margs[2]
		_ = groupNumber
		cookie := evto.Args[2]
		groupId := ConferenceCookieToIdentifier(cookie)

		///
		_, err := appctx.store.AddGroup(groupId, uint32(gopp.MustInt(groupNumber)), "")
		gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupTitle := evto.Args[2]
		groupId := evto.Margs[0]
		if ConferenceIdIsEmpty(groupId) {
			break
		}
		_, err := appctx.store.SetGroup(groupId, 0, groupTitle)
		gopp.ErrPrint(err, groupTitle)

	case "ConferencePeerName":
		// TODO
		peerNumber := gopp.MustUint32(evto.Args[1])
		peerName := evto.Margs[1]
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
		groupTitle := evto.Margs[2]
		groupId := evto.Margs[3]
		if ConferenceIdIsEmpty(groupId) {
			log.Panicln("not possible", evto.Args)
		}
		_ = groupTitle
		///
		peerPubkey := evto.Margs[1]
		_, err := appctx.store.AddPeer(peerPubkey, 0, "")
		gopp.ErrPrint(err)

	case "ConferenceMessage":
		groupId := evto.Margs[3]
		if ConferenceIdIsEmpty(groupId) {
			log.Panicln("not possible", evto.Args)
		}

		message := evto.Args[3]
		peerName := evto.Margs[0]
		peerPubkey := evto.Margs[1]
		_, err := appctx.store.GetContactByPubkey(peerPubkey)
		if err == xorm.ErrNotExist {
			peerNum := gopp.MustUint32(evto.Args[1])
			appctx.store.AddPeer(peerPubkey, peerNum, peerName)
		}
		eventId := gopp.MustInt64(evto.Margs[4])
		_, err = appctx.store.AddGroupMessage(message, "0", groupId, peerPubkey, eventId, evto.UserCode)
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
