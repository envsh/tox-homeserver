package couchcache

import (
	"gopp"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Get(key string) []byte { return ds.get(key) }

func Set(key string, value []byte) error {
	return ds.set(key, value, 0)
}

var vph *viper.Viper = viper.New()
var rtcfg interface{}
var cfgpath string

func New(path string, rtcfg_ interface{}) *viper.Viper {
	rtvip := viper.New()
	err := rtvip.AddRemoteProvider("memcouch", "123", path)
	gopp.ErrPrint(err)
	rtvip.SetConfigType("json")

	log.Println(rtvip.AllSettings())
	err = rtvip.ReadRemoteConfig()
	gopp.ErrPrint(err)
	log.Println("foo:", rtvip.Get("foo"))

	vph = rtvip
	cfgpath = path
	rtcfg = rtcfg_
	return rtvip
}

func Watch() {
	log.Println(cfgpath)
	err := vph.WatchRemoteConfigOnChannel()
	gopp.ErrPrint(err)

	vph.OnConfigChange(func(in fsnotify.Event) {
		log.Println(in)
		vph.Unmarshal(rtcfg)
	})
}
