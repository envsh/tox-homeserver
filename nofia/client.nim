{.passl:"-lnng"}

# woc, nng会启动16 个线程：https://github.com/nanomsg/nng/issues/769

proc newclient() : RpcClient =
    var cli = new(RpcClient)
    cli.devuuid = "0c5b3037-3767-4c66-b9e4-46aff8d693b1"
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


proc reqsksend(cli:RpcClient, s:string) : bool =
    var cs : cstring = $s
    var data = cs.toptr
    var rv = nng_send(cli.reqsk, data, s.len().tou64, NNG_FLAG_NONBLOCK)
    if not rv.ctrue0: lerror("cli reqsk send error", rv, nng_strerror(rv))
    return rv.ctrue0

proc getBaseInfo(cli:RpcClient) =
    var evt = Event()
    evt.EventName = "GetBaseInfo"
    evt.DeviceUuid = cli.devuuid
    var jstr = $(%*evt)
    ldebug("send... req", jstr.len(), jstr)
    discard cli.reqsksend(jstr)
    return

include "crc64.nim"
var userCodeSeq : int = 0
# 无序不重复码
# clinfo: device uuid for general
proc NextUserCode(clinfo : string) : int64 =
    var s = clinfo & "-" & $atomicInc(userCodeSeq, 1)
    return crc64(0, s, s.len.tou64).toi64

proc SelfGetName(cli:RpcClient) : string =
    return cli.binfo.MyName

proc FriendSendMessage(cli:RpcClient, ctnum: uint32, msg: string, rptid : int64) : bool =
    var evt = Event()
    evt.EventId = rptid
    evt.EventName = "FriendSendMessage"
    evt.Args.add($ctnum)
    evt.Args.add(msg)

    var jstr = $(%*evt)
    return cli.reqsksend(jstr)

proc ConferenceSendMessage(cli:RpcClient, ctnum: uint32, msg: string, rptid:int64) : bool =
    var evt = Event()
    evt.EventId = rptid
    evt.EventName = "ConferenceSendMessage"
    evt.Args.add($ctnum)
    evt.Args.add($0)
    evt.Args.add(msg)

    var jstr = $(%*evt)
    return cli.reqsksend(jstr)

proc dispatchBaseInfo(cli:RpcClient, binfo : BaseInfo) =
    ldebug("dispatch BaseInfo")
    cli.binfo = binfo
    var mdl = getnkmdl()

    mdl.SetMyInfo(binfo.MyName, binfo.ToxId, binfo.Stmsg)
    mdl.SetMyConnStatus(binfo.ConnStatus)
    mdl.SetFriendInfos(binfo.Friends)
    mdl.SetGroupInfos(binfo.Groups)

    ldebug("mdl info", "grplen", binfo.Groups.len)
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
        var fnum = parseInt(snum).tou32
        frnd.Fnum = fnum
        frnd.Status1 = fnode{"Status1"}.getint.tou32
        frnd.Name = fnode{"Name"}.getstr
        frnd.Stmsg = fnode{"Stmsg"}.getstr
        frnd.Avatar = fnode{"Avatar"}.getstr
        frnd.Seen = fnode{"Seen"}.getint.toi64
        frnd.ConnStatus = fnode{"ConnStatus"}.getint.toi32
        binfo.Friends.add(fnum, frnd)

    if not jnode.haskey("Groups"): jnode{"Groups"} = newJObject()
    for snum, gnode in jnode{"Groups"}.pairs:
        var grpo = GroupInfo()
        var gnum = parseInt(snum).tou32
        grpo.Gnum = gnum
        grpo.Mtype = gnode{"Mtype"}.getint.tou32
        grpo.GroupId = gnode{"GroupId"}.getstr
        grpo.Title = gnode{"Title"}.getstr
        grpo.Stmsg = gnode{"Stmsg"}.getstr
        grpo.Ours = gnode{"Ours"}.getbool
        #grpo.Members
        grpo.Members = initTable[string, MemberInfo]()
        if not gnode.haskey("Members"): jnode{"Members"} = newJObject()
        for pkey, pnode in gnode{"Members"}.pairs:
            var mbro = MemberInfo()
            mbro.Pnum = pnode{"Pnum"}.getint.tou32
            mbro.Pubkey = pnode{"Pubkey"}.getstr
            mbro.Name = pnode{"Name"}.getstr
            mbro.Mtype = pnode{"Mtype"}.getint
            mbro.Joints = pnode{"Joints"}.getint.toi64
            grpo.Members.add(mbro.Pubkey, mbro)

        binfo.Groups.add(gnum, grpo)

    return binfo

proc dispatchEvent(cli:RpcClient, evt : Event) =
    ldebug("dispatch common event", evt.EventName)
    var mdl = getnkmdl()

    if evt.EventName == "ConferenceMessage":
        var groupId = evt.Margs[3]
        var message = evt.Args[3]
        var peerName = evt.Margs[0]
        var groupTitle = evt.Margs[2]
        var msgo = NewMessageForGroup(evt)
        mdl.Newmsg(groupId, msgo)
    else: discard

    return

proc decodeEvent(cli:RpcClient, jnode: JsonNode) : Event =
    var evt = Event()
    fixjsonnode(evt, jnode)
    ldebug("decode common event", jnode{"EventName"}.getstr)

    # evt = jnode.to(Event) # still Error: unhandled exception: key not found:  [KeyError]
    evt.EventId = jnode{"EventId"}.getint
    evt.EventName = jnode{"EventName"}.getstr

    for e in jnode{"Args"}.getElems(): evt.Args.add(e.getstr)
    for e in jnode["Margs"].getElems(): evt.Margs.add(e.getstr)

    evt.ErrCode = jnode{"ErrCode"}.getint.toi32
    evt.ErrMsg = jnode{"ErrMsg"}.getstr
    evt.UserCode = jnode{"UserCode"}.getint
    evt.DeviceUuid = jnode{"DeviceUuid"}.getstr()

    return evt

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

    # TODO optional error
    # maybe subsk or reqsk
    # always NONBLOCK
proc readclisk(sk:nng_socket) : (string, int32) =
    var ne = cast[PNimenv](getNimenvp())
    var cli = ne.rpcli
    var buf : array[8192,char]
    var pbuf = buf.addr.toptr
    var bsz  = buf.len()
    var rv = nng_recv(sk, pbuf, bsz.addr.tou64, NNG_FLAG_NONBLOCK)
    ldebug("recved", rv, bsz)
    if bsz > buf.len(): lerror("Buf too small, need", bsz)
    var data = cast[string](buf[..(bsz-1)]) # the string result
    return (data, rv)

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


proc onsubskread(fd:AsyncFD):bool=
    ldebug("cli subsk readable", repr(fd))
    var ne = cast[PNimenv](getNimenvp())
    var cli = ne.rpcli
    var buf : array[2192,char]
    var pbuf = buf.addr.toptr
    var bsz  = buf.len()
    var rv = nng_recv(cli.subsk, pbuf, bsz.addr.tou64, NNG_FLAG_NONBLOCK)
    ldebug("recved", rv, bsz)
    if bsz > buf.len(): lerror("Buf too small, need", bsz)
    var data = cast[string](buf[..(bsz-1)]) # the string result
    cli.dispatchEventRaw(data)
    return
