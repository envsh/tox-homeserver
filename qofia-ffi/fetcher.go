package main

// 历史消息获取
import (
	"gopp"
	"log"
	"time"
	"tox-homeserver/store"
)

type Fetcher struct {
}

var hisfet = &Fetcher{}

func (this *Fetcher) loadByContactId(ct_id uint32) ([]*store.Message, error) {

	return nil, nil
}

// 加载前处理
// 根据当前的prev_batch与本地数据中的next_batch做比较，得出是本地加载还是远程加载
// 加载后处理
// 通知UI
// 写入库
// 更新该好友的同步状态信息
// 更新当前contact item的实时同步状态信息
func (this *Fetcher) LoadPrevHistory(item *RoomListItem) {
	msgos, err := appctx.GetLigTox().LoadEventsByContactId(item.GetId(), item.prevBatch)
	log.Println(msgos)
	gopp.ErrPrint(err)
	log.Println(len(msgos))

	this.RefreshPrevRealtime(item, msgos)
	this.NotifyUiPrevHistory(item, msgos)
	this.SavePrevHistory(msgos)
	this.RefreshPrevHistory(msgos)
}

func (this *Fetcher) RefreshPrevRealtime(item *RoomListItem, msgos []store.Message) {
	min := item.prevBatch
	for _, msgo := range msgos {
		if msgo.EventId <= min {
			min = msgo.EventId - 1
		}
	}
	log.Println(min, item.prevBatch)
	if min < item.prevBatch {
		log.Printf("prev_batch switch from %d to %d\n", item.prevBatch, min)
		// item.prevBatch = min
	}
}

func NewMessageFromStoreRecord(m *store.Message) *Message {
	this := &Message{}
	this.EventId = m.EventId
	this.Msg = m.Content
	this.Peer = "unknown" // ??? TODO
	cto := appctx.GetStorage().GetContactById(m.ContactId)
	if cto != nil {
		this.Peer = cto.Name
	}

	defaultTimeStringLayout := "2006-01-02 15:04:05.999999999 -0700 MST"
	tim, err := time.Parse(defaultTimeStringLayout, m.Updated[:len(defaultTimeStringLayout)])
	gopp.ErrPrint(err)
	this.Time = tim

	this.refmtmsg()
	return this
}

func (this *Fetcher) NotifyUiPrevHistory(item *RoomListItem, msgos []store.Message) {
	for i := len(msgos) - 1; i >= 0; i-- {
		msgoe := msgos[i]
		msgou := NewMessageFromStoreRecord(&msgoe)
		runOnUiThread(func() {
			item.AddMessage(msgou, true)
		})
	}
	if len(msgos) > 0 {
		uictx.mech.Trigger()
	}
}

func (this *Fetcher) SavePrevHistory(msgos []store.Message) {
	for i := len(msgos) - 1; i >= 0; i-- {
		msgoe := msgos[i]
		_ = msgoe
	}
}

func (this *Fetcher) RefreshPrevHistory(msgos []store.Message) {

}
