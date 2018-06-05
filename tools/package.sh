#!/bin/sh

# note: run in topdir
# TODO full runable package with Qt5Inline and Qt.
# TODO version info

topdir=$(pwd)
logfile=$topdir/packages/package.log

echo "This will use about 5 minutes, 1G memory, 300M space, wait moment please..."
echo `date` > $logfile
# compile server
echo "Build Linux server..."
make >> $logfile 2>&1 && mv -v bin/toxhs.tar.gz packages/

# compile client desktop linux
cd qofia-ffi
echo "Build Linux desktop client..."
make bd >> $logfile 2>&1 && \
    zip qofia-ffi.zip qofia-ffi && \
    mv -v qofia-ffi.zip ../packages/

# compile client android
echo "Build Android client..."
sh make.sh 2 >> $logfile 2>&1 && \
    cp -v ./build/build/outputs/apk/build-debug.apk qofia-ffi.apk && \
    zip qofia-ffi.apk.zip qofia-ffi.apk && \
    mv -v qofia-ffi.apk.zip ../packages/ && \
    rm -fv qofia-ffi.apk

# compile client windows
echo "Build Windows desktop client..."
sh mkwin.sh >> $logfile 2>&1 && \
    mv -v qofia-ffi-i386.exe.zip qofia-ffi-amd64.exe.zip ../packages/

echo `date` >> $logfile

