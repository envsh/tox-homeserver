package common

import (
	"bytes"
	"fmt"
	"gopp"
	"io/ioutil"
	"os"
	"runtime"
)

// libdir: /data/app/appname[-n]/lib/<archname>/
// datadir: /data/data/appname[-n]/

var archs = map[string]string{"386": "x86", "amd64": "x86_64", "arm": "arm", "mips": "mips"}

func AndroidGetLibDir() string {
	switch runtime.GOOS {
	case "android":
		bcc, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", os.Getpid()))
		gopp.ErrPrint(err)
		appdir := string(bcc[:bytes.IndexByte(bcc, 0)])
		for i := 0; i < 9; i++ {
			d := fmt.Sprintf("/data/app/%s%s/lib/%s/", appdir,
				gopp.IfElseStr(i == 0, "", fmt.Sprintf("-%d", i)), archs[runtime.GOARCH])
			if gopp.FileExist(d) {
				return d
			}
		}
	}
	return ""
}

// dsn = fmt.Sprintf("file:///data/data/io.dnesth.tofia/toxhs.sqlite")
func AndroidGetDataDir() string {
	switch runtime.GOOS {
	case "android":
		bcc, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", os.Getpid()))
		gopp.ErrPrint(err)
		appdir := string(bcc[:bytes.IndexByte(bcc, 0)])
		for i := 0; i < 9; i++ {
			d := fmt.Sprintf("/data/data/%s%s/", appdir,
				gopp.IfElseStr(i == 0, "", fmt.Sprintf("-%d", i)))
			if gopp.FileExist(d) {
				return d
			}
		}
	}
	return ""
}

// org.somegroup.appname
func AndroidGetAppProcName() string {
	switch runtime.GOOS {
	case "android":
		bcc, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", os.Getpid()))
		gopp.ErrPrint(err)
		appdir := string(bcc[:bytes.IndexByte(bcc, 0)])
		return appdir
	}
	return ""
}
