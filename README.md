# perf-go
Benchmarks of different approaches for common problem in go

## Run tests
```
make
```

## Results (On my PC ;) )

### Counter problem using [sync.atomic](https://pkg.go.dev/sync/atomic) versus [sync.mutex](https://pkg.go.dev/sync#Mutex)
|Counter|Total|ns/op|
|-|-:|-:|
|BenchmarkCounterAtomic-4|69815256|18.02|
|BenchmarkCounterMutex-4|24769442|46.93|

### Factorial recursive vs serial
|Factorial|Total|ns/op|
|-|-:|-:|
|BenchmarkFactorialSerial/FactorialSerial-20-4|113081472|10.46|
|BenchmarkFactorialSerial/FactorialRecursive-20-4|27576867|40.32|
|BenchmarkFactorialParallel/FactorialSerial-20-4|276191018|4.326|
|BenchmarkFactorialParallel/FactorialRecursive-20-4|72123966|15.11|

### Fibonacci Classic vs DP
|Fibbonacci|Total|ns/op|
|-|-:|-:|
|BenchmarkFibonnaciSerial/FibonacciSumClassic-20-4|31941|37669|
|BenchmarkFibonnaciSerial/FibonacciSumDP-20-4|59457067|18.42|
|BenchmarkFibbonacciParallel-4|75055|15622|

### Fizzbuzz three approaches
|FizzBuzz|Total|ns/op|
|-|-:|-:|
|BenchmarkFizzBuzzSer/v1-1-4|290979|3659|
|BenchmarkFizzBuzzSer/v2-1-4|250216|5175|
|BenchmarkFizzBuzzSer/v3-1-4|262671|4178|
|BenchmarkFizzBuzzPar/v1-1-4|526555|1991|
|BenchmarkFizzBuzzPar/v2-1-4|480120|2218|
|BenchmarkFizzBuzzPar/v3-1-4|494432|2205|

### Check if number is prime, Three ways worst to best
|Is Prime|Total|ns/op|
|-|-:|-:|
|BenchmarkIsPrimeSerial/PrimeV1-1-4|51702789|19.60|
|BenchmarkIsPrimeSerial/PrimeV2-1-4|48296462|22.95|
|BenchmarkIsPrimeSerial/PrimeV3-1-4|75659934|15.84|
|BenchmarkIsPrimeParallel-4|49888394|21.62|


### JSON vs XML Vs ProtoBuf
|JSON vs XML|Total|ns/op|mem/op|alloc/op|
|-|-:|-:|-:|-:|
|BenchmarkToJSON-4|332880|3214|179|3|
|BenchmarkToXML-4|99578|11726|4632|11|
|BenchmarkFromJSON-4|167470|7137|312|7|
|BenchmarkFromXML-4|46717|25749|1608|34|
|BenchmarkParToXML-4|190856|6110|4632|11|
|BenchmarkParToJSON-4|716979|1718|179|3|
|BenchmarkParFromJSON-4|336836|3518|312|7|
|BenchmarkParFromXML-4|74744|15033|1608|34|
|BenchmarkToProto-4|6285468|175.3|80|2|
|BenchmarkFromProto-4|5491707|211.5|72|2|
|BenchmarkParToProto-4|10795147|100.6|80|2|
|BenchmarkParFromProto-4|9810102|114.5|72|2|
