
import random

# TODO and other refmter
proc refmtmsgRUser(msg:Message) =
    var (sender, message, color) = ExtractRealUserMD(msg.PeerName, msg.Msg)
    #if itsfalse():
    if true:
        msg.MsgUi = message
        msg.PeerNameUi = sender
        msg.LastMsgUi = message
    return

proc refmtmsg(msg:Message) =
    msg.MsgUi = msg.Msg
    msg.PeerNameUi = msg.PeerName
    msg.TimeUi = msg.Time.totoday()
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
    this.Time = now()
    this.EventId = eventId
    if peerName == "": this.PeerName = peerid.substr(0, 8)
    this.refmtmsg()
    return this

proc NewMessageForMe(iname, itext : string) : Message =
    #var mdl = getnkmdl()

    var msgo = new(Message)
    msgo.Msg = itext
    msgo.PeerName = iname # mdl.Myname
    msgo.Time = now()
    msgo.Me = true

    msgo.refmtmsg()
    return msgo

