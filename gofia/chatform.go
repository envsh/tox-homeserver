package gofia

import (
	"log"
	"time"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/animate"
	"gomatcha.io/matcha/application"
	egview "gomatcha.io/matcha/examples/view"
	"gomatcha.io/matcha/keyboard"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
)

type ChatFormState = ContactItemState

type ChatFormView struct {
	view.Embed

	cfst *ChatFormState
	// scrollPosition *view.ScrollPosition
	Text      *text.Text
	Responder *keyboard.Responder
}

func NewChatFormView() *ChatFormView {
	this := &ChatFormView{}
	// this.cfst = &ChatFormState{msgs: arraylist.New()}
	// this.scrollPosition = &view.ScrollPosition{}
	return this
}

func (v *ChatFormView) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageMounted) {
		log.Println("hehre")
	} else if view.EntersStage(from, to, view.StageVisible) {
		log.Println("hehre")
	} else if view.EntersStage(from, to, view.StageDead) {
		log.Println("hehre")
	} else if view.ExitsStage(from, to, view.StageMounted) {
		log.Println("hehre")
	} else if view.ExitsStage(from, to, view.StageVisible) {
		log.Println("hehre")
	} else if view.ExitsStage(from, to, view.StageDead) {
		log.Println("hehre")
	}
}

func (v *ChatFormView) Build(ctx view.Context) view.Model {
	if v.cfst.group {
		return v.Buildgc(ctx)
	}
	return v.Buildfc(ctx)
}

