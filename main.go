package main

import (
	"flag"
	"tox-homeserver/server"
)

func main() {
	flag.Parse()

	SetupProfile()
	SetupTrace()

	server.SetBuildInfo(trimGOVVV(Version), getBuildInfo(true))
	go server.Main()

	SetupSignal(func() {
		StopProfile()
		StopTrace()
		server.Atexit()
	})

}
