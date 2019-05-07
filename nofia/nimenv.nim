
import os
import threadpool
import macros
import logging
import strutils
import tables

### logging initilize
macro logecho(lvl : string, msgs: varargs[untyped]): untyped =
    result = nnkStmtList.newTree()
    result.add newCall("echo", newLit("[" & $lvl & "] "))
    let lineobj = msgs.lineInfoObj()
    let sepos = lineobj.filename.rfind("/")
    let shortfname = lineobj.filename[sepos+1..lineobj.filename.len()-1]
    result[0].add(newStrLitNode(shortfname))
    result[0].add(newLit(":"))
    result[0].add(newLit(lineobj.line))
    for msg in msgs:
        result[0].add(newLit(" "))
        result[0].add(msg)
    discard # 可能等于 python的pass

macro ldebug(msgs: varargs[untyped]): untyped =
    result = quote do: logecho("DEBUG", `msgs`)
    discard # 可能等于 python的pass

macro linfo(msgs: varargs[untyped]): untyped =
    result = quote do: logecho("INFO ", `msgs`)
    discard # 可能等于 python的pass

macro lerror(msgs: varargs[untyped]): untyped =
    result = quote do: logecho("ERROR", `msgs`)
    discard

#[
echo ("123", 456)
linfo("hehehe",123)
lerror("hehehe",456)
ldebug("hehehe有",true)
ldebug("hehehe在",false)
]#

proc c2nimbool(ok:cint):bool = return if ok == 1: true else: false
proc nim2cbool(ok:bool):cint = return if ok: 1 else: 0
let ctrue = cint(1)
let cfalse = cint(0)

### begin emu code, 不能使用全局变量的部分
proc freeNimenv(ne : pointer) {.exportc.} =
    if ne == nil: return
    (cast[PNimenv](ne)).stoped = true
    return


# stop after 3 secs
proc stopafter3s(ne : PNimenv) =
    sleep(300012345)
    freeNimEnv(ne)
    discard

proc runNimenv(ne : Nimenv) =
    while ne == nil or (ne != nil and not ne.stoped):
        echo(repr(ne.stoped) & repr(ne.stoped))
        sleep(1000)
    echo("nim main proc done")


### 全局部分
var gref_of_nimenv_dont_use_other_than_this_file : Nimenv
var gptr_of_nimenv_dont_use_other_than_this_file : PNimenv

# singleton
proc newNimenvImpl() : PNimenv =
    if gptr_of_nimenv_dont_use_other_than_this_file != nil:
        return gptr_of_nimenv_dont_use_other_than_this_file
    #gnenv = cast[PNimEnv](allocShared0(sizeof(NimEnv)))
    gref_of_nimenv_dont_use_other_than_this_file = new(Nimenv)
    gptr_of_nimenv_dont_use_other_than_this_file = addr(gref_of_nimenv_dont_use_other_than_this_file)

    gptr_of_nimenv_dont_use_other_than_this_file.self = gref_of_nimenv_dont_use_other_than_this_file
    gptr_of_nimenv_dont_use_other_than_this_file.pself = gptr_of_nimenv_dont_use_other_than_this_file
    return gptr_of_nimenv_dont_use_other_than_this_file

proc newNimenv() : pointer {.exportc.} =
    return cast[pointer](newNimenvImpl())

