

const CTTYPE_FRIEND = 1
const CTTYPE_GROUP = 2

import tables

type
    PDataModel = ptr DataModel
    DataModel = ref object
        # mu
        Myid*: string
        Myname*: string
        Mystmsg*: string
        Mystno*: int
        Mysttxt*: string #  // status text
        Lastno*: int # // last valid status no, in case rpc

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

        Friendsm*: Table[string,pointer] # *thspbs.FriendInfo // uniqid =>
        Friendsv*: seq[pointer] #
        Groupsm*: Table[string,pointer] # *thspbs.GroupInfo // uniqid =>
        Groupsv*: seq[pointer] #

        Ctmsgs*: Table[string,seq[Message]] #// uniqid =>
        Hasnews*: Table[string,int] # // uniqid => , 某个联系人的未读取消息个数
        lastmsg*: Message # // Lastmsg must be not belongs to active contact chatform
        lastctname*: string # // always according with lastmsg

        repainter*: proc()

