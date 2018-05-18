package main

import (
	"flag"
	"gopp"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"runtime/trace"
	"syscall"

	_ "net/http/pprof"

	"github.com/google/gops/agent"
)

const (
	ltracep   = "trace: "
	ldebugp   = "debug: "
	linfop    = "info: "
	lwarningp = "warning: "
	lerrorp   = "error: "
	lalertp   = "alert: "
)

// replacement: dynamic set cpu pprof and trace with gops
var cpuprofile string // = filepath.Base(os.Args[0])
var tracefile string  //= "trace.out"

func init() {
	flag.StringVar(&cpuprofile, "pprof", cpuprofile, "enable CPU pprof, supply file name.")
	flag.StringVar(&tracefile, "trace", cpuprofile, "trace, supply file name.")
}

func SetupProfile() {
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}
}

func StopProfile() {
	if cpuprofile != "" {
		log.Println("save profile", cpuprofile)
		pprof.StopCPUProfile()
	}
}

func SetupTrace() {
	if tracefile != "" {
		f, err := os.Create("trace.out")
		gopp.ErrPrint(err)
		trace.Start(f)
		log.Println("traceing to trace.out...")
	}
}

func StopTrace() {
	if tracefile != "" {
		trace.Stop()
	}
}

// should block
func SetupSignal(exitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-c:
			switch s {
			case syscall.SIGINT:
				log.Println("exiting...", s)
				// os.Exit(0) // will not run defer
				if exitFunc != nil {
					exitFunc()
				}
				return // will run defer
			default:
				log.Println("unprocessed signal:", s)
			}
		}
	}
}

func SetupGops() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
}
