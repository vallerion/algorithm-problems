package _46_Permutations

// [1,2,3]

// [1 => [0=>1], 2 => [1=>1], 3 => [2=>1]]

// i=0,j=0 -> continue
// i=0,j=1 -> [1,2,3] -> [2,1,3] : [1 => [0=>1,1=>1], 2 => [1=>1,0=>1], 3 => [2=>1]] -> RECURSION([2,1,3])
// R i=0,j=1 -> [2,1,3] ->
// i=0,j=2 -> [1,2,3] -> [3,2,1] : [1 => [0=>1,1=>1,2=>1], 2 => [1=>1,0=>1], 3 => [2=>1,0=>1]] -> st = [[2,1,3],[3,2,1]]

// [[5,4,2,6],[5,4,2,6],[5,6,2,4],[5,6,2,4],[5,2,6,4],[5,2,6,4],[4,5,2,6],[4,5,2,6],[4,6,2,5],[4,6,2,5],[4,2,6,5],[4,2,6,5],[6,5,2,4],[6,5,2,4],[6,4,2,5],[6,4,2,5],[6,2,4,5],[6,2,4,5],[2,5,6,4],[2,5,6,4],[2,4,6,5],[2,4,6,5],[2,6,4,5],[2,6,4,5]]

// [[5,4,6,2],[5,4,2,6],[5,6,4,2],[5,6,2,4],[5,2,4,6],[5,2,6,4],[4,5,6,2],[4,5,2,6],[4,6,5,2],[4,6,2,5],[4,2,5,6],[4,2,6,5],[6,5,4,2],[6,5,2,4],[6,4,5,2],[6,4,2,5],[6,2,5,4],[6,2,4,5],[2,5,4,6],[2,5,6,4],[2,4,5,6],[2,4,6,5],[2,6,5,4],[2,6,4,5]]

func permute(nums []int) [][]int {

	res := make([][]int, 0)
	temp := make([]int, 0)

	perm(nums, temp, &res)

	return res
}

func perm(nums, temp []int, result *[][]int) {
	if len(temp) == len(nums) {
		tt := make([]int, len(temp))
		copy(tt, temp)
		*result = append(*result, tt)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(temp, nums[i]) {
			continue
		}

		temp = append(temp, nums[i])
		perm(nums, temp, result)
		temp = temp[:len(temp)-1]
	}
}

func contains(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}

	return false
}
