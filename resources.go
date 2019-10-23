// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/templates/migration.tmpl
// assets/templates/model.tmpl
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTemplatesMigrationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xcf\x6b\xea\x40\x10\xc7\xcf\x99\xbf\x62\x5e\xf0\x90\x3c\x34\x81\x77\x14\x3c\x3c\x6c\xe9\x55\xd0\x63\xa1\xae\x71\xdc\x2e\xee\x0f\x3b\x3b\x42\x4a\xc8\xff\x5e\x36\x15\x6a\x5b\x5a\x2d\xf4\x98\xef\x64\x3e\xfb\x65\x3e\x07\xd5\xec\x95\x26\x74\x46\xb3\x12\x13\x3c\xd4\xf5\xea\xd1\x44\x6c\x82\xa3\x88\x3b\x0e\x0e\x75\x98\x6c\x8c\xdf\x2a\x51\x60\xdc\x21\xb0\x60\x6e\x83\xce\x01\xea\x1a\xe7\x4c\x4a\xa8\xeb\x26\x58\x2d\xec\x91\x95\x5d\xa8\xd8\x28\x3b\x57\x91\x70\xd2\xf7\x2b\xb5\xb1\xe7\xf0\xdd\xd1\x37\x58\x38\xfc\xeb\x8c\x2e\xaf\x5a\x2e\x4a\x24\xe6\xc0\xd8\x41\x26\x2d\x4e\x67\xe8\xaa\x3b\x92\x55\x5b\x94\x00\x59\x94\x94\xac\x07\xc4\x2b\x6d\x29\x4a\xc8\x91\x97\x44\x58\x03\x64\x69\x96\xdf\xfb\x3c\x7d\x03\x64\x0f\xe3\xc4\xc3\x19\x4a\x5b\xdd\xb6\xd4\x14\x51\x4a\xc8\xcc\x6e\x48\xff\xcc\xd0\x1b\x9b\x9e\xca\x98\xe4\xc8\x3e\xa5\x90\xf5\x9f\x31\x5d\x87\xac\xbc\x26\x1c\xed\xe9\xf9\xdf\x18\x47\xf1\xc9\x2e\x85\x8d\xd7\xa9\x50\xf5\xdf\x0a\xf1\x5b\x95\xbe\x07\xc4\x28\x78\xaa\x7a\xf6\xf3\x6f\x96\x4c\xa5\xc8\x6f\x71\x40\x9c\x46\xde\x58\xe8\x07\x53\x37\x1c\x0e\x17\x3c\x71\xb0\x76\xa3\x9a\xfd\x07\x4d\x97\x37\xaf\x97\x94\x58\xdf\x2b\xfa\xe2\x16\xd3\x1f\x1b\x7b\x77\x81\x97\x00\x00\x00\xff\xff\x81\xdf\x2a\x26\xe8\x02\x00\x00")

func assetsTemplatesMigrationTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesMigrationTmpl,
		"assets/templates/migration.tmpl",
	)
}

