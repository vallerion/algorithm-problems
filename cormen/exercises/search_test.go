package exercises

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkLinearSearch100(b *testing.B) {
	input := GenerateSortedArray(100)

	for n := 0; n < b.N; n++ {
		rand.Seed(time.Now().UnixNano())
		item := input[rand.Intn(len(input)-1)]
		LinearSearch(input, item)
	}
}

func BenchmarkLinearSearch100k(b *testing.B) {
	input := GenerateSortedArray(100000)

	for n := 0; n < b.N; n++ {
		rand.Seed(time.Now().UnixNano())
		item := input[rand.Intn(len(input)-1)]
		LinearSearch(input, item)
	}
}

func BenchmarkBinarySearch100(b *testing.B) {
	input := GenerateSortedArray(100)

	for n := 0; n < b.N; n++ {
		rand.Seed(time.Now().UnixNano())
		item := input[rand.Intn(len(input)-1)]
		BinarySearch(input, item)
	}
}

func BenchmarkBinarySearch100k(b *testing.B) {
	input := GenerateSortedArray(100000)

	for n := 0; n < b.N; n++ {
		rand.Seed(time.Now().UnixNano())
		item := input[rand.Intn(len(input)-1)]
		BinarySearch(input, item)
	}
}
