// Code generated by go-bindata.
// sources:
// __ant__tpl__.go
// DO NOT EDIT!

package test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

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

var ___ant__tpl__Go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xc1\x6e\xdb\x30\x0c\x3d\x4b\x5f\x41\xf4\x94\xe4\x30\xa5\x5b\xd1\x43\x30\x0c\xf0\xda\x01\x29\x90\x76\xc1\x92\x9e\x55\xc2\x66\x6c\x6d\xb6\xe4\x49\xcc\xb0\xa1\xc8\xbf\x0f\x92\x9c\x39\x2e\x96\x83\xe1\xf0\xf1\x3d\x3e\x3e\x5a\x29\xe8\xb1\xfc\x81\x35\x81\xd6\xc5\xd3\x5e\xeb\xfd\x76\xa3\x35\x98\x00\xdc\x10\xf4\xde\x7d\xa7\x92\x81\xa9\xeb\x5b\x64\x92\xff\x6b\x96\x52\xa9\x58\xd8\x3e\x68\xbd\x7d\xde\x44\xba\xa7\xda\x04\x26\x0f\xf1\x3f\x78\x77\x64\xf2\xab\xd8\x06\xaa\x71\x1d\xe5\xb7\x0e\xb9\x51\x95\xf9\x65\x2a\x92\xfc\xa7\xa7\x37\x1a\xc6\x32\xf9\x03\x96\x04\xaf\x52\xac\x5d\x47\xb3\x45\x60\x7f\x2c\xf9\xf5\x34\x87\x45\x2c\x7c\xa3\x70\x6c\x59\x8a\x47\xe4\x46\x9e\xa6\x36\x76\xeb\xa9\x8d\xdd\x7a\x6a\x23\x30\xf2\x9b\xa9\x89\x32\x99\xba\x63\xe4\xd9\x22\x3e\x0b\x5f\xcf\x87\x11\x8f\x5f\xef\xbf\x6c\xa0\xf4\x84\x4c\xd0\xb9\x8a\xda\xb3\x4e\x42\xb4\x86\x6c\x33\x0a\xdc\xa7\xed\x0a\x5f\x4b\xf1\x1c\xc8\x9f\x15\x90\x1b\x28\x9d\x65\xef\xda\x96\x7c\xa6\xa7\xe2\x64\xba\x52\x90\xf9\xd0\xa0\xad\x62\xe3\xa0\x37\x5b\xfc\xd3\x9d\xc3\xf0\x3e\x64\x91\x07\x8c\xe1\x40\x8c\x1b\x7c\x06\xd3\x9c\x0b\x6c\xf4\x79\xe7\x2c\x93\x4d\x15\x63\x6b\x50\x0a\x98\x7e\x27\xb5\xc4\x99\x5d\x98\x29\x7c\x0d\xf9\x68\x80\xbd\x01\x8c\xbb\x8d\xc8\x28\x19\x19\xb9\xcf\x56\x52\x88\x02\x0e\xad\x43\xbe\xbd\x19\x91\xe0\xbc\x14\xe2\xf3\x19\x80\x97\x1e\x3d\x76\xab\xab\x8f\x1e\x6d\x4d\x2b\x58\xbe\x5b\x5e\xaf\xae\x97\xf1\xf7\xe9\xea\x45\x8a\xd3\x85\x8b\x61\x83\x0b\x23\xc3\x8e\x62\x82\x4f\xed\xfc\x3c\x3a\x36\x64\x59\x0a\x71\x37\xda\x39\xc9\x79\x0a\x6d\x38\x33\xc4\x2f\xe3\x9c\x78\x5a\x2f\x45\x30\xa2\x67\xc5\x7d\x88\xe7\xba\xbd\x49\x69\x99\x8e\x02\x63\xd7\x87\xe1\x02\xf1\xda\x70\x8c\x0f\x63\x0f\x2e\x4b\xa4\xda\xc8\x7f\xa8\x00\xb2\x82\x14\x4f\xd8\xd1\x10\xbe\x14\x45\x4d\x09\xf8\xf0\x5e\x9e\xe4\xdf\x00\x00\x00\xff\xff\xa6\x62\xed\x1d\xa1\x03\x00\x00")

func __ant__tpl__GoBytes() ([]byte, error) {
	return bindataRead(
		___ant__tpl__Go,
		"__ant__tpl__.go",
	)
}

func __ant__tpl__Go() (*asset, error) {
	bytes, err := __ant__tpl__GoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "__ant__tpl__.go", size: 929, mode: os.FileMode(420), modTime: time.Unix(1529039126, 0)}
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
	"__ant__tpl__.go": __ant__tpl__Go,
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
	"__ant__tpl__.go": &bintree{__ant__tpl__Go, map[string]*bintree{}},
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
