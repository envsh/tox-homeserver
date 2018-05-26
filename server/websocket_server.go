package server

import (
	"context"
	"encoding/json"
	"gopp"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
	"tox-homeserver/thspbs"

	"github.com/gorilla/websocket"
)

// use default options
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebsocketServer struct {
	conns   map[string]*websocket.Conn
	connsmu sync.Mutex
}

func NewWebsocketServer() *WebsocketServer {
	this := &WebsocketServer{}
	this.conns = make(map[string]*websocket.Conn)

	this.initHandler()
	return this
}

func (this *WebsocketServer) initHandler() {
	http.HandleFunc("/toxhsrpc", this.toxhsrpc)
	http.HandleFunc("/toxhspush", this.toxhspush)
	http.HandleFunc("/webdui/", this.webdui)
	http.HandleFunc("/webdui", this.webdui)
	http.HandleFunc("/echo", this.echo)
}

func (this *WebsocketServer) toxhsrpc(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err, r.URL.String())
		return
	}
	ctime := time.Now()
	raddr := c.RemoteAddr().String()
	defer func() {
		c.Close()
	}()

	// TODO 保证Req/Resp的序列？
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s of %d\n", message, mt)
		req := &thspbs.Event{}
		err = json.Unmarshal(message, req)
		gopp.ErrPrint(err, string(message))

		rsp, err := RmtCallHandlers(context.Background(), req)
		gopp.ErrPrint(err)
		rspcc, err := json.Marshal(rsp)
		gopp.ErrPrint(err)
		err = c.WriteMessage(mt, rspcc)
		gopp.ErrPrint(err)

	}
	log.Println("disconnected from:", raddr, time.Now().Sub(ctime))
}

func (this *WebsocketServer) toxhspush(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err, r.URL.String())
		return
	}
	ctime := time.Now()
	raddr := c.RemoteAddr().String()
	connkey := raddr + ctime.String()
	this.connsmu.Lock()
	this.conns[connkey] = c
	this.connsmu.Unlock()
	defer func() {
		c.Close()
		this.connsmu.Lock()
		delete(this.conns, connkey)
		this.connsmu.Unlock()
	}()

	// TODO 保证Req/Resp的序列？
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s of %d\n", message, mt)
	}
	log.Println("disconnected from:", raddr, time.Now().Sub(ctime))
}

func (this *WebsocketServer) pushevt(evt *thspbs.Event) error {
	rspcc, err := json.Marshal(evt)
	gopp.ErrPrint(err)

	this.connsmu.Lock()
	defer this.connsmu.Unlock()
	for _, c := range this.conns {
		err = c.WriteMessage(websocket.TextMessage, rspcc)
		gopp.ErrPrint(err)
	}
	return nil
}

func (this *WebsocketServer) webdui(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.String())
	bcc, err := ioutil.ReadFile("./webdui/index.html")
	gopp.ErrPrint(err)
	if err != nil {
		bcc, err = Asset("webdui/index.html")
		gopp.ErrPrint(err)
	}
	w.Write(bcc)
}

func (this *WebsocketServer) echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s of %d\n", message, mt)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
