# gotl
A Go Template Library. A bunch of utils and collections that are not part of the Golang standard library.

### Prerequisites

This library is using extensively Golang Generics, a new feature introduced in Go 1.18. Please make sure that your version of Go compiler supports this feature.

You can also use the [gotip](https://pkg.go.dev/golang.org/dl/gotip), a develpment version of Go where you likely find a support for the generics.

### Description

This library contains implementation of various data structures like stack, queue, priority queue, graph. Also, there are some common algorithms that operates on these data structures.

This is an attempt to expand the standard Go library.

### Installation
Cd to your module and:
```shell
$ go get github.com/rrowniak/gotl
```

### Usage
```go
package main

import (
	"fmt"
	"github.com/rrowniak/gotl"
)

func main() {
	var s gotl.Stack[int]
	s.Push(1)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Empty())
}
```