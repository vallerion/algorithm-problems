package _493_Reverse_Pairs

// nums[i] > nums[j] * 2
// [1,3,2,3,1]

// [3] [1]

var merged []int

func reversePairs(nums []int) int {
	merged = make([]int, len(nums))

	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	middle := start + (end-start)/2
	total := mergeSort(nums, start, middle) + mergeSort(nums, middle+1, end)

	for i := start; i <= middle; i++ {
		for j := middle + 1; j <= end && nums[i] > nums[j]*2; j++ {
			total++
		}
	}

	merge(nums, start, middle, end)

	return total
}

func merge(nums []int, start, middle, end int) {
	// left  - start to middle
	// right - middle+1 to end
	left, right := start, middle+1

	mergedIndex := 0

	for left <= middle && right <= end {
		if nums[left] > nums[right] {
			merged[mergedIndex]= nums[right]
			mergedIndex++
			//merged = append(merged, nums[right])
			right++
		} else {
			merged[mergedIndex]= nums[left]
			mergedIndex++
			//merged = append(merged, nums[left])
			left++
		}
	}

	for left <= middle {
		merged[mergedIndex]= nums[left]
		mergedIndex++
		//merged = append(merged, nums[left])
		left++
	}

	for right <= end {
		merged[mergedIndex]= nums[right]
		mergedIndex++
		//merged = append(merged, nums[right])
		right++
	}

	for i := start; i <= end; i++ {
		nums[i] = merged[i-start]
	}
}
