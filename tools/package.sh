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
    # cp -v /usr/lib/libsodium.so.23 lib/
    cp -v /usr/lib/libffi.so.6 lib/
    # make sure keep GLIBC_2.14
    cp -a ../../libs/libsodium.so* lib/
    cp -a ../../libs/libopenal.so* lib/
    cp -a ../../libs/libav*.so* lib/
    cp -a ../../libs/libsw*.so* lib/

    cd ..
    cp -a ../../tools/AppRun ./
    cp -v ../../tools/app.desktop qofia.desktop
    cp -v ../../tools/qt.conf usr/bin/
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
    mv -v qofia-ffi.apk qofia-ffi-arm7-$curver.apk
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
winlibs="libsoxr.dll xvidcore.dll libx265.dll libwebpmux-3.dll libwebp-7.dll libvpx.dll.5.0.0 libvorbisenc-2.dll libvorbis-0.dll libtheora-0.dll libtheoraenc-1.dll libtheoradec-1.dll libspeex-1.dll libopenjp2-7.dll libopencore-amrwb-0.dll libopencore-amrnb-0.dll libmp3lame-0.dll liblzma-5.dll libgsm.dll.1.0.14 libfdk-aac-1.dll libxml2-2.dll libssh.dll libmodplug-1.dll libgnutls-30.dll libgmp-10.dll libbluray-2.dll libopus-0.dll libx264-152.dll libgcrypt-20.dll libgpg-error-0.dll libtasn1-6.dll libp11-kit-0.dll libnettle-6.dll libhogweed-4.dll libogg-0.dll SDL2.dll libass-9.dll libfontconfig-1.dll libfribidi-0.dll libvidstab.dll"
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
    cp -a /usr/x86_64-w64-mingw32/bin/avcodec-58.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/avdevice-58.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/avformat-58.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/avutil-56.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/swresample-3.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/swscale-5.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/avfilter-7.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/avresample-4.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/postproc-55.dll $pkgdir/
    cp -a /usr/x86_64-w64-mingw32/bin/libx264-152.dll $pkgdir/libx264-148.dll
    for winlib in $winlibs; do
        cp -va /usr/x86_64-w64-mingw32/bin/$winlib $pkgdir/
    done
    tar jcvf $pkgdir.tar.bz2 $pkgdir/
    rm -rf $pkgdir

    pkgdir=qofia-win-i386-$curver
    mkdir $pkgdir
    mv -v qofia-ffi-i386.exe $pkgdir/
    cp -a /home/me/oss/qt.inline/packages/qtenv_win32/bin/*.dll $pkgdir/
    #cp -a /home/me/oss/qt.inline/winx32bd/libQt5Inline.so $pkgdir/Qt5Inline.dll
    #i686-w64-mingw32-strip -v -g $pkgdir/Qt5Inline.dll
    cp -a /usr/i686-w64-mingw32/bin/avcodec-58.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/avdevice-58.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/avformat-58.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/avutil-56.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/swresample-3.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/swscale-5.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/avfilter-7.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/avresample-4.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/postproc-55.dll $pkgdir/
    cp -a /usr/i686-w64-mingw32/bin/libx264-152.dll $pkgdir/libx264-148.dll
    for winlib in $winlibs; do
        cp -va /usr/i686-w64-mingw32/bin/$winlib $pkgdir/
    done
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

