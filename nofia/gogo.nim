# GC_disableMarkAndSweep()

#import threadpool
#setMinPoolSize(2)
#setMaxPoolSize(3)
import asyncdispatch
import tables
# import coro
include "nimlog.nim"

#proc GC_addStack(bottom: pointer) {.cdecl, importc.}
proc GC_removeStack(bottom: pointer) {.cdecl, importc.}
#proc GC_setActiveStack(bottom: pointer) {.cdecl, importc.}
#proc nimGC_setStackBottom(bottom: pointer) {.cdecl, importc.}
proc libgo_currtask_stack(stsize: ptr uint32) : pointer {.cdecl, importc.}
proc libgo_set_thread_createcb(fnptr:pointer) {.cdecl, importc.}
proc toaddr[T](v: ref T) : pointer = cast[pointer](v)
#[
extern void cxrt_init_env();
extern void cxrt_routine_post(void (*f)(void*), void*arg);
]#
proc cxrt_init_routine_env() {.importc.}
proc cxrt_routine_post(f:pointer, arg:pointer) {.importc.}
include "pthread_hook.nim"
var thrmap = newTable[uint,GCStackBase]() # thread_t => GCStackBase
var thrmapp = thrmap.addr
proc setupForeignThreadGc2() =
    when not defined(setupForeignThreadGc): discard
    else: setupForeignThreadGc()
proc libgo_thread_createcbfn() =
    var thrh = pthread_self()
    var sbp = new GCStackBase
    thrmap.add(thrh, sbp)
    GC_get_stack_base(sbp.sb0.addr)
    GC_register_my_thread(sbp.sb0.addr)
    sbp.sb0.gchandle = GC_get_my_stackbottom(sbp.sb0.addr)
    linfo "libgo thread created", getThreadId(), thrh
    linfo "libgo thread created", sbp.sb0.gchandle, sbp.sb0.membase, sbp.sb0.bottom
    return
libgo_set_thread_createcb(libgo_thread_createcbfn)
init_pthread_hook(cast[pointer](setupForeignThreadGc2)) # 由于链接顺序问题，pthread hook failed
pthread_setowner(1)
# create goroutines thread pool
cxrt_init_routine_env()
pthread_setowner(0)

# {.compile:"gogo.cpp".} # too slow
{.link:"gogo1.o"}
{.passl:"-L/home/me/oss/src/cxrt/libgo -llibgo -lstdc++"}
{.passl:"-L/home/me/oss/src/cxrt/bdwgc/.libs -lgc"}
{.passl:"-lpthread"}
{.passc:"-g -O0 -DGC_THREADS"}
{.passc:"-I /home/me/oss/src/cxrt/noro/include"}



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
    var argc = cast[int](pointer_array_get(arg, 1))
    for idx in 0..argc-1:
        let tyidx = 2 + idx*2
        let validx = tyidx + 1
        let akty = cast[AnyKind](pointer_array_get(arg, tyidx.cint))
        let akval = pointer_array_get(arg, validx.cint)
        if akty == akString:
            deallocShared(akval)
        elif akty == akCString:
            deallocShared(akval)
        elif akty == akInt:
            deallocShared(akval)
        else: linfo "unknown", akty
    #deallocShared(arg)
    pointer_array_free(arg)
    return

import tables
var gostacks = newTable[pointer, pointer]()
var gostacksptr = gostacks.addr

proc gcsetbottom0(arg:pointer):pointer =
    var sbi = cast[ptr GCStackBaseImpl](arg)
    GC_set_stackbottom(sbi.gchandle, sbi)
    return nil

proc gcsetbottom1(arg:pointer):pointer =
    var sbi = cast[ptr GCStackBaseImpl](arg)
    GC_set_stackbottom(sbi.gchandle, sbi)
    return nil

