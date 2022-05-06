package _1089_Duplicate_Zeros

// [1,0,2,3,0,4,5,0]
// [1,0,0,2,3,0,0,4]

func duplicateZeros(arr []int) {
	st := make([]int, len(arr))
	slow, fast := 0, 0

	for fast < len(st) {
		st[fast] = arr[slow]
		if arr[slow] == 0 && fast < len(st)-1 {
			fast++
			st[fast] = arr[slow]
		}

		slow++
		fast++
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = st[i]
	}
}
