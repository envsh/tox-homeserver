package main

import (
	"fmt"
	"gopp"
	"log"
	"sort"
	"sync/atomic"
	"tox-homeserver/thspbs"

	nk "mkuse/nuklear"

	"github.com/sasha-s/go-deadlock"
)

type MyinfoView struct {
	id    string
	name  string
	stmsg string
	sttxt string
}

func (this *MyinfoView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("我的信息", nk.NewRect(0, 0, 250, 120), nk.WINDOW_BORDER)
		if err != nil {
			ctx.LayoutRowBegin(nk.STATIC, 30, 2)
			name := gopp.IfElseStr(len(this.name) == 0, " ", this.name)
			sel0 := len(name)
			ctx.LayoutRowPush(200)
			ctx.SelectableLabel(name, 10, &sel0)
			ctx.LayoutRowPush(20)
			ctx.Label(this.sttxt, 10)
			ctx.LayoutRowEnd()
			ctx.LayoutRowDynamic(30, 1)
			stmsg := gopp.IfElseStr(len(this.stmsg) == 0, " ", this.stmsg)
			sel1 := len(stmsg)
			ctx.SelectableLabel(stmsg, 10, &sel1)
			ctx.LayoutRowStatic(30, 100, 2)
			ctx.Label("搜索框", 10)
			ctx.Label("排列过滤", 10)

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}

type MyactionView struct {
}

func (this *MyactionView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("我的控制按钮组", nk.NewRect(0, 600-50, 250, 60), nk.WINDOW_BORDER)
		if err != nil {

			ctx.LayoutRowBegin(nk.STATIC, 30, 4)
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("操作1") != nil {
			}
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("操作2") != nil {
			}
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("操作3") != nil {
			}
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("操作4") != nil {
			}
			ctx.LayoutRowEnd()
		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}

type FriendInfoView struct {
	frndinfo thspbs.FriendInfo
	grpinfo  thspbs.GroupInfo
	which    int // 1, friend, 2, group
	name     string
	stmsg    string
	mu       deadlock.RWMutex
}

func (this *FriendInfoView) setFriendInfo(fi *thspbs.FriendInfo) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.frndinfo = *fi
	this.which = 1
	this.name = fmt.Sprintf("友 %s", fi.Name)
	this.stmsg = fi.Stmsg
}

func (this *FriendInfoView) setGroupInfo(fi *thspbs.GroupInfo) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.grpinfo = *fi
	this.which = 2
	this.name = fmt.Sprintf("群 %s", fi.GetTitle())
	this.stmsg = fi.GetStmsg()
}

func (this *FriendInfoView) render() func(*nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("好友状态视图123", nk.NewRect(250, 0, 550, 90), nk.WINDOW_BORDER)
		if err != nil {
			this.mu.RLock()
			ctx.LayoutRowBegin(nk.STATIC, 30, 4)
			ctx.LayoutRowPush(30)
			ctx.ButtonLabel("空格")

			ctx.LayoutRowPush(360)
			name := gopp.IfElseStr(len(this.name) == 0, " ", this.name)
			sel0 := len(name)
			ctx.SelectableLabel(name, 5, &sel0)

			ctx.LayoutRowPush(80)
			if ctx.ButtonLabel("成员列表") != nil {
			}
			ctx.LayoutRowPush(40)
			if ctx.ButtonLabel("选项") != nil {
			}
			ctx.LayoutRowEnd()

			ctx.LayoutRowDynamic(30, 1)
			stmsg := gopp.IfElseStr(len(this.stmsg) == 0, " ", this.stmsg)
			sel1 := len(stmsg)
			ctx.SelectableLabel(stmsg, 10, &sel1)
			this.mu.RUnlock()
		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}

}

type ContectView struct {
	friendsm map[uint32]*thspbs.FriendInfo
	friendsv []*thspbs.FriendInfo
	groupsm  map[uint32]*thspbs.GroupInfo
	groupsv  []*thspbs.GroupInfo
	ctmu     deadlock.RWMutex
}

func NewcontactView() *ContectView {
	this := &ContectView{}
	this.friendsm = map[uint32]*thspbs.FriendInfo{}
	this.friendsv = []*thspbs.FriendInfo{}
	this.groupsm = map[uint32]*thspbs.GroupInfo{}
	this.groupsv = []*thspbs.GroupInfo{}
	return this
}

