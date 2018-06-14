package avhlp

import (
	"gopp"
	"log"
	"sync"
	"time"

	"golang.org/x/mobile/exp/audio/al"
)

var openOnce sync.Once

func open() {
	openOnce.Do(func() {
		err := al.OpenDevice()
		gopp.ErrPrint(err)
		errno := al.Error()
		gopp.FalsePrint(errno == 0, errno)
	})
}

type Player struct {
	src   al.Source
	bufs  []al.Buffer
	frmC  chan interface{}
	stop  bool
	first int
	wg    sync.WaitGroup
}

func NewPlayer() *Player {
	this := &Player{}
	open() // need before init src,buff...
	this.init()

	return this
}

func (this *Player) init() {
	this.src = al.GenSources(1)[0]
	this.src.SetPosition(al.Vector{0, 0, 0})
	this.src.SetGain(1.0)
	log.Println(this.src.State())
	this.bufs = al.GenBuffers(bufcnt)
	log.Println(this.src, this.bufs)
	this.frmC = make(chan interface{}, 321)
}

func (this *Player) Play() {
	this.Play1()
}

func (this *Player) Play1() {
	if true {
		for _, buf := range this.bufs {
			this.fillBuffer(buf, 10)
		}
	}
	this.src.QueueBuffers(this.bufs...)

	al.PlaySources(this.src)
	errno := al.Error()
	gopp.FalsePrint(errno == 0, errno, alerrstr(errno))

	for !this.stop {
		np := this.src.BuffersProcessed()
		nb := this.src.BuffersQueued()
		_, nb = np, nb
		if np == 0 {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		log.Println("processed:", np, nb)

		for i := int32(0); i < np && !this.stop; i++ {
			uiBuffer := make([]al.Buffer, 1)
			this.src.UnqueueBuffers(uiBuffer...)
			this.fillBuffer(uiBuffer[0], 50)
			if !this.stop {
				log.Println(uiBuffer, this.src.BuffersQueued(), np) // why uiBuffer[0].Size() crash?
				this.src.QueueBuffers(uiBuffer...)
				log.Println(uiBuffer, this.src.BuffersQueued(), np)
				// al.PlaySources(this.src)
			} else {
				break
			}
		}
	}
	log.Println("Play done.")
}

func (this *Player) Stop() {
	this.stop = true
	close(this.frmC)
	al.StopSources(this.src)
	al.DeleteSources(this.src)
	al.DeleteBuffers(this.bufs...)
}

var format = uint32(al.FormatStereo16)
var freq = int32(48000)
var bufcnt = 3

func (this *Player) PutFrame(data []byte) {
	if len(this.frmC) == cap(this.frmC) {
		log.Println("chan is full:", len(this.frmC))
	} else {
		this.frmC <- data
	}
}

func (this *Player) fillBuffer(buf al.Buffer, nframe int) {
	var data []byte
	for n := 0; n < nframe; n++ {
		frmx := <-this.frmC
		if frmx == nil {
			return
		}
		data = append(data, frmx.([]byte)...)
	}
	buf.BufferData(format, data, freq)
}

func (this *Player) RecordFrame() {

}

// reference:
// https://ffainelli.github.io/openal-example/
// https://developer.tizen.org/dev-guide/2.4/org.tizen.tutorials/html/native/multimedia/openal_tutorial_n.htm
