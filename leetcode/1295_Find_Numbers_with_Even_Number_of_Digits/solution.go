package _1295_Find_Numbers_with_Even_Number_of_Digits

func findNumbers(nums []int) int {
	evenCnt := 0

	for i := 0; i < len(nums); i++ {
		if numOfDigits(nums[i])%2 == 0 {
			evenCnt++
		}
	}

	return evenCnt
}

func numOfDigits(num int) int {
	res := 0

	for num > 0 {
		num /= 10
		res++
	}

	return res
}
