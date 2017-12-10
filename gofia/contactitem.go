package gofia

import (
	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
)

type ContactItem struct {
	view.Embed

	group  bool
	cnum   uint32
	ctid   string
	ctname string
	status uint32
	stmsg  string
	avatar string
}

func NewContactItem(group bool) *ContactItem {
	this := &ContactItem{}
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

	avabtn := view.NewButton()
	avabtn.String = "AVATAR图"
	l.Add(avabtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 40, 60, 60)
	})

	stsbtn := view.NewButton()
	stsbtn.String = "STS图"
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
	s.Top(top)
	s.Left(left)
	s.Width(width)
	s.Height(height)
}
