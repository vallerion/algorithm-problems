package _347_Top_K_Frequent_Elements

func topKFrequent(nums []int, k int) []int {
	hp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hp[nums[i]]++
	}

	bucket := make([][]int, len(nums)+1)

	for key, value := range hp {
		if bucket[value] == nil {
			bucket[value] = make([]int, 0)
		}

		bucket[value] = append(bucket[value], key)
	}

	result := make([]int, 0)

	for i := len(nums); i >= 0 && len(result) < k; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		result = append(result, bucket[i]...)
	}

	return result
}
