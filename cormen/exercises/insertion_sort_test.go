package exercises

import (
	"testing"
)

func BenchmarkInsertionSort100(b *testing.B) {
	WrapperBenchmark(100, b, InsertionSort)
}

func BenchmarkInsertionSort100k(b *testing.B) {
	WrapperBenchmark(100000, b, InsertionSort)
}

func BenchmarkInsertionSortBS100(b *testing.B) {
	WrapperBenchmark(100, b, InsertionSortBS)
}

func BenchmarkInsertionSortBS100k(b *testing.B) {
	WrapperBenchmark(100000, b, InsertionSortBS)
}

func TestInsertionSort100(t *testing.T) {
	WrapperTest(100, t, InsertionSort)
}

func TestInsertionSort100k(t *testing.T) {
	WrapperTest(100000, t, InsertionSort)
}

func TestInsertionSortBS100(t *testing.T) {
	WrapperTest(100, t, InsertionSortBS)
}

func TestInsertionSortBS100k(t *testing.T) {
	WrapperTest(100000, t, InsertionSortBS)
}
