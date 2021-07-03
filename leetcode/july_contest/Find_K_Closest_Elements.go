package july_contest

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == k {
		return arr
	}

	//index := bs(arr, 0, len(arr)-1, x)
	index := ls(arr, 0, len(arr)-1, x)
	fmt.Println(index)

	//for i := index; i < len(arr); i++ {
	//	fmt.Println(index, arr[index])
	//	if absInt(arr[i]-x) < absInt(arr[index]-x) {
	//		index = i
	//	}
	//}

	//fmt.Println(index)
	//nearestValue := arr[index]
	start, end := index, index

	for end-start < k-1 {
		if start == 0 && end == len(arr)-1 {
			break
		} else if start == 0 {
			end++
		} else if end == len(arr)-1 {
			start--
		} else if absInt(arr[start-1]-x) > absInt(arr[end+1]-x) {
			end++
		} else {
			start--
		}
	}

	fmt.Println(start, end)

	return arr[start : end+1]
}

func bs(nums []int, start, end, target int) int {
	if start >= end {
		if nums[start] == target {
			return start
		} else {
			return start
		}
	}

	q := start + ((end - start) / 2)

	if nums[q] == target {
		return q
	} else if absInt(start-q) <= 2 {
		return ls(nums, start, q-1, target)
	} else if nums[q] > target {
		return bs(nums, start, q-1, target)
	} else {
		return bs(nums, q+1, end, target)
	}
}

func ls(nums []int, start, end, target int) int {
	nearest := 0

	for i := start; i <= end; i++ {
		if nums[i] == target {
			return i
		}

		if absInt(nums[i]-target) < absInt(nums[nearest]-target) {
			nearest = i
		}
	}

	return nearest
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
