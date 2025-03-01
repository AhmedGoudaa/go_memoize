#### BenchMarking and Performance Testing
#### Between the following 
- [x] [github.com/emad-elsaid/memoize](https://github.com/emad-elsaid/memoize)
- [x] [Reflection version](https://github.com/AhmedGoudaa/go_memoize/blob/bechmarking/memoize_ref.go)

#

#### Results are as follows:
for [github.com/emad-elsaid/memoize](https://github.com/emad-elsaid/memoize)
```shell
BenchmarkDo1With_memoize-10 | 36268526 | 321.7 ns/op | 336 B/op | 8 allocs/op
```

#

for [Reflection version](https://github.com/AhmedGoudaa/go_memoize/blob/bechmarking/memoize_ref.go)

```shell
BenchmarkDo1MemRef-10 | 70152396 | 168.9 ns/op |  64 B/op | 3 allocs/op
BenchmarkDo2MemRef-10 | 57438241 | 208.6 ns/op | 104 B/op | 4 allocs/op
BenchmarkDo3MemRef-10 | 47954330 | 249.2 ns/op | 152 B/op | 5 allocs/op
BenchmarkDo4MemRef-10 | 42692116 | 282.6 ns/op | 176 B/op | 6 allocs/op
```

for this package (go_memoize)
```shell
BenchmarkDo0Mem-10 | 811289566 | 14.77 ns/op | 0 B/op | 0 allocs/op
BenchmarkDo1Mem-10 | 676579908 | 18.26 ns/op | 0 B/op | 0 allocs/op
BenchmarkDo2Mem-10 | 578134332 | 20.99 ns/op | 0 B/op | 0 allocs/op
BenchmarkDo3Mem-10 | 533455237 | 22.67 ns/op | 0 B/op | 0 allocs/op
BenchmarkDo4Mem-10 | 487471639 | 24.73 ns/op | 0 B/op | 0 allocs/op
```