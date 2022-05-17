package _1202_Smallest_String_With_Swaps

import (
	"fmt"
	"sort"
)

type DisjointSet struct {
	values []int
}

func (ds *DisjointSet) find(i int) int {
	if ds.values[i] == i {
		return i
	}
	ds.values[i] = ds.find(ds.values[i])
	return ds.values[i]
}

func (ds *DisjointSet) union(a, b int) {
	rootA, rootB := ds.find(a), ds.find(b)

	if rootA != rootB {
		ds.values[rootB] = rootA
	}
}

func Constructor(size int) *DisjointSet {
	values := make([]int, size)
	for i := 0; i < size; i++ {
		values[i] = i
	}

	return &DisjointSet{values}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(s) == 1 {
		return s
	}

	ds := Constructor(len(s))

	for i := 0; i < len(pairs); i++ {
		ds.union(pairs[i][0], pairs[i][1])
	}

	rootMap := make(map[int][]int)
	for i := 0; i < len(ds.values); i++ {
		root := ds.find(i)
		if _, ok := rootMap[root]; ok {
			rootMap[root] = append(rootMap[root], i)
		} else {
			rootMap[root] = []int{i}
		}
	}

	result := make([]byte, len(s))
	for _, indices := range rootMap {

		chars := make([]byte, 0)
		for i := 0; i < len(indices); i++ {
			chars = append(chars, s[indices[i]])
		}
		sort.SliceStable(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		fmt.Println(result)
		for i := 0; i < len(indices); i++ {
			result[indices[i]] = chars[i]
		}

	}

	return string(result)
}
