
{.compile:"pthread_hook.c".}
{.passc:"-D_GNU_SOURCE".}

# not need import threads

proc nim_pthread_getinfo(tra:pointer) {.importc.}

var thrstkbs {.threadvar.} : pointer

proc nim_pthread_proc(tra : pointer) =
    var stkbs : pointer
    thrstkbs = stkbs.addr
    when defined(setupForeignThreadGc): setupForeignThreadGc()
    echo ("nim_pthread_proc ...", tra == nil)
    nim_pthread_getinfo(tra)
    return

proc nim_pthread_create(tra:pointer) : int {.exportc.} =
    echo ("nim_pthread_create...", tra == nil)
    var nth : Thread[pointer]
    createThread(nth, nim_pthread_proc, tra)
    return 0

proc init_pthread_hook(thinitfn: pointer) {.importc.}
proc pthread_setowner(which: int) {.importc.}

# usage:
# include "pthread_hook.nim"
# pthread_setowner(1)
# init_pthread_hook()
# pthread_create
# uninit_pthread_hook()

### some about gc
type
    GCStackBase = ref object
        sb0*: GCStackBaseImpl
        sb1*: GCStackBaseImpl

    GCStackBaseImpl = object
        membase*: pointer
        regbase*: pointer
        bottom*: pointer
        gchandle*: pointer

proc GC_get_stack_base(sb:pointer) {.importc.}
proc GC_register_my_thread(sb:pointer) {.importc.}
proc GC_get_my_stackbottom(sb:pointer) : pointer {.importc.}
proc GC_set_stackbottom(gchandle:pointer, sb:pointer) {.importc.}
proc GC_call_with_alloc_lock(fn:pointer, arg: pointer) {.importc.}
proc GC_malloc(size:csize):pointer {.importc.}
proc GC_realloc(obj:pointer,size:csize):pointer {.importc.}
proc pthread_self() : uint {.importc.}

