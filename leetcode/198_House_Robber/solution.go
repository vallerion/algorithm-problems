package _198_House_Robber

// [1,2,3,1]

// [1] -> max=1
// [1,2] -> max=2
// [1,3,1] -> 1+1 or 3 -> max=3
//

// [2,1,1,2]

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	return robMax(&nums, &memo, len(nums)-1)
}

func robMax(nums, memo *[]int, i int) int {
	if i < 0 {
		return 0
	}

	if (*memo)[i] >= 0 {
		return (*memo)[i]
	}

	(*memo)[i] = max(robMax(nums, memo, i-2)+(*nums)[i], robMax(nums, memo, i-1))

	return (*memo)[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
