# Resources

Unfancy resources embedding with Go.

- No blings.
- No runtime dependency.
- Embeddable builder.

### Dude, Why?

Yes, there is quite a lot of projects that handles resource embeding but they come with more bling than you will probably ever need and you often ended up with having dependenies for your end project, not this time.

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
  -tag="": The tag to use for the generated package. Use empty [default] for no tag.
  -var="FS": The name of variable to assign the virtual-filesystem to.
```

### Techniques for "live" resources for development

```go
package main

import "net/http"
var Assets http.FileSystem 


func main() {
//Use Assets here
}
```

```go
// +build !embed

package main

import "net/http"

var Assets = http.Dir("./public")
```
Now when you build or run your project, you will have files directly served from `./public` directory. This is pretty helpful for development.

To embed your resources, do

```sh
$ resources -output="public_resources.go" -var="Assets" -tag="embed" public/*
$ go build -tags=embed
```

Now your resources should be embeded with your program!


### Contributing
Please consider opening an issue first, or just send a pull request. :)

### Credits
See [Contributors](https://github.com/omeid/go-resources/graphs/contributors).

### LICENSE
  [MIT](LICENSE).


### TODO

 - Add tests. 
 - Remove "net/http" dependency.
