package main

import (
	"fmt"
	"gopp"
	"log"
	"runtime"

	"github.com/kitech/qt.go/qtwidgets"
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

func resolveAppPath() string {
	switch runtime.GOOS {
	case "android":
		for i := 1; i < 9; i++ {
			d := fmt.Sprintf("/data/app/org.qtproject.example.go-%d/", i)
			if gopp.FileExist(d) {
				return d
			}
		}
	}
	return "/thedirshouldnotexists"
}
