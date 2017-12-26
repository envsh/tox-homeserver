package gofia

import "gomatcha.io/matcha/view"

type SettingView struct {
	view.Embed
}

func newSettingView() *SettingView {
	this := &SettingView{}
	return this
}

func (this *SettingView) Build(ctx view.Context) view.Model {
	return view.Model{}
}
