package server

import (
	"flag"
	"fmt"
	"gopp"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
	thscom "tox-homeserver/common"
	"tox-homeserver/store"

	"github.com/envsh/go-toxcore/xtox"
	"github.com/google/gops/agent"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

type appContext struct {
	tvm   *ToxVM
	rpcs  *GrpcServer
	wssrv *WebsocketServer
	st    *store.Storage
}

var appctx = &appContext{}

func Main() {
	printBuildInfo(true)
	flag.Parse()
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalln(err)
	}

	appctx.st = store.NewStorage()
	thscom.SetLogMetrics()
	go func() {
		// 为简单debug,stats,socketio,websocket使用
		wso := NewWebsocketServer()
		appctx.wssrv = wso
		log.Printf("Listen on WS: *:%d ..., %#v\n", thscom.WSPort, wso)
		err := http.ListenAndServe(fmt.Sprintf(":%d", thscom.WSPort), nil)
		gopp.ErrPrint(err)
	}()

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
