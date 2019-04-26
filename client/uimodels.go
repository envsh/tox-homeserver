package client

import (
	"fmt"
	"gopp"
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
	lastno  int    // last valid status no, in case rpc

	// for ChatForm and SendForm
	// currently active state
	Frndinfo  thspbs.FriendInfo
	Grpinfo   thspbs.GroupInfo
	Cttype    int
	Ctnum     uint32
	Ctuniqid  string // current active contact identifier ==> cur
	Ctname    string // name or title
	Ctstmsg   string
	receiptid int64

	// scrollbar y position for all contact chat session
	// 对于每个会话的值，当活动窗口时，有新消息立即滚动到最底
	// 当切换离开一个窗口时，记录当前位置
	// 当切换到一个窗口时，如果没有新消息，则使用上次记录下的位置
	Scrollbarys map[string]int

	Friendsm map[string]*thspbs.FriendInfo // uniqid =>
	Friendsv []*thspbs.FriendInfo
	Groupsm  map[string]*thspbs.GroupInfo // uniqid =>
	Groupsv  []*thspbs.GroupInfo

	Ctmsgs map[string][]*Message // uniqid =>
	// Ctmsgs  map[string][]string   // uniqid =>
	Hasnews    map[string]int // uniqid => , 某个联系人的未读取消息个数
	lastmsg    *Message       // Lastmsg must be not belongs to active contact chatform
	lastctname string         // always according with lastmsg

	repainter func()
}

func NewDataModel(repainter func()) *DataModel {
	this := &DataModel{}
	this.repainter = repainter
	this.Mysttxt = Conno2str(0)
	this.receiptid = 10000

	this.Scrollbarys = map[string]int{}

	this.Friendsm = map[string]*thspbs.FriendInfo{}
	this.Friendsv = []*thspbs.FriendInfo{}
	this.Groupsm = map[string]*thspbs.GroupInfo{}
	this.Groupsv = []*thspbs.GroupInfo{}

	this.Ctmsgs = map[string][]*Message{}
	this.Hasnews = map[string]int{}

	return this
}

func (this *DataModel) Nxtreceiptid() int64 { return atomic.AddInt64(&this.receiptid, 1) }
func (this *DataModel) emitChanged() {
	if this.repainter != nil {
		this.repainter()
	}
}

func (this *DataModel) SetMyInfo(name, id string, stmsg string) {
	defer this.emitChanged()
	this.mu.Lock()
	defer this.mu.Unlock()

	this.Myname = name
	this.Myid = id
	this.Mystmsg = stmsg
}
func (this *DataModel) SetMyConnStatus(stno int) {
	defer this.emitChanged()
	this.mu.Lock()
	defer this.mu.Unlock()

	if stno == 5 { // brokenrpc
		this.lastno = this.Mystno
		this.Mystno = stno
	} else if stno == -5 { // goodrpc
		if this.Mystno == 5 {
			this.Mystno = this.lastno
		}
	} else {
		this.Mystno = stno
	}

	this.Mysttxt = Conno2str(this.Mystno)
}
func Conno2str(stno int) string {
	switch stno {
	case 0:
		return "NOE"
	case 1:
		return "TCP"
	case 2:
		return "UDP"
	case 5:
		return "BRK" // connect between client and homeserver
	default:
		return "UNK"
	}
}
func Conno2str1(stno int) string { return Conno2str(stno)[:1] }

func (this *DataModel) SetFriendInfos(friends map[uint32]*thspbs.FriendInfo) {
	defer this.emitChanged()
	newedm := map[string]*thspbs.FriendInfo{}
	newedv := []*thspbs.FriendInfo{}
	for _, v := range friends {
		f := *v
		newedm[v.GetPubkey()] = &f
		newedv = append(newedv, &f)
	}
	sort.Slice(newedv, func(i int, j int) bool { return newedv[i].GetFnum() < newedv[j].GetFnum() })
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Friendsm = newedm
	this.Friendsv = newedv
}
func (this *DataModel) SetGroupInfos(groups map[uint32]*thspbs.GroupInfo) {
	defer this.emitChanged()
	newedm := map[string]*thspbs.GroupInfo{}
	newedv := []*thspbs.GroupInfo{}
	for _, v := range groups {
		g := *v
		newedm[v.GetGroupId()] = &g
		newedv = append(newedv, &g)
	}
	sort.Slice(newedv, func(i int, j int) bool { return newedv[i].GetGnum() < newedv[j].GetGnum() })

	this.mu.Lock()
	defer this.mu.Unlock()
	this.Groupsm = newedm
	this.Groupsv = newedv
}

