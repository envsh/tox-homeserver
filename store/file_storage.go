package store

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	thscom "tox-homeserver/common"

	"github.com/levigross/grequests"
	"github.com/peterbourgon/diskv"
	"github.com/spf13/afero"
	filetype "gopkg.in/h2non/filetype.v1"
)

type FileStorage struct {
	osfs     afero.Fs
	osfsol   *afero.MemMapFs // TODO the cache layer of osfs
	httpfs   *afero.HttpFs
	httpfsol *afero.MemMapFs // TODO the cache layer of httpfs
	dv       *diskv.Diskv    // store meta info

	dir              string // rootdir
	onFileUploadedFn func(md5str string, frndpk string, userCode string)
}

var locfso *FileStorage
var locfsonce sync.Once

func GetFS() *FileStorage {
	locfsonce.Do(func() { locfso = NewFileStorage(false) })
	return locfso
}

func GetFSC() *FileStorage {
	locfsonce.Do(func() { locfso = NewFileStorage(true) })
	return locfso
}

func NewFileStorage(client bool) *FileStorage {
	this := &FileStorage{}
	this.osfs = afero.NewOsFs()
	this.httpfs = afero.NewHttpFs(this.osfs)

	this.setBaseDir(client)

	dvo := diskv.Options{}
	dvo.BasePath = this.dir
	this.dv = diskv.New(dvo)

	return this
}

func (this *FileStorage) setBaseDir(client bool) {
	reldir := "./toxhsfiles"
	if client {
		reldir = gopp.IfElseStr(gopp.IsAndroid(), os.Getenv("EXTERNAL_STORAGE"), ".") + "/txcfiles"
		reldir = gopp.IfElseStr(gopp.IsAndroid(), "/sdcard", ".") + "/txcfiles"
		os.MkdirAll(reldir, 0744)
		os.Chmod(reldir, 0744)
	}

	absdir, err := filepath.Abs(reldir)
	gopp.ErrPrint(err)
	this.dir = absdir
}

func (this *FileStorage) ReadFile(md5str string) (data []byte, err error) {
	if !this.checkFileEscapeBaseDir(md5str) {
		return nil, fmt.Errorf("Invalid filename: %s", md5str)
	}
	fname := this.dir + "/" + md5str
	return afero.ReadFile(this.osfs, fname)
}

func (this *FileStorage) ReadFileWithName(name string) (data []byte, err error) {
	fname := this.dir + "/" + name
	return afero.ReadFile(this.osfs, fname)
}

func (this *FileStorage) SaveFile(data []byte, origName string) (string, error) {
	md5str := gopp.Md5AsStr(data)
	if !this.checkFileEscapeBaseDir(md5str) {
		return "", fmt.Errorf("Invalid filename: %s", md5str)
	}
	fname := this.dir + "/" + md5str
	this._SaveOrigName(md5str, origName)
	return md5str, afero.WriteFile(this.osfs, fname, data, 0644)
}

// for avatar
func (this *FileStorage) SaveFileWithName(name string, data []byte) error {
	md5str := name
	fname := this.dir + "/" + md5str
	this._SaveOrigName(md5str, name)
	return afero.WriteFile(this.osfs, fname, data, 0644)
}

func (this *FileStorage) _SaveOrigName(md5str, origName string) error {
	return this.dv.Write(md5str+".on", []byte(origName))
}

func (this *FileStorage) GetOrigName(md5str string) (string, error) {
	data, err := this.dv.Read(md5str + ".on")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (this *FileStorage) TempFile() afero.File {
	fh, err := afero.TempFile(this.osfs, this.dir, "txfstmp.")
	gopp.ErrPrint(err)
	return fh
}

func (this *FileStorage) TempFileName() string {
	name := fmt.Sprintf("%s/txfstmp.%s", this.dir, gopp.RandomStringAlphaMixed(7))
	return name
}

// nil for default
func (this *FileStorage) SetupHttpServer(srvmux *http.ServeMux) {
	if srvmux == nil {
		srvmux = http.DefaultServeMux
	}
	srvmux.HandleFunc("/toxhsfs/", func(w http.ResponseWriter, r *http.Request) {
		// log.Println("ohhhh", r.URL.String())
		md5str := r.URL.String()[9:]
		pos := strings.Index(md5str, ".")
		if pos == -1 {
			// direct use, maybe avatar
		} else {
			md5str = gopp.SubStr(md5str, 32)
		}

		data, err := this.ReadFile(md5str)
		gopp.ErrPrint(err, r.URL.String(), md5str)
		w.Write(data)
	})

	// because this.dir has no toxhsfs subdirectory, so it not work
	// srvmux.Handle("/toxhsfs/", http.FileServer(this.httpfs.Dir(this.dir)))

	srvmux.HandleFunc("/toxhsfs/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			log.Println("Unsported request:", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.ContentLength > thscom.MaxAutoRecvFileSize {
			log.Println("File too large:", r.ContentLength)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		data, err := ioutil.ReadAll(r.Body)
		gopp.ErrPrint(err)
		if err == nil {
			fname := r.URL.Query().Get("fileName")
			frndpk := r.URL.Query().Get("friendPubkey")
			userCode := r.URL.Query().Get("userCode")
			md5str, err := this.SaveFile(data, fname)
			gopp.ErrPrint(err)
			time.AfterFunc(1*time.Millisecond, func() {
				if this.onFileUploadedFn != nil {
					this.onFileUploadedFn(md5str, frndpk, userCode)
				}
			})
		}
	})
}

