// Unfancy resources embedding with Go.

package resources

import (
	"io"
	"os"
)

type File interface {
	io.Reader
	Stat() (os.FileInfo, error)
}

//Create a new Package.
func New() *Package {
	return &Package{
		Config: Config{
			Pkg:     "resources",
			Var:     "FS",
			Declare: true,
		},
		Files: make(map[string]File),
	}
}

//Configuration defines some details about the output Go file.
type Config struct {
	Pkg     string // Package name
	Var     string // Variable name to assign the file system to.
	Tag     string // Build tag, leave empty for no tag.
	Declare bool   // Dictates whatever there should be a defintion Variable
}

type Package struct {
	Config
	Files map[string]File
}

//Add a file to the package at the give path.
func (p *Package) Add(path string, file File) {
	p.Files[path] = file
}

//Add a file to the package at the give path, the files is the location of a file on the filesystem.
func (p *Package) AddFile(path string, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	p.Files[path] = f
	return nil
}

//Build the package
func (p *Package) Build(out io.Writer) error {
	return pkg.Execute(out, p)
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
