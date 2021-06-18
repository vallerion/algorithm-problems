package _448_Find_All_Numbers_Disappeared_in_an_Array

func findDisappearedNumbers(nums []int) []int {
	mp := make(map[int]int)
	maxNum := len(nums)
	miss := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = 1
	}

	for i := maxNum; i > 0; i-- {
		if _, ok := mp[i]; ok == false {
			miss = append(miss, i)
		}
	}

	return miss
}
