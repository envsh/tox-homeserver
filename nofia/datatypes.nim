

# 由于类型定义的可见顺序问题，把类型统一定义在此

type
    Event = ref object
        EventId*: int64
        Name*: string
        Args*: seq[string]
        Margs*: seq[string]
        #Uargs*: pointer
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
        Time*: string # time.Time?
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

