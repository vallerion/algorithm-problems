package _852_Peak_Index_in_a_Mountain_Array

// [0,1,2,3,4,10,3]
// l = 4, r = 6, [4,10,3]
//

func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 3 {
		return arr[1]
	}

	l, r := 0, len(arr)-1
	peak := 1

	for l <= r {
		q := (l + r) / 2

		if q-1 >= 0 && arr[q-1] > arr[q] {
			r = q - 1
		} else if q+1 <= len(arr)-1 && arr[q+1] > arr[q] {
			l = q + 1
		} else {
			peak = q
			break
		}
	}

	return peak
}

func peakIndexInMountainArray_linear(arr []int) int {
	index := 1

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			index = i
			break
		}
	}

	return index
}
