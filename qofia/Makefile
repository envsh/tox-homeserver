
all:


gen:
	go build -o /tmp/go-uic  go-uic/main.go
	uic untitled.ui > ui_untitled.h
	/tmp/go-uic ui_untitled.h
	uic untitled2.ui > ui_untitled2.h
	/tmp/go-uic ui_untitled2.h
	uic contact_item_view.ui > ui_contact_item_view.h
	/tmp/go-uic ui_contact_item_view.h
	uic message_item_view.ui > ui_message_item_view.h
	/tmp/go-uic ui_message_item_view.h
	qtrcc

bud:
	go build -v

run:
	./qofia

clean:
	rm -vf ui_*.h ui_*_c.h ui_*.go ui_*.cpp
	rm -vf rcc.cpp rcc_cgo_*.go
