# perf-go
Benchmarks of different approaches for common problem in go

## Run tests
```
make
```

## Results (On my PC ;) )

|Counter|Total|ns per op|
|-|-|-|
|BenchmarkCounterAtomic-4|69815256|18.02 ns/op|
|BenchmarkCounterMutex-4|24769442|46.93 ns/op|

|Factorial|Total|ns per op|
|-|-|-|
|BenchmarkFactorialSerial/FactorialSerial-20-4|113081472|10.46 ns/op|
|BenchmarkFactorialSerial/FactorialRecursive-20-4|27576867|40.32 ns/op|
|BenchmarkFactorialParallel/FactorialSerial-20-4|276191018|4.326 ns/op|
|BenchmarkFactorialParallel/FactorialRecursive-20-4|72123966|15.11 ns/op|

|Fibbonacci|Total|ns per op|
|-|-|-|
|BenchmarkFibonnaciSerial/FibonacciSumClassic-20-4|31941|37669 ns/op|
|BenchmarkFibonnaciSerial/FibonacciSumDP-20-4|59457067|18.42 ns/op|
|BenchmarkFibbonacciParallel-4|75055|15622 ns/op|

|FizzBuzz|Total|ns per op|
|-|-|-|
|BenchmarkFizzBuzzSer/v1-1-4|290979|3659 ns/op|
|BenchmarkFizzBuzzSer/v2-1-4|250216|5175 ns/op|
|BenchmarkFizzBuzzSer/v3-1-4|262671|4178 ns/op|
|BenchmarkFizzBuzzPar/v1-1-4|526555|1991 ns/op|
|BenchmarkFizzBuzzPar/v2-1-4|480120|2218 ns/op|
|BenchmarkFizzBuzzPar/v3-1-4|494432|2205 ns/op|

|Is Prime|Total|ns per op|
|-|-|-|
|BenchmarkIsPrimeSerial/PrimeV1-1-4|51702789|19.60 ns/op|
|BenchmarkIsPrimeSerial/PrimeV2-1-4|48296462|22.95 ns/op|
|BenchmarkIsPrimeSerial/PrimeV3-1-4|75659934|15.84 ns/op|
|BenchmarkIsPrimeParallel-4|49888394|21.62 ns/op|


|JSON vs XML|Total|ns per op|mem per op|alloc per op|
|-|-|-|-|-|
|BenchmarkToJSON-4|332880|3214 ns/op|179 B/op|3 allocs/op|
|BenchmarkToXML-4|99578|11726 ns/op|4632 B/op|11 allocs/op|
|BenchmarkFromJSON-4|167470|7137 ns/op|312 B/op|7 allocs/op|
|BenchmarkFromXML-4|46717|25749 ns/op|1608 B/op|34 allocs/op|
|BenchmarkParToXML-4|190856|6110 ns/op|4632 B/op|11 allocs/op|
|BenchmarkParToJSON-4|716979|1718 ns/op|179 B/op|3 allocs/op|
|BenchmarkParFromJSON-4|336836|3518 ns/op|312 B/op|7 allocs/op|
|BenchmarkParFromXML-4|74744|15033 ns/op|1608 B/op|34 allocs/op|