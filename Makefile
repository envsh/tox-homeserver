VERSION         :=      $(shell cat ./VERSION)
GOVVV=`govvv -flags -version ${VERSION}|sed 's/=/=GOVVV-/g'`

all: bd

bd:
	go build -v -o toxhs -ldflags "${GOVVV}" .
