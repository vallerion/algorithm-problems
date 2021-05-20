package dsu

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimpleNaiveFind(t *testing.T) {
	InitNaive()
	n := 100

	for i := 1; i <= n; i++ {
		makeSetNaive(i)
	}

	for i := 1; i <= n; i++ {
		if i != findNaive(i) {
			t.Errorf("%v must be in set #%v", i, i)
		}
	}
}

func TestSimpleNaiveUnion(t *testing.T) {
	InitNaive()
	n := 8

	for i := 1; i <= n; i++ {
		makeSetNaive(i)
	}

	unionNaive(1, 2) // {1,2}
	unionNaive(3, 4) // {3,4}
	unionNaive(5, 6) // {5,6}
	unionNaive(7, 8) // {7,8}

	checkUnionNaive(2, 1, t)
	checkUnionNaive(4, 3, t)
	checkUnionNaive(6, 5, t)
	checkUnionNaive(8, 7, t)
}

func checkUnionNaive(v, expected int, t *testing.T) {
	if findNaive(v) != expected {
		t.Errorf("%v must be in set #%v", v, expected)
	}
}

func BenchmarkNaive(b *testing.B) {
	InitNaive()

	for i := 1; i < b.N; i++ {
		makeSetNaive(i)
	}

	for i := 1; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		x, y := rand.Intn(i), rand.Intn(b.N)

		unionNaive(x, y)
	}
}
