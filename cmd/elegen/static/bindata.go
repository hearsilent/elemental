// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-bindata.
// sources:
// templates/README.md
// templates/identities_registry.gotpl
// templates/model.gotpl
// templates/relationships_registry.gotpl
// DO NOT EDIT!

package static

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var _templatesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x11\x02\x21\x10\x45\xd1\x7d\x47\xf1\x2c\x53\x22\x81\xc6\xbe\x0a\xc5\x47\x6a\x60\x33\xd9\xcf\x79\x2b\x31\x56\xf7\xc3\x36\x4b\x85\x8d\xbe\xb5\xa3\x09\xa1\xf3\x57\x46\x8c\x4c\x04\xa1\x3a\x75\x0a\xca\x75\xfa\x75\xbf\xcc\x24\x69\x78\x43\xcb\x3f\xcd\x7f\xd8\x13\x00\x00\xff\xff\xaa\x97\xff\x85\x4d\x00\x00\x00")

func templatesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMd,
		"templates/README.md",
	)
}

func templatesReadmeMd() (*asset, error) {
	bytes, err := templatesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1528494836, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x6b\xe4\x38\x13\xbe\xfb\x57\xd4\x1b\xc2\x8b\x0d\x1d\xe7\xb2\xec\x21\x4b\x0e\x61\x98\x81\x86\xcd\x10\x12\xd8\x4b\x93\x83\xc6\x5d\x76\xc4\xca\x92\x91\xe4\xcc\x34\x46\xff\x7d\x91\x2c\x7f\xc8\xed\xfe\x4a\x32\x3d\x30\xb9\xa4\xdb\xd2\x53\x1f\x4f\xa9\x1e\x95\xbb\x22\xd9\xbf\xa4\x40\x68\x1a\x48\x9f\x50\xa7\x9f\x04\xcf\x69\x51\x4b\xa2\xa9\xe0\xe9\x57\x52\x22\x18\x13\x45\xb4\xac\x84\xd4\x70\x51\x88\x94\x54\x42\xa2\x16\x29\x15\xd7\xc8\xb0\x44\xae\x09\xbb\x88\xa2\x57\x22\x21\x8e\x00\x00\xe8\x1a\xb9\xa6\x7a\x63\xc1\xea\x9e\x54\x70\x0b\x25\xa9\x56\x4a\x4b\xca\x8b\xe7\x1e\x93\x2e\xfd\x3e\x68\x1c\xcc\xfe\x35\xcd\x15\x48\xc2\x0b\x6c\x83\x79\xaa\x30\xa3\x39\xcd\x5c\x30\xca\x06\x32\x6c\x04\x9a\x83\x7a\x11\x35\x5b\x3f\x62\x41\x95\x46\x19\xec\x86\x14\x2e\xd3\x87\xfa\x1b\xa3\xd9\xbd\x58\x63\x88\xbd\x82\xcb\x21\x44\xb8\xb9\x85\xd4\xee\x61\xe9\xe7\xe1\xe1\xd5\x08\x70\xd1\x34\x7e\xc3\x23\x2a\x6d\x97\x8d\xb9\xb8\xb1\x31\x8c\xcd\x18\xd3\x25\xb4\x08\x5c\x21\x5f\x4f\xbd\x8f\x1e\x99\x28\xe0\x2c\x23\x1a\x0b\x21\xe9\x6f\x48\x9c\xa8\x65\x86\x3f\x85\x3c\xc2\x28\x51\x3f\x91\xb1\xab\x33\x52\xd6\x34\x3e\xaa\x4b\xba\x80\x4b\x97\xd9\x08\x74\xd7\x66\x0a\x21\xc7\xdd\xbe\x8f\x23\x76\xcf\x41\xe5\x6b\xfc\x31\xc3\xf5\xea\x79\xf5\xdc\x7e\x3c\x3b\xc7\x96\x81\x49\x7f\xf6\x54\xd0\x1c\x18\x72\x48\x97\x6d\xd8\x60\xcc\x6c\xa0\x61\xb0\x8e\xfa\x4c\x94\x95\xa8\xf9\xda\xb1\x3f\xc0\x03\x48\x67\xa9\x19\x83\x7b\xa0\x31\x2e\x32\xfb\x7f\x31\x90\x09\x66\x6f\x21\xcc\xa2\x69\x00\x99\xb2\x29\x70\xca\x16\x6f\xac\x55\x12\x45\xd7\xd7\xe0\x48\xf9\x07\xa5\xb2\x04\x4a\xd4\xb5\xe4\x0a\xf4\x0b\x42\x56\x4b\x89\x5c\xc3\xab\x5f\x13\xb9\x7b\x5c\x3a\x12\xa3\xbc\xe6\x59\x80\x8d\x13\xc8\x99\x20\xfa\xcf\x3f\xa0\xf1\x76\xfa\x0b\xe3\xee\x61\xb9\xe4\xb9\x48\x3b\x37\x36\xc3\x28\xd2\x9b\xca\x9b\xbb\x27\x9c\x14\x28\x41\x69\x59\x67\xba\x31\x91\x33\x1f\xe7\xc1\x6a\x02\xdd\x29\xfd\x22\x45\x69\x2b\x18\x73\x5b\xc6\x96\xde\x04\x66\x3b\xd9\xa5\xea\xa3\x99\xde\x39\x2b\x0b\x7f\x8e\x8e\xf1\xf6\xa9\x95\xdc\x4d\xec\xb5\x77\x73\xba\xd7\x40\xb5\x57\x9d\x9d\xe3\xdc\xbb\x96\x8e\xdb\x06\x3e\xda\xf1\xa0\x78\x2b\xf7\xf1\x48\x57\x7c\x13\x13\x3e\xe4\x17\xd3\x19\x4f\x49\xe7\x8a\xe6\x40\xe1\x16\xf2\x74\xab\x34\x84\x6f\x92\xbf\xe0\x7f\x34\x5d\xaa\xcf\x65\xa5\x37\x71\x32\x6a\xa5\x8e\x9a\x40\x34\xe6\x4c\xf5\xbc\x9f\x6c\xce\x3f\x0b\xcd\x79\x1e\xf9\x26\x39\xc0\x45\x4e\xc9\x37\x86\x71\x57\xbb\x59\x0a\xa6\xcf\x5a\x4c\xc7\x8c\xfa\x4e\x75\xf6\xd2\x57\xdf\x47\xdb\x2b\xf7\x1e\xad\x7b\xb3\xce\x65\x44\xb5\x33\xda\xd6\xe5\x31\x08\xfc\xcd\x94\xb4\xaf\xf8\x7d\x07\x24\x4e\xa2\x19\xd9\x98\x7c\x5d\x63\x4e\x6a\xa6\xb7\xcc\x72\xca\x7c\x35\x76\x11\xfd\x54\x11\xa9\xf0\x2d\x74\x6f\x23\x7f\x21\xe9\x1e\xc8\x85\xee\x48\x5c\xaa\x47\x21\xf4\x7b\x8b\xd2\x26\xf9\x9e\xd2\x7c\x58\xa5\xfc\xc5\xb6\xbf\x3c\xc1\x95\x19\xe8\x5f\x3f\x0c\xac\x3a\x03\xee\x95\xe1\x90\x1c\xb5\x95\xb5\x5d\xfb\xe4\xcc\x06\xaa\xb4\xbf\xf7\x26\xbd\xef\x4f\xd7\x44\x0a\x5a\x9d\x4b\x8e\x53\x82\x03\xc9\xcf\x87\xa3\x3e\xec\x58\xee\x3c\x5d\xe7\xd5\x8a\xff\xcf\x01\x1e\x58\x2d\x09\x03\x63\xfe\xa6\xca\x5e\xdd\x67\x3c\x98\xdb\x42\x70\x74\x9d\x66\xa0\xbf\x5d\xb5\x76\x4b\xc8\x2f\xac\x59\x40\xf9\x49\xdd\xad\xf6\xb6\xb7\x3a\xb9\xbf\x1f\x91\xb5\xe5\x7b\xa1\x95\x8a\xc7\x5e\x83\x95\xb6\x4e\x72\x3a\x5d\xc9\xb9\x3d\xd6\xd7\x2b\x91\x50\xfa\x79\xf6\x36\x70\x69\xe7\x5a\x3b\x70\xfb\xc5\xf1\xac\xed\xb6\x8d\x22\xb8\x1f\xc1\xba\x59\xbb\xfd\x16\x04\x3a\xde\x36\x4c\xdd\x9d\x77\x13\x39\x7f\x77\x8c\x79\x62\x28\xaa\xde\x2b\x61\x0c\xf0\x07\x55\xda\x2a\x36\xed\xd7\xbd\xb3\x00\x13\x5b\x75\x3f\x38\x6e\xce\x6d\x39\xff\x9b\xf4\xde\x16\x39\xed\x67\x03\xe3\xd9\x73\x53\xf4\x17\x21\xfb\xbc\xc7\x14\xda\xe2\xf9\x41\x1b\x72\x21\xdd\xf7\x82\xbe\xe2\x30\xf7\xf7\x8c\x4e\xed\x1c\xba\x4f\xc3\xdb\x74\x97\x2a\x1d\x41\xeb\x79\x65\x67\xfc\xbe\xfb\xae\x9f\x29\xf6\xbf\x02\x1f\x21\x52\xe1\x9b\x80\x95\x26\x13\xfd\x17\x00\x00\xff\xff\x11\x03\x2d\x1e\xcb\x14\x00\x00")

func templatesIdentities_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIdentities_registryGotpl,
		"templates/identities_registry.gotpl",
	)
}

