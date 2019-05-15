
{.compile:"pthread_hook.c".}
{.passc:"-D_GNU_SOURCE".}

# not need import threads

proc nim_pthread_getinfo(tra:pointer) {.importc.}

proc nim_pthread_proc(tra : pointer) =
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
