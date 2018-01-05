package gofia

import (
	"gopp"
	"log"
	"math"
	"time"

	"github.com/kitech/godsts/lists/arraylist"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/application"
	"gomatcha.io/matcha/keyboard"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pointer"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
)

type ContactMessage struct {
	msg   string
	mine  bool //
	mtype int  //
	tm    time.Time
}

type ContactItemState struct {
	group  bool
	cnum   uint32
	ctid   string
	ctname string
	status uint32
	stmsg  string
	avatar string

	// *ContactMessage
	msgs           *arraylist.List
	Text           *text.Text
	Responder      *keyboard.Responder
	ScrollPosition *view.ScrollPosition
}

func newContactItemState() *ContactItemState {
	this := &ContactItemState{}
	this.msgs = arraylist.New()
	this.Text = text.New("")
	this.Responder = &keyboard.Responder{}
	this.ScrollPosition = &view.ScrollPosition{}
	return this
}

type ContactItem struct {
	view.Embed

	OnTouch func()

	ctis *ContactItemState
}

func NewContactItem(group bool) *ContactItem {
	this := &ContactItem{}
	// this.ContactItemState = &ContactItemState{msgs: arraylist.New()}
	// this.group = group

	return this
}

func (this *ContactItem) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}
	l.Solve(func(s *constraint.Solver) { s.Height(60) })

	evtbtn := view.NewButton()
	evtbtn.String = "EVT图"
	l.Add(evtbtn, func(s *constraint.Solver) {
		// setViewGeometry4(s, 0, 0, 40, 40)
		s.CenterYEqual(l.CenterY())
		s.Left(0)
		s.Width(40)
		s.Height(40)
	})

	avtbtn := view.NewImageButton()
	//avtbtn.Image = application.MustLoadImage("ic_launcher")
	// avtbtn.Image = application.MustLoadImage("identity")
	if this.ctis.group {
		// avtbtn.Image = application.MustLoadImage("ic_launcher")
		avtbtn.Image = application.MustLoadImage("ff_icongroup_2x")
	} else {
		//avtbtn.Image = application.MustLoadImage("contact")
		avtbtn.Image = application.MustLoadImage("user_2x")
	}
	avtbtn.OnPress = func() {
		log.Println("clicked:", this.ctis)
		log.Println("view path:", ctx.Path())

		// always new view
		cf := NewChatFormView()
		cf.cfst = this.ctis
		// appctx.cfvs.Put(this.ctid, cf)
		// if !appctx.cfvs.Has(this.ctid) {
		// }
		// cfx, found := appctx.cfvs.Get(this.ctid)
		// if !found {
		//	log.Println("not found:", this.ctid)
		//} else {
		//	appctx.currV = cfx.(*ChatFormView)
		//}
		// appctx.currV = cf
		// appctx.mainV.(*TutorialView).Signal()
		appctx.app.Child = cf
		appctx.app.ChildRelay.Signal()
	}
	l.Add(avtbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 40, 60, 60)
	})

	stsbtn := view.NewImageButton()
	if this.ctis.status > 0 {
		stsbtn.Image = application.MustLoadImage("online_30")
	} else {
		stsbtn.Image = application.MustLoadImage("offline_30")
	}
	// stsbtn := view.NewButton()
	// stsbtn.String = "STS图"
	stsbtn.OnPress = func() {
		key := this.ctis.ctid
		if cfsx, found := appctx.chatFormStates.Get(key); found {
			cfst := cfsx.(*ChatFormState)
			cfv := NewChatFormView()
			cfv.cfst = cfst
			// appctx.currV = cfv
			// appctx.mainV.(*TutorialView).Signal()
			appctx.app.Child = cfv
			appctx.app.ChildRelay.Signal()
		} else {
			log.Println("chat form not found:", key)
		}
		/*
			if cfx, found := appctx.cfvs.Get(key); found {
				appctx.currV = cfx.(*ChatFormView)
				appctx.mainV.(*TutorialView).Signal()
			} else {
				log.Println("chat form not found:", key)
			}
		*/
	}
	l.Add(stsbtn, func(s *constraint.Solver) {
		// setViewGeometry4(s, 0, 100, 40, 40)
		s.CenterYEqual(l.CenterY())
		s.Left(100)
		s.Width(40)
		s.Height(40)
	})

	titlab := view.NewTextView()
	titlab.String = "TITLE TEXT字"
	titlab.String = this.ctis.ctname + "." + gopp.SubStr(this.ctis.ctid, 5)
	l.Add(titlab, func(s *constraint.Solver) {
		// setViewGeometry4(s, 0, 140, 60, 20)
		s.Top(0)
		s.Left(140)
		s.Height(20)
	})

	stslab := view.NewTextView()
	stslab.String = "STS TEXT字"
	stslab.String = this.ctis.stmsg
	stslab.Style.SetWrap(text.WrapWord)
	l.Add(stslab, func(s *constraint.Solver) {
		s.Top(20)
		s.Left(140)
		s.Height(30)
	})

	minilab := view.NewTextView()
	minilab.String = "mini never"
	l.Add(minilab, func(s *constraint.Solver) {
		s.Top(40)
		s.Left(140)
		s.Height(10)
	})

	vm := view.Model{}
	vm.Layouter = l
	vm.Children = l.Views()
	vm.Painter = &paint.Style{BackgroundColor: colornames.White}

	vm.Options = []view.Option{
		pointer.GestureList{&pointer.PressGesture{
			MinDuration: time.Second / 2,
			OnEvent: func(e *pointer.PressEvent) {
				if e.Kind == pointer.EventKindPossible {
					log.Println("Press Possible")
				} else if e.Kind == pointer.EventKindChanged {
					log.Println("Press Changed")
				} else if e.Kind == pointer.EventKindFailed {
					log.Println("Press Failed")
				} else if e.Kind == pointer.EventKindRecognized {
					log.Println("Press Recognized")
					if this.OnTouch != nil {
						this.OnTouch()
					}
				}
			},
		}},
	}

	return vm
}

