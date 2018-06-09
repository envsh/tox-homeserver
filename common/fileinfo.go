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
	mime     string
	length   int64
	md5str   string
	ext      string
	origName string
	urlval   string
}

func ParseFileInfoLine(s string) *FileInfoLine {
	parts := strings.Split(s, ";")
	if len(parts) != 6 && len(parts) != 7 {
		log.Println("fmt error:", s)
		return nil
	}
	if parts[0] != "file" && !strings.HasPrefix(parts[0], "http://") {
		log.Println("fmt error:", s)
		return nil
	}
	if len(parts) == 7 {
		return NewFileInfoLineUrl(gopp.MustInt64(parts[3]), parts[2], parts[4], parts[5], parts[6], parts[0])
	}
	return NewFileInfoLine(gopp.MustInt64(parts[2]), parts[1], parts[3], parts[4], parts[5])
}

// file;mime;len;md5;ext;origName
func (this *FileInfoLine) String() string {
	if this.urlval != "" {
		return fmt.Sprintf("%s?file;%s;%d;%s;%s;%s", this.urlval,
			this.mime, this.length, this.md5str, this.ext, this.origName)
	}
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

func NewFileInfoLineUrl(length int64, mime, md5str, ext, origName, urlval string) *FileInfoLine {
	fis := NewFileInfoLine(length, mime, md5str, ext, origName)
	fis.urlval = urlval
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
