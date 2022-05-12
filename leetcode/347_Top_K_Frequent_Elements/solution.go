package _347_Top_K_Frequent_Elements

// [1,1,1,2,2,3]

// 1 -> 3
// 2 -> 2
// 3 -> 1

// 1 -> [3]
// 2 -> [2]
// 3 -> [1]

// [1,2]

// O(n)
func topKFrequent(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	hm := make(map[int]int)
	bucket := make([][]int, len(nums))
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		hm[nums[i]]++
	}

	for k, v := range hm {
		bucket[v-1] = append(bucket[v-1], k)
	}

	i := len(bucket) - 1
	for len(result) < k {
		result = append(result, bucket[i]...)
		i--
	}

	return result
}
