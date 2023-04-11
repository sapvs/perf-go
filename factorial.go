package perf

//FactorialSerial Serial solution
func FactorialSerial(n int) int {
	fact := 1
	if n >= 0 {
		for i := 1; i <= n; i++ {
			fact *= i
		}
	}
	return fact
}

//FactorialRecursive recursive solution
func FactorialRecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}
