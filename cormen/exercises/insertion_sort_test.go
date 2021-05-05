package exercises

import (
	"testing"
)

func BenchmarkInsertionSort100(b *testing.B) {
	benchmarkInsertionSort(100, b)
}

func BenchmarkInsertionSort100k(b *testing.B) {
	benchmarkInsertionSort(100000, b)
}

func benchmarkInsertionSort(length int, b *testing.B) {
	input := GenerateArray(length)

	for n := 0; n < b.N; n++ {
		InsertionSort(input)
	}
}
