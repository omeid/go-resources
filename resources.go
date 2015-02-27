// Unfancy resources embedding with Go.

package resources

import (
	"bytes"
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
// Pkg      the package name to use.
// Var      the variable name to assign the file system to.
// Tag      the build tag for the generated file, leave empty for not tag.
// Declare  dictates whatever there should be a defintion for the variable
//          in the output file or not, it will use the type http.FileSystem.
type Config struct {
	Pkg     string
	Var     string
	Tag     string
	Declare bool
}

type Package struct {
	Config
	Files map[string]File
}

//Add a file to the package at the give path.
func (p *Package) Add(path string, file File) {
	p.Files[path] = file
}

func (p *Package) AddFile(file string, path string) error {
  f, err := os.Open(file)
  if err != nil {
	return err
  }

  p.Files[path] = f

	return nil
}

//Build the package
func (p *Package) Build() (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	return out, pkg.Execute(out, p)
}

//Write the build to a file.
func (p *Package) Write(path string) error {
  f, err := os.Create(path)
  if err != nil {
	return err
  }
  defer f.Close()

  buff, err := p.Build()

  if err != nil {
	return err
  }

  _, err = buff.WriteTo(f)
  return err
}
