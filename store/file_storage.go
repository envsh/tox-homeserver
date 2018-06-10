package store

import (
	"gopp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	locfsonce.Do(func() { locfso = NewFileStorage() })
	return locfso
}

func NewFileStorage() *FileStorage {
	this := &FileStorage{}
	this.osfs = afero.NewOsFs()
	this.httpfs = afero.NewHttpFs(this.osfs)

	reldir := "./toxhsfiles"
	absdir, err := filepath.Abs(reldir)
	gopp.ErrPrint(err)
	this.dir = absdir

	dvo := diskv.Options{}
	dvo.BasePath = this.dir
	this.dv = diskv.New(dvo)

	return this
}

func (this *FileStorage) ReadFile(md5str string) (data []byte, err error) {
	fname := this.dir + "/" + md5str
	return afero.ReadFile(this.osfs, fname)
}

func (this *FileStorage) SaveFile(data []byte, origName string) (string, error) {
	md5str := gopp.Md5AsStr(data)
	fname := this.dir + "/" + md5str
	this._SaveOrigName(md5str, origName)
	return md5str, afero.WriteFile(this.osfs, fname, data, 0644)
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

// nil for default
func (this *FileStorage) SetupHttpServer(srvmux *http.ServeMux) {
	if srvmux == nil {
		srvmux = http.DefaultServeMux
	}
	srvmux.HandleFunc("/toxhsfs/", func(w http.ResponseWriter, r *http.Request) {
		// log.Println("ohhhh", r.URL.String())
		md5str := r.URL.String()[9:]
		md5str = gopp.SubStr(md5str, 32)
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
