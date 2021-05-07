package exercises

// merge (A, i, j, p)
//	R[p]
//	cj = j
//	for
//		if i == cj && j == p
//			break
//
//		if i == cj
//			R[i+j-cj] = A[j]
//			j++
//			continue
//
//		if j == p
//			R[i+j-cj] = A[i]
//			i++
//			continue
//
//		if A[i] < A[j]
//			R[i+j-cj] = A[i]
//			i++
//			continue
//		else
//			R[i+j-cj] = A[j]
//			j++
//			continue
//
//
//
// [1,2,3,5]+[4,5,7,10] -> [1,2,3,5,4,5,7,10], i=0, j=len/2=4, mid=4, end=7
// i=0,j=4 -> 1/4 -> R[i+j-mid]->R[0]=[1] -> [1]
// i=1,j=4 -> 2/4 -> R[i+j-mid]->R[1]=[2] -> [1,2]
// i=2,j=4 -> 3/4 -> R[i+j-mid]->R[2]=[3] -> [1,2,3]
// i=3,j=4 -> 5/4 -> R[i+j-mid]->R[3]=[4] -> [1,2,3,4]
// i=3,j=5 -> 5/5 -> R[i+j-mid]->R[4]=[5] -> [1,2,3,4,5]
// i=3,j=6 -> 5/7 -> R[i+j-mid]->R[5]=[5] -> [1,2,3,4,5,5]
// i=4,j=6 -> x/7 -> R[i+j-mid]->R[6]=[7] -> [1,2,3,4,5,5,7]
// i=4,j=7 -> x/10 -> R[i+j-mid]->R[7]=[10] -> [1,2,3,4,5,5,7,10]
// i=4,j=8 -> x/x -> R[i+j-mid]->R[7]=[10] -> [1,2,3,4,5,5,7,10]

// tests
// Merge([]int{3,5,1,2}, 0, 2, 4),

func MergeIndexes(arr []int, start, mid, end int) ([]int, int) {
	result := make([]int, end+1)
	i := start
	j := mid
	iter := 0

	for {
		if i == mid && j == end+1 {
			break
		}
		iter++

		if i == mid {
			result[i+j-mid] = arr[j]
			j++
			continue
		}

		if j == end+1 {
			result[i+j-mid] = arr[i]
			i++
			continue
		}

		if arr[i] < arr[j] {
			result[i+j-mid] = arr[i]
			i++
			continue
		} else {
			result[i+j-mid] = arr[j]
			j++
			continue
		}
	}

	return result, iter
}

func Merge(first, second []int) []int {
	result := make([]int, len(first) + len(second))
	i := 0
	j := 0

	for {
		if i == len(first) && j == len(second) {
			break
		}

		if i == len(first) {
			result[i+j] = second[j]
			j++
			continue
		}

		if j == len(second) {
			result[i+j] = first[i]
			i++
			continue
		}

		if first[i] < second[j] {
			result[i+j] = first[i]
			i++
			continue
		} else {
			result[i+j] = second[j]
			j++
			continue
		}
	}

	return result
}

// MergeSort todo: working with indexes instead of copy slices can use less memory
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	q := len(arr) / 2

	left := MergeSort(arr[:q])
	right := MergeSort(arr[q:])

	return Merge(left, right)
}
