package main

import (
	"gopp"
	"io/ioutil"
	"log"
	"time"
	"tox-homeserver/couchcache"
	_ "tox-homeserver/couchcache"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// 变量名必须要大写才能解析出来
type MixConfig struct {
	Foo string
	Bar int
	Baz bool
}

func main() {
	go couchcache.Main()
	time.Sleep(1 * time.Second)
	go func() {
		path := "a.toml"
		path = "mat.toml"
		bcc, err := ioutil.ReadFile(path)
		gopp.ErrPrint(err)
		if false {
			couchcache.Set(path, []byte(`{"foo": "hehehhe", "bar":123456, "baz": false}`))
			couchcache.Set(path, bcc)
		}

		rtcfg := MixConfig{}
		rtvip := viper.New()
		rtvip.SetConfigFile(path)
		rtvip.SetConfigType("toml")
		err = rtvip.AddRemoteProvider("memcouch", "123", path)
		gopp.ErrPrint(err)

		err = rtvip.ReadInConfig()
		gopp.ErrPrint(err)
		log.Println(rtvip.AllKeys())

		err = rtvip.ReadRemoteConfig()
		gopp.ErrPrint(err)
		log.Println("foo:", rtvip.Get("foo"))
		log.Println(rtvip.AllKeys())
		log.Println(rtvip.Get("gateway"))

		jww.SetLogThreshold(jww.LevelInfo)
		err = rtvip.WriteConfigAs("rtsave.toml")
		gopp.ErrPrint(err)
		// err = rtvip.SafeWriteConfig()
		// gopp.ErrPrint(err)

		rtvip.Unmarshal(&rtcfg)
		log.Printf("%+v\n", rtcfg)

		err = rtvip.WatchRemoteConfigOnChannel()
		gopp.ErrPrint(err)
		log.Println("Zzzzzzz...")
		for {
			time.Sleep(5 * time.Second)
			log.Printf("%+v\n", rtcfg)
			log.Println("foo:", rtvip.Get("foo"))
			if false {
				break
			}
		}
		select {}
	}()
	time.Sleep(1 * time.Hour)
}
