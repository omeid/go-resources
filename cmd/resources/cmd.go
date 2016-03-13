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
	Pkg      = flag.String("package", "main", "The name of package to generate.")
	Var      = flag.String("var", "FS", "The name of variable to assign the virtual-filesystem to.")
	Tag      = flag.String("tag", "", "The tag to use for the generated package. Defaults to not tag.")
	Declare  = flag.Bool("declare", false, "Whether to declare the \"var\" or not.")
	Out      = flag.String("output", "", "The filename to write the output to.")
	TrimPath = flag.String("trim", "", "Path prefix to remove from the resulting file path")
)

func main() {

	flag.Parse()

	if *Out == "" {
		flag.PrintDefaults()
		log.Fatal("-output is required.")
	}

	config := resources.Config{
		Pkg:     *Pkg,
		Var:     *Var,
		Tag:     *Tag,
		Declare: *Declare,
	}

	res := resources.New()
	res.Config = config

	files := make(map[string]struct{})

	for _, g := range flag.Args() {
		matches, err := filepath.Glob(g)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range matches {
			files[m] = struct{}{}
		}
	}

	for file, _ := range files {
		path := strings.TrimPrefix(file, *TrimPath)
		err := res.AddFile(path, file)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := res.Write(*Out)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Done. Wrote to %s", *Out)

}
