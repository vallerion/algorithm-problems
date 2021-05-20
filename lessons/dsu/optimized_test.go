package dsu

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimplePCRankFind(t *testing.T) {
	InitPCRank()
	n := 100

	for i := 1; i <= n; i++ {
		makeSetPCRank(i)
	}

	for i := 1; i <= n; i++ {
		if i != findPCRank(i) {
			t.Errorf("%v must be in set #%v", i, i)
		}
	}
}

func TestSimplePCRankUnion(t *testing.T) {
	InitPCRank()
	n := 8

	for i := 1; i <= n; i++ {
		makeSetPCRank(i)
	}

	unionPCRank(1, 2) // {1,2}
	unionPCRank(3, 4) // {3,4}
	unionPCRank(5, 6) // {5,6}
	unionPCRank(7, 8) // {7,8}

	checkUnionPCRank(1, 2, t)
	checkUnionPCRank(2, 2, t)
	checkUnionPCRank(4, 4, t)
	checkUnionPCRank(3, 4, t)
	checkUnionPCRank(6, 6, t)
	checkUnionPCRank(8, 8, t)
}

func checkUnionPCRank(v, expected int, t *testing.T) {
	if findPCRank(v) != expected {
		t.Errorf("%v must be in set #%v", v, expected)
	}
}

func BenchmarkPCRank(b *testing.B) {
	InitPCRank()

	for i := 1; i < b.N; i++ {
		makeSetPCRank(i)
	}
	rand.Seed(time.Now().UnixNano())
	x, y := rand.Intn(b.N/2), rand.Intn(b.N)

	for i := 1; i < b.N; i++ {
		unionPCRank(x, y)
	}
}
