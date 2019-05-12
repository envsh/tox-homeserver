
import json

proc dispatchEvent(ne: PNimenv, evto: Event) =
    var mdl = ne.nkxwin.mdl
    if evto.EventName == "ConferenceMessage":
        var groupId = evto.Margs[3]
        var message = evto.Args[3]
        var peerName = evto.Margs[0]
        var groupTitle = evto.Margs[2]
        var msgo = NewMessageForGroup(evto)
        mdl.Newmsg(groupId, msgo)
        discard
    else: discard
    return

proc dispatchNormEvent(ne:PNimenv, jsonNode: JsonNode) =
    var evt = Event()
    try: evt.EventId = jsonNode["EventId"].getInt()
    except: discard
    try: evt.EventName = jsonNode["EventName"].getStr()
    except: discard
    try:
        for e in jsonNode["Args"].getElems(): evt.Args.add(e.getStr())
    except: discard
    try:
        for e in jsonNode["Margs"].getElems(): evt.Margs.add(e.getStr())
    except: discard
    try: evt.ErrCode = cast[int32](jsonNode["ErrCode"].getInt())
    except: discard
    try: evt.ErrMsg = jsonNode["ErrMsg"].getStr()
    except: discard
    try: evt.UserCode = jsonNode["UserCode"].getInt()
    except: discard
    try: evt.DeviceUuid = jsonNode["DeviceUuid"].getStr()
    except: discard

    dispatchEvent(ne, evt)
    return

proc dispatchBaseInfo(ne:PNimenv, jsonNode:JsonNode) =
    linfo("process baseinfo", )
    var mdl = ne.nkxwin.mdl
    let jso = jsonNode

    mdl.SetMyInfo(jso["Name"].getStr(), jso["ToxId"].getStr(), jso["Stmsg"].getStr())
    mdl.SetMyConnStatus(jso["ConnStatus"].getInt())

    var frndsm = initTable[string, FriendInfo]()
    var frndsv :  seq[FriendInfo]
    var grpsm = initTable[string, GroupInfo]()
    var grpsv : seq[GroupInfo]

    for k, v in jso["Groups"].pairs():
        # linfo(k, v,)
        var grpo = new(GroupInfo)
        grpo.Gnum = cast[uint32](parseInt(k))
        grpo.Title = v["Title"].getStr()
        grpo.GroupId = v["GroupId"].getStr()
        grpsv.add(grpo)
        grpsm.add(grpo.GroupId, grpo)

    mdl.Groupsm = grpsm
    mdl.Groupsv = grpsv

    for k,v in jso["Friends"].pairs():
        # linfo(k, v, )
        var frndo = new(FriendInfo)
        frndo.Fnum = cast[uint32](parseInt(k))
        #frndo.Status1 = cast[uint32](v["Status"].getInt())
        frndo.Pubkey = v["Pubkey"].getStr()
        try: frndo.Name = v["Name"].getStr()
        except: discard
        try: frndo.Stmsg = v["Stmsg"].getStr()
        except: discard
        try: frndo.ConnStatus = cast[int32](v["ConnStatus"].getInt())
        except: discard

        frndsv.add(frndo)
        frndsm.add(frndo.Pubkey, frndo)

    mdl.Friendsm = frndsm
    mdl.Friendsv = frndsv
    return

proc dispatchEventNim(ne: pointer, evtdat : cstring) {.exportc.} =
    let s : string = $evtdat
    let jsonNode = parseJson(s)
    var isnorm = false
    var isbase = false
    var iserpc = false

    isnorm = jsonNode.hasKey("EventId")
    isbase = jsonNode.hasKey("Friends")

    var one = cast[PNimenv](ne)
    if isnorm: dispatchNormEvent(one, jsonNode)
    elif isbase: dispatchBaseInfo(one, jsonNode)
    elif iserpc: linfo("unimpled event", s)
    else: linfo("unknown event", s)
    return

proc dispatchEventRespNim(ne:pointer,evtdat : cstring) {.exportc.} =
    linfo("evtdat", evtdat.len())
    return
