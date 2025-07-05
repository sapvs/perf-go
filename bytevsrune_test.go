package perf_test

import "testing"

var str1, str2 string

func BenchmarkByteToString(b *testing.B) {
	for b.Loop() {
		str1 = string([]byte{'a', 'b'})
	}
}

func BenchmarkRuneToString(b *testing.B) {
	for b.Loop() {
		str2 = string([]rune{'a', 'b'})
	}
}
