package sorting

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
// [1,2,3,5]+[4,5,7,10] -> [1,2,3,5,4,5,7,10], i=0, j=len/2=4, cj=j=4
// i=0,j=4 -> 1/4 -> R[i+j-cj]->R[0]=[1] -> [1]
// i=1,j=4 -> 2/4 -> R[i+j-cj]->R[1]=[2] -> [1,2]
// i=2,j=4 -> 3/4 -> R[i+j-cj]->R[2]=[3] -> [1,2,3]
// i=3,j=4 -> 5/4 -> R[i+j-cj]->R[3]=[4] -> [1,2,3,4]
// i=3,j=5 -> 5/5 -> R[i+j-cj]->R[4]=[5] -> [1,2,3,4,5]
// i=3,j=6 -> 5/7 -> R[i+j-cj]->R[5]=[5] -> [1,2,3,4,5,5]
// i=4,j=6 -> x/7 -> R[i+j-cj]->R[6]=[7] -> [1,2,3,4,5,5,7]
// i=4,j=7 -> x/10 -> R[i+j-cj]->R[7]=[10] -> [1,2,3,4,5,5,7,10]
// i=4,j=8 -> x/x -> R[i+j-cj]->R[7]=[10] -> [1,2,3,4,5,5,7,10]

// tests
// Merge([]int{3,5,1,2}, 0, 2, 4),

func Merge(arr []int, i, j, l int) ([]int, int) {
	result := make([]int, l)
	cj := j
	iter := 0

	for {
		if i == cj && j == l {
			break
		}
		iter++

		if i == cj {
			result[i+j-cj] = arr[j]
			j++
			continue
		}

		if j == l {
			result[i+j-cj] = arr[i]
			i++
			continue
		}

		if arr[i] < arr[j] {
			result[i+j-cj] = arr[i]
			i++
			continue
		} else {
			result[i+j-cj] = arr[j]
			j++
			continue
		}
	}

	return result, iter
}
