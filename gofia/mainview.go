package gofia

import (
	"gopp"
	"image/color"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/application"
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
	bridge.RegisterFunc("tox-homeserver/gofia NewGofiaView", func() view.View {
		// Call the TutorialView initializer.
		v := NewTutorialView()
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

	// contacts  []*ContactItem
	// contactsv []view.View
	// mvst *mainViewState
}

// This is our view's initializer.
func NewTutorialView() *TutorialView {
	log.Println("herere", appctx != nil)
	// AppOnCreate()
	this := &TutorialView{}
	this.TextColor = colornames.Red
	this.text = text.New("")
	this.responder = &keyboard.Responder{}
	appctx.logFn = this.logFn

	// this.contacts = make([]*ContactItem, 0)
	// this.contactsv = make([]view.View, 0)
	// this.mvst = &mainViewState{}

	this.registerEvents()
	// appctx.mainV = this
	return this
}

func (v *TutorialView) Lifecycle(from, to view.Stage) {
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

// Similar to React's render function. Views specify their properties and
// children in Build().
func (v *TutorialView) Build(ctx view.Context) view.Model {
	/*
		if appctx.currV != nil {
			return view.Model{Children: []view.View{appctx.currV}}
		}
	*/

	l := &constraint.Layouter{}
	if false { // for test fixed chat form view
		log.Println("111")
		// tpubkey := "398C8161D038FD328A573FFAA0F5FAAF7FFDE5E8B4350E7D15E6AFD0B993FC52"
		cf := NewChatFormView()
		l.Add(cf, func(s *constraint.Solver) {
			s.LeftEqual(l.Left())
			s.TopEqual(l.Top())
			s.RightEqual(l.Right())
			s.BottomEqual(l.Bottom())
		})
		return view.Model{Children: l.Views(), Layouter: l}
	}

	hl := &constraint.Layouter{}
	hl.Solve(func(s *constraint.Solver) {
		s.Height(50)
	})
	setbtn := view.NewImageButton()
	setbtn.Image = application.MustLoadImage("settings")
	setbtn.Image = application.MustLoadImage("applications_system")
	// setbtn := view.NewButton()
	// setbtn.String = "SET?的"
	hl.Add(setbtn, func(s *constraint.Solver) {
		s.Top(0)
		s.Left(0)
		s.Height(50)
		s.Width(50)
	})

	netbtn := view.NewTextView()
	netbtn.String = "NONE"
	if appctx.mvst.netStatus > 0 {
		netbtn.String = gopp.IfElseStr(appctx.mvst.netStatus == 1, "UDP", "TCP")
	}
	hl.Add(netbtn, func(s *constraint.Solver) {
		// setViewGeometry4(s, 0, 50, 50, 50)
		s.Left(50)
	})

	stbtn := view.NewImageButton()
	stbtn.Image = application.MustLoadImage("dot_away_36")
	// stbtn := view.NewButton()
	// stbtn.String = "STS中"
	hl.Add(stbtn, func(s *constraint.Solver) {
		setViewGeometry4(s, 0, 100, 100, 50)
		// s.Top(0)
		// s.Left(50)
		// s.Height(50)
		// s.Width(100)
	})

	nlab := view.NewTextView()
	nlab.String = "名字Tofia"
	nlab.String = appctx.mvst.nickName
	nlab.Style.SetFont(text.DefaultFont(18))
	// nlab.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}
	log.Println("align:", nlab.Style.Alignment())
	hl.Add(nlab, func(s *constraint.Solver) {
		// s.Top(0)
		s.Left(180)
		// s.Height(50)
	})

	addbtn := view.NewImageButton()
	addbtn.Image = application.MustLoadImage("add")
	addbtn.Image = application.MustLoadImage("list_add")
	// addbtn := view.NewButton()
	// addbtn.String = "F加++"
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
	log.Println("contacts:", appctx.contactStates.Size())
	vtable := &table.Layouter{}
	pkids := gopp.IV2Strings(appctx.contactStates.Keys())
	sort.Strings(pkids)
	// TODO 排序？？？
	for _, pkid := range pkids {
		ctisx, _ := appctx.contactStates.Get(pkid)
		ctis := ctisx.(*ContactItemState)
		ctv := NewContactItem(ctis.group)
		ctv.ContactItemState = ctis
		vtable.Add(ctv, nil)
	}
	/*
		for i, ctv := range appctx.contactsv {
			cell := NewTableCell()
			cell.Axis = layout.AxisY
			cell.Index = i
			if false {
				vtable.Add(cell, nil)
			} else {
				vtable.Add(ctv, nil)
			}
		}
	*/
	if appctx.contactStates.Size() < 12 {
		for i := 0; i < 12-appctx.contactStates.Size(); i++ {
			cell := NewTableCell()
			cell.Axis = layout.AxisY
			cell.Index = i
			if true {
				vtable.Add(cell, nil)
			}
		}
	}
	lstwin := view.NewScrollView()
	lstwin.ScrollAxes = layout.AxisY
	lstwin.ScrollPosition = &view.ScrollPosition{}
	lstwin.ContentLayouter = vtable
	lstwin.ContentChildren = vtable.Views()
	lstwin.OnScroll = func(p layout.Point) {
	}
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

//
func init() {
	// Registers a function with the objc bridge. This function returns
	// a view.View, which can be displayed in a MatchaViewController.
	bridge.RegisterFunc("tox-homeserver/gofia OnBackPressed", OnBackPressed)
	bridge.RegisterFunc("tox-homeserver/gofia OnKeyDown", OnKeyDown)
	bridge.RegisterFunc("tox-homeserver/gofia OnKeyUp", OnKeyUp)
	bridge.RegisterFunc("tox-homeserver/gofia OnKeyLongPress", OnKeyLongPress)
	bridge.RegisterFunc("tox-homeserver/gofia OnKeyMultiple", OnKeyMultiple)
}

// Android 点击两次Back退出应用, 2s内
// return 1 for exit main view
func OnBackPressed() int {
	log.Println("hehehhe", appctx.app.Child == nil)
	now := time.Now()
	if now.Sub(lastBackPressed).Seconds() <= 2 {
		clearLastBackPressed()
		return 1
	}

	// 切换view
	if appctx.app.Child == nil { // must main view, 开始检查2次Back
		lastBackPressed = now
	} else {
		clearLastBackPressed() //切换view，重新计数
		appctx.app.Child = nil
		appctx.app.ChildRelay.Signal()
		// appctx.mainV.(*TutorialView).Signal()
	}

	return 0
}

var lastBackPressed time.Time = time.Time{}

func clearLastBackPressed() { lastBackPressed = time.Time{} }

type KeyEvent int

func OnKeyDown(keyCode int, event KeyEvent) bool {
	log.Println(keyCode, event)
	return false
}

func OnKeyUp(keyCode int, event KeyEvent) bool {
	log.Println(keyCode, event)
	return false
}

func OnKeyLongPress(keyCode int, event KeyEvent) bool {
	log.Println(keyCode, event)
	return false
}

func OnKeyMultiple(keyCode int, count int, event KeyEvent) bool {
	log.Println(keyCode, count, event)
	return false
}
