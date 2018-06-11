#!/bin/sh

export GOOS=windows
export GOARCH=386
export CGO_ENABLED=1
export CC=/usr/bin/i686-w64-mingw32-gcc

echo -e "package main\nconst btversion = \"${GOVVV2}\"\n" > btversion.go
echo -e "const isandroid = true\n" >> btversion.go
go build -i -p 1 -v -pkgdir ~/oss/pkg/win-386 -ldflags "${GOVVV}"
if [ x"$?" == x"0" ];then
    mv -v qofia-ffi.exe qofia-ffi-i386.exe
    # zip qofia-ffi-i386.exe.zip qofia-ffi-i386.exe
fi

export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CC=/usr/bin/x86_64-w64-mingw32-gcc

go build -i -p 1 -v -pkgdir ~/oss/pkg/win-amd64 -ldflags "${GOVVV}"
if [ x"$?" == x"0" ];then
    mv -v qofia-ffi.exe qofia-ffi-amd64.exe
    # zip qofia-ffi-amd64.exe.zip qofia-ffi-amd64.exe
fi


