package exercises

// [5,3,43,4,2]
// 3: [5,5,43,4,2] -> [3,5,43,4,2]
// 43: do nothing -> [3,5,43,4,2]
// 4: [3,5,43,43,2] -> [3,5,5,43,2] ->  [3,4,5,43,2]
// 2: [3,4,5,43,43] -> [3,4,5,5,43] -> [3,4,4,5,43] -> [3,3,4,5,43] -> [2,3,4,5,43]

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		j := i - 1
		for j >= 0 && curr < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = curr
	}

	return nums
}

// [5,3,43,4,2]
// bs=0, j=0, j>=bs, A[j+1]=A[j] ... A[j+1]=curr / [5,5,43,4,2] -> [3,5,43,4,2]
// bs=1, j=1, j>=bs, A[j+1]=A[j] ... A[j+1]=curr /

// [1,4,4,4,7]
// curr=4, bs=0, j=0,

// InsertionSortBS - insertion sort with binary search
func InsertionSortBS(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		j := i - 1
		bsr := bs(nums, 0, j, curr)

		for j >= bsr {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = curr
	}

	return nums
}

func bs(nums []int, start, end, value int) int {
	if start >= end {
		if nums[start] < value {
			return start + 1
		}

		return start
	}

	q := start + (end-start)/2

	if nums[q] == value {
		return q + 1
	}

	if nums[q] > value {
		return bs(nums, start, q, value)
	} else {
		return bs(nums, q+1, end, value)
	}
}
