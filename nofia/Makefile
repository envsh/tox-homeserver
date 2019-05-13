
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

