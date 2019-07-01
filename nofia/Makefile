
all: nui

# --nomain:on --nolinking:off --app:staticlib -o:libnikui.a
# --opt:none fastest compile speed
# -d:useRealtimeGC
#		--cpu:i386 --passc:-m32 --passl:-m32
nui:
	nim --threads:on --listcmd --header -d:nimCallDepthLimit=128 \
		-p:/opt/nim/lib -p:$(HOME)/.nimble/mulib \
		--opt:none -o:nofia c nikui.nim
	cp $(HOME)/.cache/nim/nikui_d/nikui.h .

c2n:
	c2nim --typeprefixes render_x11_native.h

clean:
	rm -f nofia libnikui.a

futl:
	/opt/go/bin/go build -v -buildmode=c-archive -o libgofiautil.a ../fiautil

# some tricks
# gdb: handle SIGXCPU SIG33 SIG35 SIGPWR nostop noprint
