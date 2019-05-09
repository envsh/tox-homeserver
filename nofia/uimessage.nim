
# TODO and other refmter
proc refmtmsgRUser(this:Message) =
    return

proc refmtmsg(msg:Message) =
    msg.MsgUi = msg.Msg
    msg.PeerNameUi = msg.PeerName
    msg.TimeUi = "hehhe"
    msg.LastMsgUi = msg.Msg

    for fn in [refmtmsgRUser]: msg.fn()
    return

proc NewMessageForGroup(evto : Event) : Message =
    var groupId = evto.Args[3]
    var message = evto.Args[3]
    var peerName = evto.Margs[0]
    var groupTitle = evto.Margs[2]
    var peerId = evto.Margs[1]
    var eventId = parseInt(evto.Margs[4])

    var this = new(Message)
    this.Msg = message
    this.PeerName = peerName
    this.Time = "" # time.Now()
    this.EventId = eventId
    if peerName == "": this.PeerName = peerid.substr(0, 8)
    this.refmtmsg()
    return this

