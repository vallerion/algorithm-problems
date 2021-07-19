package _448_Find_All_Numbers_Disappeared_in_an_Array

func findDisappearedNumbers(nums []int) []int {
	hp := make(map[int]int)
	n := len(nums)
	res := make([]int, 0)

	for i := 0; i < n; i++ {
		hp[nums[i]]++
	}

	for i := 0; i < n; i++ {
		_, ok := hp[i+1]
		if ok == false {
			res = append(res, i+1)
		}
	}

	return res
}

func findDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		ind := abs(nums[i]) - 1
		if nums[ind] > 0 {
			nums[ind] = -nums[ind]
		}
	}

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[j] = i+1
			j++
		}
	}

	return nums[:j]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
