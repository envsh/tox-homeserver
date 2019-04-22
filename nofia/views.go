package main

import (
	"fmt"
	"log"

	nk "mkuse/nuklear"
)

type MyinfoView struct {
}

func (this *MyinfoView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("我的信息", nk.NewRect(0, 0, 250, 120), nk.WINDOW_BORDER)
		if err != nil {
			ctx.LayoutRowStatic(30, 100, 2)
			ctx.Label("名字", 10)
			ctx.Label("连接状态", 10)
			ctx.LayoutRowStatic(30, 100, 1)
			ctx.Label("状态文本", 10)
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
		err := ctx.Begin("好友状态视图123", nk.NewRect(250, 0, 550, 90), nk.WINDOW_BORDER)
		if err != nil {
			ctx.LayoutRowBegin(nk.STATIC, 30, 3)
			ctx.LayoutRowPush(30)
			ctx.Label("空格", 10)
			ctx.LayoutRowPush(300)
			ctx.Label("联系人名", 10)
			ctx.LayoutRowPush(60)
			if ctx.ButtonLabel("选项") != nil {
			}
			ctx.LayoutRowEnd()

			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("好友状态文本..................................", 10)
		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}

}

type ContectView struct {
}

func (this *ContectView) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		err := ctx.Begin("Hel呵呵lo", nk.NewRect(0, 120, 250, 600-160), nk.WINDOW_BORDER)
		if err != nil {
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
	msgs []string
}

func NewChatForm() *ChatForm {
	this := &ChatForm{}
	this.msgs = []string{"消息111", "消息222",
		"消息333 Unix/Linux shell script FAQ: Can you share a simple Linux shell script that shows how ",
		"消息444", "消息555", "消息666"}

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

			// draw newest n msgs
			const maxlen = 10
			msgs := this.msgs
			if len(this.msgs) > maxlen {
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

			// 1w条的时候内存倒没有问题，CPU上去了 10+%
			// 3k条以下比较好，滚动的时候使用3%上下的CPU
			for i := 1000; i < 3000; i++ {
				tmsg := fmt.Sprintf("聊天消息%d", i)
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
