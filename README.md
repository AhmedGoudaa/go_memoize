# go_memoize 

![Workflow Status](https://github.com/AhmedGoudaa/go_memoize/actions/workflows/ci.yml/badge.svg)

`go_memoize` package provides a set of functions to memoize the results of computations, allowing for efficient caching and retrieval of results based on input parameters. This can significantly improve performance for expensive or frequently called functions.

## Features
- Memoizes functions with TTL, supporting 0 to 7 comparable parameters. [List of Memoize Functions](https://github.com/AhmedGoudaa/go_memoize/blob/main/memoize.go)
- High performance, zero allocation, and zero dependencies.
- Utilizes the FNV-1a hash algorithm for caching.


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

## Functions & Usage Examples

<table>
  <tr>
    <th><code>Function</code></th>
    <th><code>Description</code></th>
    <th><code>Example</code></th>
  </tr>
  <tr>
    <td><code>Memoize</code></td>
    <td>Memoizes a function with no params</td>
    <td>
      <pre><code>
memoizedFn := Memoize(func() int { return 1 }, time.Minute)
result := memoizedFn()
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize1</code></td>
    <td>Memoizes a function with 1 param</td>
    <td>
      <pre><code>
memoizedFn := Memoize1(func(a int) int { return a * 2 }, time.Minute)
result := memoizedFn(5)
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize2</code></td>
    <td>Memoizes a function with 2 params</td>
    <td>
      <pre><code>
memoizedFn := Memoize2(func(a int, b string) string { return fmt.Sprintf("%d-%s", a, b) }, time.Minute)
result := memoizedFn(5, "example")
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize3</code></td>
    <td>Memoizes a function with 3 params</td>
    <td>
      <pre><code>
memoizedFn := Memoize3(func(a int, b string, c float64) string { return fmt.Sprintf("%d-%s-%f", a, b, c) }, time.Minute)
result := memoizedFn(5, "example", 3.14)
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize4</code></td>
    <td>Memoizes a function with 4 params</td>
    <td>
      <pre><code>
memoizedFn := Memoize4(func(a, b, c, d int) int { return a + b + c + d }, time.Minute)
result := memoizedFn(1, 2, 3, 4)
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize5</code></td>
    <td>Memoizes a function with 5 params</td>
    <td>
      <pre><code>
memoizedFn := Memoize5(func(a, b, c, d, e int) int { return a + b + c + d + e }, time.Minute)
result := memoizedFn(1, 2, 3, 4, 5)
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize6</code></td>
    <td>Memoizes a function with 6 params</td>
    <td>
      <pre><code>
memoizedFn := Memoize6(func(a, b, c, d, e, f int) int { return a + b + c + d + e + f }, time.Minute)
result := memoizedFn(1, 2, 3, 4, 5, 6)
      </code></pre>
    </td>
  </tr>
  <tr>
    <td><code>Memoize7</code></td>
    <td>Memoizes a function with 7 params</td>
    <td>
      <pre><code>
memoizedFn := Memoize7(func(a, b, c, d, e, f, g int) int { return a + b + c + d + e + f + g }, time.Minute)
result := memoizedFn(1, 2, 3, 4, 5, 6, 7)
      </code></pre>
    </td>
  </tr>
</table>


### [Benchmarking result]((https://github.com/AhmedGoudaa/go_memoize/blob/main/benchmarks/benchmark_test.go))
 Device "Apple M2 Pro"

```
goos: darwin
goarch: arm64
BenchDo0Mem | 785610645 | 15.27 ns/op | 0 B/op | 0 allocs/op
BenchDo1Mem | 633380330 | 19.02 ns/op | 0 B/op | 0 allocs/op
BenchDo2Mem | 553163156 | 21.78 ns/op | 0 B/op | 0 allocs/op
BenchDo3Mem | 482770053 | 25.11 ns/op | 0 B/op | 0 allocs/op
BenchDo4Mem | 471314336 | 25.68 ns/op | 0 B/op | 0 allocs/op
```

This project is licensed under the MIT License. See the [`LICENSE`](https://github.com/AhmedGoudaa/go_memoize/blob/main/LICENSE) file for details.
