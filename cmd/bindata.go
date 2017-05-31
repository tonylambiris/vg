// Code generated by go-bindata.
// sources:
// data/fish
// data/sh
// DO NOT EDIT!

package cmd

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

var _dataFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5f\x4f\xdb\x30\x10\x7f\xef\xa7\x38\x4a\xa5\xb6\x48\x6e\x04\x8f\x20\x26\x81\x8a\xa0\x12\xd0\x8a\xb1\xf1\x80\xba\xc8\x4d\x9c\xc4\x5a\x62\x77\xb6\x13\xa6\xa9\x1f\x7e\xb2\xd3\xc4\x71\xea\xb2\x97\x59\x42\x11\x77\xbf\xfb\xe7\xbb\xfb\xb9\xa7\x27\xc1\x86\xb2\x20\xa1\x32\x0b\x06\x83\xa4\x64\x91\xa2\x9c\x41\x95\x0e\x00\x00\x68\x02\x8c\x2b\x88\x78\xc9\x14\x8c\xb0\x48\x2b\xf8\x02\x41\x4c\xaa\x80\x95\x79\x6e\x20\xfa\x44\xbc\x28\x30\x8b\x1b\x2b\x7d\x04\x51\xa5\x60\xe6\x5f\xc2\x62\xf3\x3d\x05\x12\x65\xbc\x76\xf3\x7e\xbe\x36\x32\xf9\x41\x55\x94\xb9\x32\xe3\x11\x4b\x02\x38\x52\xb4\xc2\x8a\xb4\x52\x7d\xc2\x2a\x0d\x1b\x45\x63\x37\x9b\xa1\xbe\x6d\x4c\x8e\x5a\x7b\x54\xc6\x64\x7c\x36\x76\xb0\xb6\xaa\x3a\x4e\x5b\x8c\xfe\xb3\x57\xd5\x4d\xa8\xb9\xb4\x77\x40\x0c\x86\xa3\xef\x8b\x97\xd7\x6f\x37\x8f\xf7\xcb\x21\xd8\xf4\x3c\x29\x18\x87\xad\xe9\xa4\x73\xdd\x53\x40\xb9\x82\x8b\x8e\xb9\x24\x0a\x4c\xd5\x17\x6b\x98\x6c\xb0\x24\x0c\x17\x04\x46\xab\xb7\xf9\xd4\xf5\xa5\x81\x28\xfd\x0d\x6d\x12\xfb\xdb\xba\x58\xfb\xd5\xe1\xea\xe6\xf5\x01\x46\x0f\xcb\xa7\xbb\x60\x56\x51\xa1\x4a\x9c\xa7\x3c\xb0\x45\xb8\x6e\x43\x6b\xb8\x7c\x9c\xdf\x2f\x6b\xeb\xfa\xfb\x29\xf0\x76\xf1\xac\x71\xb7\x8b\xe7\x4f\x60\xb5\x37\xe3\xcb\x41\x35\x71\xdc\xac\x2f\x7d\x71\xf7\x91\x5c\xa4\x1e\xf6\xda\x63\xf1\x33\xa6\x02\xd0\xf6\x00\x21\x45\x14\xf4\x10\x75\xb6\x8e\xf7\xa6\x5c\x13\xc3\xe6\x59\xf7\xfe\x4f\xb7\xf7\xe1\x7c\xf1\xf5\xe6\xf6\xf1\x2e\x5c\xbd\x2c\x9f\x56\xaf\xdd\x51\x68\x66\x48\x02\x8a\x40\xaf\x60\xb8\x15\xbc\xd8\x2a\x08\x79\x1e\x87\x1d\xc1\x81\x05\xf8\x94\xa6\xfd\x7a\xc5\xf4\xf0\x4d\x6c\x06\x53\x18\xba\x4b\x70\xcc\x7b\xb3\xa9\xfe\x21\xef\x0d\xad\xa7\xd4\x6e\x6d\x26\x91\x3b\x21\xb8\xb8\x84\x67\x0e\x29\x07\xc2\x2a\x2a\x38\x2b\x08\x53\x40\x65\xbd\xdd\x04\xb0\x02\x95\x11\x28\xb8\x96\xb7\xd6\x55\x0a\x19\xc9\xb7\x3d\x3a\x81\x73\x87\x50\x4c\x2f\xda\x46\x7b\xe6\xcc\x45\xd5\x2d\xf3\xcd\x6d\x17\xe7\x43\xfd\x87\xfe\x12\x6f\xc7\x9c\x01\xe8\xb7\xe5\x5f\x16\xc4\xdf\x48\x97\x00\x88\x5d\x70\xaf\x30\x74\xeb\x27\x9e\x3d\x3c\xae\xed\x5d\xf2\xa1\xda\x98\xf7\x26\xc9\xf0\x65\xa9\xb8\x65\x71\x84\x38\x43\x15\x16\x14\x6f\x72\x02\xab\xb7\x79\x73\xd5\x52\x61\x55\x4a\x40\x88\x4a\xb4\x67\x63\x24\xcb\x8d\x54\x54\x95\xda\xd9\xb1\xd7\x66\xff\xdc\x24\x94\xc5\x80\x19\xe8\x68\x68\x1f\xad\xde\x9d\x9c\xd8\xac\x73\xb0\xaa\x50\x70\xae\x0c\x99\x1a\xfd\x47\x46\x73\x02\x8a\x48\xcd\xc6\x3d\xd0\xc9\x35\x0c\xed\x5a\xd1\xa4\x86\xa1\x04\x86\x7d\x68\x87\x4d\xdd\x3d\xd4\xe1\x19\xf9\x08\x5b\x75\x68\xb8\x7c\x12\x61\xf5\xb9\x97\xa9\xe3\x86\x26\xb0\x21\x29\x65\xe6\xa9\x36\x25\xfd\xb2\xfd\xbd\x02\x2e\xf6\x15\x78\x42\x9d\x5c\x77\xe8\xef\xaa\x5d\xac\xee\xa9\x52\xb0\xcf\xed\xa1\x0b\x97\x7c\x7a\xf6\x9d\xb6\xf4\xd5\xa7\xa0\x32\x2a\x41\x2a\x41\xb7\xd2\x30\x40\x8e\xa5\x82\x2d\x56\x99\x7e\x78\xb7\x9c\x69\x9a\x48\x04\x2f\x8c\x52\xcb\x67\xee\x1b\xd8\xeb\xc7\x64\xff\xe3\xa2\x27\xde\x81\x24\x31\x8c\xe5\x2e\x78\xff\x11\xac\xcf\x46\xbb\xdd\x78\xea\x72\xdc\xe1\x40\x0e\xfe\x06\x00\x00\xff\xff\xd4\x6b\xe6\xef\x15\x09\x00\x00")

