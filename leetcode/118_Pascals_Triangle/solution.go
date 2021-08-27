package _118_Pascals_Triangle

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	result := [][]int{{1}, {1, 1}}

	for i := 2; i < numRows; i++ {
		result = append(result, genFromPrev(&result[i-1]))
	}

	return result
}

func genFromPrev(nums *[]int) []int {
	res := []int{1}

	for i := 0; i < len(*nums)-1; i++ {
		res = append(res, (*nums)[i]+(*nums)[i+1])
	}

	res = append(res, 1)

	return res
}