func (v *ChatFormView) Buildfc(ctx view.Context) view.Model {
	l := &constraint.Layouter{}

	hl := &constraint.Layouter{}
	hl.Solve(func(s *constraint.Solver) { s.Height(60) })
	avtbtn := view.NewImageButton()
	var imgrc *application.ImageResource
	if v.cfst.group {
		imgrc = application.MustLoadImage("ff_icongroup_2x")
	} else {
		imgrc = application.MustLoadImage("user_2x")
	}
	avtbtn.Image = imgrc
	hl.Add(avtbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, 50, 50)
	})

	namlab := view.NewTextView()
	namlab.String = "NAM图+"
	namlab.String = v.cfst.ctname
	hl.Add(namlab, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 50, 120, 30)
	})
	idlab := view.NewTextView()
	idlab.String = "TID图+"
	idlab.String = v.cfst.ctid
	hl.Add(idlab, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 130, -1, 30)
	})
	stmlab := view.NewTextView()
	stmlab.String = "STM图++++++++++++++++++++"
	stmlab.String = v.cfst.stmsg
	hl.Add(stmlab, func(s *constraint.Solver) {
		setViewGeometry4(s, 30, 50, -1, 30)
	})

	//TODO mute/mic/audio/video
	vdobtn := view.NewImageButton()
	vdobtn.Image = application.MustLoadImage("video_40")
	//vdobtn := view.NewButton()
	//vdobtn.String = "VDO图+"
	hl.Add(vdobtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right())
		setViewGeometry4(s, -1, -1, 50, -1)
		s.RightEqual(hl.Right())
	})
	adobtn := view.NewImageButton()
	adobtn.Image = application.MustLoadImage("voice_40")
	//adobtn := view.NewButton()
	//adobtn.String = "ADO图+"
	hl.Add(adobtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right().Add(-50))
		setViewGeometry4(s, 0, -1, 50, 50)
		s.RightEqual(hl.Right().Add(-50))
	})
	micbtn := view.NewImageButton()
	micbtn.Image = application.MustLoadImage("mic_30")
	micbtn.PaintStyle = &paint.Style{BackgroundColor: colornames.Gray}
	//micbtn := view.NewButton()
	// micbtn.String = "C图+"
	hl.Add(micbtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right().Add(-100))
		setViewGeometry4(s, 0, -1, 30, 30)
		s.RightEqual(hl.Right().Add(-100))
	})
	mutebtn := view.NewImageButton()
	mutebtn.Image = application.MustLoadImage("volum_40")
	mutebtn.PaintStyle = &paint.Style{BackgroundColor: colornames.Gray}
	// mutebtn := view.NewButton()
	// mutebtn.String = "M图+"
	hl.Add(mutebtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right().Add(-100))
		setViewGeometry4(s, 26, -1, 30, 30)
		s.RightEqual(hl.Right().Add(-100))
	})

	hv := view.NewBasicView()
	hv.Layouter = hl
	hv.Children = hl.Views()
	l.Add(hv, func(s *constraint.Solver) {
		s.TopEqual(l.Top())
		s.LeftEqual(l.Left())
		s.RightEqual(l.Right())
	})

	//////////////// header layout end

	////// content layout begin
	// scroll of table list
	cctablo := &table.Layouter{}
	for i := 0; i < 12; i++ {
		cell := egview.NewTableCell()
		cell.Axis = layout.AxisY
		cell.Index = 1000000000 + i + 1
		cctablo.Add(cell, nil)
	}
	// TODO 设置首次加载消息最大数，然后滚动下拉持续加载
	msgcnt := 0
	v.cfst.msgs.Each(func(index int, value interface{}) {
		msgo := value.(*ContactMessage)
		msgv := NewMessageView(msgo)
		cctablo.Add(msgv, nil)
		msgcnt += 1
	})
	ccsv := view.NewScrollView()
	ccsv.ScrollAxes = layout.AxisY // 首先设置坐标轴？？？
	// v.scrollPosition.SetValue(layout.Pt(0, 100)) // hang ui why???
	//ccsv.ScrollPosition = v.scrollPosition
	ccsv.ScrollPosition = &view.ScrollPosition{}
	ccsv.ContentLayouter = cctablo
	ccsv.ContentChildren = cctablo.Views()
	ccsv.OnScroll = func(p layout.Point) {
	}
	guide := l.Add(ccsv, func(s *constraint.Solver) {
		s.BottomEqual(l.Bottom().Add(-50))
		setViewGeometry4(s, 50, 0, -1, -1)
		s.BottomEqual(l.Bottom().Add(-50))
		s.RightEqual(l.Right())
	})
	_ = guide

	scrollBottom := 200*12 + msgcnt*30 // 或者给一个无限大的值滚动到底部？
	log.Println(ccsv.ScrollPosition == nil, scrollBottom)
	if ccsv.ScrollPosition != nil {
		log.Println(ccsv.ScrollPosition.Value()) //{0, 1770.5714111328125}
		//log.Println(v.scrollPosition)
		//log.Println(v.scrollPosition.Value())
		// ccsv.ScrollPosition.SetValue(layout.Pt(0, 100)) // hang ui why???

		//log.Println("heree", v.scrollPosition.Y.Value(), scrollBottom)
		log.Println("heree", ccsv.ScrollPosition.Y.Value(), scrollBottom)
		if ccsv.ScrollPosition.Y.Value() < float64(scrollBottom) {
			a := &animate.Basic{
				Start: ccsv.ScrollPosition.Y.Value(),
				End:   float64(scrollBottom * 2 / 3),
				Dur:   time.Second / 5,
			}
			log.Println("heree", ccsv.ScrollPosition.Y.Value(), scrollBottom)
			if true {
				ccsv.ScrollPosition.Y.Run(a)
				// 12-14 22:29:57.697 13271 13271 I Go      : gofiat go-go.go:230: call name222: Call true false 0xb3175984 3 [OnScroll <int64 Value> <[]reflect.Value Value>]
			}
			log.Println("heree")
		}
	}
	////// content layout end

	////// footer layout begin
	fl := &constraint.Layouter{}
	fl.Solve(func(s *constraint.Solver) {
		setViewGeometry4(s, -1, -1, -1, 80)
		s.LeftEqual(l.Left())
		s.BottomEqual(l.Bottom())
		s.RightEqual(l.Right())
	})
	log.Println("heree")
	log.Printf("%p, %p, %p, %v\n", v.cfst, v.cfst.Text, v.cfst.Responder, v.cfst.Responder.Visible())
	ftipt := view.NewTextInput()
	ftipt.RWText = v.cfst.Text
	ftipt.KeyboardType = keyboard.TextType
	ftipt.Responder = v.cfst.Responder
	ftipt.Placeholder = "input hereeee..."
	ftipt.MaxLines = 5
	ftipt.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	ftipt.OnChange = func(t *text.Text) {
		log.Println(t.String())
	}
	ftsvl := &constraint.Layouter{}
	ftsvl.Solve(func(s *constraint.Solver) {
		setViewGeometry4(s, -1, -1, -1, 50)
		// s.TopEqual(fl.Top())
		// s.LeftEqual(fl.Left())
		// s.BottomEqual(fl.Bottom())
		// s.RightEqual(fl.Right().Add(-80))
	})
	log.Println("heree")
	ftsvl.Add(ftipt, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, -1, 50)
		s.BottomEqual(ftsvl.Bottom())
		s.RightEqual(ftsvl.Right())
	})
	log.Println("heree")
	ftsv := view.NewScrollView()
	ftsv.ScrollAxes = layout.AxisY
	ftsv.ContentChildren = []view.View{ftipt}
	ftsv.ContentLayouter = ftsvl
	log.Println("heree")
	//ftsv.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	fl.Add(ftsv, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, -1, 50)
		s.TopEqual(fl.Top())
		s.LeftEqual(fl.Left())
		s.BottomEqual(fl.Bottom())
		s.RightEqual(fl.Right().Add(-80))
	})
	log.Println("heree")

	ftstbtn := view.NewImageButton()
	// imgrc2 := application.MustLoadImage("favicon")
	imgrc2 := application.MustLoadImage("send")
	ftstbtn.Image = imgrc2
	ftstbtn.OnPress = func() {
		log.Println("OnPress, group?", v.cfst.group)
		log.Println(ftipt.RWText.String())
		if v.cfst.group {
			appctx.vtcli.ConferenceSendMessage(v.cfst.cnum, 0, ftipt.RWText.String())
		} else {
			appctx.vtcli.FriendSendMessage(v.cfst.cnum, ftipt.RWText.String())
		}
		msgo := &ContactMessage{}
		msgo.mine = true
		msgo.msg = ftipt.RWText.String()
		msgo.mtype = 0
		msgo.tm = time.Now()
		v.cfst.msgs.Add(msgo)
		ftipt.RWText.SetString("") // clear old
		v.Signal()
	}
	log.Println("heree")
	fl.Add(ftstbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, -1, 50, 50)
		s.RightEqual(fl.Right())
	})
	log.Println("heree")
	ftemojibtn := view.NewImageButton()
	ftemojibtn.Image = application.MustLoadImage("emoji_22")
	// ftemojibtn := view.NewButton()
	// ftemojibtn.String = "Emoji图+"
	fl.Add(ftemojibtn, func(s *constraint.Solver) {
		s.RightEqual(fl.Right().Add(-50))
		setViewGeometry4(s, 0, -1, 30, 30)
	})
	log.Println("heree")
	ftfilebtn := view.NewImageButton()
	ftfilebtn.Image = application.MustLoadImage("file_22")
	// ftfilebtn := view.NewButton()
	// ftfilebtn.String = "File图+"
	fl.Add(ftfilebtn, func(s *constraint.Solver) {
		s.RightEqual(fl.Right().Add(-50))
		setViewGeometry4(s, 26, -1, 30, 30)
	})
	log.Println("heree")
	fv := view.NewBasicView()
	fv.Layouter = fl
	fv.Children = fl.Views()
	l.Add(fv, func(s *constraint.Solver) {
		s.TopEqual(l.Bottom().Add(-50))
		s.BottomEqual(l.Bottom())
		s.RightEqual(l.Right())
		s.LeftEqual(l.Left())
	})
	log.Println("heree")
	////// footer layout end
	log.Println(hl.DebugStrings())
	log.Println(fl.DebugStrings())
	log.Println(l.DebugStrings())
	// log.Println(ccsv.ContentLayouter.Layout(nil))

	vml := view.Model{}
	vml.Layouter = l
	vml.Children = l.Views()
	return vml
}

func (v *ChatFormView) Buildgc(ctx view.Context) view.Model {
	if true {
		return v.Buildfc(ctx)
	}
	vml := view.Model{}
	return vml
}
