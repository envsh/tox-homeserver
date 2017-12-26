package gofia

import (
	"gopp"
	"log"
	"sort"
	"strings"
	"time"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/application"
	"gomatcha.io/matcha/bridge"
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
		// Call the MainView initializer.
		v := NewMainView()
		return v
	})
}

// Here is our root view.
type MainView struct {
	// All components must implement the view.View interface. A basic implementation
	// is provided by view.Embed.
	view.Embed

	// contacts  []*ContactItem
	// contactsv []view.View
	// mvst *mainViewState
}

// This is our view's initializer.
func NewMainView() *MainView {
	log.Println("herere", appctx != nil)
	// AppOnCreate()
	this := &MainView{}
	appctx.logFn = this.logFn

	// this.contacts = make([]*ContactItem, 0)
	// this.contactsv = make([]view.View, 0)
	// this.mvst = &mainViewState{}

	this.registerEvents()
	// appctx.mainV = this
	return this
}

func (v *MainView) Lifecycle(from, to view.Stage) {
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
func (v *MainView) BuildTestView(ctx view.Context) view.Model {
	/*
		if appctx.currV != nil {
			return view.Model{Children: []view.View{appctx.currV}}
		}
	*/

	l := &constraint.Layouter{}

	log.Println("111")
	// tpubkey := "398C8161D038FD328A573FFAA0F5FAAF7FFDE5E8B4350E7D15E6AFD0B993FC52"
	// subv  := NewChatFormView()
	subv := newLogView()
	l.Add(subv, func(s *constraint.Solver) {
		s.LeftEqual(l.Left())
		s.TopEqual(l.Top())
		s.RightEqual(l.Right())
		s.BottomEqual(l.Bottom())
	})

	return view.Model{Children: l.Views(), Layouter: l}
}

func (v *MainView) Build(ctx view.Context) view.Model {
	if false { // for test fixed chat form view
		return v.BuildTestView(ctx)
	}

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
	setbtn.Image = application.MustLoadImage("barbuttonicon_set")
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

	// contacts
	log.Println("contacts:", appctx.contactStates.Size())
	vtable := &table.Layouter{}
	vtable.StartEdge = layout.EdgeTop
	pkids := gopp.IV2Strings(appctx.contactStates.Keys())
	sort.Strings(pkids)
	// TODO 排序？？？
	for _, pkid := range pkids {
		ctisx, _ := appctx.contactStates.Get(pkid)
		ctis := ctisx.(*ContactItemState)
		ctv := NewContactItem(ctis.group)
		ctv.ctis = ctis
		vtable.Add(ctv, nil)
	}
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
	log.Println("Herehere")
	lstwin := view.NewScrollView()
	lstwin.ScrollAxes = layout.AxisY
	lstwin.ScrollPosition = &view.ScrollPosition{}
	lstwin.ContentLayouter = vtable
	lstwin.ContentChildren = vtable.Views()
	lstwin.OnScroll = func(p layout.Point) {
	}
	guide := l.Add(lstwin, func(s *constraint.Solver) {
		s.TopEqual(l.Top().Add(52))
		s.LeftEqual(l.Left())
		s.RightEqual(l.Right())
		s.BottomEqual(l.Bottom())
		s.WidthEqual(l.Width())
		log.Println(l.Width())
		// s.Debug()
	})
	log.Println(guide)
	log.Println(ctx)

	// Returns the view's children, layout, and styling.
	return view.Model{
		Children: l.Views(),
		// Children: []view.View{setbtn, stbtn, nlab, child},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Lightgray},
	}
}

func (v *MainView) logFn(s string) {
	lst := appctx.logState
	if lst.tnum >= 10 {
		matcha.MainLocker.Lock()
		parts := strings.Split(lst.text.String(), "\n")
		tails := parts[1:]
		lst.text.SetString(strings.Join(tails, "\n"))
		matcha.MainLocker.Unlock()
		lst.tnum = len(tails)
	}

	matcha.MainLocker.Lock()
	ns := lst.text.String() + s + "\n"
	lst.text.SetString(ns)
	matcha.MainLocker.Unlock()
	lst.tnum += 1

	// appctx.signalProperView(nil, true)
	// v.Signal()
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
		// appctx.mainV.(*MainView).Signal()
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
