package _56_Merge_Intervals

import "sort"

// [[1,3],[2,6],[8,10],[15,18]]

// results = [ {1,3} ]
// interval[i][0] <= interval [i-1][1] -> { min(1,2), max(3,6) } -> {1,6}

type intervalsType [][]int

func (in intervalsType) Len() int {
	return len(in)
}

func (in intervalsType) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}

func (in intervalsType) Swap(i, j int) {
	temp := in[i]
	in[i] = in[j]
	in[j] = temp
}

func merge(intervals [][]int) [][]int {
	intvls := intervalsType(intervals)
	result := make([][]int, 0)

	sort.Sort(intvls)

	result = append(result, intvls[0])

	for i := 1; i < len(intvls); i++ {
		if intvls[i][0] <= result[len(result)-1][1] {
			result[len(result)-1][0] = min(intvls[i][0], result[len(result)-1][0])
			result[len(result)-1][1] = max(intvls[i][1], result[len(result)-1][1])
		} else {
			result = append(result, intvls[i])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
