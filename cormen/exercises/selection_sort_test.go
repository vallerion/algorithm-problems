package exercises

import (
	"testing"
)

func BenchmarkSelectionSort100(b *testing.B) {
	WrapperBenchmark(100, b, SelectionSort)
}

func BenchmarkSelectionSort100k(b *testing.B) {
	WrapperBenchmark(100000, b, SelectionSort)
}

func TestSelectionSort100(t *testing.T) {
	WrapperTest(100, t, SelectionSort)
}

func TestSelectionSort100k(t *testing.T) {
	WrapperTest(100000, t, SelectionSort)
}
