package avhlp

/*
#include <AL/al.h>
#include <AL/alc.h>
*/
import "C"

import (
	// "github.com/3d0c/gmf"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"golang.org/x/mobile/exp/audio/al"
)

func AudioRec() *AudioRecorder {
	_AudioRecOnce.Do(func() {
		_AudioRec = _NewAudioRecorder()
	})
	return _AudioRec
}

var _AudioRec *AudioRecorder
var _AudioRecOnce sync.Once

var arnoseq uint64

type AudioRecorderAuto struct {
	arno uint64
}

func NewAudioRecorderAuto(f func([]byte, uint32, uint8, uint32)) *AudioRecorderAuto {
	this := &AudioRecorderAuto{}
	this.arno = atomic.AddUint64(&arnoseq, 1)
	aro := AudioRec()
	aro.connect(this.arno, f)

	runtime.SetFinalizer(this, func(obj interface{}) {
		arao := obj.(*AudioRecorderAuto)
		aro.disconnect(arao.arno)
	})
	return this
}

type AudioRecorder struct {
	capdev   *C.ALCdevice
	onFrames map[uint64]func([]byte, uint32, uint8, uint32)
	stop     bool
}

func _NewAudioRecorder() *AudioRecorder {
	this := &AudioRecorder{}
	this.onFrames = make(map[uint64]func([]byte, uint32, uint8, uint32))
	this.stop = true
	al.Error()
	this.opencapdev()
	return this
}

// can not hash a func type for map
func (this *AudioRecorder) connect(arno uint64, f func([]byte, uint32, uint8, uint32)) {
	this.onFrames[arno] = f
	if this.stop {
		this.stop = false
		this.start()
	}
}

func (this *AudioRecorder) disconnect(arno uint64) {
	log.Println(arno, this.onFrames)
	if _, ok := this.onFrames[arno]; ok {
		delete(this.onFrames, arno)
	}
	if len(this.onFrames) == 0 {
		log.Println("No one using audio capture, stop...")
		this.stop = true
	}
}

func (this *AudioRecorder) start() { go this.runCapture() }
func (this *AudioRecorder) runCapture() {
	btime := time.Now()
	log.Println("Audio capture resumed. stoped:", this.stop)
	for !this.stop {
		var samples C.ALint
		C.alcGetIntegerv(this.capdev, C.ALC_CAPTURE_SAMPLES, C.sizeof_ALint, &samples)
		if samples < AUDIO_FRAME_SAMPLE_COUNT {
			time.Sleep(5 * time.Millisecond)
			continue
		}
		buffer := make([]byte, AUDIO_CAPTURE_BUFSIZE*2) // really should be uint16, but now byte, so *2
		// dangerous, if buf is small, the stack will mass and crash randomly
		C.alcCaptureSamples(this.capdev, unsafe.Pointer(&buffer[0]), AUDIO_FRAME_SAMPLE_COUNT)
		for _, f := range this.onFrames {
			(f)(buffer, AUDIO_FRAME_SAMPLE_COUNT, AUDIO_CHANNELS, AUDIO_SAMPLE_RATE)
		}
	}
	log.Println("Audio capture paused. stoped:", this.stop, AUDIO_CAPTURE_BUFSIZE, time.Since(btime))
}

const AUDIO_SAMPLE_RATE = 48000
const AUDIO_FRAME_DURATION = 20
const AUDIO_CHANNELS = 2
const AUDIO_DEVICE_BUFSIZE = (AUDIO_FRAME_DURATION * AUDIO_SAMPLE_RATE * 4) / 1000 * AUDIO_CHANNELS
const AUDIO_FRAME_SAMPLE_COUNT = AUDIO_FRAME_DURATION * AUDIO_SAMPLE_RATE / 1000
const AUDIO_CAPTURE_BUFSIZE = AUDIO_FRAME_SAMPLE_COUNT * AUDIO_CHANNELS

func (this *AudioRecorder) opencapdev() {
	this.capdev = C.alcCaptureOpenDevice(nil, AUDIO_SAMPLE_RATE, al.FormatStereo16, AUDIO_CAPTURE_BUFSIZE)
	errno := al.Error()
	if errno != 0 {
		log.Println(errno, alerrstr(errno))
	} else {
		log.Println("Open audio capture device success.",
			AUDIO_DEVICE_BUFSIZE, AUDIO_CAPTURE_BUFSIZE, AUDIO_FRAME_SAMPLE_COUNT)
	}
	C.alcCaptureStart(this.capdev)

}

func (this *AudioRecorder) closecapdev() {
	C.alcCaptureStop(this.capdev)
	C.alcCaptureCloseDevice(this.capdev)
}

func (this *AudioRecorder) Close() {}

// https://stackoverflow.com/questions/3056113/recording-audio-with-openal
