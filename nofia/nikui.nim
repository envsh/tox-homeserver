{.passc:"-fPIC -g -O0 -DRENDER_X11_NATIVE"}
{.passl:"-lX11 -lXft -lXrender"}

# include 其他不使用全局变量的实现
{.compile: "render_x11_native.c".}
import x11/x, x11/xlib
include render_x11_native

type MixEvent = ref object
    typ*: int
    evts*: seq[TXEvent]

type
    PNkwindow = ptr Nkwindow
    Nkwindow = ref object
        rdwin*: TRenderWindow
        nkctx*: pointer
        evtch*: Channel[MixEvent]

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

proc dorepaint(nkw:PNkwindow, evts : seq[TXEvent]) =
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
            ldebug("evts sent", ok, evts.len(), nkw.evtch.ready(), nkw.evtch.peek())
            # sync()
        else: sleep(300)

    linfo("x11proc done")
    return

proc eventloop(nkw:PNkwindow) =
    linfo("eventlooping...")
    while true:
        # linfo("evtch recving...")
        var mxevt = nkw.evtch.recv()
        linfo("goty evty:", mxevt.typ, mxevt.evts.len(), nkw.evtch.peek())
        if mxevt.typ == 0: discard
        elif mxevt.typ == 1: dorepaint(nkw, mxevt.evts)
        elif mxevt.typ == 2: discard
    return

proc newNkwindow(nep:pointer): pointer {.exportc.} =
    var ne = cast[PNimenv](nep)
    var nkw = new(Nkwindow)
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

    spawn eventloop(pnkw)
    nkw.rdwin = NewRenderWindow()
    linfo("rdwin is nil", nkw.rdwin == nil)
    spawn x11proc(pnkw)
    return


### main section
discard XInitThreads()
var ne1 = newNimenvImpl()
var ne2 = newNimenvImpl()
linfo("singleton? ", ne1 == ne2)

spawn stopafter3s(ne1)

# runNimenv(ne1.self)


