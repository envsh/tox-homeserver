rm -fr ./build ./build-debug
mkdir -p ./build/libs/armeabi-v7a
cp -v libgo.so ./build/
cp -v libgo.so ./build/libs/armeabi-v7a/

time /home/me/Qt5.10.1/5.10.1/android_armv7/bin/androiddeployqt --output ./build/ --verbose --gradle --deployment bundled --android-platform android-26

find -name "*.apk" | xargs ls -lh

# libffi 自己编译的有问题，需要用这个包中的二进制版本，https://github.com/kuri65536/python-for-android/releases?after=r19
# 或者这个包中的libffi.so,  https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/python-for-android/python_r-1.zip

