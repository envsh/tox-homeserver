package main

import (
	"log"

	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

type FilteredTextEdit struct {
	*qtwidgets.QTextEdit
}

func NewFilteredTextEdit(te *qtwidgets.QTextEdit) *FilteredTextEdit {
	this := &FilteredTextEdit{}
	this.QTextEdit = te

	this.setup()
	return this
}

func (this *FilteredTextEdit) setup() {
	te := this.QTextEdit
	qtrt.Connect(te.Document().DocumentLayout(), "documentSizeChanged(const QSizeF &)", func() {
		log.Println("hehehhee")
		te.UpdateGeometry()
	})
	qtrt.Connect(te, "textChanged()", func() {
		log.Println("hehehhee")
	})
	te.InheritKeyPressEvent(func(e *qtgui.QKeyEvent) {
		log.Println("hehehhee", e.Count(), e.Key(), e.Modifiers(), e.Text())
		e.Ignore()
	})
}
