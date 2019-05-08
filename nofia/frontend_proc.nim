
import json

type Event = ref object
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


proc dispatchEvent(evt: Event) =
    if evt.Name == "ConferenceMessage": discard
    else: discard
    return

proc dispatchEventNim(evtdat : cstring) {.exportc.} =
    let s : string = $evtdat
    let jsonNode = parseJson(s)

    var evt = Event()
    try: evt.EventId = jsonNode["EventId"].getInt()
    except: discard
    try: evt.Name = jsonNode["Name"].getStr()
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

    dispatchEvent(evt)
    return

proc dispatchEventRespNim(evtdat : cstring) {.exportc.} =
    linfo("evtdat", evtdat.len())
    return
