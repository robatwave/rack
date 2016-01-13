// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// models/templates/service/mysql.tmpl
// models/templates/service/papertrail.tmpl
// models/templates/service/postgres.tmpl
// models/templates/service/redis.tmpl
// models/templates/service/syslog.tmpl
// models/templates/service/webhook.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\x7b\x6f\xdb\xba\x15\xff\x3f\x9f\x82\x20\xfa\xc7\x56\xa4\x6e\xd2\xe2\x76\x9b\xb1\x0d\x48\x9d\xb4\xf5\x9a\xa4\x99\xed\xf6\x62\x2b\x82\x0b\x45\xa6\x6d\xcd\xb2\xa8\x4b\x52\xb9\x75\x03\x7f\xf7\x1d\x3e\x24\xf1\x25\xdb\x4d\xd2\x62\xbb\xd8\xd6\xc5\xe2\x21\x79\x78\x9e\xbf\x43\x9e\xbb\x3b\x34\x25\xb3\xac\x20\x08\x27\x65\x89\xd1\x66\x73\x80\xd0\x1d\xfc\x0f\x21\x7c\xf2\xf3\x78\x42\x56\x65\x9e\x08\xf2\x86\xb2\x55\x22\x3e\x11\xc6\x33\x5a\x60\xd4\x47\xf8\xc5\xd1\xf1\xd1\xb3\xa3\xbf\xc0\x7f\xf1\xa1\x26\x1f\xd0\x62\x9a\x09\x18\xe7\xb8\x6f\x96\x80\xa5\xee\x90\x30\x6b\x20\x7c\x93\xe4\x49\x91\x12\xf6\x2c\x6d\x49\x51\x4f\xef\x19\x10\x97\x8c\xa6\x84\xf3\x2e\x5a\xfc\x1a\xd6\x5a\x0e\xf2\x8a\x0b\xc2\xe4\x86\x08\xbf\x29\xfa\xfd\xb3\x5f\xab\x24\x97\x0c\x7c\x96\x5f\x46\x64\x06\x7f\xe2\x9a\x0a\x6d\x0e\x11\xc6\xe8\x1a\xe9\x45\x36\x86\xf1\xab\x84\x25\x2b\x02\x04\x5c\x9e\x6c\x3b\xe7\xa5\xa4\xdd\x83\x6b\x87\xae\x66\xd9\xe2\xd6\x7c\x82\x8f\x93\x75\x49\x94\x44\xc7\x82\x65\xc5\xdc\x48\x53\x0d\x9d\x92\x59\x52\xe5\x42\x8d\xba\xdf\x79\xca\xb2\x52\xd4\xba\xc0\x66\x68\x73\xd8\xec\x54\x56\x91\x5d\x80\xf4\xb2\x5a\xdd\x00\x07\x91\x4d\x94\x4e\x8f\xba\xb6\x91\x52\xbc\xfa\x88\xf8\x22\x61\x84\x23\x3a\x43\x24\x49\x17\xc8\x9c\x36\xdc\xff\xac\xb8\xcd\x18\x2d\x56\xa4\x10\x71\x3e\xba\x0f\xbb\xe5\xac\xd1\xa3\xbe\x27\xeb\xef\xbd\xc5\x88\xe4\x24\xe1\xe4\x07\xe8\x6d\x44\x4a\xca\x33\x41\x59\xec\x4c\x0f\xdb\x6c\x4c\x2b\x96\x12\x94\xd2\x29\x41\xac\xdd\x26\x60\x61\x5c\xdd\x14\x44\xf0\x8e\xfd\xcf\x33\x2e\xfe\x0a\x81\x01\x3c\x6d\xf0\xa2\xdf\xd7\xc4\xfd\xfe\x70\xfa\xf7\xfb\xf0\xf4\xe9\x6a\x80\xb8\xde\x0f\xcd\x28\x43\x62\x91\x71\x24\xe3\x50\xc0\x55\x1d\x7a\x1c\xae\x3c\xc5\xfd\x61\x74\xf6\xcf\x8f\xc3\xd1\xd9\xe9\x1f\xd1\x79\xb2\xba\x99\x26\x68\x00\xfe\x46\x57\x13\x5a\x66\x29\x7a\x97\x14\xd3\x9c\x30\x64\x94\x89\xea\x15\x2d\xf6\x2e\xb2\xe2\x9c\x14\x73\xb1\x50\xcc\x1d\xdb\x43\x9e\x51\x85\xfc\x5d\x0d\x3a\x24\xd6\x0a\x0b\x68\xa4\xa4\xee\x2b\xa8\xa8\x80\x9c\x40\x36\x22\x5c\x29\xd9\x56\x1e\xb6\x64\x30\xa2\x79\xcc\x88\x6b\x26\x87\x27\x17\xfd\xbe\xa2\xb1\x38\xb9\x62\xb4\x24\x4c\x64\xc4\xb5\x08\x99\x1e\x38\xaf\x56\x44\xd2\x5f\xd1\x3c\x4b\xd7\xa7\x34\xad\x02\x87\xf7\x94\x27\xd3\xc6\x8b\x67\x90\x39\x8e\xff\x64\x6d\xa2\xed\x4e\x40\x00\x35\xf3\x3f\x3b\x43\xc8\x5b\x4f\x91\x9f\xcd\x66\x24\x55\xae\x7c\x92\xe7\xf4\x37\x6f\x35\xc3\x7a\x56\xa4\x59\x99\xe4\x3a\x3d\x8c\x09\xbb\xcd\x52\xa2\x72\x03\xce\x95\x81\xf4\x92\x55\xf2\x95\x16\xc9\x6f\xbc\x97\xd2\x95\xca\x0c\x91\x75\x4e\x52\x63\x61\x30\x8f\x0b\xde\x6f\x0f\x0e\x33\x3c\xf2\x8d\xf3\xdb\x1e\x75\x56\x86\x9c\x03\x56\x06\xcc\x3f\xc7\xee\x67\x29\x49\x2d\x6b\x57\x06\xbe\x04\x34\xe5\xfa\x12\x12\x97\x92\xc1\x74\x95\x15\xe0\x9a\x2c\x01\xa7\x0e\x64\x81\x77\x28\x48\xd1\xec\xa3\x24\x45\xe8\x28\x4a\x0a\x36\x50\x85\x25\x32\xfc\x54\xfe\xac\x0d\x53\x7f\x40\x9b\x1d\x62\xb3\x7f\xb5\x94\x9b\x30\xcb\xb5\xa6\xbd\xc5\xac\x75\x2c\xe8\xf7\xdf\x54\x85\xe6\x6a\x2f\xeb\x1e\x40\x9c\x0c\x2d\x79\xfc\xf2\x75\x95\x2e\x89\x68\x01\xc7\x3f\x68\x66\x4c\xe3\x19\x9c\x14\xfe\x0f\xf0\xca\x2d\xfd\x02\x7f\xb7\xf8\x43\xb1\x31\x22\x73\xe5\xcd\x70\xf8\xd0\xce\x60\x61\x93\xc7\xfc\x55\xf5\xa2\x4c\x07\xad\xe7\xce\xb2\x0d\x1c\x93\xb0\xe6\xf9\x4c\x41\x34\xf8\xdd\xfb\x9a\x95\x58\x6f\xd2\x69\x7e\x26\x18\xca\x55\x8c\x27\x90\x2f\x80\x4c\x0a\xf0\x15\x87\xee\x82\xac\x20\x49\x8c\xb3\xaf\x4a\x9c\xc7\x2f\xfe\xec\x0e\xd7\x01\x45\x33\xfd\x96\x88\x13\xa1\xad\x22\x88\x3a\xd2\x26\x58\x11\x78\x18\x1e\x55\x85\xc8\xb4\x0d\x17\x20\xf1\xff\x70\x77\x83\x09\x8c\xd1\x4a\xd9\xd6\xcb\x23\xdc\x6d\x0a\x71\xb0\xc6\x9a\x78\xb8\x13\xaf\x7d\x03\x29\xd7\x71\xc4\x03\x77\x0e\x29\x27\x69\xc5\x32\xb1\xc6\x1d\x4b\x71\xe9\x42\xcd\x60\x1d\xbe\x3f\x54\xa2\xac\xc4\x6e\xf4\x4c\x0d\xdd\x4e\x4e\x5d\xc2\x26\xc1\x13\x21\x20\x87\x79\x19\xfe\x53\x92\x57\x46\x97\xc6\xbc\x1a\xba\x56\xdc\x07\xf5\xbf\x9b\x03\xd8\x90\x14\x53\xb5\xae\x55\x43\xc4\x40\xbb\x2e\x29\xee\x10\x4b\x8a\x39\x41\x4f\x96\xa8\xff\x37\xd4\x3b\x2b\x04\x53\xd1\x8b\xd7\x67\xd0\x80\x1e\xe8\xaa\x12\x5c\x52\xd2\x6d\x36\x6d\xc8\x0e\xe1\xbd\xf4\x0c\xdc\x18\x54\xcb\xce\x76\xc6\x6a\x5c\xee\x31\x45\x14\x53\x0d\x2b\x2d\x17\xa4\x27\x99\x84\x01\xed\x09\x76\x62\xed\x80\xd4\x0e\xa0\xfe\xe9\x95\xfd\xdd\x41\x2b\x17\xaf\x25\x8c\x1e\x9d\x5c\x20\x41\x01\x90\x49\xab\x22\xd8\xb1\x87\x08\x17\xb0\x46\xc6\xc8\x74\x40\x2b\x27\x8c\xb7\x68\xa3\x8b\x19\x17\xd0\x04\xf8\x62\xb2\x20\xa8\x50\x53\x25\x53\x59\x01\x06\x0a\x7e\xa9\xc2\x89\x42\xfb\x02\xc6\x8d\x08\x25\xbb\x60\x67\x00\x25\x21\x86\xa0\x25\x21\x25\x62\x55\x51\x80\xa5\x20\x5a\xa0\x35\x78\x11\x4a\x4d\x9d\xd3\x9e\x66\x5f\xf5\x34\x16\xab\x34\x81\xdf\xc3\x20\xcf\x5a\x4b\x8d\x58\x69\x4d\xa2\xcd\x52\xee\x85\xcf\xe9\xfc\x2d\xa3\x55\xb9\x6d\x5a\x43\x53\xcf\xdb\xce\x97\x15\x1e\x3a\x38\x73\x52\x8e\x19\x05\x7c\x2c\x18\x49\x56\x75\x85\x1c\xcd\x36\x78\x0c\x35\x55\xa3\xd1\xe3\xd6\xc5\x3a\xce\xe2\xe6\x36\x3a\xe7\xea\x5f\x4d\xb4\xcf\x51\xda\xf0\x55\xfb\xc0\x93\x55\x52\x64\x33\xc2\x85\xe5\x04\xf0\x39\x9b\xa1\xde\xbb\x84\x5f\xe9\x79\xad\xa3\x76\x79\x4d\xdc\x62\xcf\x06\xe3\x49\xc2\x97\xa7\x92\x8b\x4c\x44\x20\x7c\x09\xbc\xf2\x0f\x2a\xdb\x39\x09\xfd\xb0\x41\x6c\x2a\x81\x5c\x47\xc0\xb8\x26\x97\xe8\xda\xdf\xc3\x22\xb6\x70\xcd\x71\xef\x68\xbf\xe4\x6f\x36\x9e\xd0\x25\x29\x76\xe6\xb7\xce\xdc\x66\xc0\x59\x07\x50\xf0\xe0\x01\xa0\xaa\x74\xa9\x66\xa8\x74\x2e\xd5\xd2\xc8\x10\x87\x90\xc1\xae\x49\x9b\x85\xea\x6f\x1e\xa9\x57\x8c\x37\xe4\xf6\x77\x6f\x4a\x03\x46\x6a\x2f\x23\x6b\x9f\x44\x4a\xdc\x5c\xb4\xd8\xb9\x87\xc8\xc0\xfe\x8b\x80\x41\x9d\x76\xb6\x02\xb8\xb8\xbd\x58\x51\xbf\xc3\x50\x2c\x3e\xb4\xa1\xc2\x7c\x30\xd5\xd7\x26\x43\xb6\x16\xd9\xee\xd3\x18\x79\xaf\xa6\xaa\x41\xa9\xda\xd8\x12\xf6\xa1\xbb\xba\x71\x25\x6b\x35\x5b\xf7\xed\xf7\x1d\x06\x5a\x1f\xea\x7f\xc2\x32\x9d\x6b\xb3\xf0\x92\xcc\x26\xf5\x33\x4e\x43\xbf\x2b\x3b\xfd\x48\x67\x08\x23\xcc\x36\x36\xc3\x70\xe1\xfb\x56\x03\x6b\x1b\x28\xd4\x06\x22\x8f\xf6\x9c\x26\xd3\xda\xa2\x22\x65\x5b\x1b\x2b\x7b\x57\x94\x89\x0b\xa8\xdd\x25\xa6\xf2\x2d\x14\x85\x82\xe9\x7b\x82\xf9\x26\x1b\x8e\x88\x4d\x7f\xe8\x0d\x68\x21\x12\x48\x08\x2c\x2e\xca\xb8\xc9\x23\x5f\x41\x97\x54\xa7\xd4\xfd\x4a\x35\x7b\xc9\x5d\x48\x20\x7a\x47\xec\xa2\xb5\xe6\xf4\x76\x4a\x7a\x52\x4f\x74\x72\x52\x3b\x27\x2e\x7d\x1f\x74\xd6\x8b\xf4\x4c\xd6\x33\xc2\x93\x93\xa5\xf0\xac\x00\x33\x90\xfe\x39\xcb\x52\x09\xe4\xb7\xdf\x3e\x3f\x74\xf5\xf6\xd6\xba\xbd\xa3\x93\xf5\x05\xb9\xcf\xc2\xdf\x81\x57\xcd\x8c\x66\xf3\x5f\x12\x57\x5c\x3f\x48\xef\x71\x84\xfe\x98\x3a\xbf\xcf\x21\xeb\x3f\xef\x75\x2d\xeb\xeb\xe0\x9b\xae\x84\x1f\xc1\x38\xbf\x99\xe1\xef\xcf\xe1\x3b\xca\x63\x2f\x03\x7b\xc9\x32\x2b\xa6\xe4\x8b\xb5\xd9\x08\x0a\x11\xba\xe2\x3f\x58\xc6\xc6\xec\xef\x73\x86\x4b\xba\x95\x3d\x6b\x4c\x5d\xa8\x91\xa9\x0a\xb6\xa6\xec\x95\x2e\x76\xa8\xd6\x68\x02\xee\x43\xdc\xcd\xad\xb8\xee\xe3\x6f\xb8\x26\x6d\xa4\xe8\x09\xcf\xd3\xb5\x5d\x8d\xf9\xe0\x45\xca\xd8\xc9\x69\x3a\x71\x9d\x5e\x8e\x35\x1a\xf0\x9e\xee\x7e\xa0\xab\x87\x25\xe4\x83\xd6\x7e\x94\x14\xe9\x57\xa5\xf7\xd2\x5e\x44\xe4\x63\x73\x7d\xe5\x56\x9e\x7e\xed\xa9\x1f\x80\x1c\xd2\xc6\xa1\x3a\x00\x2c\x56\x64\xee\x45\x48\x00\x7c\xd0\x1e\x88\xf0\xd9\x4d\x23\x48\x0f\xc2\x60\x87\xa3\x61\x31\x67\xf2\x81\x32\xa8\x19\xb6\x9a\x8d\xa1\xf2\x10\x1a\x1e\x64\x53\x36\x94\xf2\xc0\x47\x3d\xf5\x9f\xe7\x47\xe1\x55\xfb\xb0\x84\xb3\x0b\x9a\x52\xf9\xea\x80\x45\x5a\x86\x24\x6f\x18\x5d\xc9\x8d\x1f\xd3\x9a\x82\x4d\x26\xf4\xb1\xb7\x70\x76\xd8\xec\xaa\x92\x76\x03\x46\xbb\x60\xfa\x54\xa6\xc3\xa9\xc3\xac\x7c\x50\xf3\x2f\x1e\x0f\x3b\x2d\xb6\xdb\x48\xf3\x84\x8b\x2c\x6d\x51\x3a\x68\x5a\x5e\x98\xb4\xa0\x7d\xb7\xd1\x5a\x2f\xa2\x6d\x45\x60\xbe\xd9\x82\xd0\xa5\x28\xf9\x15\xf5\xc6\xe9\x82\x00\x5f\x38\x2b\xcc\xa5\xba\x57\x41\xea\x71\xa3\x93\x9a\xda\x49\x5a\xa1\x48\x65\x7f\x45\x41\xd4\x0b\xc6\x29\x03\x04\x0f\x07\xd1\xcf\x3a\x9a\xaf\xb3\x22\xb9\xc9\x89\x94\xa1\x60\x15\x39\xb4\xaf\xce\x5f\x1d\x39\x0e\xd2\xae\x63\xdf\x04\x83\xe5\x4e\x73\xd2\x4e\x7a\xf9\xea\xc8\x9b\xc6\x28\xe7\xff\xa6\x05\xa9\xb7\x68\x87\xde\x91\x24\x17\x8b\xc1\x82\xa4\x4b\xbf\x5e\xd5\x43\xeb\xc9\x02\xdc\x70\x41\xf3\xa9\xba\x15\x75\xef\xf8\x87\x52\x48\xb7\xea\x95\xee\x27\xaf\xaa\x63\xf3\xf8\x53\x8b\xae\x8d\xf0\x64\x70\xe5\x3c\x88\x74\x25\xa0\xda\xae\xdf\x64\x8c\x0b\xf9\xa3\x4e\x4a\xb1\xc7\x18\x4b\x70\x2f\x9d\xef\x1f\x8b\x45\xf4\x30\x6d\xd5\x63\x89\x44\xbe\x93\x93\x22\x2c\x09\xf7\x0c\x3d\xfa\xc0\xc3\x99\x3e\xee\xa3\x14\x29\x7e\x84\x88\x3c\x01\xda\x91\x4b\xc9\x36\x24\xb1\x3d\xe7\x7b\xc7\x30\x65\x1b\xf2\x26\x3a\x25\x7b\x70\xd6\x90\x3e\x06\x57\xc6\x3a\xbc\x5d\x02\x0e\x3d\x3d\x3d\xa0\x2c\x0b\x0f\x14\xea\xc7\xd3\xd0\x78\x7c\x1e\x99\xf6\xc3\x75\xd4\xa1\xa5\x4e\xee\xbe\x87\x9e\xa2\x1b\x01\x07\x96\xfd\x7b\xe9\xe5\xe1\x45\x79\xb0\x65\x84\x8b\xff\x73\x1d\xc6\x3d\xed\x77\xad\x43\xef\xcb\xb5\xdf\xd0\xf1\xd8\xc0\xe7\xfc\xf5\x80\xd2\x65\x46\xc6\x00\x52\x96\xf2\xd5\x88\x37\x49\xfd\xf3\x9d\xdf\xe5\x91\xcc\xd4\x8d\xa5\xbc\x8a\x77\xd6\xb0\x6c\xa5\xbe\x66\x85\x53\xfb\x9f\x81\xcd\x2e\x98\x1c\xb9\x06\xda\x5e\x13\x58\x2d\x1c\x16\x36\xdb\x55\xb8\x78\xcf\xe2\xce\xb5\x6a\xfc\x7d\xcb\x6f\x49\xea\x78\x42\xdb\xab\x15\xa9\xb3\xc3\xc8\xeb\xf1\x68\xdb\x7d\x7c\x4c\xed\x35\xff\xb8\x96\xe1\x35\x74\x6c\xef\x4f\x72\x7b\x93\xfc\x7d\xac\x4e\xa5\x30\xcf\x91\x94\x7b\x6d\x4b\xbe\xcd\xba\x40\xfd\x20\xf6\xf7\x75\xfc\x19\xc2\x6a\xfc\x09\xee\x48\xc2\x8e\xa5\x78\xb7\x92\x03\xff\x5c\xfb\xb5\xf5\x1d\xb6\x3d\x6d\x69\x20\x7b\xfc\xde\xb0\x4e\x25\xab\x51\xa2\xab\x86\x1c\x3c\xe8\xa6\xa9\x1a\x74\xd9\x7a\x43\x9e\xc6\x63\x62\xc7\x1c\x46\xe6\x12\x0a\xb2\x3a\x64\x72\x59\xfd\x45\xeb\x8f\xdd\xab\x8d\xfc\xb5\x7e\xce\xc4\x62\x8f\xb5\xd2\x17\x3b\x99\x07\x92\x93\x4a\x2c\x28\xcb\xbe\x92\x68\x05\x1d\xcc\xba\x8e\x48\xd5\xea\xf3\x8a\xca\xf5\x69\x64\x19\x1f\x63\x1d\x74\x8d\xb6\x23\xf5\x5f\x7a\x74\xeb\x03\xbe\xfd\x1c\x29\xe3\xd2\x36\x60\xdd\x6b\x42\x5d\xfb\x54\x67\x82\x52\x6d\xc1\x32\x30\xba\xef\x84\x78\xb8\x4a\xe6\xcd\x98\xfa\x61\x0d\x0e\xe8\x6a\x95\x14\xd3\x7a\xd8\xfc\xd4\xb7\x83\x16\x59\xdb\x5c\x12\xa9\x65\xdc\x06\x94\xa6\x0a\x8e\xb7\x58\xb7\x0f\x4e\x4b\xb2\x3e\x44\x4f\x6e\x65\xf2\x31\x1d\x37\xb7\x50\x6e\x38\x15\xa5\xbc\x14\x02\x32\x5d\x3f\xab\x5f\x9a\xdc\xc2\xa3\x7e\x86\xc3\xef\x87\x97\x67\xe3\xe1\x18\xc7\x7b\x30\x9a\xd0\x70\xfe\xe1\xed\x2f\x6f\x47\x1f\x3e\x5e\xe1\x8e\xae\x8b\x36\x88\x8c\x3e\x0c\xce\xc6\x63\x5f\xba\x4e\xc5\xff\x89\xe6\x10\x19\xec\x38\xd3\xd6\x51\x17\xf2\xb5\x51\xd6\xbe\x86\xc8\x3f\x61\x6f\xdb\x71\x0c\x6b\xa8\x23\x5b\x1b\x1b\xaf\x03\x97\xc3\xc0\x7e\x33\xed\x2a\xcf\x65\x3f\xe8\x9e\xd8\x5d\x1a\xc6\xeb\xe0\x28\xf6\x88\x3c\x7b\xee\x07\xbd\xc2\x47\xc2\x36\x69\xd8\xf0\xe7\x5b\x25\x79\x50\x63\x84\x4e\x3a\xa0\xe9\x04\x0e\x56\xbf\x5c\xd8\xbf\xe6\x42\x86\xf1\xcb\x7e\xdf\x34\x69\x1a\x1d\x9c\x92\x9c\xc8\x68\xdf\x20\x2b\xe0\x40\x1e\x6f\x07\xa6\x48\x25\x92\x94\x92\x60\x1a\x0d\x43\xca\xbe\xb5\x0b\x69\x3c\x49\xe6\x5e\xea\xab\x1b\x27\x30\x5f\x43\x9c\x5e\xc9\xab\xeb\xfa\xde\xb8\xee\x0b\x75\x2e\x8c\x1a\x7a\xd9\xd2\x7d\x18\xbb\x64\x0e\x2e\x3f\x63\xc1\xcf\x92\xda\x7f\x03\x00\x00\xff\xff\xd7\x35\x02\x4a\xd0\x33\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 13264, mode: os.FileMode(420), modTime: time.Unix(1452657902, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceMysqlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x5b\x6f\xda\x30\x14\x7e\xe7\x57\x58\x7e\xda\x24\xc6\x28\x95\x2a\x2d\x9a\x26\x51\xa0\x55\xa4\xb5\x43\x85\x75\x0f\x53\x1f\x8c\x7d\x40\xd6\x82\xed\xd9\x4e\x27\x56\xf1\xdf\x67\x87\x24\xd8\x84\xd0\x8b\xb4\x5e\x50\xc8\x39\xdf\xf9\xce\xdd\x7e\x7a\x42\x0c\x96\x5c\x00\xc2\x06\xf4\x23\xa7\x80\xd1\x76\xdb\x41\xe8\xc9\xfd\x23\x84\x87\x3f\x66\x73\x58\xab\x8c\x58\xb8\x92\x7a\x4d\xec\x3d\x68\xc3\xa5\xc0\x28\x41\x78\xd0\x3f\xeb\x7f\xe8\x7f\x72\x7f\xb8\xbb\x53\x9f\x12\x4d\xd6\x60\x9d\x0e\x4e\x4a\x13\xde\x48\x96\x49\xea\x2c\xb0\x99\x95\x9a\xac\x20\x90\x39\xe9\x7c\xa3\xa0\x30\x77\x9b\xaf\x17\xa0\x4b\x53\x85\x68\x0c\x4b\x92\x67\xb6\x90\x9e\xf5\x63\x89\xa1\x9a\x2b\x5b\xb9\x52\x53\x20\xb3\xe3\x40\x86\xff\x05\xf4\xee\xfa\xf2\x3d\x2e\x51\xdb\x0a\x8e\xc7\xc4\x92\x05\x31\x6d\x7e\xcc\xac\xe6\x62\xd5\xe6\x07\x51\xea\x94\x23\xa5\x2a\x62\x25\x07\x12\x2e\x21\x4d\x17\x52\x61\x2c\x11\x14\x0a\xd2\xb7\xb8\xc1\x16\x3d\x3b\xe8\xad\x39\xd5\xf2\x94\x3b\x15\x0f\xa2\x19\x31\x06\x2d\xa5\x0e\x3c\x93\x0c\x4c\xd3\xb5\xa9\x53\xfc\x23\x35\x7b\x85\x5b\x31\xe7\xcc\x35\x12\x68\xa4\x2a\x3b\x0d\x86\x59\xbe\x10\x60\xcd\x11\x02\x87\xfe\xca\x8d\xfd\xec\xda\x2e\x49\x26\xa3\x41\x92\xec\x74\x93\x24\x65\x5f\xda\x38\x1d\xe8\x7e\x3a\x42\xa6\xb4\xda\xa0\xfb\xee\x3a\xbb\xa8\xc2\x7f\x28\x77\x19\x6b\x5e\x51\x34\xc8\xef\x15\x3d\x1e\xe7\x3e\x44\xe7\xbc\x8f\xef\x74\x78\xb5\xe5\x4e\x60\x1f\x7f\xcb\xad\xca\xa3\x4c\xe2\xa9\xd4\xf6\xfc\xbc\x7f\x31\xa7\x6a\xc8\x98\xf6\x22\x67\x80\x64\x39\xec\x1e\xaf\x44\x92\x5c\x83\x1d\x5a\xeb\xbe\xff\xdc\x77\x08\xee\x22\x3c\x11\x4c\x49\x2e\x6c\xcf\x23\xc1\x18\x8c\x1e\xd0\x36\x6c\x8d\xbd\x6d\xff\xf8\x36\xdb\x05\xf2\xc0\xf0\x44\x3c\xde\x6c\xcc\xef\x2c\x9c\xcc\xc8\xf2\x1d\x2c\x7d\x22\x6a\xf9\x51\x74\xd8\xb9\xc7\xd0\xb5\xfc\x28\x3a\x6c\x93\x63\xe8\x5a\xee\xd1\x51\x15\xee\xc0\xc8\x5c\x53\x88\xea\x30\x03\x9a\x6b\x6e\x37\xd7\x5a\xe6\xea\xb9\x16\x88\x95\x83\x46\x98\x6a\xa9\x40\x5b\x0e\xf1\xb4\x38\x49\xa1\x7a\xd0\x27\x6b\x1f\x07\xaa\x16\x79\x37\x54\x8f\x18\x52\xb1\x2a\xca\xeb\x8a\x14\xe8\x20\x1f\x6c\xaa\x1c\xa5\x95\x54\x66\xde\xa0\xa5\xca\xd7\xee\x4a\xcb\x75\x59\x70\xec\xeb\xef\xdf\xcd\xe5\xe1\x9b\x11\x67\x3a\xf5\xa1\xba\x5d\xdd\x2b\x7e\x3f\x9e\x5d\xe0\x32\x57\xbb\x9f\x87\xc8\x27\x37\x1b\x29\x8b\x72\xec\xa7\x25\x00\x6c\x5b\x56\xc7\x73\x39\xbd\x1b\xbb\x8f\xf1\x65\xa8\xfc\xa2\x9c\x46\x90\x57\xe4\xb6\x00\xa5\xcc\x44\xb1\x54\x5b\xee\x64\x3c\xf5\x84\x3c\x1b\xcc\x7e\x96\x5e\x12\xc9\xb1\x23\xb7\xf6\xac\x21\xdc\x3b\x54\xa6\xa1\x22\x1b\xf9\x53\x23\xc2\x46\x07\x57\x2b\x2e\x65\x20\x2c\x5f\x72\xd0\x31\xb1\x8f\x67\x66\x09\xfd\x75\xbb\x1b\xa4\x03\xf8\x6d\x3d\x7e\xcd\x79\xef\xb6\x16\xaa\x81\x0a\xeb\x7e\x00\x9c\x88\x95\xbb\xec\xd4\xf5\x8c\xeb\x78\x43\x8c\xbb\xba\xc4\x7b\xa0\x39\xfc\x2d\x90\x78\xf9\x34\x37\x4e\x04\x8b\x47\x27\x92\xe4\x8b\x8c\xd3\x6c\x33\xa4\x6e\x9f\x18\xbe\xc8\x0a\x67\x97\x24\x33\x87\x4d\xb7\xab\x5d\xd5\x2a\x2b\x35\x88\xe5\xee\xcc\x88\x66\xbe\x98\xf6\x30\x49\xd1\xca\x71\x3b\xed\xa1\xd9\xa6\x9d\xea\x73\xdb\x71\x17\x45\x10\xcc\xdf\x0d\xff\x05\x00\x00\xff\xff\xc3\x37\xc1\xe6\x33\x0a\x00\x00")

func templatesServiceMysqlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceMysqlTmpl,
		"templates/service/mysql.tmpl",
	)
}

func templatesServiceMysqlTmpl() (*asset, error) {
	bytes, err := templatesServiceMysqlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/mysql.tmpl", size: 2611, mode: os.FileMode(420), modTime: time.Unix(1444747254, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServicePapertrailTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x7f\x6b\x1b\x39\x10\xfd\xbf\x9f\x42\x88\x42\xa1\x38\x4e\x9c\x50\x8e\x13\xdc\x1f\xbe\x34\x69\x7d\x75\x12\xb3\x4e\x7b\x70\x25\x1c\xca\xee\xd8\xd1\x79\x2d\x2d\x92\x36\x4d\x1a\xfc\xdd\x6f\x24\xed\x7a\x7f\xc6\x0e\x81\xe3\xa0\x6e\x58\xe9\x69\x34\x7a\x33\xf3\x66\x9e\x9e\x48\x02\x0b\x21\x81\x50\x03\xfa\x5e\xc4\x40\xc9\x66\xf3\x86\x90\x27\xfc\x11\x42\xc7\x7f\xce\xaf\x61\x9d\xa5\xdc\xc2\xb9\xd2\x6b\x6e\xbf\x81\x36\x42\x49\x4a\x18\xa1\xc7\x47\xa3\xa3\x83\xa3\x5f\xf1\x1f\x1d\x04\xf8\x8c\x6b\xbe\x06\x8b\x18\xca\x0a\x13\xb8\xfa\x55\xa7\xb5\x4f\x5c\xb8\x7e\xcc\xc0\x5b\x98\x5b\x2d\xe4\xb2\x38\xed\xb7\x3e\x82\x89\xb5\xc8\x6c\x79\xc7\x8c\x67\xa0\xad\xe6\x22\x25\x5f\xa3\xe9\x80\xc0\x70\x39\x24\xef\x52\xb5\x34\xa3\x61\xb6\xdd\xe3\x59\x36\x8c\xd5\x9a\x8d\x46\xc7\x27\x1f\xde\xd1\xc2\xdc\xc6\xff\xdd\x14\xbe\x45\x60\x54\xae\x63\x68\xb8\x36\xe5\xeb\xdb\x84\x9f\x3d\x40\x9c\xbb\x2b\x23\x95\x42\x8f\xab\xcc\x13\xc1\xd8\x64\x7c\xc1\x98\xc7\xd4\x3c\x9e\x69\xe5\xdc\x10\x0d\xc3\x81\x3c\x63\xf2\x35\x38\xfc\x4c\xa5\x22\x7e\xfc\xa8\x62\xfc\x96\xb6\x85\x43\x64\xc9\x6a\x20\xf5\xf8\x00\x79\x1d\xfd\x52\xbb\xc4\x83\xe6\x16\xa3\x50\x9c\xff\xde\xd8\x22\x2d\x7b\x1e\x7e\xb6\x58\x40\x6c\xbd\xef\x69\xaa\x7e\xb4\xac\x15\xae\x0b\x19\x8b\x8c\xfb\xf0\xe0\x05\x45\x06\xa0\x79\x42\x53\xcf\xcc\x90\xaf\xf9\x4f\x25\xf9\x0f\xe3\xf8\xa5\xe4\xa6\xa4\xb3\x61\x67\x1c\xdb\xe0\x3d\x9e\x33\xd6\xb0\xea\xe1\x78\xa2\x05\xdf\x34\xbe\xeb\xbb\x0d\xcb\x18\x78\x7b\xe7\x9c\x3f\xa4\xcd\x65\xc7\x64\xe0\xba\xc9\x41\x9b\x81\x80\x7c\xbc\xc4\x7c\x74\x66\x42\xa0\x4f\x53\x95\x27\x21\x91\xd1\xe1\x2f\x98\xf7\x46\x98\x0e\x33\x74\x4f\xb8\x3c\xe6\x25\x21\xf3\xc0\x5d\x61\xeb\x73\xbc\x38\xb6\x3f\x7c\x1e\x56\x51\xdf\xbb\x4f\xca\x40\xb2\x89\xbc\x57\x2b\x38\xcf\x65\x38\xd0\x8b\xbe\x79\xe6\x92\xb2\x74\x76\x5d\xf3\xfe\x19\x93\x3d\xab\x3d\x29\xf4\xdf\xd3\xb0\x0a\xc1\x66\x9f\xc0\x46\x10\x2b\x9d\x74\xe3\xde\x87\x9d\xdf\x71\x9d\x4c\x50\xd1\xb8\x55\x7a\xff\x89\x20\x5f\xb7\x80\xc2\x06\x7c\xbd\x1f\x3f\x15\xc6\x06\xec\x0e\x77\x9c\xda\xb1\x53\x04\x59\x98\xaa\xe5\x27\xad\xf2\xec\xa5\xe0\x7d\x7e\x78\xf4\x2c\xb7\x08\x3d\xbb\xc7\x1c\x35\xaf\x4e\x8c\xfe\x04\xf8\x5f\x42\x1d\xbb\x2a\x5f\x94\x55\x5e\x8b\x09\x8f\x57\xaf\x7f\x20\x0a\xe4\xb9\x64\xec\x0f\x25\x0a\xa5\xa3\x03\xf7\x3f\xd7\x92\xa1\x3c\xb2\xd6\xa5\xb8\xf9\xe4\x8e\x2f\xb6\xcd\x23\x82\xa5\x6f\x6a\x9b\x01\xa1\x3d\xdb\xe3\x38\x56\xb9\xb4\x93\xa4\x40\x18\xe7\xed\xa1\xc3\x79\x18\x29\x71\xfe\x15\x5e\xd5\x1c\xec\xf0\xbd\xd3\xe4\x9b\x96\xaa\x16\xdc\x77\xd6\x76\x6b\x71\xfd\xab\x42\x96\xab\xdb\x48\xe2\xc8\xa0\xb9\x5c\x02\x79\xbb\x1a\x90\xb7\xf7\x84\xfd\x46\x86\xe3\xe8\xd2\x84\xb9\xa1\xe0\x0d\x41\x79\x86\x6d\x11\x41\xb8\x7e\x81\xfd\xd9\xf5\xf9\x56\x7f\xdc\x8e\x01\xfe\x61\x41\xa1\x19\xf3\x79\x38\xf7\xac\x97\xe7\x9a\x1d\xa0\xea\xb7\xa4\xd3\x48\xcf\x24\xbf\x4d\x21\x71\x3b\x56\xe7\xd0\xea\xa0\x35\xd3\x63\x1d\xe6\x0b\x74\x14\xdf\xb0\xd9\xb4\x9b\x6d\x29\x94\x9e\x69\xb2\x8d\x3e\x2a\xc2\xd8\x5a\xb7\xf0\xbd\x36\x9a\x60\xf5\x2c\x01\xd5\x01\x5f\x82\x66\x3b\x3d\xd2\xb5\x00\x74\x58\x2e\x67\xca\x88\xed\x64\x73\x1d\x4d\x2e\xfe\xfe\x7c\x15\x4d\xfe\xba\xba\xac\x27\x65\xc5\x62\x9d\x72\x90\x49\x45\x70\xf7\xea\xe7\x67\x96\x92\xd6\xad\xf2\xbf\x68\x74\xf9\xcc\x65\x92\x7a\xbb\x54\xc8\x04\x1e\x86\x77\xc5\x42\x23\x14\xe5\xb8\xd4\xe5\xa6\x6f\xae\xea\xa7\x87\x9e\xaa\x04\xba\x13\xd1\xfc\xe4\xf7\x3c\x5e\x81\xed\x2b\xbc\x83\x50\x79\xb1\xc2\x96\xf6\xb0\xa3\xd0\x42\x65\xb4\x63\x71\xf2\x05\x1e\x1d\x36\xf4\xc5\xc3\x6a\x88\x1c\xfe\x14\x19\x7d\x76\x1e\x89\xb0\x38\x45\x98\x25\x24\x7a\xfc\x4f\x53\xaf\xe9\x35\xee\xa9\xdc\xab\xd7\xf1\x07\xda\xad\x9e\xc6\x28\x7a\x95\xdb\x2c\xb7\x75\xd2\x5f\x51\x54\x53\x21\x57\xed\xb0\x7d\xe3\x69\xee\x5d\xdc\xa6\xf5\x0b\x92\xa9\x33\x9e\x97\x56\x2a\x5e\x1d\xa4\xfd\x96\x37\xee\x57\x19\xfb\x37\x00\x00\xff\xff\xcd\x2b\x5b\x59\x4a\x0c\x00\x00")

func templatesServicePapertrailTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServicePapertrailTmpl,
		"templates/service/papertrail.tmpl",
	)
}

func templatesServicePapertrailTmpl() (*asset, error) {
	bytes, err := templatesServicePapertrailTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/papertrail.tmpl", size: 3146, mode: os.FileMode(420), modTime: time.Unix(1446924998, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServicePostgresTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\x5b\x6f\xe2\x38\x14\x7e\xe7\x57\x58\x7e\xda\x95\x58\x96\xb2\x17\x69\xa3\xd5\x48\x14\x68\x15\x69\xa6\x83\x0a\xd3\x79\x18\xf5\xc1\xd8\x07\x64\x4d\xb0\x2d\xdb\xe9\xa8\x53\xf1\xdf\xe7\x38\x24\x21\x26\x04\x7a\x47\x21\xe7\x7c\xdf\xb9\x1f\xfb\xe5\x85\x08\x58\x4b\x05\x84\x3a\xb0\x4f\x92\x03\x25\xbb\x5d\x8f\x90\x17\xfc\x27\x84\x8e\xbf\x2e\x96\xb0\x35\x19\xf3\x70\xa3\xed\x96\xf9\x07\xb0\x4e\x6a\x45\x49\x42\xe8\x68\x78\x35\xfc\x63\xf8\x1f\xfe\xd1\xfe\x5e\x7d\xce\x2c\xdb\x82\x47\x1d\x9a\x94\x14\x81\x24\xcb\x34\x47\x06\xb1\xf0\xda\xb2\x0d\x34\x64\x28\x5d\x3e\x1b\x28\xe8\xee\xf2\xed\x0a\x6c\x49\x55\x88\xa6\xb0\x66\x79\xe6\x0b\xe9\xd5\x30\x96\x38\x6e\xa5\xf1\x95\x2b\xb5\x09\xe2\xf6\x36\x88\x93\x3f\x81\xfc\x76\x7b\xfd\x3b\x2d\x51\xbb\x0a\x4e\xa7\xcc\xb3\x15\x73\x5d\x7e\x2c\xbc\x95\x6a\xd3\xe5\x07\x33\xe6\x9c\x23\xa5\x2a\x11\xa5\x0d\xa2\x30\x21\x6d\x17\x52\xe5\x3c\x53\x1c\x0a\xa3\xef\x71\x43\xac\x06\x7e\x34\xd8\x4a\x6e\xf5\x39\x77\x2a\x3b\x84\x67\xcc\x39\xb2\xd6\xb6\xe1\x99\x16\xe0\xda\xae\xcd\x51\xf1\x87\xb6\xe2\x0d\x6e\xc5\x36\x17\xd8\x48\x60\x89\xa9\x78\x5a\x16\x16\xf9\x4a\x81\x77\x27\x0c\x20\xfa\xa3\x74\xfe\x7f\x6c\xbb\x24\x99\x4d\x46\x49\xb2\xd7\x4d\x92\x54\x7c\xe8\xb2\x89\xa0\x87\xf9\x84\xb8\x92\xb5\x65\xee\x0b\x76\x76\x51\x85\xf7\xe4\xd9\x68\xe7\x37\x16\xf3\x74\x39\xe0\xbc\xb2\xd3\xf2\xe0\xc1\xf0\xd3\xc1\x1e\xe2\xc4\x08\x42\x90\xe7\x63\xac\x99\x7b\x0d\x7e\xfa\x39\xf7\x26\x8f\xd2\x49\xe7\xda\xfa\x7f\xfe\xfe\x6b\xb4\xe4\x66\x2c\x84\x0d\x22\x24\x60\x59\x0e\xfb\xc7\x1b\x95\x24\xb7\xe0\xc7\xde\xe3\xf7\x6f\x87\x36\xa1\x7d\x42\x67\x4a\x18\x2d\x95\x1f\x04\x24\x38\x47\xc9\x23\xd9\x35\xfb\xe3\xc0\x1d\x1e\xdf\xc7\x5d\x20\x8f\x88\x67\xea\x69\x5e\x66\xbb\x39\xa1\x11\xf9\x3d\xac\x43\x2e\x6a\x79\x17\x41\xb3\x89\x4f\x11\xd4\xf2\x2e\x82\x66\xd3\x9c\x22\xa8\xe5\x81\x20\x2a\xc7\x3d\x38\x9d\x5b\x0e\x51\x41\x16\xc0\x73\x2b\xfd\xf3\xad\xd5\xb9\xb9\xd4\x0b\xb1\x72\xa3\x23\xe6\x56\x1b\xb0\x5e\x42\x3c\x3b\x28\x29\x54\x8f\x1a\xa6\x6a\x5d\x52\x6d\xf6\x7e\x13\x11\x19\x49\xd5\xa6\x28\x35\x16\xac\xa1\x43\x42\xbc\xa9\x41\xab\x5e\x73\x9d\x05\x4e\xcf\x4d\xa8\xe3\x8d\xd5\xdb\xb2\xf8\x34\xf4\x42\x78\xb7\xd4\xc7\x6f\x26\x52\xd8\x34\x44\x8b\xcb\x7b\x50\xfc\xfe\x79\xf5\x2f\x2d\xd3\xb5\xff\x79\x8c\x7c\xc2\x39\x49\x45\x94\xe6\x30\x39\x0d\xc0\xae\x63\x97\x5c\x4a\xeb\xfd\x14\x3f\xa6\xd7\x4d\xe5\x57\xa5\x35\x82\xbc\x2d\xbd\x05\x2e\x15\x2e\x0a\xa7\xda\x7c\x67\x43\xaa\x07\xe6\x62\x3c\x87\xd1\x7a\x4d\x30\xa7\x8e\xe1\xda\xb3\x96\xf0\xe0\x50\x99\x89\xca\xd8\x24\x9c\x24\x11\x36\x3a\xcc\x3a\x71\xa9\x00\xe5\xe5\x5a\x82\x8d\x0d\x87\x78\x16\x9e\xf1\xef\x77\xfb\x71\x3a\x82\xdf\xd5\x43\xd8\x9e\xfd\x7e\x67\xad\x5a\xa8\x66\xe9\x8f\x80\x33\xb5\xc1\x0b\x10\x3d\xbd\xec\x51\xfe\x89\x39\xbc\xd1\xc4\x0b\xa1\xbd\x05\x3a\x20\xf1\x22\x6a\x6f\x9f\x08\x16\x0f\x50\x24\xc9\x57\x99\xe4\xd9\xf3\x98\xe3\x62\x71\x72\x95\x15\xfe\xae\x59\xe6\x8e\xfb\x6e\x5f\xbe\xaa\x5b\x36\xe6\x88\x08\x4f\x91\x68\xf2\x8b\x99\x6f\xe6\x29\xda\x3d\xb8\xdc\x1e\xdb\x9d\xda\xab\x3e\x77\x3d\xbc\x3f\x82\x12\xe1\xca\xf8\x2b\x00\x00\xff\xff\xa1\x5a\xfe\x19\x4a\x0a\x00\x00")

func templatesServicePostgresTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServicePostgresTmpl,
		"templates/service/postgres.tmpl",
	)
}

func templatesServicePostgresTmpl() (*asset, error) {
	bytes, err := templatesServicePostgresTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/postgres.tmpl", size: 2634, mode: os.FileMode(420), modTime: time.Unix(1444747254, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceRedisTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xd1\x6f\xda\x3e\x10\x7e\xe7\xaf\xb0\xfc\xf4\xfb\x49\x8c\x41\x27\x75\x6a\x34\x4d\x42\x0c\xaa\x48\x5b\x87\x80\x76\x0f\x55\x1f\x8c\x7d\x50\x6b\x89\x1d\xd9\x4e\xb7\xaa\xe2\x7f\xdf\xd9\x09\x10\x27\x1d\x6d\xa5\xb5\x80\x22\xdf\xdd\xf7\xdd\x7d\x77\xb6\xf3\xf4\x44\x04\x6c\xa4\x02\x42\x2d\x98\x07\xc9\x81\x92\xdd\xae\x47\xc8\x13\x7e\x09\xa1\xe3\x1f\xcb\x15\xe4\x45\xc6\x1c\xcc\xb4\xc9\x99\xbb\x01\x63\xa5\x56\x94\x24\x84\x9e\x0d\x47\xc3\x77\xc3\x0b\xfc\xd0\x7e\xe5\x3e\x67\x86\xe5\xe0\xd0\x87\x26\x35\x04\xae\x7e\x61\x8e\xad\x99\x85\xc6\x1a\xae\xae\x1e\x0b\x08\x30\x4b\x67\xa4\xda\xd6\x10\x55\x00\x6c\x58\x99\xb9\x60\x1d\xc6\x06\xcb\x8d\x2c\xdc\x3e\x83\xda\x91\x88\x9a\x81\x48\x25\xe0\x37\xad\x03\x76\xfb\x48\x9a\x2a\xeb\x98\xe2\x10\x38\xbb\x59\x9c\x4c\x02\x8d\x9c\xf1\x7b\x18\xb8\xb3\x41\x2e\xb9\xd1\x7f\x4b\x08\x1d\x57\xf7\x40\x1c\x22\x12\xbd\xc1\x54\x2a\x4e\xe2\x34\x29\xb1\xf8\x4e\x52\x73\x66\xed\x2f\x6d\xc4\x1b\x64\x89\xab\xbf\x56\x88\x2b\xc8\x7f\x48\xb0\x06\x62\x20\xd7\x0f\x20\xfe\xef\x12\x2d\xcb\xb5\x02\x67\x9f\x2f\xfc\xab\xb4\xee\x13\xb6\x39\x49\xa6\x93\xb3\x24\xa9\x7c\x93\x24\x15\x9f\x4f\xd4\x79\x33\x9f\x10\x5b\xa3\x76\xe8\x6e\x0a\xfe\x3c\xd5\x91\x05\xe3\x3d\xc5\x69\x86\x03\x72\xaf\x81\x4f\xbf\x97\xae\x28\xa3\x62\xe8\x5c\x1b\x77\xfe\xe1\xe3\xc5\x8a\x17\x63\x21\x8c\x37\x21\x00\xcb\x4a\xa8\x1e\x67\x2a\x49\x2e\xc1\x8d\x9d\x6f\xe6\x2d\xa1\x0b\x28\x32\xc9\x99\xa7\xba\x34\xba\x2c\x68\x1f\x31\x8c\xcc\x99\x79\x9c\x2a\x31\xd7\x52\xb9\x81\x07\x02\x6b\x29\xb9\x23\xbb\x66\xcf\x8e\x54\xfe\xf1\x9f\x50\x05\xa0\x16\xcf\x54\x3d\x2c\x40\x48\xdb\xdc\x3a\x11\xd1\x02\x36\x5e\xa6\x83\xdd\x47\x47\x3a\x2d\xc0\xea\xd2\x70\x88\x94\x5a\x02\x2f\x8d\x74\x8f\x55\x2e\x2f\x34\x29\x76\x6e\xb4\x6a\x6e\x74\x01\xc6\x49\x88\x47\x0a\x2d\xc1\xb5\xd5\x49\xe3\xeb\x20\xfb\xd3\xa5\xdf\x74\x8f\x18\x52\xb5\x0d\x8a\xa3\x6e\x0d\x1f\xe2\x8b\x4d\x0b\xa4\x74\x9a\xeb\xcc\x03\x3a\x1e\x64\x9c\x19\x9d\xd7\x3d\xa0\xbe\x25\x7e\x6d\xa5\xdb\x2b\x13\x29\x4c\xea\x4b\xa5\xa3\xe1\x20\xfc\xbf\x1f\x9d\xd3\x5a\xab\xea\xef\x2e\xca\x09\xa7\x37\x15\x91\xc6\x7e\x9e\x1b\x01\xbb\xce\xc0\x4f\xfc\xf1\x50\x6d\x9c\x17\x85\xcd\x98\x75\x32\x04\xec\xf7\xda\x1b\xe4\x6d\x29\xbb\xa8\x94\x0d\x28\x64\xdb\x82\x39\xec\xfb\x54\xd8\xa8\x9c\xfd\x69\x70\xb2\xa4\xce\xdc\xbe\xb2\xa4\xee\xbc\xbf\xa6\xae\x71\xe9\x34\xde\x2d\x92\xcf\x98\xcc\xf0\x10\x33\x53\xc5\xd6\x19\xf8\x36\x6c\x58\x66\xa1\xdf\x76\xfe\x26\x95\x36\xf5\x55\x74\x5d\x6c\x0d\x13\x3e\x1f\x67\xca\xd8\x35\x24\x75\xa5\xc5\xe1\xdc\x3f\x88\x10\x5d\x08\xc7\xba\x9f\x6d\xe7\x15\x5e\x69\x51\x70\xa7\xdf\x2d\x80\xa9\xda\xe2\x8d\x7a\x98\xfd\xb8\x29\x57\x65\x1e\xe2\x27\x59\x69\xeb\x8b\x92\x8e\x62\x97\x78\x86\x9b\x96\xb6\xbc\xf1\x40\x1c\x33\x0c\x8d\x59\x3a\xc6\x7f\x86\xe4\x5b\xf9\xc5\xdb\x2e\x8c\xc7\x6d\x73\x40\xa2\x7d\x8f\x07\xcb\x5d\x77\x50\x7a\xfb\xdf\x5d\x0f\x5f\x21\x40\x09\xff\xd6\xf0\x27\x00\x00\xff\xff\x64\xe5\x91\xd9\x4d\x08\x00\x00")

func templatesServiceRedisTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceRedisTmpl,
		"templates/service/redis.tmpl",
	)
}

func templatesServiceRedisTmpl() (*asset, error) {
	bytes, err := templatesServiceRedisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/redis.tmpl", size: 2125, mode: os.FileMode(420), modTime: time.Unix(1443527400, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceSyslogTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x57\x5d\x6f\xdb\x36\x14\x7d\xef\xaf\x20\x88\x02\x05\x0a\x47\x89\x13\x14\xc3\x08\xec\xc1\x6b\x92\x2e\x5b\xd6\x1a\x76\xda\x3d\x0c\x79\x60\x64\xda\xe5\x2c\x91\x02\x49\xa5\x71\x02\xff\xf7\x5d\x7e\xc8\x12\x29\xd9\x09\xf2\xb0\x01\xb1\x03\x91\x87\x97\x57\xe7\xdc\x2f\x3f\x3d\xa1\x05\x5b\x72\xc1\x10\xd6\x4c\xdd\xf3\x9c\x61\xb4\xdd\xbe\x41\xe8\x09\x3e\x08\xe1\xc9\x5f\xf3\x1b\x56\x56\x05\x35\xec\x52\xaa\x92\x9a\x6f\x4c\x69\x2e\x05\x46\x04\xe1\xd3\x93\xf1\xc9\xd1\xc9\xcf\xf0\x87\x47\x1e\x3e\xa5\x8a\x96\xcc\x00\x06\x93\x60\x02\x56\xbf\xaa\xa2\xf3\x08\x0b\x37\x9b\x8a\x39\x0b\x73\xa3\xb8\x58\x85\xd3\x6e\xeb\x9c\xe9\x5c\xf1\xca\x34\x77\xcc\x37\xba\x90\x2b\xf4\x75\x76\x3d\x42\x2c\x5b\x65\xe8\x1d\x3c\xea\x71\x56\xd1\x8a\x29\xa3\x28\x2f\x68\x55\x65\xb9\x2c\xc9\x78\x7c\x7a\xf6\xe1\x1d\x0e\xa6\xb6\xee\xff\x36\xf8\x35\x63\x5a\xd6\x2a\x67\x91\x5b\xd7\xb4\xbc\x5b\xd0\x8b\x07\x96\xd7\xf6\xba\x99\x2c\xd8\x80\x9b\xc4\x91\x40\xc8\xd5\xe4\x4f\x42\x1c\xa6\xe3\xed\x54\x49\xeb\x06\x8f\x0c\x7b\xe2\xb4\xae\x4b\x66\xf1\x53\x59\xf0\x7c\x73\x2e\x73\x78\x16\x26\xc1\x01\xb2\x61\xd4\x13\x7a\x7a\x04\x9c\x8e\x7f\xea\x5c\xe2\x40\x73\x03\x0a\x84\xf3\x7f\x47\x5b\x28\xb1\xe7\xe0\x17\xcb\x25\xcb\x8d\xf3\xbd\x28\xe4\x8f\xc4\x5a\x70\x9d\x8b\x9c\x57\xd4\x49\x03\x17\x04\xf5\xc1\x3c\xc2\x85\x63\x26\xa3\x25\x7d\x94\x82\xfe\xd0\x96\x5f\x8c\x6e\x1b\x3a\x23\x3b\x93\xdc\x78\xef\xe1\x9c\x36\x9a\xb4\x2f\x0e\x27\x12\xf8\x36\x7a\xee\xee\x46\x96\x21\x8c\xcc\x77\xeb\xfc\x31\x8e\x97\x2d\x93\x9e\xeb\x98\x83\x94\x01\x8f\xdc\x7c\x86\x58\xb4\x66\xae\xe5\x6a\x5e\xdf\xed\xe2\xea\x92\x17\x10\xa1\x89\x94\xd1\xc1\xbd\x62\x39\xcc\x4b\x04\x73\xc0\x43\xa2\x0d\xb9\x1d\x8e\x3d\x2f\x9e\x83\xb5\xc4\x0f\xee\xa3\x46\x46\x72\x25\xee\xe5\x9a\x5d\xd6\xc2\x1f\x18\x44\xdf\xee\xb9\xa4\x49\x9c\x43\xd7\xbc\xdf\x63\x72\x60\x75\x20\x80\xfe\x03\x1a\xa0\x60\x90\x8f\x8a\x81\x18\x10\x09\x9f\x94\xac\xab\x3d\xc6\x7a\x60\xa8\x4f\x8c\x96\xcf\xa0\xa7\xb5\x01\xe8\xc5\x3d\x08\xad\x5f\xcd\xee\x30\x8b\xff\x0b\x5f\x79\x21\xeb\xc5\xd2\x55\x7b\x40\x12\x5f\x91\xef\x18\x44\x73\xbe\x7e\xfd\x0b\x42\x8d\xb9\x14\x84\xfc\x2e\x79\x28\x16\x78\x64\xbf\xa9\x12\x04\x2a\x0c\x49\x2e\x85\xcd\x27\x7b\x7c\xb9\xab\xbf\x33\xb6\x72\x3d\x61\x3b\x42\x78\x60\x7b\x92\xe7\xb2\x16\xe6\x6a\x11\x10\xda\x7a\x7b\x6c\x71\x0e\x86\x1a\x9c\x7b\x0b\x57\x18\x2c\xec\xf8\xbd\x2d\x6b\xb7\x49\x61\x0a\xdc\xf7\xd6\x0e\x97\xb3\xee\x53\x8b\x6c\x56\x77\x4a\x42\xc7\x55\x54\xac\x18\x7a\xbb\x1e\xa1\xb7\xf7\x88\xfc\x82\xb2\xc9\xec\xb3\xf6\x6d\x37\xf0\x06\xa0\xba\x82\xce\x02\x20\x58\x1f\xac\x5f\x69\xc3\x39\x67\x15\x13\x0b\xfd\x25\xb0\x9b\x58\x70\x95\x60\xca\x54\xc9\xb5\x6f\xdf\x91\x62\x6d\x43\x76\x1c\x5d\xdb\xa8\x26\x03\x57\xc6\xd5\xb8\xed\x7d\xa8\xd7\xd4\x20\x68\x0c\x17\x4e\xcb\x89\x72\x9d\x3c\xe8\xff\x89\x99\x89\x31\xde\x47\xef\xd5\xae\x30\x81\x1e\x0e\xdb\xeb\x33\xd8\x5f\x0f\x5d\x01\xbe\xfd\x58\x90\xf6\xc7\x26\xb1\x9d\xb2\x16\x00\xef\x0f\xe4\x6e\xb7\xdd\x78\x6d\x09\xee\x5c\xf0\x1c\x53\x09\xcd\xd1\x4c\xe0\xc1\x84\x74\xe0\x2f\xa7\x28\x24\xa2\x75\x76\xb8\x50\xa7\x24\x84\xf5\xe6\x15\xdb\x04\x48\x78\xec\xd1\xd7\xb6\x7a\xd4\xcb\x43\x9c\xd9\x3c\x74\x85\xec\x70\xce\xc5\xc3\xc0\x6d\x5f\xa4\xb9\xcb\xf4\x90\x87\xb1\x83\xfd\x0c\x1d\x3c\x1a\xb4\x7d\xa1\x74\x00\x83\x80\x6f\xd3\x26\xa5\x61\xff\x28\xd7\xc8\x36\xc0\xf4\x81\x89\xee\x37\x2a\x16\x85\xcb\x3b\xcc\xc5\x82\x3d\x64\xdf\xc3\x42\x24\x79\x33\x45\x46\xd1\x8e\x2c\xd1\x43\xe3\x66\x08\xf9\x84\x4c\xfc\x51\x2e\x58\x7f\x50\x9c\x9f\xfd\x5a\xe7\x6b\x66\x86\x8a\xe9\x91\xaf\xa6\xb9\x84\x10\x7a\x38\x20\xa4\xaf\x76\x29\xfd\x67\x7f\xb0\x0d\xde\xc5\xe1\xb1\x76\x33\x77\xf6\xc8\x2b\xbc\x77\x44\x9b\x81\x94\xdc\x8f\x57\x02\xbc\xfd\x47\xc7\x34\xdc\xc0\x9e\xac\x5d\x37\x3a\xfd\x80\xfb\xd5\x30\x9a\xce\xbf\xd4\xa6\xaa\x4d\x97\xf0\xd7\x14\x49\x2e\xd6\xa9\x64\xdf\x68\x51\x3b\x17\xfb\x21\xb5\x3f\x8c\x7a\xbf\x56\x1a\x2b\x2d\xa7\x16\x92\xbe\xcb\x1b\xfb\x69\x8d\xfd\x1b\x00\x00\xff\xff\xa1\x2a\xd4\x1d\x59\x0d\x00\x00")

func templatesServiceSyslogTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceSyslogTmpl,
		"templates/service/syslog.tmpl",
	)
}

func templatesServiceSyslogTmpl() (*asset, error) {
	bytes, err := templatesServiceSyslogTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/syslog.tmpl", size: 3417, mode: os.FileMode(420), modTime: time.Unix(1452661719, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceWebhookTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x92\x4d\x6f\xb2\x40\x10\xc7\xef\x7c\x8a\xcd\x5e\xbc\x28\xa0\xcf\x73\x29\x37\xd3\x97\x53\x63\x8d\xa0\x9e\x71\x19\x75\x23\x30\x9b\x65\x30\x6d\x88\xdf\xbd\xbb\x4b\x51\x48\xeb\xa5\x4d\x48\xc8\xce\xcb\x6f\x66\xfe\x33\x4d\xc3\x32\xd8\xcb\x12\x18\xaf\x40\x9f\xa5\x00\xce\x2e\x17\xaf\xf1\x18\xe3\xf3\x6d\x9c\x40\xa1\xf2\x94\xe0\x05\x75\x91\xd2\x06\x74\x25\xb1\xe4\x2c\x62\x7c\x16\x4e\xc3\x49\xf8\x60\x3e\x3e\xb6\xc1\xcb\x54\xa7\x05\x90\x89\xe0\x11\xb3\xe9\xc6\xb6\xd6\xf9\xf5\x61\x9e\xc9\x87\x02\x97\x1b\x93\x96\xe5\xc1\xe5\x39\xc7\x13\x54\x42\x4b\x45\x1d\x7b\x0b\xbb\x23\xe2\x89\xad\x57\xaf\x63\x06\xfe\xc1\x67\xa3\x23\x91\xaa\xa2\x20\x38\x68\x99\xf9\x02\xcb\x33\xbe\x9b\x5f\x11\xe8\x54\x9c\x26\x36\x38\x98\xce\xfe\xfd\x1f\x71\x87\xbc\xb4\x64\xfe\x58\x57\x84\x45\x82\x4a\x8a\x5f\xf5\x31\xa4\x2d\x90\xe4\x5e\x8a\xd4\xba\xff\xcc\xf4\xbe\xb8\x7c\x05\x15\xd6\x5a\x40\x4f\xb6\x7e\xa1\xea\xc7\x22\xed\x60\x51\x14\x2f\xe2\xb8\xde\xdd\x0a\x5c\xab\x76\x9b\x32\xb1\x53\x3f\xbc\xd9\x97\x1a\x15\x68\x92\xd0\xe7\x1a\x7b\xdc\xee\x3e\xc1\x13\xd8\xa4\xc6\xf6\xb5\xe7\xd1\x50\xc2\x4e\x89\xb6\x17\x6b\x9a\x6b\x37\xd4\x2d\xfc\xbb\x46\x83\x24\x53\x9e\x50\x60\xee\x86\xb0\x2b\xe5\x3d\xe7\x73\x99\x29\x94\x25\x0d\x89\xf6\x86\x9c\x5c\x9d\x6c\x7d\xf1\xde\x6a\x52\x35\xdd\xbf\xb8\x4d\x9a\xd7\xc0\xef\xe0\x1c\xc6\x33\xb7\xde\x30\x28\x33\x7b\xf5\x9f\x01\x00\x00\xff\xff\x54\xf9\xe1\x63\x0d\x03\x00\x00")

func templatesServiceWebhookTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceWebhookTmpl,
		"templates/service/webhook.tmpl",
	)
}

func templatesServiceWebhookTmpl() (*asset, error) {
	bytes, err := templatesServiceWebhookTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/webhook.tmpl", size: 781, mode: os.FileMode(420), modTime: time.Unix(1446924998, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
	"templates/service/mysql.tmpl": templatesServiceMysqlTmpl,
	"templates/service/papertrail.tmpl": templatesServicePapertrailTmpl,
	"templates/service/postgres.tmpl": templatesServicePostgresTmpl,
	"templates/service/redis.tmpl": templatesServiceRedisTmpl,
	"templates/service/syslog.tmpl": templatesServiceSyslogTmpl,
	"templates/service/webhook.tmpl": templatesServiceWebhookTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{
		}},
		"service": &bintree{nil, map[string]*bintree{
			"mysql.tmpl": &bintree{templatesServiceMysqlTmpl, map[string]*bintree{
			}},
			"papertrail.tmpl": &bintree{templatesServicePapertrailTmpl, map[string]*bintree{
			}},
			"postgres.tmpl": &bintree{templatesServicePostgresTmpl, map[string]*bintree{
			}},
			"redis.tmpl": &bintree{templatesServiceRedisTmpl, map[string]*bintree{
			}},
			"syslog.tmpl": &bintree{templatesServiceSyslogTmpl, map[string]*bintree{
			}},
			"webhook.tmpl": &bintree{templatesServiceWebhookTmpl, map[string]*bintree{
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

