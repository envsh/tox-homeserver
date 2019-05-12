package grpctp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"gopp"
	"log"
	"strings"
	"time"

	"tox-homeserver/client/transport"
	"tox-homeserver/thscom"
	"tox-homeserver/thspbs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

func init() {
	tp := NewGrpcTransport()
	transport.RegisterTransport("grpc", tp)
}

/////
type GrpcTransport struct {
	*transport.TransportBase

	rpcli *grpc.ClientConn
}

func NewGrpcTransport() *GrpcTransport {
	this := &GrpcTransport{}
	this.TransportBase = transport.NewTransportBase()
	return this
}

func (this *GrpcTransport) Close() error {
	this.Closed = true
	this.Connedcbs = nil
	rpcli := this.rpcli
	this.rpcli = nil
	return rpcli.Close()
}

// addr: host:port
func (this *GrpcTransport) Connect(addr string) error {
	this.Srvurl = addr
	srvurl := addr
	log.Println("connecting grpc:", srvurl)

	kaopt := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                75 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	})

	// TODO optional tls flag
	certPEMBlock, err := transport.Asset(thscom.TPCertFile)
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
	for !this.Closed {
		err := this.serveBackendEventGrpcImpl()
		gopp.ErrPrint(err)
		if this.Closed {
			break
		}

		if this.Retryer == nil {
			this.Retryer = gopp.NewRetry()
		}
		retryWait := 3*time.Second + this.Retryer.NextWaitOnly()
		maxWait := 51 * time.Second
		retryWait = gopp.IfElse(retryWait > maxWait, maxWait, retryWait).(time.Duration)
		log.Println("Grpc maybe disconnected, retry after", retryWait)
		this.notifyRpcStatus(err, retryWait)
		// TODO for android, 需要在从休眠中醒来时通知并取消该sleep
		// TODO for android, 也许需要监听wifi状态
		time.Sleep(retryWait)
	}
	log.Println("Grpc serve proc done:", this.Closed, time.Since(nowt))
}

func (this *GrpcTransport) notifyRpcStatus(err error, retryWait time.Duration) {
	evto := &thspbs.Event{}
	if err != nil {
		evto.EventName = "brokenrpc"
		evto.Args = []string{err.Error(), retryWait.String()}
	} else {
		evto.EventName = "goodrpc"
	}

	jcc, err := json.Marshal(evto)
	gopp.ErrPrint(err)
	if jcc == nil {
		log.Println("Wtf:", evto)
	}
	this.RunOnData(evto, jcc)
}

func (this *GrpcTransport) serveBackendEventGrpcImpl() error {
	clio := thspbs.NewToxhsClient(this.rpcli)
	stmc, err := clio.PollCallback(context.Background(),
		&thspbs.Event{EventName: "PollCallback", DeviceUuid: this.DevUuid})
	gopp.ErrPrint(err)
	if err != nil {
		return err
	}
	this.notifyRpcStatus(nil, 0)

	var toperr error
	// success reset
	this.Retryer = nil
	cnter := uint64(0)
	for !this.Closed {
		evto, err := stmc.Recv()
		gopp.ErrPrint(err)
		if err != nil {
			toperr = err
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
		this.RunOnData(evto, jcc)
	}
	log.Println("Grpc poll got events:", cnter)
	return toperr
}

func (this *GrpcTransport) RmtCall(args *thspbs.Event) (*thspbs.Event, error) {
	cli := thspbs.NewToxhsClient(this.rpcli)
	rsp, err := cli.RmtCall(context.Background(), args, grpc.UseCompressor("gzip"))
	return rsp, err
}

func (this *GrpcTransport) GetBaseInfo() *thspbs.BaseInfo {
	cli := thspbs.NewToxhsClient(this.rpcli)
	in := &thspbs.Event{EventName: "GetBaseInfo", DeviceUuid: this.DevUuid}
	bi, err := cli.GetBaseInfo(context.Background(), in)
	gopp.ErrPrint(err)

	return bi
}
