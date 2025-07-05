package perf_test

import (
	"fmt"
	"testing"

	"github.com/vsapan/perf-go"
)

var isPrimeFuncs = []struct {
	name string
	f    func(int) bool
}{
	// {name: "PrimeV1", f: perf.IsPrimeV1},
	// {name: "PrimeV2", f: perf.IsPrimeV2},
	// {name: "PrimeV3", f: perf.IsPrimeV3},
}

var primeUnitTests = []struct {
	in   int
	want bool
}{
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{9789, false},
}

func ExampleIsPrimeV1() {
	isPrime := perf.IsPrimeV1(11)
	fmt.Printf("11 is prime = %v", isPrime)
	//Output: 11 is prime = true
}
func ExampleIsPrimeV2() {
	isPrime := perf.IsPrimeV1(11)
	fmt.Printf("11 is prime = %v", isPrime)
	//Output: 11 is prime = true
}
func ExampleIsPrimeV3() {
	isPrime := perf.IsPrimeV1(11)
	fmt.Printf("11 is prime = %v", isPrime)
	//Output: 11 is prime = true
}

func TestIsPrime(t *testing.T) {
	t.Parallel()
	for _, test := range isPrimeFuncs {
		for _, tt := range primeUnitTests {
			t.Run(fmt.Sprintf("%s-%d", test.name, tt.in), func(t *testing.T) {
				if got := test.f(tt.in); got != tt.want {
					t.Errorf("%s = %v, want %v", test.name, got, tt.want)
				}
			})
		}
	}
}

var isPrimeResult bool
var primeInput int = 999

func BenchmarkIsPrimeSerial(b *testing.B) {
	for _, test := range isPrimeFuncs {
		b.Run(fmt.Sprintf("%s-%d", test.name, b.N), func(b *testing.B) {
			for b.Loop() {
				isPrimeResult = test.f(primeInput)
			}
		})
	}
}

func BenchmarkIsPrimeParallel(b *testing.B) {
	for _, test := range isPrimeFuncs {
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				test.f(primeInput)
			}
		})
	}
}
