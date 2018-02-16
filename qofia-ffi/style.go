package main

import (
	"fmt"
	"gopp"
)

const (
	_BACKGROUND = iota + 1
	_TEXT
	_TEXT_SUB // minor text color, lighter than major text
	_HEADER_BG
	_ROOM_ITEM_BG_SELECTED
	_ROOM_ITEM_BG_HOVER
	_MSG_ITEM_BG_SELECTED
)

const (
	_STL_SYSTEM = iota + 1
	_STL_DARK
	_STL_MATERIAL
)

var palettes = map[int]map[int]string{
	_STL_MATERIAL: map[int]string{
		_BACKGROUND:            "#ffffff",
		_TEXT:                  "#000000", // #5f5f5f
		_TEXT_SUB:              "#a4a4a4",
		_HEADER_BG:             "#d6dde3",
		_ROOM_ITEM_BG_SELECTED: "#38A3D8",
		_ROOM_ITEM_BG_HOVER:    "#c8c8c8",
		_MSG_ITEM_BG_SELECTED:  "#d1efff",
	},
	_STL_DARK: map[int]string{
		_BACKGROUND:            "#000000",
		_TEXT:                  "#ffffff",
		_TEXT_SUB:              "#a4a4a4",
		_HEADER_BG:             "#d6dde3",
		_ROOM_ITEM_BG_SELECTED: "#38A3D8",
		_ROOM_ITEM_BG_HOVER:    "#c8c8c8",
		_MSG_ITEM_BG_SELECTED:  "#d1efff",
	},
	_STL_SYSTEM: map[int]string{
		_BACKGROUND:            "#eff0f1",
		_TEXT:                  "#31363b",
		_TEXT_SUB:              "#c4c9cd",
		_HEADER_BG:             "#d6dde3",
		_ROOM_ITEM_BG_SELECTED: "#3daee9",
		_ROOM_ITEM_BG_HOVER:    "#c4c9cd",
		_MSG_ITEM_BG_SELECTED:  "#d1efff",
	},
}

func GetColor(c int) string {
	if s, ok := palettes[_STL_MATERIAL][c]; ok {
		return s
	}
	return "#" + gopp.RandStrHex(6)
}

func GetBg(c int) string { return fmt.Sprintf("background:%s;", GetColor(c)) }
func GetFg(c int) string { return fmt.Sprintf("color:%s;", GetColor(c)) }
