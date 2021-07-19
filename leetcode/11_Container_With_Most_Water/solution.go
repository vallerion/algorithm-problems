package _11_Container_With_Most_Water

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	highest := 0

	for i < j {
		val := min(height[i], height[j]) * (j - i)

		if val > highest {
			highest = val
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return highest
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
