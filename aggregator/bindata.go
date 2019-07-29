// Code generated by go-bindata.
// sources:
// pkg/aggregator/data/blacklist.txt
// DO NOT EDIT!

package aggregator

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

var _blacklistTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x5b\xfb\x72\xe3\xae\xaf\xff\x3f\xef\x42\xa6\xed\x76\x2f\xdf\xc7\x91\x41\xb6\x55\x2e\x62\xb9\x38\x71\x9e\xfe\x8c\x00\xa7\xe9\xb6\xa7\xbf\x9d\xe9\x2c\xba\x18\x63\x01\xe2\x23\x89\x3c\xfd\xc0\x33\x85\x99\x4f\x4f\x41\x39\x0a\x78\x2e\xdb\xe9\x59\xfd\xf7\x5f\x46\x3e\x6b\xf6\xa7\x67\x35\x27\x44\x95\x57\x48\xa8\xa6\x5a\x0a\x87\xdc\x05\x4f\x4f\x86\x9d\x83\x94\xd5\x5d\xf7\xe9\x29\x23\x24\xbd\x62\x58\x28\xe0\xd0\x7b\xf1\x90\x0b\x26\xde\x06\x59\xfb\xfb\x9e\x23\x78\x7f\x4e\xf5\xf4\x7c\xc1\xa9\xab\x9c\xbd\x3b\xbd\xbc\x5e\x7f\xab\x8c\x69\xc3\xa4\x72\x8d\x91\x53\x39\x67\x2a\x78\x7a\xd9\xb9\xa6\xde\xfc\xa1\x1c\x96\x82\x49\x19\xf6\x40\x21\x9f\x03\x96\xd3\x8f\x0b\xec\x01\xcb\x85\x93\xed\x2f\x7e\xa5\x10\xa4\xff\x57\xca\x85\xf3\xca\xb1\x73\xef\x6f\xcb\x67\x4e\xcb\xe9\xa7\xca\x05\x63\x56\x85\x55\x2e\x90\x8a\x9a\x6a\x96\xb1\xf7\x3e\x7e\xce\x9c\xf0\x2a\xbd\xfc\x5a\x39\x46\x0a\x4b\x63\xff\xb6\x1c\x85\xf9\xdb\x83\x45\xcf\x01\x77\x0e\xcd\x78\x4d\x78\x83\xfe\x2a\x98\xb4\xc1\x79\x3d\x5f\xf7\x5b\x6f\x2f\x47\xd3\x41\x0d\x7a\xc5\xd4\xd5\x34\x98\xfe\x09\xa0\x21\x25\x64\xe9\x19\x74\x0d\x58\xe8\xaa\x12\xce\x98\xd2\xa1\x6a\x20\xd0\x3a\x9a\x1a\xf2\xd1\x9c\xcf\x6e\x3f\x81\xc9\x11\x52\x19\xac\x42\x0b\xa6\x73\xb1\x27\x30\x1b\x86\x52\x13\x46\x48\x56\xb3\x7c\x24\x69\x18\x5a\x1b\x69\x9c\x39\x55\xdf\xe7\x04\xcc\xc6\x16\xca\x31\x50\x4c\x6c\xe8\x06\x14\xba\xf6\x3c\x73\x32\x30\x39\xbc\xe0\x24\xf3\x90\x21\x18\xcf\x13\x39\x84\x18\xf3\xd0\xe1\x04\x6d\xfc\x34\xb1\xa3\x32\xde\x43\x06\x92\xc7\x09\xdd\xd9\xde\x4e\x60\xeb\x1a\xe8\x3c\xed\x27\x70\x33\x4c\x5c\xfa\xcb\x1c\x4d\x98\x4b\x06\xd7\xad\x08\x8e\xf0\xda\x28\xe9\xcd\x51\x80\x09\x02\xa1\xb6\x70\x8e\xee\x04\xce\x42\x98\x21\x79\x68\x93\x08\xce\x61\x5a\x48\xdb\xf1\x68\xa3\xde\xe0\xa0\x6c\xe0\xcb\xf8\x40\xe7\x3c\x24\x8b\x25\x07\xbc\x18\xd8\x1d\x9d\x17\x13\x84\x1d\xf0\x92\xcf\xde\x1c\xcd\x97\xd7\x33\x35\xfe\x85\x3d\x86\xfb\xc3\x97\x44\xcb\x2a\x2b\xb5\x0d\x2a\xae\x6d\x04\x6d\xea\x5c\xc1\xe4\xe9\x7a\xae\x70\x02\x0f\x37\x0e\xb2\x2b\xda\x42\x26\x3d\xbe\xc8\x17\x65\xdb\x83\xbe\xa0\x53\x5b\x42\x83\xb9\xe0\x61\xdd\x00\x4e\x81\x4e\x3c\x41\xc9\xe7\x95\xdd\x19\x73\x63\xee\x85\x74\x56\xb2\x44\x9a\x99\x02\x44\x50\x14\x42\x6e\x3d\x05\x93\x98\x8c\xca\x65\x3f\xcc\x16\xc8\xcb\xc0\xb8\xf0\xf1\x00\x79\x94\xbf\xae\x5f\x48\x27\xb8\x38\x4c\xdd\x6e\xa1\xd0\x52\x61\x82\x34\x55\xd3\xe7\x2d\x16\xb4\xa0\xda\x97\x35\x3a\x61\x30\x60\x37\x48\x85\x52\x9b\x3d\x61\xf0\x26\x13\xd0\xfb\x4f\xd6\xd2\xc6\xae\x14\xea\x86\x48\xc5\xa0\xe5\x61\xb1\x54\x22\x50\x28\xaa\x1b\xbd\x77\x58\x22\xa4\xbf\xf5\xa0\x6a\x74\x62\xd2\x63\x67\x41\xb6\x6a\xaf\xa3\x5d\x1c\x16\xd2\x7d\x2c\x6d\xb8\x65\xb7\xfd\xbb\x6b\x61\xa5\xd9\x47\x87\xd7\xb6\x90\x84\xb6\x04\x6a\xae\xce\x99\x44\x1b\xde\xb5\x1e\x67\xa1\xf5\x51\x0b\x4f\x8e\x17\x21\xda\x64\xd5\xc2\x5f\xe9\x08\xaf\x24\x98\x67\xd2\x7d\x30\xb5\xf0\x46\x06\x79\x4a\xdc\xf6\xde\xd8\x68\x1b\x6d\xa0\x1c\x79\xae\xd7\xc1\xb0\x37\x48\xb2\xac\xd9\x0e\x23\x6c\x85\x29\x14\x4c\xd8\xc7\xbe\x15\xde\x76\x5b\x63\xb3\xe6\x0d\x52\xd1\xae\x4e\xfd\xbd\xb7\xa9\x5a\x10\x7f\xdb\xbb\xba\xc9\xe7\xd5\xdb\x69\x02\xba\x42\x52\xbe\x66\xd2\x90\xd5\x92\xa0\x50\xdf\x6f\x13\x38\x30\xb5\x2d\xc7\x09\x1c\x15\xae\x89\x65\xc8\x78\x48\x3f\x6c\x93\x09\x92\x51\x09\xc1\x89\x54\x3e\x7e\x82\x52\x8b\x9a\x6b\xe8\x1d\x6c\x90\x08\x16\x32\x67\x83\xa7\x09\x41\xaf\x85\x0d\xec\x4d\x86\x26\x31\xfb\xfb\x54\xd5\x7c\x9a\x30\xa1\xc7\xb0\x43\x5e\x6d\xdf\x69\xb2\x7f\x95\x41\x70\x6a\x35\xe6\x1c\x13\x77\x8e\x78\xcd\xc3\xbc\xaa\x66\x38\x4f\xae\x62\x17\x89\x91\x79\x9e\x87\x77\xbb\xb3\x32\xcf\xe5\x02\x09\xdb\x02\x7b\xe0\xba\x5a\x88\xc3\x5d\xb7\xf9\x1e\x2a\x7b\xd6\xcc\xe2\xcf\xfb\x68\x0f\x69\xeb\x77\x35\x66\xda\x4b\x9f\x87\x47\x26\xa2\x79\x67\xde\xdd\xd9\x05\x92\x19\x76\x6d\xe7\xcb\x8a\xe0\xca\x3a\x21\xd4\x72\xf4\x7b\x8b\x49\x5c\x68\x39\x5c\xdb\x44\xb3\x4a\xe3\x44\x9b\x68\x71\x72\xd6\xcc\x47\x87\x43\xc3\x39\x12\xc3\x6b\x07\x39\xf7\xd5\xd4\x6c\x4f\x7c\x6c\x0b\x7b\x3b\x4d\x14\x1d\x04\xd4\x18\x4a\x6a\x8b\x77\xa2\x64\x9e\x7b\xe3\xf0\xc7\x13\xdd\xee\xaf\x72\xa0\xad\x9a\x13\xc9\x17\x2f\xd0\xe9\x15\xe4\xf4\x2b\xfd\x50\x90\x35\xfe\xfc\xf4\xd4\xe7\xdd\xf1\xf2\x5a\xcf\x85\x63\x6b\xca\x11\x70\x9e\x6b\x68\x44\xe1\x02\xae\x4d\xb8\xab\xa8\xf2\xdf\x2a\x76\x9f\xe8\xd6\xe8\xc4\xe2\x9e\xfb\x1b\xd9\x15\x70\x96\xfb\x48\x98\x73\xf1\x7b\x8c\x7d\x6f\x4c\x31\xf1\x73\xef\x3e\x81\xc5\x15\x2e\xdd\x11\x4f\xe9\x7e\x96\x4c\x09\xc1\xaa\xb2\xa2\xd2\x6b\x3b\xb2\xbb\xbc\xd9\x26\x94\x7d\xbc\x23\x59\x95\x0a\xb5\xaf\x4e\x5c\x56\x14\x6c\xe1\x9b\x3b\xef\xbc\x9a\xc9\xf1\xd6\xda\x19\x9d\x6b\x8d\x6a\xc8\x85\xe3\x2d\xd5\x78\xd9\x5f\x1e\x0a\xf5\x45\xd9\x2c\x5d\x67\x2c\x5c\x4b\x57\xb7\x78\x1f\x53\xb5\xee\x81\x48\x0b\x26\x45\x3e\x62\xa2\x31\xb7\x35\x05\x35\x43\x69\x06\xee\xc8\x47\xcd\x9c\xd4\x98\xde\xa1\xf3\xce\x17\x80\xf2\x8f\x70\x57\x7a\x45\x88\x6a\xc0\x83\xfe\x95\x77\x6e\x24\xe7\xb2\xe2\x64\x30\xa9\x07\x00\x21\x0a\xfd\x50\x6e\x03\xde\x95\x47\x23\x07\x92\xc8\x34\x38\xa7\x78\x56\xa6\x1e\x46\xd3\x90\x0c\x71\x6e\x40\x69\x7c\xb0\x86\x54\x50\xaf\x81\xb4\xf4\xa0\x31\xb0\xe5\x7c\x34\x37\x70\xbd\x29\xae\x07\x26\xd0\xe7\xc4\x27\xbd\xea\xda\xdc\x77\x1f\x58\x49\x35\x17\x34\x6a\x02\x6d\x1d\x85\x01\xa8\xf4\x8a\x6e\x87\x89\x42\xb6\x67\x43\x18\xa0\xcd\x94\x5e\x05\x2e\xa1\x02\x8f\xb7\xb1\x39\xf5\x4a\xb7\x95\xac\x7a\x69\x72\x6a\x43\xa2\xb2\x0b\x40\xa1\xee\x25\x35\x6d\xe4\x2e\x90\xca\x8a\x50\xc6\xf6\xd7\x0e\x21\x50\x58\x86\xab\xc8\x67\x4b\xb8\xb5\x67\x1d\x69\x9b\xf9\xd0\x22\x0f\x05\xc5\xeb\x8b\xdf\x54\xae\x5a\x7e\xeb\x5f\xc4\x46\x9c\x00\x0c\xe3\x6b\x36\x7b\x9e\xa6\xbf\x83\x98\x67\x44\x0f\x79\xa5\xb0\x77\x65\x57\xfd\x24\x9f\xdc\xde\xc0\xde\x63\x82\x05\xbb\xc8\x47\x15\xd9\x73\xd6\xeb\x41\xcb\x2a\xd5\xa8\x40\x1c\xb2\xac\x96\x7f\x78\xf0\x89\x19\x4c\xc2\xcb\xbf\xdc\x09\x12\xd8\x4f\xcc\x44\x10\xfe\x65\x1a\xfe\x8a\x05\xce\xfc\xcb\x45\x87\xe1\xd3\xdb\xe7\x84\x9f\x14\x17\xe4\xb4\xe0\xbf\x5c\x4a\x1b\x7d\x7a\x13\x6d\x9f\x07\xf4\xc6\x6b\x68\x3b\xfc\x91\x57\xa5\x11\xe0\xce\xaf\x82\xc9\x93\x00\xe2\xd2\x0d\x17\x34\x61\x5a\x70\x49\x5c\x63\xf3\x44\x9a\x43\x40\x2d\x9b\x33\x97\x6a\x88\xc7\xe4\xb0\x25\x54\x0e\x2e\x0a\xc3\xcc\x49\xcb\xb9\x52\x14\x1c\xdf\xf5\xa5\x74\x9a\xbe\x93\x6a\xfd\x9d\xd4\x98\xef\xa4\x88\xdf\x49\xe7\xf9\x3b\xe9\xb2\x7c\x27\x5d\xd7\xef\xa4\x44\xdf\x49\xdf\xde\xbe\x93\x5a\xfb\x9d\xd4\xb9\xef\xa4\xde\x7f\x27\x0d\xe1\x3b\x29\xf3\x77\xd2\x18\xbf\x93\xfe\xfd\xfb\x9d\x34\xa5\xef\xa4\x39\x7f\x27\x2d\xe5\x3b\x69\xad\xdf\x49\xb7\xed\x3b\xe9\xe5\xf2\x9d\xf4\x7a\xfd\x4e\xba\xef\xdf\x49\x6f\xb7\x21\x8d\x7b\x8b\x31\xb4\x03\xf2\x79\x6c\x9a\xc1\xa3\x90\x0b\x95\x5a\x70\xb0\x37\x30\x2b\xe7\x06\xc8\xe4\xc4\xd6\xb1\x1f\x12\xcd\xa3\xd5\x89\xd9\x9e\x73\x8d\x98\x02\x5e\xfa\x03\x35\x17\xf6\xb9\xde\x0f\x46\xbd\x4f\x98\x94\xe7\x30\x50\x84\x01\x72\x7b\x82\x60\xdb\x51\x60\x20\xb1\x81\xee\x75\x8d\xc4\x0c\x81\xba\x67\x35\xfd\xd0\xeb\x4a\xda\x68\x88\x54\x3a\xa6\x3c\x19\x94\x10\xfa\xf9\x87\x6c\x7f\x83\x6e\xa6\xa0\xe0\xef\xfb\x0b\x8d\x7c\x2b\x2c\x58\x47\xaf\x18\x16\x52\x9b\xb2\x09\x0d\x09\xce\xe8\x3a\xc9\x43\xe1\x0d\x03\x26\x76\xbc\xd0\x0e\x43\x39\x6b\x48\xcb\x1d\x07\x1f\x30\xb8\x8d\x02\x4b\x16\xa3\x5a\x0e\xb9\xa4\x6a\x0b\xa7\xbd\x0f\x61\x6b\xc9\x09\x01\x3a\xa7\xf7\x23\xcb\x50\x74\xec\x21\x0b\x72\x6b\x5d\x53\x6c\x98\x48\x8c\x64\x04\xc1\x85\x23\x78\x34\x24\x98\x3f\x98\x16\x93\x9b\x37\xb4\x57\xe8\x2d\x0e\x17\x28\xfd\x7c\x30\xb6\x34\x78\x62\x58\xe7\x57\x70\xc3\x10\xac\x33\x24\xbd\x4a\x28\xd2\x46\xc8\xba\x1d\xd1\xe0\x0e\xb2\xca\xb4\x4f\x90\xf1\x03\xa3\x67\x3e\x3e\xb2\x04\x4e\x74\xc6\x92\x53\x1d\x76\xe7\x37\x4b\x6a\x35\xe3\x65\x1e\x28\x48\xb8\xa2\xed\x38\x4c\x85\x95\xd7\x40\xbb\x5a\xb9\xa0\xeb\xe3\xf3\x14\xa0\x60\x4f\x69\x78\x77\x30\x04\x52\x37\x2a\x5b\x50\x5b\x46\xdf\x95\x73\x81\x4d\x18\xca\xa6\xdd\x8f\x3e\x73\x5d\x04\xfe\xf2\xd6\xf3\x30\x26\xd5\x38\xec\xb4\x25\x59\x82\x32\x7f\xa8\xa6\xba\x23\xe4\x8e\x98\x51\xf5\xb3\x55\xe3\x3d\x4b\xf4\x91\xf5\xdc\x79\x20\x50\x2b\xb1\x57\x12\x68\x6a\x37\xb0\xb3\xb0\x09\x82\x6a\x29\x96\x0e\x77\xa4\xe7\xe3\xf9\xb3\x9e\x4f\xf8\xa1\x7f\x99\x41\xe1\xc4\xb6\x02\x51\x73\x60\x19\xfb\x09\x0d\xd8\x65\xde\x2e\x7f\x9b\x25\xd0\x54\x99\x8f\x66\x48\x34\x55\xfa\xbd\x63\x36\x34\xf5\x61\x0e\x70\x11\xbc\x84\x7c\xa6\x72\x42\xab\x28\x6c\x98\xdb\xd9\x86\x16\x0a\x48\x18\x29\x9b\x16\x2d\xab\x45\x62\xfe\x2e\xe1\x98\x18\x6d\x51\x36\x75\xba\xf0\x19\xf1\x84\x0e\x6d\x49\x64\x79\xa3\xbe\x6e\xd0\xb5\x9d\x9f\xa3\xab\x39\x0f\x46\x28\x7d\x6b\x34\x43\x3a\x4f\x33\xa4\x15\xc2\x42\x7d\x68\x6e\xc3\x7b\xe8\x86\x1e\x53\xe6\xa0\x52\xed\xcf\x26\xd2\x62\x3b\xec\x53\x85\x89\x4b\xfb\xee\xac\x39\x15\x51\xca\x82\x32\x9a\x28\x17\x54\x07\xd6\x6c\x3d\x15\xa0\x44\x16\x5a\x74\x7b\x5e\xd2\x09\x0b\xae\xe1\x01\x3f\x63\x55\x5f\x3b\xad\x97\xfe\xf5\x35\x71\xcf\xa8\xb5\x08\xa7\x8d\xa6\x26\x8e\xb0\xf4\x59\x3c\x18\x19\x3c\x1b\x74\xb6\x6d\x41\xdc\x9a\xa3\x96\x15\x7b\x84\xfe\x78\x95\x49\x99\xe9\x3e\xf9\xd7\x98\x30\x67\xb5\xed\x1b\xdf\xda\x33\x3b\x66\xc5\x41\x70\xb6\x78\xab\xf9\x39\x98\x91\xc7\x98\x21\xf0\xc4\xdd\x4a\x33\xe4\xa2\x2e\x9c\x4c\x7f\xba\x25\xf5\xba\x60\x32\x7c\x09\x8e\xc1\x8c\x3d\x32\xa3\xa7\x40\xb9\x1c\xa9\x80\x99\x0c\xb8\x0c\x12\x0c\xb5\x51\xb4\xb8\x5c\x56\x41\xa3\x0c\x08\x5a\x1e\xdb\x75\x26\x27\x58\x87\x8b\xba\xad\xe8\xf0\xd6\xbc\xc2\x4c\x01\x82\x26\x70\x2a\x93\xaf\x0e\xee\xc1\x6a\x13\x64\xde\xfa\x0a\x9e\x49\xc6\xad\x21\x8d\x2c\x4d\xd7\xb8\x1e\x91\x83\xc2\x94\x38\x0d\xb6\x63\x28\x2d\x7c\xfe\x94\x7e\x9d\x1d\x5f\x50\xf0\x73\x82\xd8\x7c\xdd\x11\x81\x9c\x47\x3f\xa7\xb6\xcf\x55\x4c\xac\x0b\x0f\x85\x8c\xd7\x31\x06\x4e\xe5\xc2\xd9\x43\x2a\x1a\xd2\x39\x5e\x4e\x2d\xe0\xf8\xf5\xdf\x21\xbe\x5e\x70\x3a\xcc\x94\xa0\xca\x1a\x6d\x7d\x24\x44\x35\x4f\x1f\x32\x23\x83\xc7\x6c\xbf\x60\x1f\x1f\xf0\x61\xe8\x5f\x27\x94\x3b\x9b\x9b\x01\xff\x17\x5f\xd6\xdd\x17\xfc\xdf\xef\x82\x63\x28\x77\x46\x5b\xde\xaa\x30\xbb\xf7\x5e\x0f\x93\xff\x3b\xee\xc0\x66\x04\x6b\x42\x5d\x56\x28\x19\x62\x94\xa5\x73\x57\xb9\x38\x08\x77\x95\xbc\x06\x20\x97\x8f\x0d\x33\x67\x70\xd0\x07\xbf\x80\xc7\x1f\x4f\x4f\x62\xbb\x05\x82\x21\x0b\x31\xf6\x68\x76\x41\x48\x3a\xc1\x5c\xce\x35\x37\x22\x4b\x7c\x94\xd8\xf3\x59\x16\xde\x69\x91\x03\x10\x5c\xe4\x14\xda\x1a\x5f\x30\x42\xa1\xa2\xe4\x9d\x0d\x71\x2f\x98\x3c\x66\x19\xfb\x30\xd4\x82\x45\xb5\x58\x89\x6e\x63\xf4\xc2\x79\x34\xd3\xe3\x77\xde\x65\x83\xa9\x02\x5f\xee\x02\xcf\xa9\x85\x11\x88\x49\x6d\x94\xa9\xc8\x8a\xbc\x77\x79\x97\x52\x5e\xbf\x10\x67\x64\xb5\xa2\x8b\xf7\xde\x5a\x60\xfc\xcf\x0c\x1e\xda\x00\x26\x53\xd0\xec\xdf\x87\x7c\xe7\xb8\xfd\xce\x73\xe0\x27\x4e\xcb\x4a\x81\x64\xef\x2f\x58\xe2\xde\x3f\xf6\xae\x92\x48\xaf\x7f\xab\x70\xbc\xfb\x40\xdf\xbb\x59\xe1\xd6\x8f\xc2\x45\x60\xd3\x18\x78\x1f\x25\xcd\x65\xa2\xa5\xc9\x28\x75\xa3\x4b\xdb\x6e\xa8\x45\x47\xda\x0e\xb6\x98\x78\xa6\xd2\x29\x9e\x64\xa3\xf7\x12\xc6\xc2\x13\x87\x65\x24\x3d\x17\x66\xb3\x56\xcf\x69\x84\xec\x42\xc7\xc4\x2d\xd3\xdb\x74\x79\x71\xe2\x81\x21\xbd\x93\x1e\x12\x1e\xca\x8b\xcb\x55\x8f\x90\x7b\xe1\xb4\xc0\x6d\xf4\x9b\x60\x06\x0f\xbd\x93\x0a\xc9\x48\x64\xde\x97\x46\x25\x83\x85\x23\x16\x4c\x79\xaa\xa9\xfb\xd3\x55\xd6\x9b\x86\xb8\x41\x78\xcf\x8c\xad\x10\xe3\x9e\x39\xb4\x4f\x5d\x21\x19\xd5\xbe\xd5\xf3\x44\xa7\x15\x36\x8c\x35\xe7\x43\xf3\x02\x44\xb9\xa6\xb9\x93\xc6\xf3\x46\xa8\x41\x02\xe5\xe6\x03\xdf\x39\x7d\xac\x2b\x82\x9b\xa8\xb9\x1a\x69\x2e\x90\x4b\xea\x27\xfe\xca\x1e\x61\x96\x33\xc6\x91\xc5\x73\xb1\x8d\xe3\xf7\x48\xba\xd4\x34\x18\x61\x99\x21\xbc\xd1\x78\xa0\xa1\xda\x0f\x88\x66\xe5\x52\xb8\x1e\xed\x9a\xd1\xd0\xed\x9d\xf0\xe4\xba\x65\x56\xbe\x70\xc4\xa3\x59\x04\xad\xc4\x5e\x43\x01\x97\x23\xf8\x33\x8a\xe0\x56\x38\x2d\x8a\x63\x9b\xcb\x35\x23\x45\x80\xb3\xbd\x9d\xd6\x2a\x98\x75\x29\x1c\x22\x8f\x0c\xef\x5a\x3d\x04\x6e\xf8\x3b\x77\xcc\x27\x26\x5f\x6b\x30\xf8\x36\x3e\x6f\x33\x2a\x17\x1e\x33\xb8\xee\x14\xd5\x0d\x02\xdf\xe0\xec\xf1\x44\x9a\xcf\x09\x4f\xb4\x80\x59\xb0\xe4\x0b\x27\xd7\x9d\x08\x2d\xa9\xaa\xeb\xc4\xd7\x66\x4c\x12\xcb\x6c\x82\xa3\xbb\xf5\xc9\xf1\x86\x9d\x4c\xfc\x81\xaa\x9d\xba\xeb\xfe\x43\xfa\x0f\x74\x5b\x38\x8f\x0c\x19\xfb\x23\xfd\x4f\x77\xe2\x2d\xc9\x53\xc0\xd0\x3f\x7e\xa4\xc3\x66\x72\x6d\x9d\x93\x8f\x5c\x30\x14\xea\xf9\x3d\x0a\x9a\xca\x2e\xd3\xdc\x3b\x6a\x9b\xd6\xe2\xde\xbf\x28\xe8\x84\x90\xf1\x72\xb9\x1c\x6e\xa7\x8f\x26\x88\x8b\x58\x39\x9e\xb3\x3c\x32\xa3\x2d\x99\x68\xf4\x37\xf3\x03\x50\x3e\x58\xc9\xab\x0a\xc7\xc3\x99\xe4\xf8\x8e\x62\x94\x50\x30\xcd\x98\x30\xe8\xee\x51\x85\x71\x55\x40\x83\x10\xa8\x76\xd4\x18\x29\xdb\x96\xbc\x6c\x4d\x2e\x1c\xc6\xa2\xa2\x1c\x21\xd0\x0e\xaa\x15\xc6\xd4\xe4\x20\xe8\x76\x96\x53\x51\x1e\xae\x87\x43\xa7\xdb\x9a\x4b\x42\xc7\xad\x7e\xf3\xf6\x36\xc1\x94\x2d\xb7\x6c\xd1\x1b\x4f\x54\xef\x9e\xff\xad\x7a\x7b\x20\xc9\xb7\x9a\x8b\x25\xe7\x64\x3d\x91\x2c\x05\x61\x0c\x0f\x22\x76\xb6\x30\x4d\xe0\x60\x55\x09\x8d\x9a\x12\x68\x74\x58\xfa\x56\xb2\x30\x35\x33\x4d\x14\x26\x09\xc7\x52\xbd\xb3\xb4\xec\xd9\x9f\x16\x3e\xf0\xc4\x0f\xaa\x56\x4d\xfb\x47\xdb\x59\xe5\x71\x81\xb9\x43\xd3\x3b\x97\x17\x0a\xca\x97\xfc\xc8\xf4\xe4\x3e\x90\xfc\x51\xba\xab\x09\xb1\xc1\xc6\x8f\xdc\xc8\x7a\x2d\x9f\xde\xdb\xb3\x99\x6a\x2b\xd3\x23\xb7\x50\xb0\x3c\xcf\x1f\x58\x65\x3c\x67\xa1\x55\x65\x64\xe5\x58\xf0\x13\x64\x9e\xcb\xb0\x85\xa7\xa0\x32\xf8\xae\x97\x20\xd6\xdb\x81\xd8\x2c\xdc\x0e\x4b\xdc\xd2\xb1\x66\x2d\x26\xbd\xf6\x00\xc5\xe2\xd6\x7a\x95\xe3\xd2\xe2\x2e\xa8\x30\x4b\x30\x21\x7e\xbf\xc1\xaa\xaa\xf5\x51\xea\xfd\x4a\xde\x0f\xae\x47\xa5\x15\x92\xe5\x96\x95\xb4\x14\xf8\x28\x9d\xb4\xb6\x3d\x42\x14\xa1\x5e\x9e\x9e\xff\x9c\xb5\x6e\xed\xa9\xf6\xf2\x95\xb4\x23\x3b\x2c\xfd\x2b\x29\x70\x46\xd3\xdb\x81\x16\x0e\x9c\xd7\xee\x59\xed\x3f\x19\x48\x3b\x72\x6b\x59\x4d\x23\xfc\xb0\xec\xab\xbd\x97\x12\x2c\x07\x5b\x93\x00\xcc\xfe\xf4\x6d\x85\xcc\x13\xf4\x4d\x6a\x13\xe4\x20\x81\xbc\x82\xad\xb0\x63\x3f\x41\x32\xad\x8f\x84\x66\x2f\x6d\xe6\x46\x3f\x0e\x5a\x40\xf8\x4e\x06\xe3\x68\x1b\x18\x5b\x28\x9e\x3f\x20\x76\x07\xb1\x70\x54\xaf\xca\x1d\xf6\x91\x00\x41\xaf\xa8\xad\x2a\x97\x9e\x1a\xfa\x94\xb7\xe3\x7b\x0e\xed\x53\xd2\xae\x3f\xb8\xa6\x91\x81\xfb\x7f\x12\x73\xc2\xce\x74\x1d\xed\xab\x49\xe4\xdc\x78\x3b\xe2\xc4\x7b\x0f\x76\x06\x63\x01\x97\x74\xcb\x8e\x38\x4c\xed\xa8\x6b\x53\xe4\xb0\xb0\x32\xa0\xd7\xb6\x7a\x9c\xf8\x94\x3f\x2f\xbd\x19\xc9\xf0\xcc\x53\xef\xe1\x3c\xa5\x93\xa3\x52\x9c\xa0\xe1\x9e\x53\x10\x8b\xcc\x74\x1d\x0e\xc4\xd1\x46\xa0\x22\xe8\x15\x87\x90\xc2\x92\x98\xbd\x41\xcd\xa9\x61\x5f\x32\x08\xf9\x8e\xca\x9d\x55\x0b\xe7\x9a\x5d\x5d\x5a\x20\xd4\xf7\xe2\xc3\xce\x10\x6f\x2c\xa7\x7e\x10\x9c\xb6\x9e\x5c\x1e\x29\x67\x57\xb5\xdd\xa7\xea\xdc\x99\xf8\xe4\xaa\x65\x72\x7a\xcc\x64\x4b\x65\x6b\xe1\x16\x0e\x0f\x41\xaa\xab\xd7\xda\xee\x29\x78\x58\x48\x1b\xc2\x72\x5e\xfe\x0a\x11\xb0\x90\x7e\xf0\x39\x4d\xe5\xf3\x4d\x06\xe1\xc5\xc4\x8b\x04\x52\x82\xba\x3c\x84\x71\x02\xe8\x43\x21\xd4\x56\xf9\x86\x78\x4e\xec\xf0\x5e\xb5\xf0\x90\x26\x87\xbd\x2c\xdd\x49\x32\x3d\xfe\x94\x85\xd5\xcb\x5e\xb2\xa0\x04\xaa\x8d\xdb\x18\x88\x76\xf4\x59\x12\xe4\x73\x8e\xa0\x51\x88\xb2\xd0\x94\x39\x08\x4a\xf6\x70\x55\x10\x63\x82\x51\xfd\xf4\x70\xbd\x5e\xc9\x73\x2f\x5d\xb7\x0b\x06\x8a\x6e\xca\x60\xc2\x0d\xee\xe5\x83\xc6\xef\xc5\x62\x5b\xee\x8a\x06\x2d\xdf\x43\x6d\x8f\x46\x19\x96\x1d\xb7\xf4\x85\x23\x8c\x9b\xe1\x24\x00\xe7\x41\x89\x74\xab\x75\x40\xde\xa7\xba\x1f\x8a\x79\x78\xbb\x01\xf7\x3c\x9a\x7a\x93\x03\x25\xe4\xea\xc6\xfb\x16\x88\xec\x28\xab\xff\x7e\x75\xba\x80\x73\xfc\x9e\xad\x92\xf3\xef\x9d\x9f\xe1\xcd\x50\x07\x68\x9e\x66\x8c\xa9\x65\xa4\xda\xd6\xf6\x64\xf9\x26\x88\xa6\xbf\x89\x6c\xe2\xbe\x37\xee\x63\x14\xd6\x0d\x76\xff\x22\x41\x72\xeb\x21\xe0\x02\x7e\xe8\x27\x9e\xea\x46\xef\xca\xa9\x70\x6a\x7e\xb3\x75\x6e\x7b\x11\x49\x9e\x6a\xb7\x3a\xe4\x83\xe1\xec\xcd\xc9\xb3\xb6\x35\xd6\x7e\x48\x7b\x36\x33\xa7\x0b\x0f\x45\x13\x08\xd5\x5c\xcb\xc4\x12\xd9\xb7\x7e\x98\x02\xdf\x56\x7b\xa8\xcb\x62\xbb\xb5\x60\xb8\x39\x8e\xae\x12\xb0\xd5\xb8\x54\x74\xa0\xc5\xd3\x7e\x48\xfb\x34\xf1\xef\xdf\xbf\x7f\xf7\xcd\xda\xc8\xd2\x6f\xdc\xb4\xf6\xed\x76\xeb\xed\x9c\x68\x1e\x2d\x01\x74\x01\x53\x35\x9d\xde\x95\x61\x5b\x0f\xd3\x08\xa3\x5f\x3b\xe9\x2b\xa5\x66\xed\x50\xcd\xa0\x5b\xa2\xf0\x30\x47\xcd\xa4\x57\x70\x0e\x5a\x65\x42\xf4\xf6\xa9\xae\x02\x0b\xfd\x3e\x97\x58\xdf\x03\x40\xbf\x47\x07\xbb\x80\x9f\x46\x06\x71\x24\xce\xc2\x4b\x2b\xd8\x06\x08\xac\xd7\x2c\x67\x53\x9b\xd6\x80\x68\x0a\x67\x74\xce\xef\x0d\xab\xce\xc7\xe5\x01\x39\x00\x3b\x20\xe8\x11\x6d\xc0\x0d\x22\x8b\x93\xb6\xfd\xc1\x4b\x4e\xdc\x62\xc2\x83\x2c\x30\xcf\x60\x72\xbf\xcd\x24\x51\x62\x20\x4b\xca\xf7\x75\x16\xe8\x16\x68\x0f\xbc\x2d\x9c\xd8\x3c\x94\xe8\x02\x6f\x02\xa8\xd5\x4a\xaa\x60\x3f\x52\x42\x9d\x2a\x39\x83\xad\xf8\xd4\xac\x1c\xea\x0c\xbd\x58\xc6\x8a\xd5\xf3\xb3\x62\xc5\x77\xf2\xd7\x17\x54\xaa\x8d\xf8\xf3\x41\xf4\xe7\x2e\x9a\x32\xe6\x4c\x1c\xe2\x9a\x20\x8f\x89\x65\x43\x30\x61\x21\x0b\xeb\xa0\x25\xe8\xed\xab\xad\x39\x69\x9e\x25\xa8\x4d\x87\xb8\x5d\x65\x78\xb9\xe0\xf4\x40\xde\xd3\xa2\x6d\x25\x71\xfb\xd6\x30\x62\xda\x01\x3b\x1e\x60\xd3\xe0\xac\x34\xca\xd8\x83\xa6\x50\xc8\x3f\x3e\xe2\xed\xf4\xa0\x5e\xd0\x47\x07\x05\xdf\x51\xfd\x47\x40\xd3\xa9\xb2\x65\x4c\x84\x59\xb0\x1d\x07\xb7\x5f\x58\xa2\x33\x39\x2b\x99\x59\xb1\x0b\xcd\xf8\x1c\x4b\xe6\x06\xad\x38\x81\xad\xee\x9c\x63\xef\x23\x17\x64\xbd\x72\x30\x89\x33\x35\x1f\xcc\x97\xd0\x67\x76\x3e\xf1\x0d\xfa\x07\x46\x20\xd3\x5f\xf7\x7e\xa5\x20\x82\xdb\x28\xdd\xf3\x73\x51\xab\x7b\xcd\x34\xd5\x53\xc4\xe4\x1f\x26\x3f\x62\x1a\x19\x87\x88\x25\xf1\x66\xe1\xb1\xd6\xdc\x72\x50\x02\x20\xdb\xee\x6a\x94\xa5\xa2\x57\x0c\x06\x33\x2d\x3d\x7f\x35\xe2\xb4\x2c\xfe\x1f\x06\xc7\xb9\xfc\xf2\xba\x0e\xa2\x3a\x28\x82\x7a\x6f\xa7\x48\xb7\x1b\x7c\xa8\xa2\x77\x4e\xd9\x35\x8f\x6c\x58\xb4\x8f\xd8\x26\xda\x8f\x83\x17\xa4\x62\x5b\x48\x32\x48\xf4\x81\x6c\x03\xfc\x91\x4d\x46\x0b\xcf\x5d\x40\xd9\xde\xc0\x72\x80\x4e\xda\x1a\x61\x77\xb8\x74\x90\x1f\x39\x1e\x97\xea\x22\xc7\x7c\x9e\xb9\x06\xd3\x7c\x90\xd0\x75\x81\x5d\xaf\x34\x3c\x94\xa0\x82\xb5\x4e\xa3\xe4\xbe\xc0\x3f\x8c\x1a\x1a\x78\x3f\x98\xa3\x64\x20\x24\x2b\xbd\xb6\x80\xbc\xad\xac\xc6\xc1\x76\xe7\xed\x9d\x9e\x39\x81\x91\x63\xe0\xfe\xcc\x42\xcb\x3b\xb1\x9a\xe7\xa7\x3f\x4f\xe7\x3e\x1f\x9d\x65\x1d\x98\xfe\x05\x29\x70\x18\x81\x56\x23\xa2\xc3\x70\x50\x25\x08\x5a\x60\xd3\xa6\x9f\x6f\x26\xc1\xe6\xb0\x07\x32\xed\xe1\x44\x1a\xc7\x7b\x12\xe9\x15\xb3\xa5\x9e\x00\x1b\x3c\x01\xc5\x4d\x8f\x5b\x0a\x2a\x0f\xc2\x00\x7a\xb3\x61\x1a\x86\x49\x6c\xaa\xf6\x87\x6c\xa3\x05\x0a\xa7\xf6\xca\xc4\x73\xdf\xdd\xe0\x8e\x6b\x39\x59\x22\xeb\x98\xfa\x55\xa7\xa3\xb0\xde\xf2\x59\x77\x54\x10\x13\x87\x3e\x93\x89\xb3\xa7\x63\x8f\xb6\x2b\x35\x54\x6a\xb1\xa4\x7a\x79\x41\xe6\xb4\x0f\x21\xc3\xeb\x9f\xa6\x53\x83\x5e\xcf\xed\x44\x3a\xc5\x9a\xf4\x0a\x19\xdb\x22\x0c\x9c\xfa\x95\x82\xbf\x15\x1c\x95\xbd\x63\x8b\x1b\x8f\xf5\xfd\xb7\x52\x51\xd9\x73\x3b\x7a\x16\x38\xfd\xbd\x60\x6e\x2b\x26\x41\xb0\x1d\x71\x0a\xbe\xed\x13\x20\x3c\x0a\x4b\x56\xf7\x7b\x77\xad\x0f\x61\x67\x0e\x47\x00\xfb\x4e\xca\xf2\x7a\xa7\x64\xe3\x27\x88\x64\x9a\x99\x7a\x5e\x65\x81\xce\xea\x7b\x57\x3e\x3d\xc1\x6d\xe2\x64\x41\x65\x2b\xb8\x65\x04\x32\x49\x4f\xcf\x4f\x6d\x61\x27\x04\x97\x50\xa0\x43\x76\x34\xce\x89\x84\x26\x50\xc6\xd1\x5e\xc4\xd8\xe2\x26\xc4\x61\x26\x74\xd8\x88\xf6\xa8\x97\x48\xa1\x5f\x90\xe9\x65\x78\x35\x6e\xee\x89\xcf\xe9\x4f\x87\xf2\x32\x1c\x50\xc2\xe8\x48\x83\xba\x17\xb2\x12\xf6\xdb\xbb\x67\x9a\x3d\x1f\x0c\xe7\x1e\xaf\xd0\x0d\x83\x34\x3e\xa6\x96\x20\x1f\x8c\xc8\x21\xd3\x26\xae\x33\x77\x38\x91\x70\xc3\x94\x5b\xa8\xa0\x4f\x69\x86\xad\xc7\xa0\xfd\x02\x65\xc0\xb4\xec\xef\x2b\x47\x1c\x19\xd4\x53\x62\x9e\x55\x3b\x55\x45\xb1\x9f\x7c\x65\x60\xd9\x54\x25\x98\x4d\xa3\x2d\x13\x21\x90\xae\x53\xe2\xaa\x7b\xcb\x4b\xe8\xd8\x87\x54\x05\x72\x95\xbc\xf2\xa5\x8b\x32\x5e\x7b\x7e\x23\xd5\x1c\x19\xc7\x4b\x7a\x61\x44\xc5\x96\x70\x32\xfd\xaa\xef\xc1\x6c\x99\x55\x07\x47\x8e\x30\xed\xed\x3e\xdd\x91\x35\xca\x60\x94\x80\x8e\xc3\x0b\x67\x30\xbb\xaa\x09\x1c\x74\xa9\x2b\x39\xa6\x7e\x7b\x2e\x43\x78\xe3\x8c\xad\x14\x51\xc7\x4d\xd9\x0c\xa1\x40\xc2\x70\x9e\xf6\xde\xce\x0b\xcd\x45\xf0\xb0\x50\xb8\x06\x3e\xea\x44\x19\x36\x2c\x75\xc2\xf7\xfd\xfb\x91\xd3\x16\x65\xd6\x10\x32\xcc\xbd\x3a\x91\x35\x94\xb3\xac\xbf\x53\xd6\x09\x31\x14\x16\x70\xd6\xd7\x52\xd6\x89\x62\xc1\x8e\x63\xfa\x7c\xf7\x8a\xc3\x03\x03\x83\x4e\x7b\x1c\xfa\xa8\x6b\xa2\xb2\x2b\xcd\x12\x3f\xdd\x8b\x1a\xed\x93\x65\x75\xcc\x93\x6a\xf7\x3b\xd4\x8a\xe3\xac\xcc\xe8\xc1\x95\x87\x66\x07\x92\x9d\x66\xf5\xa2\x9e\xee\x6d\x39\x65\x67\x4e\xfe\xce\xc8\xde\x0b\xe0\xca\xc8\x6d\x0f\x1e\xc0\x21\xcb\x51\x89\x82\x42\xaf\x9f\xe8\xf6\xc5\xc8\x78\x95\x33\x47\x90\x41\xb7\x39\xf2\x1b\x5b\xcc\x87\x38\x8e\x5a\x4e\x46\x3e\xd6\x72\x8f\xd1\x5b\x6f\xd7\x0c\xbc\x1f\xed\x1d\x1c\xdd\xdb\x05\x31\xdc\xef\xe3\xe6\x15\x96\xc2\x3e\x37\xb7\xf5\xa1\x82\xd1\xc0\xec\xfd\xe6\xe6\xc7\xe2\x86\x2c\xba\xc6\x19\x75\x73\xe2\x0f\x64\x1b\xd5\x03\x5d\x44\x1c\x76\x58\x68\x88\xfa\xc5\x73\x4f\x09\x06\x42\x3e\x57\xdb\xd8\xad\x80\x2d\x23\xa1\x09\x35\x1f\x38\x31\x93\x6f\xb7\x30\xfb\xa1\x9f\x49\x02\x9f\x2f\xca\x2d\x99\xc2\x9a\x38\x44\x09\x96\xd8\x74\xd5\x82\x4a\xc0\xaf\xac\xf8\xe1\x12\x85\xf7\x73\xe8\x17\x4c\x14\x6f\x7d\x50\x54\x70\x03\x57\xfb\x72\x10\xcf\x97\x2d\x85\xbb\xdd\x1c\x18\xcb\x8f\x0b\xd6\xa1\xc9\x65\x23\x54\x1b\xba\xf1\x59\x6e\x2e\xd9\xec\xd3\xd4\x0e\xa2\xec\x6c\xbf\x6f\x9c\x1d\xfb\xa3\x71\xb9\x5f\xa6\x6b\x16\xf4\x40\x8e\x7a\xc2\xbb\x15\xae\xe2\xca\x01\x0d\x65\xcd\x35\x0c\x74\x97\x03\x4c\xb9\xad\x9f\x80\x0b\xdf\xc0\x60\xba\xad\x30\xea\xdf\x39\x50\x6c\xc6\x6d\xff\x5f\x4e\x99\xc1\xca\x5f\x93\xb1\x5e\x49\xfd\x30\xa3\xfd\x5e\xa8\xe8\xaf\xfe\x50\xba\x38\x2e\xea\x7c\x51\xaa\xfa\x58\x66\x51\xcf\x8f\x7a\x07\xf3\xe5\x2b\xe6\x8f\xaf\x98\xaf\x5f\x31\x7f\x7e\xc5\xfc\xfd\xc8\xbc\x90\x59\xb0\x3c\x70\x3e\xac\xc3\xc6\xc9\xd8\xaf\x5f\x77\xaa\x24\x30\x1d\x7e\x67\x5e\x99\x82\x6c\xd6\x36\x45\xed\x06\x3f\x25\x54\x0b\xf4\xfb\xe3\x99\x5d\x88\x0e\xf2\x78\x38\x1b\x8c\x5c\x0c\x4e\x2d\xf9\xde\x0d\x50\x37\x0c\x94\xea\xd8\xee\xbc\x8d\x9b\x19\xd1\x81\x85\xb1\x33\x79\xaf\x37\xd9\xad\x86\xda\xea\x11\x56\x54\x3d\x97\x34\x88\x1b\xd8\x1a\x7b\xf4\x94\xe3\xa4\xa2\xa3\xd2\xf3\x7c\x42\x65\x8d\xa1\xd7\x4e\x72\x44\x34\x35\x2a\xbf\xf7\x74\x5f\x8e\x14\x5e\x9e\x9e\x7f\x09\x80\x6e\xb1\xec\x85\x6e\x23\x45\x22\xbe\x78\xb3\xf0\xfc\xe3\xe9\x81\xda\xe4\x24\x6c\x5f\x9a\x68\x0a\xdc\x9b\x05\x36\x1a\x20\x48\x34\x0b\x1e\xdf\x5e\x28\x81\x7b\x2f\xdf\xe7\xc2\xda\x96\x0b\xe5\x91\xb3\x6d\xb1\x82\x1a\x60\x24\x97\x84\xe0\x55\x31\x87\x2c\xf1\x6e\xe1\xf5\x77\x7f\xb0\x1a\x0c\xa5\x15\x6d\x1a\xdd\xf3\x8b\xf7\xfb\x11\xb9\x06\x93\xea\xf2\x1e\x7b\xb4\x4b\x3c\x84\xc1\xf4\xeb\x40\x8d\xf4\xe0\xfb\xbd\xbb\x46\x6d\x98\xfb\xb5\xd9\xbc\xa1\x58\xaa\x6f\x86\x0d\x0b\x0b\x68\xe2\xac\xf9\x72\x2a\xca\x43\xbb\xab\xd9\xd3\x58\x45\x25\xd4\xe7\x5c\x4f\x05\xa8\x45\xa5\x8d\x09\xa5\x30\xab\x5c\x48\x60\x51\x7e\x67\xad\x7d\x3e\x8b\x51\xdd\xc9\xf6\x1b\x2f\xfd\x92\xa9\x85\x87\x3b\x7f\x05\xcd\x35\xbd\x75\x65\x0c\xe5\x31\xed\x52\x30\x3a\x5e\xd9\xa8\x25\xe0\xad\xdf\xa8\x29\x78\x7d\xbc\x68\x50\x96\xd2\x4b\x7a\xed\xf1\x15\xd8\x54\xd6\x85\xf5\xa0\x11\x6a\xe1\x91\x40\xf7\x4e\x18\xbd\x0c\x36\xd6\x5e\x59\xb1\xf9\x85\x01\x62\x64\x32\x0b\xdb\x06\x00\x86\x33\x2b\xec\x8f\xce\x39\xaa\xa5\x17\x8c\xa4\xe9\x5e\x0e\xee\xf3\xa7\x9f\x6d\x08\xf3\x49\x5d\x46\xfd\xac\x70\x1c\x40\x53\x16\x59\xe1\xf8\x70\x9e\x9c\x35\x9f\x5a\x7e\x54\x1d\x40\xf2\xfe\xe2\x5e\xc7\x7a\x79\x55\x4b\xec\x56\x1d\x1b\xb7\xfd\x86\x47\xf6\xe5\x60\xbc\x34\x86\xcc\xf2\x07\xc6\xa3\x46\xbf\x0b\xd3\xc6\xd2\x39\x0b\x06\xaa\xf9\x51\x67\xe4\x5c\xf0\xb1\xa3\x83\x97\x0e\x66\x6c\xb3\xee\xa0\xd0\x58\xf9\x25\x35\x5f\xde\x62\x8a\x92\xa1\xd4\x33\x9a\xda\x09\xad\x2c\x3b\xcc\x7d\xce\x6a\xa0\xb0\x98\x7e\xd6\x94\x4b\xae\x73\xdb\x99\x15\xce\x45\x9f\x2a\xe4\x86\x2d\xab\xee\x57\x36\xaa\x81\x9e\x94\x16\xad\xf7\x98\xb4\xda\x04\x14\x3a\xf0\xc2\xd2\x53\x30\xa7\xea\x54\x4b\x7e\xf4\x7d\x5f\x83\xc1\x24\xc0\x22\x18\x41\x84\x2d\x32\xed\x6a\x81\xa6\x9a\x65\xe5\xd6\x40\x1b\xa6\xf7\x0b\x23\x35\x38\xf2\x54\xee\x57\x7b\x6a\x88\x09\x0d\xe9\x02\x93\x43\x81\xe2\x35\x16\xf2\xa8\x46\x64\x3d\x28\xac\x8f\x54\x7d\x94\xf5\x5e\x5b\xf3\x88\x14\x1a\xe7\x16\xc1\x62\x19\xed\x1a\x16\xea\xb5\xaa\x4d\xce\x9b\x0e\x6d\xe5\x0b\x36\x70\x64\x6a\x47\xfd\x1b\x24\xb2\x7c\x33\x3d\xb1\xb3\xa1\xe3\x96\xfd\xec\x37\xde\x36\x41\x68\xb1\x08\x12\xdc\x04\x59\xb3\x75\xdc\xeb\x0a\x1b\xe6\x00\x0d\x0c\xf6\x17\x10\x2c\x09\xda\xcf\x23\x9a\xb4\xdd\x68\x50\x3d\x4c\xbb\x5f\x36\xe9\xdc\x9e\x8b\xb8\xd3\x0f\x17\xd4\x3f\xfc\xbe\x6d\x23\x74\x62\xca\x8d\x6c\xe1\x44\xa0\x34\x86\xf1\xf3\xa6\x8d\x0d\xb0\xc1\x9c\xfb\xd6\xdf\xd8\x58\xde\xb0\x39\xd3\xed\xb6\xa2\xab\xc6\x76\x03\x6d\xb7\x3a\x1d\x09\x9d\xcb\x8f\x37\xd8\xa0\x63\xcc\xce\x00\xe7\x22\x44\x4c\x06\xf3\x28\xd5\x5f\x4c\xbe\xd7\xcb\x2e\xd8\x7f\x1e\x32\x46\x77\x84\x58\x17\x9c\x94\xc4\x0f\xa1\xf6\xb3\xff\xfe\x9b\xbd\x0f\x37\x18\x84\x7b\x5f\xd5\x32\x6b\xc7\x75\x8e\xf7\xf0\x6d\xec\xbf\x0f\x02\xd1\x3e\xde\xd2\xb8\xed\x18\xe9\xb1\x60\x3f\x47\xbe\x10\x60\xfa\x20\xca\x32\x3c\xc2\x4b\x3e\x06\xd2\x4f\xd2\xb3\x6f\x1a\x35\x39\x8e\x78\xd4\x18\x2f\x7d\xe8\x0e\x47\xf2\xb0\x4d\x8d\xe2\xb4\x40\x6e\x7b\xe8\xfd\xae\x92\x4e\x78\xe9\x1f\x72\xb0\xf4\x71\x06\x5c\x38\x59\xea\x77\x70\xfa\x2f\x1c\x69\x6e\xf6\xe3\xe4\x8c\x3f\x7e\xcc\x72\xa9\xf3\x28\xfa\x5c\x2e\x2f\xe3\xd7\x2c\x5d\x70\xb9\xbc\x47\x4f\xd7\x9f\xef\xbf\xc6\xba\x5a\x68\xc5\xb4\xd3\x35\xa8\xfe\xef\xe7\x0f\x33\x69\x88\x0b\xfc\x84\xe2\xa2\x9b\x8d\xff\x05\xcb\x33\x4c\xcf\xd3\x86\xab\x7b\x9a\x7e\x17\x86\x27\x7b\x16\xfd\xf8\x0c\x74\x3c\xf8\xcb\x6a\xf0\x17\xbc\x68\xf3\xdf\x04\x3b\xba\xbf\x9f\x34\x7e\x5b\x0d\x70\xd5\xeb\x34\xf9\x25\xe8\xf4\x5b\xaf\xb7\xfd\xc9\x3e\xad\x9f\xfb\xd2\x4e\x5b\x03\xfa\xc7\x94\xe7\xc5\xcc\xe5\x07\xe0\xf4\x16\x7f\x62\xc1\x4f\xaa\xbf\xf3\x04\xd3\xfa\xa6\x7f\x68\xad\x7f\xc2\xb2\x4c\xb7\xd9\xcf\xf4\xa5\x96\x7f\x86\x55\xbf\x4e\x2f\xf0\x77\x9e\x65\xc5\x3f\x88\xc2\x4f\x98\xde\x70\x9d\x2f\xf4\x67\x7a\xfb\xe2\xe1\x29\xe2\x0f\x98\x97\xba\xe3\x17\xc2\x95\x5f\x60\xc1\x69\x5a\x1d\x6d\xfb\xbf\xf2\x3f\x79\x02\x4b\xaf\x70\x83\x4b\xfd\x39\x7d\x25\x4d\xb4\x4e\xb4\x5e\xe3\xf5\xef\x02\xf3\xd3\xf2\xdc\x5f\xf1\xe7\x09\xcc\x75\xb5\xf9\x5d\x71\xc5\x19\x5c\xf8\x05\x9a\x66\x03\x1c\x7e\xea\x5f\xf3\x2b\x5c\xd7\x2f\xba\x74\x8b\x9f\xde\x9e\xe1\x79\xb2\x7f\xdc\xb9\x73\x9f\x7f\x3d\xab\xd7\x0d\xfd\xf4\x4b\xbf\xb9\xdf\x10\x26\x40\xf8\x01\x73\xa0\xf0\xe9\x73\x75\x99\x26\xfd\x66\x7e\x98\x29\xe3\xba\x7c\x32\xa5\x2e\xd3\xac\xcd\x9b\xfb\x33\x01\xbe\xad\xf3\xf4\xcc\x9f\x86\xa0\xcb\x44\x4b\xa0\x1f\xf0\xf6\xfa\x49\x36\x97\x09\x19\x68\x5f\x9e\xc1\x3e\xeb\xe9\xb7\xf9\x57\x81\xca\x34\x55\xf3\x17\xdf\xa6\x39\x2e\x3f\xba\x5b\x1d\x06\x01\x78\xb3\x29\x68\xe3\xfe\x9a\xf5\x17\x04\xfc\x53\x3e\x3e\x2c\x2a\x01\x20\x19\xf1\x36\xab\x46\xff\x0a\xbf\xe8\xb1\x03\x03\xcb\xa2\x7f\x4e\x66\x75\x33\xf8\x3c\x1b\x7a\x7b\x8d\xbf\xa7\x4f\x9d\x98\x45\xd6\xe9\x2f\xd0\xbc\x86\x5f\xe9\xa3\xf8\xbf\x27\xd0\x18\xc8\x46\x9c\x26\x63\x5e\xe7\x5f\xe6\xb3\xfc\x4d\x22\x5e\xc0\x1f\xa0\xfd\x21\xd4\x73\x93\xea\x67\xd0\xfb\xf2\xcf\x1b\xcd\x33\x4c\x6f\x4f\x30\xe5\xff\x8e\xe3\x60\x70\x69\x66\x7c\x82\xff\x3a\xca\x13\x1e\x3e\x8b\x0d\x6e\x7a\x0d\x76\x39\xf2\x63\x83\xbf\xcc\xaf\xfa\x9f\x85\x33\x1c\x65\xbb\xb8\xd7\x34\x6f\xbf\x46\xb4\xbb\x7f\x3c\xa6\x76\xcc\x4a\x73\x8f\xc0\xf7\x95\x52\x4d\x4b\x3b\xd0\x77\x8b\xfa\xf2\xd7\xd1\xb5\xe5\x44\x76\x36\x19\x25\x72\xdd\xb9\x46\x4e\xe1\x3d\x5b\xfa\x91\x31\xb2\xa5\x07\x73\x0c\x53\x0e\x9e\x7e\xab\x96\xb2\xe1\x4b\x3f\x98\x6e\xb0\x6e\xd0\xbc\xd0\x0d\x06\x40\x6e\xae\xe8\x06\x9b\x80\xc4\xe6\x1b\x6f\xc6\x2b\x41\x7d\xfd\x81\x51\x90\x53\x81\x17\xea\x8e\xed\xd6\x7e\xff\xa2\x7c\x1a\xb7\xc4\xc5\x7e\x37\x66\xdf\x2e\xc9\xb5\x67\x04\x10\x37\x40\xf1\x7f\x01\x00\x00\xff\xff\xf1\x3c\x8a\xce\xe9\x3e\x00\x00")

func blacklistTxtBytes() ([]byte, error) {
	return bindataRead(
		_blacklistTxt,
		"blacklist.txt",
	)
}

func blacklistTxt() (*asset, error) {
	bytes, err := blacklistTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "blacklist.txt", size: 16105, mode: os.FileMode(420), modTime: time.Unix(1541751116, 0)}
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
	"blacklist.txt": blacklistTxt,
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
	"blacklist.txt": &bintree{blacklistTxt, map[string]*bintree{}},
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

