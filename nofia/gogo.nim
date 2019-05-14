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
{.link:"gogo.o"}
{.passl:"-lstdc++ -L/home/me/oss/src/cxrt/libgo -llibgo -lgc -lgccpp"}
{.passc:"-g -O0"}

proc toaddr(v: proc) : pointer = cast[pointer](v)
proc ntocstr(v: string) : cstring {.exportc.}= return v
proc ctonstr(v: cstring) : string {.exportc.} = $v

import macros

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

macro gogo2(stmt:untyped):untyped =
    result = quote do:
        if true:

            discard

proc hello(i:int, s:string, cs:cstring) =
    echo 123,"inhello, called by goroutines",getThreadId()

proc timeoutfn0(fd:AsyncFD):bool{.gcsafe.}=
    gogo hello(789, "nim lit string", "c lit string")
    gogo hello(789, "nim lit string", "c lit string")
    return false

proc timeoutfn1(fd:AsyncFD):bool{.gcsafe.}=
    gogo hello(789, "nim lit string", "c lit string")
    gogo hello(789, "nim lit string", "c lit string")
    return false

proc timeoutfn2(fd:AsyncFD):bool{.gcsafe.}=
    gogo hello(789, "nim lit string", "c lit string")
    gogo hello(789, "nim lit string", "c lit string")
    return false

if isMainModule:
    echo "main threadid",getThreadId()
    gogo hello(789, "nim lit string", "c lit string")
    addTimer(3000, false, timeoutfn0)
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

{.hint[XDeclaredButNotUsed]:off.}

