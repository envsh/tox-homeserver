
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
        LastMsgUi*: string
        Sent*: bool
        UserCode*: int64

        Links*: seq[string]
        Index*: int64


