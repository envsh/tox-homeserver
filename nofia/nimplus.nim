
proc Splitn(s:string, n:int) : seq[string] =
    var rets: seq[string]
    var sub : string
    for c in s:
        sub.add(c)
        if sub.len() >= n:
            rets.add(sub)
            sub = ""
    return rets



proc tostr(v:openArray[char]) : string = cast[string](v)
proc tou64[T](v:ptr T) : ptr uint64 = cast[ptr uint64](v)
proc tou64(v: Natural) : uint64 = cast[uint64](v)
proc toi32(v: Natural) : int32 = cast[int32](v)
proc tou32(v: Natural) : uint32 = cast[uint32](v)
proc toptr[T](v:ptr T) : pointer = cast[pointer](v)
proc toptr(v:cstring) : pointer = cast[pointer](v)

# some case c use 0 as success, 1 as failed
proc ctrue0(v : Natural) : bool = v == 0
proc cfalse0(v : Natural) : bool = v == 1
proc ctrue1(v : Natural) : bool = v == 1
proc cfalse1(v : Natural) : bool = v == 0

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

