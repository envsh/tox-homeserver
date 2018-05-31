package common

import (
	"fmt"
	"log"
	"runtime"
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
const GroupTitleSep = " ::: " // group title and group stmsg title, "@title@ ::: @stmsg@"

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
