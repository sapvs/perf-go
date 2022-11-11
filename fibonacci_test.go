package perf_test

import (
	"fmt"
	"testing"

	"github.com/s3vt/perf-go"
)

var testfun = []struct {
	name string
	f    func(int) int
}{
	{"FibonacciSumClassic", perf.FibonacciSumV1},
	{"FibonacciSumDP", perf.FibonacciSumV2},
}

func TestFibonacciSum(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"10", args{10}, 55},
		{"20", args{20}, 6765},
		{"30", args{30}, 832040},
		{"40", args{40}, 102334155},
	}

	for _, fn := range testfun {
		for _, tt := range tests {
			t.Run(fn.name+"-"+tt.name, func(t *testing.T) {
				if got := fn.f(tt.args.n); got != tt.want {
					t.Errorf("FibonacciSum-%s() = %v, want %v", tt.name, got, tt.want)
				}
			})
		}
	}
}

//Benchmarks
var inp int = 20
var output int

func BenchmarkFibonnaciSerial(b *testing.B) {
	for _, fn := range testfun {
		b.Run(fmt.Sprintf("%s-%d", fn.name, inp), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				output = fn.f(inp)
			}
		})
	}
}

func BenchmarkFibbonacciParallel(b *testing.B) {
	for _, fn := range testfun {
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				fn.f(inp)
			}
		})
	}
}
