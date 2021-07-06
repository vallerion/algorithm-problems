package _11_Container_With_Most_Water

// [1,8,6,2,5,4,8,3,7]
func maxArea(height []int) int {
	maxRange := 0

	i, j := 0, len(height)-1

	for i < j {
		rng := min(height[i], height[j])
		rng *= j - i
		maxRange = max(maxRange, rng)

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return maxRange
}

// bruteForce
func maxAreaBF(height []int) int {
	maxRange := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			rng := min(height[i], height[j])
			rng *= j - i
			maxRange = max(maxRange, rng)
		}
	}
	return maxRange
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