func assetsTemplatesMigrationTmpl() (*asset, error) {
	bytes, err := assetsTemplatesMigrationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/migration.tmpl", size: 744, mode: os.FileMode(436), modTime: time.Unix(1571782333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xcb\x6e\xdb\x3c\x10\x85\xd7\xe2\x53\x0c\x84\x00\xb1\x7f\xe4\x97\xf6\x01\xba\x08\xdc\x2e\x02\x34\x45\x5a\xb7\x5d\x87\x22\x47\xf2\x44\x24\x47\xe5\x05\x8e\x6b\xe8\xdd\x0b\x5d\x92\x3a\x89\xdd\x6c\xb2\xb3\x39\x33\xdf\x1c\x1d\x1e\x76\x52\xb5\xb2\x41\xb0\xac\xd1\x08\x41\xb6\x63\x1f\x61\x21\xb2\xbc\xa1\xb8\x49\x55\xa1\xd8\x96\x71\x4b\xee\xbe\x4c\x89\x74\x2e\x32\x0b\x43\xc9\xc8\xa9\x64\xa9\xf5\xbc\xc5\xdf\x68\xca\x4a\xaa\x16\x9d\x2e\x47\x54\x2e\x96\x42\xc4\x5d\x87\x03\xab\x2c\x61\xbf\x2f\xd6\xe4\x9a\x64\xa4\xbf\x95\x41\x49\xb3\x92\x01\xfb\x7e\xde\x9b\x9d\x2a\x87\xe8\x93\x8a\xb0\x17\x59\x66\x8b\x6b\x8d\x2e\x52\x4d\x4a\x46\x62\x27\xb2\x6c\xbf\xff\x1f\xbc\x74\x0d\xc2\x59\x8b\xbb\x0b\x38\xeb\x3c\x77\x70\xf9\x01\x8a\x5b\xcf\xdd\x47\xac\x43\xdf\xcf\x6d\x54\x83\xe3\x38\x75\x14\xd7\xe1\x93\xad\x50\x6b\xd4\x73\x7d\x3a\xfe\x22\xed\xb0\xf3\xf1\xef\x5a\xd6\xf8\x7d\xd7\x0d\x47\x77\xba\xba\xcc\x9f\xce\xbf\x7e\x5e\xb1\x49\xd6\xf5\x7d\x0e\xf7\x81\xdd\x41\x69\xfe\x86\x95\xb4\x38\x7f\xc2\x05\x5b\x8a\x68\xbb\xb8\xcb\xef\x66\x2d\xe8\xf4\x93\xae\xe9\x37\x00\x80\x2d\xae\x92\xa6\x28\xb2\x7e\xb0\xae\x2c\x61\x8d\x71\xe5\x51\x46\xfc\x29\x4d\xc2\x00\x35\xfb\xc9\xae\x42\xd4\xc9\x29\x58\x1c\x98\x76\xb0\x10\xfe\x3b\x61\xe6\xf2\x25\x71\xb1\x04\xf4\x9e\x3d\xec\x05\x40\x59\xc2\x37\xec\x8c\x54\x08\xe7\xc7\xc1\xa3\x3f\xc5\x3a\x7a\x72\xcd\xf9\xa8\x46\x3a\x8e\x1b\xf4\x90\x1c\xfd\x4a\x08\x34\xdf\x0f\xfa\x89\x77\x5d\xc3\x30\x02\x9a\x31\x8c\xee\xe3\x03\x85\x08\xec\x61\xc7\x09\x3a\x8f\x35\xfe\x65\xb0\xc3\x62\x1e\x8b\xb0\x25\x63\xa0\x42\x48\x01\xf5\xb8\x69\x6d\x52\x03\x0d\x3a\xf4\xe3\xdd\xcf\x9d\x57\xd0\x19\x49\x6e\x48\x09\xb9\x66\x9a\x92\x26\x30\x6c\xd9\xb7\x50\xa5\x08\x34\xb3\xac\x6c\x11\x0c\x73\x0b\x81\x2c\x19\xe9\x41\x1a\x03\xc1\xa4\x66\xf2\x15\xa5\xda\x00\xb9\x10\xa5\x53\x38\xc1\xb9\x86\xb8\xa1\x00\x1e\x03\x27\xaf\xb0\x18\x73\x4c\x05\x16\x97\x70\xdc\xa0\xdb\xfa\x61\x08\x5f\x3e\x60\xbb\xfa\x21\x7f\x16\xeb\xd7\x7d\x6f\xbb\x7c\x0a\xf0\xe2\x25\x14\x2f\xef\xf5\xe4\xda\xe5\x49\xe2\x98\xbd\x57\xa0\xa5\xc8\x3c\xc6\xe4\x1d\x38\x32\xa2\x7f\x8c\xe5\x8f\x4e\xbf\x73\x2c\x0f\x89\x07\xb1\x7c\x53\xed\xf3\xb9\x23\x6a\x6f\x64\x54\x1b\x50\xec\x34\x0d\x56\xbd\x83\xda\x91\xb8\x88\xea\x1f\x2d\x15\xb3\x19\xd4\xfb\xf9\x9a\x6f\x86\x85\xe3\xdc\x8a\xa7\xb7\x3f\xeb\xf4\xa2\x17\x7f\x02\x00\x00\xff\xff\x52\x01\xd0\xff\x82\x05\x00\x00")

func assetsTemplatesModelTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesModelTmpl,
		"assets/templates/model.tmpl",
	)
}

func assetsTemplatesModelTmpl() (*asset, error) {
	bytes, err := assetsTemplatesModelTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/model.tmpl", size: 1410, mode: os.FileMode(436), modTime: time.Unix(1571788620, 0)}
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
	"assets/templates/migration.tmpl": assetsTemplatesMigrationTmpl,
	"assets/templates/model.tmpl":     assetsTemplatesModelTmpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"migration.tmpl": &bintree{assetsTemplatesMigrationTmpl, map[string]*bintree{}},
			"model.tmpl":     &bintree{assetsTemplatesModelTmpl, map[string]*bintree{}},
		}},
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
