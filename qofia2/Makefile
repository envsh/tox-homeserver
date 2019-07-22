LDFLAGS0=-ldflags ""
LDFLAGS1=-ldflags "-w -s" # 对go-sqlite3和xorm很有效
# nostrip: 19M, strip: 13M, upx: 5M
all:
	gogo build $(LDFLAGS1) # --tags "libsqlite3"
	ls -lh qofia2
	#upx qofia2
	#go tool nm -size  -sort size qofia2|head

LDFLAGS2=-gccgoflags "-static-libgo"
# build way 2
# nostrip: 55M, strip: 33M
bw2:
	go build -v $(LDFLAGS2)
