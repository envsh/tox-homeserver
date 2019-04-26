package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
	"unsafe"

	nk "mkuse/nuklear"

	thscli "tox-homeserver/client"
)

type IconPool struct {
	tstico       *nk.Image
	loadiconOnce sync.Once
}

func (this *IconPool) render() func(ctx *nk.Context) {
	return func(ctx *nk.Context) {
		this.loadiconOnce.Do(func() {
			if true {
				img := nk.XsurfLoadImageFromFile("/home/me/Pictures/2019-04-18_20-31.png")
				this.tstico = img
			} else { // 不行的方式
				bcc, err := ioutil.ReadFile("/home/me/Pictures/2019-04-18_20-31.png")
				gopp.ErrPrint(err)
				this.tstico = nk.NewImage(unsafe.Pointer(&bcc[0]))
			}
		})
	}
}

type MyinfoView struct {
	w *nk.Window2
}

func (this *MyinfoView) loadicons() {
}

func (this *MyinfoView) render() func(ctx *nk.Context) {
	w := &nk.Window2{}
	w.Name = "我的信息窗口"
	w.Rect = nk.NewRect(0, 0, 250, 120)
	w.Flags = nk.WINDOW_BORDER
	this.w = w

	return func(ctx *nk.Context) {
		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {
			name := uictx.mdl.Myname
			name = gopp.IfElseStr(len(name) == 0, " ", name)
			sel0 := len(name)
			sttxt := uictx.mdl.Mysttxt
			ctx.LayoutRowBegin(nk.STATIC, 30, 3)
			ctx.LayoutRowPush(30)
			// ctx.ButtonLabel("我")
			// ctx.ButtonImage(img)
			ctx.ButtonImage(uictx.icop.tstico)
			ctx.LayoutRowPush(160)
			ctx.SelectableLabel(name, 10, &sel0)
			ctx.LayoutRowPush(20)
			ctx.Label(sttxt, 10)
			ctx.LayoutRowEnd()

			stmsg := uictx.mdl.Mystmsg
			stmsg = gopp.IfElseStr(len(stmsg) == 0, " ", stmsg)
			sel1 := len(stmsg)
			ctx.Tooltip(stmsg, 250)
			ctx.LayoutRowDynamic(30, 1)
			ctx.SelectableLabel(stmsg, 10, &sel1)
			ctx.LayoutRowStatic(30, 100, 2)
			ctx.Label("搜索框👉", nk.TEXT_LEFT)
			ctx.Label("排列过滤 ", nk.TEXT_RIGHT)

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}

type MyactionView struct {
	w *nk.Window2
}

func (this *MyactionView) render() func(ctx *nk.Context) {
	w := &nk.Window2{}
	w.Name = "我的控制按钮组"
	w.Rect = nk.NewRect(0, 600-50, 250, 60)
	w.Flags = nk.WINDOW_BORDER
	this.w = w

	return func(ctx *nk.Context) {
		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {

			ctx.LayoutRowBegin(nk.STATIC, 30, 4)
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("加好友") != nil {
			}
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("加群组") != nil {
			}
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("文件") != nil {
			}
			ctx.LayoutRowPush(50)
			if ctx.ButtonLabel("设置") != nil {
				uictx.setfmv.show = true
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
	w *nk.Window2
}

func (this *FriendInfoView) render() func(*nk.Context) {
	w := &nk.Window2{}
	w.Name = "当前群/友状态视图"
	w.Rect = nk.NewRect(250, 0, 550, 85)
	w.Flags = nk.WINDOW_NO_SCROLLBR
	this.w = w

	return func(ctx *nk.Context) {
		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {
			ctx.LayoutRowBegin(nk.STATIC, 30, 4)
			ctx.LayoutRowPush(40)
			ctx.ButtonLabel(gopp.IfElseStr(uictx.mdl.Cttype == thscli.CTTYPE_FRIEND, "好友", "群组"))

			name := uictx.mdl.Ctname
			name = gopp.IfElseStr(len(name) == 0, " ", name)
			sel0 := len(name)
			ctx.LayoutRowPush(320)
			ctx.SelectableLabel(name, 5, &sel0)

			ctx.LayoutRowPush(80)
			if ctx.ButtonLabel("成员列表") != nil {
				uictx.mblstv.show = true
			}
			ctx.LayoutRowPush(80)
			totheight := float32(3 * 40)
			if ctx.MenuBeginLabel("  选项  ", nk.TEXT_RIGHT, nk.NewVec2(120, totheight)) != nil {
				ctx.LayoutRowDynamic(30, 1)
				if ctx.MenuItemLabel("Mute", nk.TEXT_LEFT) != nil {
					log.Println("Mute")
				}
				ctx.LayoutRowDynamic(30, 1)
				if ctx.MenuItemLabel("hehe2", nk.TEXT_LEFT) != nil {
					log.Println("action2")
				}
				ctx.LayoutRowDynamic(30, 1)
				if ctx.MenuItemLabel("hehe3", nk.TEXT_LEFT) != nil {
					log.Println("action3")
				}

				ctx.MenuEnd()
			}
			ctx.LayoutRowEnd()

			// for group, it is always empty, so we borrow it to show lastmsg
			stmsg := gopp.IfElseStr(uictx.mdl.Cttype != thscli.CTTYPE_FRIEND,
				uictx.mdl.Lastmsg(), uictx.mdl.Ctstmsg)
			stmsg = gopp.IfElseStr(len(stmsg) == 0, " ", stmsg)
			sel1 := len(stmsg)
			rstmsg := ""
			if pos := strings.Index(stmsg, ": "); pos != -1 {
				rstmsg = stmsg[pos+2:]
			}
			ctx.LayoutRowBegin(nk.STATIC, 30, 2)
			if len(rstmsg) > 12 {
				ctx.Tooltip(rstmsg, 500-100)
			}
			ctx.LayoutRowPush(500 - 100)
			ctx.SelectableLabel(stmsg, nk.TEXT_LEFT, &sel1)

			ctx.Tooltip("当前/总数", 130) // this is tooltip of next widget, here is below label
			curcnt, totcnt := uictx.mdl.TotalCurrMsgcount()
			labtxt := fmt.Sprintf("消息数：%d/%d", curcnt, totcnt)
			ctx.LayoutRowPush(130)
			ctx.Label(labtxt, nk.TEXT_RIGHT)
			ctx.LayoutRowEnd()
		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}

}

type ContectView struct {
	w *nk.Window2
}

func NewcontactView() *ContectView {
	this := &ContectView{}
	return this
}

func (this *ContectView) render() func(ctx *nk.Context) {
	w := &nk.Window2{}
	w.Name = "联系人列表"
	w.Rect = nk.NewRect(0, 120, 250, 600-160)
	w.Flags = nk.WINDOW_BORDER | nk.WINDOW_SCROLL_AUTO_HIDE
	this.w = w

	return func(ctx *nk.Context) {
		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {
			for _, v := range uictx.mdl.GroupList() {
				name := fmt.Sprintf("%s・%d", v.GetTitle(), len(v.GetMembers()))
				statxt := fmt.Sprintf("%d", uictx.mdl.NewMsgcount(v.GetGroupId()))
				tiptxt := fmt.Sprintf("未读=%d, 所有=%d",
					uictx.mdl.NewMsgcount(v.GetGroupId()), uictx.mdl.Msgcount(v.GetGroupId()))
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("群")
				ctx.LayoutRowPush(150)
				if ctx.ButtonLabel(name) != nil {
					log.Println("group clicked", v.GetGnum(), name)
					uictx.mdl.Switchtoct(v.GetGroupId())
				}
				ctx.Tooltip(tiptxt, 30)
				ctx.LayoutRowPush(30)
				ctx.Label(statxt, nk.TEXT_CENTERED)
				ctx.LayoutRowEnd()
			}
			for _, v := range uictx.mdl.FriendList() {
				name := fmt.Sprintf("%s", v.GetName())
				statxt := fmt.Sprintf("%s %d",
					thscli.Conno2str1(int(v.Status)), uictx.mdl.NewMsgcount(v.GetPubkey()))
				tiptxt := fmt.Sprintf("%s %s, 未读=%d, 所有=%d",
					thscli.Conno2str(int(v.Status)), gopp.IfElseStr(v.Status == 0, "离线", "在线"),
					uictx.mdl.NewMsgcount(v.GetPubkey()), uictx.mdl.Msgcount(v.GetPubkey()))
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(30)
				ctx.ButtonLabel("友")
				ctx.LayoutRowPush(150)
				if ctx.ButtonLabel(name) != nil {
					log.Println("friend clicked", v.GetFnum(), name)
					uictx.mdl.Switchtoct(v.GetPubkey())
				}
				ctx.Tooltip(tiptxt, 30)
				ctx.LayoutRowPush(30)
				ctx.Label(statxt, nk.TEXT_LEFT)
				ctx.LayoutRowEnd()
			}

			for i := 0; i < 16; i++ {
				name := fmt.Sprintf("好友名%d", i+1)
				statxt := fmt.Sprintf("联系人%d", i+1)
				ctx.LayoutRowStatic(30, 100, 2)
				ctx.Label(name, nk.TEXT_LEFT)
				ctx.Label(statxt, nk.TEXT_LEFT)
			}

			ctx.LayoutRowDynamic(510-3*30, 1)
			ctx.Label("空白区域", nk.TEXT_CENTERED)

		}
		ctx.End()
		if ctx.WindowIsHidden("Hello") {
			return
		}
	}
}

/////
type ChatForm struct {
	w *nk.Window2
}

func NewChatForm() *ChatForm {
	this := &ChatForm{}
	return this
}

func (this *ChatForm) render() func(ctx *nk.Context) {
	w := &nk.Window2{}
	w.Name = "聊天消息列表窗口"
	w.Rect = nk.NewRect(250, 80, 550, 600-160)
	w.Flags = nk.WINDOW_BORDER
	this.w = w

	return func(ctx *nk.Context) {
		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {

			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息窗口", nk.TEXT_CENTERED)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息1", nk.TEXT_CENTERED)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息2", nk.TEXT_CENTERED)
			ctx.LayoutRowDynamic(30, 1)
			ctx.Label("聊天消息3", nk.TEXT_CENTERED)

			// draw newest n msgs
			const maxlen = 500
			uniqid := uictx.mdl.Ctuniqid
			hasnew := uictx.mdl.Hasnewmsg(uniqid)
			uictx.mdl.Unsetnewmsg(uniqid)
			msgs := uictx.mdl.GetNewestMsgs(uniqid, maxlen)
			if len(msgs) == 0 {
				// render no message
			}

			for oidx, msg := range msgs {
				ctx.LayoutRowBegin(nk.STATIC, 39, 6)
				ctx.LayoutRowPush(30)
				if !msg.Me {
					ctx.ButtonLabel(" ")
				} else {
					ctx.Label(" ", nk.TEXT_CENTERED)
				}
				name := gopp.IfElseStr(msg.Me, uictx.mdl.Myname, msg.PeerNameUi)
				ctx.LayoutRowPush(300)
				if len(msg.Links) > 0 {
					if ctx.MenuBeginLabel(name, nk.TEXT_LEFT, nk.NewVec2(450, 150)) != nil {
						for _, link := range msg.Links {
							ctx.LayoutRowDynamic(30, 1)
							if ctx.MenuItemLabel(link, nk.TEXT_LEFT) != nil {
								uictx.XdgOpen(link)
							}
						}
						ctx.MenuEnd()
					}
				} else {
					ctx.Label(name, nk.TEXT_LEFT)
				}
				ctx.LayoutRowPush(80)
				ctx.Label(msg.TimeUi, nk.TEXT_RIGHT)
				ctx.LayoutRowPush(30)
				ctx.Label(gopp.IfElseStr(msg.Sent, " ", "=>"), nk.TEXT_RIGHT)
				ctx.LayoutRowPush(30)
				if ctx.ButtonLabel("Copy") != nil {
					uictx.Copy2Clipboard(msg.MsgUi)
				}
				ctx.LayoutRowPush(30)
				if msg.Me {
					ctx.ButtonLabel("Me")
				} else {
					ctx.Label(" ", nk.TEXT_CENTERED)
				}
				ctx.LayoutRowEnd()

				wraped := gopp.Splitrnui(msg.MsgUi, 60)
				for idx, line := range wraped {
					ctx.LayoutRowBegin(nk.STATIC, 39, 3)
					ctx.LayoutRowPush(30)
					ctx.Label(gopp.IfElseStr(idx > 0, " ", fmt.Sprintf("%d", oidx)), nk.TEXT_CENTERED)
					ctx.LayoutRowPush(450)
					seln := len(line)
					ctx.SelectableLabel(line, nk.TEXT_LEFT, &seln)
					ctx.LayoutRowPush(30)
					ctx.Label(" ", nk.TEXT_CENTERED)
					ctx.LayoutRowEnd()
				}
			}

			// 1w条的时候内存倒没有问题，CPU上去了 10+%
			// 3k条以下比较好，滚动的时候使用3%上下的CPU
			for i := 1000; i < 300; i++ {
				tmsg := fmt.Sprintf("聊天消息%d\x00", i)
				ctx.LayoutRowDynamic(30, 1)
				ctx.Label(tmsg, nk.TEXT_CENTERED)
			}

			emptylen := 410 - float32(len(msgs)+10)*30
			if emptylen > 0 {
				ctx.LayoutRowDynamic(emptylen, 1)
				ctx.Label("空白区域", nk.TEXT_CENTERED)
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
	w *nk.Window2

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
	w := &nk.Window2{}
	w.Name = "消息输入发送视图"
	w.Rect = nk.NewRect(250, 520, 550, 600-510)
	w.Flags = nk.WINDOW_BORDER
	this.w = w

	return func(ctx *nk.Context) {
		err := ctx.Begin(w.Name, w.Rect, w.Flags)
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
			ctx.LayoutRowPush(520 - 80)
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
					uniqid := uictx.mdl.Ctuniqid
					rptid := thscli.NextUserCode(devInfo.Uuid)
					msg := string(this.iptbuf[:this.iptblen])
					var err error
					switch cttype {
					case thscli.CTTYPE_FRIEND:
						msgo := thscli.NewMessageForMe(msg)
						msgo.UserCode = rptid
						uictx.mdl.Newmsg(uniqid, msgo)
						_, err = vtcli.FriendSendMessage(ctnum, msg, rptid)
					case thscli.CTTYPE_GROUP:
						msgo := thscli.NewMessageForMe(msg)
						msgo.UserCode = rptid
						uictx.mdl.Newmsg(uniqid, msgo)
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

type MemberListForm struct {
	w *nk.Window2

	show bool
}

func NewMemberListForm() *MemberListForm {
	this := &MemberListForm{}
	return this
}

func (this *MemberListForm) render() func(ctx *nk.Context) {
	w := &nk.Window2{}
	w.Name = "群成员列表窗口"
	w.Rect = nk.NewRect(260, 60, 500, 600-100)
	w.Flags = nk.WINDOW_BORDER | nk.WINDOW_CLOSABLE | nk.WINDOW_MOVABLE | nk.WINDOW_TITLE
	this.w = w

	return func(ctx *nk.Context) {
		if !this.show {
			return
		}

		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {
			/*
				err := ctx.PopupBegin(nk.POPUP_STATIC, w.Name, w.Flags, w.Rect)
				if err != nil {
					ctx.PopupEnd()
				}
			*/

			mbs := uictx.mdl.CurMembers()
			ctx.LayoutRowDynamic(30, 2)
			ctx.ButtonLabel(uictx.mdl.Ctname)
			ctx.ButtonLabel(fmt.Sprintf("%d 个成员", len(mbs)))

			for _, v := range mbs {
				ctx.LayoutRowBegin(nk.STATIC, 30, 3)
				ctx.LayoutRowPush(160)
				ctx.Label(v.GetName(), nk.TEXT_CENTERED)
				ctx.LayoutRowPush(160)
				ctx.Label(v.GetPubkey(), nk.TEXT_CENTERED)
				ctx.LayoutRowPush(130)
				ctx.Label(gopp.TimeToFmt1(time.Unix(v.GetJoints(), 0)), nk.TEXT_CENTERED)
				ctx.LayoutRowEnd()
			}
		}
		ctx.End()
		this.show = !ctx.WindowIsClosed(w.Name) && !ctx.WindowIsHidden(w.Name)
	}
}

type SettingForm struct {
	w *nk.Window2

	show bool
}

func NewSettingForm() *SettingForm {
	this := &SettingForm{}
	return this
}

func (this *SettingForm) render() func(ctx *nk.Context) {
	w := &nk.Window2{}
	w.Name = "设置窗口"
	w.Rect = nk.NewRect(220, 60, 500, 600-100)
	w.Flags = nk.WINDOW_BORDER | nk.WINDOW_CLOSABLE | nk.WINDOW_MOVABLE | nk.WINDOW_TITLE
	this.w = w

	setidx := 0
	return func(ctx *nk.Context) {
		if !this.show {
			return
		}

		err := ctx.Begin(w.Name, w.Rect, w.Flags)
		if err != nil {
			/*
				err := ctx.PopupBegin(nk.POPUP_STATIC, w.Name, w.Flags, w.Rect)
				if err != nil {
					ctx.PopupEnd()
				}
			*/

			ctx.LayoutRowDynamic(30, 5)
			if ctx.ButtonLabel("基本设置") != nil {
				setidx = 0
			}
			if ctx.ButtonLabel("设置2") != nil {
				setidx = 1
			}
			if ctx.ButtonLabel("高级") != nil {
				setidx = 2
			}
			if ctx.ButtonLabel("关于") != nil {
				setidx = 3
			}
			if ctx.ButtonLabel("开发") != nil {
				setidx = 4
			}

			if setidx < 0 || setidx > 4 {
			} else if setidx == 0 {
				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("界面风格", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				sel0 := "hehehehe"
				if ctx.ComboBeginLabel(sel0, nk.NewVec2(150, 150)) != nil {
					ctx.LayoutRowDynamic(30, 1)
					ctx.Label("默认", nk.TEXT_LEFT)
					ctx.Label("黑色", nk.TEXT_LEFT)
					ctx.Label("白色", nk.TEXT_LEFT)
					ctx.Label("蓝色", nk.TEXT_LEFT)
					ctx.Label("红色", nk.TEXT_LEFT)
					ctx.ComboEnd()
				}
				ctx.LayoutRowEnd()

				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("Use HS", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				actived1 := 0
				ctx.CheckboxLabel("cb1", &actived1)
				ctx.LayoutRowEnd()

				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("开启皮肤", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				actived2 := 0
				ctx.CheckboxLabel("cb2", &actived2)
				ctx.LayoutRowEnd()

				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("ToxHS地址", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				sel1 := "hehehehe"
				if ctx.ComboBeginLabel(sel1, nk.NewVec2(150, 150)) != nil {
					ctx.LayoutRowDynamic(30, 1)
					ctx.Label("txhs.duckdns.org", nk.TEXT_LEFT)
					ctx.Label("10.0.0.31", nk.TEXT_LEFT)
					ctx.Label("127.0.0.1", nk.TEXT_LEFT)
					ctx.ComboEnd()
				}
				ctx.LayoutRowEnd()

				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("字体名称", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				sel2 := "hehehehe"
				if ctx.ComboBeginLabel(sel2, nk.NewVec2(150, 150)) != nil {
					ctx.LayoutRowDynamic(30, 1)
					ctx.Label("font1", nk.TEXT_LEFT)
					ctx.Label("font2", nk.TEXT_LEFT)
					ctx.Label("font3", nk.TEXT_LEFT)
					ctx.Label("font4", nk.TEXT_LEFT)
					ctx.Label("font5", nk.TEXT_LEFT)
					ctx.ComboEnd()
				}
				ctx.LayoutRowEnd()

				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("字体大小", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				ftsz := ctx.SliderInt(5, 14, 50, 1)
				_ = ftsz
				ctx.LayoutRowEnd()
			} else if setidx == 1 {
				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowEnd()

			} else if setidx == 2 {
				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowEnd()

			} else if setidx == 3 {
				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowEnd()
			} else if setidx == 4 {
				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("日志级别", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				sel2 := "hehehehe"
				if ctx.ComboBeginLabel(sel2, nk.NewVec2(150, 250)) != nil {
					ctx.LayoutRowDynamic(30, 1)
					ctx.Label("TRACE", nk.TEXT_LEFT)
					ctx.Label("DEBUG", nk.TEXT_LEFT)
					ctx.Label("INFO", nk.TEXT_LEFT)
					ctx.Label("WARNING", nk.TEXT_LEFT)
					ctx.Label("ERROR", nk.TEXT_LEFT)
					ctx.Label("FATAL", nk.TEXT_LEFT)
					ctx.ComboEnd()
				}
				ctx.LayoutRowEnd()

				ctx.LayoutRowBegin(nk.STATIC, 30, 2)
				ctx.LayoutRowPush(100)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowPush(200)
				ctx.Label("PlaceHolder3", nk.TEXT_LEFT)
				ctx.LayoutRowEnd()
			}
		}
		ctx.End()
		this.show = !ctx.WindowIsClosed(w.Name) && !ctx.WindowIsHidden(w.Name)
	}
}