func setViewGeometry4(s *constraint.Solver, top, left, width, height float64) {
	if top >= 0 {
		s.Top(top)
	}
	if left >= 0 {
		s.Left(left)
	}
	if width >= 0 {
		s.Width(width)
	}
	if height >= 0 {
		s.Height(height)
	}
}

/////////////////
type MessageView struct {
	view.Embed

	msg *ContactMessage

	width   int
	lineNo  int
	lineCnt int

	OnTouch func()
}

func calcLineCnt(s string) int {
	charsPerLine := 21
	slen := len(s)
	n := math.Ceil(float64(slen) / float64(charsPerLine) / 3)
	if true {
		return int(n)
	}
	return 1
}

func NewMessageView(msg *ContactMessage) *MessageView {
	this := &MessageView{}
	this.msg = msg

	this.lineCnt = calcLineCnt(msg.msg)
	return this
}

func (this *MessageView) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}
	l.Solve(func(s *constraint.Solver) {
		s.Height(float64(30 * this.lineCnt))
	})

	icobtn := view.NewImageButton()
	icobtn.Image = application.MustLoadImage("icon_avatar_40")
	icobtn.OnPress = func() {
		// TODO show contact info
	}

	msgtxt := view.NewTextView()
	msgtxt.Style.SetAlignment(text.AlignmentLeft)
	msgtxt.Style.SetWrap(text.WrapWord)
	msgtxt.PaintStyle = &paint.Style{BackgroundColor: colornames.Greenyellow}
	msgtxt.String = gopp.IfElseStr(this.msg.mine, "mine:", "frnd:") + this.msg.msg

	// 自己消息与对方消息的view排列
	if this.msg.mine {
		l.Add(icobtn, func(s *constraint.Solver) {
			setViewGeometry4(s, 0, -1, 30, 30)
			s.RightEqual(l.Right())
		})

		l.Add(msgtxt, func(s *constraint.Solver) {
			setViewGeometry4(s, 0, 40, -1, float64(30*this.lineCnt))
			s.RightEqual(l.Right().Add(-40))
		})
	} else {
		l.Add(icobtn, func(s *constraint.Solver) {
			setViewGeometry4(s, 0, 0, 30, 30)
		})

		l.Add(msgtxt, func(s *constraint.Solver) {
			setViewGeometry4(s, 0, 40, -1, float64(30*this.lineCnt))
			s.RightEqual(l.Right().Add(-40))
		})
	}

	vm := view.Model{}
	vm.Layouter = l
	vm.Children = l.Views()
	vm.Painter = &paint.Style{BackgroundColor: colornames.White}

	return vm
}

///////////
