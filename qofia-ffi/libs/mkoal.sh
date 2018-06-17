#!/bin/sh

SYSROOT=/opt/android-ndk/sysroot

# CPP="arm-linux-androideabi-gcc -E" CXXCPP="arm-linux-androideabi-gcc -E" CPPFLAGS="-I${SYSROOT}/usr/include" ./configure --prefix="${SYSROOT}/androidsys" --host=arm-linux-androideabi --target=arm-linux-androideabi CFLAGS="--sysroot $SYSROOT"

export PATH=/opt/andndk16/bin:$PATH
export CC=arm-linux-androideabi-gcc
export CXX=arm-linux-androideabi-g++
export CPP=arm-linux-androideabi-cpp
export CFLAGS="-fPIE -DANDROID -Wno-multichar -D__ANDROID_API__=16"

cmake -DCMAKE_VERBOSE_MAKEFILE=on -DEXAMPLES=off -DCMAKE_INSTALL_PREFIX=/androidsys/ ..

echo ""
echo "configure done. edit link.txt, change soname without version suffix libopenal.so"
sleep 3
make


# source: https://github.com/AerialX/openal-soft-android
