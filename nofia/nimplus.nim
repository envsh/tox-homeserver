{.hint[XDeclaredButNotUsed]:off.}

import unicode
import json

proc Splitn*(s:string, n:int) : seq[string] =
    var rets: seq[string]
    var sub = newString(0)
    for c in s:
        sub.add(c)
        if sub.len() >= n:
            rets.add(sub)
            sub = newString(0)
    if strutils.strip(sub).len > 0:
        if sub.len < 5:
            ldebug("left sub", sub.len, sub[0], sub)
        rets.add(sub)
    return rets

proc Splitrn*(s:string, n:int) : seq[string] =
    var rets: seq[string]
    var sub = newString(0)
    for c in s.runes:
        sub.add(c)
        if sub.len() >= n:
            rets.add(sub)
            sub = newString(0)
    if unicode.strip(sub).len > 0: rets.add(sub)
    return rets

# // rune support, utf8 3byte, but ui width is 2
proc Splitrnui*(s:string, n:int) : seq[string] =
    var rets: seq[string]
    var sub = newString(0)
    var subuilen = 0
    for c in s.runes:
        let uilen = if c.size == 1: 1 else: 2
        if (subuilen + uilen) > n:
            rets.add(sub)
            sub = newString(0)
            subuilen = 0

        sub.add(c)
        subuilen += uilen
    if unicode.strip(sub).len > 0: rets.add(sub)
    return rets

proc tostr*(v:openArray[char]) : string = cast[string](v)
proc tou64*[T](v:ptr T) : ptr uint64 = cast[ptr uint64](v)
proc tou64*(v: Natural) : uint64 = cast[uint64](v)
proc toi64*(v: Natural) : int64 = cast[int64](v)
proc toi32*(v: Natural) : int32 = cast[int32](v)
proc tou32*(v: Natural) : uint32 = cast[uint32](v)
proc toptr*[T](v:ptr T) : pointer = cast[pointer](v)
proc toptr*(v:cstring) : pointer = cast[pointer](v)
proc tof32*(v: Natural) : float32 = return v.float # for float, cannot use cast, for cast is bitwise cast
proc tof64*(v: Natural) : float64 = return v.float
proc tojson*[T](v: T) : string = $(%*v)
proc toaddr(v: proc) : pointer = cast[pointer](v)
proc ntocstr(v: string) : cstring {.exportc.}= return v
proc ctonstr(v: cstring) : string {.exportc.} = $v

# some case c use 0 as success, 1 as failed
proc ctrue0*(v : Natural) : bool = v == 0
proc cfalse0*(v : Natural) : bool = v == 1
proc ctrue1*(v : Natural) : bool = v == 1
proc cfalse1*(v : Natural) : bool = v == 0

#let ctrue = cint(1)
#let cfalse = cint(0)

import macros
import tables
# 让他像 newseq[int](40)
macro newarr*(T:typedesc, sz:int) : untyped =
    result = quote do:
        var arr :array[`sz`, `T`]; arr

macro newseq*(T:typedesc, sz:int) : untyped =
    result = quote do: newseq[T](sz)

macro newtable*(TK:typedesc, TV:typedesc) :untyped =
    result = quote do:
        var map = initTable[`TK`, `TV`](); map

#var buf = newarr(char, 100)
#ldebug(buf.len())
#var buf = newarr(int, 123)
#ldebug(buf.len())
#var buf = newarr(string, 123)
#ldebug(buf.len())
#var tab = newtable(int,  string)
#ldebug(tab.len())

import json
import typeinfo
import typetraits

# 默认填充在obj有的字段，但jnode中没有的字段
# var obj = SomeTime()
# fixjsonnode(obj, jnode)
# jnode.to(obj)
proc fixjsonnode*[T](obj:var T, jnode:JsonNode)=
    # fields 需要作用于Tuple 或者 Object，ref不行，所以要.base
    var anyobj = obj.toany
    if anyobj.kind == akRef: anyobj = obj.toany.base
    elif anyobj.kind == akPtr: anyobj = obj.toany.base
    for x, y in anyobj.fields():
        if jnode.haskey(x): continue
        if y.kind == akInt64: jnode{x} = newJInt(0)
        elif y.kind == akUInt64: jnode{x} = newJInt(0)
        elif y.kind == akInt32: jnode{x} = newJInt(0)
        elif y.kind == akUInt32: jnode{x} = newJInt(0)
        elif y.kind == akInt16: jnode{x} = newJInt(0)
        elif y.kind == akUInt16: jnode{x} = newJInt(0)
        elif y.kind == akString: jnode{x} = newJString("")
        elif y.kind == akTuple: jnode{x} = newJObject()
        elif y.kind == akRef: jnode{x} = newJObject()
        else: linfo("Unknown", obj.type.name, y.kind)
    return

import times

# // 中文常用格式
proc totoday*(t: DateTime) : string = t.format("HH:mm:ss")
proc totodayminite*(t: DateTime): string = t.format("HH:mm")

