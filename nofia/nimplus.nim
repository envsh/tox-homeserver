
proc Splitn(s:string, n:int) : seq[string] =
    var rets: seq[string]
    var sub : string
    for c in s:
        sub.add(c)
        if sub.len() >= n:
            rets.add(sub)
            sub = ""
    return rets

    
