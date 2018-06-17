package main

import (
	"fmt"
	"gopp"
	"log"
	"sync"
	"time"
	"tox-homeserver/avhlp"
)

type AVSession struct {
	audioPlayer   *avhlp.Player
	videoPlayer   *VideoPlayer
	audioEnabled  bool
	videoEnabled  bool
	audioRecorder *avhlp.AudioRecorderAuto
	videoRecorder *avhlp.VideoRecorderAuto
	muteVideo     bool // for self
	muteMic       bool // for self
	muteMixer     bool // for self
	btime         time.Time
	contact       string // group id or friend pubkey
	initiatorIsMe bool   // is me initate this av session

	// from local recorder
	onNewAudioFrame func(aframe []byte, samples uint32, channels uint8, samplingRate uint32)
	onNewVideoFrame func(vframe []byte, width, height uint16)
}

type AVManager struct {
	sesses map[string]*AVSession // contact =>
	sessmu sync.RWMutex
}

var avman *AVManager
var avmanOnce sync.Once

func AVMan() *AVManager {
	avmanOnce.Do(func() {
		avman = NewAVManager()
	})
	return avman
}
func NewAVManager() *AVManager {
	this := &AVManager{}
	this.sesses = make(map[string]*AVSession)
	return this
}

func (this *AVManager) NewSession(contact string, audioEnabled, videoEnabled bool,
	onNewAudioFrame func([]byte, uint32, uint8, uint32),
	onNewVideoFrame func([]byte, uint16, uint16)) error {
	if this.HasSession(contact) {
		return fmt.Errorf("Session already exists: %s", contact)
	}
	this.sessmu.Lock()
	defer this.sessmu.Unlock()

	sess := &AVSession{}
	sess.contact = contact
	sess.audioEnabled = audioEnabled
	sess.videoEnabled = videoEnabled
	sess.onNewAudioFrame = onNewAudioFrame
	sess.onNewVideoFrame = onNewVideoFrame
	sess.btime = time.Now()
	if !gopp.IsAndroid() { // see upstream bug: #11
		sess.audioPlayer = avhlp.NewPlayer()
	}

	if videoEnabled {
		sess.videoPlayer = NewVideoPlayer()
	}

	if sess.onNewAudioFrame != nil && sess.audioEnabled {
		if !gopp.IsAndroid() {
			sess.audioRecorder = avhlp.NewAudioRecorderAuto(sess.onNewAudioFrame)
		}
	}
	if sess.onNewAudioFrame != nil && sess.videoEnabled {
		if !gopp.IsAndroid() {
			sess.videoRecorder = avhlp.NewVideoRecorderAuto(sess.onNewVideoFrame)
		}
	}

	this.sesses[contact] = sess
	if sess.audioPlayer != nil {
		sess.audioPlayer.Play()
	}

	return nil
}

func (this *AVManager) HasSession(contact string) bool {
	this.sessmu.RLock()
	defer this.sessmu.RUnlock()
	_, ok := this.sesses[contact]
	return ok
}

// stop and remove
func (this *AVManager) RemoveSession(contact string, name string) error {
	if !this.HasSession(contact) {
		return fmt.Errorf("Session not found, %s", contact)
	}

	this.sessmu.Lock()
	defer this.sessmu.Unlock()
	sess := this.sesses[contact]

	if !sess.muteVideo {
		log.Println("Stop video recorder...", name)
		if sess.videoRecorder != nil {
			sess.audioRecorder = nil
		}
	}
	if !sess.muteMic {
		log.Println("Stop audio recorder...", name, sess.onNewAudioFrame)
		if sess.audioRecorder != nil {
			sess.audioRecorder = nil
		}
	}
	if sess.videoEnabled {
		log.Println("Stop video player...", name)
		sess.videoPlayer.Stop()
		sess.videoPlayer = nil
	}
	if !sess.muteMixer {
		log.Println("Stop audio player...", name)
		if sess.audioPlayer != nil {
			sess.audioPlayer.Stop()
			sess.audioPlayer = nil
		}
	}

	delete(this.sesses, contact)
	log.Printf("AVSession info(%s): eclapsed: %s, A/V: %v/%v, mic/mixer: %v/%v\n", name,
		time.Since(sess.btime), sess.audioEnabled, sess.videoEnabled, !sess.muteMic, !sess.muteMixer)

	return nil
}

func (this *AVManager) SetMuteMic(contact string, mute bool) error {
	return nil
}

func (this *AVManager) SetMuteMixer(contact string, mute bool) error {
	return nil
}

func (this *AVManager) PutAudioFrame(contact string, frame []byte) error {
	this.sessmu.RLock()
	sess := this.sesses[contact]
	this.sessmu.RUnlock()
	if sess != nil && sess.audioPlayer != nil {
		sess.audioPlayer.PutFrame(frame)
	}
	return nil
}

func (this *AVManager) PutVideoFrame(contact string, frame []byte, width, heigh int) error {
	this.sessmu.RLock()
	sess := this.sesses[contact]
	this.sessmu.RUnlock()
	if sess != nil {
		if sess.videoEnabled {
			sess.videoPlayer.PutFrame(frame, width, heigh)
		}
	}
	return nil
}
