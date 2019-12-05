// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
		},
		"/templates/default.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "default.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
			uncompressedSize: 1082,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\x08\xfb\x12\x0b\x98\x7a\x2f\xb0\x0d\xc5\xb0\xf5\x12\x0c\x43\x82\xec\xb2\x06\x81\x1a\x33\xae\x5a\x89\xca\x64\xba\x0d\xe0\xe8\xdf\x07\xd9\x86\x63\x2d\xd9\x6d\xbd\x49\x14\xf9\x1e\x1f\x1f\xd5\xb6\x50\xe2\x5e\x13\x42\xb6\xdd\xd6\xcd\xe3\x33\xee\x38\x83\x10\x7e\xb5\x2d\xc8\x15\x2b\x6e\x6a\x38\x01\xbb\xf5\xe1\x80\x1e\x42\x68\x5b\xd0\x7b\xc0\xdf\xe3\x63\xb6\xd7\x5e\x53\x15\x6b\x6e\x63\xcd\x9d\x41\xcf\xb5\xfc\xd6\x45\xe1\x04\x06\xa9\x2f\x43\x2a\x21\x84\x0d\xc4\xa4\x7b\xef\x9a\xc3\x42\x3d\xa2\xa9\xe5\xca\x79\xc6\xf2\x87\xd2\xbe\x96\x3f\x95\x69\x30\x12\x3e\x3b\x4d\x90\x41\x44\x85\x9e\xb2\x62\x98\x47\x2c\xf9\xc5\x59\xeb\xa8\x2f\x2e\x86\xd8\x04\xaf\x80\x10\xe6\x6d\x0b\x6f\x9a\x9f\xd2\x64\xb9\x44\xeb\x5e\x31\x65\xff\xae\x2c\xd6\x7d\x83\x57\xd9\xc7\xc6\x8b\xf1\x34\x1e\x66\xc9\xf0\x54\x14\x6e\x15\xa9\x0a\xfd\x7a\xb9\x18\x8a\xe5\xd7\x23\xa3\x27\x65\xd6\xcb\x05\x84\x70\x93\xdf\x74\x79\xf5\x67\x8f\x3b\xd4\xaf\xe8\x3f\xc6\xa4\xe5\x70\x49\xd0\x53\x78\xc6\x23\xf7\x1c\x5b\xa3\x6b\x1e\xe0\xbd\xa2\x0a\x41\xc6\x74\x21\x7a\x49\x42\xcc\xce\x0f\x97\x33\x86\x10\x3e\xc1\x87\xce\x85\xa8\x3d\xda\x06\xa3\x78\x38\x81\x55\xfe\xa5\x74\x6f\x04\x27\x78\x62\x6b\x06\x99\x43\x4b\x42\xdc\x11\x39\x56\xac\x1d\xa5\x44\x93\xf8\x7f\x64\x5b\xb9\xc6\xef\xf0\x56\x08\xe8\xf6\xf1\x1e\x09\xbd\x62\xe7\xfb\x61\x6e\xe6\x57\x82\xc5\x6c\x76\xc5\xa9\xe9\x2c\x4b\x4d\x95\x34\x9a\x5e\x24\x6b\x36\x38\x4c\x92\xd1\x1e\x8c\xe2\xf4\x1f\xc8\x7f\xd9\x7d\xc6\xd8\x39\x62\xa4\xce\x8f\x3c\xcf\x73\x78\x78\xaf\x9f\xf3\xb0\x01\x21\x22\xb8\xa6\x12\x8f\xc9\x16\x43\xd6\x2d\x06\x29\xdb\xa9\xe9\xe6\x32\xd5\x73\xb1\x9a\x51\x57\xd1\xfb\x37\xcd\xbb\xd8\xb1\xbf\x7a\x4a\xdc\xf9\x13\x00\x00\xff\xff\xaa\x4e\x32\x4b\x3a\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/default.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
