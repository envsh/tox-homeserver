package main

import (
	"context"
	"fmt"
	"gopp"
	"log"
	"tox-homeserver/common"
	"tox-homeserver/thspbs"

	"github.com/nats-io/nats"

	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

func main() {
	target := fmt.Sprintf("%s:%d", common.GrpcIp, common.GrpcPort)
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	gopp.ErrPrint(err)
	thsc := thspbs.NewToxhsClient(cc)
	in := &thspbs.EmptyReq{}
	info, err := thsc.GetBaseInfo(context.Background(), in)
	gopp.ErrPrint(err, info)
	log.Println(info, len(info.Friends))

	natsAddr := fmt.Sprintf("nats://%s:%d", common.GnatsIp, common.GnatsPort)
	log.Println("connecting...", natsAddr)
	nc, err := nats.Connect(natsAddr)
	gopp.ErrPrint(err, nc, nats.DefaultURL)

	nc.Subscribe(common.CBEventBusName, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	select {}
}
