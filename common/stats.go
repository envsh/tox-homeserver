package common

import (
	"log"
	"os"
	"sync"
	"time"

	metrics "github.com/rcrowley/go-metrics"
)

const (
	ByteRecv  = "bytes.recv"
	ByteSent  = "bytes.sent"
	ByteTotal = "bytes.total"
	FmsgSent  = "fmsg.sent"
	FmsgRecv  = "fmsg.recv"
	FmsgTotal = "fmsg.total"
	GmsgSent  = "gmsg.sent"
	GmsgRecv  = "gmsg.recv"
	GmsgTotal = "gmsg.total"
	MsgTotal  = "msg.total"
)

var MetReg = metrics.NewRegistry()
var logMet sync.Once

func SetLogMetrics() {
	logMet.Do(func() {
		go metrics.Log(MetReg, 30*time.Second, log.New(os.Stdout, LogPrefix, log.Lmicroseconds|log.Lshortfile))
	})
}

var smpRecv = metrics.NewUniformSample(10000)
var smpSent = metrics.NewUniformSample(10000)
var smpTotal = metrics.NewUniformSample(10000)

const csuf = ".c"
const msuf = ".m"
const hsuf = ".h"

func BytesRecved(n int) {
	metrics.GetOrRegisterCounter(ByteRecv+csuf, MetReg).Inc(int64(n))
	metrics.GetOrRegisterMeter(ByteRecv+msuf, MetReg).Mark(int64(n))
	metrics.GetOrRegisterHistogram(ByteRecv+hsuf, MetReg, smpRecv).Update(int64(n))
	metrics.GetOrRegisterHistogram(ByteTotal+hsuf, MetReg, smpTotal).Update(int64(n))
}

func BytesSent(n int) {
	metrics.GetOrRegisterCounter(ByteSent+csuf, MetReg).Inc(int64(n))
	metrics.GetOrRegisterMeter(ByteSent+msuf, MetReg).Mark(int64(n))
	metrics.GetOrRegisterHistogram(ByteSent+hsuf, MetReg, smpSent).Update(int64(n))
	metrics.GetOrRegisterHistogram(ByteTotal+hsuf, MetReg, smpTotal).Update(int64(n))
}

func MsgRecved(group bool) {
	if group {
		metrics.GetOrRegisterCounter(GmsgRecv+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(GmsgTotal+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(MsgTotal+csuf, MetReg).Inc(1)
	} else {
		metrics.GetOrRegisterCounter(FmsgRecv+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(FmsgTotal+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(MsgTotal+csuf, MetReg).Inc(1)
	}
}

func MsgSent(group bool) {
	if group {
		metrics.GetOrRegisterCounter(GmsgSent+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(GmsgTotal+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(MsgTotal+csuf, MetReg).Inc(1)
	} else {
		metrics.GetOrRegisterCounter(FmsgSent+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(FmsgTotal+csuf, MetReg).Inc(1)
		metrics.GetOrRegisterCounter(MsgTotal+csuf, MetReg).Inc(1)
	}
}
