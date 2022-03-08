package _1636_Sort_Array_by_Increasing_Frequency

import "sort"

// [2,3,1,3,2]

// [1]=1
// [2]=2
// [3]=2

// [1]=[1]
// [2]=[3,2]

// [1,3,3,2,2]

var mp map[int]int

func frequencySort(nums []int) []int {
	mp = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	sort.SliceStable(nums, func(i, j int) bool {
		if mp[nums[i]] == mp[nums[j]] {
			return nums[i] > nums[j]
		}

		return mp[nums[i]] < mp[nums[j]]
	})

	return nums
}
