
all: nui

#		--nomain:on --nolinking:off --app:staticlib -o:libnikui.a
nui:
	nim --threads:on --listcmd --header -d:nimCallDepthLimit=128 \
		-p:/opt/nim/lib -p:$(HOME)/.nimble/mulib \
		-o:nofia c nikui.nim
	cp $(HOME)/.cache/nim/nikui_d/nikui.h .

c2n:
	c2nim --typeprefixes render_x11_native.h

clean:
	rm -f nofia libnikui.a

gogo:
	# handle SIGXCPU SIG33 SIG35 SIGPWR nostop noprint
	# ./configure --enable-redirect-malloc --enable-static=yes --enable-shared=no --enable-cplusplus --disable-thread-local-alloc
	g++ -g -O0 -c gogo1.cpp -o gogo1.o -I/home/me/oss/src/cxrt/libgo/libgo
	# --gc:none
	LDFLAGS="-lgc123" nim --threads:on --listcmd --header -d:nimCallDepthLimit=128 \
		-d:debug -d:nimCoroutines --gc:refc -p:/opt/nim/lib -p:$(HOME)/.nimble/mulib \
		-o:gogo- c gogo.nim


