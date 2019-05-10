package server

import (
	"context"
	"gopp"
	"tox-homeserver/thspbs"
)

type msgbroker struct{}

func (this *msgbroker) Pubmsg(ctx context.Context, evt *thspbs.Event) {
	for name, srv := range srvp {
		err := srv.Pubmsg(ctx, evt)
		gopp.ErrPrint(err, name)
	}
}

var srvp = map[string]Serveror{} // name => Serveror

func registerServeror(name string, srv Serveror) {
	srvp[name] = srv
}

// server的真正角色应该是 transport
type Serveror interface {
	Setup() error
	Stop() error

	// reqrep event handler
	LoopCall() // handler default is RmtCallHandlerRaw

	// pubsub interface
	Pubmsg(ctx context.Context, evt *thspbs.Event) error
}

var _ Serveror = &GrpcServer{}
var _ Serveror = &NNGServer{}
var _ Serveror = &WebsocketServer{}
