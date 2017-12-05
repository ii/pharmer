// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package azure

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x9b\x5d\x73\x9b\x38\x17\xc7\xef\xf3\x29\x34\xbe\x6e\xf2\x18\xfc\xda\xdc\xb9\x7e\xe1\xc9\xa6\x93\x74\x83\xd3\xce\xee\x4e\xc7\xa3\x80\xe2\x52\x3b\x92\x2b\x89\xa4\x69\xa7\xdf\x7d\x07\xec\x60\x07\x10\xfc\x21\x64\x7b\xd3\x66\x90\xac\xf3\x93\xce\xf9\x9f\x73\x8c\xe1\xe7\x11\x21\x2d\x4e\xef\x58\xeb\x94\xb4\xe8\x8f\x50\xb2\xd6\x9b\xe8\x12\xe3\xf7\xad\x53\xf2\xcf\x11\x21\x84\xb4\x7c\x76\x1f\x5f\x25\xa4\xf5\x8d\xb6\x8e\x08\xf9\x1c\xcf\x91\x6c\x19\x08\xae\x92\x79\x3f\xe3\x7f\x09\x69\xad\x85\x47\x75\x20\x78\xb4\xe6\xc7\x40\x2e\x03\x1e\xd0\xdd\x02\xc9\xc7\xa2\xb1\x29\x55\x9a\x5c\xbb\xfb\xa1\x1f\x82\xb3\xfd\x7a\xf1\x25\x46\x95\x0e\x55\x6b\x77\xe1\x73\xfc\xff\xaf\x37\x2f\xb2\x47\x6c\xc0\xa2\x8d\x9b\x3c\x13\x0f\xb9\xe6\xc6\x8c\x6b\x49\xd7\x65\x3b\xf4\xb6\xd3\xaa\x6c\xf2\x6c\xbd\x0e\xb8\x08\x54\x9e\xd5\x0b\x21\xf5\x17\x02\xda\xe6\xd1\xe4\x1a\x00\x73\xf6\x9d\xe6\x5a\x77\x45\x88\x5b\x57\xd1\xe4\x1a\xd6\x3f\x31\xa5\x73\x6d\xec\x39\x8c\x53\x72\x30\x1e\x98\xd2\x35\x28\xc6\x74\x1d\xdc\x0a\x69\x88\xb5\x18\x00\x30\x5c\x79\xdf\xcf\xe3\x37\x63\xb0\x24\xb8\xb7\x26\x2b\x04\xf7\x9f\x21\xbb\x61\x1e\x19\x07\xfa\x31\x37\xc6\x29\xa7\x3e\x25\x91\xb2\x8a\x83\x3c\x9e\x17\x49\xab\x42\x90\x09\x29\xb8\x16\x05\x66\x77\x0e\x06\x2c\xef\xfc\x8b\x1b\x77\xa9\x20\x1f\x68\xb8\x16\xc4\xd5\x54\xb3\x3c\x88\x77\x92\xfe\x08\xd6\x24\x0e\xf9\x42\x84\x9b\x78\x62\x1c\xed\x15\x34\x2e\xd9\x9a\x72\xdf\x2c\xf1\x69\x28\xc5\x86\x95\xcb\x9b\x6d\xe7\xc1\x86\x2f\x98\xfe\xc2\x64\x64\x3b\x57\xe1\x71\x9c\x01\xb6\xa3\x48\xab\x6a\x7a\x4c\xa5\x1f\xdc\xde\xe6\x99\xbd\x3e\x27\x91\xe5\x42\x93\xe1\x2a\x32\x8a\x9b\x7b\x2f\xb8\x2f\xb8\xc1\x5a\xb9\x5b\xc3\x55\x45\x97\xba\x01\x5f\xd2\x8d\x90\xb9\xd1\x14\xdb\x8b\x04\x42\x46\xea\x30\xa5\x98\xf2\x66\x34\x95\x46\x33\x61\xf3\xff\x17\x7c\x49\xce\x05\x5f\x1a\x6b\x63\xa9\xe5\xea\x46\x2f\xd8\xc3\xf6\x28\xc9\x27\xba\x66\xb9\x11\x35\x0a\x55\xa4\xcd\x00\x48\x23\xf4\x69\x6a\xb5\x4c\xf2\x31\xf0\xb4\x90\xf9\x89\x7a\x6f\x3d\xf1\x00\x86\x90\x78\xa1\x4a\x46\x5b\x3d\x8a\x37\xc4\xa5\x81\xa6\x77\xb9\x34\x7f\xd0\x0d\xe5\xe5\xe7\xf0\x35\x9a\x56\xcd\xf6\xa5\xa2\xab\x02\x93\xa5\xe2\x8a\x4d\xe6\xe8\x2b\x69\x07\x03\xae\x34\xe5\x1e\x9b\x3f\x6e\x58\x4e\x53\xa8\x56\x61\x1c\xe6\x9a\x72\x9f\x4a\x7f\x31\x6a\xef\xcd\xf9\x4c\x79\x32\xd8\x24\x3a\xc9\x9b\xe3\x51\xcd\x96\x42\x3e\xc6\x3e\x3b\x56\x4c\x06\x87\xd1\xe4\x6d\xa2\xd5\xad\xfd\xf6\xe8\x5d\xeb\x94\xb4\x4f\x06\xfd\xe1\xde\x4a\xa0\x56\xad\x53\x62\xb7\x73\xcf\x2a\xc3\x67\x01\x7c\xd6\xcb\xf8\xac\x93\x41\x2f\x8d\x37\x00\xf1\x6c\x00\xcf\xae\x86\x67\xa7\xf0\x3a\x27\x19\x3a\xab\xd3\xc3\xf0\x3a\x00\x5e\xa7\x1a\x5e\x37\x85\x37\xc8\x78\x76\x08\xc2\x75\x01\xb8\x6e\x35\xb8\x61\xda\xb5\xdd\x34\x5d\xbf\x0d\xd2\xf5\x00\xba\xde\xcb\x3c\x9b\xa5\x83\x1d\xdb\x07\xe8\xfa\x2f\x73\xac\x9d\xd5\x2c\xea\xd9\x01\x40\x37\x78\x99\x67\x7b\xfd\xda\x9e\x1d\x02\x74\xc3\x62\x3a\x72\x4c\x3c\x71\xb7\x09\x35\x3b\x0e\xb8\x66\x5c\x05\xf7\x8c\x3c\xa5\xde\x1a\xec\x9d\xa1\x8d\xb1\xbf\x05\xd8\xdf\x36\xc9\x6e\xf5\xd3\x41\x6b\xd9\xb5\xe9\x2d\xa4\xda\x58\x25\xe5\xe6\xb7\x9d\xbd\x05\xd5\xa2\x92\x62\xf4\xfb\x4e\x7f\x02\xe0\x4f\x0c\xf4\x13\xb0\x94\xe6\xd4\xaa\x1e\x56\x49\x27\x40\x25\x9d\x18\x2a\xa9\x89\x2e\x9d\x6f\x33\xa5\xca\x6a\x83\x70\x40\x1d\x9d\x18\xea\xa8\x09\x2e\x9d\x6e\xb3\xc5\xc0\x46\xe9\x80\x42\x3a\x31\x14\x52\x13\x5d\x5a\x36\xd9\x62\xd0\x45\xe9\x10\xd9\x4c\x4c\xb2\x41\x5d\x9b\x53\x4a\x61\x3e\x24\xf2\xac\x8a\xa1\x07\x14\x53\x98\x0f\x09\x3e\xab\x62\xf4\x95\xa7\x45\xdc\xbf\x48\xf8\x59\x15\xe3\x0f\x49\x7c\x43\x98\x70\x71\x0f\xf9\xf8\xd9\xb4\x67\x94\xf7\xf6\x7f\x90\x00\x31\x4a\xbb\x06\x65\x83\x89\x10\x83\xec\xd4\x80\x6c\x32\x21\x62\x94\xdd\x1a\x94\x0d\x26\xc6\x1e\x46\xd9\xab\x13\x96\x69\xf9\x64\xf5\x8d\xab\x07\x95\x4f\x1d\xfd\x34\x9a\xc7\x41\xce\x3a\x0a\x6a\x34\x9f\x83\x9c\x75\x44\xd4\x68\x5e\x07\x39\xeb\xc8\xa8\xd9\xfc\x0e\x2a\xc9\xaa\x23\x25\xbb\x9d\x09\xd1\x76\x4e\x8c\x82\xa8\x2e\xd2\x0c\xb9\xa6\x66\xc8\xad\x5f\x85\x06\x20\x1f\x72\x8e\xae\xe9\x10\x4d\x7c\xe5\xf5\xa7\x0b\xe2\x21\xbd\x90\x6b\xea\x85\x4c\x78\x40\xe5\x19\x82\x7c\x48\x2f\xe4\x9a\x7a\x21\x13\x5f\x79\xcd\xe9\xf5\xd1\xf0\xc3\xe2\xaf\x6a\x00\x96\xa7\x71\xf8\x04\xa1\x76\xdc\x35\xf6\xe3\xa8\x8f\x5f\x72\x86\x50\x10\x1a\x3b\x72\xd4\xcb\xd9\xd4\x6d\x59\xe0\x37\x7d\x17\xea\xc9\x5d\x63\x53\x6e\xcc\x33\x40\xd6\xb6\x6d\x54\xcb\x60\x5f\xe1\x16\xf4\x15\xee\x4b\x1a\x73\x38\x25\x82\x9c\x05\x7d\x85\x99\xb3\xc1\xd4\x08\x62\x16\xb4\x15\x66\xcc\x26\x53\x24\xc8\x59\xd0\x56\x98\x39\x1b\x4c\x95\x60\x53\xe1\x16\x34\x15\x05\xe1\x59\xde\xa0\x57\x50\x3b\xac\xa4\x5a\x52\x6a\x34\xb9\xa3\xa4\xb5\xc4\xd4\x68\x92\x47\x49\x6b\xe9\xa9\xd9\x64\x8f\xa2\xd6\x92\x54\xc3\x49\x1f\x95\x55\x51\xb3\x5e\x10\xab\x40\xb7\x6e\x0f\xb1\x66\x7d\x06\xf4\x4a\x33\x43\xa7\x34\x03\xcb\x52\xe6\x24\x2d\x2c\x40\x67\xc0\x19\xce\x0c\xc7\x67\x62\x4b\xeb\x3c\x23\xf3\x0e\x16\x91\x33\xa0\xf9\x98\x19\x5a\x0f\x13\x5b\x5a\xd9\x19\x61\xf7\xb1\x00\x9c\x01\x3f\xfe\xcd\x0c\x3f\xfe\x99\xd8\x32\x3f\x3a\x67\xb5\x0c\x26\xc8\x99\x05\xfc\xae\xfb\x6c\x12\x14\x73\x69\x01\x77\xb2\xfa\x05\xf3\xe2\xcc\x52\x08\xa0\x32\x00\xaa\xba\xaa\x00\x9d\x6b\x23\x74\x76\x55\xba\x52\x5d\x80\xce\xed\x22\x74\xdd\xaa\x74\xa5\xca\x40\x33\xca\x10\xc1\x1b\x56\xc5\x2b\x17\x07\x9a\x55\xac\x3e\x14\x7b\xfd\xca\xc1\x57\x2e\x0f\x30\xb9\x38\x40\xc1\x70\x0c\x05\xc3\x01\x83\x2f\xdb\xd2\x74\x86\x20\x1d\x50\x32\x1c\x43\xc9\x30\xd1\xa5\x83\x2f\xdb\xc6\x0c\xfa\x98\x3a\x1c\xe0\x4b\xb5\x63\xf8\x4a\x6d\xa2\xcb\xc4\x5e\xb6\x71\xb1\x7a\x1d\x4c\x1e\x0e\x50\xd4\x1c\x43\x51\x33\xf1\x65\x22\xcf\xb6\xb3\x35\xb7\x3d\xc0\xf4\xe1\x00\x0f\x2c\x39\x86\x07\x96\x4c\x80\x9d\x4c\xea\xeb\x66\x0b\xaf\xd5\x05\xe3\x0f\xb9\xf7\xe9\x98\xee\x7d\x3a\xe8\xad\xa7\xda\x3d\xbf\x83\xdc\xfb\x74\x4c\xf7\x3e\x8d\x7c\xe5\x12\xe9\x81\xdd\x81\x83\xdc\xfd\x74\x4c\x77\x3f\x8d\x80\x88\x4a\xda\xf0\x19\x22\x32\x31\xdd\x00\x35\x22\x22\x42\xb1\xd1\x2f\x4c\x8e\x8b\x28\xc5\x35\x49\xc5\xc4\x88\x68\xa5\x9b\x7c\x4f\x4a\x9e\xd7\xf5\x24\xf3\x19\xd7\x01\x5d\xe7\x3c\xad\xbb\x91\xe2\x3e\xf0\x99\x8c\x0c\x8f\x92\xd7\xc2\x9e\x56\xdc\xac\xe9\xe3\x4c\xc8\x3b\xaa\xa3\xf1\xdb\x80\xad\x0f\xde\x53\xa0\x9c\x0b\x1d\x3f\x76\x1c\xad\xfb\xb4\x62\xb4\xe6\x17\x2a\xef\x98\x3c\xa1\x9b\x8d\xf2\x84\xcf\x4e\x3c\x71\xf7\x3f\x6f\x1d\x2a\xcd\xe4\xf1\x9e\x26\x5a\x32\x59\xcd\xf4\x31\x9f\xab\xf4\x47\x76\x9f\xf8\x95\x80\xc4\x5c\xcf\x9f\x62\xde\xd3\x6c\xdf\x71\xf3\x04\xbf\x0d\x96\xf1\x26\xff\xbe\xbe\x9a\x2e\xe6\xd3\x8b\xd1\xc5\x7c\x71\x36\x39\x00\x88\x56\x12\x32\x3a\xd7\xed\x0b\x72\x0b\xcd\x38\xe5\x7a\x11\xf8\xcf\x27\x7d\x55\x5b\x4f\x6e\x87\xd3\x4b\xac\xe9\x0d\x8b\x39\xe7\xf1\x30\x39\x4b\x7d\x3a\xe0\x9b\x50\x6f\x3f\xfe\x3d\x79\xc4\xfa\x60\x37\xe5\xec\xee\xf5\x3b\x77\x7c\x75\xf6\x61\x7e\x76\x79\x51\xb2\x03\x15\xde\x24\xc1\x67\xdc\xc7\xe1\x24\xe3\x6e\xdc\x83\x49\xaf\xb0\xa7\xf1\xfb\xb3\x69\xa9\x3f\xbc\x75\xc0\x0a\xfc\xb1\x1d\x36\xee\x60\x1c\x0f\xbf\x1e\xbb\x3b\x1d\x5f\x4d\xe7\x00\xbf\x62\x9e\x64\xba\x68\x0f\x6e\xce\x8c\xf4\x3e\xf2\xe6\x24\x7b\xd9\x50\xa5\x1e\x84\xf4\x0f\xf6\xb3\xfb\x2b\xff\xfd\x81\x4c\x16\x70\xb5\x90\x74\xf9\x6a\xc9\x40\x6d\x97\x7f\x0d\x65\xbb\xf3\xcb\xab\x91\x33\x5d\x8c\xc6\xe3\xcb\xeb\x8b\x42\x7f\xec\x28\x16\xd4\xf3\x44\xc8\x0d\x1e\xc9\x1d\x4c\x9c\x11\x1f\x16\xd9\x9d\x16\x19\xe5\xcd\x6d\x40\xf0\xbb\x2d\x9d\x4f\xff\x42\xb6\xb3\x62\x8f\xf9\x5b\xc9\x0c\x14\x6f\x83\x9c\xa7\xe7\xe3\xf1\x95\x14\x9f\x55\x78\xc3\x24\x67\x9a\xa9\x8f\x4c\xaa\xfc\xd7\x88\xef\xb7\x23\xd1\xc2\xd6\xc9\xf0\xc4\xfc\x04\x6f\x6a\x74\xfb\xee\xf2\x41\xac\xf9\x2c\xba\xa0\x65\xc8\x0e\xca\xca\x37\x9a\xbd\xb6\x91\xc2\xdf\x5d\x3d\x3a\xe4\x8f\xb9\x8f\x7e\xfd\x1b\x00\x00\xff\xff\x47\xac\x93\xe2\x2d\x3d\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 15661, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
