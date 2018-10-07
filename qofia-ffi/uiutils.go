package main

import (
	"bytes"
	"gopp"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"unsafe"

	store "tox-homeserver/store"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"

	"github.com/holys/initials-avatar"
	"github.com/issue9/identicon"
)

func Repolish(w qtwidgets.QWidget_ITF) {
	syl := w.QWidget_PTR().Style()
	syl.Unpolish(w)
	syl.Polish(w)
}

// let QScrollArea with QScroller installed can only vertical scroll
func SetScrollContentTrackerSize(sa *qtwidgets.QScrollArea) {
	wgt := sa.Widget()
	sa.InheritResizeEvent(func(arg0 *qtgui.QResizeEvent) {
		osz := arg0.OldSize()
		nsz := arg0.Size()
		if false {
			log.Println(osz.Width(), osz.Height(), nsz.Width(), nsz.Height())
		}
		if osz.Width() != nsz.Width() {
			wgt.SetMaximumWidth(nsz.Width())
		}
		// this.ScrollArea_2.ResizeEvent(arg0)
		arg0.Ignore() // I ignore, you handle it. replace explict call parent's
	})
}

func SetQLabelElideText(lab *qtwidgets.QLabel, txt string, suff string, skipTooltip ...bool) {
	// font := lab.PaintEngine().Painter().Font()
	font := lab.Font()
	rect := lab.Rect()
	sz1 := lab.Size()
	sz2 := lab.SizeHint()
	gm := lab.Geometry()

	elwidth := int(gopp.MaxU32([]uint32{uint32(rect.Width()), uint32(sz2.Width())}))
	elwidth = gopp.IfElseInt(elwidth > 500, rect.Width(), elwidth)
	elwidth = rect.Width()
	elwidth = gopp.IfElseInt(elwidth < 150, sz2.Width(), elwidth)
	elwidth = gopp.IfElseInt(elwidth > 150, elwidth-10, elwidth)
	if false {
		log.Println(rect.Width(), sz1.Width(), sz2.Width(), gm.Width(), lab.ObjectName(), txt)
	}

	fm := qtgui.NewQFontMetrics(font)
	elwidth -= gopp.IfElseInt(elwidth < 150, 0, fm.Width__(suff))
	etxt := fm.ElidedText__(txt, qtcore.Qt__ElideRight, elwidth)
	if false {
		log.Println(len(txt), len(etxt), elwidth, fm.Width__(suff), len(suff), suff)
	}

	lab.SetText(etxt + suff)
	if len(skipTooltip) == 0 {
		lab.SetToolTip(txt + suff)
	}
}

func SetQWidgetDropable(w qtwidgets.QWidget_ITF, dropable bool) {
	w.QWidget_PTR().InheritDragEnterEvent(func(event *qtgui.QDragEnterEvent) {
		if dropable {
			event.AcceptProposedAction()
		}
	})
}

// case paintEvent crash
func PointerStep(p unsafe.Pointer, offset uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + offset)
}

// the offset calc can used for both Qt4 and Qt5
func NewQPainter(w qtwidgets.QWidget_ITF) *qtgui.QPainter {
	ptr := PointerStep(w.QWidget_PTR().GetCthis(), 2*unsafe.Sizeof(uintptr(0)))
	ptdev := qtgui.NewQPaintDeviceFromPointer(ptr)
	return qtgui.NewQPainter_1(ptdev)
}

// TODO need margin of sides
// TODO how get round background mask?
// for friend
func GetIdentIcon(s string) *qtgui.QIcon {
	back := color.RGBA{240, 240, 240, 255}

	// from github
	fore1 := color.RGBA{228, 131, 172, 255} // pink, 255 no transpart
	fore2 := color.RGBA{144, 181, 233, 255} // blue

	fore := gopp.IfElse(rand.Int()%2 == 1, fore1, fore2).(color.RGBA)
	imgo, err := identicon.Make(64, back, fore, []byte(s))

	// fores := []color.Color{color.Black, color.RGBA{200, 2, 5, 100}, color.RGBA{2, 200, 5, 100}}
	// ii, err := identicon.New(64, back, fores...)
	// imgo := ii.Make([]byte(s))

	gopp.ErrPrint(err, s)

	w2 := store.GetFSC().TempFile()
	err = png.Encode(w2, imgo)
	gopp.ErrPrint(err)
	gopp.ErrPrint(w2.Sync())
	defer os.Remove(w2.Name())
	defer w2.Close()

	idico := qtgui.NewQIcon_2(w2.Name())
	gopp.FalsePrint(!idico.IsNull(), "gen idico failed.")
	// log.Println(w2.Name(), idico.IsNull(), s)

	if false { // null result?
		w := bytes.NewBuffer([]byte{})
		err = png.Encode(w, imgo)
		data := w.Bytes()
		qimgo := qtgui.QImage_FromData(unsafe.Pointer(&data[0]), len(data), "PNG")
		if qimgo.IsNull() {
		}
		idico := qtgui.NewQIcon_1(qtgui.QPixmap_FromImage(qimgo, 0))
		log.Println(len(data), qimgo.IsNull(), idico.IsNull())
	}

	return idico
}

