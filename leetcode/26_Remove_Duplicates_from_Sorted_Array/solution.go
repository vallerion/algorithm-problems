package _26_Remove_Duplicates_from_Sorted_Array

// [0,0,1,1,1,2,2,3,3,4]
// [0,1,2,3,4,2,2,3,3,4]

// i=0,j=1 -> if nums[i]!=nums[j]: nums[i+1]=nums[j],i++,j++, else: j++
// i=0,j=2 -> nums[i+1]=nums[j],i++,j++
// i=1,j=3 -> j++
// i=1,j=4 -> j++
// i=1,j=5 -> nums[i+1]=nums[j],i++,j++
// i=2,j=6 -> j++
// i=2,j=7 -> nums[i+1]=nums[j],i++,j++
// i=3,j=8 -> j++
// i=3,j=9 -> nums[i+1]=nums[j],i++,j++
// i=4,j=10

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}


	i, j := 0, 1

	for j < len(nums) {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}

		j++
	}

	return i+1
}
