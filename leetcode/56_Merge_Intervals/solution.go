package _56_Merge_Intervals

import (
	"sort"
)

type Intervals [][]int

func (i Intervals) Len() int           { return len(i) }
func (i Intervals) Less(a, b int) bool { return i[a][0] < i[b][0] }
func (i Intervals) Swap(a, b int)      { i[a], i[b] = i[b], i[a] }

func merge(intervals [][]int) [][]int {
	data := Intervals(intervals)
	sort.Sort(data)

	merged := make([][]int, 0)

	for i := 0; i < data.Len(); i++ {
		if len(merged) == 0 {
			merged = append(merged, data[i])
			continue
		}

		if merged[len(merged)-1][1] >= data[i][0] {
			merged[len(merged)-1][1] = max(data[i][1], merged[len(merged)-1][1])
		} else {
			merged = append(merged, data[i])
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
