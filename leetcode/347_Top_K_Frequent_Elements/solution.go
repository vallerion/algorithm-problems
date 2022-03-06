package _347_Top_K_Frequent_Elements

// [1,1,1,2,2,3,4,4] 2

// [1]=3
// [2]=2
// [3]=1
// [0,0,0,0,0,0]
// [0,3,2,1,0,0]

func topKFrequent(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if k == len(nums) {
		return nums
	}

	mp := make(map[int]int)
	arr := make([][]int, len(nums)+1)

	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 0)
	}

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	for k, v := range mp {
		arr[v] = append(arr[v], k)
	}

	result := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) > 0 {
			result = append(result, arr[i]...)
		}

		if len(result) == k {
			break
		}
	}

	return result
}
