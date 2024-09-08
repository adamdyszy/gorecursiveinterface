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

I made new benchmarks with optimized version, but it doesn't work when we want global init and usage without knowing the type that will be used. So I added new category to benchmarks WrapperGlob, here are results:

```console
go test -benchmem -run=^$ -bench . github.com/adamdyszy/gorecursiveinterface/implementation/calc
goos: windows
goarch: amd64
pkg: github.com/adamdyszy/gorecursiveinterface/implementation/calc
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkPointerDirectChainCalc-8                       494060974                2.376 ns/op           0 B/op              0 allocs/op
BenchmarkPointerWrapperChainCalc-8                      173702611                7.510 ns/op           0 B/op              0 allocs/op
BenchmarkPointerWrapperGlobChainCalc-8                   2606378               461.8 ns/op           256 B/op              8 allocs/op
BenchmarkPointerManyOperationsDirectCalc-8               5091511               243.6 ns/op             0 B/op              0 allocs/op
BenchmarkPointerManyOperationsWrapperCalc-8              3136461               358.0 ns/op             0 B/op              0 allocs/op
BenchmarkPointerManyOperationsWrapperGlobCalc-8            54733             22070 ns/op           12800 B/op            400 allocs/op
BenchmarkPointerDirectCalc-8                            33220843                33.94 ns/op            0 B/op              0 allocs/op
BenchmarkPointerWrapperCalc-8                           33519552                32.98 ns/op            0 B/op              0 allocs/op
BenchmarkPointerWrapperGlobCalc-8                       34595391                33.48 ns/op            0 B/op              0 allocs/op
BenchmarkPointerMixedDirectChainCalc-8                    126073              9025 ns/op               0 B/op              0 allocs/op
BenchmarkPointerMixedWrapperChainCalc-8                   126804              9460 ns/op               0 B/op              0 allocs/op
BenchmarkPointerMixedWrapperGlobChainCalc-8               119727              9603 ns/op             256 B/op              8 allocs/op
BenchmarkPointerHeavyInitMixedDirectChainCalc-8           130136              9251 ns/op               0 B/op              0 allocs/op
BenchmarkPointerHeavyInitMixedWrapperChainCalc-8          129837              9394 ns/op               0 B/op              0 allocs/op
BenchmarkPointerHeavyInitMixedWrapperGlobChainCalc-8      117151              9623 ns/op             256 B/op              8 allocs/op
BenchmarkDirectChainCalc-8                               2640057               390.8 ns/op            16 B/op              1 allocs/op
BenchmarkWrapperChainCalc-8                              3205514               396.7 ns/op            16 B/op              1 allocs/op
BenchmarkWrapperGlobChainCalc-8                          1000000              1012 ns/op             304 B/op             13 allocs/op
BenchmarkManyOperationsDirectCalc-8                      2379189               528.8 ns/op            16 B/op              2 allocs/op
BenchmarkManyOperationsWrapperCalc-8                     1658659               818.8 ns/op            16 B/op              2 allocs/op
BenchmarkManyOperationsWrapperGlobCalc-8                   43846             27029 ns/op           14416 B/op            602 allocs/op
BenchmarkDirectCalc-8                                   86552608                11.77 ns/op            0 B/op              0 allocs/op
BenchmarkWrapperCalc-8                                  84598789                14.44 ns/op            0 B/op              0 allocs/op
BenchmarkWrapperGlobCalc-8                              78587498                13.97 ns/op            0 B/op              0 allocs/op
BenchmarkMixedDirectChainCalc-8                           488338              2498 ns/op              15 B/op              1 allocs/op
BenchmarkMixedWrapperChainCalc-8                          501043              2092 ns/op              15 B/op              1 allocs/op
BenchmarkMixedWrapperGlobChainCalc-8                      384610              3063 ns/op             304 B/op             13 allocs/op
BenchmarkHeavyInitMixedDirectChainCalc-8                  535161              2042 ns/op              16 B/op              2 allocs/op
BenchmarkHeavyInitMixedWrapperChainCalc-8                 508276              2123 ns/op              16 B/op              2 allocs/op
BenchmarkHeavyInitMixedWrapperGlobChainCalc-8             450966              2700 ns/op             304 B/op             14 allocs/op
```

So as we can see the main problem is related with not knowing the type needed for global variable.