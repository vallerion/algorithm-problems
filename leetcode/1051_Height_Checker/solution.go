package _1051_Height_Checker

// 1,1,4,2,1,3

// 0 -> 3
// 1 -> 1
// 2 -> 1
// 3 -> 1

func heightChecker(heights []int) int {
	bucket := make([]int, 101)

	for i := 0; i < len(heights); i++ {
		bucket[heights[i]-1]++
	}

	res := 0
	bucketPtr := 0

	for i := 0; i < len(heights); i++ {
		for bucket[bucketPtr] == 0 {
			bucketPtr++
		}

		if heights[i] != bucketPtr+1 {
			res++
		}
		bucket[bucketPtr]--
	}

	return res
}

// heightCheckerF - first version, O(n) but uses a lot of memory
func heightCheckerF(heights []int) int {
	bucket := make([][]int, 101)

	for i := 0; i < len(heights); i++ {
		if bucket[heights[i]-0] == nil {
			bucket[heights[i]-0] = make([]int, 0)
		}
		bucket[heights[i]-0] = append(bucket[heights[i]], heights[i])
	}

	cp := make([]int, 0)

	for i := 0; i < len(bucket); i++ {
		if bucket[i] == nil {
			continue
		}

		cp = append(cp, bucket[i]...)
	}

	res := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != cp[i] {
			res++
		}
	}

	return res
}
