import threadpool
import asyncdispatch
setMinPoolSize(2)
setMaxPoolSize(3)

#[
extern void cxrt_init_env();
extern void cxrt_routine_post(void (*f)(void*), void*arg);
]#
proc cxrt_init_routine_env() {.importc.}
proc cxrt_routine_post(f:pointer, arg:pointer) {.importc.}
include "pthread_hook.nim"
pthread_setowner(1)
init_pthread_hook()
# create goroutines thread pool
cxrt_init_routine_env()
pthread_setowner(0)

# {.compile:"gogo.cpp".}
{.link:"gogo1.o"}
{.passl:"-lstdc++ -L/home/me/oss/src/cxrt/libgo -llibgo -lgc -lgccpp"}
{.passc:"-g -O0"}

include "nimlog.nim"

type
    # mirror of C TGenericSeq
    PNSeq = ptr NSeq
    NSeq = object
        len*: int
        reserved*: int
        data*: pointer

proc toaddr(v: proc) : pointer = cast[pointer](v)
proc ntocstr(v: string) : cstring {.exportc.} = return v
proc ctonstr(v: cstring) : string {.exportc.} = $v
proc todptr[T](v: seq[T]) : pointer = ((cast[PNSeq](v)).data).addr
proc todptr(v: string) : pointer = ((cast[PNSeq](v)).data).addr

import macros
import typeinfo

proc cxrt_routine_post_emu(fnptr:pointer, arg:pointer) =
    echo "routine post...", repr(arg)
    return

macro gogo(stmt : untyped) : untyped =
    var retstmt = newTree(nnkStmtList)
    # gen callback proc
    var cbprocstmt = newTree(nnkStmtList)
    var argtuplety = newTree(nnkTupleTy)
    for idx, s in stmt:
        if idx == 0: continue
        if s.kind == nnkStrLit:
            argtuplety.add(newIdentDefs(ident("a" & $idx), ident("string")))
        elif s.kind == nnkIntLit:
            argtuplety.add(newIdentDefs(ident("a" & $idx), ident("int")))
        ### TODO other more types

    cbprocstmt.add(newVarStmt(ident("fnargptr"),
                              newTree(nnkCast).add(newTree(nnkPtrTy).add(argtuplety)).add(ident("arg"))))
    cbprocstmt.add(newVarStmt(ident("fnarg"), newTree(nnkDerefExpr).add(ident("fnargptr"))))
    for idx, s in stmt:
        if idx == 0: continue
        cbprocstmt.add(newVarStmt(ident("v" & $idx),
                                  newTree(nnkBracketExpr).add(ident("fnarg"), newIntLitNode((idx-1)))))
    var callorigfn = newCall(stmt[0])
    for idx, s in stmt:
        if idx == 0: continue
        callorigfn.add(ident("v" & $idx))
    cbprocstmt.add(callorigfn)
    cbprocstmt.add(newTree(nnkReturnStmt).add(newEmptyNode()))
    var fnprms = newseq[NimNode]()
    fnprms.add(newEmptyNode())
    fnprms.add(newIdentDefs(ident("arg"), ident("pointer")))
    var fnval = newProc(params=fnprms, body=cbprocstmt)
    var fnvar = newVarStmt(ident("gogofn"), fnval)
    var fnptrvar = newVarStmt(ident("fnptr"), newCall(ident("toaddr"), ident("gogofn"))) # dont use .addr to get proc addr
    retstmt.add(fnvar)
    retstmt.add(fnptrvar)

    # gen callout stmt
    var parval = newPar()
    for idx, s in stmt:
        if idx == 0: continue
        parval.add(s)
    var tpvar = newVarStmt(ident("fnarg"), parval)
    retstmt.add(tpvar)
    var calladdr = newCall(ident("addr"), ident("fnarg"))
    var tpaddrvar = newVarStmt(ident("fnargptr"), calladdr)
    retstmt.add(tpaddrvar)
    var tp2var = newVarStmt(ident("fnargptr2"),
                                newCall(ident("alloc0"), newCall(ident("sizeof"), ident("fnarg"))))
    retstmt.add(tp2var)
    retstmt.add(newCall(ident("copyMem"), ident("fnargptr2"), ident("fnargptr"),
                        newCall(ident("sizeof"), ident("fnarg"))))
    for idx, s in stmt:
        if idx == 0: continue
        if s.kind == nnkStrLit:
            retstmt.add(newCall(ident("GC_ref"),
                                newTree(nnkBracketExpr).add(ident("fnarg"), newIntLitNode((idx-1)))))
    retstmt.add(newCall(ident("GC_fullCollect")))
    retstmt.add(newCall(ident("cxrt_routine_post"), ident("fnptr"), ident("fnargptr2")))
    var topstmt = newIfStmt((ident("true"), retstmt))
    echo repr(topstmt)
    result = topstmt
    # result = quote do: 123


