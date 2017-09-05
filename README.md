# Resources [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/omeid/go-resources)  [![Build Status](https://travis-ci.org/omeid/go-resources.svg?branch=master)](https://travis-ci.org/omeid/go-resources) [![Go Report Card](https://goreportcard.com/badge/github.com/omeid/go-resources?bust=true)](https://goreportcard.com/report/github.com/omeid/go-resources)
Unfancy resources embedding with Go.

- No blings.
- No runtime dependency.
- Idiomatic Library First design.

### Dude, Why?

Yes, there is quite a lot of projects that handles resource embedding but they come with more bling than you ever need and you often end up with having dependencies for your end project, not this time.

### Installing

Just go get it!

```sh
$ go get github.com/omeid/go-resources/cmd/resources
```


### Usage

```
$ resources -h
Usage resources:
  -declare
        whether to declare the -var (default false)
  -fmt
        run output through gofmt, this is slow for huge files (default false)
  -output filename
        filename to write the output to
  -package name
        name of the package to generate (default "main")
  -tag tag
        tag to use for the generated package (default no tag)
  -trim prefix
        path prefix to remove from the resulting file path in the virtual filesystem
  -var name
        name of the variable to assign the virtual filesystem to (default "FS")
  -width number
        number of content bytes per line in generetated file (default 12)
```

### Optimization
Generating resources result in a very high number of lines of code, 1mb of resources result about 5mb of code at over 87 thousand lines of code, _don't worry, the size of data stored in your binary is exactly same as the resources (eg. 1mb of resources only increases your binary size by 1mb)_, compiling this many lines of code takes time and slows down the compiler.  
To avoid recompiling the resources every time and leverage the `go build` cache, generate your resources into a standalone package and then import it, this will allow for faster iteration as you don't have to wait for the resources to be compiled with every change.

##### "Live" development of resources 
For fast iteration and improvement of your resources, you can work around the compile with the following technique: 

```go
package main

import "net/http"
var Assets http.FileSystem 

func main() {
  if Assets == nil {
    panic("No Assets. Have you generated the resources?")
  }

  //Use Assets here
}
```

```go
// +build !embed

package main

import (
	"net/http"

	live "github.com/omeid/go-resources/live"
)


var Assets = live.Dir("./public")
```
Now when you build or run your project, you will have files directly served from `./public` directory.

And then to embed your resources, do

```sh
$ resources -output="public_resources.go" -var="Assets" -tag="embed" public/*
$ go build -tags=embed
```

Now your resources should be embedded with your program!  
Of course, you may use any `var` name or tag you please.

### Go Generate
There is a few reasons to avoid resource embedding in Go generate,
first Go Generate is for generating Go source code from your code, generally the resources you want to embed aren't effected by the Go source directly and as such generating resources are out of the scope of Go Generate.
Second, You're unnecessarily slowing down code iterations by blocking `go generate` for resource generation.

But if you must use, put the `//go:generate resources` followed by the usual flags on the command-line somewhere in your Go files.

# Resources, The Library [![GoDoc](https://godoc.org/github.com/omeid/go-resources?status.svg)](https://godoc.org/github.com/omeid/go-resources)
The resource generator is written as a library and isn't bound to filesystem by the way of accepting files in the form 
```go
type File interface {
      io.Reader
      Stat() (os.FileInfo, error)
}
```
along with a helper method that adds files from filesystem, this allows to integrate go-resources with ease in your workflow when the when the provided command doesn't fit well, for an example see the [Gonzo binding](https://github.com/go-gonzo/resources/blob/master/resources.go) of go-resources.  
Please refer to the [GoDoc](https://godoc.org/github.com/omeid/go-resources) for complete documentation.


### Strings

The generated FileSystem also implements an `String(string) (string, bool)` method that allows you to read the content of a file as string, to use that
instead of defining your file Assets variable as simply an http.FileSystem, do the following:
```go

type Resources interface {
	http.FileSystem
	String(string) (string, bool)
}

var Assets Resources
```

Now you can call `Assets.String(someFile)` and get the content as string with a boolean value indicating whatever the file was found or not.


===

### Contributing
Please consider opening an issue first, or just send a pull request. :)

### Credits
See [Contributors](https://github.com/omeid/go-resources/graphs/contributors).

### LICENSE
  [MIT](LICENSE).


### TODO
 - Add tests. 