func (this *FileStorage) OnFileUploaded(f func(md5str string, frndpk string, userCode string)) {
	oldfn := this.onFileUploadedFn
	this.onFileUploadedFn = func(md5str string, frndpk string, userCode string) {
		if f != nil {
			f(md5str, frndpk, userCode)
		}

		if oldfn != nil {
			oldfn(md5str, frndpk, userCode)
		}
	}
}

func setpxy4ro(ro *grequests.RequestOptions) {
	urlo, err := url.Parse("http://127.0.0.1:8117")
	gopp.ErrPrint(err)
	ro.Proxies = map[string]*url.URL{"http": urlo, "https": urlo}
}

func (this *FileStorage) DownloadToFile(urltxt string) (string, error) {
	md5str, err := this.DownloadToFileImpl(urltxt, false)
	if err != nil {
		md5str, err = this.DownloadToFileImpl(urltxt, true)
	}
	return md5str, err
}

func (this *FileStorage) DownloadToFileImpl(urltxt string, usepxy bool) (string, error) {
	ro := &grequests.RequestOptions{}
	ro.RequestTimeout = 40 * time.Second
	ro.RedirectLimit = 40
	if usepxy {
		setpxy4ro(ro)
	}
	resp, err := grequests.Get(urltxt, ro)
	gopp.ErrPrint(err, urltxt)
	gopp.ErrPrint(resp.Error, urltxt)
	if err != nil {
		return "", err
	}

	fname := this.dir + "/" + filepath.Base(urltxt) + ".part"
	err = resp.DownloadToFile(fname)
	gopp.ErrPrint(err, fname)

	if err == nil {
		md5str := thscom.Md5file(fname)
		err = os.Rename(fname, this.dir+"/"+md5str)
		gopp.ErrPrint(err, fname)
		this._SaveOrigName(md5str, filepath.Base(urltxt))
		return md5str, err
	}
	return "", err
}

func (this *FileStorage) GetFilePath(md5str string) string {
	return this.dir + "/" + md5str
}

// file://
func (this *FileStorage) GetUrlDir() string { return fmt.Sprintf("file://%s", this.dir) }

func (this *FileStorage) RemoveFile(md5str string) error {
	if !this.checkFileEscapeBaseDir(md5str) {
		return fmt.Errorf("Invalid filename: %s", md5str)
	}
	path := this.GetFilePath(md5str)
	return this.osfs.Remove(path)
}

func (this *FileStorage) checkFileEscapeBaseDir(md5str string) bool {
	path := this.GetFilePath(md5str)
	abspath, err := filepath.Abs(path)
	gopp.ErrPrint(err)
	if strings.HasPrefix(abspath, this.dir) {
		return true
	}
	return false
}

func (this *FileStorage) NewFileInfoLine4Md5(md5str string) *thscom.FileInfoLine {
	fname := this.GetFilePath(md5str)
	ftyo, err := filetype.MatchFile(fname)
	gopp.ErrPrint(err, fname)
	fi, err := os.Stat(fname)
	gopp.ErrPrint(err, fname)
	this_ := &thscom.FileInfoLine{}
	this_.Mime = ftyo.MIME.Value
	this_.Ext = ftyo.Extension
	this_.Length = fi.Size()
	this_.OrigName = gopp.Retn(this.GetOrigName(md5str))[0].(string)
	this_.Md5str = md5str
	return this_
}
