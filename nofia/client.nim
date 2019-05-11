{.passl:"-lnng"}

proc newclient() : RpcClient =
    var cli = new(RpcClient)
    cli.rurl = "tcp://10.0.0.32:2081"

    var rv = nng_sub0_open(cli.subsk.addr)
    ldebug("sub0 open", rv, cli.subsk, sizeof(cli.subsk))
    var recvfd: int
    var getoptsz: int = sizeof(recvfd)
    var rv1 = nng_getopt(cli.subsk, NNG_OPT_RECVFD, recvfd.addr,  getoptsz.addr.tou64)
    ldebug("sub0 recvfd", rv1, recvfd, getoptsz, nng_strerror(rv1))
    cli.subrfd = cast[AsyncFD](recvfd)

    rv = nng_setopt(cli.subsk, NNG_OPT_SUB_SUBSCRIBE, nil, 0)
    rv = nng_req0_open(cli.reqsk.addr)
    ldebug("req0 open", rv)
    var rv2 = nng_getopt(cli.reqsk, NNG_OPT_SENDFD, recvfd.addr, getoptsz.addr.tou64)
    ldebug("req0 sendfd", rv2, recvfd, getoptsz, nng_strerror(rv1))
    cli.reqwfd = cast[AsyncFD](recvfd)

    rv2 = nng_getopt(cli.reqsk, NNG_OPT_RECVFD, recvfd.addr, getoptsz.addr.tou64)
    cli.reqrfd = cast[AsyncFD](recvfd)
    ldebug("req0 recvfd", rv2, recvfd, getoptsz, nng_strerror(rv1))

    return cli

proc connect(cli:RpcClient) =
    var rurl = cli.rurl
    var flag = NNG_FLAG_NONBLOCK.toi32
    flag = 0
    var rv = nng_dial(cli.subsk, rurl, nil, flag)
    ldebug("sub0 dial", rv, nng_strerror(rv))

    rurl = replace(rurl, "2081", "2082")
    rv = nng_dial(cli.reqsk, rurl, nil, flag)
    ldebug("req0 dial", rv, nng_strerror(rv))
    return

proc onsubskread(fd:AsyncFD):bool=
    ldebug("cli subsk readable", repr(fd))
    return

proc onreqskread(fd:AsyncFD):bool {.gcsafe.}=
    var ne = cast[PNimenv](getNimenvp())
    var cli = ne.rpcli
    var buf : array[8192,char]
    var pbuf = buf.addr.toptr
    var bsz  = buf.len()
    var rv = nng_recv(cli.reqsk, pbuf, bsz.addr.tou64, NNG_FLAG_NONBLOCK)
    ldebug("recved", rv, bsz)
    var data = cast[string](buf[..(bsz-1)]) # the string result
    return


proc reqsksend(cli:RpcClient, s:string)=
    var cs : cstring = $s
    var data = cs.toptr
    var rv = nng_send(cli.reqsk, data, s.len().tou64, NNG_FLAG_NONBLOCK)
    ldebug("cli reqsk send", rv, s.len(), nng_strerror(rv))
    return

proc getBaseInfo(cli:RpcClient) =
    var evt = Event()
    evt.Name = "GetBaseInfo"
    var jstr = $(%*evt)
    ldebug("send... req", jstr.len(), jstr)
    cli.reqsksend(jstr)
    return
