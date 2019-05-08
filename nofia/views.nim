

proc nkwndhello0(nkw:PNkwindow, name : string) {.gcsafe.} =
    var cname : cstring = "hehehhehehee"
    if nk_begin(nkw.nkctx, cname, nk_rect(x:0, y:0, w:300, h:200),
                NK_WINDOW_TITLE or NK_WINDOW_BORDER or
                NK_WINDOW_MOVABLE or NK_WINDOW_SCALABLE) == ctrue:
        nk_layout_row_static(nkw.nkctx, 30, 200, 1)
        nk_label(nkw.nkctx, "hehehh呵呵ehe", 1)
        discard
    discard nk_end(nkw.nkctx)
    return

proc nkwndhello1(nkw:PNkwindow, name : string) {.gcsafe.} =
    return

proc nkwndhello2(nkw:PNkwindow, name : string) {.gcsafe.} =
    return

proc MyInfoView(nkw:PNkwindow, name:string) {.gcsafe.} =
    var ctx = nkw.nkctx
    var mdl = nkw.mdl

    if ctx.nk_begin(name, nk_rect(x:0,y:0,w:250,h:120), NK_WINDOW_BORDER) == ctrue:
        ctx.nk_layout_row_begin(NK_STATIC, 30, 3)
        ctx.nk_layout_row_push(30)
        discard ctx.nk_button_label("ICO")
        ctx.nk_layout_row_push(160)
        let sel0 : int = 0
        var MyName : cstring = mdl.MyName
        if mdl.MyName == "": MyName = " "
        discard ctx.nk_selectable_label(MyName, 10, sel0.unsafeAddr)
        ctx.nk_layout_row_push(20)
        ctx.nk_label("heheh", 10)
        ctx.nk_layout_row_end()

        ctx.nk_tooltip("ehehe", )
        ctx.nk_layout_row_dynamic(30, 1)
        var Mystmsg : cstring = mdl.MyStmsg
        if mdl.MyStmsg == "": Mystmsg = " "
        discard ctx.nk_selectable_label(MyStmsg, 10, sel0.unsafeAddr)
        ctx.nk_layout_row_static(30, 100, 2)
        ctx.nk_label("搜索框👉", NK_TEXT_LEFT)
        ctx.nk_label("排列过滤 ", NK_TEXT_RIGHT)
        discard
    discard ctx.nk_end()
    return

proc MyActionView(nkw:PNkwindow, name:string) {.gcsafe.} =
    var ctx = nkw.nkctx
    if ctx.nk_begin(name, nk_rect(x:0, y:600-50, w:250, h:60), NK_WINDOW_BORDER) == ctrue:
        ctx.nk_layout_row_begin(NK_STATIC, 30, 4)
        ctx.nk_layout_row_push(50)
        if ctx.nk_button_label("加友") == ctrue: discard
        ctx.nk_layout_row_push(50)
        if ctx.nk_button_label("加群") == ctrue: discard
        ctx.nk_layout_row_push(50)
        if ctx.nk_button_label("文件") == ctrue: discard
        ctx.nk_layout_row_push(50)
        if ctx.nk_button_label("设置") == ctrue: discard
        ctx.nk_layout_row_end()
        discard
    discard ctx.nk_end()
    return

