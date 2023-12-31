<p align="center">
<img src="./assets/photos/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/erfanmomeniii/heap?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.1-red?style=for-the-badge&logo=none" alt="version" />
</p>

# heap
<i>heap</i> is a lightweight package that provides binary min and max binary heap in Go. for seeing usages check [this](https://pkg.go.dev/github.com/erfanmomeniii/heap).

# Documentation

## Install

```bash
go get github.com/erfanmomeniii/heap
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/heap"
``` 

## Quick Start
The following examples illustrates how to use this package for creating max and min binary heap tree:

### Max heap
```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/heap"
)

func main() {
	h := heap.NewMax()

	h.Insert(4)
	h.Insert(12)
	h.Insert(10)
	
	max,_:= h.GetMax()
	fmt.Println(max)
	// 12
	
	h.Delete()
	
	max,_=h.GetMax()
	fmt.Println(max)
	// 10
}
```

### Min heap 
```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/heap"
)

func main() {
	h := heap.NewMin()

	h.Insert(4)
	h.Insert(12)
	h.Insert(10)
	
	min,_:= h.GetMin()
	fmt.Println(min)
	// 4
	
	h.Delete()
	
	min,_=h.GetMin()
	fmt.Println(min)
	// 10
}
```
We can also update the value of a heap node by using an alias. Here is an example:
```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/heap"
)

func main() {
	h := heap.NewMin()

	h.Insert(4, "number 1")
	h.Insert(12, "number 2")
	h.Insert(10, "number 3")

	m, a := h.GetMin()
	fmt.Println(m,a)
	// 4
	// number 1

	h.Update("number 1", 13)
	
	m, a = h.GetMin()
	fmt.Println(m,a)
	// 10
	// number 3
}
```

## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
