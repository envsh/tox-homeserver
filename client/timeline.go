package client

import (
	"context"
	"encoding/json"
	"fmt"
	"gopp"
	"log"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"
)

func (this *LigTox) LoadEventsByContactId(pubkey string, prev_batch int64) ([]store.Message, error) {
	args := &thspbs.Event{}
	args.Name = "LoadEventsByContactId"
	args.Args = []string{pubkey, fmt.Sprintf("%d", prev_batch)}

	cli := thspbs.NewToxhsClient(this.rpcli)
	rsp, err := cli.RmtCall(context.Background(), args)
	gopp.ErrPrint(err)
	if err != nil {
		return nil, err
	}
	log.Println(rsp.Args)
	rets := []store.Message{}
	err = json.Unmarshal([]byte(rsp.Args[0]), &rets)
	gopp.ErrPrint(err)
	if err != nil {
		return nil, err
	}
	return rets, nil
}
