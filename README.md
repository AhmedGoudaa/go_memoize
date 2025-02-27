# go_memoize Package

`go_memoize` package provides a set of functions to memoize the results of computations, allowing for efficient caching and retrieval of results based on input parameters. This can significantly improve performance for expensive or frequently called functions.  This package uses the FNV-1a hash algorithm and has zero dependencies.
This package uses the FNV-1a hash algorithm and has zero dependencies.

## Installation

To install the package, use `go get`:

```sh
go get github.com/yourusername/go_memoize
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
    "github.com/yourusername/go_memoize"
)

func main() {
    computeFn := func(a int, b string) string {
        // Simulate an expensive computation
        time.Sleep(2 * time.Second)
        return fmt.Sprintf("%d-%s", a, b)
    }

    memoizedFn := memoize.Memoize2(computeFn, 10*time.Second)

    // First call will compute the result
    result := memoizedFn(5, "example")
    fmt.Println(result) // Output: 5-example

    // Subsequent calls within 10 seconds will use the cached result
    result = memoizedFn(5, "example")
    fmt.Println(result) // Output: 5-example
}
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.