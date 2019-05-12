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

proc reqsksend(cli:RpcClient, s:string)=
    var cs : cstring = $s
    var data = cs.toptr
    var rv = nng_send(cli.reqsk, data, s.len().tou64, NNG_FLAG_NONBLOCK)
    ldebug("cli reqsk send", rv, s.len(), nng_strerror(rv))
    return

proc getBaseInfo(cli:RpcClient) =
    var evt = Event()
    evt.EventName = "GetBaseInfo"
    var jstr = $(%*evt)
    ldebug("send... req", jstr.len(), jstr)
    cli.reqsksend(jstr)
    return

proc dispatchBaseInfo(cli:RpcClient, binfo : BaseInfo) =
    ldebug("dispatch BaseInfo")
    cli.binfo = binfo
    var mdl = getnkmdl()

    mdl.SetMyInfo(binfo.MyName, binfo.ToxId, binfo.Stmsg)
    mdl.SetMyConnStatus(binfo.ConnStatus)
    mdl.SetFriendInfos(binfo.Friends)
    mdl.SetGroupInfos(binfo.Groups)

    return

import typeinfo
import typetraits

proc decodeBaseInfo(cli:RpcClient, jnode: JsonNode) : BaseInfo =
    ldebug("decode BaseInfo")
    var binfo = BaseInfo()
    # try fill valid zero value for not exist fields
    fixjsonnode(binfo, jnode)

    binfo.EventId = jnode{"EventId"}.getInt.toi64
    binfo.EventName = jnode{"EventName"}.getstr
    binfo.ToxId = jnode{"ToxId"}.getstr
    binfo.MyName = jnode{"MyName"}.getstr
    binfo.Stmsg = jnode{"Stmsg"}.getstr
    binfo.Status1 = jnode{"Status1"}.getInt.tou32
    binfo.ConnStatus = jnode{"ConnStatus"}.getInt.toi32
    binfo.NextBatch = jnode{"NextBatch"}.getInt.toi64
    binfo.ToxVersion = jnode{"ToxVersion"}.getstr

    binfo.Friends = initTable[uint32,FriendInfo]()
    binfo.Groups = initTable[uint32,GroupInfo]()

    if not jnode.haskey("Friends"): jnode{"Friends"} = newJObject()
    for snum, fnode in jnode{"Friends"}.pairs:
        var frnd = FriendInfo()
        var num = parseInt(snum).tou32
        frnd.Fnum = num
        frnd.Status1 = fnode{"Status1"}.getint.tou32
        frnd.Name = fnode{"Name"}.getstr
        frnd.Stmsg = fnode{"Stmsg"}.getstr
        frnd.Avatar = fnode{"Avatar"}.getstr
        frnd.Seen = fnode{"Seen"}.getint.toi64
        frnd.ConnStatus = fnode{"ConnStatus"}.getint.toi32
        binfo.Friends.add(num, frnd)

    if not jnode.haskey("Groups"): jnode{"Groups"} = newJObject()
    for snum, gnode in jnode{"Groups"}.pairs:
        var grpo = GroupInfo()
        var num = parseInt(snum).tou32
        grpo.Gnum = num
        grpo.Mtype = gnode{"Mtype"}.getint.tou32
        grpo.GroupId = gnode{"GroupId"}.getstr
        grpo.Title = gnode{"Title"}.getstr
        grpo.Stmsg = gnode{"Stmsg"}.getstr
        grpo.Ours = gnode{"Ours"}.getbool
        #grpo.Members
        grpo.Members = initTable[string, MemberInfo]()
        if not gnode.haskey("Members"): jnode{"Members"} = newJObject()
        for snum, pnode in gnode{"Members"}.pairs:
            var mbro = MemberInfo()
            var pnum = parseInt(snum).tou32
            mbro.Pnum = pnum
            mbro.Pubkey = pnode{"Pubkey"}.getstr
            mbro.Name = pnode{"Name"}.getstr
            mbro.Mtype = pnode{"Mtype"}.getint
            mbro.Joints = pnode{"Joints"}.getint.toi64
            grpo.Members.add(mbro.Pubkey, mbro)

        binfo.Groups.add(num, grpo)

    return binfo

proc dispatchEvent(cli:RpcClient, evt : Event) =
    return

proc decodeEvent(cli:RpcClient, jnode: JsonNode) : Event =
    return

proc dispatchEventRaw(cli:RpcClient, data:string)=

    var jnode = parseJson(data)
    var jname = jnode{"EventName"}.getstr

    if jname == "GetBaseInfoResp":
        ldebug("got BaseInfo", cli == nil, jnode == nil)
        var binfo = cli.decodeBaseInfo(jnode["Binfo"])
        cli.dispatchBaseInfo(binfo)
    elif jname == "":
        linfo("Empty event name")
    else:
        var evt : Event = cli.decodeEvent(jnode)
        cli.dispatchEvent(evt)

    return

proc onreqskread(fd:AsyncFD):bool {.gcsafe.}=
    var ne = cast[PNimenv](getNimenvp())
    var cli = ne.rpcli
    var buf : array[8192,char]
    var pbuf = buf.addr.toptr
    var bsz  = buf.len()
    var rv = nng_recv(cli.reqsk, pbuf, bsz.addr.tou64, NNG_FLAG_NONBLOCK)
    ldebug("recved", rv, bsz)
    if bsz > buf.len(): lerror("Buf too small, need", bsz)
    var data = cast[string](buf[..(bsz-1)]) # the string result
    cli.dispatchEventRaw(data)
    return
