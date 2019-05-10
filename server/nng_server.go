package server

import (
	"context"
	"encoding/json"
	"gopp"
	"log"
	"mkuse/go-nng"
	"tox-homeserver/thspbs"
	"unsafe"
)

// req/req
// pub/sub
type NNGServer struct {
	pubsk nng.Nng_socket
	repsk nng.Nng_socket
}

func newNNGServer() *NNGServer {
	this := &NNGServer{}
	return this
}

func testhhh() {
	var sock nng.Nng_socket
	r := nng.Nng_rep0_open(&sock)
	log.Println(r)
}

func (this *NNGServer) Setup() error {
	// althrough it listen, but not block
	r := nng.Nng_pub0_open(&this.pubsk)
	log.Println(r)
	r = nng.Nng_listen(this.pubsk, "tcp://0.0.0.0:2081", nil, 0)
	log.Println(r)

	r = nng.Nng_rep0_open(&this.repsk)
	log.Println(r)
	r = nng.Nng_listen(this.repsk, "tcp://0.0.0.0:2081", nil, 0)
	log.Println(r)
	return nil
}
func (this *NNGServer) Stop() error {
	return nil
}
func (this *NNGServer) LoopCall() {
	go this.repproc()
}

func (this *NNGServer) repproc() {
	var rbuf = make([]byte, 512)
	var rblen int
	for {
		r := nng.Nng_recv(this.repsk, unsafe.Pointer(&rbuf[0]), (*nng.Size_t)(unsafe.Pointer(&rblen)), 0)
		log.Println(r, rblen)

		req := &thspbs.Event{}
		err := json.Unmarshal(rbuf[:rblen], req)
		gopp.ErrPrint(err, string(rbuf[:rblen]))

		rsp, err := RmtCallHandlers(context.Background(), req)
		gopp.ErrPrint(err)
		rspcc, err := json.Marshal(rsp)
		gopp.ErrPrint(err)
		r = nng.Nng_send(this.repsk, unsafe.Pointer(&rspcc[0]), uint64(len(rspcc)), 0)
		log.Println(r, len(rspcc))
	}
}

func (this *NNGServer) Pubmsg(ctx context.Context, evt *thspbs.Event) error {
	bcc, err := json.Marshal(evt)
	gopp.ErrPrint(err)

	r := nng.Nng_send(this.repsk, unsafe.Pointer(&bcc[0]), uint64(len(bcc)), 0)
	log.Println(r, len(bcc))
	return err
}
