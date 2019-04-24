package main

import (
	"fmt"
	"gopp"
	"log"

	nk "mkuse/nuklear"

	thscli "tox-homeserver/client"
)

type MyinfoView struct {
}

func (this *MyinfoView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("我的信息", nk.NewRect(0, 0, 250, 120), nk.WINDOW_BORDER)
		if err != nil {
			name := uictx.mdl.Myname
			name = gopp.IfElseStr(len(name) == 0, " ", name)
			sel0 := len(name)
			sttxt := uictx.mdl.Mysttxt
			ctx.LayoutRowBegin(nk.STATIC, 30, 2)
			ctx.LayoutRowPush(200)
			ctx.SelectableLabel(name, 10, &sel0)
			ctx.LayoutRowPush(20)
			ctx.Label(sttxt, 10)
			ctx.LayoutRowEnd()

			stmsg := uictx.mdl.Mystmsg
			stmsg = gopp.IfElseStr(len(stmsg) == 0, " ", stmsg)
			sel1 := len(stmsg)
			ctx.LayoutRowDynamic(30, 1)
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
}

func (this *FriendInfoView) render() func(*nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("好友状态视图123", nk.NewRect(250, 0, 550, 85), nk.WINDOW_NO_SCROLLBR)
		if err != nil {
			ctx.LayoutRowBegin(nk.STATIC, 30, 4)
			ctx.LayoutRowPush(40)
			ctx.ButtonLabel("空格")

			name := uictx.mdl.Ctname
			name = gopp.IfElseStr(len(name) == 0, " ", name)
			sel0 := len(name)
			ctx.LayoutRowPush(360)
			ctx.SelectableLabel(name, 5, &sel0)

			ctx.LayoutRowPush(80)
			if ctx.ButtonLabel("成员列表") != nil {
			}
			ctx.LayoutRowPush(40)
			if ctx.ButtonLabel("选项") != nil {
			}
			ctx.LayoutRowEnd()

			stmsg := uictx.mdl.Ctstmsg
			stmsg = gopp.IfElseStr(len(stmsg) == 0, " ", stmsg)
			sel1 := len(stmsg)
			ctx.LayoutRowDynamic(30, 1)
			ctx.SelectableLabel(stmsg, 10, &sel1)
		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}

}

type ContectView struct {
}

func NewcontactView() *ContectView {
	this := &ContectView{}
	return this
}

func (this *ContectView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("Hel呵呵lo", nk.NewRect(0, 120, 250, 600-160), nk.WINDOW_BORDER)
		if err != nil {
			for _, v := range uictx.mdl.GroupList() {
				name := fmt.Sprintf("群 %s", v.GetTitle())
				statxt := fmt.Sprintf("%d", uictx.mdl.NewMsgcount(v.GetGroupId()))
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("III")
				ctx.LayoutRowPush(150)
				if ctx.ButtonLabel(name) != nil {
					log.Println("group clicked", v.GetGnum(), name)
					uictx.mdl.Switchtoct(v.GetGroupId())
				}
				ctx.LayoutRowPush(30)
				ctx.Label(statxt, 10)
			}
			for _, v := range uictx.mdl.FriendList() {
				name := fmt.Sprintf("友 %s", v.GetName())
				statxt := fmt.Sprintf("%s %d",
					thscli.Conno2str1(int(v.Status)), uictx.mdl.NewMsgcount(v.GetPubkey()))
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("III")
				ctx.LayoutRowPush(150)
				if ctx.ButtonLabel(name) != nil {
					log.Println("friend clicked", v.GetFnum(), name)
					uictx.mdl.Switchtoct(v.GetPubkey())
				}
				ctx.LayoutRowPush(30)
				ctx.Label(statxt, 10)
			}

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
}

func NewChatForm() *ChatForm {
	this := &ChatForm{}
	return this
}

func (this *ChatForm) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("Hel呵呵lo2", nk.NewRect(250, 80, 550, 600-160), nk.WINDOW_BORDER)
		if err != nil {

			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息窗口", 10)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息1", 10)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息2", 10)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息3", 10)

			// draw newest n msgs
			const maxlen = 500
			uniqid := uictx.mdl.Ctuniqid
			hasnew := uictx.mdl.Hasnewmsg(uniqid)
			uictx.mdl.Unsetnewmsg(uniqid)
			msgs := uictx.mdl.GetNewestMsgs(uniqid, maxlen)
			if len(msgs) == 0 {
				// render no message
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

			// 1w条的时候内存倒没有问题，CPU上去了 10+%
			// 3k条以下比较好，滚动的时候使用3%上下的CPU
			for i := 1000; i < 300; i++ {
				tmsg := fmt.Sprintf("聊天消息%d\x00", i)
				ctx.LayoutRowDynamic(30, 1)
				ctx.Label(tmsg, 10)
			}

			emptylen := 410 - float32(len(msgs)+10)*30
			if emptylen > 0 {
				ctx.LayoutRowDynamic(emptylen, 1)
				ctx.Label("空白区域", 10)
			}

			if hasnew {
				// seem no any useful affect
				// ctx.InputScroll(nk.NewVec2(100000, 10000))
				ctx.ForceScroll(100000, 100000) // seem ok
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
	this.iptbuf = make([]byte, 320)
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
			ctx.LayoutRowPush(530 - 80)
			newlen := this.iptblen
			active := ctx.EditString(nk.EDIT_FIELD, this.iptbuf, &newlen, len(this.iptbuf))
			if this.iptblen != newlen {
				this.iptblen = newlen
				log.Println("text", string(this.iptbuf[:newlen]), newlen)
			} else if active == 1 {
			}
			ctx.LayoutRowPush(80)
			if ctx.ButtonLabel("发送按钮") != nil {
				if this.iptblen > 0 {
					cttype := uictx.mdl.Cttype
					ctnum := uictx.mdl.Ctnum
					rptid := uictx.mdl.Nxtreceiptid()
					msg := string(this.iptbuf[:this.iptblen])
					var err error
					switch cttype {
					case thscli.CTTYPE_FRIEND:
						_, err = vtcli.FriendSendMessage(ctnum, msg, rptid)
					case thscli.CTTYPE_GROUP:
						err = vtcli.ConferenceSendMessage(ctnum, 0, msg, rptid)
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
