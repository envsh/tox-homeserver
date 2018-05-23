package main

// 历史消息获取
import (
	"gopp"
	"log"
	"math/rand"
	"time"
	thscli "tox-homeserver/client"
	"tox-homeserver/common"
	"tox-homeserver/store"
)

// 两类取数据
// 从远程取使用前缀pull
// 从本地取使用前缀load
type Fetcher struct {
}

var hisfet = &Fetcher{}

func (this *Fetcher) pullByContactId(ct_id uint32) ([]*store.Message, error) {

	return nil, nil
}

// 加载前处理
// 根据当前的prev_batch与本地数据中的next_batch做比较，得出是本地加载还是远程加载
// 加载后处理
// 通知UI
// 写入库
// 更新该好友的同步状态信息
// 更新当前contact item的实时同步状态信息
func (this *Fetcher) PullPrevHistoryByRoom(item *RoomListItem) {
	this.PullPrevHistoryById(item.GetId(), item.timeline.PrevBatch)
}

func (this *Fetcher) PullPrevHistoryById(pubkey string, prev_batch int64) {
	msgos, err := appctx.GetLigTox().PullEventsByContactId(pubkey, prev_batch)
	log.Println(msgos)
	gopp.ErrPrint(err)
	log.Println(len(msgos))

	item := uictx.iteman.Get(pubkey)
	if item == nil {
		log.Println("wtf...", pubkey)
		return
	}

	// 判断 >0 防止可能出现错误导致这里也错? 可以用err判断吧
	// if len(msgos) > 0 && len(msgos) < common.PullPageSize {
	if len(msgos) < common.PullPageSize {
		// think as no more data
		item.timeline.PrevBatch = 0 // 该room同步结束
	}

	this.RefreshPrevRealtime(item, msgos)
	this.NotifyUiPrevHistory(item, msgos)
	this.SavePrevHistory(msgos)
	this.RefreshPrevStorageByItem(item, pubkey)
}

func (this *Fetcher) RefreshPrevRealtime(item *RoomListItem, msgos []store.Message) {
	min := item.timeline.PrevBatch
	for _, msgo := range msgos {
		if msgo.EventId <= min {
			min = msgo.EventId - 1
		}
	}
	log.Println(min, item.timeline.PrevBatch)
	if min < item.timeline.PrevBatch {
		log.Printf("%s's prev_batch switch from %d to %d\n", item.GetName(), item.timeline.PrevBatch, min)
		item.timeline.PrevBatch = min
	}
}

// 从库中查询sync info,尝试做更新/合并数据库信息
func (this *Fetcher) RefreshPrevStorageByItem(item *RoomListItem, pubkey string) {
	this.RefreshPrevStorageTimeLine(&item.timeline, pubkey, item.GetName())
}

func (this *Fetcher) RefreshPrevStorageTimeLine(itemtl *thscli.TimeLine, pubkey string, itemName string) {
	st := appctx.GetStorage()
	_ = st

	sis, err := st.GetTimeLinesByPubkey(pubkey)
	gopp.ErrPrint(err, sis)
	tls := thscli.SyncInfos2TimeLines(sis)

	mrgcnt := 0
	rtl := &thscli.TimeLine{NextBatch: itemtl.NextBatch, PrevBatch: itemtl.PrevBatch}
	for _, tl := range tls {
		ntl, can := rtl.Merge(tl)
		log.Println("mrgres:", rtl, tl, can)
		if can {
			rtl = ntl
			mrgcnt += 1
		} else {
			break
		}
	}

	// TODO 删除与添加事务起来
	c, err := st.GetContactByPubkey(pubkey)
	gopp.ErrPrint(err)
	// err = st.DeleteTimeLinesByPubkey(pubkey)
	// gopp.ErrPrint(err, pubkey)
	if mrgcnt > 0 {
		log.Println("mrgsome:", mrgcnt, rtl, tls[mrgcnt:])
		// 删除合并掉的
		for i := 0; i < mrgcnt; i++ {
			err := st.DeleteSyncInfoById(sis[i].Id)
			gopp.ErrPrint(err, sis[i])
		}
	} else {
		// save one
	}
	// 写入新的部分
	err = st.AddSyncInfo(c.Id, rtl.NextBatch, rtl.PrevBatch)
	gopp.ErrPrint(err, c)
	log.Println("rt vs. st timeline:", itemtl, rtl, mrgcnt, itemName)
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
		_, err := appctx.GetStorage().AddMessage(&msgoe)
		gopp.ErrPrint(err, i)
	}
}

// by BaseInfo
func pullAllRoomsLatestMessages() {
	// rand wait range: [20,70]ms
	log.Println("Loading all rooms's latest messages....", len(vtcli.Binfo.Friends), len(vtcli.Binfo.Groups))
	btime := time.Now()
	time.Sleep(time.Duration(rand.Int()%50+20) * time.Millisecond)
	for _, frnd := range vtcli.Binfo.Friends {
		hisfet.PullPrevHistoryById(frnd.Pubkey, vtcli.Binfo.NextBatch)
	}

	for _, grp := range vtcli.Binfo.Groups {
		hisfet.PullPrevHistoryById(grp.GroupId, vtcli.Binfo.NextBatch)
	}
	log.Println("Load all rooms's latest messages done.", len(vtcli.Binfo.Friends),
		len(vtcli.Binfo.Groups), time.Since(btime))
	// about 350ms with 7 contacts
}
