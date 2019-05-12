{.passc:"-fPIC -g -O0 -xc -DRENDER_X11_NATIVE"}
{.passl:"-lX11 -lXft -lXrender"}

{.compile: "render_x11_native.c.ngo".}
import x11/x, x11/xlib
include render_x11_native

{.passc:"-I/usr/include/freetype2"}
{.compile:"nuklear_x11_all.c.ngo".}
include nuklear_x11_all

# Nimenv depend some types, put before it
import asyncdispatch
import asyncfutures

include "nimlog.nim"
include "nimplus.nim"
include "datatypes.nim"
include "uimodels.nim"

type MixEvent = ref object
    typ*: int
    evts*: seq[TXEvent]

import tables
type
    PNkwindow = ptr Nkwindow
    Nkwindow = ref object
        rdwin*: TRenderWindow
        xfont*: XFont
        nkctx*: nk_context
        dpyfd*: AsyncFD
        evtch*: Channel[MixEvent]
        wnds*: Table[string, proc (nkw:PNkwindow, name:string) {.gcsafe.}]
        wndrunner*: proc (nkw:PNkwindow) {.gcsafe.}
        mdl*: DataModel
        sdipt*: nk_editor_state
        ximst*: xim_state

# like global vars, but put in struct
type
    PNimenv = ptr Nimenv
    Nimenv = ref object
        self*: Nimenv # 为了能够用指针取到引用
        pself*: PNimenv
        argc*: int
        argv*: string
        stoped*: bool
        nkxwin*: Nkwindow
        asyevts*: Table[string, proc (pne: PNimenv)]
        rpcli*: RpcClient


#include "frontend_proc.nim"
include "nimenv.nim"
include "uimessage.nim"
include "client.nim"
include "views.nim"

proc docjkinput(nkw:PNkwindow, evt : PXEvent) =
    var ximst = nkw.ximst
    var ksym : TKeySym
    var status: int

    ximst.gotlen = XmbLookupStringFixed(nkw.rdwin.ic, cast[PXKeyPressedEvent](evt),
                                        ximst.ximbuf.addr, ximst.ximbuf.len.cint, ksym.addr,
                                        cast[PStatus](status.addr))
    if ximst.gotlen > 0:
        var cjkiptxt = ximst.ximstr()
        linfo("xim inputed", ximst.gotlen, cjkiptxt.len, cjkiptxt)
        for r in cjkiptxt.runes: nkw.nkctx.nk_input_unicode(r)
    return

proc dorepaint(nkw:PNkwindow, evts : seq[TXEvent]) =
    var rdwin = nkw.rdwin
    nk_input_begin(nkw.nkctx)
    for evt1 in evts:
        var evt = evt1
        var evtp = evt.addr
        var kevtp = cast[PXKeyEvent](evtp)
        if evtp.theType == KeyPress and kevtp.keycode == 0:
            linfo("not your(nk) food", evtp.theType, kevtp.keycode)
            docjkinput(nkw, evtp)
            continue
        nk_xlib_handle_event(rdwin.dpy, rdwin.screen, rdwin.win, evt.addr)
    nk_input_end(nkw.nkctx)

    # GUI

    # APP widgets here
    if nkw.wndrunner != nil: nkw.wndrunner(nkw)

    # Draw
    discard XClearWindow(rdwin.dpy, rdwin.win)
    nk_xlib_render(rdwin.win, nk_color(r:30, g:30, b:30))
    discard XFlush(rdwin.dpy)
    return

proc x11proc(nkw:PNkwindow) =
    var stop = false
    while not stop:
        var evt : TXEvent
        var evts = newSeq[TXEvent]()
        var ret1 = XPending(nkw.rdwin.dpy)
        # ldebug("pending", ret1)
        while ret1 > 0:
            ret1 = ret1 - 1
            let ret = XNextEvent(nkw.rdwin.dpy, addr(evt))
            if XFilterEvent(evt.addr, None) == ctrue: continue # 为啥没这行会调不出输入法
            # ldebug("nxtevt:", ret, "evty:", evt.theType, "pending left", ret1, cast[pointer](nkw.rdwin.dpy))
            if evt.theType == ClientMessage:
                lerror("some error occurs", evt.theType)
                # stop = true # TODO
            elif ret == 0: evts.add(evt)
        # ldebug("evtcnt", evts.len())
        if evts.len() > 0:
            # ldebug("evts sending...", evts.len(), nkw.evtch.ready())
            var mevt = MixEvent(typ:1, evts:evts)
            var ok = nkw.evtch.trysend(mevt)
            #var ok = nkw.evtch.trysend("hehhe" & repr(evts.len()))
            # ldebug("evts sent", ok, evts.len(), nkw.evtch.ready(), nkw.evtch.peek())
            # sync()
        else: sleep(150)

    linfo("x11proc done")
    return

