package server

import (
	"fmt"
	"gopp"
	"log"
	"net/http"
	"reflect"
	"sort"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper

	defvpr   *viper.Viper
	rootPath string
}

var rtcfg *Config
var rtcfgOnce sync.Once

func GetConfig() *Config {
	rtcfgOnce.Do(func() {
		rtcfg = NewConfig()
	})
	return rtcfg
}

func NewConfig() *Config {
	this := &Config{}
	this.init()
	return this
}

func (this *Config) init() {
	defvpr := viper.New()
	defvpr.SetDefault("timezone", 8)
	defvpr.SetDefault("auto_accept_friend_request", true)
	defvpr.SetDefault("auto_accept_group_invite", true)
	defvpr.SetDefault("auto_accept_recieve_file", true)
	defvpr.SetDefault("file_server_port", 8099)
	defvpr.SetDefault("rpc_server_port", 2080)
	defvpr.SetDefault("listen_address", "0.0.0.0")
	defvpr.SetDefault("transport_cert_pubkey", "")
	defvpr.SetDefault("transport_cert_privkey", "")

	rtvip := Viper2Http(defvpr, "/rtcfg/")
	this.defvpr = defvpr
	this.Viper = rtvip
}

/////
func Viper2Http(defvpr *viper.Viper, rootPath string) (rtvpr *viper.Viper) {
	rtvpr = viper.New()
	for key, valuex := range defvpr.AllSettings() {
		rtvpr.SetDefault(key, valuex)
	}

	// "/rtcfg/" // ?act=list, set,
	http.HandleFunc(rootPath, (&Viper4Http{rootPath, defvpr, rtvpr}).adminHandler)
	return
}

type Viper4Http struct {
	rootPath string
	defvpr   *viper.Viper
	rtvpr    *viper.Viper
}

func (this *Viper4Http) adminHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	gopp.ErrPrint(err)
	log.Println(r.Form)
	log.Println(r.FormValue("act"))
	act := r.FormValue("act")
	if act == "" {
		this.showIndex(w, r)
	} else if act == "reset" {
		this.resetValue(w, r)
	} else if act == "set" {
		this.setValue(w, r)
	}
}

func (this *Viper4Http) showIndex(w http.ResponseWriter, r *http.Request) {
	wrout := func(format string, args ...interface{}) {
		s := fmt.Sprintf(format, args...)
		n, err := w.Write([]byte(s))
		gopp.ErrPrint(err, n)
	}

	wrout("<html><head></head><body>")
	allsets := this.rtvpr.AllSettings()
	wrout("Total Count: %d<br/>", len(allsets))
	keys := this.rtvpr.AllKeys()
	sort.Slice(keys, func(i int, j int) bool { return keys[i] < keys[j] })

	for _, key := range keys {
		valx := this.rtvpr.Get(key)
		wrout("<form method=post action=\"?\">")
		wrout(fmt.Sprintf("%s<input type=hidden name=key value='%s'/> ", key, key))
		wrout(fmt.Sprintf("= <input type=text name=value value='%v'>(default: %v, type: %s)",
			valx, this.defvpr.Get(key), reflect.TypeOf(valx).String()))
		wrout("<input type=submit name=act value='set'/>")
		wrout("<input type=submit name=act value='reset'/>")
		wrout("</form>")
	}

	wrout("</body></html>")
}

func (this *Viper4Http) setValue(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Form)
	err := r.ParseForm()
	gopp.ErrPrint(err)

	key := r.FormValue("key")
	value := r.FormValue("value")
	this.rtvpr.Set(key, value)
	log.Println(key, value)

	err = this.rtvpr.WriteConfigAs("toxhs.toml")
	gopp.ErrPrint(err)
	this.redirectTo(this.rootPath+"?", w)
}

func (this *Viper4Http) resetValue(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Form)
	err := r.ParseForm()
	gopp.ErrPrint(err)

	key := r.FormValue("key")
	value := this.defvpr.Get(key)
	this.rtvpr.Set(key, value)
	log.Println(key, value)

	this.redirectTo(this.rootPath+"?", w)
}

func (this *Viper4Http) redirectTo(link string, w http.ResponseWriter) {
	w.Header().Set("Location", link)
	w.WriteHeader(301)
}
