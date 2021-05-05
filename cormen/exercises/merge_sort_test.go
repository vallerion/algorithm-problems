package exercises

import (
	"testing"
)

func BenchmarkMergeSort100(b *testing.B) {
	benchmarkMergeSort(100, b)
}

func BenchmarkMergeSort100k(b *testing.B) {
	benchmarkMergeSort(100000, b)
}

func benchmarkMergeSort(length int, b *testing.B) {
	input := GenerateArray(length)

	for n := 0; n < b.N; n++ {
		MergeSort(input)
	}
}
