// Different versions of fizzbuzz implementation for benchmarking
package perf_test

import (
	"fmt"
	"testing"

	"github.com/s3vt/perf-go"
)

var result []string

var fizzBuzzFuncs = []struct {
	name string
	f    func(int) []string
}{
	{name: "v1", f: perf.FizzbuzzV1},
	{name: "v2", f: perf.FizzbuzzV2},
	{name: "v3", f: perf.FizzbuzzV3},
}

const COUNT = 50

func BenchmarkFizzBuzzSer(b *testing.B) {
	for _, ffn := range fizzBuzzFuncs {
		b.Run(fmt.Sprintf("%s-%d", ffn.name, b.N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = ffn.f(COUNT)
			}
		})
	}
}

func BenchmarkFizzBuzzPar(b *testing.B) {
	for _, ffn := range fizzBuzzFuncs {
		b.Run(fmt.Sprintf("%s-%d", ffn.name, b.N), func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					ffn.f(COUNT)
				}
			})
		})
	}
}
