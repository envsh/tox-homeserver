package gofia

import (
	"log"

	"gomatcha.io/matcha/bridge"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/view"
)

func init() {
	bridge.RegisterFunc("tox-homeserver/gofia New", func() view.View {
		app := NewApp()
		return NewAppView(app)
	})
}

type App struct {
	Stack Stack

	ChildRelay *comm.Relay
	Child      view.View
}

func NewApp() *App {
	log.Println("hehrere")
	AppOnCreate()
	app := &App{}
	app.Stack = NewStack()
	app.ChildRelay = &comm.Relay{}

	v := NewMainView()
	app.Stack.SetViews(v)

	return app
}

type AppView struct {
	view.Embed

	app     *App
	backKey comm.Id
}

func NewAppView(app *App) *AppView {
	log.Println("hehrere")
	this := &AppView{}
	this.Embed = view.NewEmbed(app)
	this.app = app
	// appctx.mainV = this
	appctx.app = app
	return this
}

var InterBackRelay *comm.Relay

func (v *AppView) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageMounted) {
		InterBackRelay = &comm.Relay{}
		v.backKey = InterBackRelay.Notify(func() {
			v.app.Stack.Pop()
			v.app.Child = nil
			v.app.ChildRelay.Signal()
		})
		v.Subscribe(v.app.ChildRelay)
	} else if view.ExitsStage(from, to, view.StageMounted) {
		InterBackRelay.Unnotify(v.backKey)
		InterBackRelay = nil
		v.Unsubscribe(v.app.ChildRelay)
	}
}

func (v *AppView) Build(ctx view.Context) view.Model {
	// If user has selected an example, display it.
	if v.app.Child != nil {
		return view.Model{Children: []view.View{v.app.Child}}
	}

	// Otherwise display the stack view
	// var stack view.View = NewStackViewWithStack(v.app.Stack)
	// return view.Model{Children: []view.View{stack}}
	nv := NewMainView()
	return view.Model{Children: []view.View{nv}}
}
