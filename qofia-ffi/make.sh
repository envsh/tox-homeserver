source ~/triline/shell/android-ndk-env.sh
source ~/triline/shell/android-go-env.sh
# export CGO_LDFLAGS="$CGO_LDFLAGS -lopenal"

env | grep CGO_
set -x


go install -v -i -pkgdir ~/oss/pkg/android_arm github.com/kitech/qt.go/qtqt
go install -v -i -pkgdir ~/oss/pkg/android_arm github.com/kitech/qt.go/qtrt
go install -v -i -pkgdir ~/oss/pkg/android_arm github.com/mattn/go-sqlite3

# go build -p 1 -v  -pkgdir ~/oss/pkg/android_arm .
rm -vf libmain.so
# time go install -p 1 -v -i  -pkgdir ~/oss/pkg/android_arm tox-homeserver/gofia
time go build -p 1 -v -i -pkgdir ~/oss/pkg/android_arm -buildmode=c-shared -o libmain.so .
chmod +x libmain.so

$CC -xc andwrapmain.c.nogo -shared   -o libgolem.so -lmain -L. -Wl,-soname,libgolem.so
ccret=$?

if [ x"$1" == x"2" ] && [ x"$ccret" == x"0" ]; then
    sleep 2;
    sh make2.sh
fi


