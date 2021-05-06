package exercises

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

type SortAlgo func([]int)[]int

func WrapperTest(length int, t *testing.T, fn SortAlgo) {
	original := GenerateArray(length)

	input := make([]int, len(original))
	copy(input, original)

	got := fn(input)

	want := make([]int, len(original))
	copy(want, original)
	sort.Ints(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func WrapperBenchmark(length int, b *testing.B, fn SortAlgo) {
	for n := 0; n < b.N; n++ {
		fn(GenerateArray(length))
	}
}

func GenerateArray(length int) []int {
	var result []int

	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		result = append(result, rand.Intn(length * 2))
	}

	return result
}

func GenerateSortedArray(length int) []int {
	nums := GenerateArray(length)
	sort.Ints(nums)

	return nums
}