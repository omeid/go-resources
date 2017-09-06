# Resources [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/omeid/go-resources)  [![Build Status](https://travis-ci.org/omeid/go-resources.svg?branch=master)](https://travis-ci.org/omeid/go-resources) [![Go Report Card](https://goreportcard.com/badge/github.com/omeid/go-resources?bust=true)](https://goreportcard.com/report/github.com/omeid/go-resources)
Unfancy resources embedding with Go.

- No blings.
- No runtime dependency.
- Idiomatic Library First design.

### Dude, Why?

Yes, there is quite a lot of projects that handles resource embedding
but they come with more bling than you ever need and you often end up
with having dependencies for your end project. Not this time.

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

Generating resources result in a very high number of lines of code, 1MB
of resources result about 5MB of code at over 87,000 lines of code. This
is caused by the chosen representation of the file contents within the
generated file.

Instead of a (binary) string, `resources` transforms each file into an
actual byte slice. For example, a file with content `Hello, world!` will
be represented as follows:

``` go
FS = &FileSystem{
  "/hello.txt": File{
    data: []byte{
      0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64,
      0x21,
    },
    fi: FileInfo{
      name:    "hello.txt",
      size:    13,
      modTime: time.Unix(0, 1504640959536230658),
      isDir:   false,
    },
  },
}
```

While this seems wasteful, the compiled binary is not really affected.
_If you add 1MB of resources, your binary will increase 1MB as well_.

However, compiling this many lines of code takes time and slows down the
compiler. To avoid recompiling the resources every time and leverage the
compiler cache, generate your resources into a standalone package and
then import it, this will allow for faster iteration as you don't have
to wait for the resources to be compiled with every change.

``` sh
mkdir -p assets
resources -declare -var=FS -package=assets -output=assets/assets.go your/files/here
```

``` go
package main

import "importpath/to/assets"

func main() {
  data, err := assets.FS.Open("your/files/here")
  // ...
}
```

##### "Live" development of resources

For fast iteration and improvement of your resources, you can work
around the compile with the following technique:

First, create a normal `main.go`:

```go
package main

import "net/http"

var Assets http.FileSystem

func main() {
  if Assets == nil {
    panic("No Assets. Have you generated the resources?")
  }

  // use Assets here
}
```

Then, add a second file in the same package (`main` here), with the
following content:

```go
// +build !embed

package main

import (
	"net/http"

	"github.com/omeid/go-resources/live"
)

var Assets = live.Dir("./public")
```

Now when you build or run your project, you will have files directly
served from `./public` directory.

To create a *production build*, i.e. one with the embedded files, build
the resouces with `-tag=embed` and add the `embed` tag to `go build`:

```sh
$ resources -output=public_resources.go -var=Assets -tag=embed public/*
$ go build -tags=embed
```

Now your resources should be embedded with your program!
Of course, you may use any `var` or `tag` name you please.

### Go Generate

There is a few reasons to avoid resource embedding in `go generate`.

First `go generate` is for generating Go source code from your code,
generally the resources you want to embed aren't effected by the Go
source directly and as such generating resources are slightly out of the
scope of `go generate`.

Second, you're unnecessarily slowing down code iterations by blocking
`go generate` for resource generation.

# Resources, The Library [![GoDoc](https://godoc.org/github.com/omeid/go-resources?status.svg)](https://godoc.org/github.com/omeid/go-resources)

The resource generator is written as a library and isn't bound to
filesystem by the way of accepting files in the form

```go
type File interface {
      io.Reader
      Stat() (os.FileInfo, error)
}
```

along with a helper method that adds files from filesystem.

This allows to integrate `resources` with ease in your workflow when the
when the provided command doesn't fit well, for an example see the [Gonzo
binding](https://github.com/go-gonzo/resources/blob/master/resources.go)
`resources`.

Please refer to the [GoDoc](https://godoc.org/github.com/omeid/go-resources)
for complete documentation.

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

---

### Contributing

Please consider opening an issue first, or just send a pull request. :)

### Credits

See [Contributors](https://github.com/omeid/go-resources/graphs/contributors).

### LICENSE

[MIT](LICENSE).

### TODO

- Add tests.
