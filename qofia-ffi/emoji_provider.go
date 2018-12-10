package main

import (
	"gopp"
	"log"
	"strings"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

type EmojiElem struct {
	Unicode   string
	Shortname string
}

type _EmojiProvider struct {
	People   []EmojiElem
	Nature   []EmojiElem
	Food     []EmojiElem
	Activity []EmojiElem
	Travel   []EmojiElem
	Objects  []EmojiElem
	Symbols  []EmojiElem
	Flags    []EmojiElem
}

func (this *_EmojiProvider) ToVec() ([]string, [][]EmojiElem) {
	return []string{"people", "nature", "food", "activity", "travel", "objects", "symbols", "flags"},
		[][]EmojiElem{this.People, this.Nature, this.Food, this.Activity,
			this.Travel, this.Objects, this.Symbols, this.Flags}
}
func (this *_EmojiProvider) ToMap() map[string][]EmojiElem {
	return map[string][]EmojiElem{
		"people": this.People, "nature": this.Nature, "food": this.Food, "activity": this.Activity,
		"travel": this.Travel, "objects": this.Objects, "symbols": this.Symbols, "flags": this.Flags}
}
func (this *_EmojiProvider) ShowNames() []string {
	return []string{"Smileys & People", "Animals & Nature", "Food & Drink", "Activity",
		"Travel & Places", "Objects", "Symbols", "Flags"}
}

var EmojiProvider = &_EmojiProvider{
	People:   _EmojiProvider_people,
	Nature:   _EmojiProvider_nature,
	Food:     _EmojiProvider_food,
	Activity: _EmojiProvider_activity,
	Travel:   _EmojiProvider_travel,
	Objects:  _EmojiProvider_objects,
	Symbols:  _EmojiProvider_symbols,
	Flags:    _EmojiProvider_flags,
}

/////
type EmojiCategory struct {
	*Ui_EmojiCategory

	OnEmojiSelected func(string, string)

	items []*qtwidgets.QTableWidgetItem
}

func NewEmojiCategory(category string, emojivec []EmojiElem) *EmojiCategory {
	this := &EmojiCategory{}

	this.Ui_EmojiCategory = NewUi_EmojiCategory2()
	cols := 7
	rows := gopp.IfElseInt(len(emojivec)%cols == 0 || len(emojivec) < cols,
		len(emojivec)/cols, len(emojivec)/cols+1)

	this.TableWidget.HorizontalHeader().SetVisible(false)
	this.TableWidget.VerticalHeader().SetVisible(false)
	this.TableWidget.SetFixedSize1(cols*50+20, rows*50+20)
	this.TableWidget.SetRowCount(rows)
	this.TableWidget.SetColumnCount(cols)
	for i := 0; i < cols; i++ {
		this.TableWidget.SetColumnWidth(i, 50)
	}
	for i := 0; i < rows; i++ {
		this.TableWidget.SetRowHeight(i, 50)
	}

	var emafnt = qtgui.NewQFont1p("Emoji One")
	var itmfnt *qtgui.QFont
	for i, e := range emojivec {
		col := i % cols
		row := i / cols

		witem := qtwidgets.NewQTableWidgetItem1p(e.Unicode)
		witem.SetSizeHint(qtcore.NewQSize1(42, 42))
		witem.SetTextAlignment(qtcore.Qt__AlignCenter | qtcore.Qt__AlignHCenter)
		if itmfnt == nil {
			itmfnt = witem.Font()
			pxsz := itmfnt.PixelSize()
			itmfnt.SetPixelSize(gopp.IfElseInt(pxsz <= 0, 29, pxsz*3))
			emafnt.SetPixelSize(gopp.IfElseInt(pxsz <= 0, 29, pxsz*3))
			if gopp.IsAndroid() {
				itmfnt.SetPixelSize(gopp.IfElseInt(pxsz <= 0, 29, pxsz*3) - 8)
				emafnt.SetPixelSize(gopp.IfElseInt(pxsz <= 0, 29, pxsz*3) - 8)
			}
		}
		witem.SetFont(itmfnt)
		if emafnt.ExactMatch() {
			witem.SetFont(emafnt)
		}
		witem.SetToolTip(e.Shortname)
		this.TableWidget.SetItem(row, col, witem)
		this.items = append(this.items, witem)
		// log.Println(category, i, row, col, e.Unicode)
	}

	qtrt.Connect(this.TableWidget, "cellClicked(int, int)", func(r, c int) {
		idx := r*cols + c
		log.Println("clicked:", r, c, idx, len(emojivec))
		if idx >= len(emojivec) {
			return
		}
		elem := emojivec[idx]
		log.Println("clicked:", r, c, elem.Unicode)
		if this.OnEmojiSelected != nil {
			this.OnEmojiSelected(elem.Unicode, elem.Shortname)
		}
	})

	this.Label.SetText(strings.Title(category))
	return this
}

func (this *EmojiCategory) clear() {
	// disconnect signals
	qtrt.Disconnect(this.TableWidget, "cellClicked(int, int)")
	// set member to nil
	this.OnEmojiSelected = nil
	this.Ui_EmojiCategory = nil
	this.items = nil
}
