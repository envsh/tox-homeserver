package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gopp"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shengdoushi/base58"
	filetype "gopkg.in/h2non/filetype.v1"
)

type FileInfoLine struct {
	Mime     string
	Length   int64
	Md5str   string
	Ext      string
	OrigName string
	Urlval   string
}

func ParseFileInfoLine(s string) *FileInfoLine {
	parts := strings.Split(s, ";")
	if len(parts) != 6 && len(parts) != 7 {
		log.Println("fmt error:", s)
		return nil
	}
	if parts[0] != "txc" && !strings.HasPrefix(parts[0], "http://") {
		log.Println("fmt error:", s)
		return nil
	}
	if len(parts) == 7 {
		return NewFileInfoLineUrl(gopp.MustInt64(parts[3]), parts[2], parts[4], parts[5], parts[6], parts[0])
	}
	return NewFileInfoLine(gopp.MustInt64(parts[2]), parts[1], parts[3], parts[4], parts[5])
}

// txc;mime;len;md5;ext;origName
func (this *FileInfoLine) String() string {
	if this.Urlval != "" {
		return fmt.Sprintf("%s?txc;%s;%d;%s;%s;%s", this.Urlval,
			this.Mime, this.Length, this.Md5str, this.Ext, this.OrigName)
	}
	return fmt.Sprintf("txc;%s;%d;%s;%s;%s", this.Mime, this.Length, this.Md5str, this.Ext, this.OrigName)
}

func (this *FileInfoLine) ToType() string {
	if strings.HasPrefix(this.Mime, "image/") {
		return MSGTYPE_IMAGE
	} else if strings.HasPrefix(this.Mime, "audio/") {
		return MSGTYPE_AUDIO
	} else if strings.HasPrefix(this.Mime, "video/") {
		return MSGTYPE_VIDEO
	}
	return MSGTYPE_FILE
}

func NewFileInfoLine(length int64, mime, md5str, ext, origName string) *FileInfoLine {
	fis := &FileInfoLine{}
	fis.Mime = mime
	fis.Length = length
	fis.Md5str = md5str
	fis.Ext = ext
	fis.OrigName = origName
	return fis
}

func NewFileInfoLineUrl(length int64, mime, md5str, ext, origName, urlval string) *FileInfoLine {
	fis := NewFileInfoLine(length, mime, md5str, ext, origName)
	fis.Urlval = urlval
	return fis
}

func NewFileInfoLineForFile(fname string) *FileInfoLine {
	ftyo, err := filetype.MatchFile(fname)
	gopp.ErrPrint(err, fname)
	fi, err := os.Stat(fname)
	gopp.ErrPrint(err, fname)
	this := &FileInfoLine{}
	this.Mime = ftyo.MIME.Value
	this.Ext = ftyo.Extension
	this.Length = fi.Size()
	this.OrigName = filepath.Base(fname)
	this.Md5str = Md5file(fname)
	return this
}

func Md5file(fname string) string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

// Bitcoin: 123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
// base58.BitcoinAlphabet
// IPFS: 123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
// base58.IPFSAlphabet
// Ripple: rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz
// base58.RippleAlphabet
// Flickr: 123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
// base58.FlickrAlphabet
func Base58Encode(data []byte) string {
	return base58.Encode(data, base58.FlickrAlphabet)
}
func Base58EncodeFromHex(h string) string {
	md5bin, err := hex.DecodeString(h)
	gopp.ErrPrint(err)
	return Base58Encode(md5bin)
}
func Base58Decode(s string) ([]byte, error) {
	return base58.Decode(s, base58.FlickrAlphabet)
}
