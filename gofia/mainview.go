package gofia

import (
	"image/color"
	"log"
	"strconv"
	"strings"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/bridge"
	"gomatcha.io/matcha/keyboard"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
)

// UIUIUUIUIUIUIUI
func init() {
	// Registers a function with the objc bridge. This function returns
	// a view.View, which can be displayed in a MatchaViewController.
	bridge.RegisterFunc("tox-homeserver/gofia New", func() view.View {
		// Call the TutorialView initializer.
		v := NewTutorialView()
		v.TextColor = colornames.Red
		return v
	})
}

// Here is our root view.
type TutorialView struct {
	// All components must implement the view.View interface. A basic implementation
	// is provided by view.Embed.
	view.Embed
	TextColor color.Color
	text      *text.Text
	tnum      int
	responder *keyboard.Responder

	contacts  []*ContactItem
	contactsv []view.View
	mvst      *mainViewState
}

// This is our view's initializer.
func NewTutorialView() *TutorialView {
	AppOnCreate()
	this := &TutorialView{}
	this.text = text.New("")
	this.responder = &keyboard.Responder{}
	appctx.logFn = this.logFn

	this.contacts = make([]*ContactItem, 0)
	this.contactsv = make([]view.View, 0)

	this.mvst = &mainViewState{}
	this.registerEvents()
	return this
}

// Similar to React's render function. Views specify their properties and
// children in Build().
func (v *TutorialView) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}

	hl := &constraint.Layouter{}
	hl.Solve(func(s *constraint.Solver) {
		s.Height(50)
	})
	setbtn := view.NewButton()
	setbtn.String = "SET?的"
	hl.Add(setbtn, func(s *constraint.Solver) {
		s.Top(0)
		s.Left(0)
		s.Height(50)
		s.Width(50)
	})

	stbtn := view.NewButton()
	stbtn.String = "STS中"
	hl.Add(stbtn, func(s *constraint.Solver) {
		s.Top(0)
		s.Left(50)
		s.Height(50)
		s.Width(100)
	})

	nlab := view.NewTextView()
	nlab.String = "名字Tofia"
	nlab.String = v.mvst.nickName
	nlab.Style.SetFont(text.DefaultFont(18))
	// nlab.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	log.Println(nlab.Style.Alignment())
	hl.Add(nlab, func(s *constraint.Solver) {
		// s.Top(0)
		s.Left(150)
		// s.Height(50)
	})

	addbtn := view.NewButton()
	addbtn.String = "F加++"
	hl.Add(addbtn, func(s *constraint.Solver) {
		s.Top(0)
		s.RightEqual(hl.Right())
		// s.Height(50)
		s.Width(100)
	})

	hv := view.NewBasicView()
	hv.Layouter = hl
	hv.Children = hl.Views()
	l.Add(hv, func(s *constraint.Solver) {
		s.Top(0)
		s.Left(0)
		s.Height(50)
		s.RightEqual(l.Right())
	})

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
		v.Signal()
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
		s.Top(100)
		s.Left(0)
		s.WidthEqual(l.Width())
		s.Height(200)
	})

	// contacts
	log.Println("contacts:", len(v.contactsv), len(v.contacts))
	vtable := &table.Layouter{}
	for i, ctv := range v.contactsv {
		cell := NewTableCell()
		cell.Axis = layout.AxisY
		cell.Index = i
		if false {
			vtable.Add(cell, nil)
		} else {
			vtable.Add(ctv, nil)
		}
	}
	if len(v.contactsv) < 12 {
		for i := 0; i < 12-len(v.contactsv); i++ {
			cell := NewTableCell()
			cell.Axis = layout.AxisY
			cell.Index = i
			if true {
				vtable.Add(cell, nil)
			}
		}
	}
	lstwin := view.NewScrollView()
	lstwin.ContentLayouter = vtable
	lstwin.ContentChildren = vtable.Views()
	l.Add(lstwin, func(s *constraint.Solver) {
		s.Top(200)
		s.Left(0)
		s.WidthEqual(l.Width())
		s.BottomEqual(l.Bottom())
	})

	// Returns the view's children, layout, and styling.
	return view.Model{
		Children: l.Views(),
		// Children: []view.View{setbtn, stbtn, nlab, child},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Lightgray},
	}
}

func (v *TutorialView) logFn(s string) {
	if v.tnum >= 10 {
		matcha.MainLocker.Lock()
		parts := strings.Split(v.text.String(), "\n")
		tails := parts[1:]
		v.text.SetString(strings.Join(tails, "\n"))
		matcha.MainLocker.Unlock()
		v.tnum = len(tails)
	}

	matcha.MainLocker.Lock()
	ns := v.text.String() + s + "\n"
	v.text.SetString(ns)
	matcha.MainLocker.Unlock()
	v.tnum += 1
	v.Signal()
}

///
type TableCell struct {
	view.Embed
	Axis  layout.Axis
	Index int
}

func NewTableCell() *TableCell {
	return &TableCell{}
}

func (v *TableCell) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}
	l.Solve(func(s *constraint.Solver) {
		if v.Axis == layout.AxisY {
			s.Height(200)
		} else {
			s.Width(200)
		}
	})

	label := view.NewTextView()
	label.String = strconv.Itoa(v.Index)
	l.Add(label, func(s *constraint.Solver) {
	})

	border := view.NewBasicView()
	border.Painter = &paint.Style{BackgroundColor: colornames.Gray}
	l.Add(border, func(s *constraint.Solver) {
		s.Height(1)
		s.LeftEqual(l.Left())
		s.RightEqual(l.Right())
		s.BottomEqual(l.Bottom())
	})

	return view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.White},
	}
}
