package couchcache

import (
	"bytes"
	"io"
	"log"

	"github.com/spf13/viper"
)

var rcp *remoteConfigProvider

func init() {
	viper.SupportedRemoteProviders = append(viper.SupportedRemoteProviders, "memcouch")
	rcp = &remoteConfigProvider{}
	rcp.dCh = make(chan string)
	viper.RemoteConfig = rcp
	log.Println("inited rtflag pkg", viper.SupportedRemoteProviders)
}

type remoteConfigProvider struct {
	dCh         chan string
	isWatchChan bool
}

func (rc *remoteConfigProvider) changed(k string) {
	log.Println(rc.isWatchChan)
	if rc.isWatchChan {
		rc.dCh <- k
	}
}

func (rc *remoteConfigProvider) Get(rp viper.RemoteProvider) (io.Reader, error) {
	// couchcache.Get(rp.Path())
	log.Println(rp.Path())
	buf := bytes.NewBuffer([]byte{})
	d := ds.get(rp.Path())
	buf.Write(d)
	log.Println(string(d))
	return buf, nil
}

func (rc remoteConfigProvider) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	log.Println(rp.Path())
	return nil, nil
}

func (rc *remoteConfigProvider) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	log.Println(rp.Path())
	rc.isWatchChan = true

	quit := make(chan bool)
	quitwc := make(chan bool)
	viperResponsCh := make(chan *viper.RemoteResponse)
	// need this function to convert the Channel response form crypt.Response to viper.Response
	go func(cr <-chan string, vr chan<- *viper.RemoteResponse, quitwc <-chan bool, quit chan<- bool) {
		log.Println("process watch channel:", rp.Provider(), rp.Path())
		defer func() { rc.isWatchChan = false }()

		for {
			select {
			case <-quitwc:
				log.Println("quit watch chan", rp.Provider(), rp.Path())
				quit <- true
				return
			case resp := <-cr:
				log.Println("changed", resp)
				if resp != rp.Path() {
					break
				}

				value := ds.get(resp)
				log.Println("changed", resp, string(value))
				vr <- &viper.RemoteResponse{
					Error: nil,
					Value: value,
				}
			}
		}
	}(rc.dCh, viperResponsCh, quitwc, quit)

	return viperResponsCh, quitwc

	// return nil, nil
}

////////////
