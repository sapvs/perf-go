# perf-go
Benchmarks of different approaches for common problem in go


|Case | Total | ns per op |
|-----|-------|-----------|
|BenchmarkCounterAtomic-4|69815256| 18.02 ns/op|
|BenchmarkCounterMutex-4|24769442| 46.93 ns/op|
|BenchmarkFactorialSerial/FactorialSerial-20-4 |113081472| 10.46 ns/op|
|BenchmarkFactorialSerial/FactorialRecursive-20-4|27576867| 40.32 ns/op|
|BenchmarkFactorialParallel/FactorialSerial-20-4 |276191018|4.326 ns/op|
|BenchmarkFactorialParallel/FactorialRecursive-20-4|72123966| 15.11 ns/op|
|BenchmarkFibonnaciSerial/FibonacciSumClassic-20-4 | 31941| 37669 ns/op|
|BenchmarkFibonnaciSerial/FibonacciSumDP-20-4|59457067| 18.42 ns/op|
|BenchmarkFibbonacciParallel-4|75055| 15622 ns/op|
|BenchmarkFizzBuzzSer/v1-1-4 |290979|3659 ns/op|
|BenchmarkFizzBuzzSer/v2-1-4 |250216|5175 ns/op|
|BenchmarkFizzBuzzSer/v3-1-4 |262671|4178 ns/op|
|BenchmarkFizzBuzzPar/v1-1-4 |526555|1991 ns/op|
|BenchmarkFizzBuzzPar/v2-1-4 |480120|2218 ns/op|
|BenchmarkFizzBuzzPar/v3-1-4 |494432|2205 ns/op|
|BenchmarkIsPrimeSerial/PrimeV1-1-4|51702789| 19.60 ns/op|
|BenchmarkIsPrimeSerial/PrimeV2-1-4|48296462| 22.95 ns/op|
|BenchmarkIsPrimeSerial/PrimeV3-1-4|75659934| 15.84 ns/op|
|BenchmarkIsPrimeParallel-4|49888394| 21.62 ns/op|