func templatesIdentities_registryGotpl() (*asset, error) {
	bytes, err := templatesIdentities_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5323, mode: os.FileMode(420), modTime: time.Unix(1554955843, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x6b\x73\xdc\x36\x92\x9f\x3d\xbf\x02\x99\x52\xbc\x33\xde\x11\xe5\xcf\x93\xd5\x56\xf9\x1c\x27\xa5\x3a\x3b\x51\x45\xbe\xec\x07\x97\x6b\x0d\x91\xcd\x11\x62\x0e\x40\x03\xa0\x6c\xdd\xd4\xfc\xf7\x2b\x3c\x48\x00\x7c\x0c\x41\x3d\x7c\xba\x3a\xf9\x43\xa2\xc1\xa3\xd1\x2f\x34\xba\x1b\x60\x97\x38\xfd\x8c\x37\x80\x76\x3b\x94\x5c\x80\x4c\x5e\x33\x9a\x93\x4d\xc5\xb1\x24\x8c\x26\xbf\xe1\x2d\xa0\xfd\x7e\x36\x23\xdb\x92\x71\x89\x16\x33\x84\x10\x9a\xe7\x5b\x39\x37\x7f\x6d\x58\x82\x4b\xc6\x41\xb2\x84\xb0\x13\x28\x60\x0b\x54\xe2\xa2\xee\x25\xf2\xaa\xba\x4c\x52\xb6\x3d\xd9\x12\x99\x5e\x41\x51\x5c\x9d\xa4\xac\xbc\x11\x92\x57\xa9\xac\x38\x98\x81\xbb\xdd\x31\xe2\x98\x6e\x00\x25\x17\x25\xa4\xc9\xfb\x9b\x12\xce\x39\xbb\x26\x19\x70\xa1\xd6\xd7\xd0\x14\x8a\x68\xbf\x77\x53\x80\x66\xe8\xd8\xf6\xb6\x41\xfc\x89\x0b\x92\x69\x22\x22\x01\xed\xf7\xb3\xe5\x6c\xb6\xdb\xa1\xa3\x02\x4b\x10\xf2\x4f\xe0\x82\x30\x8a\xd6\xa7\x16\xe2\x5b\xdd\xfc\x4a\x4a\x4e\x2e\x2b\x09\xa2\x1e\xa0\xd8\xb3\xdb\xd9\xc5\x8f\xc8\x0a\x1d\x01\xad\xb6\x6a\xde\x65\x45\x8a\xec\x0d\xad\xb6\xc2\x80\x68\x83\xde\xef\x67\x27\x27\x8a\xf3\x7a\x86\xa6\x1a\xed\xf7\x88\x43\xc9\x41\x00\x95\x02\xc9\x2b\x40\x25\x13\x82\x5c\x16\x80\xae\x71\x51\x81\x40\x39\xe3\x08\xd7\x58\x68\x62\xcc\xf4\x06\x33\x2b\xb4\x79\x32\x93\x0a\x62\x07\xbe\x90\x9c\xd0\xcd\x6c\x96\x32\x2a\x6a\x91\x3a\xf6\x1d\x51\xbc\x85\x15\x3a\xd2\xab\x29\x2a\xcc\xe4\x3f\xcd\xe2\x96\x85\x16\x6d\x6a\x56\x6a\x63\x6c\xa6\xaa\x01\xe6\xaf\xfd\x3e\xb1\x8b\xb8\x29\x1d\xac\x4e\x0d\x29\xf5\x8c\x5a\x38\x4e\x36\xee\xef\x99\xc2\x96\xe4\x88\x32\x69\x65\xf3\x8e\x65\x50\x24\x3f\x83\xc4\xe9\x15\x64\x8e\xb1\x7e\xef\x1b\x2a\x89\xbc\xb1\xcc\x39\xcb\x40\xff\x6c\xa3\xde\xb4\xb3\x5c\xff\x66\x97\x7f\x41\x2a\x93\xd9\x35\xe6\x71\xf0\x4e\x51\xb3\x09\x92\xa6\x71\xa7\x89\x51\x43\xd7\xa8\xd1\x40\x0f\xd4\x1f\x20\xa4\xea\xdd\xef\xe7\x2b\x3d\xf4\x35\x96\xb0\x61\xfc\x66\xdd\x37\x94\x55\x3c\x6d\x84\x6c\xc6\x9f\x9b\x5d\xbc\xee\x82\xb6\x3d\x6e\x24\x27\xd7\x58\xaa\x91\xed\x81\xa6\x63\xbf\x5f\xcd\xf6\x13\x79\xbd\xdb\xf5\x8d\x38\x13\x7f\x30\x26\xc7\x64\x71\x5e\x54\x1c\x17\x68\xbf\x7f\x4b\x84\xf4\xa5\x81\x51\xa1\x5a\x58\x1e\x31\xb7\x51\xf4\x98\x35\x3e\x7c\x7c\x31\x38\x52\x11\x7c\x72\x82\x3c\xed\x90\x15\xa7\x46\x35\x48\xaf\x6a\x08\x44\xa8\xfe\xa9\xb0\x4d\x66\x79\x45\x53\xb4\x60\x91\xb8\x2c\x9b\x95\x16\xcb\x7e\xbd\xd1\x32\x33\x58\x0c\xc3\x74\xea\x37\x33\xf8\xbf\x66\xa5\xc3\x1d\xa3\x92\x11\x2a\x81\x23\xc9\x10\x46\xca\xfc\x6a\x84\xe3\x50\x9c\x4e\x92\x5a\xbc\x87\x9c\x9c\xe0\xcb\x02\x44\x4d\x93\x46\x63\x7d\x8a\x70\x59\x02\xcd\x16\x71\xc0\x77\xfb\x15\x62\x49\x92\x2c\x7d\xb6\x3c\x57\xa0\x2c\xe1\xaf\x34\x34\x0b\x54\x04\x62\x92\x4c\xff\xc4\x88\xc2\x57\xb3\xba\x95\xe3\x43\xf1\xc1\xe0\xb2\xa8\xd7\x4f\x92\xa4\x9f\x25\xa3\xac\x62\x95\xbc\x23\xa7\xd4\x91\xf1\xef\x95\x62\x85\x02\x64\xec\x7c\x8d\x97\xb1\x4d\xf5\x3a\xcd\x32\xac\x92\x7a\x42\xb2\x38\xb4\x5b\x96\x06\xfe\x3e\xd0\x53\x56\x49\x2b\x0e\xbd\xdf\x52\x46\xaf\x81\x4b\x5f\x1a\x5a\x13\xe9\x10\xdd\xb7\x63\xb7\xfa\xef\xb0\xda\x69\x4c\x42\x7e\x6e\xf1\x67\x58\x1c\x18\xbe\x42\x05\xd0\x05\x5b\x3a\x16\x12\x35\xed\xe5\x4f\x88\xa0\x7f\xd8\xbe\x9f\x10\xf9\xfb\xdf\x43\x16\x7e\x20\x1f\xd1\x29\x62\x1f\xc8\xc7\x83\xac\xf9\x19\x72\x5c\x15\xf2\x77\x9e\x01\x0f\xcc\x4c\x66\x3a\x10\x53\x3d\x84\x6e\x50\x4e\xa0\xc8\x44\xad\xad\x29\xa3\x12\xe8\x2d\xf8\xe3\x2f\xb8\x58\xa2\x0f\x1f\x8d\x1b\xd0\xb2\x31\x75\xb3\x23\xa9\x71\x6d\xcc\x2a\xbf\x5b\xb4\x9c\x1f\xd4\xf6\x6a\x6a\x8f\xac\x39\xe8\xc2\xa3\x2a\x3c\xd6\x0d\x8b\x0c\x4b\xde\xb3\x8b\x12\x73\x01\x01\x3b\x22\x8d\xba\x55\x32\xc8\x94\x6a\x19\x30\xb1\xfb\xfa\xe4\x04\xfd\xde\x35\xe5\xe8\x2b\x29\x0a\xc4\x68\x71\xa3\x59\x8e\x6d\xd7\x86\x5c\x03\xb5\x22\x49\xd0\x6f\xcc\xfc\x89\xb6\x80\xa9\x40\x4a\x81\x38\xd8\x26\x01\xb7\x10\x52\xcd\x82\x85\x15\x7a\x92\x24\x46\x1e\xb1\x46\x42\x2b\xf5\x14\xfa\xef\xac\xe5\x49\x0b\x67\x65\x74\x92\xc5\x8b\x11\x1c\xd0\x7e\x7f\xd8\x74\xd4\xda\xe4\xeb\xc2\xb5\x6d\xbb\xeb\x56\xb0\xb0\x17\x4b\x44\xa8\xec\x3b\x64\x41\x26\xaf\xce\xcf\xce\x68\xce\x12\xcf\x57\x37\x7e\xbe\x55\x5c\x2f\x6c\x18\xf3\x35\xdb\x3e\xe6\x56\x0d\x51\x44\xe0\xf6\xb4\xda\x0f\x1c\x77\x69\xac\x1f\x5f\xa5\x0a\xff\xd6\x16\x3d\xbc\x35\x77\x3b\x1d\x3d\xbc\x67\xbf\x68\x35\x3d\x52\xd4\xa2\x1c\x17\x02\x74\x50\x14\x10\xa6\xf8\xa2\x57\xae\x01\x28\x7e\x7d\xfa\x4b\x30\xba\x9e\x1f\xcf\xd1\x56\x6c\x54\xf8\xa8\xff\xbe\xd4\x8d\xff\xd6\xb4\x59\x39\xcd\x3f\x59\x59\xfe\x06\x5f\x47\xf8\x53\x7b\x2a\xea\x6c\x1e\x3e\x6f\x14\x4e\x5a\xd6\x23\x00\x17\xcb\xc3\x40\x5a\x22\x7f\x7e\x68\xac\xd3\x7a\x9f\x11\xeb\x03\x7a\xb2\x9a\x79\x76\xee\xd8\x8f\x0b\x15\xdf\xd5\xe6\x12\x8c\x7b\x71\x24\x5a\x8c\x88\x6d\x19\x98\x54\xeb\x8f\x8b\x2b\x56\x15\xd9\xbf\x38\x91\x70\x46\x89\x24\xb8\x20\xff\x0d\x5c\x89\x53\x07\x9a\x6a\x29\x63\x78\x5b\x2a\x70\x94\x9c\x57\x97\x05\x49\x15\x35\xa8\x05\xf6\x88\x50\xa2\xad\xc8\xd7\x1e\xb0\x20\x03\xe0\xed\xb9\x24\xb7\xd3\x83\xf6\xbe\xb6\x63\xdf\xf6\xc7\x35\xd9\xbd\x17\x13\xf4\xf5\xba\xee\x43\x51\x5d\x6d\x38\x0e\x6a\xcb\x3d\x39\xe9\x28\xf0\xd2\xfd\x5d\x36\x85\xb0\x9c\xb4\xdc\x05\x13\x2f\x07\x74\xfd\x4d\xa0\x8a\x92\x2f\x55\x1d\xb2\xa8\x39\x13\x69\x55\x53\x16\x4b\x14\xba\x08\x26\xcc\x33\x73\x3d\x6c\x6a\xe5\xac\x4d\x78\xd2\x2c\xe0\x06\x25\xaf\xeb\xf3\xb9\xde\xc7\x8d\x94\x95\xe1\x69\x81\x98\x77\x32\x33\xb7\x61\xd8\x05\x48\x0f\x4b\x01\xf2\x61\x18\x16\x2c\xb3\x20\x19\xaa\x0f\xec\x38\xae\xc5\xb1\x0b\x9d\x22\x92\xdd\x07\x53\xfa\x8e\xd5\x2b\xcc\xb3\x94\x65\x90\xb5\x0f\x58\x6d\xcc\xa3\x19\x71\xeb\x53\x75\x3a\x15\x83\xce\x73\x9d\x32\x18\x70\xa2\xa3\x68\xd1\xc4\x3c\x6e\x6f\x79\x22\xc7\x7c\x15\xac\x7b\x45\xca\x49\x29\x5d\x16\xf2\x67\x96\x86\x61\x08\x4b\x2b\x6d\xe9\xf4\x18\xe5\x17\xba\xfd\x12\xab\x11\x3f\xb3\xb4\x63\x44\x2c\xe7\x3e\x29\xa2\x44\xfa\x1f\x38\xfd\x2c\x49\xfa\x59\x1c\xc0\xee\x53\x90\x8f\x9a\x48\x7a\xec\x26\xd6\x38\x0e\x21\x9b\x6f\x65\x72\x51\x72\x42\x65\xbe\x98\xff\xe3\x47\xb1\xfe\x51\xfc\x73\xae\xe2\x6b\x77\x30\x68\xb1\xb9\x26\x63\x3f\x97\x1d\x51\x45\xfa\x69\x8d\xcc\xcc\x29\xff\x2b\x50\xe0\x58\xc2\xaf\x20\x25\x70\x94\x74\x0e\xf1\x93\x13\xf4\x2b\x48\x45\x62\xc7\x70\xb4\x63\xa9\xce\x00\xbb\xd7\x39\xa4\x40\xae\xdb\x76\xef\xe8\x00\xcf\x06\x56\x5c\x2c\xc3\x75\xea\xfc\x6e\xc8\x52\x63\xf6\x3a\xa7\x42\xcb\xc5\xee\xb2\xe0\xe2\x00\x0b\x2e\x06\x58\xd0\x98\xfe\x92\xb3\x12\xb8\x3a\xb2\x23\x18\x81\x2a\xa1\x34\xc1\x05\x7d\xfa\xe0\x88\x67\xcf\x00\x36\x0b\x6f\x8f\xf7\x32\xaa\x39\x38\xfa\x59\x84\x4e\x91\x07\xa1\xb5\x35\xda\x3b\x03\xd3\x0c\x2d\x86\xb6\xc7\xb2\xdb\x65\x52\xb6\x4b\xcb\xcf\xde\x78\x5c\x98\xa6\xfe\x93\x42\x4d\xba\xaa\xc7\x43\x56\x67\x7b\xee\x35\x94\x1e\xd9\xc9\x51\x11\xb4\x19\xe2\xc7\xd1\x35\xd3\x49\xae\x63\x5e\x33\x79\x89\x4e\x4f\xd1\x4b\x2f\xf0\x3d\x39\x41\x94\x15\x84\xca\x35\xda\x30\x73\x25\x26\x9a\xce\x3a\x9a\x18\x8f\x7c\x3d\x88\xa8\xe7\xea\xea\xa0\x5d\x08\x27\x76\xf5\x63\x8d\x9e\xf7\x2b\xce\xaa\xb3\x66\xcb\xe7\xde\xfb\xf1\xb8\x28\x55\x1c\x10\x45\xcb\xde\x4f\x31\xe6\x2e\xc1\x68\x05\xe0\x68\x15\x5f\x89\x4c\xaf\x50\x7e\x5f\xe4\xa7\x58\x40\x78\x66\xae\x83\x7e\x43\x47\x97\x19\x7a\x13\x3d\x5f\xf4\xb3\x69\x39\x89\x4d\x56\xe6\xa2\xb4\x91\xee\x39\x56\x14\xe2\xb2\x2c\x4c\x8e\x9d\x32\x8a\x28\x29\x9c\xd3\x89\x51\x44\x66\xa4\xce\x55\x4f\x0c\x53\xf4\xe2\x0b\xbb\x41\x0f\x29\xbb\xe7\x9b\x1e\x9b\x68\x4e\xb8\x6b\xce\x51\xf6\xd7\xd1\x9e\x99\x67\x1b\x49\x8e\x7e\x30\x2b\xfb\x87\xe2\x99\x78\xf3\xa5\xc2\xc5\xc2\x3f\x29\x97\x9e\xf8\x4b\x4c\x49\xba\x98\xa7\x98\x2a\x4b\x54\x6a\xe6\xe5\x9c\x6d\x11\x46\x86\x8a\xaf\x44\x5e\xa1\x8c\xe4\x39\x70\xa0\xb2\xb9\x82\x99\x07\xb9\x23\xc1\x74\x50\x6d\x56\x8f\xcb\x3c\xb5\xef\x3c\xdb\xb4\x88\x01\xe3\xfb\xc3\xa9\x16\xa7\x97\x0a\x1b\x32\xd2\x2f\x06\x40\xcc\x1a\x1d\x0a\x75\x6b\xd0\xaf\xd7\xee\x2e\x94\xad\x2b\x9d\x0c\xa0\x34\xb7\x18\x64\xe4\x16\x43\xdf\xbe\xc6\x3a\x6c\x76\xa1\xd8\x04\x0a\xc9\x11\x53\x36\xd2\xf0\xe4\xd9\x33\xbb\x1d\x28\x29\x66\xcf\xf6\xb3\xd9\x33\x9b\x94\x3c\x9c\x60\xd9\xcf\x9e\xb1\xa4\x5e\xf9\x8c\x4a\xb6\x60\x95\x5c\xce\x66\xcf\x7a\x72\xe6\x6e\x90\x22\x9e\x80\x08\x8f\x6c\x42\xed\xce\x31\x67\xcb\x41\x1a\x26\x33\xa5\x46\x6d\x6c\xfc\x6e\x36\x7b\x26\x31\xdf\x80\x5c\x21\xe0\x3a\xe1\x13\x3c\x79\x48\x34\x87\xd9\x72\xf6\x8c\xe4\x7a\xc0\x0f\x8e\x81\x66\x43\x04\xfe\xe6\x7f\x51\x7d\x3e\x49\xa6\x65\xae\x45\x7e\x68\xfd\x35\xfa\x51\xcc\xf5\xc2\xcb\xa5\x11\xc2\x0b\x73\xaf\xf3\xc2\xe0\x74\xe8\x3e\x47\x6f\x0d\x9b\x7e\x35\xcf\x27\x74\xb8\x4c\x32\xcb\xe7\xb4\xe2\x66\x1b\xd2\x9c\xf1\xad\x89\x0c\x84\x64\x1c\x32\xc7\x79\x47\x66\x74\xdc\x68\x97\x5a\x2c\x15\xd6\x8c\xd7\xca\xa5\x7f\x68\xc3\xe4\x6c\xd9\x1b\xdd\xb6\xab\xb3\x05\x5f\x2a\xc2\x21\x7b\x73\x68\x60\x1d\x83\x4f\x3a\x61\xbc\x0b\xed\xf7\x1c\x53\x41\x14\xd5\x41\x5f\xf2\xe6\x5b\xc9\x04\xb8\x34\xa9\x6d\xfe\xc3\xe2\x14\x8e\x86\x2f\xc8\xbc\x73\x98\x1b\x5f\x64\xee\x99\x1a\xab\x22\x0e\xf5\x9a\x1f\x35\x28\x1b\x9b\x84\xc1\xe1\x80\xe1\x59\xfe\x14\x6a\x94\xf3\x4d\x02\x56\x9d\xb6\x1a\x12\x7b\x57\xa9\xb4\xc6\x19\x28\x2f\xf0\xf4\x68\x61\x1c\x2d\x1c\x3d\x40\xab\xed\x7c\x79\x77\x72\x0c\x5f\x86\xce\xe4\xef\x40\x96\x23\x49\x92\x2d\x4c\x12\xd0\x7b\xb2\x85\xc7\x2a\x9e\x6f\x12\x38\xc5\xc5\x7c\xe9\xb7\x16\x44\xc8\xf9\x72\x02\x85\x6f\x2c\x98\x47\x43\xa5\xa3\x85\x50\x09\x1b\xe0\x93\x04\x76\x46\xe5\x23\xa4\x24\x2f\x18\x96\x93\xe8\xf8\x45\xcd\x78\x1c\x94\x1c\x22\x8c\x43\xde\x43\x56\x3f\x9e\x89\x3b\x0d\x06\x30\x86\x1a\x53\xb8\xe3\xde\xe0\x90\xbf\xd5\x1b\xa1\xd5\xf8\x0e\x97\xce\xa2\xd9\xb0\x46\x54\x97\xde\xcb\x89\x7e\x9f\xcf\xa1\xe8\x88\x14\xd5\xe5\x38\x45\x31\x54\xf9\x91\xc7\x20\x7d\xc9\xab\xa2\x60\x5f\x21\x7b\x7d\xc5\x48\x0a\x22\x46\x95\x8c\x35\x3e\xa3\xfa\xf5\xc4\x24\x9b\xbc\x72\xd9\x4f\x35\xef\x2f\x46\x68\x07\x81\x4f\xf3\x15\x9a\x7f\x52\xd0\xf6\x2b\xed\xb5\xbc\xaa\x24\xdb\xd8\x54\x4e\x76\x40\x2d\xef\x20\x64\x87\x03\xe6\x51\x2c\x38\xc7\x52\x59\xb7\xb8\x7d\xb4\xd2\x99\xcb\xf6\x1a\x9f\x7a\x9a\xdf\x81\x10\x78\x03\xa6\x57\x75\x7a\xae\xc1\x03\x90\xbd\x91\x28\x79\x87\xbf\xbd\x05\xba\x91\x57\xe8\x65\x0c\xe1\xef\xf0\x37\xb2\xad\xb6\x66\x4a\x2c\xf9\xaa\xd5\xad\xa3\x5a\xf4\xdd\xf1\x43\x51\x44\xe8\x24\x8a\x08\xbd\x25\x45\xcd\x3a\x0f\x4e\x11\xfe\xa6\x9f\xd2\xa2\x97\xc9\xcb\x21\x27\x31\xfe\x24\xb0\x22\x9c\x70\x10\x34\x12\xfc\xd3\x3e\xb4\xbd\x37\x72\x0b\x01\x0d\xbd\x11\x38\x47\x1f\xc2\x2b\x15\x5c\x2c\x5a\x58\x2f\xef\x59\x4a\x63\x5a\x78\x9f\x32\x33\x4a\x3a\x5d\x66\x35\x16\xf7\x2f\xb3\x48\x94\x6f\x23\x32\x87\xf4\xf7\x11\x59\xf3\xc8\x22\x41\x9d\x2f\x03\xf4\x23\x0c\xfd\x32\xff\xd0\xe7\x01\x8e\x11\x0a\xdc\xb6\x26\x56\x53\xee\xbd\xab\x70\xe4\x9b\xc6\x58\x8f\x2b\x96\xcc\xe3\x01\x17\x2b\x82\x09\x01\xbd\xd7\x0d\xa5\x1a\xb1\x26\xcf\x67\x62\x71\xc7\x07\xff\x6d\xff\xeb\x4a\x48\xb6\xad\xd3\xf7\x0e\x42\xe2\xb2\x86\x5b\x5c\x96\x84\x6e\xf4\x07\x02\xfa\x6a\xd7\x41\x7a\x67\xba\x12\xfb\x7f\x34\x77\x9f\x85\x74\xd0\x69\xe5\x14\x6b\xa8\xfd\xa2\xb0\x70\x6b\x81\xb0\x7b\x63\x71\x1f\xc7\xed\x4d\x40\xe8\x0f\x2f\xd1\x3f\xbd\x0b\x01\x9b\xa0\x0a\x87\xf8\x99\x49\x0b\x03\x7a\xe6\x36\xb3\xa1\x33\xcb\xcb\xa2\xc5\x3f\x76\x28\x21\x25\x39\x49\x35\x67\x7f\x61\xbc\x49\x71\x04\x97\x37\x4d\x6b\x30\xbc\xb9\xdd\x35\x59\x33\xf7\x6d\x89\xfe\x56\xe3\x33\xdc\xd4\xa9\x9c\xb1\x4b\xd4\x21\x1c\x16\x1a\x50\xf7\x1a\x66\x00\x1d\x97\x5c\xbc\x5e\x21\xf6\xd9\x8a\x7f\x70\x61\x97\xcd\x79\x87\xcb\x0f\x6a\xa9\x8f\x3f\xa9\x69\x1d\x4e\x5f\xfb\x4c\x3e\x39\x41\xff\x02\x94\xb2\xaa\xc8\x34\x6f\x73\x42\x33\x44\xe4\x0a\x09\x86\x0a\x90\x7f\x13\x28\xbd\x82\xf4\x33\x62\xf6\x49\x28\xfb\x0a\xdc\xdc\x38\x10\x9a\xc1\x37\xc8\x90\x28\x21\x45\x5b\x5c\xfa\x32\x3b\x84\xe7\x5b\x05\xe2\x35\x16\xd0\x83\x70\xfd\x7c\xbd\x97\x21\x22\x90\x61\x5e\x15\x85\x27\x23\x11\x8e\xdc\xe2\x32\x52\x5a\x03\x6b\x2d\x96\x0a\xc6\x07\x23\xac\x8f\xb1\xb2\x8a\x20\x3f\xa0\xda\x65\x19\x2b\x18\xd4\x56\x73\x69\x32\xa0\x9c\xf6\xb2\x91\x08\x44\x04\xc2\xe8\x1a\xf8\x0d\xc2\xd9\x35\xa6\x29\x64\x48\x31\x40\xa3\x27\xaf\xb0\x44\x37\xac\xb2\xb7\xc8\x5a\xd2\x14\x20\x43\x97\x95\x44\x84\x22\xc1\xb6\xa0\x00\xe9\xe9\x35\x2b\x51\x25\x40\x8b\x3a\xee\xbd\x88\xcd\x61\x86\x84\x84\x2a\xaf\x3f\xbf\xc8\x71\x0a\xbb\x26\x75\x6e\x2f\xc3\xf4\xb0\x5d\xcb\x6e\x47\x66\x29\x0f\xdd\x7f\x1d\xbe\x66\xef\x31\x7f\xbd\xb6\x67\xe4\x61\x6b\x20\xd2\xce\x2b\x57\x5c\xea\x0b\xaf\x46\xb4\x4a\x90\x87\x13\xf2\x63\x1f\x59\x85\xeb\x9d\x4e\x51\xd4\x5d\xfb\x6c\x9c\x94\x09\xf6\x1e\x5b\x36\x33\x14\x0a\xad\x67\x08\xc7\xfe\x07\x86\x6d\xa6\xcf\xd7\xa3\x96\xcf\x7f\x46\x74\x3c\x18\xcb\xaa\x7f\x7e\xfb\xba\x1b\x7b\xaa\x98\x33\x80\xd5\xba\xc3\x0c\xc3\xf4\xf5\x40\xfa\xe0\x78\xbf\x9f\x14\xe2\x3b\x87\xb2\x99\xb6\x6f\x7c\x93\x55\x97\xb6\x56\x2e\xc0\x61\xe7\x77\xac\x7b\xf3\x06\x07\xa9\xab\x17\xe8\x55\x7b\xf5\x2f\xe8\x58\x0f\x88\x2b\x6e\x09\x0e\x5a\x72\xbf\xd3\xe2\x26\x58\xc1\x6b\x37\x14\xb4\x46\x46\x41\x0f\x9e\xb9\x79\xfd\x7e\xbb\x81\xde\x1a\x39\x05\x7a\xed\x9e\xb7\xa1\xeb\x76\x0d\x5d\xbf\x06\x0e\x46\x07\xcf\x81\x93\xf8\x05\x4b\x0e\x69\x5b\xda\xae\xb5\x26\xc5\x1b\x15\x09\x37\x78\xcc\xe6\x00\x37\xcd\xeb\x9e\xf7\x66\xad\x47\x66\x51\x2b\x79\x97\x3e\x75\x97\x6d\x32\xb8\xbb\xfe\x28\x70\xbf\x90\x42\x02\xd7\xd7\x7b\x5e\xaf\x6b\x35\x40\x83\x51\x71\x70\x19\x07\xb2\xa1\xff\x09\x81\x4e\xba\x56\x0b\xd7\x1f\x15\x05\xd7\xbe\x3f\xf3\x7a\x4c\x8b\x81\xd7\xf4\x46\xc1\xf2\x9e\xc2\x7a\xbd\xae\xd5\xc0\x0c\x46\x45\xc1\xf5\x73\x53\x4d\x67\xd3\xb8\xee\xe6\xaf\x22\x81\x76\xb6\x49\xdd\xb6\xee\x24\x54\xa2\x20\x7a\x09\x27\x07\xb2\x6e\x5c\x77\x93\x52\x91\x40\xbb\x68\xda\xb6\x75\x27\x87\x10\x03\xb1\x6d\x39\x3d\x83\x39\xc9\x4e\x6a\xc3\xd4\x56\xf4\xa6\xd1\xe0\xe6\x8f\x89\x02\x7a\xce\xc9\x16\xf3\x9b\x96\x9a\xbb\x56\x03\x36\x18\x15\x05\xf7\x0f\xc0\x59\xdb\xa0\xd7\x6d\x6b\x9b\xca\x6d\x46\x44\x42\x0c\x6f\x85\x0d\x44\xd3\xb6\x6e\x27\x87\xa3\x20\x5e\x40\xca\x21\xf8\xb0\xc2\xb4\xd4\x1f\xa6\xd8\xde\x48\x58\xed\x6d\x7d\xe1\x6d\xeb\x8b\x49\xdb\xfa\xc2\x3c\x07\xf0\x61\xe9\x16\x0b\xab\xee\x8d\x83\x55\x5d\xda\x37\x91\x0e\x98\x69\xaa\x3f\x8b\x6f\x06\xc4\x69\x61\xe7\x2e\x5f\xfd\x6b\x1a\x0d\x8a\xfe\x98\x38\xa0\x2d\x14\x3d\xfc\x46\x91\xb3\x2b\x0c\x67\x22\xc6\x7d\xf0\xfe\xa8\xf2\x3b\x38\xe3\x03\x0b\x3f\x6e\xaf\xdc\x84\xf1\xc9\x93\x4b\xfe\xe4\x92\x3f\xb9\xe4\x4f\x2e\xf9\x93\x4b\x8e\x9e\x5c\xf2\x27\x97\xfc\xc9\x25\xaf\x21\x3e\xb9\xe4\x4f\x2e\xf9\xa8\x4b\xbe\xeb\x7e\x82\x78\x87\x0f\xad\xcc\xed\x5d\x7c\xc1\x8e\xfe\x1a\x59\xb1\x10\x4c\x59\x89\x49\xeb\x7d\xf8\x38\xf6\x49\xc5\xbd\x55\xcd\x9a\x82\xd7\xff\x6a\xed\xac\x69\x15\x66\x6e\x47\xde\xad\xea\x68\x4d\x59\xe2\x41\xaa\x69\x7d\x0f\xce\x3c\x50\x65\xad\xdb\xf3\xee\x6e\xf5\xb5\x46\x77\xd7\x77\xa8\xb2\x35\x4d\x00\xff\x5f\x6b\x6d\x4d\xe3\xd2\xe3\xae\x21\x60\xbe\xf0\x3d\x2f\x30\x09\xab\x41\x4c\x3a\x1b\x82\xb2\x5b\x0f\xbb\xe9\x2d\xae\x8f\x4b\xed\x92\x06\xab\x83\x0a\x78\x9b\x62\x56\xd3\xb8\x73\xfb\x92\x56\xe3\xee\x47\x4f\x09\xab\xee\x57\xe0\x87\x6a\x59\x25\x51\x5e\x47\x50\xd2\xaa\x67\x1f\x8c\xeb\x7f\x6f\x69\x2b\xc9\x2b\x08\x1f\x03\xde\x7f\x75\xab\x28\x06\xfa\x35\xae\x22\x78\x91\x34\xa5\xae\xc6\xc7\x2e\x96\x51\x9f\x15\xef\x82\x83\x7e\x7c\xc2\xae\x56\x90\xa8\xa2\x4a\x56\x27\x3a\x1f\x2d\x47\x7c\x8f\xfb\x10\x15\x96\x62\x6b\x26\x05\x68\x8f\x54\x02\x8a\x27\xe5\x36\x05\x94\x48\x1e\x5d\x0d\xa8\xfb\x19\x8f\x5f\x38\x69\xef\xb3\xeb\xc5\x03\x15\x64\x8a\x2d\xb1\x74\xdf\xfc\xfd\x5e\xf5\x96\x9e\xf7\x15\x5c\xba\x73\x09\xa5\x18\x0a\xef\x64\xcb\xfb\x8e\xf4\x52\xb7\xb4\x30\xb3\x72\xb9\x0d\x82\x7d\x47\xb1\x6e\xe9\xab\xa8\x61\x8f\xe2\xd1\x92\x7d\x0d\xb3\x27\xdd\x04\x35\x7b\xa6\x5d\xe1\xc1\x3e\x0a\x0e\x0f\xee\xc1\x8f\xf5\xc7\xbe\xd5\xef\x9c\x1b\xe1\x11\xff\x7f\xac\xb0\x4f\x23\xe3\x07\x2b\xef\x33\xc4\xd0\xc7\x59\xdf\x07\x67\x19\x07\xd1\xb8\xfd\xfd\xe5\x7e\xa2\x98\xf6\x70\x45\x7f\x9e\xef\xc6\xab\xfe\x44\x96\x8a\x88\xf6\x3d\x62\xed\x81\x5f\x36\x22\xca\x0d\x69\xb6\x6e\x4c\xf1\x88\x28\x3f\xe5\xc1\x4a\x48\x3c\x18\xb3\x5c\x39\x89\x98\x59\x0f\x5f\x54\x62\x1c\x8b\x88\xd2\x12\x31\xb5\x57\x7c\x95\xfd\x9f\x00\x00\x00\xff\xff\xdb\x53\xae\x35\x7e\x64\x00\x00")

func templatesModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotpl,
		"templates/model.gotpl",
	)
}

