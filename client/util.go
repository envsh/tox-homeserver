package client

import (
	"fmt"
	"hash/crc64"
	"strings"
	"sync/atomic"
	"time"

	"tox-homeserver/thscom"
)

var userCodeSeq uint64 = 0
var sumtab *crc64.Table = crc64.MakeTable(uint64(time.Now().UnixNano()))

func NextUserCode(clinfo string) int64 {
	v := crc64.Checksum([]byte(fmt.Sprintf("%s-%d", clinfo, atomic.AddUint64(&userCodeSeq, 1))), sumtab)
	return int64(v)
}

func HttpUrl() string {
	srvurl := appctx.vtcli.srvurl
	hurl := fmt.Sprintf("http://%s:%d", strings.Split(srvurl, ":")[0], thscom.WSPort)
	return hurl
}
func HttpFsUrl() string {
	return fmt.Sprintf("%s/toxhsfs", HttpUrl())
}
func HttpFsUrlFor(name string) string {
	return fmt.Sprintf("%s/%s", HttpFsUrl(), name)
}
func HttpFsUrlForUpload() string {
	return HttpFsUrlFor("upload?")
}
