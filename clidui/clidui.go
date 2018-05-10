// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"gopp"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
	"tox-homeserver/thspbs"

	"github.com/gorilla/websocket"
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

var addr = flag.String("addr", "localhost:8089", "http service address")

var mw *Ui_MainWindow
var contacts = make(map[string]string)

type UiThreadRunner struct {
	tmer    *qtcore.QTimer
	procs   []func()
	procsmu sync.Mutex
}

func NewUiThreadRunner() *UiThreadRunner {
	this := &UiThreadRunner{}
	this.tmer = qtcore.NewQTimer(nil)
	this.procs = make([]func(), 0)

	qtrt.Connect(this.tmer, "timeout()", func() {
		var f func()
		this.procsmu.Lock()
		if len(this.procs) > 0 {
			f = this.procs[0]
			this.procs = this.procs[1:]
		}
		this.procsmu.Unlock()
		if f != nil {
			log.Println("hehehhe")
			f()
		}
	})

	return this
}

func (this *UiThreadRunner) Run(f func()) {
	if !this.tmer.IsActive() {
		this.tmer.Start(3)
	}
	this.procsmu.Lock()
	defer this.procsmu.Unlock()
	this.procs = append(this.procs, f)
}

var uitrunner *UiThreadRunner

func main() {
	flag.Parse()
	log.SetFlags(0)

	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw = NewUi_MainWindow2()
	mw.QWidget_PTR().Show()
	uitrunner = NewUiThreadRunner()
	uitrunner.Run(func() {})
	go wsproc()
	app.Exec()
}

func wsproc() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/toxhs"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s\n", message)
			uitrunner.Run(func() { mw.PlainTextEdit.AppendPlainText(string(message)) })
			evt := &thspbs.Event{}
			err = json.Unmarshal(message, evt)
			gopp.ErrPrint(err)
			if err == nil {
				switch evt.Name {
				case "FriendMessage":
					if _, ok := contacts[evt.Margs[0]]; !ok {
						uitrunner.Run(func() {
							contacts[evt.Margs[0]] = evt.Margs[1]
							mw.ComboBox.AddItem__(evt.Margs[0])
						})
					}
				}
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			if false {
				err := c.WriteMessage(websocket.TextMessage, []byte("hehe:"+t.String()))
				if err != nil {
					log.Println("write:", err)
					return
				}
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
