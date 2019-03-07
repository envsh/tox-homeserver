#!/bin/sh

SYSROOT=/opt/android-ndk/sysroot

# CPP="arm-linux-androideabi-gcc -E" CXXCPP="arm-linux-androideabi-gcc -E" CPPFLAGS="-I${SYSROOT}/usr/include" ./configure --prefix="${SYSROOT}/androidsys" --host=arm-linux-androideabi --target=arm-linux-androideabi CFLAGS="--sysroot $SYSROOT"

export PATH=/opt/andndk16/bin:$PATH
if [[ -f /opt/andndk16/bin/arm-linux-androideabi-gcc ]]; then
    export CC=arm-linux-androideabi-gcc
    export CXX=arm-linux-androideabi-g++
    export CPP=arm-linux-androideabi-cpp
else # android-x86
    export CC=i686-linux-android-gcc
    export CXX=i686-linux-android-g++
    export CPP=i686-linux-android-cpp
fi
export CFLAGS="-fPIE -DANDROID -Wno-multichar -D__ANDROID_API__=16"

cmake -DCMAKE_VERBOSE_MAKEFILE=on -DEXAMPLES=off -DCMAKE_INSTALL_PREFIX=/opt/androidsys/ .

echo ""
echo "configure done. edit link.txt, change soname without version suffix libopenal.so"
sleep 3
# make && make install 


# source: https://github.com/AerialX/openal-soft-android
