package _67_Add_Binary

// a = "1010", b = "1011"
// 1 + 0 + 0 = 1 > 1 -> 1, add = 0
// 1 + 1 + 0 = 2 > 1 -> 0, add = 1
// 0 + 0 + 1 = 1 > 1 -> 1, add = 0
// 1 + 1 + 0 = 2 > 1 -> 0, add = 1
// 1 0 1 0 1

// a = "1111", b = "1111", add=0
// 1+1+0=2, 2%2=0, add=2>>1=1
// 1+1+1=3, 3%2=1, add=3>>1=1
// 1+1+1=3, 3%2=1, add=3>>1=1
// 1+1+1=3, 3%2=1, add=3>>1=1
// 0+0+1=1, 1%2=1, add=1>>1=0

func addBinary(a string, b string) string {
	ar := []byte(a)
	br := []byte(b)
	res := ""

	i := len(ar) - 1
	j := len(br) - 1
	currSum := byte(0)
	add := byte(0)

	for i >= 0 || j >= 0 || add > 0 {
		dar := byte(0)
		dbr := byte(0)

		if i >= 0 {
			dar = ar[i] - '0'
		}
		if j >= 0 {
			dbr = br[j] - '0'
		}

		currSum = dar + dbr + add
		add = currSum >> 1

		res = string(currSum%2+'0') + res

		i--
		j--
	}

	return res
}
