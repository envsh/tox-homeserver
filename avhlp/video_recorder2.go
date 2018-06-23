package avhlp

/*
 */
import "C"

import (
	"errors"
	"fmt"
	"gopp"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/3d0c/gmf"
)

// TODO TODO test for android ffmpeg video record, but no lucky
func VideoRec2() *VideoRecorder2 {
	_VideoRecOnce2.Do(func() {
		_VideoRec2 = _NewVideoRecorder2()
	})
	return _VideoRec2
}

var _VideoRec2 *VideoRecorder2
var _VideoRecOnce2 sync.Once

var vrnoseq2 uint64

type VideoRecorder2Auto struct {
	vrno uint64
}

func NewVideoRecorder2Auto(f func([]byte, uint16, uint16)) *VideoRecorder2Auto {
	this := &VideoRecorder2Auto{}
	this.vrno = atomic.AddUint64(&vrnoseq, 1)
	vro := VideoRec2()
	vro.connect(this.vrno, f)

	runtime.SetFinalizer(this, func(obj interface{}) {
		vrao := obj.(*VideoRecorder2Auto)
		vro.disconnect(vrao.vrno)
	})
	return this
}

type VideoRecorder2 struct {
	devname  string
	iptctx   *gmf.FmtCtx
	onFrames map[uint64]func([]byte, uint16, uint16)
	stop     bool

	suvcap  *exec.Cmd
	rawctxp *C.char
	cretp   C.int
	cstopp  C.int
}

func _NewVideoRecorder2() *VideoRecorder2 {
	this := &VideoRecorder2{}
	// x11grab#:0.0
	this.devname = "/dev/video0"
	// this.iptctx = gmf.NewCtx()
	this.onFrames = make(map[uint64]func([]byte, uint16, uint16))
	this.stop = true

	return this
}

// can not hash a func type for map
func (this *VideoRecorder2) connect(arno uint64, f func([]byte, uint16, uint16)) {
	this.onFrames[arno] = f
	if this.stop {
		this.stop = false
		this.start()
	}
}

func (this *VideoRecorder2) disconnect(vrno uint64) {
	log.Println(vrno, this.onFrames)
	if _, ok := this.onFrames[vrno]; ok {
		delete(this.onFrames, vrno)
	}
	if len(this.onFrames) == 0 {
		log.Println("No one using video capture, stop...")
		this.stop = true
	}
}

func (this *VideoRecorder2) opencapdev() error {
	subprocstopped := false
	go func() {
		this.cstopp = 0
		sucmd := findsupath()
		suvcapexe := fmt.Sprintf("%s/libsuvcapd.so", getLibDirp())
		realcmd := fmt.Sprintf("%s %d %d %d %s", suvcapexe,
			uintptr(unsafe.Pointer(&this.rawctxp)), uintptr(unsafe.Pointer(&this.cretp)),
			uintptr(unsafe.Pointer(&this.cstopp)), this.devname)
		log.Println("Realcmd:", realcmd)
		cmdo := exec.Command(sucmd, "-c", realcmd)
		this.suvcap = cmdo
		err := cmdo.Run()
		gopp.ErrPrint(err)
		output, err := cmdo.CombinedOutput()
		log.Println("Subproc exited.", cmdo.ProcessState.String(), err, string(output))
		subprocstopped = true
	}()

	btime := time.Now()
	for {

		time.Sleep(5 * time.Millisecond)
		if subprocstopped {
			log.Println("maybe subproc has error", this.cretp)
			return errors.New("Unexcepted exit")
		}
		if this.cretp == 1 {
			log.Println(this.rawctxp)
			// this.iptctx = gmf.NewOpenedInputCtxFromPointer(unsafe.Pointer(this.rawctxp))
			break
		}
		if time.Since(btime).Seconds() > 10 {
			log.Println("Wait subproc init timeout.")
			return errors.New("Wait subproc")
		}

	}
	log.Println("Stmcnt:", this.iptctx.StreamsCnt())
	/*
		this.trySudoPerm() // for android
		this.iptctx = gmf.NewCtx()
		err := this.iptctx.OpenInput(this.devname)
		gopp.ErrPrint(err)
		if err == nil {
			log.Println("Open video capture device success:", this.devname)
		}
		log.Println("Stmcnt:", this.iptctx.StreamsCnt())
		return err
	*/
	return nil
}
func (this *VideoRecorder2) closecapdev() {
	this.cstopp = 1
	this.rawctxp = nil
	this.iptctx = nil
	/*
		this.iptctx.CloseInputAndRelease()
		this.iptctx = nil
	*/
	log.Println("Video capture device closed:", this.devname)
}

func (this *VideoRecorder2) start() { go this.runCapture() }
func (this *VideoRecorder2) runCapture() {
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
	dstcctx, swsctx, dstFrame := makeRawFrame2(iptstm)
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

func makeRawFrame2(srcVideoStream *gmf.Stream) (*gmf.CodecCtx, *gmf.SwsCtx, *gmf.Frame) {
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

func (this *VideoRecorder2) Close() {}
func (this *VideoRecorder2) trySudoPerm() {
	if gopp.IsAndroid() {
		findsu := func() string {
			paths := []string{"/system/bin/su"}
			for _, p := range paths {
				if gopp.FileExist(p) {
					return p
				}
			}
			return "su"
		}
		supath := findsu()
		_ = supath
		chmodvx := func(n int) {
			// rw+all
			cmd := exec.Command("/system/bin/su", "-c", fmt.Sprintf("/system/bin/sh -c \"chmod 666 /dev/video%d\"", n))
			output, err := cmd.CombinedOutput()
			if err != nil && strings.Contains(err.Error(), "No such file or directory") {
			} else {
				log.Println(err, strings.Replace(string(output), "\n", " ", -1))
			}
		}
		for i := 0; i < 50; i++ {
			chmodvx(i)
		}
	}
}

// https://stackoverflow.com/questions/3056113/recording-audio-with-openal
