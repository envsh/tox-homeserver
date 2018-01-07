package common

import (
	"log"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
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
var logMet1 sync.Once

var ffn = func(n float64) string { return humanize.FormatFloat("#.##", n) }

func logMetFancy() {
	recvMeter := metrics.GetOrRegisterMeter(ByteRecv+msuf, MetReg)
	sentMeter := metrics.GetOrRegisterMeter(ByteSent+msuf, MetReg)
	tplstr := "meter: %s, count: %d, 1-min rate: %s, 5-min rate: %s, 15-min rate: %s, mean rate: %s\n"
	log.Printf(tplstr, ByteRecv+msuf, recvMeter.Count(),
		ffn(recvMeter.Rate1()), ffn(recvMeter.Rate5()),
		ffn(recvMeter.Rate15()), ffn(recvMeter.RateMean()))
	log.Printf(tplstr, ByteSent+msuf, sentMeter.Count(),
		ffn(sentMeter.Rate1()), ffn(sentMeter.Rate5()),
		ffn(sentMeter.Rate15()), ffn(sentMeter.RateMean()))
}

func SetLogMetrics() {
	logMet1.Do(func() {
		freq := 30 * time.Second
		// go metrics.Log(MetReg, freq, log.New(os.Stdout, LogPrefix, log.Lmicroseconds|log.Lshortfile))
		go func() {
			for {
				time.Sleep(freq)
				logMetFancy()
			}
		}()
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
