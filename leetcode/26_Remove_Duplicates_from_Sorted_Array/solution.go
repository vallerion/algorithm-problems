package _26_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
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
