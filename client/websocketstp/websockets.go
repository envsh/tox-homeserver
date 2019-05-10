package wstp

import (
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"strings"
	"time"

	"tox-homeserver/client/transport"
	"tox-homeserver/thscom"
	"tox-homeserver/thspbs"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

func init() {
	tp := NewWebsocketTransport()
	transport.RegisterTransport("websockets", tp)
}

/////
type WebsocketTransport struct {
	*transport.TransportBase

	wsc4rpc  *websocket.Conn
	wsc4push *websocket.Conn
}

func NewWebsocketTransport() *WebsocketTransport {
	this := &WebsocketTransport{}
	this.TransportBase = transport.NewTransportBase()
	return this
}

func (this *WebsocketTransport) Close() error {
	return nil
}

// addr: host:port
func (this *WebsocketTransport) Connect(addr string) (err error) {
	srvurl := addr
	wsurl := fmt.Sprintf("ws://%s:%d", strings.Split(srvurl, ":")[0], thscom.WSPort)
	this.wsc4rpc, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/toxhsrpc", wsurl), nil)
	gopp.ErrPrint(err, wsurl)
	this.wsc4push, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s/toxhspush", wsurl), nil)
	gopp.ErrPrint(err, wsurl)
	return
}

func (this *WebsocketTransport) Start() error {
	go this.serveBackendEventWS()
	return nil
}

// should block
func (this *WebsocketTransport) serveBackendEventWS() {
	var err error
	for !this.Closed {
		err = this.serveBackendEventWSImpl()
		for retry := 1; ; retry++ {
			log.Println("Websocket maybe disconnect, retry 3 secs...")
			time.Sleep(time.Duration(retry+retry) * time.Second)
			err = this.Connect(this.Srvurl)
			gopp.ErrPrint(err, "ws reconnect error")
			if err == nil {
				log.Println("Websocket reconnect success.")
				break
			}
			if err != nil && retry > 10000 {
				goto funcend
			}

		}
	}
funcend:
	log.Println("Websocket given up!!!")
}
func (this *WebsocketTransport) serveBackendEventWSImpl() error {
	var errtop error
	for !this.Closed {
		c := this.wsc4push
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			errtop = err
			break
		}
		log.Printf("wsrecv: %s\n", message)

		jso, err := simplejson.NewJson(message)
		gopp.ErrPrint(err)
		evto := &thspbs.Event{}
		err = json.Unmarshal(message, evto)
		gopp.ErrPrint(err)
		if rdatao, ok := jso.CheckGet("data"); ok {
			rmessage, _ := rdatao.Encode()
			this.RunOnData(evto, rmessage)
		} else if _, ok := jso.CheckGet("Name"); ok {
			this.RunOnData(evto, message)
		} else {
			log.Println("Unknown packet:", string(message))
		}
	}
	log.Println("Websocket poll done")
	return errtop
}

func (this *WebsocketTransport) RmtCall(args *thspbs.Event) (*thspbs.Event, error) {
	data, err := json.Marshal(args)
	gopp.ErrPrint(err)
	err = this.wsc4rpc.WriteMessage(websocket.TextMessage, data)
	gopp.ErrPrint(err)
	mt, rdata, err := this.wsc4rpc.ReadMessage()
	gopp.ErrPrint(err, mt)
	rsp := &thspbs.Event{}
	err = json.Unmarshal(rdata, rsp)
	gopp.ErrPrint(err)
	return rsp, err
}

func (this *WebsocketTransport) GetBaseInfo() *thspbs.BaseInfo {
	var binfo *thspbs.BaseInfo

	in := &thspbs.Event{}
	in.Name = "GetBaseInfo"
	in.DeviceUuid = this.DevUuid
	rsp, err := this.RmtCall(in)
	gopp.ErrPrint(err)

	binfo = &thspbs.BaseInfo{}
	err = json.Unmarshal([]byte(rsp.Args[0]), binfo)
	gopp.ErrPrint(err)

	return binfo
}
