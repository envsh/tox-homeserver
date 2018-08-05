package main

import (
	"fmt"
	"gopp"
	"log"
	"runtime"
	"time"

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

// 中文常用格式
func Time2Today(t time.Time) string {
	return t.Format("15:04:05")
}

func Time2TodayMinute(t time.Time) string {
	return t.Format("15:04")
}

func condWait(timeoutms int, f func() bool) {
	for {
		time.Sleep(time.Duration(timeoutms) * time.Millisecond)
		if f() {
			break
		}
	}
}

func CallStateString(state uint32) string {
	states := map[uint32]string{
		0: "NONE", 1: "ERROR", 2: "FINISH", 4: "SENDING_A", 8: "SENDING_V",
		16: "ACCEPT_A", 17: "ACCEPT_V",
	}
	if s, ok := states[state]; ok {
		return s
	}
	return "Unknown"
}

// more special symbols: http://www.fhdq.net/
func NameNumSep() string {
	// fmt.Sprintf(" (%d)", len(ct.Members)),
	// fmt.Sprintf(" [%d]", len(ct.Members)),
	// fmt.Sprintf(" <%d>", len(ct.Members)),
	// fmt.Sprintf(" -%d", len(ct.Members)),
	// fmt.Sprintf(" +%d", len(ct.Members)),
	// fmt.Sprintf(" | %d", len(ct.Members)),
	// fmt.Sprintf(" %%%d", len(ct.Members)),
	// fmt.Sprintf(" Σ%d", len(ct.Members)),
	// fmt.Sprintf(" √%d", len(ct.Members)),
	// fmt.Sprintf(" ·%d", len(ct.Members)),
	// fmt.Sprintf(" •%d", len(ct.Members)),
	// fmt.Sprintf(" ●%d", len(ct.Members)),
	// fmt.Sprintf(" ◦%d", len(ct.Members)),
	// fmt.Sprintf(" ○%d", len(ct.Members)),
	return "◦"
}
