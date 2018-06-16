package avhlp

import (
	"C"
	"gopp"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/3d0c/gmf"
)

func VideoRec() *VideoRecorder {
	_VideoRecOnce.Do(func() {
		_VideoRec = _NewVideoRecorder()
	})
	return _VideoRec
}

var _VideoRec *VideoRecorder
var _VideoRecOnce sync.Once

var vrnoseq uint64

type VideoRecorderAuto struct {
	vrno uint64
}

func NewVideoRecorderAuto(f func([]byte, uint16, uint16)) *VideoRecorderAuto {
	this := &VideoRecorderAuto{}
	this.vrno = atomic.AddUint64(&vrnoseq, 1)
	vro := VideoRec()
	vro.connect(this.vrno, f)

	runtime.SetFinalizer(this, func(obj interface{}) {
		vrao := obj.(*VideoRecorderAuto)
		vro.disconnect(vrao.vrno)
	})
	return this
}

type VideoRecorder struct {
	devname  string
	iptctx   *gmf.FmtCtx
	onFrames map[uint64]func([]byte, uint16, uint16)
	stop     bool
}

func _NewVideoRecorder() *VideoRecorder {
	this := &VideoRecorder{}
	// x11grab#:0.0
	this.devname = "/dev/video0"
	// this.iptctx = gmf.NewCtx()
	this.onFrames = make(map[uint64]func([]byte, uint16, uint16))
	this.stop = true

	return this
}

// can not hash a func type for map
func (this *VideoRecorder) connect(arno uint64, f func([]byte, uint16, uint16)) {
	this.onFrames[arno] = f
	if this.stop {
		this.stop = false
		this.start()
	}
}

func (this *VideoRecorder) disconnect(vrno uint64) {
	log.Println(vrno, this.onFrames)
	if _, ok := this.onFrames[vrno]; ok {
		delete(this.onFrames, vrno)
	}
	if len(this.onFrames) == 0 {
		log.Println("No one using video capture, stop...")
		this.stop = true
	}
}

func (this *VideoRecorder) opencapdev() error {
	this.iptctx = gmf.NewCtx()
	err := this.iptctx.OpenInput(this.devname)
	gopp.ErrPrint(err)
	if err == nil {
		log.Println("Open video capture device success:", this.devname)
	}
	log.Println("Stmcnt:", this.iptctx.StreamsCnt())
	return err
}
func (this *VideoRecorder) closecapdev() {
	this.iptctx.CloseInputAndRelease()
	this.iptctx = nil
	log.Println("Video capture device closed:", this.devname)
}

func (this *VideoRecorder) start() { go this.runCapture() }
func (this *VideoRecorder) runCapture() {
	btime := time.Now()
	if err := this.opencapdev(); err != nil {
		return
	}
	defer this.closecapdev()

	// stm, err := iptctx.GetStream(0)
	iptstm, err := this.iptctx.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	gopp.ErrPrint(err)
	log.Printf("iptstm: id: %d, idx: %d, A/V: %v/%v, nbfrms: %d, stmty: %d, ccset: %v\n",
		iptstm.Id(), iptstm.Index(), iptstm.IsAudio(), iptstm.IsVideo(), iptstm.NbFrames(),
		iptstm.Type(), iptstm.IsCodecCtxSet())

	iptcctx := iptstm.CodecCtx()
	dstcctx, swsctx, dstFrame := makeRawFrame(iptstm)
	// dstc, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	// gopp.ErrPrint(err, gmf.AV_CODEC_ID_RAWVIDEO)
	// dstcctx := gmf.NewCodecCtx(dstc)

	srcpkti := 0
	pktch := this.iptctx.GetNewPackets()
	for !this.stop {
		srcpkt := <-pktch
		srcpkti++
		frmi := 0
		ok := true
		frm := <-srcpkt.Frames(iptcctx)
		{
			if this.stop {
				frm.Free()
				frm.Release()
				break
			}
			swsctx.Scale(frm, dstFrame)
			pkt, err := dstFrame.Encode(dstcctx)
			gopp.ErrPrint(err, ok, frm)

			log.Printf("srcpkti: %d, frmi: %d, ok: %v, pktstmidx: %d, iptstmidx: %d\n",
				srcpkti, frmi, ok, pkt.StreamIndex(), iptstm.Index())
			if pkt.StreamIndex() != iptstm.Index() {
				log.Println("steam index not match,", pkt.StreamIndex(), iptstm.Index())
				pkt.Free()
				pkt.Release()
				continue
			}
			log.Printf("Vcap: channels: %d, fmt:%d, keyfrm: %d, w/h: %d/%d, ts: %d, size: %d/%d, LineSize: %d/%d\n",
				frm.Channels(), frm.Format(), frm.KeyFrame(), frm.Width(), frm.Height(),
				frm.TimeStamp(), pkt.Size(), len(pkt.Data()), frm.LineSize(0), frm.LineSize(1))

			dat := pkt.Data()
			for _, f := range this.onFrames {
				_, _ = f, dat
				// (f)(dat, uint16(frm.Width()), uint16(frm.Height()))
				(f)(dat, uint16(dstFrame.Width()), uint16(dstFrame.Height()))
			}

			frm.Free()
			frm.Release()
			pkt.Free()
			pkt.Release()
		}

		srcpkt.Free()
		srcpkt.Release()
	}

	log.Println("Video capture stoped:", this.stop, AUDIO_CAPTURE_BUFSIZE, time.Since(btime))
}

func makeRawFrame(srcVideoStream *gmf.Stream) (*gmf.CodecCtx, *gmf.SwsCtx, *gmf.Frame) {
	codec, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	if err != nil {
		log.Fatal(err)
	}

	cc := gmf.NewCodecCtx(codec)
	// defer gmf.Release(cc)

	cc.SetTimeBase(gmf.AVR{Num: 1, Den: 1})

	// This should really be AV_PIX_FMT_RGB32, but then red and blue channels are switched
	cc.SetPixFmt(gmf.AV_PIX_FMT_RGB24).SetWidth(srcVideoStream.CodecCtx().Width()).SetHeight(srcVideoStream.CodecCtx().Height())
	if codec.IsExperimental() {
		cc.SetStrictCompliance(gmf.FF_COMPLIANCE_EXPERIMENTAL)
	}

	if err := cc.Open(nil); err != nil {
		log.Fatal(err)
	}

	swsCtx := gmf.NewSwsCtx(srcVideoStream.CodecCtx(), cc, gmf.SWS_BICUBIC)
	// defer gmf.Release(swsCtx)

	dstFrame := gmf.NewFrame().
		SetWidth(srcVideoStream.CodecCtx().Width()).
		SetHeight(srcVideoStream.CodecCtx().Height()).
		SetFormat(gmf.AV_PIX_FMT_RGB24) // see above

	if err := dstFrame.ImgAlloc(); err != nil {
		log.Fatal(err)
	}
	return cc, swsCtx, dstFrame
}

func (this *VideoRecorder) Close() {}

// https://stackoverflow.com/questions/3056113/recording-audio-with-openal
