// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
		this.tmer.Start(10)
	}
	this.procsmu.Lock()
	defer this.procsmu.Unlock()
	this.procs = append(this.procs, f)
}

var uitrunner *UiThreadRunner
var rpcinch = make(chan string, 128)
var rpcoutch = make(chan string, 128)

func main() {
	flag.Parse()
	log.SetFlags(log.Flags() | log.Lshortfile)

	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	mw = NewUi_MainWindow2()
	mw.QWidget_PTR().Show()

	uitrunner = NewUiThreadRunner()
	uitrunner.Run(func() {})

	go wsrpcproc(rpcinch, rpcoutch)
	go wspushproc()

	qtrt.Connect(mw.PushButton, "clicked(bool)", func(bool) {
		log.Println("hehhehe")
		currtext := mw.ComboBox.CurrentText()
		ct := findContactByName(currtext)
		gopp.NilPrint(ct, currtext)
		if ct != nil {
			msg := mw.LineEdit_2.Text()
			if ct.IsGroup() {
				req := &thspbs.Event{}
				req.Name = "ConferenceSendMessage"
				req.Args = []string{gopp.ToStr(ct.Pnum), "0", msg, ct.Pubkey}
				rpcCallObj(req)
			} else {
				req := &thspbs.Event{}
				req.Name = "FriendSendMessage"
				req.Args = []string{gopp.ToStr(ct.Pnum), msg, ct.Pubkey}
				rpcCallObj(req)
			}
			var line = fmt.Sprintf("me -> %s: %s", ct.Name, msg)
			if false {
				appendOutput(line)
			}
			mw.LineEdit_2.Clear()
		}
	})
	// todo not supported string param
	// qtrt.Connect(mw.ComboBox, "currentTextChanged(const QString &)", func() {})
	qtrt.Connect(mw.ComboBox, "currentIndexChanged(int)", func(idx int) {
		log.Println(idx)
		text := mw.ComboBox.CurrentText()
		cti := findContactByName(text)
		if cti != nil {
			mw.LineEdit.SetText(fmt.Sprintf("%s:%d", cti.Pubkey, cti.Pnum))
		}
	})

	app.Exec()
}

func wspushproc() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/toxhspush"}

	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	} else {
		uitrunner.Run(func() { mw.Label_8.SetText("None") })
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s\n", message)
			processResponse(string(message))
		}
		log.Println("done")
		uitrunner.Run(func() { mw.Label_8.SetText("None") })
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

func rpcCallObj(req interface{}) string {
	data, err := json.Marshal(req)
	gopp.ErrPrint(err)
	return rpcCall(string(data))
}
func rpcCall(req string) string {
	rpcinch <- req
	resp := <-rpcoutch
	return resp
}

func GetBaseInfo() string {
	req := &thspbs.Event{}
	req.Name = "GetBaseInfo"
	data, err := json.Marshal(req)
	gopp.ErrPrint(err)
	return rpcCall(string(data))
}

// should block
func wsrpcproc(rpcinch chan string, rpcoutch chan string) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/toxhsrpc"}

	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	} else {
		uitrunner.Run(func() { mw.Label_8.SetText("None") })
	}
	defer c.Close()
	go func() {
		resp := GetBaseInfo()
		processResponse(resp)
	}()

	for msg := range rpcinch {
		err = c.WriteMessage(websocket.TextMessage, []byte(msg))
		gopp.ErrPrint(err)
		mt, resp, err := c.ReadMessage()
		gopp.ErrPrint(err, mt)
		if err != nil {
			break
		}
		rpcoutch <- string(resp)
	}
	log.Println("done")
}

var gbinfo *thspbs.BaseInfo

