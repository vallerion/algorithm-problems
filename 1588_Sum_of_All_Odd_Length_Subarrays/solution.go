package _1588_Sum_of_All_Odd_Length_Subarrays


// arr = [1,4,2,5,3], i = 0, j = 0, sum = 0, res[]
// j = 0,2,4 ... i += 2
// i=0, j=i+2 -> sum += arr[i]
// i=0, j=2 -> sum += sum(arr[i..j]), sum=1+4+2=7
// i=0, j=4 -> sum += sum(arr[0..4]), sum=1+4+2+5+3=15
// sum += arr[i], sum+=4
// i=1, j=3 -> sum += sum(arr[1..3]), sum=4+2+5=11

func sumOddLengthSubarrays(arr []int) int {
	total := 0
	for i:=0; i< len(arr); i++ {
		total += arr[i]
		for j:=i+2; j<len(arr); j+=2 {
			total += sum(arr[i:j+1])
		}
	}

	return total
}

func sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}