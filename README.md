# Resources

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
```sh
$ resources -help
Usage of resources:
  -declare=false: Whether to declare the "var" or not.
  -output="": The filename to write the output to.
  -package="main": The name of package to generate.
  -tag="": The tag to use for the generated package. Defaults to not tag.
  -trim="": Path prefix to remove from the resulting file path
  -var="FS": The name of variable to assign the virtual-filesystem to.
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

	here "github.com/omeid/go-here"
  // Here is used to find the absolute path relative to the source code
  // this allows you to use this trick in non-main packages and call
  // go run/test et al from any location.
)


var Assets = http.Dir(here.Abs("./public"))
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



===

### Contributing
Please consider opening an issue first, or just send a pull request. :)

### Credits
See [Contributors](https://github.com/omeid/go-resources/graphs/contributors).

### LICENSE
  [MIT](LICENSE).


### TODO
 - Add tests. 
