package gofia

import (
	"golang.org/x/image/colornames"
	egview "gomatcha.io/matcha/examples/view"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/view"
)

type ChatFormState = ContactItemState

type ChatFormView struct {
	view.Embed

	cfst *ChatFormState
}

func NewChatFormView() *ChatFormView {
	this := &ChatFormView{}
	this.cfst = &ChatFormState{}

	return this
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
	avtbtn := view.NewButton()
	avtbtn.String = "ICO图"
	hl.Add(avtbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, 50, -1)
	})

	namlab := view.NewTextView()
	namlab.String = "NAM图+"
	hl.Add(namlab, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 50, 80, 30)
	})
	idlab := view.NewTextView()
	idlab.String = "TID图+"
	hl.Add(idlab, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 130, -1, 30)
	})
	stmlab := view.NewTextView()
	stmlab.String = "STM图++++++++++++++++++++"
	hl.Add(stmlab, func(s *constraint.Solver) {
		setViewGeometry4(s, 30, 50, -1, 30)
	})

	//TODO mute/mic/audio/video
	vdobtn := view.NewButton()
	vdobtn.String = "VDO图+"
	hl.Add(vdobtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right())
		setViewGeometry4(s, -1, -1, 50, -1)
		s.RightEqual(hl.Right())
	})
	adobtn := view.NewButton()
	adobtn.String = "ADO图+"
	hl.Add(adobtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right().Add(-50))
		setViewGeometry4(s, 0, -1, 50, -1)
		s.RightEqual(hl.Right().Add(-50))
	})
	mutebtn := view.NewButton()
	mutebtn.String = "M图+"
	hl.Add(mutebtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right().Add(-100))
		setViewGeometry4(s, 0, -1, 30, 30)
		s.RightEqual(hl.Right().Add(-100))
	})
	micbtn := view.NewButton()
	micbtn.String = "C图+"
	hl.Add(micbtn, func(s *constraint.Solver) {
		s.RightEqual(hl.Right().Add(-100))
		setViewGeometry4(s, 30, -1, 30, 30)
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
	ccsv := view.NewScrollView()
	ccsv.ContentLayouter = cctablo
	ccsv.ContentChildren = cctablo.Views()
	ccsv.ScrollAxes = layout.AxisY
	l.Add(ccsv, func(s *constraint.Solver) {
		s.BottomEqual(l.Bottom().Add(-50))
		setViewGeometry4(s, 50, 0, -1, -1)
		s.BottomEqual(l.Bottom().Add(-50))
		s.RightEqual(l.Right())
	})
	////// content layout end

	////// footer layout begin
	fl := &constraint.Layouter{}
	fl.Solve(func(s *constraint.Solver) {
		setViewGeometry4(s, -1, -1, -1, 80)
		s.LeftEqual(l.Left())
		s.BottomEqual(l.Bottom())
		s.RightEqual(l.Right())
	})

	ftipt := view.NewTextInput()
	ftipt.Placeholder = "input hereeee..."
	ftipt.MaxLines = 5
	ftipt.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	ftsvl := &constraint.Layouter{}
	ftsvl.Solve(func(s *constraint.Solver) {
		setViewGeometry4(s, -1, -1, -1, 50)
		// s.TopEqual(fl.Top())
		// s.LeftEqual(fl.Left())
		// s.BottomEqual(fl.Bottom())
		// s.RightEqual(fl.Right().Add(-80))
	})
	ftsvl.Add(ftipt, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, -1, 50)
		s.BottomEqual(ftsvl.Bottom())
		s.RightEqual(ftsvl.Right())
	})
	ftsv := view.NewScrollView()
	ftsv.ContentChildren = []view.View{ftipt}
	ftsv.ContentLayouter = ftsvl
	ftsv.ScrollAxes = layout.AxisY
	//ftsv.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	fl.Add(ftsv, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 0, -1, 50)
		s.TopEqual(fl.Top())
		s.LeftEqual(fl.Left())
		s.BottomEqual(fl.Bottom())
		s.RightEqual(fl.Right().Add(-80))
	})

	ftstbtn := view.NewButton()
	ftstbtn.String = "SNT图+"
	fl.Add(ftstbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, -1, 50, 50)
		s.RightEqual(fl.Right())
	})

	ftemojibtn := view.NewButton()
	ftemojibtn.String = "Emoji图+"
	fl.Add(ftemojibtn, func(s *constraint.Solver) {
		s.RightEqual(fl.Right().Add(-50))
		setViewGeometry4(s, 0, -1, 30, 30)
	})

	ftfilebtn := view.NewButton()
	ftfilebtn.String = "File图+"
	fl.Add(ftfilebtn, func(s *constraint.Solver) {
		s.RightEqual(fl.Right().Add(-50))
		setViewGeometry4(s, 30, -1, 30, 30)
	})

	fv := view.NewBasicView()
	fv.Layouter = fl
	fv.Children = fl.Views()
	l.Add(fv, func(s *constraint.Solver) {
		s.TopEqual(l.Bottom().Add(-50))
		s.BottomEqual(l.Bottom())
		s.RightEqual(l.Right())
		s.LeftEqual(l.Left())
	})
	////// footer layout end

	vml := view.Model{}
	vml.Layouter = l
	vml.Children = l.Views()
	return vml
}

func (v *ChatFormView) Buildgc(ctx view.Context) view.Model {

	vml := view.Model{}
	return vml
}
