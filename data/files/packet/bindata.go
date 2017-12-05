// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package packet

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x5f\x4f\xdb\x3c\x14\xc6\xef\xfb\x29\x8e\x72\x9d\xb7\x2f\x49\x4b\x55\xf5\x2e\x84\x6c\x2a\x0c\x16\xd1\x32\x54\x4d\x28\x72\x93\x03\x33\x4d\x6c\xcf\x76\x0b\x1d\xea\x77\x9f\x9c\x40\x5c\xd2\x50\xfe\x68\x37\xa0\xfa\xb1\x9f\xe7\x77\x62\x1f\xfb\xb1\x03\xe0\x30\x52\xa0\x33\x02\x47\x90\x74\x81\xda\x71\xcd\x18\xb2\x95\x33\x82\x9f\x1d\x00\x00\x27\xc3\x95\xd3\x01\xb8\x2e\x15\x89\xb7\x94\x33\x55\xab\x8f\xe5\x5f\x00\x27\xe7\x29\xd1\x94\x33\x63\x75\x3e\x0b\xe1\x0c\xb5\xe4\x2e\x9c\xcf\x5c\xb8\x9c\x04\xa5\x6d\x39\xaf\x32\x30\xb3\x2e\x27\x10\x11\xa5\xad\xf4\x87\x33\xb4\xce\xe5\x10\xde\x4b\xcf\x79\xfa\x79\x5d\xfe\xdf\xb8\xaf\xe7\x4e\x08\x83\x13\xae\xd0\x85\x30\xd8\x17\x7b\x85\x6f\xc4\xaa\xbb\xf4\x03\xb1\x41\xa1\x34\xca\x8c\x14\x2e\x9c\x7f\x6b\x8b\x8c\x2e\xe1\x2a\x9a\x4c\xf7\x46\x92\x42\x7d\x20\x72\xca\x17\x6b\xee\xc2\x49\xdc\x16\x17\x4c\xc6\x01\x44\xc1\x1b\x81\x4c\xea\x66\x60\xbd\xcb\x94\x29\x4d\x58\x8a\xd3\xb5\xc0\x96\xbd\x56\x8b\xa5\xc9\x99\x13\x89\x05\x6a\x92\x27\x07\x36\x29\x43\x95\x4a\x2a\x6a\xd0\x59\x1c\xc1\xc1\x08\xfa\xf0\x30\x1c\xc0\xa0\x3f\xa7\x1a\x42\x2e\x51\xb9\x30\xfc\x7a\x04\xc7\xc7\x17\x3d\xb8\x08\xce\xec\xfa\x94\x68\xbc\xe5\x72\x6d\x16\x1f\x11\x89\xe6\x20\x91\x7c\x4b\x17\x26\xbb\x6f\xcb\x26\x85\x33\x82\xa1\xcd\xa7\x6a\x61\x06\x0e\x5a\xbf\xe1\x0e\xb9\xb7\x9f\xdc\x6b\x25\xef\xf9\xff\x10\xbd\xe7\x37\xd9\x3d\xff\x9d\xf0\xfe\x7e\x78\x7f\x04\x7e\x0b\xbd\x7f\x38\xa8\xf0\xfb\x10\x85\xe1\x67\x4a\xf0\x9b\x35\xf8\x87\x83\x66\x11\xfe\xf0\xbd\x5b\xd0\xdb\x5f\x45\x6f\x04\xde\x60\xb7\x0a\xcf\x1f\x3e\x57\xf1\x89\x0a\xbc\x41\xa3\x02\xcf\xdf\x39\x42\xf5\x36\xd4\x6d\x91\x4a\xcc\x90\x69\x4a\xf2\x96\xa6\x10\x92\xaf\x68\x86\xd2\x04\xc7\xf6\x2e\x7d\x76\x14\x39\x59\x7f\xe1\xb2\x20\xda\x4c\xb8\xa1\x98\x67\x56\x27\x8c\x71\x5d\x36\xb7\x31\x7e\xb4\x4d\x2a\x7e\x11\x59\xa0\xec\x12\x21\x54\xca\x33\xec\xa6\xbc\xf8\x3f\xcd\x97\xe6\xc2\xf9\xcf\xe2\x18\xcb\xe7\x5e\xde\xd4\xae\x65\xc8\xcb\xae\xb7\xd6\xd5\x2d\x9f\x72\x76\x43\x6f\x4b\xe4\x20\x3c\x8d\xa6\x49\x7c\xf1\xfd\x24\x0a\xa7\xc9\xf8\xb8\xa6\xab\xbc\xb8\x2c\xec\x2b\x91\x08\xc9\xef\x30\xd5\x09\xcd\x5e\x4e\xbb\x53\xd5\xce\x3d\xe9\x4d\x97\x9c\xcc\xb1\x84\x8d\x2b\x1d\xc6\x8d\xf5\x94\x89\x65\xf9\x81\x34\x3e\x68\xa7\x56\x36\xee\xfb\x2b\x08\xe2\x71\x72\x1a\xcd\xf6\xe2\x13\x41\x93\x05\xae\xdb\xd9\x89\xa0\xa7\x4d\xad\x06\x0f\xe2\x31\xec\xa8\x35\xb5\x20\x4a\xdd\x73\x99\x6d\x91\xbf\x72\xc3\x2e\x96\x73\x94\x0c\x35\xaa\x1f\x28\x55\xfb\x93\xba\xaa\x14\x63\xec\x75\x87\xdd\xd7\x2f\xd9\x86\x5a\xbd\xde\x5b\xa7\xc8\xbc\xe0\x23\xd0\x72\x89\x16\xdb\xf9\x4d\x76\xc7\x84\xe4\xd9\xd3\x68\x67\x9b\xbf\xe4\xee\x6c\x3a\x7f\x03\x00\x00\xff\xff\xcf\x43\x8f\xc9\x31\x08\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 2097, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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
