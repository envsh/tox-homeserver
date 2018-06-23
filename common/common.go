package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"runtime"
	"strings"
	"time"
)

const CBEventBusName = "cbevt"

// const GrpcIp = "127.0.0.1"
var GrpcIp = "10.0.0.31"

const GrpcPort = uint16(2080)

const WSPort = uint16(8099)

var GrpcAddr = fmt.Sprintf("%s:%d", GrpcIp, GrpcPort)
var WSAddr = fmt.Sprintf("ws://%s:%d", GrpcIp, WSPort)
var WSAddrlo = fmt.Sprintf("ws://%s:%d", "127.0.0.1", WSPort)

const DefaultUserName = "ToxUser"
const GroupTitleSep = "-:::-" // group title and group stmsg title, "@title@-:::-@stmsg@"

const LogPrefix = "[gofiat] "

func init() {

	switch runtime.GOOS {
	case "android":
		GrpcIp = "10.0.0.31"
		GrpcAddr = fmt.Sprintf("%s:%d", GrpcIp, GrpcPort)
	case "linux":
		fallthrough
	default:
		GrpcIp = "127.0.0.1"
		GrpcAddr = fmt.Sprintf("%s:%d", GrpcIp, GrpcPort)

	}
}

func init() { log.SetPrefix(LogPrefix) }

const PullPageSize = 20
const DefaultTimeLayout = "2006-01-02 15:04:05.999999999 -0700 MST"

func NowTimeStr() string { return time.Now().Format(DefaultTimeLayout) }

const MaxMessageLen = 1372
const MaxImageLen = 1024 * 1024 * 50 // 50M???
const MaxAutoRecvFileSize = MaxImageLen

const UiNameLen = 32
const UiStmsgLen = 45
const MaxOfflineMessageTTL = 56 * 86400 * time.Second
const MaxOfflineMessageCount = 123

const HalfMaxUint32 = math.MaxUint32 / 2
const FileHelperName = "FileHelper"
const FileHelperFnum = math.MaxUint32 - 1

var FileHelperPk = md5topk(FileHelperName)

func md5topk(s string) string {
	md5bin := md5.Sum([]byte(s))
	mymd5 := strings.ToUpper(hex.EncodeToString(md5bin[:]))
	return mymd5 + mymd5
}

func IsFixedSpecialContact(fnum uint32) bool {
	return fnum < math.MaxUint32 && fnum >= FileHelperFnum
}

// m.text, m.image, m.audio, m.video, m.location, m.emote
// m.file
const (
	MSGTYPE_TEXT     = "m.text"
	MSGTYPE_IMAGE    = "m.image"
	MSGTYPE_AVATAR   = "m.avatar"
	MSGTYPE_AUDIO    = "m.audio"
	MSGTYPE_VIDEO    = "m.video"
	MSGTYPE_LOCATION = "m.location"
	MSGTYPE_EMOTE    = "m.emote"
	MSGTYPE_FILE     = "m.file"
)
