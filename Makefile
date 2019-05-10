VERSION         :=      $(shell cat ./VERSION)
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`
GOVVV2=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'|sed 's/main./gofia./g'`
CPWD=$(shell pwd)

all: bd

bd: com
	go-bindata -nocompress -pkg server -o server/webdui_bindata.go webdui/ data/server.crt data/server.key
	PKG_CONFIG_PATH=/opt/toxcore-static2/lib64/pkgconfig/ CGO_LDFLAGS="-lopus -lsodium" \
	CC=/usr/bin/clang CXX=/usr/bin/clang++ \
		go build -p 1 -i -v -o bin/toxhs -ldflags "${GOVVV}" -pkgdir ~/oss/pkg/linux_amd64_clang .
	tar zcvf bin/toxhs.tar.gz bin/toxhs

democ: com
	go build -v -o bin/democ ./demos/native-console-grpc-client/

com:
	protoc -I./server ./server/ths.proto --go_out=plugins=grpc:./thspbs/
	go-bindata -nocompress -pkg transport -o client/transport/certs_bindata.go data/server.crt
	# cd ${HOME}/golib/src/github.com/go-xorm/cmd/xorm && xorm reverse -s sqlite3 "${CPWD}/data/toxhs.sqlite" templates/goxorm "${CPWD}/gofia/"
	go install -v ./thspbs/ ./thscom/ ./client/ ./client/grpctp ./client/websocketstp ./store/ ./gomain2c/
	#go install -v ./avhlp/


allfia: gofiab tofiab tofiai
gofiab: #build
	echo -e "package gofia\nconst btversion = \"${GOVVV2}\"\n" > gofia/btversion.go
	echo -e "const isandroid = true\n" >> gofia/btversion.go
	# matcha build --target android/arm -v -x --ldflags "${GOVVV2}" tox-homeserver/gofia
	matcha build --target android/386 -v -x --ldflags "${GOVVV2}" tox-homeserver/gofia
	ls -l ${HOME}/golib/src/gomatcha.io/matcha/android/matchabridge.aar
	cd ./bin/ && unzip -o ${HOME}/golib/src/gomatcha.io/matcha/android/matchabridge.aar
	ls -l ./bin/jni/armeabi*/

gofiac: # check quickly
	echo -e "package gofia\nconst btversion = \"${GOVVV2}\"\n" > gofia/btversion.go
	echo -e "const isandroid = true\n" >> gofia/btversion.go
	go build -v --ldflags "${GOVVV2}" tox-homeserver/gofia
	ls -l ${HOME}/golib/src/gomatcha.io/matcha/android/matchabridge.aar
	cd ./bin/ && unzip -o ${HOME}/golib/src/gomatcha.io/matcha/android/matchabridge.aar
	ls -l ./bin/jni/armeabi*/

tofiab: # build
	cd tofia && ./gradlew build  --console plain --build-cache --warn build
	find ./tofia -name "*.apk"|xargs ls -lh
tofiai: # install
	adb install -r ./tofia/app/build/build/outputs/apk/debug/app-debug.apk
tofiac: # clean
	rm -vf ./tofia/app/build/build/outputs/apk/*/*.apk
	rm -vf ./tofia/app/build/build/outputs/apk/*/*.apk.tar.gz

emu:
	go build -v -buildmode=c-shared -o bin/libtoxcore.so ./toxemu/
	chmod +x bin/libtoxcore.so
	rm -f ~/.config/tox/tox_save.lock

emuc:
	go build -v -o /tmp/toxemu.out ./toxemu/

dep:
	dep init -v -gopath -no-examples

wc:
	ls *.go server/*.go client/*.go common/*.go qofia-ffi/*.go gomain2c/*.go avhlp/*.go \
			| grep -v ui_ | grep -v rcc_rc.go | xargs wc -l

lint:

pprof:

# go tool pprof ./hyperkube http://172.16.3.232:10251/debug/pprof/profile
# convert -density 1200 dot_away.svg -size 72x72 dot_away_72.png
# rsvg-convert -o transfer.png  transfer.svg
