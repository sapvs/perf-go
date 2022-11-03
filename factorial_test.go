package perf_test

import (
	"fmt"
	"testing"

	"github.com/s3vt/perf-go"
)

var factorialfuncs = []struct {
	name string
	funx func(int) int
}{
	{name: "FactorialSerial", funx: perf.FactorialSerial},
	{name: "FactorialRecursive", funx: perf.FactorialRecursive},
}

var factorialUnitTests = []struct {
	name string
	n    int
	want int
}{
	{"1", 1, 1},
	{"2", 2, 2},
	{"3", 3, 6},
	{"5", 5, 120},
	{"20", 20, 2432902008176640000},
}

func ExampleFactorialSerial() {
	factorial := perf.FactorialSerial(5)
	fmt.Printf("Factorial of 5 is %d", factorial)
	//Output: Factorial of 5 is 120
}

func ExampleFactorialRecursive() {
	number := 5
	factorial := perf.FactorialRecursive(number)
	fmt.Printf("Factorial of %d is %d", number, factorial)
	//Output: Factorial of 5 is 120
}

func TestFactorial(t *testing.T) {
	t.Parallel()
	for _, fun := range factorialfuncs {
		for _, tt := range factorialUnitTests {
			t.Run(fmt.Sprintf("%s-%s", fun.name, tt.name), func(t *testing.T) {
				if got := fun.funx(tt.n); got != tt.want {
					t.Errorf("%s(%d) = %v, want %v", fun.name, tt.n, got, tt.want)
				}
			})
		}
	}
}

// Benchmarks
var out int
var inputArg int = 20

func BenchmarkFactorialSerial(b *testing.B) {
	for _, fun := range factorialfuncs {
		b.Run(fmt.Sprintf("%s-%d", fun.name, inputArg), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				out = fun.funx(inputArg)
			}
		})
	}
}

func BenchmarkFactorialParallel(b *testing.B) {
	for _, fun := range factorialfuncs {
		b.Run(fmt.Sprintf("%s-%d", fun.name, inputArg), func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					fun.funx(inputArg)
				}
			})
		})
	}
}

var fact int

func FuzzFactorial(f *testing.F) {
	for _, tc := range []int{1, 2, 3, 4, 5} {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, a int) {
		fact = perf.FactorialSerial(a)
	})

}
