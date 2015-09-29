// Code generated by go-bindata.
// sources:
// data/aws-simple/build/build-php.sh
// data/aws-simple/build/template.json.tpl
// data/aws-simple/deploy/main.tf.tpl
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package phpapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsSimpleBuildBuildPhpSh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x55\xdd\x6e\x1b\x37\x13\xbd\xdf\xa7\x98\x28\xc6\x97\xaf\x45\x28\x2a\x06\x5c\x14\xf2\x0f\xaa\xc8\x4a\x6d\x20\x90\x0c\xa9\x4d\x2f\x5c\x23\xa0\x96\xb3\x2b\xc6\xbb\x24\x43\x72\x25\xcb\x8e\xde\xbd\x43\xae\x2c\x5b\x2e\xea\x2b\x71\xc9\xf9\x39\x73\xe6\xcc\xe8\xed\x1b\x3e\x57\x9a\xcf\x85\x5f\x64\x99\xc7\x00\xcc\x80\x36\x8d\xde\x1e\xd1\x39\xbc\x53\xe9\x68\x95\xc5\x42\xa8\x6a\x7b\x1d\x9c\xc8\x31\xcb\xe8\x64\xdc\xff\x7f\x82\x87\x0c\x00\x2a\x93\x8b\x0a\xbc\x69\x5c\x8e\x85\xaa\xf0\xf4\xe0\xc3\xd3\x75\xa5\x34\x6a\x73\x7a\x70\x18\xaf\x30\x5f\x18\xe8\x8c\xa6\xd3\xc9\x14\x44\x80\x83\x87\x27\xa7\x4d\xff\xe0\xa1\xb5\xdd\x1c\xc3\x67\xe1\x03\xf9\x97\xbe\xdf\x89\x6e\xa5\x43\x0b\x26\x04\x03\x7c\x29\x1c\xa7\x07\xee\xd7\x9e\x7e\xe0\x07\x84\x84\x4d\xc3\x61\x2f\xdb\x64\x84\xce\xc2\xbb\x04\x0e\x3a\x07\x0f\x1f\x07\xb3\x8b\xaf\xb3\xc9\x9f\xd3\xe1\x68\xd3\x89\x17\x9f\x2f\xc7\xa3\xf1\x64\xd3\x79\x07\x84\x21\xcb\x0c\xc6\x12\xe8\xe1\xb7\x0e\x1c\x9e\xfd\xef\x03\x85\xa3\xa0\x25\x3a\x60\xa1\xcd\x77\x06\x5c\xe2\x92\xeb\xa6\xaa\x8e\x61\x93\x99\x2a\x39\xb4\x65\x5c\x47\x8b\x1b\x20\xe7\xf8\x94\xbd\x85\xbc\x32\x8d\x64\xb9\xd1\x85\x2a\x21\x17\x1a\x94\x0e\xe8\x0a\x74\x08\x2b\x15\x16\x20\x6c\x80\xdc\xd4\xb5\xd0\xd2\x83\x2a\x40\x85\x77\x1e\x7c\x50\x55\x45\x96\x60\x9d\xa1\x3a\xbd\xa7\x24\xd0\xf9\x4b\xa8\xa0\x74\x09\x05\x15\xb2\x17\x96\x30\x51\x08\x5b\x61\xc0\x6e\xb7\xdb\xc9\x1a\x4d\xfe\x70\x7d\x0d\xac\xd8\x92\xa3\xe6\x3c\x79\x70\xa5\x7d\x10\x3a\x47\x3e\x37\x26\xb0\x42\x69\xe5\x17\x28\xe1\xe6\xe6\x18\xa4\x21\x5a\x7d\x85\x44\x6b\xaf\x7b\x94\x49\xa3\xa9\xa7\x31\xef\x40\xca\x98\x36\x22\x25\xce\x8d\x57\xc1\x38\x85\x1e\x08\x32\x34\x56\x8a\x08\x2a\xe5\xc5\x3b\x6b\x5c\x80\xf3\xd1\xc7\xcb\xc1\xf8\xeb\xa7\xe9\x64\xfc\xc7\x68\x7c\x7e\xaa\x8d\x4e\x45\x8b\x3c\xa8\x25\x12\xc1\xe0\x1b\x69\x62\x3c\x56\x92\xb4\x52\x08\x04\xb6\xfe\xd7\x4b\x02\x4b\x44\xb0\x35\x09\xa9\x08\x2b\xe1\x90\x11\x23\x16\x5d\xa0\xfc\x2c\xd2\x66\xf4\x93\x97\x94\x2c\x7a\xee\x30\xae\xa3\xa3\xb5\xa2\x6f\xb4\x74\xf8\x8d\xdb\x85\x3d\x62\x47\xdd\x5f\xa8\x2b\x33\xc4\xda\x47\xde\xe6\x48\x35\x7d\x6f\x94\x23\x12\xe8\xd3\x3a\x5c\xa2\x0e\xd0\x69\xb4\x68\xc2\x82\x8e\x2a\x27\x70\x12\xac\xc8\x6f\x45\x89\xbe\x43\xce\x49\x4a\x1e\x4c\x43\x7a\x28\x5e\x82\xed\xee\x55\x71\x8b\xeb\x6d\x7d\xaf\x94\x9d\x48\xbe\x6c\xdd\x23\xd1\x57\x17\x57\xf0\x65\x38\xf3\xef\x61\x40\x69\x17\xf8\x3e\x31\x6d\x08\x8e\xdb\xe1\x48\x7c\xbf\xc2\x97\x48\x9e\x87\x10\x6b\xa6\x71\x9b\x6f\xbf\x59\x6d\x24\x4b\x77\x7f\x53\xb3\xe7\xf7\x0e\x4a\x9a\xe8\x1a\x5d\xde\x38\x45\x83\x39\x6f\x54\x25\x19\x09\x2e\x56\x4e\xdf\xd1\x8a\x9e\xda\x43\xe2\xaf\xce\xdd\x9a\x84\xd0\x9e\xd7\xfe\x7b\xd5\x1e\x0b\x5b\xb7\x87\x52\xb6\xbf\x0e\x85\x8c\xa3\xdb\x7e\xd9\x92\x2c\xdb\x42\x47\x77\x21\x29\x21\x29\xca\xa6\x3a\x52\x11\xf5\xad\x54\x34\x63\x16\xb8\x77\x4b\x1e\xc7\x88\xba\x69\xdb\xb7\x20\x1c\xdc\xdf\x91\x98\x43\x6d\x77\x4f\xdd\x50\xde\x03\x1b\xbe\xb0\x4f\x39\x66\x18\x52\x02\x52\x4a\xad\xbc\x57\x46\xef\x13\x46\xa3\xba\xd2\xc0\xa6\xb0\x5a\xad\x18\xb5\x41\xf4\x5f\x46\xa1\x49\x6c\xc7\xe7\xf9\x35\x8f\x53\x66\x3c\xba\xee\x37\x6f\x34\xd0\xcc\x44\x89\x10\x33\x2f\x1a\x38\x7c\x34\x8b\x39\x89\x40\x99\x80\x3f\x52\xc9\xfc\x0c\x16\x21\x58\xdf\xe7\x9c\xba\xb6\x8b\x69\x5c\xc9\xb7\x1d\xa4\x46\xff\x88\xc4\xc5\xd0\x5b\xc8\xf5\x12\x76\x96\x76\x41\x7c\xf0\xc6\xc7\xad\x47\xeb\x34\x6d\xed\xc7\xc7\x6c\x0b\x67\xda\x68\x1d\xb1\x3c\xde\xef\xf4\xd9\x62\x7a\x0c\xcb\x9a\x1d\x07\xff\x15\xf1\x49\x56\x6c\x65\xdc\x2d\x05\x65\xb1\x53\x7b\xcc\xd0\x9b\x36\x8c\xb6\x62\x7b\xd8\x8d\x3b\x8d\x67\xa1\xb2\xb6\x29\xc3\xb4\xac\x48\x67\xa9\xf3\x51\x8e\x7b\x3d\x71\x35\x70\x62\x83\x6f\x95\xca\x69\x84\x69\xc6\x51\x8b\x79\x85\x92\xf7\x7a\x3d\x0a\x5f\x88\xa6\x0a\xdd\xb8\xf5\xb2\x8c\xe6\x12\x4e\x4e\x06\x57\x83\xe1\xc5\x68\x38\x19\x7f\x22\xc6\x5a\xad\x20\xbe\x16\x68\xa7\x9e\x18\xe5\xf9\x2a\xa7\xc1\xfe\x1d\x35\xc1\x8e\xe3\x3e\x5f\xc3\x84\x0c\xb3\x93\x2f\xca\x85\x46\x54\x17\x86\xfe\x7b\x7e\xee\xff\xda\x3b\x23\xee\xce\x4d\xde\xd4\x34\x20\x53\x5a\xa1\x2f\x75\x03\x70\x72\x4e\xcb\x24\x4f\xeb\x67\x5f\x3c\xd1\x15\x60\xda\x6e\x1b\x88\x84\x96\x4e\x10\x51\x32\x3a\xf1\x9d\xd7\x59\x76\xc2\x9f\x65\x3d\xcb\x9e\x4a\xdc\x12\x49\xb4\xc5\xe5\xfc\xa6\x93\xfd\x13\x00\x00\xff\xff\x89\xa6\x02\xdd\xb2\x07\x00\x00"

func dataAwsSimpleBuildBuildPhpShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildBuildPhpSh,
		"data/aws-simple/build/build-php.sh",
	)
}

func dataAwsSimpleBuildBuildPhpSh() (*asset, error) {
	bytes, err := dataAwsSimpleBuildBuildPhpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/build-php.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x5d\x6f\x9b\x30\x14\x7d\xe7\x57\x58\x48\xe9\x53\x80\x6c\xad\xa6\x69\xaf\xfb\x19\x55\x44\x0d\xb8\xe1\x2a\xb6\xb1\x7c\x4d\xa6\x16\xf9\xbf\xef\x9a\xef\x4c\x0d\xcd\x9a\x17\x23\x9f\xc3\x39\xc7\x87\xeb\x74\x11\xa3\x5f\xac\x40\xe7\x86\x97\x67\x61\xf3\x8b\xb0\x08\x8d\x8e\x7f\xb1\xf8\x90\xfe\x4c\x0f\xf1\x3e\x1a\x38\x17\x6e\x81\x17\x52\x20\x41\xc3\x6b\xb4\xc9\xff\x60\xce\xcb\x52\x20\xe6\x67\xf1\x46\x88\x6e\xa5\xdc\xaf\x51\x14\xa5\x15\xee\x16\x6a\xc5\x69\x30\xbb\x42\x50\xb6\x27\xca\xe3\xea\x11\xe8\xf7\xfd\x14\xc4\xd8\xe6\x02\x21\x23\x25\x25\xc2\xf3\xf8\x56\xb7\x63\xaf\x8d\x65\x15\x58\x06\x9a\x1e\x5b\x5d\x71\x47\xac\x9c\x76\x30\x2d\x5a\x90\x15\xdb\xf9\x89\x3c\xae\x24\xe7\xde\x8c\x08\xa7\xc5\x5a\x48\x19\xef\x17\x00\xb4\x04\x1d\xa0\xe7\x58\x9d\x83\x6c\x62\x58\xe6\x94\xc9\x1a\xe7\x9a\x6c\x31\x48\xba\x2e\x38\xcb\xa6\x31\xe9\x6f\xda\x75\xc2\x32\xef\xe3\xe3\xa8\xe4\xf7\xb7\x3d\x5f\x41\x8a\xb5\x25\x36\xad\x2d\x7b\x84\x34\x83\xa5\xf7\xd9\x1a\xaf\x04\x3a\xd0\xbd\x6b\x20\xfd\x47\x9a\x3b\xc2\x6c\x15\x50\x56\xf7\x1e\xdd\x7b\xf6\xf0\xc0\x0a\x8e\x35\x4b\x33\xc5\x41\xa7\x58\x7f\xd0\xc5\x8e\x09\x5d\x85\xef\xb5\xf5\x49\x36\xea\xd9\x31\x1a\xd4\x82\x42\x28\x52\xa0\x14\x2d\xd2\x39\x5f\xe6\xc1\x79\xa1\x33\x0f\x1e\x2b\xda\x3d\x4d\x26\xdc\x98\xd4\x9d\xde\xbf\x54\x18\x96\x16\x8c\x0b\x50\x3f\x6e\x89\xa9\x4d\x38\xfd\x24\xd5\xaf\xc7\x69\x8c\x7b\xca\x38\xc2\xf3\x7d\xd2\x5c\xf5\xd2\x21\xca\xac\x3c\x1b\x72\xc5\xdf\xa9\x74\x51\xe0\x82\x5d\xdd\xbe\x5b\xbd\x5c\x5f\xd3\xed\x72\xe2\xab\x1b\xbb\xa5\xb8\x10\x3f\x51\x9c\x6f\xf9\x96\xda\x40\xfa\x2c\x5b\x3f\x01\x39\x57\x30\xf4\x01\xc9\xf7\x6f\x3f\x1e\x0f\xd5\xd3\xd3\xc2\x01\x8d\x8e\x6b\x62\x4d\xb5\x95\x8f\xa9\xe4\xf6\x24\x56\x32\x58\xe7\xc1\x79\xaa\xbb\x2d\x68\x76\xdb\x55\xa9\x0a\xf2\x09\xeb\xba\xf0\x44\x63\xfd\x6f\x76\x5a\x69\x88\xb8\x32\x1f\x25\x1e\xfe\xb2\x8e\x51\xe4\xa3\xbf\x01\x00\x00\xff\xff\x4f\xd6\x14\x86\x64\x05\x00\x00"

func dataAwsSimpleBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildTemplateJsonTpl,
		"data/aws-simple/build/template.json.tpl",
	)
}

func dataAwsSimpleBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsSimpleBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x52\xc1\x8e\xdb\x20\x10\xbd\xe7\x2b\x46\x74\x8f\x5d\x27\xed\xb1\x52\xcf\xbd\xb5\x1f\x50\xad\x10\xc6\x24\x45\x8b\x01\xc1\x90\xca\xb2\xfc\xef\x1d\xa0\x16\xc1\xbb\xbd\x36\x39\xf9\xcd\x63\xde\xcc\x9b\xf7\x01\xbe\x29\xab\x82\x40\x35\xc1\xb8\xc0\x0f\x44\xf7\x11\x26\x07\xd6\x21\xa8\x49\x23\xcc\xc2\x26\x61\xcc\x72\x3a\xdd\x45\xd0\x62\x34\x0a\x98\xb6\xd7\x20\xb8\x9e\x18\xac\xdb\x03\x2c\x7e\x47\x2e\xa4\x54\x31\xf2\x57\xb5\xbc\x53\x8c\x4a\x06\x85\xff\x28\x06\x75\xd3\xce\x1e\x0a\x44\xe5\x56\xcc\xaa\xc0\x8f\x0f\x66\x7d\x60\x6a\x1b\x51\x58\xa9\x38\x2e\x3e\xd3\x61\x52\x57\x91\x0c\xc2\x57\x60\xf8\x79\x98\xb5\x0c\x8e\xc1\xe3\x8b\x98\x46\x4b\xd3\xf8\x34\x1a\x2d\x0f\xdd\xee\x5e\x72\xa9\xa7\xf0\x0e\xfc\x77\xed\x93\x0f\xee\xae\x27\x15\xca\xf4\x04\x9d\x00\xda\xf2\x59\xf5\x69\xa5\x87\x43\x6f\xca\xc6\x88\xd6\x6c\xe8\x69\x0d\x2f\xb4\x6a\x08\xe4\x5f\x47\xab\x38\x51\x68\x88\xa0\xa2\x4b\x41\x36\x7f\x53\xd0\xb8\xf0\x5b\x70\xc9\x33\x02\xbd\xaf\x93\x65\x0f\x6b\x9f\x75\xad\x1f\xdb\xf6\x5c\x5b\xee\xc7\x2c\x9a\x75\xc1\xa6\x57\xbf\xa9\x44\x35\x6d\x6f\x24\x17\x4b\x3f\x00\x5a\x1f\x9d\x74\xa6\x8e\xf7\xfc\xa9\x80\xd7\xe0\x66\xee\x5d\xc0\x02\x5e\x0a\x86\x6e\x47\x1a\x96\xad\xe5\xa3\x71\xf2\x35\x12\xf6\x93\x5d\x86\xf2\x3f\x5f\xd8\x0b\xd5\xb7\xac\xa6\xfe\x9b\xd8\x1b\x1b\xf7\x28\x3d\x1a\x48\x81\x83\xf6\x6b\xf7\x98\x75\xf1\xad\x4b\x5f\x2b\x77\x70\xbd\x7d\x0d\x1d\x79\xdc\xf5\xe9\xb2\x58\x88\x7b\xf2\x0f\x82\x3b\x5c\x4f\x92\xcf\xd3\x1f\x9d\x3a\xd7\x2d\x9f\xd6\xb7\x89\x18\x68\x9d\x21\x9f\xf3\x25\x3f\x46\x71\x23\x7f\xe1\x7b\x16\xe9\x82\xc1\xaa\x29\x2e\xa1\x4f\x08\x2c\x05\x53\x3d\xb8\x0b\x93\x0a\xf5\x17\xa2\xff\x72\x3e\x57\x89\x7d\xc7\xd2\xbc\x2e\xc0\x27\x1b\xb7\x73\x0e\xe8\x9f\x00\x00\x00\xff\xff\x5f\x73\x79\x4b\x5f\x04\x00\x00"

func dataAwsSimpleDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleDeployMainTfTpl,
		"data/aws-simple/deploy/main.tf.tpl",
	)
}

func dataAwsSimpleDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsSimpleDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x5f\x53\xd4\x48\x10\x7f\xcf\xa7\x68\x03\x8a\x56\x91\x44\x29\xf5\x01\x85\x3a\x44\x3c\xa8\xf2\xc0\x62\xf1\x5e\xee\xae\xd6\xd9\xa4\x93\xcc\x91\xcc\x8c\x33\x93\xd5\xdc\x92\xef\x7e\xdd\x93\x2c\x0b\x6a\x79\x57\x05\xc5\xa4\xa7\xff\xfe\xba\xfb\x37\x6c\xc1\xaf\xa8\xd0\x0a\x8f\x05\x2c\x7a\xb8\xf0\x5e\xef\x42\xa1\x41\x69\x0f\x58\x48\xff\x20\xda\x8a\xb6\xe0\xaa\x96\x0e\xe8\xc7\xd7\x08\xbf\x8b\xca\x0a\xe5\x4b\xd9\x20\x54\xdf\xda\x42\xa9\x6d\xd0\x2a\x70\x89\x8d\x36\x2d\x2a\x0f\xba\x24\x17\x9e\x5d\x08\x63\x1a\x99\x0b\x2f\xb5\xca\x1c\xda\xa5\xcc\x31\x85\x33\x0f\xae\xd6\x5d\x53\x84\xa0\x0b\x84\x5a\xa8\x22\xe1\xe0\x58\xa4\x70\xa5\xa1\xd5\x85\x2c\x7b\x76\x4b\x7e\xee\x84\xdf\x85\xce\x61\x88\x76\x64\x0c\x0b\xd2\x28\x9a\xae\xd3\x5c\xab\x52\x56\x9d\xc5\xc7\xf1\x5e\xfc\x84\x2b\xba\x19\x45\x37\x11\xc0\x78\x4a\x97\x6d\xba\xd0\x5f\xe1\x00\xe2\x5a\xb8\x5a\xe6\xda\x9a\xcc\x58\xcc\xa5\xc3\x97\xcf\xe3\x88\x14\xb7\xe0\x54\x3b\x2a\x40\x35\x3d\x28\xf4\x5f\xb4\xbd\xbe\x67\x3e\xc9\x20\x36\x56\x2e\x09\x87\xf9\x24\x88\x77\x41\x9a\x7d\x88\x57\x2b\x06\x62\x2e\xcd\x5c\x14\x85\x45\xe7\x60\x18\x26\xc7\x33\xf4\x9d\x01\x01\xae\x57\x39\xe1\x57\xea\xa6\x40\x0b\xa5\xd5\x2d\xe8\xce\x02\x7b\x91\xaa\x82\x42\x52\x42\x5e\x5b\x2a\x5f\x43\xb6\x1c\xab\xbb\x97\xc3\xe8\x60\x3e\x39\xd8\xa1\x90\x46\xf8\x3a\x5d\x3b\x18\x86\x9d\x5d\x88\xd7\x96\xf1\x2e\xd9\x02\xe8\x2f\xd4\x37\xca\xef\x56\x0a\x95\xd5\x9d\xb9\x23\x19\x93\x3c\x51\x62\x41\x6d\x9e\xcd\x4e\x41\x54\xdc\x4a\x6a\xef\x17\x61\x0b\x76\xec\x34\xb5\xdf\x7b\x3e\x4e\xd5\x53\xad\x06\x55\x81\x2a\x97\xe8\x42\x05\x6e\x93\xa9\x73\x75\x3a\x59\xcf\x47\x5f\x07\xe0\x6d\x87\x63\xa0\x77\xba\x53\x45\x98\x0b\x58\x77\x6e\xfc\x7a\x2c\x4b\x10\xaa\x7f\x42\x5a\xab\x87\x61\xba\x08\x11\x90\x8a\x8e\x6b\x8b\x39\x49\x5c\x4a\x38\xc3\xc3\x81\xd4\xf8\x9e\x5a\x9a\x69\x1a\xc7\x6c\xa3\x95\x10\x30\x64\xde\x68\x6d\xd2\x63\x92\x7a\x02\x8b\x9b\xf1\x73\x28\xd9\x59\x40\x90\x0e\xf7\x54\x8d\xd5\x4b\xe9\x38\xc3\xd8\xd5\xd8\x34\xdc\x71\xd5\x48\x85\x84\x61\x5e\xc0\xd6\x8a\x0c\x06\x78\xf4\x08\x16\x34\x5a\xd3\x67\xd6\x0a\xa9\x52\x57\xc7\x63\x31\x04\x15\xd7\x43\x49\x07\x08\xde\x6b\x51\x80\x68\x9a\xd0\xfe\xd2\x8a\x8a\x77\xc7\x41\x8d\x16\x43\xdd\x84\xc2\x3d\x80\xd3\x0d\x24\x6b\x6d\xc6\x85\xe7\x6d\x63\x1d\x10\xe1\xca\x27\xc9\x8d\x45\x8a\x32\x0c\x3f\xcc\xe0\x4c\x39\xcf\x09\x2c\x3a\x49\xcb\x88\x6a\x29\xad\x56\x6c\xf5\x7f\x2b\xdf\x76\xb9\x95\xc6\xcf\x69\xcd\xa3\xef\x6d\x18\xd3\x7d\x23\x2c\x85\xc0\xc6\x85\x95\x34\xbb\xa0\x6f\xc6\x89\xbc\xb3\x8d\xb7\x3a\x59\xb7\xa0\x46\x75\xc9\xb3\xbd\xf4\xe9\x73\x46\x8d\x12\x8e\xf8\x37\xba\x13\x89\x2c\x5e\xbf\x9e\x1d\x5f\x9e\x7d\xb8\x8a\xb6\x1e\x64\x0b\xa9\x32\xc6\x3c\x8a\x1c\x7a\x48\x98\xc9\x3a\x35\x1d\xd1\x5a\xfc\x2a\xc3\xd1\x48\x83\xa5\x90\xcd\x24\xf6\x56\xe4\x34\x8a\x74\xd2\xf6\xf1\x13\x58\x71\x4e\x8d\xce\x45\x43\x63\xde\xd9\x1c\x99\x5d\x0e\xb6\x9f\x6d\xc4\x5c\xb0\xd2\x07\xdb\x7b\x2c\xc2\xbc\xd6\x10\x9f\x5c\x5e\x5e\x5c\x82\xf0\xb0\xbd\xda\x18\x0d\xfb\xdb\xab\x51\x77\x78\x05\xef\x05\x91\x49\xa3\x2b\xb7\xcf\xc5\xd0\xce\xa1\x01\x1e\x53\x5e\x6c\x9b\xd1\x45\xe6\x7a\x47\x7f\xe0\x06\x7c\xc8\x4d\xc1\xde\xd3\x68\x88\x28\x3b\x03\x3b\x21\x39\x88\xb7\x57\x6f\x8e\x66\xa7\xf3\xd9\xc5\xc7\xcb\xe3\x93\x21\x66\xc1\xfb\xb3\xf3\x93\xf3\x8b\x21\xde\x01\xca\x21\x8a\x34\x72\x09\x74\xf1\x4b\x0c\x7b\x87\x8f\x9e\x91\x3b\x72\x5a\x11\xfc\x89\x1f\xe3\x1d\x42\x46\x73\x92\xa9\xae\x69\x5e\xc1\x10\xe9\x26\x18\x8c\x65\xfc\xc1\x1a\x7f\x01\x19\xf3\x15\x31\xee\x6f\xe2\x1a\x81\x50\xa3\x85\xf7\x35\x95\xf7\x69\xe2\x08\xa0\x95\xfe\x04\x95\xa6\x5d\x1f\x59\xaa\x09\x24\xc5\x7c\x4c\x54\xca\x82\xb0\x35\xa3\x57\xda\x89\x5b\x0e\x82\x43\x8a\x5f\xeb\x16\xd7\x92\x2c\xe5\x8e\xd9\x9c\xa3\x1d\x4f\xeb\xcf\xbc\xc2\xbc\x13\xe6\x9b\x70\xa3\xec\xa9\x08\xa9\x22\x22\x84\x07\x23\x74\xf1\x47\x87\x6f\xcf\x67\xd4\xe1\x18\x32\xf4\x79\x46\x09\xf1\x6f\x31\x1f\x27\x0f\x0e\xef\x54\x49\x69\xa9\x68\xdd\xaa\x3b\x86\x37\xe0\x3a\x1a\x45\x8f\x08\x89\xf8\x2f\x37\xe4\x40\xe3\x68\x30\x3d\x5f\x0c\x02\x10\xb3\x7b\x61\x7d\x54\x4a\x02\xbf\x81\xf8\xa8\x08\x0c\x29\x8c\xa7\x2b\xa3\x9d\x24\x02\x67\x4a\xa4\x77\x0d\x3a\xc3\x7c\xa4\xaa\x34\x4d\xe3\x08\xbf\x1a\x6d\x3d\xbc\x3d\x79\x73\x76\x74\x3e\x7f\x77\x79\x71\x7e\x75\x72\xfe\xf6\x40\x69\x25\x99\xa0\x44\xee\xe5\x12\xa3\x75\x48\xf2\x97\x10\xe5\x8e\x2e\x28\xdd\xfe\xbb\x1b\x39\x6d\x70\xd2\x83\xe9\x7d\x4d\xac\xe7\x74\xe9\x89\x73\x31\xa1\x15\x34\x68\x3d\xa7\xf1\x03\x59\x92\xeb\xb6\xa5\x95\x66\x47\x34\x6d\xca\x71\x5e\x49\xed\xbd\x71\x9b\x20\x45\x91\xf0\xfd\x6d\x49\x7d\x88\x63\xc4\xbe\x56\xf4\xb8\xfd\x9d\x99\xda\xbc\x48\x5e\xa4\x2f\x23\x7e\xdd\xb0\x75\x3c\x0d\xf4\x9c\x5b\xfc\xdc\xd1\x7c\x14\xfc\x49\xcf\xeb\x92\x89\x2a\xee\x94\xe8\xb8\x25\x9e\xff\x1d\xa0\x3b\x23\xf2\x6b\x7a\x16\x5c\x4c\xc6\x61\xd0\x1d\xf1\x20\xff\xeb\xf0\x6d\x6d\xe9\xbd\xa2\xaf\xb1\x9f\xe0\xf8\x09\x4a\xa1\x27\x13\xb9\x71\x5f\x3e\x9c\x7e\x08\xad\x70\x9d\xe1\x32\xc3\x0b\x36\x85\x0f\x5d\xf9\x19\xaa\x54\x22\xfc\x49\x63\xb0\xf8\xc7\x42\x45\x4b\xd1\xa2\xcd\x69\x58\xc5\xc4\x9a\x09\xbd\xf1\x5c\x14\x7d\xb3\x16\x5d\x8d\x87\x00\x4d\x9b\xdb\x9e\x46\x62\x3c\xf7\xee\x73\x33\x1e\x4b\xd3\x8e\x87\xaa\x18\xff\x32\x47\x33\x67\x8c\x5f\xa6\x22\xcd\xef\x6a\x38\xd6\x2d\x75\x01\x6d\x48\x98\x97\xcb\xb7\x26\x0a\xe1\x12\x37\x83\xd0\xb8\xfd\x2c\xa3\xec\xf3\xb5\xa2\xb6\x55\x36\x55\x42\xbb\x74\xc3\xce\x6f\x2b\x6d\x97\x70\xab\x67\x6a\x61\x21\xeb\x1c\x13\x12\x31\x5d\x20\xd4\xf5\x65\x34\xf1\xec\xbf\x01\x00\x00\xff\xff\x74\x5a\xdc\x2e\x36\x0a\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/aws-simple/build/build-php.sh": dataAwsSimpleBuildBuildPhpSh,
	"data/aws-simple/build/template.json.tpl": dataAwsSimpleBuildTemplateJsonTpl,
	"data/aws-simple/deploy/main.tf.tpl": dataAwsSimpleDeployMainTfTpl,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-simple": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"build-php.sh": &bintree{dataAwsSimpleBuildBuildPhpSh, map[string]*bintree{
				}},
				"template.json.tpl": &bintree{dataAwsSimpleBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsSimpleDeployMainTfTpl, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
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

