import macros
import logging
import strutils

proc `$`*(x: pointer):string = return repr(x)

### logging initilize
const LOGLVL_TRACE = 0
const LOGLVL_DEBUG = 1
const LOGLVL_INFO = 2
const LOGLVL_WARNING = 3
const LOGLVL_ERROR = 4
const LOGLVL_FATAL = 5

# nim plus log level
var nploglevel = LOGLVL_DEBUG
proc llevel(lvl: int) =
    if lvl >= LOGLVL_TRACE and lvl <= LOGLVL_FATAL:
        nploglevel = lvl

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
        if msg.typeKind() == ntyPointer: result[0].add(newLit(repr(msg)))
        else: result[0].add(msg)
    #discard # 可能等于 python的pass

macro ldebug(msgs: varargs[untyped]): untyped =
    result = quote do: logecho("DEBUG", `msgs`)
    #discard # 可能等于 python的pass

macro linfo(msgs: varargs[untyped]): untyped =
    result = quote do: logecho("INFO ", `msgs`)
    #discard # 可能等于 python的pass

macro lerror(msgs: varargs[untyped]): untyped =
    result = quote do: logecho("ERROR", `msgs`)
    # discard


#[
echo ("123", 456)
linfo("hehehe",123)
lerror("hehehe",456)
ldebug("hehehe有",true)
ldebug("hehehe在",false)
]#
