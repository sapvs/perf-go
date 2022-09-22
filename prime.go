package perf

import (
	"math"
)

// IsPrimeV1 worst
func IsPrimeV1(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeV2 better
func IsPrimeV2(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeV3 best
func IsPrimeV3(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}

	if num > 2 && num%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(num)))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