func (this *ContectView) setFriendInfos(friends map[uint32]*thspbs.FriendInfo) {
	newedm := map[uint32]*thspbs.FriendInfo{}
	newedv := []*thspbs.FriendInfo{}
	for k, v := range friends {
		f := *v
		newedm[k] = &f
		newedv = append(newedv, &f)
	}
	sort.Slice(newedv, func(i int, j int) bool { return newedv[i].GetFnum() < newedv[j].GetFnum() })
	this.ctmu.Lock()
	defer this.ctmu.Unlock()
	this.friendsm = newedm
	this.friendsv = newedv
}
func (this *ContectView) setGroupInfos(groups map[uint32]*thspbs.GroupInfo) {
	newedm := map[uint32]*thspbs.GroupInfo{}
	newedv := []*thspbs.GroupInfo{}
	for k, v := range groups {
		g := *v
		newedm[k] = &g
		newedv = append(newedv, &g)
	}
	sort.Slice(newedv, func(i int, j int) bool { return newedv[i].GetGnum() < newedv[j].GetGnum() })
	this.ctmu.Lock()
	defer this.ctmu.Unlock()
	this.groupsm = newedm
	this.groupsv = newedv
}

func (this *ContectView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("Hel呵呵lo", nk.NewRect(0, 120, 250, 600-160), nk.WINDOW_BORDER)
		if err != nil {
			this.ctmu.RLock()
			for _, v := range this.groupsv {
				name := fmt.Sprintf("群 %s", v.GetTitle())
				statxt := fmt.Sprintf("%d", 1)
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("III")
				ctx.LayoutRowPush(150)
				if ctx.ButtonLabel(name) != nil {
					log.Println("group clicked", v.GetGnum(), name)
					uictx.fiview.setGroupInfo(v)
					uictx.chatform.switchto(v.GetGroupId())
					uictx.sendv.setcontact(CTTYPE_GROUP, v.GetGnum())
				}
				ctx.LayoutRowPush(30)
				ctx.Label(statxt, 10)
			}
			for _, v := range this.friendsv {
				name := fmt.Sprintf("友 %s", v.GetName())
				statxt := fmt.Sprintf("%d", v.GetStatus())
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("III")
				ctx.LayoutRowPush(150)
				if ctx.ButtonLabel(name) != nil {
					log.Println("friend clicked", v.GetFnum(), name)
					uictx.fiview.setFriendInfo(v)
					uictx.chatform.switchto(v.GetPubkey())
					uictx.sendv.setcontact(CTTYPE_FRIEND, v.GetFnum())
				}
				ctx.LayoutRowPush(30)
				ctx.Label(statxt, 10)
			}
			this.ctmu.RUnlock()
			for i := 0; i < 16; i++ {
				name := fmt.Sprintf("好友名%d", i+1)
				statxt := fmt.Sprintf("联系人%d", i+1)
				ctx.LayoutRowStatic(30, 100, 2)
				ctx.Label(name, 10)
				ctx.Label(statxt, 10)
			}

			ctx.LayoutRowDynamic(510-3*30, 1)
			ctx.Label("空白区域", 10)

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}

/////
type ChatForm struct {
	msgs   map[string][]string
	uniqid string // current active contact identifier
	mu     deadlock.RWMutex
}

func (this *ChatForm) switchto(uniqid string) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.uniqid = uniqid
}
func (this *ChatForm) newmsg(uniqid string, msg string) {
	this.mu.Lock()
	defer this.mu.Unlock()
	// this.which = which
	this.msgs[uniqid] = append(this.msgs[uniqid], msg)
}

func NewChatForm() *ChatForm {
	this := &ChatForm{}
	this.msgs = map[string][]string{"": {
		"消息111", "消息222",
		"消息333 Unix/Linux shell script FAQ: Can you share a simple Linux shell script that shows how ",
		"消息444", "消息555", "消息666",
	},
	}

	return this
}

