package main

import (
	"log"

	"github.com/kitech/qt.go/qtandroidextras"
)

func hello() {
	log.Println(qtandroidextras.QAndroidBinder__Normal)
}

func testRunOnAndroidThread() {
	qtandroidextras.RunOnAndroidThread(func() {
		log.Println("this is run on android thread.")
	})
}
