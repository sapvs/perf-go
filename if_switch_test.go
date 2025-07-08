package perf

import "testing"

var fux = []struct {
	name string
	f    func(int) bool
}{
	{"If", MapWithIf},
	{"Elseif", MapWithElseif},
	{"Switch", MapWithSwitch},
}

var result bool

func Benchmark(b *testing.B) {
	for _, f := range fux {
		b.Run(f.name, func(b *testing.B) {
			for b.Loop() {
				result = f.f(100_000)
			}
		})
	}
}
