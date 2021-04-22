package __1

import "testing"

func benchmarkQuickFind(length int, b *testing.B) {
	pairs := GeneratePairs(length)

	for n := 0; n < b.N; n++ {
		QuickFind(pairs)
	}
}

func BenchmarkQuickFind5 (b *testing.B) {
	benchmarkQuickFind(5, b)
}

func BenchmarkQuickFind100 (b *testing.B) {
	benchmarkQuickFind(100, b)
}

func BenchmarkQuickFind100000 (b *testing.B) {
	benchmarkQuickFind(100000, b)
}