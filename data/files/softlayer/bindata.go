// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package softlayer

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5c\x5d\x53\xdc\x3a\x12\x7d\xe7\x57\xa8\xe6\x89\xad\x72\xd8\xf1\xe7\x9a\xbc\xcd\x25\x24\x97\x4b\x48\xd8\x40\x76\x37\xb5\x95\xa2\x14\x5b\xcc\x38\xe3\x91\x1c\x49\x06\x26\x29\xfe\xfb\x2d\x1b\x98\x0f\xcb\xb4\xac\xd0\x54\xa5\xf2\x92\x80\xc6\xa3\x3e\xad\xee\x3e\x67\xba\x19\xfb\xc7\x0e\x21\x23\x4e\x17\x6c\xf4\x92\x8c\x94\xb8\xd4\x25\x5d\x32\x39\xf2\x9a\x65\xc6\xaf\xd4\xe8\x25\xf9\xff\x0e\x21\x84\x8c\x72\x76\xd5\x2e\x13\x32\xfa\x46\x1f\x7e\xaa\xa4\xc8\x47\x3b\x84\x7c\x6e\xdf\x20\xd9\xb4\x10\x7c\xfd\x9e\x1f\xed\xbf\x84\x8c\x4a\x91\x51\x5d\x08\xde\x18\x99\x2c\x94\x66\x32\xa7\x0b\x8f\xbc\x63\x7a\xc6\x64\x49\x79\xae\xee\x37\x5c\x6d\xd2\x5c\x49\x17\x1b\xcb\xdf\x05\x67\xeb\x9d\xdb\x25\xba\x50\x63\x7f\x75\xc5\xfd\x42\x38\xba\xff\xfd\x73\xfb\xff\xad\xf7\x38\x92\x83\x19\xe3\x9c\x16\x1e\x39\xe2\x79\x41\xfb\x10\x64\x33\x06\x22\xc8\x66\x6c\xec\x0f\x37\xf8\x8a\x96\x25\x55\x1e\xf9\x78\x36\xe9\xb3\x96\xd3\x12\xb4\x96\xd3\x72\xdb\xdf\x66\x21\xee\x2e\x24\xdd\x85\xfd\xce\x82\x3f\x1e\x8e\xf8\xb5\xa4\x7c\x7e\x59\x4b\xed\x91\x37\x4c\x2e\x28\x5f\xf6\x01\xbf\x94\x14\x04\x7e\x29\xe9\x38\x18\x6e\xf4\x4f\xc1\xa7\xe4\x58\xf0\xa9\x47\x0e\x66\x05\xef\x8d\xcc\x6c\x3e\x05\x4d\xce\xe6\x53\x37\x93\xb5\xd2\x82\x3f\x1a\x9a\x99\xa8\x61\x73\xa2\x76\x31\xf7\x56\xf0\xbc\xb5\x76\xdc\x67\xac\x14\x1c\x34\x56\x0a\xee\x62\xec\x84\x95\x5f\x44\x2d\x39\xf3\xc8\xa4\x56\x5a\xd2\xb2\x3f\xd9\x17\x0c\x4e\xbf\x05\x2b\x5d\x92\xfd\xdf\x35\x93\x4c\x53\x29\x3c\x72\xc2\x6e\x8a\x4c\xf4\xdb\xbc\xb1\xd8\xbc\x71\xb1\x79\x52\x94\x94\x7b\xe4\x48\xd3\xb2\x37\x51\x17\x85\xc5\xc5\xc2\xc9\xc5\x13\xc1\xb5\x64\xb4\xf4\xc8\x01\xe5\x34\xef\x3f\x55\x4b\x30\x17\x82\xbb\x98\x7c\xaf\x4a\xe1\x91\x77\x42\x5e\xd3\x5e\x0f\x85\x82\x3d\x14\xca\xc9\xc3\x53\x2a\x0b\xe5\x91\x86\x06\x32\xd6\x67\xaf\xa2\x12\xb4\x57\x51\xe9\x62\xef\x8c\x72\xf2\x97\x50\xec\xd1\x42\x54\x5f\x33\xd0\x9e\xfa\x9a\x6d\x73\x64\xb3\xe0\xa0\x09\x67\x54\x90\x53\x5a\x37\x87\xfc\x18\x02\x2a\x60\x04\x54\x38\x79\xcc\x44\x5d\x7a\xe4\x58\x48\xd6\x9b\x3f\x8a\x59\xcc\x31\x47\x73\x54\xeb\x12\x38\x5f\x06\x53\xb9\x62\xd4\xc9\x5c\xc1\xa7\xb4\x12\x92\x79\x64\xf5\x63\xaf\x59\x0e\xd3\xb9\xe2\x53\x27\xb3\xcb\x9c\xb3\xa5\x85\xef\xd4\x32\x87\x6d\x2e\x73\x17\x9b\xe7\x62\xbe\x14\x1e\xf9\x8b\x56\x94\xf7\x99\xd3\x62\x0e\x9a\xd3\x62\xee\xc2\xea\xe7\x42\x0a\xae\x05\x44\x3d\x5a\xc0\xb5\xa9\x85\x53\x6d\xfe\x97\xaa\x59\xc1\xa7\xad\x4c\xbe\x3a\x78\x34\x83\xae\x73\xb8\x42\xaf\xf3\x4e\x85\x36\x0b\x51\x07\xc6\xea\x23\x65\xc1\x95\x6e\xb8\xe7\x7c\x59\xb1\x9e\x0f\x96\x6a\x5e\x37\x36\xfd\xcc\x5f\xac\x8d\xe6\x4c\x65\xb2\xa8\x1e\x70\xfb\x24\x13\x92\x11\xca\x73\xe2\xbf\xf9\x83\x7c\x98\x9c\xac\x2f\xcd\xa8\x66\x53\x21\x97\xf7\x9f\x01\x64\xb9\x24\xff\x29\xa4\xae\x69\x49\xce\x98\xbc\x62\x1b\x07\x98\x55\x8d\x29\x7f\xed\x31\x5d\x34\xbf\xf7\x1e\xdb\x0a\x56\x30\x08\x56\x80\x0e\x2b\x80\x61\x45\x83\x60\x45\xe8\xb0\x22\x18\x56\x32\x08\x56\x82\x0e\x2b\x81\x61\xa5\x83\x60\xa5\xe8\xb0\x52\x18\x96\x3f\x2c\xb9\x7c\xfc\xec\xf2\x2d\xe9\xe5\x0f\x0b\xa4\x8f\x1f\x49\x1f\x0c\x65\x00\xd1\x44\xd0\x02\x53\x48\x3c\x11\xb8\xf0\x44\x00\xf1\xc4\x26\xae\xa7\x87\xb2\x8b\x0b\x8c\x64\x00\x11\xc5\x26\xae\xa7\x33\x45\x17\x17\xc8\x14\x01\xc4\x14\x9b\xb8\x9e\x9e\x60\x5d\x5c\x96\xfc\x02\xa8\x62\x13\xd7\xd3\xb9\xa2\x8b\x0b\xe4\x8a\x00\xe4\x8a\xad\xc4\xc7\xcf\x30\x98\x2c\x02\x90\x2c\xb6\xa0\xe1\x07\x13\x66\x8b\x08\x62\x8b\x08\x95\x2d\x22\x17\xb6\x88\x20\xb6\x88\x50\xd9\xa2\x8b\x0b\x0c\x65\x04\xb1\x45\x84\xca\x16\x5d\x5c\x20\x5b\x44\x10\x5b\x44\xa8\x6c\xd1\xc5\x65\xc9\x2f\x80\x2d\x22\x54\xb6\xe8\xe2\x02\xd9\x22\x02\xd9\x22\xc2\x65\x0b\x23\xf3\x2d\x29\x06\xb1\x45\x84\xcb\x16\x06\x34\x4b\x34\xc3\x81\xa7\x16\xe2\x9f\x5a\x68\x2b\xcc\x81\x99\x16\xe1\xa7\x5a\x64\xc9\xb5\x64\x20\x69\x24\xf8\xac\x91\x80\xb4\x91\x42\xf4\x9f\xa2\xd2\x7f\xea\x42\xff\x29\x44\xff\x29\x2a\xfd\x77\x71\x81\x59\x96\x42\xf4\x9f\xa2\xd2\x7f\x17\x97\x25\x8e\x00\x65\xa4\xa8\xf4\xdf\xc5\x05\x12\x46\x0a\xd1\x7f\x8a\x4a\xff\x5d\x5c\x60\x49\xa6\x20\xfd\xa7\xb8\xf4\x6f\x64\xbe\x25\xc5\x20\xfa\x4f\x71\xe9\xdf\x80\x66\x89\x26\x44\xff\x29\x2e\xfd\x77\xa1\xc1\xf4\x9f\x82\xf4\x9f\xe2\xd2\xbf\x51\x9a\x96\x5c\x83\xe8\x3f\xc5\xa5\x7f\xa3\x3a\xe1\x69\x14\x38\x2c\xf0\x71\xa7\x05\xbe\xd3\xb8\xc0\x07\xe7\x05\x3e\xee\xc0\xc0\x80\x06\xcf\x7e\xc0\x91\x81\x8f\x3b\x33\x30\xa0\xd9\x02\x0a\x8d\xa5\x70\xc7\x06\x06\x34\x78\xc4\x08\x0e\x0e\x7c\xdc\xc9\x81\x01\x0d\x1e\x33\xc2\xb3\x03\x1f\x79\x78\x60\x16\x82\x2d\xdd\xc0\x59\x23\xf2\xfc\xc0\x44\x67\x0b\x2b\x24\x0a\x5b\xe8\x10\x54\xc1\x40\x07\xcb\x42\x53\xaa\x43\xb3\x0e\x41\x18\xcc\x62\xb5\xe5\x1d\x24\x0d\xdb\xe5\xfa\x0c\x54\x62\x11\x87\x04\x14\x87\x04\x57\x1c\x12\x27\x71\x48\x40\x71\x48\x70\xc5\xa1\x0b\x0d\xce\xb8\x04\x14\x87\x04\x57\x1c\xba\xd0\x6c\x01\x85\x78\x24\xc1\x15\x87\x2e\x34\x98\x45\x12\x50\x1c\x12\x5c\x71\xe8\x42\x83\x8b\x34\x81\xc5\x21\x41\x16\x07\xa3\x10\x6c\xe9\x06\x8a\x43\x82\x2c\x0e\x06\x3a\x5b\x58\x41\x71\x48\x90\xc5\xa1\x8b\xce\x22\x0e\x09\x2c\x0e\x09\xb2\x38\x18\xc5\x6a\xcb\x3b\x50\x1c\x12\x64\x71\x30\xea\xd5\x2a\x0e\xc1\xd0\xc3\xf3\x83\x67\x38\x3d\x3f\x00\x8f\x2f\x04\x5b\x9b\x10\xb7\xb5\x09\x9d\x5a\x9b\x10\x6c\x6d\x42\xdc\xd6\xc6\x80\x06\x96\x44\x08\xb6\x36\x21\x6e\x6b\x63\x40\x03\x33\x2e\x04\x5b\x9b\x10\xb7\xb5\x31\xa0\x81\x34\x17\x82\xad\x4d\x88\xdb\xda\x18\xd0\xac\x65\x30\x34\xd9\x10\xd4\xcb\x2c\x04\x5b\xba\x41\xea\x15\x22\xb7\x36\x26\x3a\x5b\x58\x21\xf5\x0a\x91\x5b\x1b\x03\x1d\xac\x5e\x21\xdc\xda\x84\xc8\xad\x8d\x59\xac\xb6\xbc\x83\xd4\x2b\x44\x6e\x6d\xcc\x7a\xb5\x71\x09\xa8\x5e\x9d\xb2\x78\x86\xd3\xb3\xab\x57\x10\x0d\x16\x89\xe8\x39\x64\x22\x02\xb3\x2f\x06\x9b\xc3\x18\xb7\x39\x8c\x9d\x9a\xc3\x18\x6c\x0e\x63\xdc\xe6\xd0\x80\x66\x3b\x35\xa0\x28\x62\xdc\xe6\xd0\x80\x06\x96\x44\x0c\x36\x87\x31\x6e\x73\x68\x40\x03\x79\x38\x06\x9b\xc3\x18\xb7\x39\x34\xa0\x81\x75\x1a\xc3\xcd\x61\x8c\xdc\x1c\x9a\x85\x60\x2d\xd2\xa1\x41\x45\x90\x57\x13\x9d\x2d\xac\x90\xbc\xc6\xc8\xcd\xa1\x81\x0e\x96\xd7\x18\x6e\x0e\x63\xe4\xe6\xd0\x2c\x56\x5b\xde\x41\xf2\x1a\x23\x37\x87\x66\xbd\xda\xb8\x04\x94\xd7\x18\xbb\x39\xec\xa9\x0b\xdb\xf1\x81\xf2\xba\x2d\x12\x08\xf2\x6a\xca\x04\x2c\xaf\x47\x5c\xb3\x92\xfc\x8f\x09\x4e\x0e\xc3\x17\x7e\xf0\xaf\x31\xb9\x0a\x1f\x85\x7b\x56\xf0\x69\xc9\x48\xef\x9b\xc8\x6e\x44\x0e\x1a\x5f\x3c\x12\xee\xc5\x63\xf2\xe6\xcf\xef\xff\x00\x1d\xf9\x83\x4a\x46\x4e\x98\x46\xfa\x5e\xce\x06\x28\x72\x18\xbf\x08\x92\x00\x74\xe5\x55\x73\x86\xfd\xef\x21\xbb\x7e\xf0\xe0\x4a\xb0\x17\x61\xb8\xe2\x38\x11\xdf\x3c\xe0\x07\x58\xd1\x4f\xb8\x12\x91\x5d\x3f\x59\xbb\xe2\xa3\xb8\xe2\x56\xa2\xbf\x65\x82\xfd\x44\x50\x7e\xd5\x98\x34\xb0\x62\xa7\x4a\x59\xbf\x85\xec\x06\xe3\xb5\x27\x21\x86\x27\xc1\xd8\x8d\x60\xbb\xb8\xf6\xdd\x5d\xd9\xbf\x77\x25\x5a\xbb\x92\xa0\xb8\x62\x7e\x1b\xd3\xd9\x15\xc7\xfc\xda\xbf\xcb\xaf\x20\xc5\x76\xc5\xfc\x66\xd1\x6f\x13\x15\xb7\x52\xf9\x95\x3d\x09\x62\xf0\x63\xf0\x6f\x9c\x5f\xbf\x32\x15\xaf\x5c\x59\xdd\x85\x99\x49\x96\x33\xae\x0b\x5a\xf6\xdc\x83\x59\x49\x71\x55\xe4\x4c\xb6\x6a\xb8\xf5\x04\x91\x3b\xd7\x0a\x55\x95\x74\xf9\x5a\xc8\x05\xd5\xcd\x35\x97\x05\x2b\x37\xee\xbc\xa5\x9c\x0b\xdd\xde\x64\xda\xec\xfd\xb0\x6b\xb3\xef\x8c\xca\x05\x93\x7b\xb4\xaa\x54\x26\x72\xb6\x97\x89\xc5\x3f\xb3\xb2\x56\x9a\xc9\x17\x6b\x44\xcd\x96\x0f\x77\x8f\xde\xae\x76\x6d\x8d\x6c\xdf\x79\xba\xde\xfa\xee\xd9\x26\x99\xe0\x97\xc5\xb4\x45\xfd\xfe\xf5\xf9\xdb\xc9\xa7\xc3\x0f\x17\x93\xd3\xa3\x8b\x8f\x67\x87\x1f\xde\x4d\x4e\x0e\x57\x10\xef\x36\x14\xb2\x39\x9b\xf5\x33\x52\x2e\x68\x55\x5c\xd4\x8a\xc9\xf6\xe9\x29\x5b\xd7\x7e\x55\x77\x61\xec\x7f\xb5\xa4\x5f\x58\x8b\x7a\x72\x7a\x44\x3e\xf6\x5e\x52\xf0\xaa\x6e\xcf\x4a\xb3\x1b\x3d\x5a\xbd\x72\xeb\x39\x3b\x73\x7c\xf8\x69\x88\x1f\x73\xb6\xec\x77\x81\x56\xc5\x71\xf7\xb5\x2d\x07\x8c\x57\x57\xd8\x2b\xaa\xd4\xb5\x90\xf9\x06\xfe\xfb\x9f\xba\x37\xf9\xce\xeb\x2f\x4c\x72\xa6\xfb\xee\xf0\xbd\x62\x52\x3d\xfc\xe1\x6c\x2f\xdd\x1b\x3f\xfe\x67\xb5\xed\x57\xef\x1f\x5e\xb3\x91\x4f\x39\xbb\x1a\xbd\x24\x5a\xd6\x6c\x8d\x77\xf4\x8d\x9a\x6b\xed\x43\x6d\xee\x56\x77\x36\x81\xb7\x80\x77\x6e\xff\x0e\x00\x00\xff\xff\x6c\x18\xc3\x8a\x33\x47\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 18227, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
