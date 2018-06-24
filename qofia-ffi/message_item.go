package main

import (
	"fmt"
	"gopp"
	"image"
	"os"
	"path/filepath"
	"strings"

	thscli "tox-homeserver/client"
	thscom "tox-homeserver/common"

	humanize "github.com/dustin/go-humanize"
)

// Ui_MessageItemView's wrapper
type MessageItem struct {
	*Ui_MessageItemView

	Sent     bool
	UserCode int64
}

func NewMessageItem() *MessageItem {
	this := &MessageItem{}
	this.Ui_MessageItemView = NewUi_MessageItemView2()

	return this
}

/////
// TODO move Message struct here???
func (this *Message) IsFile() bool {
	return strings.HasPrefix(this.Msg, "txc;")
}
func (this *Message) IsImage() bool {
	return strings.HasPrefix(this.Msg, "txc;image/")
}

func (this *Message) GetFileInfoLine() *thscom.FileInfoLine {
	return thscom.ParseFileInfoLine(this.Msg)
}

// QLabel not support <img src='http://' />
// locdir can not be prefix file://
func Msg2FileText(fil *thscom.FileInfoLine, locdir string) string {
	absdir, _ := filepath.Abs(locdir)

	picw := previewWidth(fil.Mime, fmt.Sprintf("%s/%s", absdir, fil.Md5str))
	fsrc := fmt.Sprintf("%s", thscli.HttpFsUrlFor(fil.Md5str)) + "." + fil.Ext
	itext := fmt.Sprintf("<a href='%s'><img width='%d' src='file://%s/%s' alt='%s'/></a><br/>%s (%s; %s)",
		fsrc,
		picw, absdir, fil.Md5str, fsrc, fil.OrigName, fil.Mime, humanize.Bytes(uint64(fil.Length)))
	return itext
}

func GetUrlByHash(md5str string) string {
	return thscli.HttpFsUrlFor(md5str)
}

func GetLocalPathByHash(md5str string, locdir string) string {
	absdir, _ := filepath.Abs(locdir)
	return fmt.Sprintf("file://%s/%s", absdir, md5str)
}

func previewWidth(mime_ string, fname string) int {
	// picw := gopp.IfElseInt(strings.HasPrefix(fil.Mime, "image/"), 200, 50)
	maxval := 150
	if strings.HasPrefix(mime_, "image/") {
		fp, err := os.Open(fname)
		gopp.ErrPrint(err, fname)
		defer fp.Close()
		imgcfg, _, err := image.DecodeConfig(fp)
		gopp.ErrPrint(err, fname)
		if imgcfg.Width <= maxval && imgcfg.Height <= maxval {
			return imgcfg.Width
		}
		if imgcfg.Height > imgcfg.Width {
			// w/h = x/200
			return imgcfg.Width * maxval / imgcfg.Height
		} else {
			return maxval
		}
	} else {
		return 50
	}
}