proc eventloop(nkw:PNkwindow) =
    linfo("eventlooping...")
    while true:
        # linfo("evtch recving...")
        var mxevt = nkw.evtch.recv()
        # ldebug("goty evty:", mxevt.typ, mxevt.evts.len(), nkw.evtch.peek())
        if mxevt.typ == 0: discard
        elif mxevt.typ == 1: dorepaint(nkw, mxevt.evts)
        elif mxevt.typ == 2: discard
    return

proc newNkwindow(nep:pointer): pointer {.exportc.} =
    var ne = cast[PNimenv](nep)
    var nkw = new(Nkwindow)
    nkw.wnds = initTable[string, proc (nkw:PNkwindow, name:string)]()
    nkw.mdl = newDataModel()
    nkw.sdipt = neweditstate()
    nkw.ximst = newximstate()
    createnkwndprocs(nkw.addr)

    linfo("chan is nil", cast[pointer](addr(nkw.evtch)))
    ldebug("chan is ready", nkw.evtch.ready())
    nkw.evtch.open(8)
    ldebug("chan is ready", nkw.evtch.ready())
    ne.nkxwin = nkw
    return cast[pointer](addr(ne.nkxwin))

proc NkwindowOpen(nep:pointer) {.exportc.} =
    var ne = cast[PNimenv](nep)
    var nkw = ne.nkxwin
    var pnkw = addr(ne.nkxwin)

    # 启动顺序，事件接收线程， 打开XDisplay, nk xlib初始化， 启动 x11事件接收线程
    spawn eventloop(pnkw) # shoule before NewRenderWindow, or recv nothing
    nkw.rdwin = NewRenderWindow()
    linfo("rdwin is nil", nkw.rdwin == nil, nkw.rdwin)
    var dpyfd = ConnectionNumber(nkw.rdwin.dpy)
    linfo("dpyfd:", dpyfd)
    nkw.dpyfd = cast[AsyncFD](dpyfd)

    var xft = nk_xfont_create(nkw.rdwin.dpy, "*")
    var rdwin = nkw.rdwin
    var cmap = rdwin.cmap # 这种方式得到的值不对，可能是结构体在nim与c中不匹配。确认，XIM,XIC大小不对
    # cmap = RenderWindowCmap(rdwin)
    var vis = rdwin.vis
    # vis = RenderWindowVis(rdwin)
    var nkctx = nk_xlib_init(xft, rdwin.dpy, rdwin.screen, rdwin.win, vis, cmap, 800, 600)
    nkw.xfont = xft
    nkw.nkctx = nkctx

    spawn x11proc(pnkw)
    return


### main section
discard XInitThreads()
var ne1 = newNimenvImpl()
var ne2 = newNimenvImpl()
linfo("singleton? ", ne1 == ne2)

spawn stopafter3s(ne1)

discard newNkwindow(getNimenvp())
NkwindowOpen(getNimenvp())
# nk_x11_event_handle(ne1.nkxwin.rdwin)

# runNimenv(ne1.self)
var cli = newclient()
cli.connect()
ne1.rpcli = cli

### tryout eventloop
proc onchkx11evt(pne:PNimenv) =
    linfo("x11 check event timedout")
    return
proc onrtx11evt(fd: AsyncFD) : bool =
    linfo("x11 has rt event")
    return false
proc chkx11toh(fd: AsyncFD): bool =
    onchkx11evt(ne1)
    return false

proc initAsyevtTable() =
    var ne = ne1
    #ne.asyevts.add($ne.nkxwin.dpyfd, onrtx11evt)
    # ne.asyevts.add("chkx11evt", onchkx11evt)
    addTimer(50000, false, chkx11toh)
    # too much, disable
    #register(ne.nkxwin.dpyfd)
    #addRead(ne.nkxwin.dpyfd, onrtx11evt)
    register(cli.subrfd)
    addRead(cli.subrfd, onsubskread)
    register(cli.reqrfd)
    addRead(cli.reqrfd, onreqskread)
    return


initAsyevtTable()
# cli.reqsksend("hehehhee")
cli.getBaseInfo()

# type CallBack = proc (fd: AsyncFD) : bool

while true:
    poll(300000)
    #linfo("poll timeout")

