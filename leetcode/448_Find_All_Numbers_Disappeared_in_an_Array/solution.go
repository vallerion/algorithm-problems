package _448_Find_All_Numbers_Disappeared_in_an_Array

// 4,3,2,7,8,2,3,1
// 7,3,2,-1,8,2,3,1
// 3,3,2,-1,8,2,-1,1
// 3,-1,-1,-1,8,2,-1,1

// 1,1

func findDisappearedNumbers(nums []int) []int {
	i := 0

	for i < len(nums) {
		if nums[i] == -1 || nums[nums[i]-1] == -1 {
			i++
			continue
		}

		if i != nums[i]-1 {
			temp := nums[nums[i]-1]
			nums[nums[i]-1] = -1
			nums[i] = temp
		} else {
			nums[nums[i]-1] = -1
		}
	}

	res := make([]int, 0)
	for j := 0; j < len(nums); j++ {
		if nums[j] != -1 {
			res = append(res, j+1)
		}
	}

	return res
}

func findDisappearedNumbersWithExtraSpace(nums []int) []int {
	hh := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		hh[nums[i]]++
	}

	res := make([]int, 0)

	for i := 1; i < len(hh); i++ {
		if hh[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}