### begin gogo2
type
    # mirror of typeinfo.Any
    SysAny = object
        value*: pointer
        rawTypePtr*: pointer

include "ffi.nim"

proc ffi_call2(cif:pffi_cif, fn:proc, rvalue: Any , avalue: pointer) =
    var fnptr = cast[pointer](fn)
    var rval2 = cast[SysAny](rvalue)
    ffi_call(cif, fnptr, rval2.value, avalue)
    return

# nim typeinfo.AnyKind to ffi_type*
proc nak2ffipty(ak: AnyKind) : pffi_type =
    if ak == akInt:
        return ffi_type_sint64.addr
    elif ak == akString:
        return ffi_type_pointer.addr
    elif ak == akCString:
        return ffi_type_pointer.addr
    elif ak == akPointer:
        return ffi_type_pointer.addr
    elif ak == akRef:
        return ffi_type_pointer.addr
    elif ak == akPtr:
        return ffi_type_pointer.addr
    elif ak == akSequence:
        return ffi_type_pointer.addr
    else: echo "unknown", ak
    return nil

proc gogorunner_cleanup(arg :pointer) =
    linfo "gogorunner_cleanup", repr(arg)
    var argv = cast[seq[pointer]](arg)
    var argc = cast[int](argv[1])
    for idx in 0..argc-1:
        let tyidx = 2 + idx*2
        let validx = tyidx + 1
        let akty = cast[AnyKind](argv[tyidx])
        if akty == akString:
            deallocShared(argv[validx])
        elif akty == akCString:
            deallocShared(argv[validx])
        elif akty == akInt:
            deallocShared(argv[validx])
        else: linfo "unknown", akty
    return

# pack struct, seq[pointer], which, [0]=fnptr, 1=argc, 2=a0ty, 3=a0val, 4=a1ty, 5=a1val ...
proc gogorunner(arg : pointer) =
    linfo "gogorunner", repr(arg)
    var argv = cast[seq[pointer]](arg)
    var fnptr  = argv[0]
    var argc = cast[int](argv[1])

    var atypes : seq[pffi_type]
    var avalues : seq[pointer]
    for idx in 0..argc-1:
        let tyidx = 2 + idx*2
        let validx = tyidx + 1
        let akty = cast[AnyKind](argv[tyidx])
        atypes.add(nak2ffipty(akty))
        linfo "recv val", idx, akty, argv[validx]
        if akty == akString:
            var cs = cast[cstring](argv[validx])
            var ns : string = $cs
            GC_ref(ns)
            avalues.add(ns.addr)
        elif akty == akCString:
            avalues.add(argv[validx])
        elif akty == akInt: avalues.add(argv[validx])
        else: linfo "unknown", akty, argv[validx]
        discard

    var cif : ffi_cif
    var rvalue : uint64
    # dump_pointer_array(argc.cint, atypes.todptr())
    var ret = ffi_prep_cif(cif.addr, FFI_DEFAULT_ABI, argc.cuint, ffi_type_pointer.addr, atypes.todptr())

    # dump_pointer_array(argc.cint, avalues.todptr())
    ffi_call(cif.addr, fnptr, rvalue.addr, avalues.todptr())
    gogorunner_cleanup(arg)
    return

