package client

import (
	"gopp"
	"log"
	"net/http"
	"sync"
	"time"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
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
	appctx.srvtp.OnData(func(jso *simplejson.Json, data []byte) { appctx.dispatchEvent(jso) })

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
func (this *AppContext) dispatchEvent(jso *simplejson.Json) {
	evtName := jso.Get("Name").MustString()
	switch evtName {
	case "SelfConnectionStatus":
	case "FriendRequest":
		///
		pubkey := jso.Get("Args").GetIndex(0).MustString()
		_, err := appctx.store.AddFriend(pubkey, 0, "", "")
		gopp.ErrPrint(err, jso.Get("Args"))

	case "FriendMessage":
		// jso.Get("Args").GetIndex(0).MustString()
		msg := jso.Get("Args").GetIndex(1).MustString()
		fname := jso.Get("Margs").GetIndex(0).MustString()
		pubkey := jso.Get("Margs").GetIndex(1).MustString()
		_ = fname
		///
		eventId := int64(gopp.MustInt(jso.Get("Margs").GetIndex(2).MustString()))
		_, err := appctx.store.AddFriendMessage(msg, pubkey, eventId)
		gopp.ErrPrint(err)

	case "FriendConnectionStatus":
		fname := jso.Get("Margs").GetIndex(0).MustString()
		pubkey := jso.Get("Margs").GetIndex(1).MustString()
		_, _ = fname, pubkey

	case "ConferenceInvite":
		groupNumber := jso.Get("Margs").GetIndex(2).MustString()
		_ = groupNumber
		cookie := jso.Get("Args").GetIndex(2).MustString()
		groupId := ConferenceCookieToIdentifier(cookie)

		///
		_, err := appctx.store.AddGroup(groupId, uint32(gopp.MustInt(groupNumber)), "")
		gopp.ErrPrint(err)

	case "ConferenceTitle":
		groupTitle := jso.Get("Args").GetIndex(2).MustString()
		groupId := jso.Get("Margs").GetIndex(0).MustString()
		if ConferenceIdIsEmpty(groupId) {
			break
		}
		_, err := appctx.store.SetGroup(groupId, 0, groupTitle)
		gopp.ErrPrint(err, groupTitle)

	case "ConferencePeerName":
		// TODO
		peerNumber := gopp.MustUint32(jso.Get("Args").GetIndex(1).MustString())
		peerName := jso.Get("Margs").GetIndex(1).MustString()
		peerPubkey := jso.Get("Margs").GetIndex(1).MustString()
		_, err := appctx.store.AddPeer(peerPubkey, 0, "")
		gopp.ErrPrint(err)
		if store.IsUniqueConstraintErr(err) {
			_, err = appctx.store.UpdatePeer(peerPubkey, peerNumber, peerName)
			gopp.ErrPrint(err)
		}
	case "ConferencePeerListChange":
		// TODO
	case "ConferenceNameListChange": // TODO depcreated
		groupTitle := jso.Get("Margs").GetIndex(2).MustString()
		groupId := jso.Get("Margs").GetIndex(3).MustString()
		if ConferenceIdIsEmpty(groupId) {
			log.Println("empty")
			break
		}
		_ = groupTitle
		///
		peerPubkey := jso.Get("Margs").GetIndex(1).MustString()
		_, err := appctx.store.AddPeer(peerPubkey, 0, "")
		gopp.ErrPrint(err)

	case "ConferenceMessage":
		groupId := jso.Get("Margs").GetIndex(3).MustString()
		if ConferenceIdIsEmpty(groupId) {
			break
		}

		message := jso.Get("Args").GetIndex(3).MustString()

		//
		peerName := jso.Get("Margs").GetIndex(0).MustString()
		peerPubkey := jso.Get("Margs").GetIndex(1).MustString()
		_, err := appctx.store.GetContactByPubkey(peerPubkey)
		if err == xorm.ErrNotExist {
			peerNum := gopp.MustUint32(jso.Get("Args").GetIndex(1).MustString())
			appctx.store.AddPeer(peerPubkey, peerNum, peerName)
		}
		eventId := int64(gopp.MustInt(jso.Get("Margs").GetIndex(4).MustString()))
		_, err = appctx.store.AddGroupMessage(message, "0", groupId, peerPubkey, eventId)
		gopp.ErrPrint(err, jso)

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
