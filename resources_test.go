package resources

import "testing"

//go:generate go build -o testdata/resources github.com/omeid/go-resources/cmd/resources
//go:generate testdata/resources -package test -var FS  -output testdata/generated/store_prod.go  testdata/*.txt testdata/*.sql

func TestPackage(t *testing.T) {

	// err = p.Add(path string, file File)
	// err = p.AddFile(path string, file string)
	// err = p.Build(out io.Writer)
	// err = p.Write(path string)

}
