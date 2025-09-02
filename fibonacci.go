package perf

// FibonacciSumV1 non DP solution
func FibonacciSumV1(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciSumV1(n-1) + FibonacciSumV1(n-2)
}

var fibmap map[int]int = make(map[int]int)

// FibonacciSumV2 DP solution
func FibonacciSumV2(n int) int {
	if n < 2 {
		return n
	}
	_, ok := fibmap[n]
	if !ok {
		// log.Printf("not found for %d\n", n)
		fibmap[n] = FibonacciSumV2(n-1) + FibonacciSumV2(n-2)
	}
	// log.Printf("found for %d\n", n)
	return fibmap[n]
}

// FibonacciSumV3 Bottom up
func FibonacciSumV3(n int) int {
	if n < 2 {
		return n
	}

	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n+1; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
} // FibonacciSumV4 space optimized
func FibonacciSumV4(n int) int {
	if n < 2 {
		return n
	}

	prev := 0
	curr := 1

	for i := 2; i < n+1; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}
