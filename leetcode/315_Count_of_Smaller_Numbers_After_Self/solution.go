package _315_Count_of_Smaller_Numbers_After_Self

// [5,2,6,1]
// [5] [2] [6] [1]
// [2,5] | [1 0 0 0]
// [1,6] | [1 0 1 0]
// [1,2,5,6] | [2 1 1 0]

// [2,5,2,6,1]

// [2] [5] [2] [6] [1]
// [2]+[5] 			-> [2,5] 		| [0 0 0 0 0]
// [2,5]+[2]		-> [2,2,5] 		| [0 0 0 0 0]
// [6]+[1]			-> [1,6]		| [0 0 0 1 0]
// [2,2,5]+[1,6]	-> [1]

// 0 4 -> 2

type original struct {
	value, initialIndex int
}

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// hash map is bad idea, because values in array can be duplicated
	originals := make([]original, len(nums)) // store initial positions in array

	for i := 0; i < len(nums); i++ {
		originals[i].value = nums[i]
		originals[i].initialIndex = i
	}

	counts := make([]int, len(nums)) // result array

	mergeSort(originals, counts, 0, len(nums)-1)

	return counts
}

func mergeSort(originals []original, counts []int, start, end int) {
	if start >= end {
		return
	}

	q := start + (end-start)/2

	mergeSort(originals, counts, start, q)
	mergeSort(originals, counts, q+1, end)

	merge(originals, counts, start, q, end)
}

func merge(originals []original, counts []int, start, middle, end int) {
	if start >= end {
		return
	}

	result := make([]original, 0)

	left := start       // left array - start to middle
	right := middle + 1 // right array - middle+1 to end

	leftLessThenRight := 0
	for left < middle+1 && right <= end {
		if originals[left].value > originals[right].value {
			leftLessThenRight++

			result = append(result, originals[right])
			right++
		} else {
			initialPosition := originals[left].initialIndex
			counts[initialPosition] += leftLessThenRight

			result = append(result, originals[left])
			left++
		}
	}

	for left < middle+1 {
		initialPosition := originals[left].initialIndex
		counts[initialPosition] += leftLessThenRight

		result = append(result, originals[left])
		left++
	}

	for right <= end {
		result = append(result, originals[right])
		right++
	}

	j := 0
	for i := start; i <= end; i++ {
		originals[i] = result[j]
		j++
	}
}
