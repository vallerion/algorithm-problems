package _27_Remove_Element

// [0,1,2,2,3,0,4,2], val = 2
// j = 7, i = 0 -> if nums[j] == val { j--, continue } if nums[i] == val {swap(nums[i], nums[j]), j-- } i++
// j = 7, i = 0 -> j--
// j = 6, i = 0 -> i++
// j = 6, i = 1 -> i++
// j = 6, i = 2 -> [0,1,4,2,3,0,2,2], i++, j--
// j = 5, i = 3 -> [0,1,4,0,3,2,2,2], i++, j--
// j = 4, i = 4 -> [0,1,4,0,3,2,2,2], i++, j--

// [3,2,2,3], 3
// i=0,j=3 -> j--
// i=0,j=2 -> swap, j--, j++ -> [2,2,3,3]
// i=1,j=1


// [1,1,1,2,1], 1
// i=0,j=4 -> j--
// i=0,j=3 -> swap, j--, i++ -> [2,1,1,1,1]
// i=1,j=2 -> j--
// i=1,j=1
func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	i, j := 0, len(nums)-1

	for i < j {
		if nums[j] == val {
			j--
			continue
		}

		if nums[i] == val {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j--
		}

		i++
	}

	count := i

	for count < len(nums) && nums[count] != val {
		count++
	}

	return count
}
