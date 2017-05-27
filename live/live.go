// Package live implements a live implementation of go-resources http.FileSystem compatible FileSystem.
package live

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
)

// Resources describes an instance of the go-resources which is an extension of
// http.FileSystem
type Resources interface {
	http.FileSystem
	String(string) (string, bool)
}

// Dir returns an Resources implementation that servers the files from the
// provided dir location, it will expand the path relative to the caller.
func Dir(dir string) Resources {

	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("failed to get Caller details")
	}

	dir = filepath.Join(filepath.Dir(filename), dir)
	return &resources{http.Dir(dir)}
}

type resources struct {
	http.FileSystem
}

func (r *resources) String(name string) (string, bool) {

	file, err := r.Open(name)
	if err != nil {
		return "", false
	}

	content, err := ioutil.ReadAll(file)

	if err != nil {
		return "", false
	}

	return string(content), true
}
