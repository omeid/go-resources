// Unfancy resources embedding with Go.

package resources

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

//Create a new Package.
func New() *Package {
	return &Package{
		Config: Config{
			Pkg:     "resources",
			Var:     "FS",
			Declare: true,
		},
		Files: make(map[string]*Entry),
	}
}

//Configuration defines some details about the output Go file.
type Config struct {
	Pkg     string // Package name
	Var     string // Variable name to assign the file system to.
	Tag     string // Build tag, leave empty for no tag.
	Declare bool   // Dictates whatever there should be a defintion Variable
}

type Entry struct {
	RenderedFileInfo string
	RenderedData     string
	FileInfos        []string
}

func NewEntry(f *os.File) *Entry {
	entry := &Entry{}

	buf := &bytes.Buffer{}
	fileInfoTpl.Execute(buf, f)
	entry.RenderedFileInfo = buf.String()
	buf.Reset()
	fileDataTpl.Execute(buf, f)
	entry.RenderedData = buf.String()

	return entry
}

type Package struct {
	Config
	Files map[string]*Entry
}

//Add a file to the package at the give path, the files is the location of a file on the filesystem.
func (p *Package) AddFile(pathKey string, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	e := NewEntry(f)
	p.Files[pathKey] = e
	f.Close()

	dirKey := filepath.Dir(pathKey)
	dir := filepath.Dir(path)

	if dirKey == "/" {
		dirKey = "."
	}

	// if dir and path are equal, we're already the root
	if dirKey == pathKey {
		return nil
	}
	if _, exists := p.Files[dirKey]; !exists {
		p.AddFile(dirKey, dir)
	}
	p.Files[dirKey].FileInfos = append(p.Files[dirKey].FileInfos, e.RenderedFileInfo)

	return nil
}

//Build the package
func (p *Package) Build(out io.Writer) error {
	return pkgTpl.Execute(out, p)
}

//Write the build to a file, you don't need to call Build.
func (p *Package) Write(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return p.Build(f)
}