// for group
func GetInitAvatar(name string) *qtgui.QIcon {
	// fontFile := "./resource/Hiragino_Sans_GB_W3.ttf"

	cfg := avatar.Config{}
	cfg.FontFile = fontFile
	cfg.MaxItems = 2
	avto := avatar.NewWithConfig(cfg)
	name = strings.TrimLeft(name, " ~!#$%^&*()_+`-={}[]\\|:\";'<>?,./")
	data, err := avto.DrawToBytes(name, 64)
	gopp.ErrPrint(err, len(name), name)

	fname := store.GetFSC().TempFileName()
	err = ioutil.WriteFile(fname, data, 0644)
	gopp.ErrPrint(err, len(fname), fname, len(data))
	defer os.Remove(fname)

	idico := qtgui.NewQIcon_2(fname)
	gopp.FalsePrint(!idico.IsNull(), "gen idico failed.", len(name), name)
	return idico
}

// fontforge: font name: FZLTHJW--GB1-0,  family: FZLanTingHeiS-R-GB
var fontFile = "./resource/fzlt.ttf"

func PrepareFont() {
	rcfile := ":/resource/fzlt.ttf"

	locFile := store.GetFSC().GetFilePath("fzlt.ttf")
	fi, err := os.Stat(locFile)
	if err != nil || fi.Size() == 0 /*|| crc1 != crc2 */ {
		fp := qtcore.NewQFile_1(rcfile)
		fp.Open(qtcore.QIODevice__ReadOnly)
		data := qtcore.NewQIODeviceFromPointer(fp.GetCthis()).ReadAll().Data_fix()
		qtcore.NewQIODeviceFromPointer(fp.GetCthis()).Close()

		if len(data) > 0 {
			fontFile = locFile
			err := ioutil.WriteFile(fontFile, []byte(data), 0644)
			gopp.ErrPrint(err, fontFile)
		} else {
			log.Println("maybe read font rc file error.", len(data), rcfile)
		}
	} else {
		fontFile = locFile
	}
	log.Println(fontFile)
}

func FindProperFontFile(w *qtwidgets.QWidget) {
	// qtgui.QFontDatabase__SimplifiedChinese
	fi := w.FontInfo()
	fnto := w.Font()
	log.Println(fi, fnto)
	log.Println(fnto.Family(), fnto.ToString(), fnto.RawName())
	log.Println(fnto.StyleName())
	log.Println(fnto.Key(), fnto.LastResortFamily())
	if false { //failed
		fnto2 := qtgui.NewQFont_1_(fnto.Family())
		log.Println(fnto2.RawName())
	}
	fflst := qtgui.QFontDatabase_ApplicationFontFamilies(qtgui.QFontDatabase__Any)
	fflstx := qtcore.NewQStringListxFromPointer(fflst.GetCthis())
	log.Println(fflstx.Count_1())
	log.Println(fflstx.ConvertToSlice())
}

func setAutoHeightForTextEdit(te *qtwidgets.QTextEdit) {
	font := te.Font()
	fm := qtgui.NewQFontMetrics(font)

	minh := 24
	avgh := (fm.Height() + minh) / 2
	maxh := avgh * 4

	qtrt.Connect(te.Document().DocumentLayout(),
		"documentSizeChanged(const QSizeF &)", func(sz *qtcore.QSizeF) {
			newh := int(sz.Height())
			if newh >= minh && newh <= maxh {
				// te.UpdateGeometry()
				te.SetFixedHeight(int(sz.Height()) + 2)
			}
		})
}
func unsetAutoHeightForTextEdit(te *qtwidgets.QTextEdit) {
	qtrt.Disconnect(te.Document().DocumentLayout(), "documentSizeChanged(const QSizeF &)")
}
