
SYSROOT=/opt/android-ndk/sysroot

# CPP="arm-linux-androideabi-gcc -E" CXXCPP="arm-linux-androideabi-gcc -E" CPPFLAGS="-I${SYSROOT}/usr/include" ./configure --prefix="${SYSROOT}/androidsys" --host=arm-linux-androideabi --target=arm-linux-androideabi CFLAGS="--sysroot $SYSROOT"

export PATH=/opt/andndk16/bin:$PATH
if [[ -f /opt/andndk16/bin/arm-linux-androideabi-gcc ]]; then
    export CC=arm-linux-androideabi-gcc
    export CXX=arm-linux-androideabi-g++
    export CPP=arm-linux-androideabi-cpp

    ./configure --prefix=/opt/androidsys --host arm-linux-androideabi
else # android-x86
    export CC=i686-linux-android-gcc
    export CXX=i686-linux-android-g++
    export CPP=i686-linux-android-cpp

    ./configure --prefix=/opt/androidsys --host i686-linux-android
fi


make && make install


# replace configure:21196 soname_spec=$lt_soname_spec to soname_spec=libffi.so
#xxx broken arm elf, patchelf --set-soname libffi.so /opt/androidsys/lib/libffi.so
