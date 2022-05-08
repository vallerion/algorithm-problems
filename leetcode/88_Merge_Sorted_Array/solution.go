package _88_Merge_Sorted_Array

// [1,2,5,0,0] [0,3] st=[]
// [1,2,5,0,3]
// 0,3 [0,2,5,0,3] [1]
// 1,4 [0,1,5,0,3] [2]
// 2,4 [0,1,2,0,5] [3]
// 3,5 [0,1,2,3,5] [3]

// [1,2,3,0,0,0]
// [1,2,3,0,1,2]
// [0,1,1,2,3,2] [2,3]

func merge(nums1 []int, m int, nums2 []int, n int) {
	hp := make([]int, 0)

	i, j := 0, 0

	for i < m || j < n {
		if i >= m {
			hp = append(hp, nums2[j])
			j++
		} else if j >= n {
			hp = append(hp, nums1[i])
			i++
		} else if nums1[i] <= nums2[j] {
			hp = append(hp, nums1[i])
			i++
		} else {
			hp = append(hp, nums2[j])
			j++
		}
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = hp[i]
	}
}
