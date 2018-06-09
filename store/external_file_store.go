package store

import (
	"bytes"
	"errors"
	"gopp"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/levigross/grequests"
)

type ExtFileStore struct {
}

var extfso *ExtFileStore
var extfsonce sync.Once

func GetExtFS() *ExtFileStore {
	extfsonce.Do(func() { extfso = NewExtFileStore() })
	return extfso
}

func NewExtFileStore() *ExtFileStore {
	this := &ExtFileStore{}

	return this
}

// TODO more usable external public file store
func (this *ExtFileStore) getPutUrl() string {
	hetxt := "fromtox-nobanme-github-com-envsh=1"
	u := "https://fars.ee/?" + hetxt
	return u
}

// should block
func (this *ExtFileStore) PutFileByName(fname string, f func()) (uname string, err error) {
	bname := filepath.Base(fname)

	var r *os.File
	for i := 0; i < 3; i++ {
		r, err = os.Open(fname)
		gopp.ErrPrint(err, fname)
		uname, err = this.PutFileByReaderFarsee(r, bname)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	for i := 0; i < 3; i++ {
		r, err = os.Open(fname)
		gopp.ErrPrint(err, fname)
		uname, err = this.PutFileByReaderVimcn(r, bname, false)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	for i := 0; i < 3; i++ {
		r, err = os.Open(fname)
		gopp.ErrPrint(err, fname)
		uname, err = this.PutFileByReaderVimcn(r, bname, true)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	return
}

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() (err error) {
	//we don't actually have to do anything here, since the buffer is
	// just some data in memory
	//and the error is initialized to no-error
	return
}

func NewClosingBuffer(data []byte) *ClosingBuffer {
	return &ClosingBuffer{bytes.NewBuffer(data)}
}

// should block
func (this *ExtFileStore) PutFileByData(data []byte, bname string) (uname string, err error) {
	for i := 0; i < 3; i++ {
		r := NewClosingBuffer(data)
		uname, err = this.PutFileByReaderFarsee(r, bname)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	for i := 0; i < 3; i++ {
		r := NewClosingBuffer(data)
		uname, err = this.PutFileByReaderVimcn(r, bname, false)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 0; i < 3; i++ {
		r := NewClosingBuffer(data)
		uname, err = this.PutFileByReaderVimcn(r, bname, true)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	return
}

func (this *ExtFileStore) PutFileByReaderFarsee(r io.ReadCloser, bname string) (uname string, err error) {
	ro := &grequests.RequestOptions{}
	ro.Params = map[string]string{
		// "friendPubkey": item.GetId(),
		// "userCode":     gopp.ToStr(userCode),
		// "fileName":     filepath.Base(fname),
	}
	fileo := grequests.FileUpload{}
	fileo.FileName = bname
	fileo.FileContents = r
	fileo.FieldName = "c"
	ro.Files = append(ro.Files, fileo)
	ro.RequestTimeout = 10 * time.Second
	ro.DialTimeout = 5 * time.Second
	ro.TLSHandshakeTimeout = 5 * time.Second

	u := this.getPutUrl()
	resp, err := grequests.Post(u, ro)
	gopp.ErrPrint(err, u)
	if false {
		log.Println(u, resp.StatusCode, resp.String(), resp.Header)
		req := resp.RawResponse.Request
		log.Println(req.ContentLength, req.Header)
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusMovedPermanently:
	default:
		err = errors.New("http status:" + gopp.ToStr(resp.StatusCode))
		return
	}

	data := resp.Bytes()
	retm := this.parseResultFarsee(data)
	log.Println(retm)
	uname = retm["url"]
	return
}
func (this *ExtFileStore) parseResultFarsee(data []byte) (retm map[string]string) {
	retm = make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		kv := strings.Split(line, ": ")
		retm[kv[0]] = kv[1]
	}
	return
}

func (this *ExtFileStore) PutFileByReaderVimcn(r io.ReadCloser, bname string, usepxy bool) (uname string, err error) {
	ro := &grequests.RequestOptions{}
	ro.Params = map[string]string{
		// "friendPubkey": item.GetId(),
		// "userCode":     gopp.ToStr(userCode),
		// "fileName":     filepath.Base(fname),
	}
	fileo := grequests.FileUpload{}
	fileo.FileName = bname
	fileo.FileContents = r
	fileo.FieldName = "name"
	ro.Files = append(ro.Files, fileo)
	ro.RequestTimeout = 20 * time.Second
	ro.DialTimeout = 10 * time.Second
	ro.TLSHandshakeTimeout = 10 * time.Second
	if usepxy {
		urlo, err := url.Parse("http://127.0.0.1:8117")
		gopp.ErrPrint(err)
		ro.Proxies = map[string]*url.URL{"http": urlo, "https": urlo}
	}

	u := this.getPutUrl()
	u = "https://img.vim-cn.com/"
	resp, err := grequests.Post(u, ro)
	gopp.ErrPrint(err, u)
	if false {
		log.Println(u, resp.StatusCode, resp.String(), resp.Header)
		req := resp.RawResponse.Request
		log.Println(req.ContentLength, req.Header)
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusMovedPermanently:
	default:
		err = errors.New("http status:" + gopp.ToStr(resp.StatusCode))
		return
	}

	data := resp.Bytes()
	log.Println(string(data))
	uname = strings.TrimSpace(string(data))
	return
}
