// Code generated for package i18n by go-bindata DO NOT EDIT. (@generated)
// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x59\xcd\x72\xa3\x3a\x16\xde\xf3\x14\x0a\xa9\xae\xd9\xf4\x74\xdd\x75\x76\x34\x26\x09\x13\x1b\x28\xc0\xb9\x9d\xd9\x50\x32\xc8\xb6\x26\x20\x51\x92\x48\xda\x77\x37\xef\x35\xef\x34\xaf\x30\x75\x24\x01\xc2\x76\x6e\xe7\xd6\xac\xba\x6d\xeb\xfc\xe8\xfc\x7c\xe7\x3b\xca\x6d\xcd\xbb\x8e\x33\x2f\x09\x36\x51\x15\xfd\x88\x8b\xb2\xb8\x43\x7e\x82\x3b\x82\x70\x2b\x08\x6e\x4e\x88\xfc\xa4\x52\x49\xdf\x8b\xb3\x2a\x49\xcb\xf9\x50\xd6\x12\x2c\x09\xda\xd3\xb6\x45\x94\x21\x75\x24\x48\x90\x03\x95\x4a\x9c\x50\x9c\x21\x6e\xbe\x92\x27\xa9\x48\x87\x24\x51\x8a\xb2\x03\xea\xf1\x81\xf8\x9e\xe7\xdd\xd6\xed\x20\x15\x11\x5e\xb8\xde\x16\x65\x94\x57\xab\x68\x1d\x95\x51\x75\x1f\xc4\xeb\x68\x75\x87\xfc\x1a\x33\xc4\xb8\x42\x0d\x69\x89\x22\xc8\x1e\x07\x43\xf5\x20\x04\x61\x0a\x49\x85\x15\xf1\x27\x05\x71\xa1\xdd\xcb\xb7\x49\x12\x27\x0f\x77\xc8\x2f\x8f\x8e\x98\xd4\xca\xc4\xc0\x18\x65\x87\x0b\xa1\x75\x1a\x06\xeb\x3b\xe4\xc7\x5d\xcf\x85\x9a\xa4\x6a\xcc\x40\x6a\x47\xd0\xd0\x1f\x04\x6e\x48\x33\x4b\x26\x69\x75\xbf\x8e\x7e\xdc\x21\x3f\xc4\xec\x6f\x0a\xed\x29\x6b\xd0\xbe\x25\x3f\x11\xed\x7d\x2f\x49\xab\x4d\xa0\xcf\x05\xcf\x41\xbc\x0e\xbe\xaf\xa3\x3b\xe4\xdf\x63\xda\x92\x06\x29\x6e\x4e\xe3\x37\x4c\x5b\xbc\x6b\x09\xea\xb0\xb6\xc7\x78\x63\x83\x23\x48\x43\x98\xa2\xb8\xf5\x16\x71\xa9\xf2\xa8\x48\xb7\x79\xb8\x54\x66\x42\x74\x83\xca\x23\x11\x04\xae\x8a\x19\xc2\x52\xf2\x9a\x62\x45\x1a\x74\xe4\x52\xa1\x81\x35\x44\x20\x75\xa4\x12\xbd\x92\x93\xff\x81\xda\xea\x9f\x69\xf2\x97\x74\xff\xc1\x19\xb9\xa2\xfb\x3e\xd8\xae\xcb\x2a\xcc\xa3\x55\x94\x94\x71\xb0\xae\xc2\x20\xd1\x81\x36\x66\xef\x90\xbf\x22\x7b\x3c\xb4\x0a\xcd\x37\x75\xa2\x6d\x8c\x36\xbe\xa9\xca\xf0\x31\x0a\x9f\xe6\xc2\xd0\x69\x9d\xa5\x18\x94\xea\x2c\xaa\x4b\x4e\x57\xaf\xd4\xff\x1f\x24\xc4\x15\x77\xc4\xf7\xb2\xa0\x28\x7e\x4f\xf3\xd5\xe4\x4c\xb2\x5d\x43\xd2\x7b\x2c\xe5\x3b\x17\x0d\x1a\x4b\x6e\x47\xd0\xae\xc5\xec\xf5\xbf\xff\xf9\xb7\xef\x65\x79\xfc\x1c\x94\x51\xf5\x14\xbd\x9c\x0b\x82\x27\xbd\xa0\x6f\x58\x11\xb8\xb8\xe3\xc5\x2c\xee\xdd\x42\xf8\xbd\xc7\xb4\x28\xab\x60\x9d\x47\xc1\xea\x65\xee\xa0\x47\xc8\xcc\x79\x9b\xd9\xcc\x68\x89\xe9\xd2\x57\x13\x62\x32\x0b\x39\xb1\x2a\x9c\xc4\xbc\x53\x75\xd4\x01\xb0\xb5\x7c\x4d\x6f\xf5\xfd\xa5\xca\xf2\xf4\x1f\x51\x58\xfe\x5f\x26\x7a\xc1\xff\x45\x6a\xe5\x7b\xc5\x4b\x51\x46\x9b\xca\x02\xc5\x7d\xba\x4d\x56\xd7\x71\xa2\xe5\x35\x6e\x01\x24\xf6\x54\x48\xa5\x03\x95\xa4\x20\xe7\x76\x4b\xc2\x51\xdc\xcf\x4d\xe2\x7b\x71\x61\x1a\x55\xdf\xc1\x81\x08\x6a\xba\xd6\x28\x05\x87\x7d\x13\xef\x78\x93\xa5\x79\x59\x45\x79\x9e\xe6\x63\xce\x12\x8e\x1a\xac\x30\x5c\xd3\x8a\xbd\x63\x89\xf6\x7c\x60\xcd\x0d\xb2\x9e\xd6\x47\x52\xbf\x6a\x3f\xed\x91\x3d\x6d\xc9\xcd\x52\x29\xa8\xab\x9e\x83\xf5\x16\x3c\x8d\xba\x5e\x9d\x8c\x5e\xce\x50\x4b\x19\x41\x5f\xe4\xd9\xf9\xb4\xac\xc2\x74\x93\xe9\x1c\x8c\x72\x31\xab\x79\xd7\x6b\x7c\xfb\x53\xe1\xdf\xf3\x34\x79\xa8\xee\xd3\x7c\x13\x94\x56\x4c\x08\x52\x2b\x64\xbc\xe3\xa2\xc3\xea\x43\x61\xa7\x0b\xdd\xac\x84\x4e\x0b\x71\x65\x22\xf0\xa1\x0e\x5b\x25\xcb\xb4\x9a\xac\x7f\x42\xda\x56\x5b\xb2\xdd\xdc\x21\x3f\x40\x8a\x2b\xdc\x22\xbe\x47\x5f\x24\x12\xfc\x5d\xc2\x7f\xf5\xf5\xb1\x20\x08\xef\x18\x5c\xa7\xfd\x8a\xe4\x2b\x00\xa9\xd6\x33\xc1\xd3\x58\xb7\x71\xb2\x6c\x8a\x1d\xc0\xa9\x2e\x55\x41\x24\x1f\x44\x0d\x4e\x7c\x45\x82\x60\xc9\xd9\xdd\x07\xfe\x14\xc1\xf3\x12\xeb\x24\x7e\xb3\x05\x7f\x26\xec\x79\xb7\x80\x23\x33\x82\x40\x1c\x36\x41\x19\x3e\x8e\x28\x30\x42\x08\x95\x88\x8e\xd9\xf1\xbd\x34\x8f\x1f\xe2\xc4\x06\xde\x3d\xcf\x05\x3d\x50\x86\xdb\x8f\x04\xb7\xc5\x3c\x9b\x82\xb0\x8c\xb5\xa3\xe5\x08\x67\x76\x98\x11\x06\x6d\xe1\x94\x2d\x67\x0a\xd7\x4a\x17\x2e\x6e\x3a\xca\x60\x14\x63\xc5\xc5\x8d\x55\xe8\x66\x2f\xe1\x48\x0e\xf5\x51\x2b\xd4\xfd\x17\xac\x36\x71\x72\x89\xd3\x60\xb4\xb1\x58\xad\x95\x1a\x17\x2e\xb0\xfa\x66\xe9\x74\x1e\xad\x83\x32\x5a\x39\xf0\xb2\x05\xb1\x23\x06\xd7\x5d\x10\xb1\xd8\xa1\x5d\x58\xaf\x82\x6c\xf2\x60\x9b\xad\x82\xc9\x83\xb6\xc1\xfd\xb9\x61\xd2\x50\x63\xf7\x39\xca\xe3\xfb\x97\x2a\x4c\x57\x0e\x7d\x78\x26\x82\xee\x69\x8d\x15\xe5\x0c\xd5\xbc\x21\x88\x08\xc1\x85\xef\x45\x9b\x20\x5e\x57\xab\xb8\xb0\x28\xb3\xc1\xb4\x1d\xd9\x89\xd4\x25\xd8\x50\xf9\xc9\xc0\x8e\xda\xdc\xf4\x46\x1d\x28\xec\xb0\xaa\x8f\x68\xaf\x4b\xcb\xc0\x1b\x4c\xb2\xa9\x7e\x0a\xf8\x34\xf9\x0a\xa1\xf9\x93\x31\x36\xd6\xc8\xb9\x12\x8d\x6b\x77\xc8\x7f\x17\x9c\x1d\xe6\x41\x87\xb8\x70\x44\x8c\x83\x7a\xe2\x4c\xce\x9d\x4f\x9c\x24\xad\x9e\xb6\xdf\xa3\x2c\xae\x74\x11\xcc\x90\x2d\x89\xb9\xf4\xeb\xb0\x23\x3d\x45\xb8\x6b\x28\x73\x4a\x96\x8d\xe4\xae\x18\xc3\xa7\xe1\xdc\x55\x38\xe7\x7f\x9a\x02\x42\xaa\x85\xe2\x2c\x46\x2d\x3f\x50\x86\x28\x33\x48\x06\x19\xb3\x73\xc2\x56\x07\xea\x30\xc3\x07\xd2\x11\xa6\x60\xa2\x02\xcb\xe4\x6c\x1c\x68\x79\xf4\x10\xa7\xc9\x67\x09\x12\x32\xc2\xbf\x1a\x69\xc0\x6b\xc0\x14\xfc\x3b\x1a\x02\x6e\xf4\x69\x33\x9a\x18\xfd\x6a\x6e\xb6\x98\x2d\xa9\xa8\x99\x51\xa1\x29\x83\x83\x0d\x92\x9d\xe0\x57\xc6\x53\xcd\xd9\x9e\x1e\x06\xa1\x63\xa6\xcb\x2c\xde\x04\x0f\xd1\xc7\xaa\x68\x87\x0f\xe4\x73\x8a\xb2\xaa\x78\x4c\x73\x33\x6e\xe4\xb0\xdf\xd3\x9a\x02\xe9\x8e\x7b\x08\x0b\xef\x09\x93\x0a\xd7\xaf\xde\x43\x54\x8e\x19\x18\xeb\x31\xe1\x63\x90\xf5\x58\x80\xf3\x36\x8f\x1b\xd2\xed\x88\x98\x80\x22\x58\x41\xf5\x7f\x91\x68\xc2\x86\x1d\x21\x0c\xe1\x46\x13\x6d\x17\x4e\x46\xd4\xfa\x22\x17\x08\xa8\xf5\x5b\x9a\x65\x4d\x4c\xe4\x73\x1c\x59\x1f\x32\xcf\xb1\xb6\xae\xd0\xce\x51\xf6\x31\x28\x2a\x9b\x1e\x67\xe0\x39\xa9\x9c\x52\x13\x5e\xc1\x43\xef\x16\xf8\xbc\x97\x00\x2e\x8d\xd4\xcf\x6e\x27\x55\x19\x14\x4f\x30\x0c\x9b\x46\x93\x7e\xe8\x59\xbb\xe8\xe8\x8f\x63\xd5\xd8\x7d\xe5\x6b\x6f\x12\xf6\x8e\xa9\x42\x54\xa1\x86\x33\xf2\x0d\x0c\xec\x70\xfd\x3a\xf4\x41\x5d\xf3\x81\x29\x2f\x0b\xf2\x60\x53\x45\x9b\xac\x7c\x39\x4f\x5b\x8f\x05\xee\x88\x22\x42\xb7\x7b\x59\x15\xdb\x2c\x33\xd9\xdd\x32\x39\xf4\xc0\x23\xa0\x86\x4f\x3d\x2c\x54\x4b\xbe\xbd\x40\x52\x83\x68\x13\x99\xfc\x1e\x84\x4f\xdb\xac\x0a\xc2\x30\xdd\x26\x7f\x85\x56\x2e\x1c\xff\x34\xbf\xf4\x6e\xa1\x65\x46\xe3\xd9\x3a\x48\xae\x31\x64\xe3\xa3\xb1\x03\xe7\xe7\xca\x1a\x24\x69\x46\x68\x19\x77\xbc\x71\x79\x71\xcc\x2c\x77\xa3\x4f\xdc\x46\x5b\xf9\xfc\x25\xac\xe9\xef\x3a\x06\x67\xa1\xfc\xc5\x7d\x4c\xdc\x10\xb6\x81\xfb\xab\x37\x1b\x8d\xc4\xeb\xa8\x70\x37\x19\x0b\x08\xb6\x04\xd5\x6c\x08\x98\x2f\xda\x91\x3d\x17\x04\xc9\x77\xaa\xea\x23\x2c\xf1\xea\xc2\x93\x05\x8c\x19\x2b\x97\x1b\xf8\x8e\x80\x30\x08\x92\x06\x0d\xbd\x6e\x5c\x47\x2c\x8f\x8a\x32\xcd\xa3\x4b\x39\x41\xa4\xe2\x82\xb2\xc3\xb2\xd5\x73\x4b\xf5\x2e\xf3\xe5\x5c\xf3\x97\x97\x9b\x37\x90\xeb\x0b\xd2\xdc\xff\xd3\x3a\x34\xa6\x79\x47\x5a\x0e\x53\x4f\xf1\x25\x4e\x97\xb0\x09\xf3\x9e\x08\xcb\x3e\x26\x6c\xe8\x89\x80\x11\xa7\xd1\xc1\x10\xd2\x0b\x78\x7b\xb4\xfb\xd6\x04\x6f\xbe\x37\xb1\x5f\x03\x99\x91\x3d\x39\x5e\x5f\x17\x9e\x81\x4c\xdd\x11\x97\x3a\x33\x5b\x9e\x8e\xce\xb3\xc6\xbd\x94\xf9\x7e\xde\xa0\x8e\xb0\x77\x0b\x7c\xcc\xd0\xb5\x09\xa5\x4d\x21\x15\x27\x56\x1f\x05\x67\xf4\x0f\x48\xb1\x24\xc2\x50\xaa\xdf\x2c\xb9\x5b\xa7\x0f\x71\x72\x2e\xb3\x75\x39\x2d\x50\x92\x1b\x7b\x7a\x26\x69\xe5\xfc\x90\xd4\x0b\x7e\xa4\x3b\xaa\x24\x82\x33\xd6\xc6\x5e\xf0\x0e\x58\xc4\x01\x0a\x8c\xb2\x6f\x9f\xa1\xc4\xae\x89\xaa\x28\x83\x72\xeb\x3c\x6a\x19\x6e\x6d\x4c\x58\x52\x03\x8d\x4b\xa5\x17\xc6\x85\x46\xef\x73\x48\x87\x7d\x94\x4a\xa4\xb0\x7c\x3d\x87\x6f\x88\xd8\x5b\x17\xea\x19\xeb\x3d\x6f\xaa\x30\x4d\xee\xe3\x87\xf9\x19\x20\x74\xa7\xef\x05\x3b\x9b\x05\xce\x9f\xc8\xca\xf3\xc9\xfd\x51\x65\x36\xa4\x6f\xf9\xa9\xd3\x83\xa0\xc5\xec\x93\x15\xea\xdd\xd2\x1e\x78\xc0\xe4\x27\xd8\x23\x4c\x11\x41\x1a\xd8\xdd\x25\x39\x68\x95\xe0\x42\x4b\x6b\x25\x67\xb0\xd3\xbe\x43\x2e\xe6\x63\x5f\x51\xbf\x7c\x07\xc0\x07\x4c\xd9\xf4\xaa\xe8\x2e\xfe\x71\x06\xab\x21\x78\x58\xd7\x43\x4f\x49\x83\x30\x6b\x1c\x27\x05\xd1\x9a\x1a\x2d\x1c\x27\xcf\xc1\x3a\x86\x70\xc4\xbd\xd9\xa2\xde\x70\x4b\x9b\x91\xc4\x38\x4f\x35\x6c\x00\xf6\x01\xdb\xe6\x81\x30\xb8\xba\xb9\x06\x6e\x1a\x41\xa4\x24\xda\xe2\x6f\xdf\x2e\xc9\x91\x54\x58\xe8\xcb\xd8\x93\xda\x1b\x39\xec\x18\x81\x4a\xd5\x61\xfa\x7b\xcf\x79\x0b\xe6\xb2\x34\x5d\x5f\xcd\x53\x9c\x21\x38\xe3\xb0\x9c\x2b\x33\x62\x7e\x10\xd4\xf4\x71\x79\xeb\x89\x9d\x18\x02\x2c\x95\x38\x79\xc0\xbc\x8a\x32\x7f\xb9\x7c\x5a\x2a\xdd\xd7\x58\xbe\x37\xaf\x73\x58\xd4\x47\xaa\x48\xad\x06\x41\xf4\xb0\xbf\x28\xb6\x0b\x85\xdb\x62\xba\xc2\x27\xd4\xe9\xf8\x43\x4f\xfa\x5e\x12\xfd\xd8\x16\xd5\x36\x59\xbc\xe8\x90\x9f\x83\x84\x7a\x61\xa4\xb6\x34\xdf\xa1\xfc\xce\x06\x3c\x15\x0b\x98\xd3\x85\x72\xe3\x0c\xce\x09\xf8\xc7\xe1\x91\xe5\xe9\x73\xbc\x8a\xf2\x89\x82\xba\x03\xa4\x16\x44\x47\x19\x4a\x6e\x50\x1c\xac\xd5\xa8\x03\x96\x65\x63\xdb\x61\x36\xe0\xb6\x3d\xc1\x97\x74\x7f\x5a\xbe\x23\x48\x07\x7f\xcb\x97\x2c\x5a\x98\xd0\xb0\x3b\xdd\x1c\x74\x4d\x5c\xea\x06\xa5\xac\x3d\x8d\x9f\x25\x02\xb8\xff\x8a\x96\xb0\xfb\x6b\x1a\x35\x0f\x9f\xfd\x9f\xd0\x28\xa4\x41\x8d\x1c\xae\x11\x0f\x18\xbb\x4b\xce\x30\x0d\x62\x5c\x2b\xfd\x76\x69\x7a\x5d\x12\x29\x61\xab\x2a\xa2\xa2\x00\x36\xff\x14\xbd\x2c\x66\x82\xfd\x5d\x3f\x74\xda\x1e\xb3\x1c\x3d\xcb\x53\x98\xb4\x4e\x27\x8e\x67\xcd\xc6\xce\xd9\x1b\x11\xd2\x65\x8f\x5a\x0c\xc8\x75\x92\xba\x9b\x94\x3a\x7b\xe0\xd8\xe9\x57\x25\xc5\x11\x9e\xc8\x3a\x17\xf3\x83\x26\x38\xad\xb8\xc0\x07\xe2\xe9\x02\x00\xb7\xa1\x06\x7e\x14\x31\x50\xb5\xc2\xfc\x06\xa2\x6f\x14\x98\xd0\x79\xad\x6b\x14\x8f\x93\x2a\xfa\x11\x85\xdb\x32\x4e\x13\xf3\x3a\x05\x00\x3e\xc5\x88\xfc\x24\xf5\xa0\x1c\x6e\x65\xff\xf6\x60\x9d\x98\xd0\x62\xaa\x53\xd4\x62\xf8\xfe\x9b\xa6\xe7\x35\xef\x7a\xce\x08\x53\x5e\x98\x6e\xb2\x34\x89\x92\x72\x5a\xc5\x0d\x78\xdb\xdf\xcf\x7c\xfb\x86\x56\xdc\xfc\xbd\x82\xf4\x04\x9b\xf1\x35\xa1\xf5\x37\xff\x7f\x01\x00\x00\xff\xff\x75\xf7\xc1\xc0\xb4\x19\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 6580, mode: os.FileMode(420), modTime: time.Unix(1659003327, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x58\xcd\x52\xdb\x4a\x16\xde\xeb\x29\x54\xa6\xee\xee\xce\xad\x59\x67\xa7\xd8\x82\x68\xb0\x25\x95\x24\x73\xc3\x6c\x54\x84\x78\x32\x4c\x82\xed\xc2\x30\x35\x77\x56\x71\xc2\x8f\x01\x1b\x3b\x09\x66\x02\x38\x01\x13\x08\xbe\x10\xdb\x40\x12\xec\xf8\x07\x5e\x46\xdd\x92\x57\xbc\xc2\x54\x77\x4b\xed\xb6\x1c\x48\x36\x14\x96\xce\x39\x7d\xfa\xfc\x7c\xe7\x3b\x1a\x99\x4e\xcc\xce\x26\xe2\x9c\x2c\x44\x44\x53\x7c\x28\xe9\x86\x7e\x8f\x0f\x80\x42\xce\x3e\x3e\x03\x8d\x0b\x50\x7d\x0b\x4a\x95\x00\x27\xa9\xa6\xac\x18\x7d\x01\xa7\xde\x00\xa5\x8a\xfd\xb9\x6d\xb7\xf7\x9c\xda\x95\xdd\xad\xf5\xca\x5f\x7b\xef\x0e\x40\xf9\x14\x2c\x6f\x5b\xed\x37\xa0\xf5\x46\x52\x03\x1c\x37\x32\xfd\x6c\x21\x35\x1f\x9b\xe3\x82\xe1\xa8\x6e\x88\x9a\x19\x12\xc3\xa2\x21\x9a\xa3\x82\x14\x16\x43\xf7\xf8\x00\xfc\xdf\x3e\xfc\x5c\x04\x99\xfd\xde\xf6\x21\xe8\xbe\x01\xab\x39\x7b\xed\x12\x3e\x4f\xdb\x3b\x8b\xbd\xdd\x65\xfb\xea\x30\x40\x55\x25\x1d\x3b\xa1\x45\x65\x59\x92\xc7\xee\xf1\x01\x22\x60\x35\x73\xa0\x54\x71\xae\x0b\x4e\x39\x6b\x35\xab\x37\x9d\xf4\x90\x4a\x58\x09\x0a\x61\x74\xaf\x7a\x07\x2c\x1d\x11\x35\xf7\xe0\xdc\x8a\xdd\x3a\xee\x2b\xc8\x8a\x39\x1a\x16\x1f\x22\xc7\x4a\x27\xce\xcb\x2e\x5c\xbd\x02\x99\x33\xd0\xf9\x06\x9f\x1f\xf3\x33\xc9\x00\x27\x2b\x66\x44\xc0\xa2\xc2\x84\x20\x85\x85\xfb\x61\xd1\x2f\x9c\xaf\xdb\x9b\x15\x7b\x67\x91\x9f\x9d\x42\x57\xe7\x9d\xb5\x17\xf6\x8b\x6f\x01\x0e\x45\x63\x2e\xf6\x38\x16\x9f\x9f\x99\x7a\xc6\x0d\x04\xc2\xd4\x44\x5d\x89\x6a\x41\x64\xcb\x8d\xc5\xe1\xb9\xf3\xe5\xe8\xa6\x93\x76\xea\x47\xf6\xf1\xdb\xde\xeb\x23\xab\xb9\x0e\x4b\xab\x60\xe9\xb3\x93\xde\xb4\x9a\x6d\x58\x6a\x05\x6e\x31\x62\xfe\x5d\x91\x7f\xd6\x12\xf1\x16\x64\xb1\xb1\x51\x21\x1a\x36\xcc\xa0\x26\x86\x44\xd9\x90\x84\xb0\x19\x14\x64\x1c\x40\x72\x0e\x0a\x79\xfb\xad\x53\x3b\x04\x2b\x55\x98\xab\x59\xcd\x9c\xf3\xb2\x4b\x0e\xc1\x51\xc7\x45\x14\x7c\x20\x06\xc7\xfb\xf9\x25\x27\x92\x82\x22\x0a\x56\x73\xc3\xde\xac\xc0\x4c\x03\x3d\xdc\x6d\x82\x42\x36\xc0\xa9\x82\xae\xff\xae\x68\x21\x7a\xa0\x1c\x0d\x93\x84\x2d\xdb\xfb\x69\x4f\xaf\x65\xff\xd9\xc2\x07\xa9\x9a\x34\x21\x18\xa2\x39\x2e\x4e\xfa\x35\xbc\x1b\xfa\x34\x38\x6e\xe4\x9f\x89\xd4\x3c\xf7\x40\xd1\x0d\x53\x08\x6b\xa2\x10\x9a\xec\x97\x33\x09\x27\x53\xef\x6e\x5c\xb1\x34\xbd\xca\x70\x38\xa9\x9e\xdd\xce\x93\x70\x7a\x35\x3b\x6c\xc0\xbc\x3f\x69\xaa\x9a\xf2\x37\x31\x68\xfc\xac\xad\xf2\x37\x7b\xb7\x86\xdd\xd7\x27\x75\x43\x8c\x98\x6e\x1b\x8e\x2a\x51\x39\xe4\x76\xe1\x52\x86\xf4\x1c\x2c\x7d\x82\xa5\x96\xa4\x92\x44\x28\x48\x94\x2d\x50\x49\xe5\x9d\xaf\x8b\xb0\x55\x40\x91\xb9\xfc\x1c\xe0\x24\x9d\x74\x05\x76\x11\xdb\x72\x3d\xb0\x9a\x1b\xa4\x7f\x79\x49\xe5\xc1\xf2\x85\x7d\x92\xbe\xe9\x64\xdd\x02\xcf\x14\xc0\xda\x1e\x6c\xb4\xc1\xda\x7e\x80\xc4\x52\x8a\xa8\x8a\x66\x98\xa2\xa6\x29\x9a\x97\x03\x58\x3a\xa1\xad\x40\x5a\xce\xde\x59\x84\xc5\x33\x98\xab\xe1\xbb\x36\xe0\x87\xe7\x70\xef\x88\xbc\x82\x5b\x2b\x56\xfb\x12\xbb\xcd\x1a\x44\xa6\xcc\x09\x21\x1c\x45\xde\xff\x92\xe2\x9d\x72\x16\x96\x56\xed\x3f\x5b\xc4\x8e\x4f\x58\x31\xcc\xa0\x12\x51\x71\xcc\x7d\x4a\x58\x1c\x61\x44\x2d\x0b\x8b\x5f\x06\xf5\x7e\xd7\x14\x79\xcc\x1c\x55\xb4\x88\x60\x50\x0d\xfb\xb4\x0e\xf2\x1f\xe0\x7e\x07\x74\xf2\x56\x33\x07\xab\x1f\xec\xb2\xef\x3c\xa6\x43\xd8\x7c\xb8\x27\xae\x5e\xa1\xe3\x32\x67\xa0\xbe\xdc\x7b\x7d\x34\xa8\xe9\x56\xc0\x5d\x6a\x24\xed\x83\x6a\x6e\x09\xc9\xd1\x08\xaa\x9d\xa5\x73\xde\x77\x39\x58\xfd\x00\x9a\xcd\x9b\x4e\xd6\x69\x7c\x76\xae\x57\x5c\x65\x8a\x07\x5e\x05\x4a\xf8\x3c\x92\x67\x52\x0e\xc8\x90\x75\xfd\x0e\xd5\xbd\x5b\x8b\x59\xb0\xb1\x07\x76\xf7\x6f\x3a\x3b\xbf\xa4\xbe\xeb\x84\x2e\x4c\x88\xd4\xca\x8f\xf4\x39\x6e\x64\x21\x15\x9b\xeb\x37\x38\xba\x78\x44\x30\x82\x0f\x68\x77\xf7\x36\xb7\x9d\x7a\x3d\xc0\x29\x9a\x34\x26\xc9\x6e\x48\xa9\xc8\xc6\xde\xa0\x54\x54\xef\xc3\xba\x10\x34\x24\xec\x0b\xc1\x13\x58\x3a\x01\x05\x84\x66\xa4\xc8\x9c\xf4\x26\x1a\x50\xb5\xb2\x5d\x58\x06\xaf\xde\xe2\x0a\xc3\xda\x6c\xec\xd1\x14\xa8\x1e\x12\x7d\x2c\x21\x84\x22\x92\x7c\x1b\xee\xf1\x53\x8f\x67\x67\xe2\x3c\x11\x27\x18\xe3\x1c\x9c\x32\x08\xc8\x7a\xa7\x89\x61\xc1\x10\x43\x4c\xd3\xbb\x6e\x5e\x94\x29\xfa\x7a\xb9\x0e\x87\x04\x95\x1e\x1a\x55\x43\x02\x3e\x14\x3d\x1d\x38\xcc\xba\xae\xc1\xcd\x6f\xf8\xa4\x09\x51\x93\x46\x27\xcd\xa0\x12\x62\x86\x69\xef\x24\xeb\xd4\xd3\x4c\xb4\xc4\x88\x20\x85\xcd\x90\xa4\xbb\x30\xd0\x7b\x51\xb3\xda\x97\x64\x62\x3b\x07\xa7\xf6\xc7\xf4\x6d\xe1\xf2\x74\xd9\x64\x10\x6d\x90\xfd\xd6\x5b\xca\x51\xe8\x72\x61\x9f\x26\x58\x47\xbf\xfa\xf0\xef\x21\x3d\xc5\x7e\x92\x4e\x0f\xf8\x07\x75\x31\x8a\xb0\x5a\x30\xb3\x35\x98\x7e\xe2\x14\x06\x6e\xe2\x90\x5d\x3b\x67\x70\x5b\x56\xcc\xf1\xe8\x7d\x51\x95\x4c\x9c\xc7\x3e\x46\x0e\x92\x15\xab\x59\x25\xff\xf0\x4f\x17\x1e\xc5\x92\x33\xfc\xd4\xec\xe3\x99\x38\x4f\x8e\x62\xad\xf4\x73\x47\xed\x90\x94\x91\x38\x0d\xd9\xb1\xb7\xdb\xa0\x5b\xb4\xae\xcb\x30\x5d\xc7\xc5\x3f\x17\x7b\x32\x93\x88\x7b\xe3\x40\x13\xc7\x24\x45\xfe\xa9\x99\x0f\xb2\x2d\xb0\xb7\xc7\x8e\x03\x66\x52\x73\x23\xff\x4d\xc4\x63\x9e\x55\x34\xed\x7f\xce\xa6\x67\x61\x60\xca\xbc\xac\xd8\xdd\x0b\xa7\x56\x06\x99\xd7\x83\xc4\x89\x80\xb9\xb3\xd1\x00\xf9\x2d\x17\x9a\xf0\x74\x63\x31\xbc\xb7\x94\xb3\xbb\x64\x46\x49\x11\x61\x4c\xbc\x4d\xb1\x58\x02\x2f\xf3\xb7\x29\xaa\xa6\xfe\x40\xd1\x50\x94\xa5\x24\xef\x0d\x28\x8e\x1b\x49\x24\x63\xf1\xd4\xfc\xd4\xf4\x53\x6e\x4c\x34\xbc\xe0\x79\x35\xd2\x87\x59\x1c\x29\x14\x94\xe4\x5c\xe2\x5f\xb1\xe9\xf9\x48\x6c\xf6\x51\x6c\x8e\xf6\xa2\x10\x72\x41\xd6\xad\x2a\x7c\x77\x6f\x86\xb1\x0d\xcb\xe0\x31\x05\x14\x32\xc7\x08\x85\xf0\xec\x53\xa2\xe4\x81\xf9\x2d\x68\x41\x2a\x65\x88\x25\x79\x5a\x0f\x04\xdd\x74\xc3\x8d\x54\xb0\x30\x4b\x22\x6e\x3a\xe9\x21\x5d\x6e\x24\x9e\x78\x1c\xe3\x64\xd4\xf5\x1e\x8b\x71\xf9\xb0\x69\x08\xfa\x38\x9e\x0e\x97\x56\x7b\x8b\x70\x4e\xf8\x26\x67\x75\x4b\x68\x38\x94\x2a\x70\xf5\xd8\x29\x67\x7f\xe5\x9d\x7a\xc3\xae\xae\x82\xab\x25\x50\x7b\x69\xb5\x3f\x91\xc7\x68\x38\xd6\x8b\xbf\xa1\x03\x1e\x4d\x4d\x3f\x5d\x48\x0a\xd3\xd3\x89\x85\xf8\x3c\xa7\x0a\x9a\x10\x31\xc5\x88\x6a\x4c\x22\xdb\xf9\x17\xb0\x78\xe6\x65\x08\x5d\x59\x8f\xaa\x2a\x49\x1d\x9a\x42\x9b\x75\x98\x45\xcc\xdd\x3e\x6f\x83\xf7\xeb\x01\xce\xc7\x05\xe1\x7e\xb9\x77\x92\x65\x90\xc3\x2d\xe0\xfb\x42\x70\x3c\xaa\x9a\x42\x30\xa8\x44\xe5\x9f\xe5\x4b\xe0\x70\xc5\x6a\x77\x9d\x2f\x1f\x41\xbe\x71\x0b\x6b\xe2\x46\x92\xcf\xa6\x68\xf3\xa9\x61\x41\xbe\x8b\xcc\xb1\x8d\x80\x20\xc5\x6b\x75\xab\xb9\xee\x2d\x1a\x55\xab\x7b\x6d\x6f\x56\xfc\xbc\xfb\x07\x9e\xfa\x0c\xfb\x3c\x65\x96\xa4\xfb\x38\xf4\xbe\xa8\xdc\xe1\xb1\x2f\x04\x77\x7a\xec\x59\x93\xc2\xa2\xce\xf2\x65\x17\xd7\x5c\xdb\xc8\x1e\x21\x64\x60\x39\x07\x32\x2b\x30\x77\xc0\x1e\x32\x80\x0f\xc4\x22\x5d\xc3\x48\x95\x11\xe9\xef\xac\x61\x9a\xa8\x1b\x8a\x26\xfa\xc4\x61\xfa\x00\x1c\xe6\x3c\x71\xda\x61\x5a\x2c\x95\x58\x98\x9b\x8e\x0d\xc7\x99\xb9\xc6\x1d\xce\xb3\xc5\xe5\xa3\xe0\xfd\x8e\x1b\x20\xdc\xe7\xef\xad\xd6\x86\xaf\xef\x9c\xeb\x5d\xc4\xb1\xaa\x87\xa4\x8d\x18\x8a\x3a\x84\x17\x1e\x75\xce\x79\x83\x88\x52\x2f\x02\x40\xa2\xc7\xf3\x30\xe9\x02\x8d\x0b\x0f\x7f\x70\x51\x0e\xdb\x63\x4b\x86\xb1\xea\x6b\x94\x61\x3d\x36\x59\x8c\x1e\x37\xf2\xec\xf1\x54\x92\xb0\x0c\x8a\x78\xee\x86\x55\xc8\xc2\xea\x11\xc8\x9c\xa1\xb6\x75\xa9\x46\xeb\xaf\x2e\x23\x09\x2b\x63\x92\xec\xd7\xa0\x84\x84\x58\xc7\x61\xc1\xd2\x7d\x9a\x41\x66\xad\xfd\x31\x0d\xab\x07\xe8\x15\x51\xb1\xb7\xdb\xbd\xed\x65\xfe\x16\x6a\xc6\x9a\x30\x75\x43\x30\xa2\x7a\x3f\xc3\x98\xd6\x21\x09\xdc\x2d\x33\x29\x2e\x28\xe9\x18\xf0\xfc\x28\x88\x6e\xe4\xd1\x01\xb8\x7a\x0a\xf3\x79\xab\x59\xb5\x77\x16\xad\x76\x1b\xac\x95\x91\xf6\xbf\x67\x83\x89\xf8\x3f\x66\x9e\x70\x13\x11\x33\xa8\xc8\xa3\xd2\x58\x7f\x0f\x24\x43\x89\xe1\x13\x7d\x19\xff\x77\x0b\x2a\xda\xaf\x1e\x26\x67\x77\xd6\x10\x37\x32\x93\x44\x43\xaf\xff\xb9\x05\x6f\x70\xf6\xce\xa2\xa4\xc2\xda\x57\xc4\x91\x1a\x17\xb0\xb4\x8a\x7f\xd1\x0d\xcc\xa9\x37\x7a\x2b\x39\xb8\x75\x46\xa4\xe9\x37\x19\xdf\x96\x87\x8a\xec\xe0\x14\xe4\xf6\x31\xb1\xcb\x52\x92\xda\x5b\x59\x83\x9b\x57\x58\x4b\x92\x27\x84\xb0\x14\xc2\xf2\xa0\x74\x06\xde\x3d\xa7\x6b\x8e\x37\x8f\x69\xaa\xf7\x60\xa6\x80\x1d\x23\x82\xb0\x78\xd6\x5b\xc9\xa3\x0a\xe1\xe9\x24\x77\xbe\x36\xc0\xf1\x3a\x79\x0f\x5e\x67\x41\xb5\x60\x77\x5f\xd1\x8b\xfe\x25\x99\x48\x3c\x43\x26\x55\x45\x09\x0f\x45\x51\x4a\xf2\xf0\x7c\xff\xbb\x2c\x07\x2d\x9d\xcc\xb7\xa1\x80\x4b\xa6\x52\xf3\x73\x7f\x70\x88\x0a\xe8\x86\x36\x39\xbc\xc9\x3b\xf5\x23\xf8\xfe\x12\xbe\x77\x47\x10\xce\x3c\x5a\x66\xdd\x23\x68\x77\xfa\xf4\xa3\xba\x18\xba\x43\x1b\x56\x3f\x38\x07\xa7\x1e\x94\xca\xe2\xc3\xa8\x6e\x46\x65\x36\xf2\x72\xec\x3f\x0b\x29\xde\xb9\x7e\x0f\x37\x8e\x08\xfd\x63\x62\xea\xe6\xce\xa9\x17\xdd\xb0\xb8\x78\x4f\x61\xce\xc3\x49\x55\x53\x26\xa4\x90\xa8\x51\x92\xe3\xac\x9c\x80\xb5\x0a\xac\x94\x41\x27\x0f\x32\xbb\xa0\xdd\xa2\xdf\xc5\x48\x66\xe1\xea\x3a\x58\xab\x90\xbd\x80\xdd\xeb\x18\x04\x32\x26\x55\x91\x1a\xa4\x53\x9a\x00\x10\x9d\xd5\xa8\x60\xdb\x4b\xe4\x15\x8f\x40\xee\xd7\x41\xac\xf9\xf1\xac\xee\x03\xeb\x0f\x06\x15\x4e\xb3\xbd\xb3\x48\x1e\xda\xd5\x2d\xbb\x78\x84\x26\x07\xee\xf1\xfe\x28\x48\xc5\x52\x29\x44\x9d\x75\x51\xd7\x11\xef\x1b\x17\x27\x07\xd0\xce\x7d\xcf\x3f\x8d\xfd\xc1\xa3\x42\x29\x66\x5c\x36\xa7\x6a\x0a\x1a\x0f\x4c\x99\x7b\xa2\x04\x86\x9c\xee\x27\x90\xdd\x22\xee\xb9\x2a\x88\x8a\xc9\x0a\x4b\x9e\x29\x01\xb4\xdb\xaf\x40\x6d\x87\xae\xe5\x30\xb3\xe5\x7d\xe2\x41\x3e\xce\x27\xe6\xa6\x9e\xc4\x38\x9c\x37\xe4\x25\x4a\xdd\x43\x5d\x32\x58\x18\x42\x7f\x5f\x54\x60\xbe\x60\x5d\xed\x82\xe2\x72\x80\xc3\xb8\x25\xc9\xa6\xf8\x50\x0c\x46\x0d\x49\x41\x6b\x0a\xf9\xf4\xe9\x8e\xc6\xd2\x2a\x81\x2b\xca\xda\x70\x54\x50\x15\xd9\x95\x1c\x78\xbe\x4d\x6a\xe9\xb7\xdf\x30\x5f\x9b\x4e\xcc\x26\x13\xf1\x58\x7c\x9e\x0b\x2a\x11\x55\x91\x45\xd9\xa0\xfb\x91\xdd\x5e\x44\xa3\xdc\xf3\x84\xd8\x00\xeb\xd7\xbd\x95\x1c\x38\xcc\x31\x90\xf4\xff\x00\x00\x00\xff\xff\x2c\x00\x3d\xca\x07\x16\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 5639, mode: os.FileMode(420), modTime: time.Unix(1659003331, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
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
