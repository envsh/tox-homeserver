package gofia

import (
	"gopp"
	"log"
	"time"

	"github.com/kitech/godsts/lists/arraylist"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/application"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
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
	msgs *arraylist.List
}
type ContactItem struct {
	view.Embed

	*ContactItemState
}

func NewContactItem(group bool) *ContactItem {
	this := &ContactItem{}
	this.ContactItemState = &ContactItemState{msgs: arraylist.New()}
	this.group = group

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
	avtbtn.Image = application.MustLoadImage("ic_launcher")
	avtbtn.Image = application.MustLoadImage("identity")
	if this.group {
		//avtbtn.Image = application.MustLoadImage("group")
	} else {
		//avtbtn.Image = application.MustLoadImage("contact")
	}
	avtbtn.OnPress = func() {
		log.Println("clicked:", this.ContactItemState)
		log.Println("view path:", ctx.Path())

		// always new view
		cf := NewChatFormView()
		cf.cfst = this.ContactItemState
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
	stsbtn.Image = application.MustLoadImage("dot_online22")
	// stsbtn := view.NewButton()
	// stsbtn.String = "STS图"
	stsbtn.OnPress = func() {
		key := this.ctid
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
	titlab.String = this.ctname + "." + this.ctid[:5]
	l.Add(titlab, func(s *constraint.Solver) {
		// setViewGeometry4(s, 0, 140, 60, 20)
		s.Top(0)
		s.Left(140)
		s.Height(20)
	})

	stslab := view.NewTextView()
	stslab.String = "STS TEXT字"
	stslab.String = this.stmsg
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

//
type MessageView struct {
	view.Embed

	msg *ContactMessage
}

func NewMessageView(msg *ContactMessage) *MessageView {
	this := &MessageView{}
	this.msg = msg

	return this
}

func (this *MessageView) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}
	l.Solve(func(s *constraint.Solver) {
		s.Height(20)
	})

	msgtxt := view.NewTextView()
	msgtxt.String = gopp.IfElseStr(this.msg.mine, "mine:", "frnd:") + this.msg.msg
	l.Add(msgtxt, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, -1, 20)
	})

	vm := view.Model{}
	vm.Layouter = l
	vm.Children = l.Views()
	vm.Painter = &paint.Style{BackgroundColor: colornames.White}

	return vm
}

///////////
