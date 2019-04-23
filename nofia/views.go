package main

import (
	"fmt"
	"log"
	"sort"
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
			ctx.LayoutRowStatic(30, 100, 2)
			// ctx.Label("名字", 10)
			ctx.Label(this.name, 10)
			// ctx.Label("连接状态", 10)
			ctx.Label(this.sttxt, 10)
			ctx.LayoutRowStatic(30, 100, 1)
			// ctx.Label("状态文本", 10)
			ctx.Label(this.stmsg, 10)
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
			ctx.LayoutRowBegin(nk.STATIC, 30, 3)
			ctx.LayoutRowPush(30)
			ctx.Label("空格", 10)

			ctx.LayoutRowPush(400)
			ctx.Label(this.name, 10)

			ctx.LayoutRowPush(60)
			if ctx.ButtonLabel("选项") != nil {
			}
			ctx.LayoutRowEnd()

			ctx.LayoutRowDynamic(30, 1)
			ctx.Label(this.stmsg, 10)
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
			const maxlen = 10
			msgs := this.msgs[this.uniqid]
			if len(msgs) > maxlen {
				msgs = msgs[len(msgs)-10:]
			}
			for idx, msg := range msgs {
				tmsg := fmt.Sprintf("%d %s", idx, msg)
				ctx.LayoutRowBegin(nk.STATIC, 39, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("|")
				ctx.LayoutRowPush(430)
				ctx.Label(tmsg, 10)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("|")
				ctx.LayoutRowEnd()
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

type SendForm struct {
	iptbuf  []byte
	iptblen int
	iptres  []byte
}

func NewSendForm() *SendForm {
	this := &SendForm{}
	this.iptbuf = make([]byte, 32)
	return this
}

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
			newlen := this.iptblen
			ctx.EditString(nk.EDIT_FIELD, this.iptbuf, &newlen, len(this.iptbuf))
			if this.iptblen != newlen {
				this.iptblen = newlen
				log.Println("text", string(this.iptbuf[:newlen]), newlen)
			}
			ctx.LayoutRowPush(80)
			if ctx.ButtonLabel("发送按钮") != nil {
			}
			ctx.LayoutRowEnd()

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}
