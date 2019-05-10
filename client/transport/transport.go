package transport

import (
	"gopp"
	"log"
	"reflect"

	"tox-homeserver/thspbs"
)

// by import _ "xxx/client/xxxtp" for register this value

var dfttp Transport

func DftTP() Transport { return dfttp }
func RegisterTransport(name string, tp Transport) {
	if dfttp != nil {
		log.Println("Duplicate register transport", reflect.TypeOf(dfttp))
		return
	}
	dfttp = tp
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
	RmtCall(*thspbs.Event) (*thspbs.Event, error)
	GetBaseInfo() *thspbs.BaseInfo // temporary
}

type TransportBase struct {
	Name        string
	DevUuid     string
	Srvurl      string
	Datacbs     []func(*thspbs.Event, []byte)
	Disconncdbs []func(error)
	Connedcbs   []func()
	Closed      bool

	Retryer *gopp.Retryer
}

func NewTransportBase() *TransportBase {
	this := &TransportBase{}
	this.Datacbs = make([]func(*thspbs.Event, []byte), 0)
	this.Disconncdbs = make([]func(error), 0)
	this.Connedcbs = make([]func(), 0)
	return this
}

func (this *TransportBase) OnData(f func(*thspbs.Event, []byte)) {
	this.Datacbs = append(this.Datacbs, f)
}
func (this *TransportBase) OnDisconnected(f func(err error)) {
	this.Disconncdbs = append(this.Disconncdbs, f)
}
func (this *TransportBase) OnConnected(f func()) {
	this.Connedcbs = append(this.Connedcbs, f)
}

func (this *TransportBase) RunOnData(evto *thspbs.Event, data []byte) {
	for _, datacb := range this.Datacbs {
		datacb(evto, data)
	}
}
func (this *TransportBase) runOnDisconnected(err error) {
	for _, disconncb := range this.Disconncdbs {
		disconncb(err)
	}
}
func (this *TransportBase) runOnConnected() {
	for _, connedcb := range this.Connedcbs {
		connedcb()
	}
}
