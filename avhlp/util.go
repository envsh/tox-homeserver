package avhlp

import (
	"fmt"

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
