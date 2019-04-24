package client

import (
	"fmt"
	"sort"
	"sync/atomic"
	"tox-homeserver/thspbs"

	"github.com/sasha-s/go-deadlock"
)

const CTTYPE_FRIEND = 1
const CTTYPE_GROUP = 2

// 绘制UI界面的快照数据
// 共享给所有的views
type DataModel struct {
	mu      deadlock.RWMutex
	Myid    string
	Myname  string
	Mystmsg string
	Mystno  int
	Mysttxt string // status text

	// for ChatForm and SendForm
	// currently active state
	Frndinfo  thspbs.FriendInfo
	Grpinfo   thspbs.GroupInfo
	Cttype    int
	Ctnum     uint32
	Ctuniqid  string
	Ctname    string // name or title
	Ctstmsg   string
	receiptid int64

	// scrollbar y position for all contact chat session
	// 对于每个会话的值，当活动窗口时，有新消息立即滚动到最底
	// 当切换离开一个窗口时，记录当前位置
	// 当切换到一个窗口时，如果没有新消息，则使用上次记录下的位置
	Scrollbarys map[string]int

	Friendsm map[uint32]*thspbs.FriendInfo
	Friendsv []*thspbs.FriendInfo
	Groupsm  map[uint32]*thspbs.GroupInfo
	Groupsv  []*thspbs.GroupInfo

	Ctmsgs  map[string][]string // uniqid =>
	Hasnews map[string]int      // uniqid => , 某个联系人的未读取消息个数
	// uniqid  string // current active contact identifier ==> cur
}

func NewDataModel() *DataModel {
	this := &DataModel{}
	this.Scrollbarys = map[string]int{}

	this.Friendsm = map[uint32]*thspbs.FriendInfo{}
	this.Friendsv = []*thspbs.FriendInfo{}
	this.Groupsm = map[uint32]*thspbs.GroupInfo{}
	this.Groupsv = []*thspbs.GroupInfo{}

	this.Ctmsgs = map[string][]string{}
	this.Hasnews = map[string]int{}

	return this
}

func (this *DataModel) Nxtreceiptid() int64 { return atomic.AddInt64(&this.receiptid, 1) }

func (this *DataModel) SetMyInfo(name, id string, stmsg string) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.Myname = name
	this.Myid = id
	this.Mystmsg = stmsg
}
func (this *DataModel) SetMyConnStatus(stno int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.Mystno = stno
	this.Mysttxt = Conno2str(stno)
}
func Conno2str(stno int) string {
	switch stno {
	case 0:
		return "NOE"
	case 1:
		return "TCP"
	case 2:
		return "UDP"
	default:
		return "UNK"
	}
}

func (this *DataModel) SetFriendInfos(friends map[uint32]*thspbs.FriendInfo) {
	newedm := map[uint32]*thspbs.FriendInfo{}
	newedv := []*thspbs.FriendInfo{}
	for k, v := range friends {
		f := *v
		newedm[k] = &f
		newedv = append(newedv, &f)
	}
	sort.Slice(newedv, func(i int, j int) bool { return newedv[i].GetFnum() < newedv[j].GetFnum() })
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Friendsm = newedm
	this.Friendsv = newedv
}
func (this *DataModel) SetGroupInfos(groups map[uint32]*thspbs.GroupInfo) {
	newedm := map[uint32]*thspbs.GroupInfo{}
	newedv := []*thspbs.GroupInfo{}
	for k, v := range groups {
		g := *v
		newedm[k] = &g
		newedv = append(newedv, &g)
	}
	sort.Slice(newedv, func(i int, j int) bool { return newedv[i].GetGnum() < newedv[j].GetGnum() })

	this.mu.Lock()
	defer this.mu.Unlock()
	this.Groupsm = newedm
	this.Groupsv = newedv
}

// current
func (this *DataModel) SetFriendInfo(fi *thspbs.FriendInfo) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Frndinfo = *fi
	this.Cttype = CTTYPE_FRIEND
	this.Ctname = fmt.Sprintf("友 %s", fi.Name)
	this.Ctstmsg = fi.Stmsg
}

// current
func (this *DataModel) SetGroupInfo(fi *thspbs.GroupInfo) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Grpinfo = *fi
	this.Cttype = CTTYPE_GROUP
	this.Ctname = fmt.Sprintf("群 %s", fi.GetTitle())
	this.Ctstmsg = fi.GetStmsg()
}

func (this *DataModel) Switchtoct(uniqid string) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Ctuniqid = uniqid
}

func (this *DataModel) Setcontact(cttype int, ctnum uint32) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Cttype = cttype
	this.Ctnum = ctnum
}

func (this *DataModel) Newmsg(uniqid string, msg string) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.Ctmsgs[uniqid] = append(this.Ctmsgs[uniqid], msg)
	this.Hasnews[uniqid] += 1
}

func (this *DataModel) Hasnewmsg(uniqid string) bool {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.Hasnews[uniqid] > 0
}
func (this *DataModel) NewMsgcount(uniqid string) int {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.Hasnews[uniqid]
}

func (this *DataModel) Msgcount(uniqid string) int {
	return 0
}

// like: limit m, offset n
func (this *DataModel) Getmsgs(uniqid string, limit int, start ...int) {

}
