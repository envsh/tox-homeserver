package common

import (
	"encoding/hex"
	"fmt"
	"gopp"
	"log"
	"strings"

	"github.com/shengdoushi/base58"
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
