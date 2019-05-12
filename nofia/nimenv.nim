
import os
import threadpool
import macros
import logging
import strutils
import tables


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
var grptr_of_nimenv_dont_use_other_than_this_file : pointer

# singleton
proc newNimenvImpl() : PNimenv =
    if gptr_of_nimenv_dont_use_other_than_this_file != nil:
        return gptr_of_nimenv_dont_use_other_than_this_file
    #gnenv = cast[PNimEnv](allocShared0(sizeof(NimEnv)))
    var ne = new(Nimenv)
    ne.asyevts = initTable[string, proc (pne: PNimenv)]()

    gref_of_nimenv_dont_use_other_than_this_file = ne
    gptr_of_nimenv_dont_use_other_than_this_file = addr(gref_of_nimenv_dont_use_other_than_this_file)
    grptr_of_nimenv_dont_use_other_than_this_file = cast[pointer](addr(gref_of_nimenv_dont_use_other_than_this_file))

    gptr_of_nimenv_dont_use_other_than_this_file.self = gref_of_nimenv_dont_use_other_than_this_file
    gptr_of_nimenv_dont_use_other_than_this_file.pself = gptr_of_nimenv_dont_use_other_than_this_file

    return gptr_of_nimenv_dont_use_other_than_this_file

proc newNimenv() : pointer {.exportc.} =
    return cast[pointer](newNimenvImpl())

# 解偶
proc getNimenvp() : pointer {.gcsafe.} = return grptr_of_nimenv_dont_use_other_than_this_file

# 解偶
proc getrpcli() : RpcClient = (cast[PNimenv](getNimenvp())).rpcli
proc getnkwnd() : NKWindow = (cast[PNimenv](getNimenvp())).nkxwin
proc getnkmdl() : DataModel = (cast[PNimenv](getNimenvp())).nkxwin.mdl


