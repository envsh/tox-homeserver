package main

import (
	"log"

	"qt.go/qtwidgets"
)

func dbgpolicy(name string, p *qtwidgets.QSizePolicy) {
	log.Println(name, p.ControlType(),
		p.ExpandingDirections(),
		p.HorizontalPolicy(),
		p.VerticalPolicy(),
		p.HorizontalStretch(),
		p.VerticalStretch(),
		p.RetainSizeWhenHidden())

}
