package _941_Valid_Mountain_Array

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	if arr[0] <= arr[1] || arr[len(arr)-1] >= arr[len(arr)-2] {
		return false
	}

	increasing := true

	for i := 1; i < len(arr); i++ {
		if increasing && arr[i] > arr[i-1] {
			continue
		}
		if increasing && arr[i] < arr[i-1] {
			increasing = false
		}
		if !increasing && arr[i] < arr[i-1] {
			continue
		}
		if !increasing && arr[i] > arr[i-1] {
			return false
		}
		if arr[i] == arr[i-1] {
			return false
		}
	}

	return true
}
