
proc hello(i:int, s:string, cs:cstring) =
    echo 123

proc toaddr(v: proc) : pointer = cast[pointer](v)
proc ntocstr(v: string) : cstring {.exportc.}= return v
proc ctonstr(v: cstring) : string {.exportc.} = $v

import macros

proc cxrt_routine_post(fnptr:pointer, arg:pointer) =
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
    var fnptrvar = newVarStmt(ident("fnptr"), newCall(ident("addr"), ident("gogofn")))
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
    retstmt.add(newCall(ident("cxrt_routine_post"), ident("fnptr"), ident("fnargptr")))
    echo repr(retstmt)
    result = retstmt
    # result = quote do: 123

if isMainModule:
    gogo hello(789, "nim lit string", "c lit string")
# expand expr:
# caller convert
# =>      var arg = (a0, a1, a2)
# =>      var argp = vtp.addr
# =>      var gogofn = proc (p:pointer): ...
# =>      hello_gogo(argp)
# callee convert
# => hello_gogo(argp:pointer)
# => var tpvptr = cast[ptr tuple[x0:int, x1:string, x2:string]](arpg)
# => var (a0, a1, a2) = tpvptr[]
# => hello(a0, a1, a2)


{.hint[XDeclaredButNotUsed]:off.}

