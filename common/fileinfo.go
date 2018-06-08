package common

import (
	"fmt"
	"gopp"
	"log"
	"strings"
)

type FileInfoLine struct {
	mime     string
	length   int64
	md5str   string
	ext      string
	origName string
}

func ParseFileInfoLine(s string) *FileInfoLine {
	parts := strings.Split(s, ";")
	if len(parts) != 6 {
		log.Println("fmt error:", s)
		return nil
	}
	if parts[0] != "file" {
		log.Println("fmt error:", s)
		return nil
	}
	return NewFileInfoLine(gopp.MustInt64(parts[2]), parts[1], parts[3], parts[4], parts[5])
}

// file;mime;len;md5;ext;origName
func (this *FileInfoLine) String() string {
	return fmt.Sprintf("file;%s;%d;%s;%s;%s", this.mime, this.length, this.md5str, this.ext, this.origName)
}

func NewFileInfoLine(length int64, mime, md5str, ext, origName string) *FileInfoLine {
	fis := &FileInfoLine{}
	fis.mime = mime
	fis.length = length
	fis.md5str = md5str
	fis.ext = ext
	fis.origName = origName
	return fis
}
