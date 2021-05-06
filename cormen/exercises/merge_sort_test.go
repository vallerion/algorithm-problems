package exercises

import (
	"testing"
)

func BenchmarkMergeSort100(b *testing.B) {
	WrapperBenchmark(100, b, MergeSort)
}

func BenchmarkMergeSort100k(b *testing.B) {
	WrapperBenchmark(100000, b, MergeSort)
}

func TestMergeSort100(t *testing.T) {
	WrapperTest(100, t, MergeSort)
}

func TestMergeSort100k(t *testing.T) {
	WrapperTest(100000, t, MergeSort)
}
