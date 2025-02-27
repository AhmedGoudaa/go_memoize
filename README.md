# go_memoize Package

![Workflow Status](https://github.com/AhmedGoudaa/go_memoize/actions/workflows/ci.yml/badge.svg)

`go_memoize` package provides a set of functions to memoize the results of computations, allowing for efficient caching and retrieval of results based on input parameters. This can significantly improve performance for expensive or frequently called functions.

#### This package uses the FNV-1a hash algorithm and has zero dependencies.

### Note: All Memoize functions take functions with parameters that are only comparable.

## Installation

To install the package, use `go get`:

```sh
go get github.com/AhmedGoudaa/go_memoize
```

## Usage

### Basic Memoization

The `Memoize` function can be used to memoize a function with no parameters:

```go
computeFn := func() int {
    // Expensive computation
    return 42
}

memoizedFn := Memoize(computeFn, 10*time.Second)
result := memoizedFn()
```

### Memoization with Parameters

The package provides functions to memoize functions with up to 7 parameters. Here are some examples:

#### One Parameter

```go
computeFn := func(a int) int {
    // Expensive computation
    return a * 2
}

memoizedFn := Memoize1(computeFn, 10*time.Second)
result := memoizedFn(5)
```

#### Two Parameters

```go
computeFn := func(a int, b string) string {
    // Expensive computation
    return fmt.Sprintf("%d-%s", a, b)
}

memoizedFn := Memoize2(computeFn, 10*time.Second)
result := memoizedFn(5, "example")
```

#### Three Parameters

```go
computeFn := func(a int, b string, c float64) string {
    // Expensive computation
    return fmt.Sprintf("%d-%s-%f", a, b, c)
}

memoizedFn := Memoize3(computeFn, 10*time.Second)
result := memoizedFn(5, "example", 3.14)
```

### Cache Management

The `Cache` struct is used internally to manage the cached entries. It supports setting, getting, and deleting entries, as well as computing new values if they are not already cached or have expired.

## Example

Here is a complete example of using the `memoize` package:

```go
package main

import (
    "fmt"
    "time"
    m "github.com/AhmedGoudaa/go_memoize"
)

func main() {
    computeFn := func(a int, b string) string {
        // Simulate an expensive computation
        time.Sleep(2 * time.Second)
        return fmt.Sprintf("%d-%s", a, b)
    }

    memoizedFn := m.Memoize2(computeFn, 10*time.Second)

    // First call will compute the result
    result := memoizedFn(5, "example")
    fmt.Println(result) // Output: 5-example

    // Subsequent calls within 10 seconds will use the cached result
    result = memoizedFn(5, "example")
    fmt.Println(result) // Output: 5-example
}
```

## List of Memoize Functions

- `Memoize`: Memoizes a function with no parameters.
- `Memoize1`: Memoizes a function with 1 parameter.
- `Memoize2`: Memoizes a function with 2 parameters.
- `Memoize3`: Memoizes a function with 3 parameters.
- `Memoize4`: Memoizes a function with 4 parameters.
- `Memoize5`: Memoizes a function with 5 parameters.
- `Memoize6`: Memoizes a function with 6 parameters.
- `Memoize7`: Memoizes a function with 7 parameters.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```