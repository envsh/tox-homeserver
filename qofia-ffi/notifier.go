package main

import (
	"gopp"
	"log"
	"sync"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

// 单向通知。一般用于从goroutine到qt的事件
type Notifier struct {
	rb     *qtcore.QBuffer
	rbmu   sync.RWMutex
	noteFn func()
}

// 使用注意，在需要接收通知的qt线程中创建实例
func NewNotifier(f func()) *Notifier {
	this := &Notifier{}
	this.noteFn = f
	this.rb = qtcore.NewQBuffer(nil)
	ok := this.rb.Open(qtcore.QIODevice__ReadWrite | qtcore.QIODevice__Unbuffered)
	gopp.FalsePrint(ok, "wtf")

	this.setup()
	return this
}

func (this *Notifier) setup() {
	qtrt.Connect(this.rb, "readyRead()", func() {
		this.rbmu.Lock()
		this.rb.ReadAll()
		this.rbmu.Unlock()
		if this.noteFn != nil {
			this.noteFn()
		}
	})
}

func (this *Notifier) Close() {
	qtrt.Disconnect(this.rb, "readyRead()")
	this.rbmu.Lock()
	this.rb.Close()
	this.rbmu.Unlock()
}

func (this *Notifier) Trigger() {
	this.rbmu.Lock()
	n := this.rb.Write1(".")
	this.rbmu.Unlock()
	if n != 1 {
		log.Println("error trigger:", n)
	}
}
