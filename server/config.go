package server

import (
	"flag"
	"sync"
	"time"

	"mkuse/appcm"

	"github.com/spf13/pflag"
)

var metrics_server = ""
var log_level int = 0

func init() {
	flag.StringVar(&metrics_server, "metrics-server", metrics_server, "like: localhost:8086")
	flag.IntVar(&log_level, "log-level", log_level, "[0-5]")
}

type Config struct {
	*appcm.Config
}

var rtcfg *Config
var rtcfgOnce sync.Once

func GetCfg() *Config {
	rtcfgOnce.Do(func() { rtcfg = NewConfig() })
	return rtcfg
}

func NewConfig() *Config {
	this := &Config{}
	this.Config = appcm.GetConfig()
	this.init()
	return this
}

func (this *Config) init() {
	defvpr := this

	defvpr.BindPFlag("metrics_server", pflag.PFlagFromGoFlag(flag.Lookup("metrics-server")))
	defvpr.BindPFlag("log_level", pflag.PFlagFromGoFlag(flag.Lookup("log-level")))
	defvpr.SetDefault("log_level", 0)
	defvpr.SetDefault("log_file", "") // default stderr
	defvpr.SetReadOnly("timezone", true)
	defvpr.SetDefault("timezone", 8)
	defvpr.SetDefault("auto_accept_friend_request", true)
	defvpr.SetDefault("auto_accept_group_invite", true)
	defvpr.SetDefault("auto_accept_recieve_file", true)
	defvpr.SetDefault("auto_follow_someone", true)
	defvpr.SetDefault("file_server_port", 8099)     // ro
	defvpr.SetDefault("rpc_server_port", 2080)      // ro
	defvpr.SetDefault("listen_address", "0.0.0.0")  // ro
	defvpr.SetDefault("transport_cert_pubkey", "")  // ro
	defvpr.SetDefault("transport_cert_privkey", "") // ro

	defvpr.SetDefault("enable_ipv6", false)
	defvpr.SetDefault("enable_udp", true)
	defvpr.SetDefault("enable_tcp_relay", false)
	defvpr.SetDefault("enable_lan_discovery", false)
	defvpr.SetDefault("black_list", "")

	defvpr.SetDefault("proxy_uri", "") // http:// or socks5://
	defvpr.SetDefault("reconnect", "change any char of this value to restart server.")
	defvpr.SetDefault("use_ipfs_file_store", true)

	defvpr.BindPFlag("enable_pprof", pflag.PFlagFromGoFlag(flag.Lookup("pprof")))
	defvpr.BindPFlag("enable_trace", pflag.PFlagFromGoFlag(flag.Lookup("trace")))
	defvpr.SetDefault("uptime", time.Now().String()) // ro
	defvpr.SetDefault("toxcore_version", "0.0.0")
	defvpr.SetDefault("toxhs_version", "0.0.0")
	defvpr.SetDefault("build_info", "")

	defvpr.SetDefault("test_strs1", []string{"1", "2"})
	defvpr.SetDefault("test_ints1", []int{1, 2})

}