proc FriendInfoView(nkw:PNkwindow, name:string) {.gcsafe.} =
    var ctx = nkw.nkctx
    var mdl = nkw.mdl

    if ctx.nk_begin(name, nk_rect(x:250, y:0, w:550, h:85), NK_WINDOW_NO_SCROLLBAR) == ctrue:
        ctx.nk_layout_row_begin(NK_STATIC, 30, 4)
        ctx.nk_layout_row_push(40)
        discard ctx.nk_button_label("友群")

        var ctname = mdl.Ctname
        if ctname == "": ctname = " "
        var sel0 = 0
        ctx.nk_layout_row_push(320)
        discard ctx.nk_selectable_label(ctname, 5, sel0.unsafeAddr)

        ctx.nk_layout_row_push(80)
        if ctx.nk_button_label("成员列表") == ctrue: discard

        ctx.nk_layout_row_push(80)
        if ctx.nk_menu_begin_label(" 选项 ", NK_TEXT_RIGHT, nk_vec2(x:120, y: 150))==ctrue:
            ctx.nk_layout_row_dynamic(30, 1)
            if ctx.nk_menu_item_label("Mute", NK_TEXT_LEFT) == ctrue:
                linfo("Mute")
            ctx.nk_layout_row_dynamic(30, 1)
            if ctx.nk_menu_item_label("hehe2", NK_TEXT_LEFT) == ctrue:
                linfo("action2")
            ctx.nk_layout_row_dynamic(30, 1)
            if ctx.nk_menu_item_label("hehe3", NK_TEXT_LEFT) == ctrue:
                linfo("action3")
            ctx.nk_menu_end()
        ctx.nk_layout_row_end()

        ctx.nk_layout_row_begin(NK_STATIC, 30, 2)
        ctx.nk_layout_row_push(500-100)
        var stmsg = if mdl.Cttype != CTTYPE_FRIEND: mdl.Lastmsg() else: mdl.Ctstmsg
        if stmsg == "": stmsg = " "
        discard ctx.nk_selectable_label(stmsg, NK_TEXT_LEFT, sel0.unsafeAddr)

        ctx.nk_tooltip("当前/总数",)
        ctx.nk_layout_row_push(130)
        ctx.nk_label("hoiifddf", NK_TEXT_RIGHT)
        ctx.nk_layout_row_end()
        discard
    discard ctx.nk_end()

proc ContactView(nkw:PNkwindow, name:string) {.gcsafe.} =
    var ctx = nkw.nkctx
    var mdl = nkw.mdl

    if ctx.nk_begin(name, nk_rect(x:0, y:120, w:250, h:600-160),
                    NK_WINDOW_BORDER or NK_WINDOW_NO_SCROLLBAR) == ctrue:
        for v in mdl.GroupList():
            var name = v.Title
            var tiptxt = "hehe"
            ctx.nk_layout_row_begin(NK_STATIC, 30, 3)
            ctx.nk_layout_row_push(30)
            discard ctx.nk_button_label("群")
            ctx.nk_layout_row_push(150)
            if ctx.nk_button_label(name) == ctrue:
                linfo("group clicked", name)
                mdl.Switchtoct(v.GroupId)
            ctx.nk_layout_row_push(30)
            ctx.nk_label("hehhe", NK_TEXT_CENTERED)
            ctx.nk_layout_row_end()

        for v in mdl.FriendList():
            var name = v.Name
            var statxt = "hehhe"
            var tiptxt = "hehehe"
            ctx.nk_layout_row_begin(NK_STATIC, 30, 3)
            ctx.nk_layout_row_push(30)
            discard ctx.nk_button_label("友")
            ctx.nk_layout_row_push(150)
            if ctx.nk_button_label(name) == ctrue:
                linfo("friend clicked", name)
                mdl.Switchtoct(v.Pubkey)
            ctx.nk_layout_row_push(30)
            ctx.nk_label(statxt, NK_TEXT_LEFT)
            ctx.nk_layout_row_end()

        ctx.nk_layout_row_dynamic(510-3*30, 1)
        ctx.nk_label("空白区域", NK_TEXT_CENTERED)
        discard
    discard ctx.nk_end()

