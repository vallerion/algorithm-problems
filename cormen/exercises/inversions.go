package exercises

// Inversions
// Build algo to count of inversions in integers array.
// Inversion it means i < j, A[i] > A[j],
// where A - input array, i, j <= n, n - length of array.
// Make it in O(n * log n).
// Hint: modify merge sort.
//
// Ex:
// input: [2,3,8,6,1] -> output: 5
// Explanation: [2,1], [3,1], [8,6], [8,1], [6,1]
// ---------
// 1. Max count inversions in worse case?
// Arithmetic progression: n*(n-1)/2
// 2. It easy do in brute force way:
// 	  a. Take first element
//	  b. Compare it with each next
//	  c. Increase counter if conditions are met
//	  d. We have O(n^2) :(
// 3. Okay, lets modify merge sort,
//	  just add counter in merge functions.
//	  Test case: [15,10,8,6,1,0], it must return 15
// 4. Ok, i did it, but with internet
func Inversions(nums []int) int {
	temp := make([]int, len(nums))
	return inversions(nums, temp, 0, len(nums)-1)
}

func inversions(nums, temp []int, start, end int) int {
	counter := 0

	if start >= end {
		return counter
	}

	q := (start + end) / 2

	counter += inversions(nums, temp, start, q)
	counter += inversions(nums, temp, q+1, end)

	counter += mergeCounter(nums, temp, start, q+1, end)

	return counter
}

func mergeCounter(nums, temp []int, start, mid, end int) int {
	i := start
	j := mid
	counter := 0

	for i < mid && j < end+1 {
		if i == mid {
			temp[i+j-mid] = nums[j]
			j++
			continue
		}

		if j == end+1 {
			temp[i+j-mid] = nums[i]
			i++
			continue
		}

		if nums[i] < nums[j] {
			temp[i+j-mid] = nums[i]
			i++
			continue
		} else {
			temp[i+j-mid] = nums[j]
			j++
			counter += mid - i
			continue
		}
	}

	return counter
}
