package main

import (
	"flag"
	"tox-homeserver/server"
)

func main() {
	flag.Parse()

	SetupProfile()
	SetupTrace()

	go server.Main()

	SetupSignal(func() {
		StopProfile()
		StopTrace()
		server.Atexit()
	})

}
