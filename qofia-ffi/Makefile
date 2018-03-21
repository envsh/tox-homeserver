VERSION         :=      $(shell cat ./VERSION)
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`
GOVVV2=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'|sed 's/main./gofia./g'`
CPWD=$(shell pwd)

all:


gen:
	# go-uic scroll_widget.ui
	go-uic ./contact_item_view.ui
	go-uic ./message_item_view.ui
	go-uic untitled.ui
	go-uic create_room.ui
	go-uic add_friend.ui

	go-rcc ./rcc.qrc

bd:
	echo -e "package main\nconst btversion = \"${GOVVV2}\"\n" > btversion.go
	echo -e "const isandroid = true\n" >> btversion.go
	go build -p 1 -v -ldflags "${GOVVV}"

droid: droidso droidapk

droidso:
	sh make.sh
droidapk:
	sh make2.sh

droidins:
	adb install -r ./build/build/outputs/apk/build-debug.apk

run:
	./qofia-ffi

clean:
	rm -vf ui_*.h ui_*_c.h ui_*.go ui_*.cpp
	rm -vf rcc.cpp *_rc.go
