package server

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 进程内的

type PubSubPool struct {
	pname  string
	mu     sync.RWMutex
	vconns map[string]*PSConn
}

func (this *PubSubPool) get(vaddr string) *PSConn {
	this.mu.RLock()
	defer this.mu.RUnlock()
	if pso, ok := this.vconns[vaddr]; ok {
		return pso
	}
	return nil
}

func (this *PubSubPool) put(vaddr string, pso *PSConn) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.vconns[vaddr] = pso
}

func (this *PubSubPool) del(vaddr string) {
	this.mu.Lock()
	defer this.mu.Unlock()
	delete(this.vconns, vaddr)
}

func (this *PubSubPool) each(f func(*PSConn)) {
	this.mu.RLock()
	defer this.mu.Unlock()
	cs := []*PSConn{}
	for _, c := range this.vconns {
		cs = append(cs, c)
	}

	// this.mu.Unlock()
	for _, c := range cs {
		f(c)
	}
}

var psp = &PubSubPool{pname: "default", vconns: map[string]*PSConn{}}

type PSConn struct {
	subfns map[string]func(interface{})
	subchs map[string]chan interface{}
	pscno  string
}

var cpscno uint64

func connps() *PSConn {
	pscno := atomic.AddUint64(&cpscno, 1)
	s := fmt.Sprintf("%v", pscno)
	c := &PSConn{}
	c.pscno = s
	c.subfns = map[string]func(interface{}){}
	c.subchs = map[string]chan interface{}{}
	psp.put(s, c)
	return c
}

// TODO
func connpsWithName(vname string) *PSConn { return nil }

func (this *PSConn) Publish(topic string, m interface{}) {
	psp.each(func(c *PSConn) {
		c.dispatch(this, topic, m)
	})
}

// 不要在sub回调函数中执行pub操作，会导致死锁
func (this *PSConn) Subscribe(topic string, subfn func(interface{})) {
	this.subfns[topic] = subfn
}

func (this *PSConn) ChanSubscribe(topic string, subch chan interface{}) {
	this.subchs[topic] = subch
}

func (this *PSConn) ChanSubscribe2(topic string) (subch chan interface{}) {
	subch = make(chan interface{}, 0)
	this.subchs[topic] = subch
	return
}

func (this *PSConn) dispatch(puber *PSConn, topic string, m interface{}) {
	if subfn, ok := this.subfns[topic]; ok {
		subfn(m)
	}
	if subch, ok := this.subchs[topic]; ok {
		subch <- m
	}
}

func (this *PSConn) close() {
	psp.del(this.pscno)
}
