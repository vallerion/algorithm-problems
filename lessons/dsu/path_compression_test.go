package dsu

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimplePCFind(t *testing.T) {
	InitPC()
	n := 100

	for i := 1; i <= n; i++ {
		makeSetPC(i)
	}

	for i := 1; i <= n; i++ {
		if i != findPC(i) {
			t.Errorf("%v must be in set #%v", i, i)
		}
	}
}

func TestSimplePCUnion(t *testing.T) {
	InitPC()
	n := 8

	for i := 1; i <= n; i++ {
		makeSetPC(i)
	}

	unionPC(1, 2) // {1,2}
	unionPC(3, 4) // {3,4}
	unionPC(5, 6) // {5,6}
	unionPC(7, 8) // {7,8}

	checkUnionPC(1, 1, t)
	checkUnionPC(2, 1, t)
	checkUnionPC(4, 3, t)
	checkUnionPC(6, 5, t)
	checkUnionPC(8, 7, t)
}

func checkUnionPC(v, expected int, t *testing.T) {
	if findPC(v) != expected {
		t.Errorf("%v must be in set #%v", v, expected)
	}
}

func BenchmarkPC(b *testing.B) {
	InitPC()

	for i := 1; i < b.N; i++ {
		makeSetPC(i)
	}

	for i := 1; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		x, y := rand.Intn(i), rand.Intn(b.N)

		unionPC(x, y)
	}
}
