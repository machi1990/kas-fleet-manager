// Code generated by go-bindata.
// sources:
// ../../openapi/.managed-services-api.yaml.swp
// ../../openapi/managed-services-api.yaml
// DO NOT EDIT!

package openapi

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

var _OcmExampleServiceYamlSwp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9b\x6b\xac\x24\xc7\x55\xc7\x7f\x3b\xef\xf7\x20\x1c\x45\xb6\x10\xa2\xc2\xc6\xba\x18\xf6\xce\xdc\xb5\x13\x63\x06\xfc\x76\x9c\xac\xf1\xee\x9a\xbd\xde\x35\x28\x84\x7b\xeb\x76\xd7\xdc\x69\x6f\x4f\xf7\xb8\xbb\x66\xf7\xce\x1a\x83\x90\x12\x19\x27\x11\xa0\xd8\x89\x84\x82\x43\x88\x8d\x12\xc0\x44\x01\x23\x25\x01\x4b\x08\x02\x48\x49\x0c\x28\x60\x91\x18\x04\xc2\x81\x3c\x1c\x2b\x58\x3c\x04\xc8\x24\x12\xea\xc7\xcc\xf4\xcc\xce\xf5\x7d\x90\x5d\x2f\x71\xff\xbf\xf4\x54\xd5\xa9\xaa\x73\xaa\x4e\x55\x9d\x73\xaa\x66\x63\xe5\xd4\x91\xa3\xe2\xba\xd6\x61\x80\xef\x80\x07\xef\x79\xfe\x6d\xd9\x0f\x1d\xe2\xd9\x02\x68\xeb\xac\x65\xdb\x96\x64\x07\xec\x96\xee\x67\xc6\x84\xed\x4d\xb7\xed\x7b\x46\x7b\xd3\xd2\xb6\xdc\x68\x19\x4a\xb5\x3c\x65\xf6\xa4\x6e\x19\x6e\xbf\xed\x2b\xef\x8c\x65\xa8\xb6\x6f\x6e\x2c\xbb\x46\x7f\x59\x6d\xc9\xfe\xc0\x56\xcb\xe3\x7c\x77\xa0\x1c\x39\xb0\xda\x0b\xca\x5a\x23\xd9\xb7\x77\x62\x23\x45\x8a\x57\x1f\x86\xba\xbb\x7c\x5d\x9d\x6b\xae\x3e\xbc\x12\x24\x0f\x7e\xef\xeb\xc4\x65\xdf\x79\xf2\x95\xe6\x2a\x45\x8a\x14\x29\x52\xa4\x48\x91\x22\x45\x8a\x14\x17\x11\x7a\x90\xe3\x67\x81\x4c\x9c\x1e\xc5\xdf\x03\x73\xdf\x5c\xfc\x7d\x47\xfc\x3d\x37\x57\x9e\x8f\xbf\xe5\xf8\xfb\x8d\xb9\xf2\x6c\xfc\x5d\x8d\xbf\x99\x03\xb3\xe5\x29\x52\xa4\x48\x91\x22\x45\x8a\x14\x29\x52\xa4\x48\x91\x22\x45\x8a\x0b\x07\x69\xc2\x6b\x80\xd7\x64\xa2\xfb\xff\xb1\xff\xff\xaf\x4d\x78\xa1\x09\x9f\x69\xc2\x63\x4d\x78\x4f\x13\x1e\x6a\x42\xaf\x09\x37\x35\xe1\x8a\x26\xfc\x4f\x03\xfe\xad\x01\xcf\x34\xe0\xcf\x1b\xf0\x78\x03\xde\xdf\x80\x41\x03\xd6\x1b\xf0\xa3\x0d\xe8\x34\xe0\x07\x1a\xf0\xef\x75\xf8\xc7\x3a\xfc\x71\x1d\x9e\xa8\x83\x57\x07\xa3\x0e\xd7\xd6\xe1\x60\x1d\xf2\x75\x78\xb1\x06\x4f\xd5\xe0\xb7\x6b\xa0\x6b\x20\x6b\x70\x6b\x0d\x96\x6b\xf0\x52\x15\x5e\xa8\xc2\xef\x57\xe1\xc3\x55\x78\xa8\x0a\x7e\x15\x6e\xaa\x42\xab\x0a\x07\xaa\xf0\x4c\x05\x7e\xbd\x02\xc3\x0a\x9c\xae\xc0\xed\x15\xb8\xb6\x02\xaf\xab\x40\xa3\x02\x2f\x94\xe1\x99\x32\x7c\xaa\x0c\x1f\x2e\xc3\xc3\x65\x58\x2b\xc3\xb1\x32\xdc\x58\x06\x51\x86\x7a\x19\x5e\x2c\xc1\xb3\x25\x78\xac\x04\x8f\x94\x60\xa3\x04\xc7\x4b\x70\xb8\x04\xdf\x55\x82\xaf\x16\xe1\xd9\x22\x3c\x51\x84\x5f\x2e\xc2\xfd\x45\x50\x45\xb8\xba\x08\x07\x8b\xf0\xb5\x02\x7c\xbe\x00\x9f\x2c\xc0\x63\x05\x18\x14\x40\x16\xe0\xea\x02\x7c\x4f\x01\x0e\x14\xe0\x85\x3c\x7c\x22\x0f\xbf\x91\x87\xb5\x3c\xdc\x99\x87\x56\x1e\x2e\xcf\xc3\x97\x72\xf0\xf1\x1c\x7c\x34\x07\x0f\xe5\x60\x94\x03\x33\x07\xab\x39\xb8\x2d\x07\x97\xe7\x20\x9f\x83\xaf\x64\xe1\x73\x59\x78\x7f\x16\xde\x9d\x85\xb5\x2c\xdc\x99\x85\x56\x16\x2e\xcf\xc2\x97\x32\xf0\x37\x19\x78\x34\x03\xef\xcc\x80\x9d\x81\x53\x19\x38\x98\x89\xe6\xfb\xb2\xcc\x62\x9d\x10\x31\x96\xde\xb0\xf2\x86\xa5\xce\x24\x39\xc6\xeb\x3d\xd5\xed\x88\xa5\x83\x6d\xc3\xed\x0f\x5c\x47\x39\xda\x6f\xfb\x46\x4f\xf5\xa5\xdf\x7e\x93\xe7\xb9\xde\xd2\x5c\x95\xa8\x70\xb6\x21\x39\x18\xd8\x96\x21\xb5\xe5\x3a\xed\x7b\x7d\xd7\x49\x96\x1a\xae\xa3\x95\xa3\x93\x59\xa6\xf2\x0d\xcf\x1a\x04\xe4\x1d\x71\xd2\x91\x43\xdd\x73\x3d\xeb\x9c\x32\x85\x76\xc5\x40\x79\x5d\xd7\xeb\x0b\x77\xa0\xbc\xb0\xc5\xa4\x04\xd7\x5c\x8a\x12\xdc\x3c\xd4\x3d\xa1\xdd\xd3\xca\x11\x96\x2f\x2c\xe7\x8c\xb4\x2d\x33\xc9\xf5\xe1\x3d\x72\x7d\x9b\xe5\xb8\xbe\x1c\x5e\x70\xc6\xc7\xfd\x88\xae\x3b\x74\x4c\xb1\x31\x12\x49\xbe\xaf\x5e\x59\x99\xf0\xed\x29\x7f\xe0\x3a\xbe\xf2\xa7\x6d\x2d\x8b\x5b\x94\xf4\x94\xd7\x11\x6f\x7d\x5b\x9c\xe9\x2b\x63\xe8\x59\x7a\x34\x26\xf2\x87\xfd\xbe\xf4\x46\x1d\xf1\x66\xa5\x85\x74\x84\x39\xee\x6f\xda\xd3\xa6\x0a\xd9\x6b\x6f\xf3\x04\xa8\x7d\xe6\x70\x7b\x5c\xcb\x6f\xdf\x6f\x99\x0f\x5c\x8a\xf3\xef\x88\xa1\xa3\xb6\x06\xca\xd0\xca\x14\x2a\xe8\x52\xb8\x86\x31\xf4\x3c\x65\x0a\xc3\x53\x52\x5b\xce\xa6\xd0\x3d\x35\x19\x80\xe9\x20\xbf\x31\x31\xc8\x97\x90\x48\x13\xcd\x90\xb6\xa7\xa4\x39\x12\x6a\xcb\xf2\xb5\x9f\x54\xea\x1f\xba\x14\xf9\x4e\x37\x93\x57\x86\xeb\x53\x01\x97\x61\xd5\x48\xff\xfd\xc9\x02\x48\xb2\xbe\x57\x55\xbf\x58\xfb\xe0\xad\xc1\x1a\x55\x33\x5b\xdf\xe1\x97\xdb\xfa\xf6\xcb\xf9\xf9\x7c\x6f\xcf\xf5\x79\x3c\x7b\xea\xbe\xa1\xe5\x29\xb3\x23\xb4\x37\x54\x2c\x14\x64\xb2\x6c\x4d\xa9\x25\xd3\x7a\xca\xd7\xb7\xb8\xe6\x68\x7f\x9b\x77\x34\x3a\x42\x0a\x47\x9d\x9d\xdd\xc1\x06\xae\x3f\x61\x70\x79\xe1\x78\x0c\xa4\x27\xfb\x4a\x2b\xcf\x6f\xbb\x9e\xa9\xbc\x5b\x46\x4b\xbb\xa5\xf7\x95\xf4\x8c\xde\xee\xc9\xad\x73\x6a\xd7\xc4\x03\xb9\x39\x21\x9e\x66\x5f\x82\xeb\xea\xe4\x76\xe7\xca\x25\x7f\x82\xa4\x3b\xf1\xb6\x9b\xc2\x9d\x96\xaf\x2f\x38\xf3\xe2\x8e\xd5\xe3\xc7\x84\xf4\x3c\x39\x12\x6e\x77\x6a\x7a\xb9\x1b\xf7\x2a\x23\x79\x92\x7f\x0b\xcd\xbc\x13\x4a\x0f\x3d\xc7\x17\x52\xd8\x96\xaf\x93\xdd\xfa\x7b\x36\xf6\x3a\x0c\xa4\xee\x85\x8c\xcc\x08\x16\xaa\xbf\xa3\x85\xe9\xf6\xa5\xe5\xb0\x2c\x86\x9e\xdd\x11\xed\x79\xb2\x55\x2d\x37\x03\xab\x2b\x68\x5a\x79\x63\xb2\x9e\xd6\x03\xbf\xd3\x0e\x18\x68\xf9\x5a\x6e\xaa\x96\x3b\x50\x8e\xdf\xb3\xba\xe1\x63\xf5\xf9\x46\x8e\x4a\xcb\x11\xdf\x37\xf0\x5c\x73\x68\x04\x39\x57\xbd\x4c\x73\xb3\x0d\x45\x74\x21\xf7\xc1\x37\x6c\x6d\xa5\xb5\xd2\x3a\x3c\xdf\xc5\xf1\x5b\x8f\x8a\x37\x45\xc3\x20\x56\xa3\x61\x10\x37\xdf\x75\x04\x21\xb4\xa5\x6d\xb5\x3d\x81\xe5\x74\xdd\x0e\xf1\xeb\xf9\x8e\xb8\xa6\xb5\xd2\x5a\x09\xfd\xff\x6b\xf3\xf0\x81\x42\xe4\xff\x8f\xef\xe7\x03\xbf\xff\x63\x4d\x78\xb4\x09\xef\x6a\xc2\xb9\x26\x18\x4d\x38\xd9\x84\xa3\x4d\xb8\xa1\x09\xed\x26\x54\x9b\xf0\x9f\x0d\xf8\x6a\x03\xfe\xba\x01\x9f\x6e\xc0\x27\x1a\xf0\x91\x06\x1c\x6d\xc0\x8d\x0d\xb8\xaa\x01\xaf\x6d\x40\xb6\x01\x5f\xaf\xc3\x3f\xd5\xe1\xaf\xe2\x58\xc0\xbb\xea\xf0\x40\x1d\x36\xeb\x70\xa2\x0e\x37\xd6\xe1\xaa\x3a\x7c\x77\x1d\x2a\x75\x78\xa9\x06\xff\x5c\x83\xbf\xa8\xc1\x83\x35\xb8\xa1\x06\x2f\x56\xe1\xf1\x2a\xfc\x7c\x15\x1e\xac\xc2\x56\x15\x7e\xb2\x0a\x77\x55\xe1\x78\x15\x2a\x55\xf8\xbd\x0a\x3c\x59\x81\xc7\x2b\x70\xaa\x02\x77\x54\xe0\x48\x05\x0e\x54\xe0\xa9\x32\xfc\x66\x19\x7e\xa5\x0c\x0f\x96\xe1\x5c\x19\xba\x65\x38\x55\x86\x9b\xca\xf0\xfd\x65\x78\xae\x04\xef\x2d\xc1\x91\x12\x7c\xb3\x08\xdf\x28\xc2\xd7\x8b\xf0\x85\x22\x3c\x5d\x84\xcf\x16\x61\xab\x08\x67\x8b\xd0\x2b\xc2\x6d\x45\xb8\xae\x08\x3f\x58\x84\x17\x0b\xf0\x47\x05\x78\xb2\x10\x8d\xeb\xa3\x85\x8b\x15\xcd\x49\x91\x22\x45\x8a\x6f\x17\x8c\x4d\x0a\x3d\x1a\xa8\x8e\xf0\xb5\x67\x39\x9b\x2c\x36\x7f\x1c\x57\xa8\xad\xc0\xfa\xb1\xb4\x08\x1d\x87\xe0\x2c\x3f\x6b\xd9\xb6\xd8\x50\x91\x5d\xa4\xcc\xd6\x84\xfa\x48\x37\x8c\xae\x4c\x2c\x79\x61\xf9\xce\x92\x16\x03\xcf\x3d\x63\x99\xca\x3c\x24\x5c\x4f\x58\x11\xcd\x19\x69\x0f\x55\x60\xcf\xa9\xfe\x40\x8f\x0e\x05\x79\xce\x94\xb3\xf5\xf5\xf5\xc9\xef\xa1\xaf\x3c\x47\xf6\x95\x90\xbe\x71\x48\x74\x2d\xcf\xd7\xc7\xe2\x64\x92\xde\xbf\xcf\x9e\xd6\x3f\xee\x09\xcb\x89\x18\x0e\x6c\x5e\x4f\x69\xcf\x52\x67\x94\x90\xb6\x2d\xa4\x61\xb8\x43\x47\xfb\x51\xb1\x0a\xc3\x6d\x93\x3e\xd6\xa4\x63\xae\x45\x9d\x88\x20\xa3\xb3\x33\x4f\xdb\x32\x71\xbb\xeb\x89\xd8\xb2\x3a\xb4\x3f\x76\xa6\xf3\xb0\x31\xd4\x62\xe8\x8f\xc3\x57\x41\x99\x1f\x98\x75\x41\x22\xb0\x48\x85\xd4\xda\xb3\x36\x86\x5a\xf9\xa2\x2d\x0c\xd7\x1e\xf6\x9d\x71\x79\xdc\xc1\x74\x96\x7c\xab\x6f\xd9\x32\xe4\x24\x28\xf7\x47\x8e\x96\x5b\x63\xea\xb5\x88\xcb\x8d\xd1\x9a\x30\x6c\x39\xf4\x55\x50\x20\x1d\xb1\xfa\x63\x77\x0a\x5f\x4b\xad\xfa\xca\xd1\x87\x26\x6d\xad\x0e\x94\x61\x75\x2d\xe5\x87\x95\xc7\x75\x85\xe1\x59\x5a\x79\x96\x6c\x89\xbb\xe7\x7a\xb0\xfc\x19\xf5\x60\x81\xc9\xfc\xd3\xcb\x09\x8f\x39\xf2\xb4\xbb\xd2\xf6\xc7\xae\xb6\xe5\x74\xc4\x7d\x43\xe5\x8d\xe2\x74\x38\x50\x22\xf6\x6a\xc3\xbc\xf8\x77\x67\x2f\x8a\xee\x85\x26\x73\x52\x99\x83\xc9\x49\x8c\x5f\x20\xa1\xd4\x61\x4e\x30\x3d\xa2\x27\xfd\xc0\x97\xea\x5b\x7e\x60\x55\x06\x83\xe9\x2b\x35\x5e\x1a\x17\x76\x45\xf8\xc3\x8d\xc9\x70\xad\xd9\x72\x43\xd9\x7e\xeb\xb4\x1a\x89\xeb\xc5\x52\xd7\x75\x97\x84\x74\xcc\x85\x34\x51\x0f\xd7\x8b\xa5\x0d\x99\x70\xda\xf6\xa7\xb6\xa1\xe6\x24\xfa\xf0\x45\xd8\x49\xa4\xbf\xeb\x5d\xd7\xbd\x7e\x43\x7a\xeb\x53\x45\x09\xf5\x20\x8c\x24\x4c\x94\x43\x18\xd2\x11\xd2\xf6\xdd\xc4\x5e\x22\x5c\x47\x78\xca\x96\x81\xa3\xed\x29\xdf\x1d\x7a\x86\x6a\xed\xb0\x06\x6d\xeb\xb4\x12\x4b\xfd\xd1\x95\xdb\xcb\xe4\x6b\xe9\xe9\x68\xeb\xd2\x3d\xb1\xde\x1f\xad\x77\xf6\x21\xef\x44\x13\xc2\x56\xe4\x84\x83\x04\x4b\x7b\x5b\xa1\x33\x54\xe3\x75\x2a\x76\xb9\x50\xcf\xf6\x94\xa7\xf6\xb3\x4a\xe7\xa6\xe1\xe2\xad\xd1\xa8\xe3\x30\x2b\xfa\xd9\x79\xd9\x46\x84\xe8\x5b\x8e\xd5\x1f\xf6\x3b\x62\x25\x11\x66\xeb\xca\xa1\xad\x3b\xe2\xf0\xca\xca\xdc\xfa\xb6\x1c\xad\x36\x95\xb7\x70\x81\xcf\xb9\x93\x5b\x41\xb3\xc2\x19\xf6\x37\x94\x17\xc8\xec\x29\xc3\xf5\x4c\x3f\x9e\xf5\xa1\xe7\xbc\xbc\x1c\xd6\xb9\x88\xc7\xe0\xc7\xae\x65\x38\xbc\x40\x86\xfd\x4a\x70\x97\xdc\x54\xe7\xb1\x1f\x39\xfb\x67\x7b\xca\x99\xc9\x50\x5b\x86\x52\xa6\x2f\xfc\x58\x09\x4c\x31\x08\x6a\x4f\x84\xd8\x4e\xca\x80\x2a\x8a\x2e\xca\x4d\xb5\x87\xbd\x74\x61\x7c\x74\x86\xfb\x40\xdf\x2c\x73\xca\x78\x82\x8d\x81\xd4\xbd\x19\x2e\xe2\x48\x8f\x65\x06\xcd\x2f\x0e\x11\x2e\xe0\x48\x44\xd2\x4e\xe3\x27\x03\xcf\x1d\x28\x4f\x27\x72\xa2\x5a\x51\x14\x26\xcc\x1a\x47\x85\xee\x92\xda\xe8\x9d\x88\x82\xb5\xf3\x71\xa5\xc5\x01\xcd\xf3\x83\xcd\x81\xff\xaf\x80\xf7\xc6\xf7\xff\xe3\xf7\xfd\x7f\xd2\x84\xdf\x69\xc2\xc3\x4d\x78\xa0\x09\x77\x34\xe1\x87\x9b\x50\x69\xc2\x7f\x35\xe0\xd9\x06\x7c\xaa\x01\xbf\xd4\x80\x77\x34\xc0\x68\xc0\x6a\x03\x6e\x88\xef\xfb\x69\xc0\xf3\x75\xf8\x42\x1d\x9e\xaa\xc3\x47\xea\xd0\xab\xc3\x3d\x75\x78\x4b\x1d\x2e\xab\x03\x75\xf8\x62\x0d\x9e\xae\xc1\xfb\x62\xff\xfe\x64\x0d\x6e\xa9\xc1\xeb\x6b\xd0\xa8\xc1\xdf\x55\xe1\xe9\x2a\xfc\x6a\x15\xde\x5d\x05\xa7\x0a\x3f\x5e\x85\x2b\xab\xf0\xda\x2a\xfc\x7d\x05\x3e\x53\x81\x27\x2a\xf0\xbe\x0a\x18\x15\xb8\xbb\x02\x97\x57\x20\x5f\x81\xaf\x94\xe1\x73\x65\xf8\x40\x19\x7e\xb1\x0c\xeb\x65\x38\x5a\x86\x76\x19\xae\x28\xc3\x97\x4b\xf0\xf9\x12\x3c\x5a\x82\x77\x96\xc0\x2e\xc1\xa9\x12\x1c\x2c\xc1\x65\x25\xf8\x5a\x11\xfe\x21\xf6\xf7\xff\xb0\x08\x1f\x8d\xef\xfc\x7f\xae\x08\x67\x8a\x20\x8b\x70\xac\x08\x3f\x52\x84\x95\x22\x5c\x51\x84\x6a\x11\xfe\xa3\x00\x5f\x2e\xc0\x5f\x16\xe0\x4f\x0b\xf0\xbb\x05\xf8\x60\x01\xde\x53\x80\xfb\x0b\xd0\x2f\xc0\xc9\x02\xbc\xb9\x00\xed\x02\x88\x02\x94\x0a\xf0\xdf\x79\x78\x3e\x0f\x7f\x9b\x87\xcf\xe6\xe1\x0f\xf2\xf0\x5b\x79\xf8\x60\x1e\x1e\xce\xc3\x4f\xe5\xe1\x48\x1e\xde\x98\x87\x2b\xf3\x50\xcb\xc3\x4b\x39\x78\x2e\x07\x9f\xce\xc1\xc7\x72\xf0\x6b\x39\x78\x24\x07\xf7\xe6\xe0\x64\x0e\x6e\xce\xc1\x4a\xe2\xdd\xc0\x33\x59\xf8\xb3\x2c\x7c\x32\x0b\x6f\xcf\x82\x9b\x85\xb7\x66\xe1\x2d\x59\x68\x67\xe1\x8a\x6c\xf4\x87\x8d\x2f\x66\xe0\xe9\x0c\x3c\x99\x89\xf4\xe3\x91\x0c\xbc\x3d\x7e\x47\xf0\x13\x19\xb8\x2d\x33\x6b\xa5\x47\x90\xb6\x7d\xbc\x3b\xab\x8e\x96\x56\xfd\xb9\x58\x79\xa4\xd4\x61\xb0\x31\xcc\x5f\x68\xf6\x07\xe8\xba\x5e\x5f\xea\x8e\x30\xa5\x56\xcb\xda\x4a\x9c\x61\xdb\xae\x28\x21\x86\x83\x80\xdc\x5c\x93\x73\xeb\x62\x5f\x8d\x19\xd1\x5d\xcf\x79\x8d\x6d\x43\x3e\xb7\x9e\x17\xad\xe8\xe5\xf3\xd7\xf4\x4e\x2b\xf6\x78\x48\x7a\x42\x75\x95\xa7\x1c\x63\x72\x15\x91\x18\xec\xf1\x9a\xde\xcb\xa5\xd3\x82\xa8\xf8\x2e\xe6\x6a\x21\xe1\xb7\x44\xca\x64\x9c\x3b\x21\x5a\xc8\x66\x50\xb6\xab\x09\x98\xdc\x17\xac\x45\x7b\xf1\x8e\x15\x3c\x25\xe7\x62\xe5\xdb\xeb\x82\x6b\xaa\x8b\x3e\xb3\xa1\xf8\xc9\xf8\x7a\x38\xf4\x89\xb4\x76\xb5\xb4\x13\xe9\xc4\x89\x19\x24\x27\xe7\x63\x94\x3c\x6d\x39\xe6\xfc\x09\x78\xde\x21\x35\x7b\xc0\x8b\xa8\x8b\x9d\xc9\x92\x76\xc6\xf6\x54\xb3\xe7\xf4\x36\x83\x1d\xb0\xb9\x87\x43\x71\x5e\x3d\x16\xb6\xd9\x0b\xe6\x61\x2f\x1d\x6f\x4b\x64\xed\x85\xb7\xb9\x29\x0e\xe8\xe2\xd9\x1f\x57\xd9\x08\x6f\x4c\x6e\x8f\xb7\xa7\x3b\xee\xb9\x3b\x69\xb0\xa8\x4e\x5c\x3e\xd3\x7e\x4f\xeb\x41\x98\x11\xdf\xb6\x30\xbd\x66\x59\x0d\x6b\xf9\x1d\xa6\xea\x16\x75\xb4\xe3\x35\xa7\x65\x2e\xc5\x96\xd4\xff\xa7\x2b\xce\x70\xbb\x0f\x1c\x9a\xf4\xb5\xcc\x05\xe6\xfb\x98\x3b\xbd\x17\x0c\x7d\xcb\xa9\xad\x6e\x99\x0b\x84\x48\xdf\x0f\x5e\x10\x09\xd2\x27\x3f\xaf\xe0\xd3\xc7\xd8\xb8\x14\xfe\xd0\x30\x94\xef\x77\x87\xb6\x3d\xb5\x8b\x76\xba\x1c\xdf\xbd\x3c\x49\xef\xee\x22\x3f\x0a\x3a\x19\x4b\x68\x5e\x88\xc7\x41\x51\xe3\xc9\xc7\x9d\xf1\x89\xa3\xa7\x61\x96\xdd\x8d\x54\xfa\x9e\xe6\xff\xbe\x57\x07\xfe\xff\x37\x1b\x70\x7d\x33\xf2\xff\xc7\xff\xdf\xff\x97\x26\x3c\xd7\x84\x8f\x37\xe1\x43\x4d\xf8\x85\x38\x0e\xa0\x9a\x70\xa2\x19\xd3\xa7\x48\x91\x22\x45\x8a\x14\x29\x52\xa4\x48\x91\x22\x45\x8a\x14\xaf\x06\xcc\xfa\xef\xfb\xbb\xf8\x5c\x48\x78\x21\x2f\xd3\x92\xff\x29\xe9\xf0\xbf\x01\x00\x00\xff\xff\x5c\xa7\x00\x06\x00\x60\x00\x00")

