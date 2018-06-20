package server

import (
	"gopp"
	"log"
	"tox-homeserver/thspbs"

	tox "github.com/TokTok/go-toxcore-c"
)

func _NewEvent4AV(name string, t *tox.Tox, friendNumber uint32) *thspbs.Event {
	frndnm, _ := t.FriendGetName(friendNumber)
	frndpk, _ := t.FriendGetPublicKey(friendNumber)

	evto := &thspbs.Event{Uargs: &thspbs.Argument{}, Name: name}

	evto.Uargs.FriendNumber = friendNumber
	evto.Uargs.FriendName = frndnm
	evto.Uargs.FriendPubkey = frndpk

	return evto
}

func (this *ToxVM) setupEventsForAV() {
	tav := this.tav
	t := this.t

	var audioBitRate uint32 = 48
	var videoBitRate uint32 = 64

	tav.CallbackCall(func(_ *tox.ToxAV, friendNumber uint32, audioEnabled bool, videoEnabled bool, userData interface{}) {
		log.Println(friendNumber, audioEnabled, videoEnabled)
		evto := _NewEvent4AV("Call", t, friendNumber)

		evto.Uargs.AudioEnabled = int32(gopp.IfElseInt(audioEnabled, 1, 0))
		evto.Uargs.VideoEnabled = int32(gopp.IfElseInt(videoEnabled, 1, 0))

		this.pubmsg(evto)

		// TODO
		_, err := tav.Answer(friendNumber, audioBitRate, videoBitRate)
		gopp.ErrPrint(err)
	}, nil)

	tav.CallbackCallState(func(_ *tox.ToxAV, friendNumber uint32, state uint32, userData interface{}) {
		log.Println("CallState", friendNumber, state)
		evto := _NewEvent4AV("CallState", t, friendNumber)
		evto.Uargs.CallState = state

		this.pubmsg(evto)
	}, nil)

	tav.CallbackAudioBitRate(func(_ *tox.ToxAV, friendNumber uint32, audioBitRate uint32, userData interface{}) {
		log.Println("AudioBitRate", friendNumber, audioBitRate)
		evto := _NewEvent4AV("AudioBitRate", t, friendNumber)

		evto.Uargs.AudioBitRate = audioBitRate

		this.pubmsg(evto)
	}, nil)

	tav.CallbackVideoBitRate(func(_ *tox.ToxAV, friendNumber uint32, videoBitRate uint32, userData interface{}) {
		log.Println("VideoBiteRate", friendNumber, videoBitRate)
		evto := _NewEvent4AV("VideoBiteRate", t, friendNumber)

		evto.Uargs.VideoBiteRate = videoBitRate

		this.pubmsg(evto)
	}, nil)

	tav.CallbackAudioReceiveFrame(func(_ *tox.ToxAV, friendNumber uint32, pcm []byte, sampleCount int, channels int, samplingRate int, userData interface{}) {
		// log.Println("AFrame:", friendNumber, len(pcm), sampleCount, channels, samplingRate)
		evto := _NewEvent4AV("AudioReceiveFrame", t, friendNumber)

		evto.Uargs.Pcm = pcm
		evto.Uargs.SampleCount = int32(sampleCount)
		evto.Uargs.Channels = int32(channels)
		evto.Uargs.SamplingRate = int32(samplingRate)

		this.pubmsg(evto)
	}, nil)

	tav.CallbackVideoReceiveFrame(func(_ *tox.ToxAV, friendNumber uint32, width uint16, height uint16, data []byte, userData interface{}) {
		// log.Println("VFrame:", friendNumber, width, height, len(data))
		evto := _NewEvent4AV("VideoReceiveFrame", t, friendNumber)

		evto.Uargs.Width = int32(width)
		evto.Uargs.Height = int32(height)
		evto.Uargs.VideoFrame = data

		this.pubmsg(evto)
	}, nil)
}

func (this *ToxVM) onGroupAudioFrame(t *tox.Tox, groupNumber uint32, peerNumber uint32, pcm []byte, samples uint, channels uint8, sample_rate uint32, userData interface{}) {
	groupTitle, _ := t.ConferenceGetTitle(groupNumber)
	grouppk, _ := t.ConferenceGetIdentifier(groupNumber)
	peerpk, _ := t.ConferencePeerGetPublicKey(groupNumber, peerNumber)
	peerName, _ := t.ConferencePeerGetName(groupNumber, peerNumber)

	evto := &thspbs.Event{Uargs: &thspbs.Argument{}, Name: "ConferenceAudioRecieiveFrame"}

	evto.Uargs.GroupNumber = groupNumber
	evto.Uargs.GroupTitle = groupTitle
	evto.Uargs.GroupIdentity = grouppk
	evto.Uargs.PeerNumber = peerNumber
	evto.Uargs.PeerName = peerName
	evto.Uargs.PeerPubkey = peerpk

	evto.Uargs.Pcm = pcm
	evto.Uargs.SampleCount = int32(samples)
	evto.Uargs.Channels = int32(channels)
	evto.Uargs.SamplingRate = int32(sample_rate)

	this.pubmsg(evto)
}
