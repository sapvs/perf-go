package perf_test

import (
	"fmt"
	"testing"

	"github.com/vsapan/perf-go"
)

var testfun = []struct {
	name string
	f    func(int) int
}{
	{"FibonacciSumV1", perf.FibonacciSumV1},
	{"FibonacciSumV2", perf.FibonacciSumV2},
	{"FibonacciSumV3", perf.FibonacciSumV3},
	{"FibonacciSumV4", perf.FibonacciSumV4},
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

// Benchmarks
var inp int = 30
var output int

func BenchmarkFibonnaciSerial(b *testing.B) {
	for _, fn := range testfun {
		b.Run(fmt.Sprintf("%s-%d", fn.name, inp), func(b *testing.B) {
			for b.Loop() {
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