func OcmExampleServiceYamlSwpBytes() ([]byte, error) {
	return bindataRead(
		_OcmExampleServiceYamlSwp,
		".managed-services-api.yaml.swp",
	)
}

func OcmExampleServiceYamlSwp() (*asset, error) {
	bytes, err := OcmExampleServiceYamlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".managed-services-api.yaml.swp", size: 24576, mode: os.FileMode(420), modTime: time.Unix(1575573384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ocmExampleServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x73\xd3\x48\x12\x7f\xf7\xa7\xe8\xaa\xbb\x2b\xdf\x55\x25\xb6\x73\x70\x0f\xe7\x2a\x1e\x80\x83\x2b\x28\x20\x2c\x81\xdd\x87\xad\xad\xa4\x3d\x6a\x59\x43\xa4\x19\x31\xd3\x4a\x62\x76\xf7\xbb\x6f\xcd\x8c\x2c\xc9\xb2\x2c\x64\x16\x70\xd8\x4a\x9e\xa2\x71\x77\x4f\xff\xfa\xbf\x5a\x3a\x27\x85\xb9\x9c\xc3\xbd\xc9\x6c\x32\x1b\x49\x15\xeb\xf9\x08\x80\x25\xa7\x34\x87\xd3\xc7\x2f\xe1\xc9\x0d\x66\x79\x4a\x70\x46\xe6\x4a\x0a\x82\x87\xaf\x9f\x8d\x00\x22\xb2\xc2\xc8\x9c\xa5\x56\x7d\x64\x57\x64\xac\x27\x99\x4d\x66\x93\x93\x91\x25\xe3\x4e\xe6\xa3\x63\x28\x4c\x3a\x87\x84\x39\xb7\xf3\xe9\x14\x73\x39\x71\x9a\xd8\x44\xc6\x3c\x11\x3a\x6b\xdf\xf0\x12\xa5\x82\x7f\xe6\x46\x47\x85\x70\x27\xff\x82\x20\xaa\x4b\x90\x65\x5c\x52\xbf\xb8\x33\xc6\xa5\x54\xcb\x96\x90\x69\x9b\x4c\x14\xc6\x90\x62\x88\x74\x86\x52\x8d\x72\xe4\xc4\x3a\xeb\xb8\x6b\xa6\x5a\x64\xc7\x14\x40\x1f\xdb\x00\x7a\x7a\x75\x32\x8d\xa4\xd2\x16\x0b\xe3\x09\x01\x96\xc4\xe1\x1f\x00\x5b\x64\x19\x9a\xd5\x1c\xde\x10\x17\x46\x59\x40\x48\xa5\x65\xd0\x31\x54\x4c\x6b\x52\x12\x85\x91\xbc\x5a\xb3\x02\x1c\xc3\x23\x42\x43\x66\x0e\x3f\xff\x52\x1e\x1a\xb2\xb9\x56\x96\x6c\x4d\x35\xfe\xf7\x6c\x36\xae\x1f\x5b\x70\x1e\xc2\xf3\xb3\xd3\x57\x80\xc6\xe0\xaa\x79\x2b\xe8\xc5\x7b\x12\x6c\x1b\x7c\x42\x2b\x26\xc5\x4d\x51\x00\x98\xe7\xa9\x14\xe8\x84\x4d\xdf\x5b\xad\x36\x7f\x05\xb0\x22\xa1\x0c\xdb\xa7\x00\x7f\x37\x14\xcf\x61\xfc\xb7\xa9\xd0\x59\xae\x15\x29\xb6\xd3\x40\x6b\xa7\xff\x2b\x75\x78\x21\x2d\x8f\x6b\x1c\xf7\x67\x27\x3d\x38\x0a\x4e\x80\xf5\x25\x29\x90\x16\xa4\xba\xc2\x54\x46\x87\x50\xfe\x89\x31\xda\x6c\x68\x7d\x6f\xb7\xd6\xef\x14\x16\x9c\x68\x23\x3f\x52\x04\xac\x21\x27\x13\x6b\x93\x81\xce\xc9\x78\xb5\x6e\x03\x82\xff\xf4\xc5\xcf\x3b\x45\x37\x39\x09\xa6\x08\xc8\xf1\x81\x16\x3e\x43\x0e\x6f\xfb\x1c\x0d\x66\xc4\x64\xaa\x5c\x38\xee\x64\xae\xe9\xa6\x39\x2e\x69\x3c\x94\xd8\xca\x8f\x7b\x10\x13\x1a\x91\x0c\x26\xd7\x26\x22\xf3\x68\x15\xe8\x73\x6d\xb7\xeb\xc5\x63\x43\xc8\x04\x08\x8a\xae\xab\xa4\xdd\xaf\x52\x7c\x28\xc8\xf2\x23\x1d\x35\xe8\x36\x5c\xbb\x4e\x43\x88\x90\xb1\x22\x71\x7c\xd2\x50\x34\x07\x36\x05\x8d\x7a\x7c\xdc\xef\xe1\x6e\xff\x0e\x29\x0b\xe3\xde\x5a\xd7\x53\x23\x82\xcd\x0e\x12\x99\x6d\xdd\x7d\x61\xe8\x49\xab\x1f\x5d\xf9\xf2\x2a\x84\xb4\xb2\xb7\x27\xaf\xee\x2a\xf1\x01\x11\xfc\x77\x37\x82\x2a\x5d\x31\x35\x84\xd1\x0a\xe8\x46\xda\xc3\x34\xf0\xbd\x3a\xc8\x43\x05\xc5\xae\x26\x02\xc2\xa5\xac\x9b\xca\x38\xa1\x76\x99\x3b\x0c\xa4\x41\x73\xde\xf4\x57\x19\xfd\xbe\x7b\xd8\xfb\x3f\x31\xa0\xaa\x67\xad\xc5\x0a\xaa\x14\xf9\x3a\x63\x5e\x15\x1c\xb1\x2e\x54\xb4\x71\xe1\x37\x35\x63\x67\x1d\xbc\x2b\x26\x87\x41\x70\x7f\x37\x82\x57\xba\x8e\xce\x6b\xc9\x09\xd8\x9c\x84\x8c\x25\x45\x20\xa3\xef\xa5\xb2\xdc\xd6\xd9\x34\x47\x16\xc9\x56\x51\x78\x97\x47\x7e\xa2\x53\x5f\x69\x9c\x0b\xf2\xa3\xda\xaf\xb7\x6c\xac\x7b\xed\xac\xf2\x26\xc0\xe8\x1f\xf1\x86\xd4\xb9\xa2\x44\x6b\x0b\x21\xc8\xda\xb8\x48\xd3\xd5\xad\x29\x78\x77\x83\xdf\x37\xd6\xfa\xae\x56\xdf\x0a\x10\x7f\xc1\xe9\x75\xab\xc7\xf8\xc2\xe3\x26\xd6\x5b\x31\xad\x6e\xaf\x41\x3e\xb9\x7b\x90\xd1\x78\x54\xff\xe2\x98\xd6\x0d\xe8\xcc\x49\x5f\xd7\xe2\xb2\x03\x95\x6a\xf1\x2a\xa7\xb0\x71\x1d\x35\xb4\xa6\x39\x2c\x3c\x59\x79\x18\x1e\x9e\x6a\x93\x21\xcf\xe1\xf9\x4f\x6f\x47\x6b\x78\xa5\xd0\x53\xbf\x74\x7c\x43\x31\x19\x52\x82\x36\xa5\x87\x8d\x64\x79\x94\x1b\x97\xa0\x2c\x9b\xad\x41\x46\x4d\x2b\x05\x26\xcb\x46\xaa\x65\x75\x7c\x29\xd5\xa7\x89\x12\x67\xa0\x3e\xa2\x17\xb2\xde\xc8\x0c\xd4\x6d\xd0\xc5\x39\x2e\x69\x9b\x48\x2a\xa6\x25\xd5\x71\x64\xe5\xc7\x01\x54\xac\x19\xd3\x4f\x91\x55\x5d\xbf\x31\x5a\x38\x4d\x1b\x8f\x4e\xa7\xc6\xa3\xbb\xbc\xf1\xe8\x6f\x69\x3c\x4b\xa6\x2c\x24\xad\x0f\xc1\xb5\x5c\x4c\xd3\xd3\xb8\x7f\x0d\xb7\x0e\xdd\x56\x08\xd4\x2b\xb3\x0e\x43\x77\x9b\xda\xe5\x59\x44\x9b\x09\xd3\x69\x6e\x87\x1f\xb7\x32\x6e\x07\x69\xd5\x0f\xce\x37\xc3\xac\x83\xc1\x43\x6f\xc6\xc8\x1e\xf0\x9b\x3b\xef\xbd\x30\x7b\xcb\x77\x29\xe6\x57\xfb\x1b\xe7\x1d\xa4\x83\xcb\xc9\xba\x4e\x1f\xc8\xb3\xbe\xb5\x51\x27\xce\x2d\x8f\x89\xb0\xee\x3b\x47\x1e\x44\x0e\x10\x97\x85\xc9\x4d\x8e\xc7\x2c\x33\x6a\xfc\x5a\xce\x93\x7f\x56\x58\xf3\xd3\xc6\xf7\x14\x1d\x1b\x3a\xae\xff\xfa\x75\xdd\x9c\x7c\xbb\xa6\xfc\x3d\x4b\x68\x87\xef\x5b\xd6\x6f\xf7\xba\x3a\x53\x15\xba\x66\x54\x8d\xa3\x52\xcd\xdd\x5b\x58\x52\x3e\x6e\x74\xf4\xb7\x09\xb9\xb9\x49\xc7\x60\x48\x68\x13\xb5\x4b\x65\xf3\x05\xa9\xdd\x9b\xb7\xa2\xa1\x59\xd1\x83\x0e\x8d\x7a\xea\xb4\xf8\x50\x90\x59\x75\xa9\xf1\x1a\x97\x04\xaa\xc8\x16\x64\x6a\x5d\xc2\xc7\xc1\xeb\x84\xd4\xc6\x01\xdd\x08\xa2\xc8\x36\x06\x3f\x77\x4b\xb3\x56\x77\x2b\xda\xee\x19\x11\xc5\x58\xa4\x3c\x87\x93\xea\x28\x93\x4a\x66\x45\x56\x1f\xd5\x76\x88\x31\xb5\x41\x7e\xb3\x23\x05\x94\x8d\xab\x7b\x51\xbe\xc4\x1b\x27\x7e\x0b\xa8\x75\xa3\xb8\xf1\xdf\x44\x3f\x13\xc1\x6c\xb6\x8d\x61\xd6\x87\xc1\x7f\xa1\x69\xa1\xf0\x67\x3b\x70\x74\x09\x69\xa1\xfb\xed\xb8\xd2\xe1\xac\x74\x8d\xf5\x9b\xcc\x20\x18\x84\x91\x4c\x46\xe2\xc4\x07\x9d\x5d\x29\xc6\x1b\x67\x03\x4e\xa4\xad\x83\x19\xa4\x6d\xf4\xfe\x4c\xa6\x68\x9c\x75\xb8\xc5\x42\x70\x7e\x9d\x90\xa1\x73\x10\x29\x16\x96\xdc\x29\x2a\x38\xfb\xe1\x05\x58\x46\xa6\x8c\x14\x1f\x55\x82\x0a\xbb\xde\xaa\x3a\xa8\x76\x2d\xc2\x0d\xa0\x80\xcc\x46\x2e\x0a\x26\x0b\x53\x10\x3a\x2d\x32\xb5\x49\x85\x42\xe8\x42\xf1\x04\x2a\x71\x4f\xb5\x81\x72\x21\x7a\x04\x52\x81\xff\x80\x55\xfa\xd0\x48\xba\x22\x57\x41\x9a\xbc\x36\xbc\xaa\x20\x14\x96\x8c\x13\x5e\x43\x64\x34\x7e\x80\xf6\x04\x17\xd9\xea\x62\x3e\xaa\x7e\xbc\xb8\xb8\xb0\x1f\xd2\x06\x8a\xc0\x0c\xa9\xbc\x24\x18\x67\xab\x7f\x8c\x9b\xa4\x35\xdf\xdb\x6d\xa3\x83\x40\x05\x98\x5a\x0d\x0b\x0a\x43\x38\x45\xa0\x5d\x62\xa5\x7e\x73\x60\xc8\xea\xc2\x08\x9a\x7c\x06\x48\x5b\x2c\xaa\x30\xb0\x90\xe2\x82\x52\xf2\x5b\xd7\x8b\x58\xeb\x07\x0b\x34\x17\x47\x3b\x31\x35\x79\xcf\x3d\xab\x9d\x5c\xd2\x0a\x1e\xc0\x38\xd6\x7a\x0c\xa8\xa2\x4e\x9a\x2b\x4c\x0b\x72\x54\x0b\x34\x3b\xac\xf0\x2c\xb8\xaf\x19\x59\x6a\xcc\xae\xd8\x5e\xc9\x88\xa2\x23\xd0\x06\x64\xa0\x09\xd2\xa4\x05\xca\x72\x5e\x1d\xb9\xb3\xfa\x4d\x78\xcb\x97\x9c\x20\xfb\x13\xe7\x10\x48\xd0\xba\xd7\xe8\x4c\x5a\x2b\xb5\x72\x06\xb2\x44\x70\x2d\xd3\x14\x16\xb5\x9f\x43\x76\x53\x34\x19\x5a\x4b\xcb\x8f\xa2\x9b\x29\x5a\x1e\x7e\x85\x1c\x0d\xde\x5d\xac\xbe\x78\x96\xae\x05\x0f\x4b\xd4\x45\xc1\x7b\x27\x6b\x2b\x4d\xf7\x0c\xe0\xca\xab\xfe\xe7\x10\xb7\xeb\x44\x1b\x90\x8a\x68\x45\x77\xf4\x9d\x9a\xcf\xbb\x13\xce\x51\x45\xe7\x10\x4b\x63\x19\x86\x2b\x71\x14\x38\x5e\xf5\xea\xf4\xa5\x32\x42\x69\xa0\x1b\xf7\x1e\x2f\x39\x40\x08\x05\xcc\x47\xfc\xba\xb8\x0c\x0a\xf4\x3f\x02\x00\x00\xff\xff\xb1\xbd\xae\xaa\x7c\x25\x00\x00")

func ocmExampleServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_ocmExampleServiceYaml,
		"managed-services-api.yaml",
	)
}

func ocmExampleServiceYaml() (*asset, error) {
	bytes, err := ocmExampleServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "managed-services-api.yaml", size: 9596, mode: os.FileMode(436), modTime: time.Unix(1575573384, 0)}
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
	".managed-services-api.yaml.swp": OcmExampleServiceYamlSwp,
	"managed-services-api.yaml": ocmExampleServiceYaml,
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
	".managed-services-api.yaml.swp": &bintree{OcmExampleServiceYamlSwp, map[string]*bintree{}},
	"managed-services-api.yaml": &bintree{ocmExampleServiceYaml, map[string]*bintree{}},
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

