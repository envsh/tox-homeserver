source ~/triline/shell/android-ndk-env.sh
source ~/triline/shell/android-go-env.sh

set -x
# go install -p 1 -v  -pkgdir ~/oss/pkg/android_arm tox-homeserver/gofia
# go build -p 1 -v  -pkgdir ~/oss/pkg/android_arm .
go build -p 1 -v  -pkgdir ~/oss/pkg/android_arm -buildmode=c-shared -o libmain.so .
chmod +x libmain.so

mv andwrapmain.c.nogo andwrapmain.c
$CC andwrapmain.c -shared   -o libgo.so -lmain -L.
mv andwrapmain.c andwrapmain.c.nogo

