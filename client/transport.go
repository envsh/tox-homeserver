package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"strings"
	"time"
	thscom "tox-homeserver/common"
	"tox-homeserver/thspbs"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
)

func assert_interface() {
	var tp1 *WebsocketTransport
	var tp1i Transport = tp1
	var tp2 *GrpcTransport
	var tp2i Transport = tp2
	_, _, _, _ = tp1, tp1i, tp2, tp2i
}

type Transport interface {
	Connect(host string) error
	Start() error
	Close() error
	OnData(func(*thspbs.Event, []byte))
	OnDisconnected(func(error))
	OnConnected(func())
	// WriteMessage([]byte) error
	// RecvMessage() ([]byte, error)
	rmtCall(*thspbs.Event) (*thspbs.Event, error)
}

type _TransportBase struct {
	name        string
	srvurl      string
	datacbs     []func(*thspbs.Event, []byte)
	disconncdbs []func(error)
	connedcbs   []func()
	closed      bool

	retryer *gopp.Retryer
}

func newTransportBase() *_TransportBase {
	this := &_TransportBase{}
	this.datacbs = make([]func(*thspbs.Event, []byte), 0)
	this.disconncdbs = make([]func(error), 0)
	this.connedcbs = make([]func(), 0)
	return this
}

func (this *_TransportBase) OnData(f func(*thspbs.Event, []byte)) {
	this.datacbs = append(this.datacbs, f)
}
func (this *_TransportBase) OnDisconnected(f func(err error)) {
	this.disconncdbs = append(this.disconncdbs, f)
}
func (this *_TransportBase) OnConnected(f func()) {
	this.connedcbs = append(this.connedcbs, f)
}

func (this *_TransportBase) runOnData(evto *thspbs.Event, data []byte) {
	for _, datacb := range this.datacbs {
		datacb(evto, data)
	}
}
func (this *_TransportBase) runOnDisconnected(err error) {
	for _, disconncb := range this.disconncdbs {
		disconncb(err)
	}
}
func (this *_TransportBase) runOnConnected() {
	for _, connedcb := range this.connedcbs {
		connedcb()
	}
}

const sock_crt = "data/server.crt"

/////
type GrpcTransport struct {
	*_TransportBase

	rpcli *grpc.ClientConn
}

func NewGrpcTransport() *GrpcTransport {
	this := &GrpcTransport{}
	this._TransportBase = newTransportBase()
	return this
}

func (this *GrpcTransport) Close() error {
	this.closed = true
	this.connedcbs = nil
	rpcli := this.rpcli
	this.rpcli = nil
	return rpcli.Close()
}

// addr: host:port
func (this *GrpcTransport) Connect(addr string) error {
	this.srvurl = addr
	srvurl := addr
	log.Println("connecting grpc:", srvurl)

	kaopt := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                75 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	})

	// TODO optional tls flag
	certPEMBlock, err := Asset(sock_crt)
	gopp.ErrPrint(err)

	certp := x509.NewCertPool()
	if ok := certp.AppendCertsFromPEM(certPEMBlock); !ok {
		gopp.FalsePrint(ok, "cert pool add pem csr error", ok, len(certPEMBlock))
	}
	// credo := credentials.NewClientTLSFromCert(certp, strings.Split(addr, ":")[0])
	// credo := credentials.NewClientTLSFromCert(certp, "")
	credo := credentials.NewTLS(&tls.Config{ServerName: "", RootCAs: certp, InsecureSkipVerify: true})
	certopt := grpc.WithTransportCredentials(credo)

	rpcli, err := grpc.Dial(srvurl /*grpc.WithInsecure(),*/, grpc.WithTimeout(5*time.Second), kaopt, certopt)
	gopp.ErrPrint(err, rpcli)
	if err != nil {
		return err
	}

	// ping test, seems grpc is lazy connect
	cc := rpcli
	thsc := thspbs.NewToxhsClient(cc)
	in := &thspbs.EmptyReq{}
	_, err = thsc.Ping(context.Background(), in)
	gopp.ErrPrint(err)
	if err != nil {
		return err
	}
	this.rpcli = rpcli

	return nil
}

func (this *GrpcTransport) Start() error {
	go this.serveBackendEventGrpc()
	return nil
}

// should block
func (this *GrpcTransport) serveBackendEventGrpc() {
	nowt := time.Now()
	for !this.closed {
		this.serveBackendEventGrpcImpl()
		if this.closed {
			break
		}

		if this.retryer == nil {
			this.retryer = gopp.NewRetry()
		}
		retryWait := 3*time.Second + this.retryer.NextWaitOnly()
		log.Println("Grpc maybe disconnected, retry after", retryWait)
		// TODO for android, 需要在从休眠中醒来时通知并取消该sleep
		// TODO for android, 也许需要监听wifi状态
		time.Sleep(retryWait)
	}
	log.Println("Grpc serve proc done:", this.closed, time.Since(nowt))
}

func (this *GrpcTransport) serveBackendEventGrpcImpl() {
	clio := thspbs.NewToxhsClient(this.rpcli)
	stmc, err := clio.PollCallback(context.Background(),
		&thspbs.Event{Name: "PollCallback", DeviceUuid: appctx.devo.Uuid})
	gopp.ErrPrint(err)
	if err != nil {
		return
	}

	// success reset
	this.retryer = nil
	cnter := uint64(0)
	for !this.closed {
		evto, err := stmc.Recv()
		gopp.ErrPrint(err)
		if err != nil {
			break
		}
		cnter++

		jcc, err := json.Marshal(evto)
		gopp.ErrPrint(err)
		if jcc == nil {
			log.Println("Wtf:", evto)
			continue
		}

		if strings.Contains(string(jcc), "AudioReceiveFrame") {
			// log.Println("grpcrecv:", "AudioReceiveFrame", len(jcc))
		} else if strings.Contains(string(jcc), "VideoReceiveFrame") {
			// log.Println("grpcrecv:", "VideoReceiveFrame", len(jcc))
		} else if strings.Contains(string(jcc), "ConferenceAudioRecieiveFrame") {
		} else {
			log.Println("grpcrecv:", string(jcc))
		}
		this.runOnData(evto, jcc)
	}
	log.Println("Grpc poll got events:", cnter)
}

func (this *GrpcTransport) rmtCall(args *thspbs.Event) (*thspbs.Event, error) {
	cli := thspbs.NewToxhsClient(this.rpcli)
	rsp, err := cli.RmtCall(context.Background(), args, grpc.UseCompressor("gzip"))
	return rsp, err
}

/////
type WebsocketTransport struct {
	*_TransportBase

	wsc4rpc  *websocket.Conn
	wsc4push *websocket.Conn
}

func NewWebsocketTransport() *WebsocketTransport {
	this := &WebsocketTransport{}
	this._TransportBase = newTransportBase()
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
	for !this.closed {
		err = this.serveBackendEventWSImpl()
		for retry := 1; ; retry++ {
			log.Println("Websocket maybe disconnect, retry 3 secs...")
			time.Sleep(time.Duration(retry+retry) * time.Second)
			err = this.Connect(this.srvurl)
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
	for !this.closed {
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
			this.runOnData(evto, rmessage)
		} else if _, ok := jso.CheckGet("Name"); ok {
			this.runOnData(evto, message)
		} else {
			log.Println("Unknown packet:", string(message))
		}
	}
	log.Println("Websocket poll done")
	return errtop
}

func (this *WebsocketTransport) rmtCall(args *thspbs.Event) (*thspbs.Event, error) {
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