proc ChatForm(nkw:PNkwindow, name:string) {.gcsafe.} =
    var ctx = nkw.nkctx
    var mdl = nkw.mdl

    if ctx.nk_begin(name, nk_rect(x:250, y:80, w:550, h:600-160), NK_WINDOW_BORDER) == ctrue:
        ctx.nk_layout_row_dynamic(30, 1)
        ctx.nk_label("聊天消息窗口", NK_TEXT_CENTERED)
        ctx.nk_layout_row_dynamic(30, 1)
        ctx.nk_label("聊天消息1", NK_TEXT_CENTERED)
        ctx.nk_layout_row_dynamic(30, 1)
        ctx.nk_label("聊天消息2", NK_TEXT_CENTERED)
        ctx.nk_layout_row_dynamic(30, 1)
        ctx.nk_label("聊天消息3", NK_TEXT_CENTERED)

        # // draw newest n msgs
        const maxlen = 500
        var uniqid = mdl.Ctuniqid
        var hasnew = mdl.Hasnewmsg(uniqid)
        mdl.Unsetnewmsg(uniqid)
        var msgs = mdl.GetNewestMsgs(uniqid, maxlen)
        if msgs.len() == 0: # render no message
            discard

        for oidx, msg in msgs:
            ctx.nk_layout_row_begin(NK_STATIC, 39, 6)
            ctx.nk_layout_row_push(30)
            if msg.Me: discard ctx.nk_button_label(" ")
            else: ctx.nk_label(" ", NK_TEXT_CENTERED)
            var name = if msg.Me: mdl.Myname else: msg.PeerNameUi
            ctx.nk_layout_row_push(300)
            ctx.nk_label(name, NK_TEXT_LEFT)

            ctx.nk_layout_row_push(80)
            ctx.nk_label(msg.TimeUi, NK_TEXT_LEFT)
            ctx.nk_layout_row_push(30)
            ctx.nk_label(if msg.Sent: " " else: "=>", NK_TEXT_LEFT)
            ctx.nk_layout_row_push(30)
            if ctx.nk_button_label("Copy") == ctrue:
                linfo("Copy action")
            ctx.nk_layout_row_push(30)
            if msg.Me: discard ctx.nk_button_label("Me")
            else: ctx.nk_label(" ", NK_TEXT_CENTERED)
            ctx.nk_layout_row_end()

            continue

        ctx.nk_layout_row_dynamic(510-3*30, 1)
        ctx.nk_label("空白区域", NK_TEXT_CENTERED)
        discard
    discard ctx.nk_end()

proc SendForm(nkw:PNkwindow, name:string) {.gcsafe.} =
    var ctx = nkw.nkctx
    if ctx.nk_begin(name, nk_rect(x:250, y:520, w:550, h:600-510), NK_WINDOW_BORDER) == ctrue:
        ctx.nk_layout_row_begin(NK_STATIC, 30, 7)
        for i in 0..5:
            ctx.nk_layout_row_push(50)
            if ctx.nk_button_label("操作"&repr(i)) == ctrue:
                discard
        ctx.nk_layout_row_push(50)
        if ctx.nk_button_label("粘贴") == ctrue:
            discard
        ctx.nk_layout_row_end()

        ctx.nk_layout_row_begin(NK_STATIC, 30, 2)
        ctx.nk_layout_row_push(520-80)
        var newlen : int = 5
        var active = ctx.nk_edit_string(NK_EDIT_FIELD, "jifeiwajfwf", newlen.unsafeAddr, 6, nil)
        ctx.nk_layout_row_push(80)
        if ctx.nk_button_label("发送按钮") == ctrue: discard
        ctx.nk_layout_row_end()
        discard
    discard ctx.nk_end()

#[
]#


proc nkwndrunproc(nkw:PNkwindow) {.gcsafe.} =
    for name, wp in nkw.wnds.pairs: wp(nkw, name)
    return

proc createnkwndprocs(nkw:PNkwindow) =
    nkw.wnds.add("hello0",nkwndhello0)
    nkw.wnds.add("hello1",nkwndhello1)
    nkw.wnds.add("hello2",nkwndhello1)
    nkw.wnds.add("MyInfoView",MyInfoView)
    nkw.wnds.add("MyActionView",MyActionView)
    nkw.wnds.add("FriendInfoView",FriendInfoView)
    nkw.wnds.add("ContactView",ContactView)
    nkw.wnds.add("ChatForm",ChatForm)
    nkw.wnds.add("SendForm",SendForm)
    nkw.wndrunner = nkwndrunproc
    return
