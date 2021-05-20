package dsu

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimpleRankFind(t *testing.T) {
	InitRank()
	n := 100

	for i := 1; i <= n; i++ {
		makeSetRank(i)
	}

	for i := 1; i <= n; i++ {
		if i != findRank(i) {
			t.Errorf("%v must be in set #%v", i, i)
		}
	}
}

func TestSimpleRankUnion(t *testing.T) {
	InitRank()
	n := 8

	for i := 1; i <= n; i++ {
		makeSetRank(i)
	}

	unionRank(1, 2) // {1,2}
	unionRank(3, 4) // {3,4}
	unionRank(5, 6) // {5,6}
	unionRank(7, 8) // {7,8}

	assertEquals(2, 2, t)
	assertEquals(4, 4, t)
	assertEquals(6, 6, t)
	assertEquals(8, 8, t)

	assertEquals(findRank(2), findRank(1), t)
	assertEquals(findRank(4), findRank(3), t)
	assertEquals(findRank(6), findRank(5), t)
	assertEquals(findRank(8), findRank(7), t)

	assertNotEquals(findRank(8), findRank(1), t)
	assertNotEquals(findRank(6), findRank(4), t)

	unionRank(6, 4)
	assertEquals(findRank(6), findRank(4), t)
}

func assertEquals(v, expected int, t *testing.T) {
	if v != expected {
		t.Errorf("%v must be in set #%v", v, expected)
	}
}

func assertNotEquals(v, expected int, t *testing.T) {
	if v == expected {
		t.Errorf("%v must not be in set #%v", v, expected)
	}
}

func BenchmarkRank(b *testing.B) {
	InitRank()

	for i := 1; i < b.N; i++ {
		makeSetRank(i)
	}

	for i := 1; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		x, y := rand.Intn(i), rand.Intn(b.N)

		unionRank(x, y)
	}
}
