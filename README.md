# gorecursiveinterface
Working golang recursive interface for structures that return themselves

So I had a problem where I wanted to create an interface that already existing libraries fullfil but this interface included recursive methods. After many tries I found a working solution with wrapper and creation of recursive methods inside a function.

## Benchmarks

So I created bunch of simple benchmarks to test value and pointer based methods with and without wrapper with some basic arithmetic and these are the results:

```console
go test -benchmem -run=^$ -bench . github.com/adamdyszy/gorecursiveinterface/implementation/calc
goos: windows
goarch: amd64
pkg: github.com/adamdyszy/gorecursiveinterface/implementation/calc
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkPointerDirectChainCalc-8                       500280364                2.448 ns/op           0 B/op    
      0 allocs/op
BenchmarkPointerWrapperChainCalc-8                       2108157               529.8 ns/op           512 B/op    
      8 allocs/op
BenchmarkPointerManyOperationsDirectCalc-8               4970100               242.3 ns/op             0 B/op    
      0 allocs/op
BenchmarkPointerManyOperationsWrapperCalc-8                46998             31395 ns/op           25600 B/op    
    400 allocs/op
BenchmarkPointerDirectCalc-8                            31519392                32.48 ns/op            0 B/op    
      0 allocs/op
BenchmarkPointerWrapperCalc-8                           35576426                33.25 ns/op            0 B/op    
      0 allocs/op
BenchmarkPointerMixedDirectChainCalc-8                    126039              9218 ns/op               0 B/op    
      0 allocs/op
BenchmarkPointerMixedWrapperChainCalc-8                   116311             10126 ns/op             512 B/op    
      8 allocs/op
BenchmarkPointerHeavyInitMixedDirectChainCalc-8           131287              9079 ns/op               0 B/op    
      0 allocs/op
BenchmarkPointerHeavyInitMixedWrapperChainCalc-8          119268              9641 ns/op             512 B/op    
      8 allocs/op
BenchmarkDirectChainCalc-8                              1000000000               0.3788 ns/op          0 B/op    
      0 allocs/op
BenchmarkWrapperChainCalc-8                              1974439               650.2 ns/op           544 B/op    
     12 allocs/op
BenchmarkManyOperationsDirectCalc-8                     27694885                42.35 ns/op            0 B/op    
      0 allocs/op
BenchmarkManyOperationsWrapperCalc-8                       40173             34592 ns/op           27200 B/op    
    600 allocs/op
BenchmarkDirectCalc-8                                   93287932                11.23 ns/op            0 B/op    
      0 allocs/op
BenchmarkWrapperCalc-8                                  72433149                16.26 ns/op            0 B/op    
      0 allocs/op
BenchmarkMixedDirectChainCalc-8                           728155              1810 ns/op               0 B/op    
      0 allocs/op
BenchmarkMixedWrapperChainCalc-8                          388551              2623 ns/op             544 B/op    
     12 allocs/op
BenchmarkHeavyInitMixedDirectChainCalc-8                  612369              1721 ns/op               0 B/op    
      0 allocs/op
BenchmarkHeavyInitMixedWrapperChainCalc-8                 385116              2648 ns/op             544 B/op    
     12 allocs/op
PASS
ok      github.com/adamdyszy/gorecursiveinterface/implementation/calc   39.248s
```

The worst use cases are value based builders that are called in a loop for chaining. The more we use given wrapper and call the NoRecursion interface part methods the closer it gets to the performance of direct calls. Also when we are dealing with pointers it also looks better. Also worth mentioning is the memory impact that using wrapper has on the tests. It does introduce overhead.