package main

import (
	"fmt"

	"github.com/kitech/qt.go/qtcore"
)

func (this *MainWindow) initAboutPage() {
	this.initAboutUi()
	this.initAboutSignals()
	this.initAboutEvents()
}

func (this *MainWindow) initAboutUi() {
	this.Label_37.SetText(Version)
	this.Label_38.SetText(GitCommit)
	this.Label_40.SetText(qtcore.QVersion())
	this.Label_42.SetText(qtcore.QT_VERSION_STR)
}

func (this *MainWindow) initAboutSignals() {
}

func (this *MainWindow) initAboutEvents() {
}

func (this *MainWindow) setCoreVersion(ver string) {
	this.Label_39.SetText(fmt.Sprintf("v%s", ver))
}
