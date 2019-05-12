
# TODO use mutex for write

const CTTYPE_FRIEND = 1
const CTTYPE_GROUP = 2

import tables
import algorithm

type
    PDataModel = ptr DataModel
    DataModel = ref object
        # mu
        Myid*: string
        Myname*: string
        Mystmsg*: string
        Mystno*: int
        Mysttxt*: string #  // status text
        lastno*: int # // last valid status no, in case rpc

        # // for ChatForm and SendForm
        # // currently active state
        Frndinfo*: FriendInfo
        Grpinfo*:  GroupInfo
        Cttype*: int
        Ctnum*: uint32
        Ctuniqid*: string # // current active contact identifier ==> cur
        Ctname*: string # // name or title
        Ctstmsg*: string
        receiptid*: int64

        # // scrollbar y position for all contact chat session
        # // 对于每个会话的值，当活动窗口时，有新消息立即滚动到最底
        # // 当切换离开一个窗口时，记录当前位置
        # // 当切换到一个窗口时，如果没有新消息，则使用上次记录下的位置
        Scrollbarys*: Table[string,int]

        Friendsm*: Table[string,FriendInfo] # *thspbs.FriendInfo // uniqid =>
        Friendsv*: seq[FriendInfo] #
        Groupsm*: Table[string,GroupInfo] # *thspbs.GroupInfo // uniqid =>
        Groupsv*: seq[GroupInfo] #

        Ctmsgs*: Table[string,seq[Message]] #// uniqid =>
        Hasnews*: Table[string,int] # // uniqid => , 某个联系人的未读取消息个数
        lastmsg*: Message # // Lastmsg must be not belongs to active contact chatform
        lastctname*: string # // always according with lastmsg

        repainter*: proc()


proc newDataModel() : DataModel =
    var mdl = DataModel()
    mdl.Scrollbarys = initTable[string,int]()
    mdl.Friendsm = initTable[string,FriendInfo]()
    mdl.Friendsv = newseq[FriendInfo]()
    mdl.Groupsm = initTable[string,GroupInfo]()
    mdl.Groupsv = newseq[GroupInfo]()
    mdl.Ctmsgs = initTable[string,seq[Message]]()
    mdl.Hasnews = initTable[string, int]()
    return mdl

proc Nxtreceiptid(mdl:DataModel) : int64 =
    return 0
    # { return atomic.AddInt64(&this.receiptid, 1) }

proc  emitChanged(mdl:DataModel) =
    #if this.repainter != nil {
    #	this.repainter()
    return

proc SetMyInfo(this:DataModel, name:string, id: string, stmsg: string) =
    #defer this.emitChanged()
    #this.mu.Lock()
    #defer this.mu.Unlock()

    this.Myname = name
    this.Myid = id
    this.Mystmsg = stmsg
    return

proc Conno2str(stno: int) : string =
    if stno == 0: return "NOE"
    elif stno == 1: return "TCP"
    elif stno == 2: return "UDP"
    # // connect between client and homeserver
    elif stno == 5: return "BRK"
    else: return "UNK"
    return ""

proc Conno2str1(stno: int) : string =
    let s = Conno2str(stno)
    return s.substr(0, 0)

proc SetMyConnStatus(this:DataModel, stno: int) =
    #defer this.emitChanged()
    #this.mu.Lock()
    #defer this.mu.Unlock()

    if stno == 5:
        this.lastno = this.Mystno
        this.Mystno = stno
    elif stno == -5:
        if this.Mystno == 5: this.Mystno = this.lastno
    else: this.Mystno = stno

    this.Mysttxt = Conno2str(this.Mystno)
    return

proc friendcmper(x, y : FriendInfo) :int =
    if x.Fnum < y.Fnum: return 1
    elif x.Fnum > y.Fnum: return -1
    else: return 0

proc SetFriendInfos(this:DataModel, friends: Table[uint32,FriendInfo]) =
    var newedm = initTable[string,FriendInfo]()
    var newedv = newseq[FriendInfo]()
    for k, v in friends.pairs():
        var f = v
        newedm.add(v.Pubkey, f)
        newedv.add(f)

    # sort
    sort(newedv, friendcmper)

    this.Friendsm = newedm
    this.Friendsv = newedv
    return

proc groupcmper (x, y : GroupInfo) :int =
    if x.Gnum < y.Gnum: return 1
    elif x.Gnum > y.Gnum: return -1
    else: return 0