# pack struct, seq[pointer], which, [0]=fnptr, 1=argc, 2=a0ty, 3=a0val, 4=a1ty, 5=a1val ...
proc gogorunner(arg : pointer) =
    setupForeignThreadGc2()
    var thrh = pthread_self()
    var sbp = thrmapp[][thrh]
    var stksize : uint32
    var stkbase = libgo_currtask_stack(stksize.addr)
    var stkbottom = cast[pointer](cast[uint64](stkbase) + stksize.uint64)
    if stkbase != nil: # still nolucky let nim GC work
        if not gostacksptr[].haskey(stkbase):
            gostacksptr[].add(stkbase, nil)
            #GC_addStack(stkbase)
        #GC_setActiveStack(stkbase)
        nimGC_setStackBottom(stkbase)
        sbp.sb1.membase = stkbottom
        sbp.sb1.gchandle = sbp.sb0.gchandle
        GC_call_with_alloc_lock(gcsetbottom1.toaddr, sbp.sb1.addr)
    else: linfo("wtf nil stkbase")

    linfo "gogorunner", arg, "stk:", stksize, stkbase, stkbottom, "ostk:", thrstkbs
    var fnptr = pointer_array_get(arg, 0)
    var argc = cast[int](pointer_array_get(arg, 1))
    assert(argc == 3, $argc)

    var atypes : seq[pffi_type]
    var avalues : seq[pointer]
    for idx in 0..argc-1:
        let tyidx = 2 + idx*2
        let validx = tyidx + 1
        let akty = cast[AnyKind](pointer_array_get(arg, tyidx.cint))
        let akval = pointer_array_get(arg, validx.cint)
        atypes.add(nak2ffipty(akty))
        if akty == akString:
            var cs = cast[cstring](akval)
            var ns : string = $cs
            GC_ref(ns)
            avalues.add(ns.addr)
        elif akty == akCString:
            var cs = cast[cstring](akval)
            avalues.add(cs.addr)
        elif akty == akInt: avalues.add(akval)
        else: linfo "unknown", akty, akval
        discard

    var cif : ffi_cif
    var rvalue : uint64
    # dump_pointer_array(argc.cint, atypes.todptr())
    var ret = ffi_prep_cif(cif.addr, FFI_DEFAULT_ABI, argc.cuint, ffi_type_pointer.addr, atypes.todptr())

    # dump_pointer_array(argc.cint, avalues.todptr())
    ffi_call(cif.addr, fnptr, rvalue.addr, avalues.todptr())
    gogorunner_cleanup(arg)
    nimGC_setStackBottom(thrstkbs)
    GC_call_with_alloc_lock(gcsetbottom0.toaddr, sbp.sb0.addr)
    return

# packed to passby format
proc packany(fn:proc, args:varargs[Any, toany]) =
    var ecnt = (2+args.len()*2+2)
    var pargs = pointer_array_new(ecnt.cint)
    pointer_array_set(pargs, 0, cast[pointer](fn))
    pointer_array_set(pargs, 1, cast[pointer](args.len()))

    for idx in 0..args.len-1:
        var arg = args[idx]
        var tyidx = 2+idx*2
        var validx = tyidx+1
        pointer_array_set(pargs, tyidx.cint, cast[pointer](arg.kind))
        var sarg = cast[SysAny](arg)
        if arg.kind == akInt:
            var v = allocShared0(sizeof(int))
            copyMem(v, sarg.value, sizeof(int))
            pointer_array_set(pargs, validx.cint, v)
        elif arg.kind == akString:
            var ns = arg.getString()
            var cs : cstring = $ns
            var v = allocShared0(ns.len()+1)
            copyMem(v, cs, ns.len())
            pointer_array_set(pargs, validx.cint, v)
        elif arg.kind == akCString:
            var cs = arg.getCString()
            var v = allocShared0(cs.len()+1)
            copyMem(v, cs, cs.len())
            pointer_array_set(pargs, validx.cint, v)
        else: linfo "unknown", arg.kind

    linfo "copy margs", 2+args.len*2, pargs # why refc=1
    cxrt_routine_post(gogorunner.toaddr(), pargs)
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
    linfo 123,"inhello, called by goroutines"#,getThreadId()
    var p : pointer
    p = GC_malloc(5678)
    p = GC_malloc(6789)
    p = GC_malloc(7890)
    p = GC_malloc(8901)
    return

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
    linfo "main threadid"#, getThreadId()
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

