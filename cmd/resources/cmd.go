// Unfancy resources embedding with Go.

package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/omeid/go-resources"
)

var (
	pkg      = flag.String("package", "main", "`name` of the package to generate")
	varName  = flag.String("var", "FS", "`name` of the variable to assign the virtual filesystem to")
	tag      = flag.String("tag", "", "`tag` to use for the generated package (default no tag)")
	declare  = flag.Bool("declare", false, "whether to declare the -var (default false)")
	out      = flag.String("output", "", "`filename` to write the output to")
	trimPath = flag.String("trim", "", "path `prefix` to remove from the resulting file path in the virtual filesystem")
	width    = flag.Int("width", 12, "`number` of content bytes per line in generetated file")
	gofmt    = flag.Bool("fmt", false, "run output through gofmt, this is slow for huge files (default false)")
)

type nope struct{}

func main() {
	flag.Parse()

	if *out == "" {
		flag.PrintDefaults()
		log.Fatal("-output is required.")
	}

	config := resources.Config{
		Pkg:     *pkg,
		Var:     *varName,
		Tag:     *tag,
		Declare: *declare,
		Format:  *gofmt,
	}
	resources.BlockWidth = *width

	res := resources.New()
	res.Config = config

	files := make(map[string]nope)

	for _, g := range flag.Args() {
		matches, err := filepath.Glob(g)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range matches {
			files[m] = nope{}
		}
	}

	for file := range files {
		path := strings.TrimPrefix(file, *trimPath)
		err := res.AddFile(path, file)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := res.Write(*out)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Done. Wrote to %s", *out)

}
