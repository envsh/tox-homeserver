

# {.compile:"gogo.cpp".} # too slow
{.link:"gogo1.o"}
{.passl:"-L/home/me/oss/src/cxrt/libgo -llibgo -lstdc++"}
{.passl:"-L/home/me/oss/src/cxrt/bdwgc/.libs -lgc"}
{.passl:"-lpthread"}
{.passc:"-g -O0 -DGC_THREADS"}
{.passc:"-I /home/me/oss/src/cxrt/noro/include"}

proc cxrt_init_routine_env() {.importc.}
proc cxrt_routine_post(f:pointer, arg:pointer) {.importc.}

include "pthread_hook.nim"
var thrmap = newTable[uint,GCStackBase]() # thread_t => GCStackBase
var thrmapp = thrmap.addr
proc setupForeignThreadGc2() =
    when not defined(setupForeignThreadGc): discard
    else: setupForeignThreadGc()

#[
extern void cxrt_init_env();
extern void cxrt_routine_post(void (*f)(void*), void*arg);
]#

init_pthread_hook(cast[pointer](setupForeignThreadGc2)) # 由于链接顺序问题，pthread hook failed
pthread_setowner(1)
# create goroutines thread pool
pthread_setowner(0)
import os
sleep(1)

cxrt_init_routine_env()

proc gcsetbottom0(arg:pointer):pointer =
    var sbi = cast[ptr GCStackBaseImpl](arg)
    GC_set_stackbottom(sbi.gchandle, sbi)
    return nil

proc gcsetbottom1(arg:pointer):pointer =
    var sbi = cast[ptr GCStackBaseImpl](arg)
    GC_set_stackbottom(sbi.gchandle, sbi)
    return nil

proc gogorunnerenter(arg:pointer) =
    setupForeignThreadGc2()
    var thrh = pthread_self()
    var sbp = thrmapp[][thrh]
    var stksize : uint32
    var stkbase : pointer #libgo_currtask_stack(stksize.addr)
    var stkbottom = cast[pointer](cast[uint64](stkbase) + stksize.uint64)
    if stkbase != nil: # still nolucky let nim GC work
        sbp.sb1.membase = stkbottom
        sbp.sb1.gchandle = sbp.sb0.gchandle
        GC_call_with_alloc_lock(gcsetbottom1.toaddr, sbp.sb1.addr)
    else: linfo("wtf nil stkbase")

proc gogorunnerleave(arg:pointer) =
    GC_call_with_alloc_lock(gcsetbottom0.toaddr, sbp.sb0.addr)