func (this *ChatForm) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("Hel呵呵lo2", nk.NewRect(250, 90, 550, 600-170), nk.WINDOW_BORDER)
		if err != nil {

			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息窗口", 10)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息1", 10)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息2", 10)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息3", 10)

			this.mu.RLock()
			// draw newest n msgs
			const maxlen = 30
			msgs := this.msgs[this.uniqid]
			if len(msgs) > maxlen {
				msgs = msgs[len(msgs)-30:]
			}

			for idx, msg := range msgs {
				tmsg := fmt.Sprintf("%d %s", idx, msg)
				wraped := gopp.Splitrnui(tmsg, 60)
				for idx, line := range wraped {
					ctx.LayoutRowBegin(nk.STATIC, 39, 3)
					ctx.LayoutRowPush(30)
					if idx == 0 {
						ctx.ButtonLabel("|")
					} else {
						ctx.Label(" ", 1)
					}
					ctx.LayoutRowPush(450)
					seln := len(line)
					ctx.SelectableLabel(line, gopp.IfElseInt(idx == 0, 1, 5), &seln)
					ctx.LayoutRowPush(30)
					if idx == 0 {
						ctx.ButtonLabel("|")
					} else {
						ctx.Label(" ", 1)
					}
					ctx.LayoutRowEnd()
				}
			}
			this.mu.RUnlock()

			// 1w条的时候内存倒没有问题，CPU上去了 10+%
			// 3k条以下比较好，滚动的时候使用3%上下的CPU
			for i := 1000; i < 3000; i++ {
				tmsg := fmt.Sprintf("聊天消息%d\x00", i)
				ctx.LayoutRowDynamic(30, 1)
				ctx.Label(tmsg, 10)
			}

			emptylen := 410 - float32(len(msgs)+4)*30
			if emptylen > 0 {
				ctx.LayoutRowDynamic(emptylen, 1)
				ctx.Label("空白区域", 10)
			}

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}

const CTTYPE_FRIEND = 1
const CTTYPE_GROUP = 2

type SendForm struct {
	cttype   int // contact type, group or friend
	ctnum    uint32
	mu       deadlock.RWMutex
	msgrspid int64 // 用于发送结果回执
	iptbuf   []byte
	iptblen  int
	iptres   []byte
}

func NewSendForm() *SendForm {
	this := &SendForm{}
	this.iptbuf = make([]byte, 320)
	this.msgrspid = 10000
	return this
}

func (this *SendForm) setcontact(cttype int, ctnum uint32) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.cttype = cttype
	this.ctnum = ctnum
}

func (this *SendForm) nextrspid() int64 { return atomic.AddInt64(&this.msgrspid, 1) }

func (this *SendForm) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("消息输入发送视图", nk.NewRect(250, 520, 550, 600-490), nk.WINDOW_BORDER)
		if err != nil {

			ctx.LayoutRowBegin(nk.STATIC, 30, 7)
			for i := 0; i < 7; i++ {
				txt := fmt.Sprintf("操作%d", i+1)
				ctx.LayoutRowPush(50)
				if ctx.ButtonLabel(txt) != nil {
				}
			}
			ctx.LayoutRowEnd()

			ctx.LayoutRowBegin(nk.STATIC, 30, 2)
			ctx.LayoutRowPush(500 - 80)
			cjkiptxt := uictx.app.PeekCJKInputString()
			newlen := this.iptblen
			active := ctx.EditString(nk.EDIT_FIELD, this.iptbuf, &newlen, len(this.iptbuf))
			// log.Println(active, len(cjkiptxt), cjkiptxt)
			if this.iptblen != newlen {
				this.iptblen = newlen
				log.Println("text", string(this.iptbuf[:newlen]), newlen)
			} else if active == 1 && len(cjkiptxt) > 0 {
				// copy(this.iptbuf[this.iptblen:], []byte(cjkiptxt))
				// this.iptblen += len(cjkiptxt)
			}
			ctx.LayoutRowPush(80)
			if ctx.ButtonLabel("发送按钮") != nil {
				if this.iptblen > 0 {
					this.mu.RLock()
					cttype := this.cttype
					ctnum := this.ctnum
					this.mu.RUnlock()
					rspid := this.nextrspid()
					msg := string(this.iptbuf[:this.iptblen])
					var err error
					switch cttype {
					case CTTYPE_FRIEND:
						_, err = vtcli.FriendSendMessage(ctnum, msg, rspid)
					case CTTYPE_GROUP:
						err = vtcli.ConferenceSendMessage(ctnum, 0, msg, rspid)
					default:
						err = fmt.Errorf("Unseted cttype %d %d", cttype, ctnum)
					}
					gopp.ErrPrint(err, cttype, ctnum)
					if err == nil {
						this.iptblen = 0
					}
				}
			}
			ctx.LayoutRowEnd()

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}
