package client

import (
	"fmt"
	"hash/crc64"
	"sync/atomic"
	"time"
)

var userCodeSeq uint64 = 0
var sumtab *crc64.Table = crc64.MakeTable(uint64(time.Now().UnixNano()))

func NextUserCode(clinfo string) int64 {
	v := crc64.Checksum([]byte(fmt.Sprintf("%s-%d", clinfo, atomic.AddUint64(&userCodeSeq, 1))), sumtab)
	return int64(v)
}