func templatesModelGotpl() (*asset, error) {
	bytes, err := templatesModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotpl", size: 25726, mode: os.FileMode(420), modTime: time.Unix(1555102917, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x4d\x6b\xf3\x46\x10\xc7\xef\xfa\x14\x83\x11\x25\x29\xae\xda\xe4\xb9\x19\x72\x28\x71\x79\xc8\x21\x21\xb8\x2f\x17\xe3\xc3\xd6\x1a\xdb\xdb\xae\x57\xca\x7a\xe5\x60\x84\xbe\x7b\x59\xeb\xc5\x92\xbc\x2b\x29\x26\x86\x06\x66\x02\x01\x6b\x67\xe6\x3f\xb3\x9a\x5d\xff\x70\xcc\x96\xff\xb2\x35\x42\x9a\x42\xf0\x3b\xea\xe0\x31\x92\x2b\xbe\x4e\x14\xd3\x3c\x92\xc1\x0b\xdb\x22\x64\x99\xe7\xf1\x6d\x1c\x29\x0d\xa3\x75\x14\xb0\x38\x52\xa8\xa3\x80\x47\x3f\xa3\xc0\x2d\x4a\xcd\xc4\xc8\xf3\xf6\x4c\x81\x42\x71\x8c\xdb\x6d\x78\xbc\x9b\xe1\x9a\xef\xb4\x3a\x40\xe5\x15\xcc\x6c\xeb\x9e\xb7\x4a\xe4\x12\xb8\xe4\xfa\xe6\x16\x52\xcf\x03\x00\x47\xa6\x87\xbe\x5c\x69\x96\x87\xa7\x29\x28\x26\xd7\x08\x3e\x4a\xcd\xf5\xc1\xf4\x31\x06\xbf\xcc\x0a\x93\x87\xbc\xdb\x46\x12\xd3\x68\x1e\xfc\x13\xf0\x15\xec\x36\x51\x22\xc2\x3c\x33\xaa\xba\x27\xf8\x26\xb8\x9e\x1b\xfc\xe0\x35\xf9\x5b\xf0\xe5\x73\x14\x62\x91\xc6\xda\xc2\x3c\x4d\x1b\x71\x59\xf6\x14\xe6\x1f\x17\xf0\x00\x3f\xd8\xdb\x4b\x8f\xf9\x6a\xa5\xad\x35\xdc\x08\x94\xa7\x86\x82\x47\x85\x4c\xe3\x2d\xfc\x52\x36\x61\x2c\x7f\x38\x81\x2d\x8b\xe7\x3b\xad\xb8\x5c\x2f\x7e\xb4\x2b\x3c\xc9\x55\x04\x27\x99\x52\xaa\xd8\xc3\x98\x29\x94\x7a\x0c\x3e\x5b\x96\xbb\xd7\x56\xae\xcb\x3a\xf7\xf0\x49\x4a\xfb\x46\xe6\x02\x8d\x4d\x6c\x27\x1c\x99\x8d\x2b\xfc\xb2\x6c\x04\x13\xd7\x66\x99\x56\x9a\x9d\xd4\xea\x29\x3a\x08\xa6\x18\x2b\x5c\x32\x8d\x61\x5b\xc7\xd8\x69\x75\x02\x5a\x25\x38\xb6\xa6\x43\x69\x0d\x6e\x29\xbd\x32\xc5\xb6\xa8\x51\x4d\x71\x65\x46\xdc\xec\x9f\x3b\xaa\x7a\xb1\xee\xe8\x60\x86\x6f\x09\x57\x18\xb6\x5e\x76\x69\xe5\x72\x15\xba\x9b\xd4\x4e\xcd\x0b\xbe\x9f\x16\x0a\x57\xb3\x74\x73\x96\xc7\xd8\x7c\x61\xfe\xf2\xd1\x39\xdf\xd3\x7a\xed\xc5\xa4\xf0\x31\xf8\xe2\xee\x38\x20\x03\x3a\xb0\x95\xef\xda\x10\x71\xe7\xe8\xb7\x5e\x6b\x5f\xa5\xb6\x6a\xef\x8f\xd5\x8a\xbb\xae\xdc\xd6\x82\xee\x7b\x0a\xca\x8b\x1a\x52\x92\xad\xac\x6f\x79\x59\xf7\x7d\x0a\x50\x9e\x0d\xf1\xcd\x9c\x8b\xf3\x51\xb5\x09\xa1\x0c\x7b\xd2\x66\xdd\x89\x86\x25\xe9\xf7\xea\x90\xe9\x0f\xee\xf6\xb0\xa4\xbe\xbd\xe4\x20\x0f\x39\x92\xbf\x49\xad\x38\xee\x1c\x03\x51\x3f\x89\xf3\xc5\xe9\x2c\x5a\x32\xd9\x2f\xae\xda\x5c\xc4\x7d\x67\xab\x28\xc5\x39\x36\x1f\x94\x2f\xcd\x7c\x59\x4d\x8a\x4b\xb8\xa4\x82\x8e\x61\xfb\xe3\x10\x9f\xdc\xcd\x87\x6e\xf7\xf2\xce\x8c\x83\x29\xae\x58\x22\xf4\x5f\x4c\x24\x67\x5f\x01\x75\xab\xfb\x55\x42\xad\xe0\x1e\xc1\xfe\xe1\xe2\x2b\xc0\xb7\xaa\x83\x11\xca\x64\x3b\xea\x2a\xea\x57\x21\xa2\x77\x0c\x1f\x37\x11\x5f\xe2\xf1\x65\x7f\xf0\x42\xfa\x67\x0c\xfe\xfe\xf8\x86\xe3\xa0\x99\xac\xef\x1a\x38\xee\xc0\xbe\xff\x06\xe8\x98\xf7\xd2\xfa\xcf\xe4\x80\x7b\xdb\x8f\x83\xe7\x44\x68\x1e\x8b\xce\xd7\x58\xfa\xb8\xbe\x67\x07\x0a\x5b\x4a\xee\x88\xf8\x98\xb7\x63\xa9\x95\xc4\xe1\x65\x79\x5c\x0b\xb4\xac\x3a\x01\xef\xcf\x38\x3c\x07\xbc\xfc\xe1\x95\x01\x2f\x17\x21\xc0\x73\x28\x11\xe0\x11\xe0\x11\xe0\x5d\x20\x43\x80\x57\x55\xf1\x89\x80\x47\x7c\x07\xc4\x77\xc4\x77\xc3\xbd\xaf\xc8\x77\xaf\x4c\x2f\x37\x44\x67\x44\x67\x44\x67\xce\x6a\x89\xce\x6c\x46\x74\x46\x74\x56\x33\xa2\x33\xa2\x33\xa2\x33\x77\x92\xab\xff\xfa\x36\x45\x81\x67\xbf\xbe\xe5\x0f\xaf\xcc\x77\xb9\x08\xf1\x9d\x43\x89\xf8\x8e\xf8\x8e\xf8\xee\x02\x19\xe2\xbb\xaa\x0a\xe2\xbb\xb6\x11\xdf\x35\x8d\xf8\x6e\x78\xc4\x97\xe4\xbb\xef\xa8\x5b\xb7\xcb\x0c\xcd\x41\xdf\x5f\x1b\xef\xbe\xa3\x26\xb6\x73\x28\x11\xdb\x11\xdb\x11\xdb\x5d\x20\x43\x6c\x57\x55\x41\x6c\xd7\x36\x62\xbb\xa6\x11\xdb\x0d\x8f\xf8\xaa\x6c\xf7\xcc\xe4\xc1\xc1\x77\x66\xe9\xfa\x8c\x67\x54\x88\xf3\x1c\x4a\xc4\x79\xc4\x79\xc4\x79\x17\xc8\x10\xe7\x55\x55\x10\xe7\xb5\x8d\x38\xaf\x69\xc4\x79\xc3\x23\xfe\x37\x9c\x67\x48\x85\xe8\x0c\x88\xce\x88\xce\x9c\xd5\x12\x9d\xd9\x8c\xe8\x8c\xe8\xac\x66\x44\x67\x44\x67\x44\x67\xee\x24\x9f\xf1\x2b\x5c\xfe\xbf\xf5\x30\x4d\xcb\x4f\x99\xf7\x5f\x00\x00\x00\xff\xff\xb5\x07\x23\xe4\x6c\x46\x00\x00")

func templatesRelationships_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRelationships_registryGotpl,
		"templates/relationships_registry.gotpl",
	)
}

func templatesRelationships_registryGotpl() (*asset, error) {
	bytes, err := templatesRelationships_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 18028, mode: os.FileMode(420), modTime: time.Unix(1554956824, 0)}
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
	"templates/README.md": templatesReadmeMd,
	"templates/identities_registry.gotpl": templatesIdentities_registryGotpl,
	"templates/model.gotpl": templatesModelGotpl,
	"templates/relationships_registry.gotpl": templatesRelationships_registryGotpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"README.md": &bintree{templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": &bintree{templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": &bintree{templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": &bintree{templatesRelationships_registryGotpl, map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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

