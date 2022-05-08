package _1346_Check_If_N_and_Its_Double_Exist

func checkIfExist(arr []int) bool {
	mp := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = i
	}

	for i := 0; i < len(arr); i++ {
		v, ok := mp[arr[i]*2]
		if ok && v != i {
			return true
		}
	}

	return false
}
