package server

import (
	"gopp"
	"log"
	"sync"

	"tox-homeserver/store"

	"github.com/go-xorm/xorm"
)

// on friend connect status > 0
// 1: select unsents from db
// 2: keep unsents in memory
// for group offline is when just created and can not send message. otherwise group is always online.

type OfflineMessageManager struct {
	mu     sync.RWMutex
	msgs   map[string][]store.Message // pubkey =>
	loaded bool
}

var offmsgman *OfflineMessageManager
var offmsgmanOnce sync.Once

func OffMsgMan() *OfflineMessageManager { return NewOfflineMessageManager() }

func NewOfflineMessageManager() *OfflineMessageManager {
	offmsgmanOnce.Do(func() {
		offmsgman_ := &OfflineMessageManager{}
		offmsgman_.msgs = make(map[string][]store.Message)
		offmsgman = offmsgman_
	})
	return offmsgman
}

func (this *OfflineMessageManager) GetByPubkey(pubkey string) (offmsgs []store.Message) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	offmsgs = this.msgs[pubkey]
	return
}

func (this *OfflineMessageManager) DeleteMessage(pubkey string, msgid int64) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	if msgs, ok := this.msgs[pubkey]; ok {
		newmsgs := []store.Message{}
		for _, msg := range msgs {
			if msg.Id != msgid {
				newmsgs = append(newmsgs, msg)
			}
		}
		this.msgs[pubkey] = newmsgs
	} else {
		return xorm.ErrNotExist
	}

	return nil
}

func (this *OfflineMessageManager) AddMessage(pubkey string, msg *store.Message) error {
	this.mu.Lock()
	defer this.mu.Unlock()
	if _, ok := this.msgs[pubkey]; !ok {
		this.msgs[pubkey] = nil
	}
	this.msgs[pubkey] = append(this.msgs[pubkey], *msg)
	return nil
}

func (this *OfflineMessageManager) LoadFromStorage(st *store.Storage) {
	if this.loaded {
		return
	}

	this.loaded = true
	unsents, err := st.FindUnsentMessages()
	gopp.ErrPrint(err)

	grped := make(map[int64][]store.Message)
	for _, unsent := range unsents {
		if _, ok := grped[unsent.RoomId]; !ok {
			grped[unsent.RoomId] = nil
		}
		grped[unsent.RoomId] = append(grped[unsent.RoomId], unsent)
	}

	this.mu.Lock()
	defer this.mu.Unlock()
	unsentsCount := 0
	for id, msgs := range grped {
		c := st.GetContactById(id)
		gopp.NilPrint(c, "Unknown contact:", id)
		if c != nil {
			this.msgs[c.Pubkey] = msgs
			unsentsCount += len(msgs)
		}
	}
	log.Println("Load unsents:", "contact(s):", len(this.msgs), "message(s):", unsentsCount)
}
