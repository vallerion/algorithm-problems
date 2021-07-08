package _26_Remove_Duplicates_from_Sorted_Array

// [0,0,1,1,1,2,2,3,3,4]
// ...
// [0,1,1,1,1,2,2,3,3,4]
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	j := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func removeDuplicatesSlow(nums []int) int {
	counter := 0

	for i := 1; i < len(nums); i++ {
		if counter >= len(nums)-i {
			break
		}

		if nums[i] == nums[i-1] {
			moveToEnd(&nums, i)
			counter++
			i--
		}
	}

	return len(nums) - counter
}

func moveToEnd(nums *[]int, index int) {
	for i := index; i < len(*nums)-1; i++ {
		temp := (*nums)[i]
		(*nums)[i] = (*nums)[i+1]
		(*nums)[i+1] = temp
	}
}
