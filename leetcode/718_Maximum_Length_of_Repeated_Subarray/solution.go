package _718_Maximum_Length_of_Repeated_Subarray

// [1,2,3,2,1]
// [3,2,1,4,7]

// 0,2 -> 1,3
// 1,1 -> 2,2
// 2,0 -> 3,1
// 3,1 -> 4,2
// 4,2 -> 5,3

// 0 0 0 0 0 0
// 0 0 0 1 0 0
// 0 0 1 0 0 0
// 0 1 0 0 0 0
// 0 0 2 0 0 0
// 0 0 0 3 0 0

func findLength(nums1 []int, nums2 []int) int {
	matrix := make([][]int, len(nums1)+1)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, len(nums2)+1)
	}

	maxLength := 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				matrix[i+1][j+1] = matrix[i][j] + 1
			}

			if matrix[i+1][j+1] > maxLength {
				maxLength = matrix[i+1][j+1]
			}
		}
	}

	return maxLength
}
