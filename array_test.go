package capable

import (
	"testing"
)

func arrayArgTest(a [capacity]int) {
	for _ = range a {

	}
}
func sliceArgTest(a []int) {
	for _ = range a {

	}
}

func BenchmarkTest_array(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = array()
	}
}

func BenchmarkTest_sliceLimit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sliceLimit()
	}
}

func BenchmarkTest_slice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slice()
	}
}

func BenchmarkTest_array_arg(b *testing.B) {
	a := array()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrayArgTest(a)
	}
}

func BenchmarkTest_sliceLimit_arg(b *testing.B) {
	a := sliceLimit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceArgTest(a)
	}
}

func BenchmarkTest_slice_arg(b *testing.B) {
	a := slice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceArgTest(a)
	}
}
