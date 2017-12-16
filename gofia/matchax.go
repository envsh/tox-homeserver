package gofia

import (
	"runtime"

	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/android"
	"gomatcha.io/matcha/view/ios"
)

// Wrapper for the ios.Stack and android.Stack
type Stack interface {
	SetViews(vs ...view.View)
	Views() []view.View
	Push(vs view.View)
	Pop()
}

func NewStack() Stack {
	var s Stack
	if runtime.GOOS == "android" {
		s = &android.Stack{}
	} else {
		s = &ios.Stack{}
	}
	return s
}

func NewStackView() view.View {
	if runtime.GOOS == "android" {
		return android.NewStackView()
	} else {
		return ios.NewStackView()
	}
}

func NewStackViewWithStack(s Stack) view.View {
	if runtime.GOOS == "android" {
		sv := android.NewStackView()
		sv.Stack = s.(*android.Stack)
		return sv
	} else {
		sv := ios.NewStackView()
		sv.Stack = s.(*ios.Stack)
		return sv
	}
}
