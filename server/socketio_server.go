package server

/*
基于socketio的实现，这样就能够支持浏览器与原生桌面或者原生移动app了
思路，为不同的平台实现简单的UI，只在一个窗口中接收所有的消息列表。
如果可能的话，可以提供一个回复消息的方式。
大概是，把用户名做成可点击的，然后暂存到发送消息区域，然后输入消息并发送给这个人/群。
*/

import (
	"gopp"
	"log"
	"net/http"
	"sync"

	// "github.com/googollee/go-engine.io" // checkout v1.x
	// "github.com/googollee/go-socket.io" // checkout v1.x
	"gopkg.in/googollee/go-engine.io.v1"
	"gopkg.in/googollee/go-socket.io.v1"
)

type SocketioServer struct {
	rpcsrv  *socketio.Server
	pushsrv *socketio.Server
	conns   map[string]socketio.Conn // conn id => conn
	connsmu sync.Mutex
}

func NewSocketioServer() *SocketioServer {
	this := &SocketioServer{}

	var eiopt *engineio.Options
	srv, err := socketio.NewServer(eiopt)
	gopp.ErrPrint(err)
	this.rpcsrv = srv

	srv, err = socketio.NewServer(eiopt)
	gopp.ErrPrint(err)
	this.pushsrv = srv

	this.initHandlers()
	return this
}

// not need, main.go has listened
func (this *SocketioServer) Serve() {
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (this *SocketioServer) init() {
	this.initHandlers()
}

func (this *SocketioServer) initHandlers() {
	this.initHandlerRpc()
}

func (this *SocketioServer) initHandlerRpc() {
	srv := this.rpcsrv
	// nsp: /rpc, /push,
	// conn: /
	srv.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID(), s.RemoteAddr(), s.URL(), s.Namespace(), s.RemoteHeader())
		s.Emit("hehe", "haha")
		return nil
	})
	srv.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	srv.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	srv.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	srv.OnError("/", func(e error) {
		gopp.ErrPrint(e, "meet error:")
	})
	srv.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Println("closed", msg)
	})
	go srv.Serve()
	// defer srv.Close()

	// http.Handle("/socket.io/", srv)
	http.Handle("/socket.io/", srv)
	http.Handle("/asset/", http.FileServer(http.Dir("./asset")))

}

// TODO ?
func (this *SocketioServer) initHandlerPush() {
}
