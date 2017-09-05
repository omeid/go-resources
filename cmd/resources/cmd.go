// Unfancy resources embedding with Go.

package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/omeid/go-resources"
)

var (
	pkg      = "main"
	varName  = "FS"
	tag      = ""
	declare  = false
	out      = ""
	trimPath = ""
	width    = resources.BlockWidth
	gofmt    = false
)

type nope struct{}

func main() {
	flag.StringVar(&pkg, "package", pkg, "`name` of the package to generate")
	flag.StringVar(&varName, "var", varName, "`name` of the variable to assign the virtual filesystem to")
	flag.StringVar(&tag, "tag", tag, "`tag` to use for the generated package (default no tag)")
	flag.BoolVar(&declare, "declare", declare, "whether to declare the -var (default false)")
	flag.StringVar(&out, "output", out, "`filename` to write the output to")
	flag.StringVar(&trimPath, "trim", trimPath, "path `prefix` to remove from the resulting file path in the virtual filesystem")
	flag.IntVar(&width, "width", width, "`number` of content bytes per line in generetated file")
	flag.BoolVar(&gofmt, "fmt", gofmt, "run output through gofmt, this is slow for huge files (default false)")
	flag.Parse()

	if out == "" {
		flag.PrintDefaults()
		log.Fatal("-output is required.")
	}

	config := resources.Config{
		Pkg:     pkg,
		Var:     varName,
		Tag:     tag,
		Declare: declare,
		Format:  gofmt,
	}
	resources.BlockWidth = width

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

	t0 := time.Now()

	for file := range files {
		path := strings.TrimPrefix(file, trimPath)
		err := res.AddFile(path, file)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := res.Write(out); err != nil {
		log.Fatal(err)
	}

	log.Printf("Finished in %v. Wrote %d resources to %s", time.Since(t0), len(files), out)
}
