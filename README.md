# Go Collection Utils

Go Collection Utils is a utility library for working with collections in Go. It provides a set of functions for common collection operations such as `Find`, `Map`, `Reduce`, and `Filter`.

This library is an alternative to [`go-funk`](https://github.com/thoas/go-funk) and leverages Go generics instead of `any` and reflection for better type safety and performance.

## Table of Contents

- [Features](#features)
  - [Selection](#selection)
  - [Transformation](#transformation)
  - [Aggregation](#aggregation)
  - [Statistics](#statistics)
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

### Selection

- **All**: Check if all elements in a collection satisfy a given predicate.
- **Any**: Check if any element in a collection satisfies a given predicate.
- **Contains**: Check if a collection contains a given element.
- **Count**: Count the number of elements in a collection that satisfy a given predicate.
- **Distinct**: Get distinct elements from a collection.
- **Filter**: Select elements from a collection that satisfy a given predicate.
- **Find**: Locate an element in a collection that matches a given predicate.

### Transformation

- **FlatMap**: Transform each element in a collection using a provided function and flatten the result.
- **GroupBy**: Group elements in a collection by a given key.
- **Map**: Transform each element in a collection using a provided function.

### Aggregation

- **Reduce**: Accumulate values in a collection into a single result using a reducer function.

### Statistics

- **Average**: Calculate the average (mean) of elements in a collection.
- **Max**: Find the maximum value in a collection.
- **Median**: Calculate the median of elements in a collection.
- **Min**: Find the minimum value in a collection.
- **Mode**: Calculate the mode (most frequent element) of elements in a collection.
- **Product**: Calculate the product of elements in a collection.
- **Range**: Calculate the range of elements in a collection.
- **Standard Deviation**: Calculate the standard deviation of elements in a collection.
- **Sum**: Calculate the sum of elements in a collection.
- **Variance**: Calculate the variance of elements in a collection.

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

To see what commands are available, run:

```shell
make help
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
