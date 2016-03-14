package resources

import (
	"bytes"
	"io"
	"os"
	"time"
)

func makefile(in File) (File, error) {
	var buf *bytes.Buffer
	_, err := io.Copy(buf, in)

	if err != nil {
		return nil, err
	}

	stat, err := in.Stat()

	fi := fileInfo{
		name   : stat.Name(),
		size   : stat.Size(),
		mode   : stat.Mode(),
		modTime: stat.ModTime(),
		isDir  : stat.IsDir(),
		sys    : stat.Sys(),

		files: nil,
	}
	return file{buf, fi}, nil
}

type file struct {
	*bytes.Buffer
	fi   fileInfo
}

func (f file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, os.ErrNotExist
}

func (f file) Stat() (os.FileInfo, error) {
	return &f.fi, nil
}

type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}

	files []os.FileInfo
}

func (f *fileInfo) Name() string {
	return f.name
}
func (f *fileInfo) Size() int64 {
	return f.size
}

func (f *fileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *fileInfo) ModTime() time.Time {
	return f.modTime
}

func (f *fileInfo) IsDir() bool {
	return f.isDir
}

func (f *fileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return f.files, nil
}

func (f *fileInfo) Sys() interface{} {
	return f.sys
}
