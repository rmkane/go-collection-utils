# Go Collection Utils

Go Collection Utils is a utility library for working with collections in Go. It provides a set of functions for common collection operations such as `Find`, `Map`, `Reduce`, and `Filter`.

This library is an alternative to `go-funk` and leverages Go generics instead of `any` and reflection for better type safety and performance.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Find](#find)
  - [Map](#map)
  - [Reduce](#reduce)
  - [Filter](#filter)
- [Development](#development)
- [License](#license)

## Features

The Go Collection Utils library provides the following features:

### Basic Operations

- **Filter**: Select elements from a collection that satisfy a given predicate.
- **Map**: Transform each element in a collection using a provided function.
- **Reduce**: Accumulate values in a collection into a single result using a reducer function.
- **Find**: Locate an element in a collection that matches a given predicate.

### Utility Operations

- **All**: Check if all elements in a collection satisfy a given predicate.
- **Any**: Check if any element in a collection satisfies a given predicate.
- **Contains**: Check if a collection contains a given element.
- **Count**: Count the number of elements in a collection that satisfy a given predicate.
## Installation

To install the library, use `go get`:

```sh
go get github.com/rmkane/go-collection-utils
```

## Usage

### Find

```go
package main

import (
    "fmt"

    "github.com/rmkane/go-collection-utils"
)

func main() {
    collection := []int{1, 2, 3, 4, 5}
    result, found := collections.Find(collection, func(x int) bool {
        return x%2 == 0
    })
    fmt.Println(result, found) // Output: 2 true
}
```

### Map

```go
package main

import (
	"fmt"

	"github.com/rmkane/go-collection-utils"
)

func main() {
    collection := []int{1, 2, 3, 4, 5}
    result := collections.Map(collection, func(x int) int {
        return x * 2
    })
    fmt.Println(result) // Output: [2 4 6 8 10]
}
```

### Reduce

```go
package main

import (
	"fmt"

	"github.com/rmkane/go-collection-utils"
)

func main() {
    collection := []int{1, 2, 3, 4, 5}
    result := collections.Reduce(collection, func(acc, x int) int {
        return acc + x
    }, 0)
    fmt.Println(result) // Output: 15
}
```

### Filter

```go
package main

import (
	"fmt"

	"github.com/rmkane/go-collection-utils"
)

func main() {
    collection := []int{1, 2, 3, 4, 5}
    result := collections.Filter(collection, func(x int) bool {
        return x%2 == 0
    })
    fmt.Println(result) // Output: [2 4]
}
```

## Development

To see what commands are available, run `make help`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
