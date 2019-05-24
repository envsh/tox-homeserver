
# {.compile:"gogo.cpp".} # too slow
{.link:"gogo1.o"}
{.passl:"-L/home/me/oss/src/cxrt/libgo -llibgo -lstdc++"}
{.passl:"-L/home/me/oss/src/cxrt/bdwgc/.libs -lgc"}
{.passl:"-lpthread"}
{.passc:"-g -O0 -DGC_THREADS"}
{.passc:"-I /home/me/oss/src/cxrt/noro/include"}

proc libgo_currtask_stack(stsize: ptr uint32) : pointer {.cdecl, importc.}
proc libgo_set_thread_createcb(fnptr:pointer) {.cdecl, importc.}

proc libgo_thread_createcbfn() =
    var thrh = pthread_self()
    var sbp = new GCStackBase
    thrmap.add(thrh, sbp)
    GC_get_stack_base(sbp.sb0.addr)
    GC_register_my_thread(sbp.sb0.addr)
    sbp.sb0.gchandle = GC_get_my_stackbottom(sbp.sb0.addr)
    linfo "libgo thread created", thrmap.len(), getThreadId(), thrh
    linfo "libgo thread created", thrmap.len(), sbp.sb0.gchandle, sbp.sb0.membase, sbp.sb0.bottom
    return

libgo_set_thread_createcb(libgo_thread_createcbfn)

