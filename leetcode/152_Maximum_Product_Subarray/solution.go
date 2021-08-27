package _152_Maximum_Product_Subarray

// [-1,-2,-3,0]

// -1 2 -6 1
// -6 6 -3 1

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	product := 1
	maxProd := 0

	for i := 0; i < len(nums); i++ {
		product *= nums[i]

		if product > maxProd {
			maxProd = product
		}

		if nums[i] == 0 {
			product = 1
		}
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		product *= nums[i]

		if product > maxProd {
			maxProd = product
		}

		if nums[i] == 0 {
			product = 1
		}
	}

	return maxProd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(n^2)
func maxProductON2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxProd := nums[0]

	for i := 0; i < len(nums); i++ {
		currentProduct := nums[i]
		if maxProd < currentProduct {
			maxProd = currentProduct
		}

		for j := i + 1; j < len(nums); j++ {
			currentProduct *= nums[j]
			if maxProd < currentProduct {
				maxProd = currentProduct
			}
		}
	}

	return maxProd
}
