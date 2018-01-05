package gofia

import (
	"gopp"

	"gomatcha.io/matcha/application"
)

var iconCache = make(map[string]*application.ImageResource)

func getIconRef(name string) *application.ImageResource {
	if rc, ok := iconCache[name]; ok {
		return rc
	}

	rc, err := application.LoadImage(name)
	gopp.ErrPrint(err, name)
	if err != nil {
		iconCache[name] = rc
	}
	return rc
}

// TODO load from real avatar
func getAvatarRef(name string) *application.ImageResource {
	return nil
}

type SvgResource struct {
}

func init() {
	//	var rc *application.ImageResource
}
