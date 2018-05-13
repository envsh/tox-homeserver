rm -fr ./build ./build-debug
mkdir -p ./build/libs/armeabi-v7a
cp -v libgolem.so ./build/
cp -v libgolem.so ./build/libs/armeabi-v7a/

mkdir -p ./build/java ./build/src/debug/java
MYQTDIR=/home/me/Qt5.10.1/5.10.1/
QTSRCDIR=$MYQTDIR/Src
# 如果没有相应的qtandroidextras的java文件，直接加载Qt5AndroidExtras.so会崩溃。
cp -va $QTSRCDIR/qtandroidextras/src/jar/src/* ./build/src/
cp -va $QTSRCDIR/qtconnectivity/src/android/bluetooth/src/* ./build/src/
cp -va $QTSRCDIR/qtconnectivity/src/android/nfc/src/* ./build/src/
cp -va $QTSRCDIR/qtwebview/src/jar/src/* ./build/src/
cp -va $QTSRCDIR/qtpurchasing/src/jar/src/* ./build/src/

time $MYQTDIR/android_armv7/bin/androiddeployqt --output ./build/ --verbose --gradle --deployment bundled --android-platform android-26

#apk compress useless
#tar zcvf ./build/build/outputs/apk/build-debug.apk.tar.gz ./build/build/outputs/apk/build-debug.apk
find -name "*.apk*" | xargs ls -lh

# libffi 自己编译的有问题，需要用这个包中的二进制版本，https://github.com/kuri65536/python-for-android/releases?after=r19
# 或者这个包中的libffi.so,  https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/python-for-android/python_r-1.zip

# -deployment-settings.json extra line: "qml-root-path":"qmlapp",

