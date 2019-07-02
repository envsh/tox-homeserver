package main

import (
	"log"
	"runtime"
	"time"

	"tox-homeserver/gomain2c"
)

// for android, but other OS still ok
func init() { gomain2c.Set(main) }

func main() {
	log.Println("Enter main...")

	if runtime.GOOS == "linux" {
		SetupGops()
	}
	dm = newDaemon()
	dm.appctx.OpenStrorage()

	go func() {
		time.Sleep(2 * time.Second)
		// app.Exit(0)
	}()
	// Execute app
	qofiaui_main()
	// select {}
}

func DumpCallers(pcs []uintptr) {
	log.Println("DumpCallers...", len(pcs))
	for idx, pc := range pcs {
		pcfn := runtime.FuncForPC(pc)
		file, line := pcfn.FileLine(pc)
		log.Println(idx, pcfn.Name(), file, line)
	}
	if len(pcs) > 0 {
		log.Println()
	}
}
