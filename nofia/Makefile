
all: nui
	rm -f nofia
	@echo " " >> main1.go
	#  -ldflags "-w -s"
	go build -v -i -x -pkgdir ~/oss/pkg/linux_amd64 -ldflags "-v -linkmode external" # -linkshared

#		--nomain:on --nolinking:off --app:staticlib -o:libnikui.a
nui:
	nim --threads:on --listcmd --header \
		-p:/opt/nim/lib -p:$(HOME)/.nimble/mulib \
		c nikui.nim
	cp $(HOME)/.cache/nim/nikui_d/nikui.h .

c2n:
	c2nim --typeprefixes render_x11_native.h

clean:
	rm -f nofia libnikui.a

