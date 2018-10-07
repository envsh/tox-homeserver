package main

// 历史消息获取
import (
	"gopp"
	"log"
	"math/rand"
	"time"
	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"
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
func (this *Fetcher) PullPrevHistoryByRoomItem(item *RoomListItem) {
	this.PullPrevHistoryById(item.GetId(), item.timeline.PrevBatch)
}

func (this *Fetcher) PullPrevHistoryById(pubkey string, prev_batch int64) {
	msgos, err := appctx.GetLigTox().PullEventsByContactId(pubkey, prev_batch)
	gopp.ErrPrint(err)

	item := uictx.iteman.Get(pubkey)
	if item == nil {
		log.Println("wtf...", pubkey)
		return
	}

	// 判断 >0 防止可能出现错误导致这里也错? 可以用err判断吧
	// if len(msgos) > 0 && len(msgos) < common.PullPageSize {
	if len(msgos) < thscom.PullPageSize {
		// think as no more data
		item.timeline.PrevBatch = 0 // 该room同步结束
	}

	log.Println()
	this.RefreshPrevRealtime(item, msgos)
	log.Println()
	this.NotifyUiPrevHistory(item, msgos)
	log.Println()
	this.SavePrevHistory(msgos)
	log.Println()
	this.RefreshPrevStorageByItem(item, pubkey)
	log.Println()
}

func (this *Fetcher) RefreshPrevRealtime(item *RoomListItem, msgos []store.MessageJoined) {
	min := item.timeline.PrevBatch
	for _, msgo := range msgos {
		if msgo.EventId <= min {
			min = msgo.EventId - 1
		}
	}

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

	rtl := &thscli.TimeLine{NextBatch: itemtl.NextBatch, PrevBatch: itemtl.PrevBatch}
	rtl, mrgcnt := thscli.MergeTimeLinesCount(rtl, tls)

	// TODO 删除与添加事务起来
	// 删除合并掉的
	if mrgcnt > 0 {
		log.Println("mrgsome:", mrgcnt, rtl, "left:", len(tls[mrgcnt:]))
		for i := 0; i < mrgcnt; i++ {
			err := st.DeleteSyncInfoById(sis[i].Id)
			gopp.ErrPrint(err, sis[i])
		}
	}

	// 写入新的部分
	c, err := st.GetContactByPubkey(pubkey)
	gopp.ErrPrint(err)
	err = st.AddSyncInfo(c.Id, rtl.NextBatch, rtl.PrevBatch)
	gopp.ErrPrint(err, c)
	log.Println("runtime/storage timeline:", itemtl, rtl, mrgcnt, itemName)
}

func NewMessageFromStoreRecord(m *store.MessageJoined) *Message {
	this := &Message{}
	this.EventId = m.EventId
	this.Msg = m.Content
	this.PeerName = gopp.IfElseStr(m.PeerName == "", thscom.DefaultUserName, m.PeerName)
	this.Sent = m.Sent > 0
	this.UserCode = m.UserCode

	defaultTimeStringLayout := thscom.DefaultTimeLayout
	tm, err := time.Parse(defaultTimeStringLayout, m.Updated)
	gopp.ErrPrint(err)
	this.Time = tm

	this.Me = m.PeerPubkey == vtcli.SelfGetPublicKey()

	this.refmtmsg()
	return this
}

func (this *Fetcher) NotifyUiPrevHistory(item *RoomListItem, msgos []store.MessageJoined) {
	for i := 0; i < len(msgos); i++ {
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

func (this *Fetcher) SavePrevHistory(msgos []store.MessageJoined) {
	for i := 0; i < len(msgos); i++ {
		msgoe := msgos[i]
		_ = msgoe
		_, err := appctx.GetStorage().AddMessageJoined(&msgoe)
		gopp.FalsePrint(store.IsUniqueConstraintErr(err), err, i)
	}
}

// by BaseInfo
func pullAllRoomsLatestMessages() {
	// rand wait range: [20,70]ms
	log.Println("Loading all rooms's latest messages....", len(vtcli.Binfo.Friends), len(vtcli.Binfo.Groups))
	btime := time.Now()
	time.Sleep(time.Duration(rand.Int()%50+20) * time.Millisecond)
	for idx, frnd := range vtcli.Binfo.Friends {
		log.Println("pulling...", idx)
		hisfet.PullPrevHistoryById(frnd.Pubkey, vtcli.Binfo.NextBatch)
		log.Println("pulled", idx)
	}

	for idx, grp := range vtcli.Binfo.Groups {
		log.Println("pulling...", idx)
		hisfet.PullPrevHistoryById(grp.GroupId, vtcli.Binfo.NextBatch)
		log.Println("pulled", idx)
	}
	log.Println("Load all rooms's latest messages done.", len(vtcli.Binfo.Friends),
		len(vtcli.Binfo.Groups), time.Since(btime))
	// about 350ms with 7 contacts
}
