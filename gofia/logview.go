package gofia

import (
	"gopp"
	"image/color"
	"log"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/application"
	"gomatcha.io/matcha/keyboard"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
)

type LogState struct {
	text *text.Text
	tnum int
}

func newLogState() *LogState {
	this := &LogState{}
	this.text = text.New("")
	return this
}

type LogView struct {
	view.Embed

	TextColor color.Color
	text      *text.Text
	responder *keyboard.Responder

	logst *LogState
}

func newLogView() *LogView {
	this := &LogView{}
	this.TextColor = colornames.Red
	this.text = text.New("")
	this.responder = &keyboard.Responder{}

	this.logst = appctx.logState
	return this
}

func (v *LogView) Lifecycle(from, to view.Stage) {
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

func (v *LogView) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}

	log.Println("Herehere")
	hl := &constraint.Layouter{}
	hl.Solve(func(s *constraint.Solver) {
		s.Height(50)
	})
	log.Println("Herehere")
	setbtn := view.NewImageButton()
	log.Println("Herehere")
	// setbtn.Image = application.MustLoadImage("settings")
	// setbtn.Image = application.MustLoadImage("app_system")
	setbtn.Image = application.MustLoadImage("barbuttonicon_back_2x")
	log.Println("Herehere")
	// setbtn := view.NewButton()
	// setbtn.String = "SET?的"
	setbtn.OnPress = func() {}
	hl.Add(setbtn, func(s *constraint.Solver) {
		s.Top(0)
		s.Left(0)
		s.Height(50)
		s.Width(50)
	})

	log.Println("Herehere")
	netbtn := view.NewTextView()
	netbtn.String = "NONE"
	if appctx.mvst.netStatus > 0 {
		netbtn.String = gopp.IfElseStr(appctx.mvst.netStatus == 1, "UDP", "TCP")
	}
	hl.Add(netbtn, func(s *constraint.Solver) {
		// setViewGeometry4(s, 0, 50, 50, 50)
		s.Left(50)
	})

	log.Println("Herehere")
	stsbtn := view.NewImageButton()
	if appctx.mvst.netStatus > 0 {
		stsbtn.Image = application.MustLoadImage("online_30")
	} else {
		stsbtn.Image = application.MustLoadImage("offline_30")
	}
	// stbtn := view.NewButton()
	// stbtn.String = "STS中"
	hl.Add(stsbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 80, 50, 50)
		// s.Top(0)
		// s.Left(50)
		// s.Height(50)
	})

	log.Println("Herehere")
	nlab := view.NewTextView()
	nlab.String = "名字Tofia"
	nlab.String = appctx.mvst.nickName
	nlab.Style.SetFont(text.DefaultFont(18))
	// nlab.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	log.Println("align:", nlab.Style.Alignment())
	hl.Add(nlab, func(s *constraint.Solver) {
		// s.Top(0)
		s.Left(130)
		// s.Height(50)
	})

	log.Println("Herehere")
	addbtn := view.NewImageButton()
	// addbtn.Image = application.MustLoadImage("add")
	// addbtn.Image = application.MustLoadImage("list_add")
	addbtn.Image = application.MustLoadImage("contacts_add_friend_2x")
	// addbtn := view.NewButton()
	// addbtn.String = "F加++"
	addbtn.OnPress = func() {}
	hl.Add(addbtn, func(s *constraint.Solver) {
		s.Top(0)
		s.RightEqual(hl.Right())
		s.Height(50)
		s.Width(50)
	})

	log.Println("Herehere")
	hv := view.NewBasicView()
	hv.Layouter = hl
	hv.Children = hl.Views()
	l.Add(hv, func(s *constraint.Solver) {
		s.Top(0)
		s.Left(0)
		s.Height(50)
		s.RightEqual(l.Right())
	})

	log.Println("Herehere")
	// Create a new textview.
	child := view.NewTextView()
	child.String = "Hello World"
	child.Style.SetTextColor(v.TextColor)
	child.Style.SetFont(text.DefaultBoldFont(50))
	child.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}

	// Layout is primarily done using constraints. More info can be
	// found in the matcha/layout/constraints docs.
	l.Add(child, func(s *constraint.Solver) {
		s.Top(50)
		s.Left(0)
	})

	log.Println("Herehere")
	border := view.NewBasicView()
	border.Painter = &paint.Style{BackgroundColor: colornames.Gray}
	l.Add(border, func(s *constraint.Solver) {
		s.Height(1)
		s.Top(115)
		s.LeftEqual(l.Left())
		s.RightEqual(l.Right())
		// s.BottomEqual(l.Bottom())
	})

	log.Println("Herehere")
	logipt := view.NewTextInput()
	logipt.RWText = v.text
	logipt.Placeholder = "input log here:"
	logipt.KeyboardType = keyboard.TextType
	logipt.MaxLines = 5
	logipt.Responder = v.responder
	//logipt.OnSubmit = func(t *text.Text) {
	// v.responder.Dismiss()
	// t.SetString("")
	//}
	logipt.OnChange = func(t *text.Text) {
		// v.Signal()
		logipt.Signal()
	}

	/*
		go func() {
			time.Sleep(3 * time.Second)
			matcha.MainLocker.Lock()
			v.text.SetString("aaaa在aaaa" + "\n")
			matcha.MainLocker.Unlock()
			v.Signal()
		}()
	*/

	// logipt.RWText.SetString("hehe呵呵")
	l.Add(logipt, func(s *constraint.Solver) {
		s.Top(115)
		s.Left(0)
		s.WidthEqual(l.Width())
		// s.Height(201)
		s.BottomEqual(l.Bottom())
	})

	return view.Model{
		Layouter: l,
		Children: l.Views(),
	}
}
