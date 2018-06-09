package main

import (
	"fmt"
	"gopp"
	"log"
	"os"
	"path/filepath"
	"strings"

	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"

	humanize "github.com/dustin/go-humanize"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
	"github.com/levigross/grequests"
	filetype "gopkg.in/h2non/filetype.v1"
)

func (this *MainWindow) initRoomFile() {
	this.initRoomFileUi()
	this.initRoomFileSignals()
	this.initRoomFileEvents()
}
func (this *MainWindow) initRoomFileUi() {

}
func (this *MainWindow) initRoomFileSignals() {
	qtrt.Connect(this.ToolButton_8, "clicked(bool)", func(bool) {
		dir := gopp.IfElseStr(gopp.IsAndroid(), os.Getenv("EXTERNAL_STORAGE"), os.Getenv("HOME"))
		fname := qtwidgets.QFileDialog_GetOpenFileName(this.QWidget_PTR(), "Select file", dir, "*.*", "*.*", 0)
		log.Println(fname)
		if fname != "" {
			this.sendFile(fname)
		}
	})
}
func (this *MainWindow) initRoomFileEvents() {

}

// TODO load last file info
func (this *MainWindow) initRoomFileStorage() {

}

func (this *MainWindow) sendFile(fname string) {
	item := uictx.msgwin.item
	if item == nil {
		log.Println("Dont know send to who.")
		return
	}

	fi, err := os.Stat(fname)
	gopp.ErrPrint(err)
	if fi.Size() > thscom.MaxAutoRecvFileSize {
		logtxt := fmt.Sprintf("File size too large, %d > %d", fi.Size(), thscom.MaxAutoRecvFileSize)
		log.Println(logtxt)
		ShowToast(logtxt, 1)
		return
	}

	ftyo, err := filetype.MatchFile(fname)
	gopp.ErrPrint(err, ftyo)

	picw := gopp.IfElseInt(strings.HasPrefix(ftyo.MIME.Value, "image/"), 200, 50)
	itext := fmt.Sprintf("<a href='%s'><img width='%d' src='%s' alt='%s'/></a><br/>%s (%s; %s)",
		fname,
		picw, fname, fname, filepath.Base(fname), ftyo.MIME.Value, humanize.Bytes(uint64(fi.Size())))
	userCode := thscli.NextUserCode(devInfo.Uuid)
	msgo := NewMessageForMe(itext)
	msgo.UserCode = userCode
	item.AddMessage(msgo, false)

	ro := &grequests.RequestOptions{}
	ro.Params = map[string]string{
		"friendPubkey": item.GetId(),
		"userCode":     gopp.ToStr(userCode),
		"fileName":     filepath.Base(fname),
	}
	fileo := grequests.FileUpload{}
	fileo.FileName = filepath.Base(fname)
	fileo.FileContents, _ = os.Open(fname)
	ro.Files = append(ro.Files, fileo)

	u := thscli.HttpFsUrlForUpload()
	resp, err := grequests.Put(u, ro)
	gopp.ErrPrint(err, u)
	log.Println(resp)
}
