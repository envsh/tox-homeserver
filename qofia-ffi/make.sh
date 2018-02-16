source ~/triline/shell/android-ndk-env.sh
source ~/triline/shell/android-go-env.sh

env | grep CGO_
set -x

# go install -p 1 -v  -pkgdir ~/oss/pkg/android_arm tox-homeserver/gofia
# go build -p 1 -v  -pkgdir ~/oss/pkg/android_arm .

go install -v -pkgdir ~/oss/pkg/android_arm github.com/gonuts/ffi
go install -v -pkgdir ~/oss/pkg/android_arm github.com/kitech/qt.go/qtqt
go install -v -pkgdir ~/oss/pkg/android_arm github.com/kitech/qt.go/qtrt
go install -v -pkgdir ~/oss/pkg/android_arm github.com/mattn/go-sqlite3

time go build -p 1 -v  -pkgdir ~/oss/pkg/android_arm -buildmode=c-shared -o libmain.so .
chmod +x libmain.so

mv andwrapmain.c.nogo andwrapmain.c
$CC andwrapmain.c -shared   -o libgo.so -lmain -L. -Wl,-soname,libgo.so
mv andwrapmain.c andwrapmain.c.nogo

