VERSION         :=      $(shell cat ./VERSION)
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`
GOVVV2=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'|sed 's/main./gofia./g'`
CPWD=$(shell pwd)

all:


gen:
	go build -o /tmp/go-uic  go-uic/main.go
	uic untitled.ui > ui_untitled.h
	/tmp/go-uic ui_untitled.h
	uic untitled2.ui > ui_untitled2.h
	/tmp/go-uic ui_untitled2.h
	uic contact_item_view.ui > ui_contact_item_view.h
	/tmp/go-uic ui_contact_item_view.h
	uic message_item_view.ui > ui_message_item_view.h
	/tmp/go-uic ui_message_item_view.h
	qtrcc

bud:
	go install -v tox-homeserver/gofia
	echo -e "package main\nconst btversion = \"${GOVVV2}\"\n" > btversion.go
	echo -e "const isandroid = true\n" >> btversion.go
	go build -v -ldflags "${GOVVV}"

run:
	./qofia

clean:
	rm -vf ui_*.h ui_*_c.h ui_*.go ui_*.cpp
	rm -vf rcc.cpp rcc_cgo_*.go
