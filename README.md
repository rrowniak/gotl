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
	"github.com/rrowniak/gotl/cont"
)

func main() {
	var s cont.Stack[int]
	s.Push(1)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Empty())
}
```

### Features
#### Stack
A generic implementation of the stack data structure. A self explanatory example:
```go
var s cont.Stack[string]
s.Push("bottom")
s.Push("top")

if s.Top() != "top" {
    t.Errorf("Expected 'top', got '%s'", s.Top())
}

s.Pop()

if s.Top() != "bottom" {
    t.Errorf("Expected 'bottom', got '%s'", s.Top())
}

s.Pop()

if !s.Empty() {
    t.Error("Stack is expected to be empty")
}
```
#### Queue
A generic implementation of the queue data structure. A self explanatory example:
```go
var q cont.Queue[int]
q.Push(1)
q.Push(2)

if q.Len() != 2 {
    t.Errorf("Expected queue len == 2, got %d", q.Len())
}

for i := 1; i < 3; i++ {
    if q.Front() != i {
        t.Errorf("Expected to have %d, got %d", i, q.Front())
    }
    q.Pop()
}

if !q.Empty() {
    t.Errorf("Queue should be empty now, but has %d element(s)", q.Len())
}
```
#### Priority queue
Implementation of a heap based priority queue.
#### Circular buffer
Concept explanation [here](https://en.wikipedia.org/wiki/Circular_buffer).
#### Graph
A graph library.
#### Algos
Algorithms and handy utils.
