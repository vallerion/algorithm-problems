package _496_Next_Greater_Element_I

// [4,1,2], [1,3,4,2]

// stack
// [2]
// [4]
// [4,3]
// [4,3,1]

// map
// 2 -> -1
// 4 -> -1
// 3 -> 4
// 1 -> 3

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ht := make(map[int]int)
	st := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] < nums2[i] {
			st = st[:len(st)-1]
		}

		if len(st) > 0 {
			ht[nums2[i]] = st[len(st)-1]
		} else {
			ht[nums2[i]] = -1
		}
		st = append(st, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		v, ok := ht[nums1[i]]
		if ok {
			nums1[i] = v
		}
	}

	return nums1
}
