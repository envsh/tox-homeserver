#!/bin/sh

# note: run in topdir
# TODO full runable package with Qt5Inline and Qt.
# TODO version info

topdir=$(pwd)
logfile=$topdir/packages/package.log
curver=$(cat VERSION)

echo "This will use about 5 minutes, 1G memory, 500M space, wait moment please..."
echo `date` > $logfile

# compile server
function build_server()
{
    echo "Build Linux server..."
    make >> $logfile 2>&1 && mv -v bin/toxhs.tar.gz packages/
}
function package_server()
{
    echo "Package Linux server..."
    cp -v bin/toxhs.tar.gz packages/toxhs-amd64-$curver.tar.gz
    true;
}

# compile client desktop linux
function build_desktop_linux()
{
    cd qofia-ffi
    echo "Build Linux desktop client..."
    make bd >> $logfile 2>&1 && \
        cp -v qofia-ffi ../packages/
    cd $topdir
}
function package_desktop_linux()
{
    echo "Package linux desktop client..."
    cd $topdir/packages
    cp -v ../qofia-ffi/qofia-ffi ./ # test line

    # construct AppDir for AppImage
    pkgdir=qofia-linux-amd64-$curver
    mkdir $pkgdir/usr/bin -p
    strip -v -g qofia-ffi
    mv -v qofia-ffi $pkgdir/usr/bin/
    cp -v /home/me/oss/qt.inline/lnx64/libQt5Inline.so $pkgdir/usr/
    strip -v -g $pkgdir/usr/libQt5Inline.so
    cd $pkgdir/usr/
    linuxdeployqt ./libQt5Inline.so -qmake=/home/me/Qt5.10.1/5.10.1/gcc_64/bin/qmake
    unlink AppRun
    mv -v libQt5Inline.so lib/
    cp -v /usr/lib/libopus.so.0 lib/
    cp -v /usr/lib/libsodium.so.23 lib/
    cp -v /usr/lib/libffi.so.6 lib/

    cd ..
    cp -a ../../tools/app.sh AppRun
    cp -v ../../tools/app.desktop qofia.desktop
    touch qofia.png
    cd ..

    /home/me/Downloads/appimagetool-x86_64.AppImage -v $pkgdir/
    echo "tar bz2..."
    tar jcf $pkgdir.tar.bz2 $pkgdir/
    rm -rf $pkgdir
    true;
}

# compile client android
function build_android()
{
    cd qofia-ffi
    echo "Build Android client..."
    sh make.sh 2 >> $logfile 2>&1 && \
        cp -v ./build/build/outputs/apk/build-debug.apk ../packages/qofia-ffi.apk
    cd $topdir
}
function package_android()
{
    cd $topdir/packages
    cp -v ../qofia-ffi/build/build/outputs/apk/build-debug.apk ./qofia-ffi.apk # test line
    mv -v qofia-ffi.apk qofia-ffi-arm-$curver.apk
    true;
}

# compile client windows
function build_desktop_windows()
{
    cd qofia-ffi
    echo "Build Windows desktop client..."
    sh mkwin.sh >> $logfile 2>&1 && \
        mv -v qofia-ffi-i386.exe.zip qofia-ffi-amd64.exe.zip ../packages/
    cd $topdir
}
function package_desktop_windows()
{
    cd $topdir/packages
    cp -v ../qofia-ffi/qofia-ffi-amd64.exe ./ # test line
    cp -v ../qofia-ffi/qofia-ffi-i386.exe ./ # test line

    pkgdir=qofia-win-amd64-$curver
    mkdir $pkgdir
    mv -v qofia-ffi-amd64.exe $pkgdir/
    cp -a /home/me/oss/qt.inline/packages/qtenv_win64/bin/*.dll $pkgdir/
    #cp -a /home/me/oss/qt.inline/winx64bd/libQt5Inline.so $pkgdir/Qt5Inline.dll
    #x86_64-w64-mingw32-strip -v -g $pkgdir/Qt5Inline.dll
    tar jcvf $pkgdir.tar.bz2 $pkgdir/
    rm -rf $pkgdir

    pkgdir=qofia-win-i386-$curver
    mkdir $pkgdir
    mv -v qofia-ffi-i386.exe $pkgdir/
    cp -a /home/me/oss/qt.inline/packages/qtenv_win32/bin/*.dll $pkgdir/
    #cp -a /home/me/oss/qt.inline/winx32bd/libQt5Inline.so $pkgdir/Qt5Inline.dll
    #i686-w64-mingw32-strip -v -g $pkgdir/Qt5Inline.dll
    tar jcvf $pkgdir.tar.bz2 $pkgdir/
    rm -rf $pkgdir

    true;
}

######

package_server;
package_desktop_linux;
package_desktop_windows;
package_android;

echo `date` >> $logfile

