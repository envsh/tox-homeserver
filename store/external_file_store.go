package store

import (
	"bytes"
	"gopp"
	"io"
	"log"
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

func (this *ExtFileStore) PutFileByName(fname string, f func()) (uname string, err error) {
	r, err := os.Open(fname)
	gopp.ErrPrint(err, fname)
	return this.PutFileByReader(r, filepath.Base(fname))
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

func (this *ExtFileStore) PutFileByData(data []byte, bname string) (uname string, err error) {
	r := NewClosingBuffer(data)
	return this.PutFileByReader(r, bname)
}

func (this *ExtFileStore) PutFileByReader(r io.ReadCloser, bname string) (uname string, err error) {
	for i := 0; i < 3; i++ {
		uname, err = this.PutFileByReaderImpl(r, bname)
		gopp.ErrPrint(err, bname, i)
		if err == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	return
}

func (this *ExtFileStore) PutFileByReaderImpl(r io.ReadCloser, bname string) (uname string, err error) {
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
	log.Println(u, resp.StatusCode, resp.String())

	data := resp.Bytes()
	retm := this.parseResult(data)
	log.Println(retm)
	uname = retm["url"]
	return
}

func (this *ExtFileStore) parseResult(data []byte) (retm map[string]string) {
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
