{.passc:"-fPIC -g -O0 -xc -DRENDER_X11_NATIVE"}
{.passl:"-lX11 -lXft -lXrender"}

# include 其他不使用全局变量的实现
{.compile: "render_x11_native.c.ngo".}
import x11/x, x11/xlib
include render_x11_native

{.passc:"-I/usr/include/freetype2"}
{.compile:"nuklear_x11_all.c.ngo".}
include nuklear_x11_all

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
        evtch*: Channel[MixEvent]
        wnds*: Table[string, proc (nkw:PNkwindow, name:string) {.gcsafe.}]
        wndrunner*: proc (nkw:PNkwindow) {.gcsafe.}

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

include "nimenv.nim"
include "views.nim"

proc dorepaint(nkw:PNkwindow, evts : seq[TXEvent]) =
    var rdwin = nkw.rdwin
    nk_input_begin(nkw.nkctx)
    for evt1 in evts:
        var evt = evt1
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
            # ldebug("nxtevt:", ret, "evty:", evt.theType, "pending left", ret1, repr(cast[pointer](nkw.rdwin.dpy)))
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
        else: sleep(300)

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
    createnkwndprocs(nkw.addr)

    linfo("chan is nil", repr(cast[pointer](addr(nkw.evtch))))
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
    linfo("rdwin is nil", nkw.rdwin == nil)

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

# runNimenv(ne1.self)


