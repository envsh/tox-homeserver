// Code generated by go-bindata.
// sources:
// data/server.crt
// DO NOT EDIT!

package transport

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataServerCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIDETCCAfkCFB4LcPT5dqWF/M1ukgyYxgCA5fxtMA0GCSqGSIb3DQEBCwUAMEUx
CzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl
cm5ldCBXaWRnaXRzIFB0eSBMdGQwHhcNMTgxMDA3MTIzNjAxWhcNMjgxMDA0MTIz
NjAxWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UE
CgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAueTZd7X7pWSePOrr0QNYVeObXMqLr+TIK66f4rh1fOHXKDtX
DFffjgtHTkyz7fcjiNPVu56x2uuY4T00eWdgtdYTdlF9EpUbUD/ZyAvWj/agMHmR
eT5fCjaPLny1NMtgTXYZWWvXr3U2BJkOfPgW3yOdpSCRwASmSkBzxWzQrHeOa6MF
dKPtwq2yLjhmH1xGKMpGjySLqJvGfZctx6dxNZ8Kq0GbhveR8H0V74D3yXYVg6VX
IwYeOLK8aG63wiCFM17NzcZdbHSEJ8csig7dKBiCUPcSxCVie/9R6HmKKkuBYrpQ
9XbikvftxohKRPxNYRtyZOKX4G/5AhQem0PCGQIDAQABMA0GCSqGSIb3DQEBCwUA
A4IBAQAKGOitRUWTmgXX0l3e9TeJge0rqxxqUoVRnVjPfLk+DRcMnuWD17xzHPew
y+Y1XxjXp5UfSREvI5I4d1JlwAZWJCyrX14jaRoCCmdPR39t9WdLWOQSXPDX3ai3
ca5MUZL5vosN0yWL7p3Kx86a+B5SYwV4TM8vGrnqfx0KUMY/wRIxKDLCsbdUuLFz
6dCAoVXV/7xS8vZrxqvS56gRJdOH27bhS7zG6XTNe7zVCQaix3qmfMK4j3OrUfJz
UYDa5F6m45/rrhz5fokFVBVizBZYrqREPXwwmF8YOCqdwtL2/ocfJMjPuXfTq2Y1
F8PNPkeT6PsXwZbE4ibfZoU6F5rT
-----END CERTIFICATE-----
`)

func dataServerCrtBytes() ([]byte, error) {
	return _dataServerCrt, nil
}

func dataServerCrt() (*asset, error) {
	bytes, err := dataServerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/server.crt", size: 1123, mode: os.FileMode(420), modTime: time.Unix(1538925010, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/server.crt": dataServerCrt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"server.crt": &bintree{dataServerCrt, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

