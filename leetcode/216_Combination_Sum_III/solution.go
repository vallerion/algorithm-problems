package _216_Combination_Sum_III

func combinationSum3(k int, n int) [][]int {
	temp := make([]int, 0)
	res := make([][]int, 0)
	backtracking(1, 0, k, n, temp, &res)
	return res
}

func backtracking(start, sum, k, n int, temp []int, result *[][]int) {
	if len(temp) >= k {
		if sum != n || len(temp) > k {
			return
		}
		tt := make([]int, len(temp))
		copy(tt, temp)
		*result = append(*result, tt)
		return
	}

	for i := start; i < 10; i++ {
		sum += i
		temp = append(temp, i)
		backtracking(i+1, sum, k, n, temp, result)
		sum -= i
		temp = temp[:len(temp)-1]
	}
}
