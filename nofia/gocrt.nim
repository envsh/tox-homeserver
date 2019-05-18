{.passl:"-L. -lgocrt"}

proc goinit(cnimcallfnptr : pointer) {.importc.}
proc gomain() {.importc.}
proc gogoimp(fn:pointer, args:pointer) {.importc:"gogo".} # 不直接调用的
proc gochannew(n:int) :pointer {.importc.}

include "ffi.nim"

###
proc gogorunner(arg:pointer)

proc cnimcallimpl(fnptr: pointer, args: pointer) {.exportc.} =
    gogorunner(args)
    return

proc cnimcall(fnptr: pointer, args: pointer) {.exportc.} =
    setupForeignThreadGc()
    linfo fnptr, args
    # echo "cnimcall",repr(fnptr),repr(args)
    # fnptr(args)
    var e = newAsyncEvent()
    cnimcallimpl(fnptr, args)
    return

goinit(cnimcall)