func (this *DataModel) FriendList() (rets []*thspbs.FriendInfo) {
	this.mu.RLock()
	defer this.mu.RUnlock()
	for _, e := range this.Friendsv {
		t := *e
		rets = append(rets, &t)
	}
	return
}
func (this *DataModel) GroupList() (rets []*thspbs.GroupInfo) {
	this.mu.RLock()
	defer this.mu.RUnlock()
	for _, e := range this.Groupsv {
		t := *e
		rets = append(rets, &t)
	}
	return
}
func (this *DataModel) CurMembers() (rets []*thspbs.MemberInfo) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	for _, v := range this.Grpinfo.GetMembers() {
		t := *v
		rets = append(rets, &t)
	}
	sort.Slice(rets, func(i int, j int) bool { return rets[i].GetPubkey() > rets[j].GetPubkey() })
	return
}

// current
func (this *DataModel) setFriendInfo(fi *thspbs.FriendInfo) {
	// this.mu.Lock()
	// defer this.mu.Unlock()
	this.Frndinfo = *fi
	this.Cttype = CTTYPE_FRIEND
	this.Ctname = fi.Name
	this.Ctstmsg = fi.Stmsg
	this.Ctnum = fi.GetFnum()
}

// current
func (this *DataModel) setGroupInfo(fi *thspbs.GroupInfo) {
	// this.mu.Lock()
	// defer this.mu.Unlock()
	this.Grpinfo = *fi
	this.Cttype = CTTYPE_GROUP
	this.Ctname = fi.GetTitle()
	this.Ctstmsg = fi.GetStmsg()
	this.Ctnum = fi.GetGnum()
}

func (this *DataModel) Switchtoct(uniqid string) {
	defer this.emitChanged()
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Ctuniqid = uniqid

	for _, v := range this.Groupsm {
		if v.GroupId == uniqid {
			this.setGroupInfo(v)
			return
		}
	}
	for _, v := range this.Friendsm {
		if v.GetPubkey() == uniqid {
			this.setFriendInfo(v)
			return
		}
	}
}

const maxinmemmsgcnt = 5000

func (this *DataModel) Newmsg(uniqid string, msg *Message) {
	defer this.emitChanged()
	this.mu.Lock()
	defer this.mu.Unlock()

	this.Ctmsgs[uniqid] = append(this.Ctmsgs[uniqid], msg)
	this.Hasnews[uniqid] += 1

	if uniqid != this.Ctuniqid {
		this.lastmsg = msg
		if cto, ok := this.Groupsm[uniqid]; ok {
			this.lastctname = cto.GetTitle()
		} else if cto, ok := this.Friendsm[uniqid]; ok {
			this.lastctname = cto.GetName()
		}
	}
}

func (this *DataModel) Lastmsg() string {
	this.mu.RLock()
	defer this.mu.RUnlock()
	msgo := this.lastmsg
	if msgo == nil {
		return ""
	}

	return fmt.Sprintf("%s> %s: %s", this.lastctname, msgo.PeerNameUi, msgo.MsgUi)
}

func (this *DataModel) Hasnewmsg(uniqid string) bool {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.Hasnews[uniqid] > 0
}
func (this *DataModel) Unsetnewmsg(uniqid string) {
	defer this.emitChanged()
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Hasnews[uniqid] = 0
}
func (this *DataModel) NewMsgcount(uniqid string) int {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.Hasnews[uniqid]
}
func (this *DataModel) Msgcount(uniqid string) int {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return len(this.Ctmsgs[uniqid])
}

func (this *DataModel) TotalCurrMsgcount() (cur, tot int) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	for _, v := range this.Ctmsgs {
		tot += len(v)
	}
	return len(this.Ctmsgs[this.Ctuniqid]), tot
}

// like: limit m, offset n
func (this *DataModel) Getmsgs(uniqid string, limit int, start ...int) {

}

func (this *DataModel) GetNewestMsgs(uniqid string, limit int) []*Message {
	this.mu.RLock()
	defer this.mu.RUnlock()

	msgs := this.Ctmsgs[uniqid]
	totcnt := len(msgs)

	rets := []*Message{}
	startpos := gopp.Max([]int{0, totcnt - 1 - limit}).(int)
	for idx := startpos; idx < totcnt && idx < totcnt; idx++ {
		rets = append(rets, msgs[idx])
	}

	return rets
}
