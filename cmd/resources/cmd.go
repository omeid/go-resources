package main

import (
	"flag"
	"path/filepath"

	"github.com/omeid/go-resources"
)

var (
	Pkg        = flag.String("package", "main", "The name of package to generate.")
	Var        = flag.String("var", "FS", "The name of variable to assign the virtual-filesystem to.")
	Tag        = flag.String("tag", "embed", "The tag to use for the generated package. Use empty for no tag.")
	Declare    = flag.Bool("declare", false, "Whatever to declare the \"var\" or not.")
	FileFormat = flag.String("filenameFormat", resources.FilenameFormat, "The template to use for generated files name.")
)

func main() {

	flag.Parse()

	config := resources.Config{*Pkg, *Var, *Tag, *Declare}

	files := map[string]struct{}{}

	for _, g := range flag.Args() {
		files, err := filepath.Glob(g)
	}
}
