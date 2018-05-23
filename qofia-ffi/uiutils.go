package main

import "github.com/kitech/qt.go/qtwidgets"

func Repolish(w qtwidgets.QWidget_ITF) {
	syl := w.QWidget_PTR().Style()
	syl.Unpolish(w)
	syl.Polish(w)
}
