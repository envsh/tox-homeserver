{.passl:"-L. -lgocrt -l:libgo.a -lpthread"}

proc goinit(cnimcallfnptr : pointer) {.importc.}
proc gomainloop() {.importc.}
proc gogoimp(fn:pointer, args:pointer) {.importc:"gogo".} # 不直接调用的
proc gochannew(n:int) :pointer {.importc.}

# include "ffi.nim"

###
proc gogorunner(arg:pointer)

proc gogorunnerenter(arg:pointer) =
    return

proc gogorunnerleave(arg:pointer) =
    #GC_call_with_alloc_lock(gcsetbottom0.toaddr, sbp.sb0.addr)
    return

proc cnimcallimpl(fnptr: pointer, args: pointer) {.exportc.} =
    gogorunner(args)
    return

proc cnimcall(fnptr: pointer, args: pointer) {.exportc.} =
    setupForeignThreadGc2()
    linfo fnptr, args
    # echo "cnimcall",repr(fnptr),repr(args)
    # fnptr(args)
    var e = newAsyncEvent()
    cnimcallimpl(fnptr, args)
    return

goinit(cnimcall)

if isMainModule:
    discard
