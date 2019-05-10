package server

import (
	"flag"
	"fmt"
	"gopp"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"tox-homeserver/store"
	"tox-homeserver/thscom"

	"github.com/envsh/go-toxcore/xtox"
	"github.com/google/gops/agent"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

type appContext struct {
	tvm    *ToxVM
	rpcs   *GrpcServer
	wssrv  *WebsocketServer
	nngsrv *NNGServer
	st     *store.Storage

	brker *msgbroker
}

var appctx = &appContext{brker: &msgbroker{}}

func Main() {
	log.Println(BuildInfo)
	flag.Parse()
	vcfg := GetCfg()
	vcfg.Set("toxcore_version", xtox.VersionStr())
	vcfg.Set("toxhs_version", Version)
	vcfg.Set("build_info", BuildInfo)
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalln(err)
	}

	appctx.st = store.NewStorage()
	thscom.SetLogMetrics()
	addFileHelperAsContact(appctx.st)

	log.Println("Loading unsents messages...")
	OffMsgMan().LoadFromStorage(appctx.st)

	go func() {
		// 为简单debug,stats,socketio,websocket使用
		wso := NewWebsocketServer()
		appctx.wssrv = wso
		store.GetFS().SetupHttpServer(nil)
		log.Printf("Listen on WS: *:%d ..., %#v\n", thscom.WSPort, wso)
		err := http.ListenAndServe(fmt.Sprintf(":%d", thscom.WSPort), nil)
		gopp.ErrPrint(err)
	}()
	nngsrv := newNNGServer()
	appctx.nngsrv = nngsrv
	nngsrv.Setup()
	nngsrv.LoopCall()

	rpcs := newGrpcServer()
	appctx.rpcs = rpcs

	appctx.tvm = newToxVM()

	go xtox.Run(appctx.tvm.t, appctx.tvm.tav)
	time.Sleep(50 * time.Millisecond)
	rpcs.run()
}

func Atexit() {
	tvm := appctx.tvm
	if tvm != nil {
		if tvm.tav != nil {
			log.Println("Stop toxav...")
			tvm.tav.Kill()
		}
		if tvm.t != nil {
			log.Println("Stop tox...")
			tvm.t.Kill()
		}
	}
}

func addFileHelperAsContact(st *store.Storage) {
	st.AddFriend(thscom.FileHelperPk, thscom.FileHelperFnum, thscom.FileHelperName, "")
}

///// build info
var Version string
var BuildInfo string

func SetBuildInfo(version, info string) { Version, BuildInfo = version, info }
