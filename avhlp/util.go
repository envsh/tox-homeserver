package avhlp

import (
	"bytes"
	"fmt"
	"gopp"
	"io/ioutil"
	"os"
	"runtime"

	"golang.org/x/mobile/exp/audio/al"
)

func to_al_format(channels int16, samples int16) int {
	stereo := channels > 1

	switch samples {
	case 16:
		if stereo {
			return al.FormatStereo16
		} else {

			return al.FormatMono16
		}
	case 8:
		if stereo {
			return al.FormatStereo8
		} else {
			return al.FormatMono8
		}
	}
	return -1
}

func alerrstr(eno int32) string {
	switch eno {
	case 0:
		return "AL_NO_ERROR"
	// case al.INVALID_DEVICE:
	//	return "AL_INVALID_DEVICE"
	// case al.INVALID_CONTEXT:
	//	return "AL_INVALID_CONTEXT"
	case al.InvalidName:
		return "AL_INVALID_NAME"
	case al.InvalidEnum:
		return "AL_INVALID_ENUM"
	case al.InvalidValue:
		return "AL_INVALID_VALUE"
	case al.InvalidOperation:
		return "AL_INVALID_OPERATOR"
	case al.OutOfMemory:
		return "AL_OUT_OF_MEMORY"
		/* ... */
	default:

	}
	return fmt.Sprintf("Unknown error code: %d", eno)
}

func findsupath() string {
	paths := []string{"/system/bin/su"}
	for _, p := range paths {
		if gopp.FileExist(p) {
			return p
		}
	}
	return "su"
}

func getLibDirp() string {
	// go arch name => android lib name
	archs := map[string]string{"386": "x86", "amd64": "x86_64", "arm": "arm", "mips": "mips"}

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