proc SetGroupInfos(this:DataModel, groups: Table[uint32,GroupInfo]) =
    var newedm = initTable[string,GroupInfo]()
    var newedv = newseq[GroupInfo]()
    for k, v in  groups.pairs:
        var g = v # TODO deepcopy?
        newedm.add(v.GroupId, g)
        newedv.add(g)

    # sort
    sort(newedv, groupcmper)

    this.Groupsm = newedm
    this.Groupsv = newedv
    return

proc FriendList(mdl:DataModel) : seq[FriendInfo] = return mdl.Friendsv

proc GroupList(mdl:DataModel) : seq[GroupInfo] =
    return mdl.Groupsv

# // current
proc setFriendInfo(this:DataModel, fi: FriendInfo) =
# 	// this.mu.Lock()
# 	// defer this.mu.Unlock()
    this.Frndinfo = fi
    this.Cttype = CTTYPE_FRIEND
    this.Ctname = fi.Name
    this.Ctstmsg = fi.Stmsg
    this.Ctnum = fi.Fnum
    return

# // current
proc setGroupInfo(this :DataModel,fi: GroupInfo) =
# 	// this.mu.Lock()
# 	// defer this.mu.Unlock()
    this.Grpinfo = fi
    this.Cttype = CTTYPE_GROUP
    this.Ctname = fi.Title
    this.Ctstmsg = fi.Stmsg
    this.Ctnum = fi.Gnum
    return

proc Switchtoct(mdl:DataModel, uniqid:string) =
    mdl.Ctuniqid = uniqid
    for k, v in mdl.Groupsm:
        if v.GroupId == uniqid:
            mdl.setGroupInfo(v)
            return
    for k, v in mdl.Friendsm:
        if v.Pubkey == uniqid:
            mdl.setFriendInfo(v)
            return
    return

const maxinmemmsgcnt = 5000
proc Newmsg(this:DataModel, uniqid: string, msg : Message) =
    var zeromsgs : seq[Message]
    if not this.Ctmsgs.hasKey(uniqid): this.Ctmsgs.add(uniqid, zeromsgs)
    if not this.Hasnews.hasKey(uniqid): this.Hasnews.add(uniqid, 0)

    this.Ctmsgs[uniqid].add(msg)
    this.Hasnews[uniqid] += 1

    if uniqid != this.Ctuniqid:
        this.lastmsg = msg
        if this.Groupsm.hasKey(uniqid):
            var cto = this.Groupsm[uniqid]
            this.lastctname = cto.Title
        elif this.Friendsm.hasKey(uniqid):
            var cto = this.Friendsm[uniqid]
            this.lastctname = cto.Name
        else: discard

    return

proc Lastmsg(this:DataModel): string =
    var msgo = this.lastmsg
    if msgo == nil: return ""
    return this.lastctname & "$ " & msgo.PeerNameUi & msgo.MsgUi

proc Hasnewmsg(this:DataModel, uniqid:string):bool =
    if not this.Hasnews.hasKey(uniqid): return false
    return this.Hasnews[uniqid] > 0

proc Unsetnewmsg(this:DataModel, uniqid: string) =
    if not this.Hasnews.hasKey(uniqid): return
    this.Hasnews[uniqid] = 0
    return

proc NewMsgcount(this:DataModel, uniqid:string) : int =
    if not this.HasNews.haskey(uniqid): return 0
    return this.HasNews[uniqid]

proc Msgcount(this:DataModel, uniqid:string):int =
    if not this.Ctmsgs.hasKey(uniqid): return 0
    return this.Ctmsgs[uniqid].len()

proc TotalCurrMsgCount(this:DataModel) : (int, int) =
    var cur = 0
    var tot = 0
    if this.Ctmsgs.hasKey(this.Ctuniqid): cur = this.Ctmsgs[this.Ctuniqid].len()
    for k, v in this.Ctmsgs: tot += v.len()
    return (cur, tot)

# // like: limit m, offset n
# func (this *DataModel) Getmsgs(uniqid string, limit int, start ...int) {

# }

proc GetNewestMsgs(this:DataModel, uniqid:string, limit:int) : seq[Message] =
    var zeromsgs : seq[Message]
    if not this.Ctmsgs.hasKey(uniqid): return zeromsgs

    var msgs = this.Ctmsgs[uniqid]
    var totcnt = msgs.len()

    var rets :seq[Message]
    var startpos = max(0, totcnt - 1 - limit)
    for idx in startpos..totcnt-1:
        rets.add(msgs[idx])

    return rets

