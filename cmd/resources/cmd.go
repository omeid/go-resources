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
	pkg      = flag.String("package", "main", "The name of package to generate.")
	varName  = flag.String("var", "FS", "The name of variable to assign the virtual-filesystem to.")
	tag      = flag.String("tag", "", "The tag to use for the generated package. Defaults to not tag.")
	declare  = flag.Bool("declare", false, "Whether to declare the \"var\" or not.")
	out      = flag.String("output", "", "The filename to write the output to.")
	trimPath = flag.String("trim", "", "Path prefix to remove from the resulting file path")
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
	}

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