# packed to passby format
proc packany(fn:proc, args:varargs[Any, toany]) =
    var pargs : seq[pointer]
    pargs.add(cast[pointer](fn))
    pargs.add(cast[pointer](args.len()))

    for idx, arg in args:
        var sarg = cast[SysAny](arg)
        if arg.kind == akInt:
            var v = allocShared0(sizeof(int))
            copyMem(v, sarg.value, sizeof(int))
            pargs.add(cast[pointer](akInt))
            pargs.add(v)
        elif arg.kind == akString:
            var ns = arg.getString()
            var cs : cstring = $ns
            var v = allocShared0(ns.len()+1)
            copyMem(v, cs, ns.len())
            pargs.add(cast[pointer](akString))
            pargs.add(v)
        elif arg.kind == akCString:
            var cs = arg.getCString()
            var v = allocShared0(cs.len()+1)
            copyMem(v, cs, cs.len())
            pargs.add(cast[pointer](akCString))
            pargs.add(v)
        else: linfo "unknown", arg.kind
        linfo "add val", idx, arg.kind, pargs[pargs.len-1]

    GC_ref(pargs)
    linfo "copy margs", pargs.len, cast[pointer](pargs)
    cxrt_routine_post(gogorunner.toaddr(), cast[pointer](pargs))
    return

macro gogo2(stmt:typed) : untyped =
    var nstmt = newStmtList()
    for idx, s in stmt:
        if idx == 0: continue
        linfo "aaa ", repr(s), " ", s.kind
        nstmt.add(newVarStmt(ident("a" & $idx), s))
    var packanycall = newCall(ident("packany"), stmt[0])
    for idx, s in stmt:
        if idx == 0: continue
        packanycall.add(ident("a" & $idx))

    nstmt.add(packanycall)
    var topstmt = newIfStmt((ident("true"), nstmt))
    linfo repr(topstmt)
    result = topstmt

### end gogo2

###
proc hello(i:int, s:string, cs:cstring) =
    linfo 123,"inhello, called by goroutines",getThreadId()

const usegogon = 2 # , 2
proc timeoutfn0(fd:AsyncFD):bool{.gcsafe.} = return false
proc timeoutfn1(fd:AsyncFD):bool{.gcsafe.} =
    if usegogon == 1:
        gogo hello(789, "nim lit string", "c lit string")
        gogo hello(789, "nim lit string", "c lit string")
    elif usegogon == 2:
        gogo2 hello(789, "nim lit string", "c lit string")
        gogo2 hello(789, "nim lit string", "c lit string")
    return false

proc timeoutfn2(fd:AsyncFD):bool{.gcsafe.} =
    if usegogon == 1:
        gogo hello(789, "nim lit string", "c lit string")
        gogo hello(789, "nim lit string", "c lit string")
    elif usegogon == 2:
        gogo2 hello(789, "nim lit string", "c lit string")
        gogo2 hello(789, "nim lit string", "c lit string")
    return false

proc timeoutfn3(fd:AsyncFD):bool{.gcsafe.} =
    if usegogon == 1:
        gogo hello(789, "nim lit string", "c lit string")
        gogo hello(789, "nim lit string", "c lit string")
    elif usegogon == 2:
        gogo2 hello(789, "nim lit string", "c lit string")
        gogo2 hello(789, "nim lit string", "c lit string")
    return false

if isMainModule:
    var seqinterp : seq[int]
    seqinterp.add(123)
    echo "main threadid",getThreadId()
    #gogo hello(789, "nim lit string", "c lit string")
    gogo2 hello(789, "nim lit string", "c lit string")
    addTimer(30000000, false, timeoutfn0)
    addTimer(2000, false, timeoutfn1)
    addTimer(1000, false, timeoutfn2)
    while true: poll(5000)

# expand expr:
#[
if true:
  var gogofn = proc (arg: pointer) =
    var fnargptr = cast[ptr tuple[a1: int, a2: string, a3: string]](arg)
    var fnarg = fnargptr[]
    var v1 = fnarg[0]
    var v2 = fnarg[1]
    var v3 = fnarg[2]
    hello(v1, v2, v3)
    return

  var fnptr = toaddr(gogofn)
  var fnarg = (789, "nim lit string", "c lit string")
  var fnargptr = addr(fnarg)
  var fnargptr2 = alloc0(sizeof(fnarg))
  copyMem(fnargptr2, fnargptr, sizeof(fnarg))
  cxrt_routine_post(fnptr, fnargptr2)
]#

# 这个pragma好友只在isMainModule生效
# 而且还只能声明一次，pragma already present
{.hint[XDeclaredButNotUsed]:off.}

