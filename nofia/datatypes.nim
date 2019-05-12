

# 由于类型定义的可见顺序问题，把类型统一定义在此
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
        Seen* : int64
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


# TODO 使用类型继承
type
    BaseEvent = ref object
        EventId*: int64
        Name*: string # event name

type
    BaseInfo = ref object
        EventId*: int64
        EventName*: string # event name
        ToxId*:string
        MyName*: string
        Stmsg*: string
        Status1*: uint32
        Friends*: Table[uint32,FriendInfo]
        Groups*: Table[uint32,GroupInfo]
        ConnStatus*: int32
        NextBatch*: int64
        ToxVersion*: string

type
    Argument = ref object
        FriendNumber* : uint32
        FriendName* : string
        FriendAddress* : string
        FriendPubkey* : string
        GroupNumber* : uint32
        GroupTitle* : string
        GroupIdentity* : string
        GroupCookie* : string
        GroupType* : int32
        PeerNumber* : uint32
        PeerName* : string
        PeerPubkey* : string
        MsgType* : int32
        Message* : string
        IsTyping* : int32
        MessageId* : uint32
        Status1* : int32
        StatusText* : string
        ConnStatus* : int32
        Nospam* : uint32
        FileNumber* : uint32
        FileControl* : int32
        FileIdentity* : string
        Sent* : int32
        MimeMsgType* : string
        MimeValue* : string
        NextBatch* : int64
        PrevBatch* : int64
        AudioEnabled* : int32
        VideoEnabled* : int32
        CallState* : uint32
        AudioBitRate* : uint32
        VideoBiteRate* : uint32
        Pcm* : seq[byte] # []byte
        SampleCount* : int32
        Channels* : int32
        SamplingRate* : int32
        Width* : int32
        Height* : int32
        VideoFrame* : seq[byte] # []byte
        TimeStamp* : int64

type
    Event = ref object
        EventId*: int64
        EventName*: string
        Args*: seq[string]
        Margs*: seq[string]
        Uargs*: Argument
        ErrCode*: int32
        ErrMsg*: string
        UserCode*: int64
        TrackId*: int64
        SpanId*: int64
        Json*: string
        DeviceUuid*: string

type
    PMessage = ptr Message
    Message = ref object
        Msg*: string
        PeerName*: string
        Time*: DateTime # time.Time?
        EventId*: int64

        Me*: bool
        MsgUi*: string
        PeerNameUi*: string
        TimeUi*: string
        LastMsgUi*: string
        Sent*: bool
        UserCode*: int64

        Links*: seq[string]
        Index*: int64


import ponng
type
    RpcClient = ref object
        rurl*: string
        subsk*: nng_socket
        subrfd*: AsyncFD
        subwrd*: AsyncFD
        reqsk*: nng_socket
        reqrfd*: AsyncFD
        reqwfd*: AsyncFD
        binfo*: BaseInfo
        devuuid*: string