func dataFishBytes() ([]byte, error) {
	return bindataRead(
		_dataFish,
		"data/fish",
	)
}

func dataFish() (*asset, error) {
	bytes, err := dataFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/fish", size: 2325, mode: os.FileMode(420), modTime: time.Unix(1496219525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\x5d\x4f\x2a\x3b\x14\x7d\x9f\x5f\xb1\xad\x4d\x04\x13\x24\xdc\x47\xcc\x24\xe2\x85\x28\x89\xc0\x44\xbd\xfa\x60\xbc\x93\x3a\x53\x98\xe6\x30\xed\xa4\x2d\x78\x72\xc4\xff\x7e\xd2\x0e\xcc\x67\xc1\x79\x21\xb3\xbf\xd6\xec\xb5\x57\x77\x39\x3f\xeb\x7f\x30\xde\x57\x89\xe7\x79\xdb\x55\xa7\x0b\x5f\x1e\x00\x40\x44\x14\x05\x84\x07\x08\x18\xb7\x06\xf3\x90\x48\xb3\x2d\xd1\xb4\x5b\x58\xcc\x13\x6e\x57\xe1\xc1\x03\x08\x7f\xdd\x0c\xff\xf9\x46\xb5\x88\xeb\xeb\xe2\x35\xa6\xc7\x8b\x94\xbe\x1f\xcb\x5c\xd6\xb3\x69\x94\x08\x88\xc8\x7a\xcd\xf8\xca\x7c\xe5\x86\xac\x61\xbb\x82\x48\xa4\x29\xe1\x71\x2d\x74\x6f\x33\x6e\x7c\x63\x3d\x54\x91\xc8\xf3\xbe\x3d\xaf\xda\x48\x41\x04\x5b\xc2\x1b\xf4\x38\x20\xfc\x32\x7d\x7c\xfe\x6f\xf4\x70\xb7\x40\xf0\x7e\x0d\x3a\xa1\x25\x31\xdb\x55\xa5\x31\x6b\x5d\x32\x2f\x2f\xfe\x3b\x13\x52\x43\x91\xeb\xe3\x41\xa5\xec\x9f\xd3\x65\xdb\xc9\x9d\x0f\xa2\x28\x27\x29\x05\x1c\xbc\x8e\xbb\x27\xa1\xc2\x60\xf4\x7c\xef\xe3\xfb\xc5\x6c\xd2\xbf\xda\x32\x69\x58\x59\x89\x7e\x09\x58\x4b\x0b\xcb\xbc\xc5\xc3\x38\x78\x1a\xf8\x38\x78\x1a\x1c\x8f\xb8\x5b\xe4\xe5\xf3\xdf\x53\x71\xb7\xd3\xb9\x09\xbb\x9d\xce\x4f\xe0\xd9\x5a\xb6\x52\x35\xe8\x00\x52\xef\x69\xe8\x00\xdd\xc3\xd4\x03\x8d\xb2\xab\x41\xb6\x18\xca\xbf\x65\x68\xd1\xd0\xb1\x61\x84\xe3\xe9\xd3\xe8\xf6\x61\x12\x06\x8f\x8b\x59\xf0\x7c\x74\x36\x86\x28\xd4\x29\xf3\xba\x60\x68\x43\xb5\xc1\xa4\xbf\x62\x26\xa1\x97\x41\xf3\xf3\x94\x8c\xfa\x8d\x88\x9c\xa7\xbd\x18\x4b\x4d\x35\xe4\xf8\x93\x6e\xcc\x71\x98\x48\x29\xe4\x10\xe6\x02\x56\x02\x28\xdf\x32\x29\x78\x4a\xb9\x06\xa6\xf2\x83\x4c\x81\x68\x93\x07\xa9\x30\xf6\xaa\x98\x13\xba\xce\x8a\x77\x49\xf5\x46\x72\x18\x1c\x7a\x6a\x74\x8f\x5b\xca\x71\xcc\xc5\xa1\x09\xd7\x9c\x5d\x12\x6b\x0d\xb0\x09\x58\x88\x66\xc3\x15\xad\xe8\xbf\x2d\x69\x87\xe8\x5c\x6a\x75\x2a\xbd\x71\xae\xec\x88\x82\xc7\xc9\x4b\x25\x36\x78\x1d\xfb\x08\xb5\xcd\x56\x41\xe1\xbf\x8b\xd9\x6c\x34\x1f\xfb\xb8\xfe\xee\x85\x66\xd2\x99\x14\x69\xa6\xc3\xfd\x72\x2a\xa6\x7d\x0e\x4b\xc6\x63\x20\x1c\xc8\x46\x8b\xde\x5e\x0d\x4c\x70\x58\xb2\x35\x2d\x05\x81\x10\x9c\xf9\x80\xf0\x69\xe8\xb6\x50\x7e\x48\xa8\xa9\x38\x07\x32\x6b\x07\x81\x13\xeb\xd5\x01\x90\x6b\xa7\xaa\x1c\x27\x69\xb6\x6a\x0e\x63\xe9\x30\xdd\x16\xab\xd8\x50\xdd\xe0\xd0\xc1\x99\xe7\xb5\x33\x0b\x1e\x4b\xe2\x42\x29\x84\xf6\x0d\x9e\x75\x7c\x26\x6c\x4d\x6d\x5f\x8d\x90\x9c\x50\xdb\x50\x2c\x8a\x76\xf2\xd3\xb7\x6c\x87\x57\x16\x6c\x9b\x04\xf3\x70\xfa\x19\x16\x21\xa1\xd9\xe0\x3e\xee\x44\x44\x9f\x2e\x55\xbf\xe5\x0e\xa3\xf6\xeb\xa7\xbf\x27\x00\xe1\x76\xfd\xbd\x24\x4e\xac\x89\xca\x81\x2f\x2e\x5e\x47\xa1\x5a\xfc\x7e\x8a\x8e\xf9\x36\xbc\xe7\xa0\x13\xa6\x40\x69\xc9\x32\x65\xd7\xcc\x9a\x28\x0d\x19\xd1\x89\xb9\x83\x33\xc1\xcd\x2e\x5a\x4a\x91\x5a\xa7\xb1\x5f\x35\xff\x68\x94\xf3\xea\xd8\x95\xd6\xa4\x0a\x76\xa0\x68\x0c\x17\x6a\xd7\x7f\xfb\xbf\xff\x7e\x89\x77\xbb\x8b\x9c\xb2\x58\x70\xab\x9b\x28\xc9\x3e\xe3\x70\xb9\xe1\x91\x49\x52\x7e\x07\x7f\x35\x4c\x6f\x37\xef\xdf\x80\xda\xda\x41\x5d\x97\xa2\xbc\xbf\x01\x00\x00\xff\xff\x89\x52\x9e\x79\x2f\x09\x00\x00")

func dataShBytes() ([]byte, error) {
	return bindataRead(
		_dataSh,
		"data/sh",
	)
}

func dataSh() (*asset, error) {
	bytes, err := dataShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sh", size: 2351, mode: os.FileMode(420), modTime: time.Unix(1496219525, 0)}
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
	"data/fish": dataFish,
	"data/sh": dataSh,
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
		"fish": &bintree{dataFish, map[string]*bintree{}},
		"sh": &bintree{dataSh, map[string]*bintree{}},
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

