package exercises

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSelectionSort100(b *testing.B) {
	benchmarkSelectionSort(100, b)
}

func BenchmarkSelectionSort100k(b *testing.B) {
	benchmarkSelectionSort(100000, b)
}

func benchmarkSelectionSort(length int, b *testing.B) {
	input := GenerateArray(length)

	for n := 0; n < b.N; n++ {
		SelectionSort(input)
	}
}

func GenerateArray(length int) []int {
	var result []int

	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		result = append(result, rand.Intn(length))
	}

	return result
}

