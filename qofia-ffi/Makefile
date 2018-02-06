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

	go-rcc ./rcc.qrc

bd:
	go install -v tox-homeserver/gofia-ffi
	echo -e "package main\nconst btversion = \"${GOVVV2}\"\n" > btversion.go
	echo -e "const isandroid = true\n" >> btversion.go
	go build -p 1 -v -ldflags "${GOVVV}"

run:
	./qofia-ffi

clean:
	rm -vf ui_*.h ui_*_c.h ui_*.go ui_*.cpp
	rm -vf rcc.cpp *_rc.go
