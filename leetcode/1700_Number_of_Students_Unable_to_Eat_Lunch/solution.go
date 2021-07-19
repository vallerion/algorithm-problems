package _1700_Number_of_Students_Unable_to_Eat_Lunch

// [1,1,1,0,0,1] [1,0,0,0,1,1]

// [1,1,0,0,1] [0,0,0,1,1]
// [1,0,0,1,1] [0,0,0,1,1]
// [0,0,1,1,1] [0,0,0,1,1]
// [0,1,1,1]   [0,0,1,1]
// [1,1,1]     [0,1,1]

func countStudents(students []int, sandwiches []int) int {
	hp := make(map[int]int)

	for i := 0; i < len(students); i++ {
		hp[students[i]]++
	}

	for len(sandwiches) > 0 {
		top := sandwiches[0]
		sandwiches = sandwiches[1:]

		if hp[top] == 0 {
			break
		}

		hp[top]--
	}

	return hp[0] + hp[1]
}
