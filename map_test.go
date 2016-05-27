package capable

import (
	"testing"
)

func mapLimit() {
	a := make(map[int]int, capacity)
	for i := 0; i < capacity; i++ {
		a[i] = i
	}
}

func mapNo() {
	a := make(map[int]int)
	for i := 0; i < capacity; i++ {
		a[i] = i
	}
}

func BenchmarkTest_mapLimit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapLimit()
	}
}

func BenchmarkTest_mapNo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapNo()
	}
}
