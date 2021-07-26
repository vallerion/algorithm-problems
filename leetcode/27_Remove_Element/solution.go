package _27_Remove_Element

// [0,1,2,2,3,0,4,2] // 0
// [2,1,2,2,3,0,4,2]
// ...
// [2,1,2,2,3,4,4,2], j=5

// [0,1,2,2,3,0,4,2] // 0
// [1,2,2,3,4,2,4,2]
// ...
// [
func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	i, j := 0, 0

	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	return i
}
