package _1299_Replace_Elements_with_Greatest_Element_on_Right_Side

// [17,18,5,4,6,1]
// 6, max=1 -> [17,18,5,4,1,1]
// 4, max=6 -> [17,18,5,6,1,1]
// 5, max=6 -> [17,18,6,6,1,1]
// 18, max=6 -> [17,6,6,6,1,1]
// 17, max=18 -> [18,6,6,6,1,1]

func replaceElements(arr []int) []int {
	if len(arr) == 1 {
		arr[0] = -1
		return arr
	}

	max := arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		arr[i] = max

		if temp > max {
			max = temp
		}
	}

	arr[len(arr)-1] = -1

	return arr
}
