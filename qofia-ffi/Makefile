VERSION         :=      $(shell cat ./VERSION)
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`
GOVVV2=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'|sed 's/main./gofia./g'`
CPWD=$(shell pwd)



all:
	usage="make < gen bd clean >"

gen:
	# PATH=/home/me/Qt5.10.1/5.10.1/gcc_64/bin:$PATH
	# go-uic scroll_widget.ui
	go-uic ./contact_item_view.ui
	go-uic ./message_item_view.ui
	go-uic untitled.ui
	go-uic create_room.ui
	go-uic add_friend.ui
	go-uic emoji_panel.ui
	go-uic emoji_category.ui

	cgo-rcc -pkg qrcin ./rcc.qrc
	mv rcc_rc.go qrcin/

# need min 1.2G memory with GOGC=5 and minqt
bd:
	echo -e "package main\nconst btversion = \"${GOVVV2}\"\n" > btversion.go
	echo -e "const isandroid = true\n" >> btversion.go
	# why CC=clang CXX=clang++ can not build. got: must specify -pkgdir
	CC=/usr/bin/clang CXX=/usr/bin/clang++ GOGC=10 \
		go build -p 1 -v -i  -gcflags="-c 1"  -ldflags "${GOVVV}" -pkgdir ~/oss/pkg/linux_amd64_clang
	# GOGC=10 go build -p 1 -v -i -gcflags="-c 1"  -ldflags "${GOVVV}"

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


