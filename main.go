package main

import (
	"flag"
	"fmt"
	"gopp"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/google/gops/agent"
	"github.com/kitech/go-toxcore/xtox"
	gnatsd "github.com/nats-io/gnatsd/server"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

type appContext struct {
	tvm  *ToxVM
	rpcs *GrpcServer
}

var appctx = &appContext{}

func main() {
	printBuildInfo(true)
	flag.Parse()
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalln(err)
	}

	go func() {
		err := http.ListenAndServe(":8089", nil)
		gopp.ErrPrint(err)
	}()

	rpcs := newGrpcServer()
	appctx.rpcs = rpcs

	appctx.tvm = newToxVM()

	go xtox.Run(appctx.tvm.t)
	time.Sleep(50 * time.Millisecond)
	go runLocalNatsd()
	time.Sleep(50 * time.Millisecond)
	rpcs.run()
}

// should block
func runLocalNatsd() {
	fs := flag.NewFlagSet("nats-server", flag.ExitOnError)
	opts, err := gnatsd.ConfigureOptions(fs, []string{"--port", "4111"}, func() {}, func() {}, func() {})
	gopp.ErrPrint(err, "gnatsd options error")

	ndsrv := gnatsd.New(opts)
	ndsrv.ConfigureLogger()
	if err := gnatsd.Run(ndsrv); err != nil {
		gopp.ErrPrint(err, "gnatsd exited")
	}
}

// build info
var GitCommit, GitBranch, GitState, GitSummary, BuildDate, Version string

func printBuildInfo(full bool) { log.Println(getBuildInfo(full)) }
func getBuildInfo(full bool) string {
	trim := func(s string) string {
		if strings.HasPrefix(s, "GOVVV-") {
			return s[6:]
		}
		return s
	}
	commit := trim(GitCommit)
	branch := trim(GitBranch)
	// state := trim(GitState)
	summary := trim(GitSummary)
	date := trim(BuildDate)
	version := trim(Version)

	if full {
		return fmt.Sprintf("govvv: v%s branch:%s git:%s build:%s summary:%s, ",
			version, branch, commit, date, summary)
	}
	return fmt.Sprintf("govvv: v%s git:%s build:%s", version, commit, date)
}