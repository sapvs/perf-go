package perf

//FibonacciSumV1 non DP solution
func FibonacciSumV1(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciSumV1(n-1) + FibonacciSumV1(n-2)
}

var fibmap map[int]int = make(map[int]int)

//FibonacciSumV2 DP solution
func FibonacciSumV2(n int) int {
	if n < 2 {
		return n
	}
	_, ok := fibmap[n]
	if !ok {
		fibmap[n] = FibonacciSumV2(n-1) + FibonacciSumV2(n-2)
	}
	return fibmap[n]
}