func processResponse(data string) {
	resp := &thspbs.Event{}
	err := json.Unmarshal([]byte(data), resp)
	gopp.ErrPrint(err)

	jso := resp
	if jso.Name == "GetBaseInfoResp" {
		binfo := &thspbs.BaseInfo{}
		err := json.Unmarshal([]byte(resp.Args[0]), binfo)
		gopp.ErrPrint(err)
		gbinfo = binfo

		uitrunner.Run(func() {
			mw.Label_2.SetText(binfo.Name)
			mw.Label_4.SetText(binfo.Stmsg)
			mw.Label_8.SetText(gopp.IfElseStr(binfo.ConnStatus > 0, "OK", "None"))

			for _, friendo := range binfo.Friends {
				found := findContact(friendo.Pubkey) != nil
				if !found {
					putContact(friendo.Name, friendo.Pubkey, friendo.Fnum, thspbs.MemberInfo_FRIEND)
					mw.ComboBox.AddItem__(friendo.Name)
				}
			}
			for _, groupo := range binfo.Groups {
				found := findContact(groupo.GroupId) != nil
				if !found {
					putContact(groupo.Title, groupo.GroupId, groupo.Gnum, thspbs.MemberInfo_GROUP)
					mw.ComboBox.AddItem__(groupo.Title)
				}
			}
		})

	} else if jso.Name == "FriendSendMessageResp" {
	} else if jso.Name == "ConferenceSendMessageResp" {
	} else if jso.Name == "FriendMessage" {
		var line = jso.Margs[0] + ": " + jso.Args[1]
		appendOutput(line)
	} else if jso.Name == "ConferenceMessage" {
		var line = jso.Margs[2] + "'s " + jso.Margs[0] + ": " + jso.Args[3]
		appendOutput(line)
		uitrunner.Run(func() {
			found := findContact(jso.Margs[3]) != nil
			if !found {
				putContact(jso.Margs[2], jso.Margs[3], uint32(gopp.MustInt(jso.Args[0])), thspbs.MemberInfo_GROUP)
				mw.ComboBox.AddItem__(jso.Margs[2])
			}
		})
	} else if jso.Name == "ConferenceTitle" {
		var line = jso.Name + " change to " + jso.Args[2] + " by " + jso.Margs[1]
		appendOutput(line)
		putContact(jso.Args[2], jso.Margs[2], uint32(gopp.MustInt(jso.Args[0])), thspbs.MemberInfo_GROUP)
	} else if jso.Name == "ConferenceNamePeerName" {
		var line = jso.Args[2] + " joined in " + jso.Margs[2]
		appendOutput(line)
	} else if jso.Name == "SelfConnectionStatus" {
		var line = jso.Name + " " + jso.Margs[0]
		appendOutput(line)
	} else if jso.Name == "FriendConnectionStatus" {
		var line = jso.Name + " " + jso.Margs[0] + " " + jso.Margs[2]
		appendOutput(line)
	} else if jso.Name == "ConferenceInvite" {
		var line = jso.Name + " to ??? " + " by " + jso.Margs[0]
		appendOutput(line)
	} else if jso.Name == "ConferencePeerListChange" { // TODO leave???
	} else if jso.Name == "FriendSendMessage" {
		var line = "me -> " + jso.Margs[0] + ": " + jso.Args[1]
		appendOutput(line)
	} else if jso.Name == "ConferenceSendMessage" {
		var line = "me -> " + jso.Margs[0] + ": " + jso.Args[2]
		appendOutput(line)
	} else {
		appendOutput(data)
	}

}

func appendOutput(line string) {
	nowt := time.Now()
	nowts := fmt.Sprintf("%d:%d:%d", nowt.Hour(), nowt.Minute(), nowt.Second())
	line = nowts + " " + line
	uitrunner.Run(func() {
		mw.PlainTextEdit.AppendPlainText(line)
		// TODO scroll to bottom
	})
}

var contacts = make(map[string]*thspbs.ContactInfo) // identity =>
func findContact(id string) *thspbs.ContactInfo {
	if minfo, ok := contacts[id]; ok {
		return minfo
	}
	return nil
}
func findContactByName(name string) *thspbs.ContactInfo {
	for _, minfo := range contacts {
		if minfo.Name == name {
			return minfo
		}
	}
	return nil
}
func putContact(name, id string, pnum uint32, mtype thspbs.MemberInfo_MemType) {
	ctinfo := &thspbs.ContactInfo{}
	ctinfo.Name = name
	ctinfo.Pubkey = id
	ctinfo.Pnum = pnum
	ctinfo.Mtype = mtype
	contacts[id] = ctinfo
}
