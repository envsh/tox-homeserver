

const CTTYPE_FRIEND = 1
const CTTYPE_GROUP = 2

import tables

const MemberInfo_UNKNOWN  = 0
const MemberInfo_FRIEND = 1
const MemberInfo_GROUP = 2
const MemberInfo_PEER = 3

type
    MemberInfo = ref object
        Pnum*: uint32
        Pubkey* : string
        Name* : string
        Mtype* : int
        # // extra
        Joints* : int64

type
    FriendInfo = ref object
        Fnum* : uint32
        Status1* : uint32 # 为啥叫Status编译器报错
        Pubkey* : string
        Name* : string
        Stmsg* : string
        Avatar* : string
        Seen* : uint64
        ConnStatus* : int32

type
    GroupInfo = ref object
        Gnum* : uint32
        Mtype* : uint32
        GroupId* : string
        Title* : string
        Stmsg* : string
        Ours* : bool
        Members* : Table[string,MemberInfo]


type
    BaseInfo = ref object
        ToxId*:string
        Name*: string
        Stmsg*: string
        Status*: uint32
        Friend*: Table[uint32,FriendInfo]
        Groups*: Table[uint32,GroupInfo]
        ConnStatus*: int32
        NextBatch*: int64
        ToxVersion*: string

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
        # Frndinfo thspbs.FriendInfo
        # Grpinfo thspbs.GroupInfo
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
    return s.substr(0, 1)

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

