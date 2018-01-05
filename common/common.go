package common

import (
	"fmt"
	"log"
	"runtime"
)

const CBEventBusName = "cbevt"

// const GrpcIp = "127.0.0.1"
var GrpcIp = "10.0.0.31"

const GrpcPort = uint16(2080)

// const GnatsIp = "10.0.0.6"
const GnatsIp = "10.0.0.31"
const GnatsPort = uint16(4111) //uint16(4222)

var GrpcAddr = fmt.Sprintf("%s:%d", GrpcIp, GrpcPort)
var GnatsAddr = fmt.Sprintf("nats://%s:%d", GnatsIp, GnatsPort)
var GnatsAddrlo = fmt.Sprintf("nats://%s:%d", "127.0.0.1", GnatsPort)

const DefaultUserName = "Tox User"
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
