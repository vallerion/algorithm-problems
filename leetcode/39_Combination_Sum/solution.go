package _39_Combination_Sum

func combinationSum(candidates []int, target int) [][]int {

	temp := make([]int, 0)
	res := make([][]int, 0)

	backtracking(candidates, temp, &res, 0, 0, target)

	return res
}

func backtracking(candidates, temp []int, result *[][]int, inx, sum, target int) {
	if sum == target {
		tt := make([]int, len(temp))
		copy(tt, temp)
		*result = append(*result, tt)
		return
	}

	for i := inx; i < len(candidates); i++ {
		if candidates[i]+sum > target {
			continue
		}

		temp = append(temp, candidates[i])
		sum += candidates[i]
		backtracking(candidates, temp, result, i, sum, target)
		sum -= candidates[i]
		temp = temp[:len(temp)-1]
	}
}
