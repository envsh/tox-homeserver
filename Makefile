VERSION         :=      $(shell cat ./VERSION)
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`
GOVVV2=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'|sed 's/main./gofia./g'`

all: bd

bd: com
	go build -v -o bin/toxhs -ldflags "${GOVVV}" .

democ: com
	go build -v -o bin/democ ./examples/

com:
	protoc -I. ths.proto --go_out=plugins=grpc:./thspbs/
	go install -v ./thspbs/ ./common/ ./client/


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
	cd ./bin/ && unzip ${HOME}/golib/src/gomatcha.io/matcha/android/matchabridge.aar
	ls -l ./bin/jni/armeabi*/

tofiab: # build
	cd tofia && ./gradlew build  --console plain --build-cache --warn build
	find ./tofia -name "*.apk"|xargs ls -lh
tofiai: # install
	adb install -r ./tofia/app/build/outputs/apk/debug/app-debug.apk
tofiac: # clean
	rm -vf ./tofia/app/build/outputs/apk/*/*.apk

lint:

pprof:

# go tool pprof ./hyperkube http://172.16.3.232:10251/debug/pprof/profile
# convert -density 1200 dot_away.svg -size 72x72 dot_away_72.png
# rsvg-convert -o transfer.png  transfer.svg
