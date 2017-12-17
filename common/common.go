package common

import (
	"fmt"
	"runtime"
)

const CBEventBusName = "cbevt"

// const GrpcIp = "127.0.0.1"
var GrpcIp = "10.0.0.31"

const GrpcPort = uint16(2080)
const GnatsIp = "10.0.0.6"
const GnatsPort = uint16(4222)

var GrpcAddr = fmt.Sprintf("%s:%d", GrpcIp, GrpcPort)
var GnatsAddr = fmt.Sprintf("nats://%s:%d", GnatsIp, GnatsPort)

const DefaultUserName = "Tox User"
const GroupTitleSep = " ::: " // group title and group stmsg title, "@title@ ::: @stmsg@"

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
