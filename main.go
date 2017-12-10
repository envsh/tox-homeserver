package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/gops/agent"
	"github.com/kitech/go-toxcore/xtox"
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

	rpcs := newGrpcServer()
	appctx.rpcs = rpcs

	appctx.tvm = newToxVM()

	go xtox.Run(appctx.tvm.t)

	time.Sleep(50 * time.Millisecond)
	rpcs.run()
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
