package _2365_Task_Scheduler_II

// A B B C   2

// A -> 0 / 1
// B -> 1 / 2
// B -> 4 / 5
// C -> 5 / 6

// A B A B C   2

// A -> 0 / 1
// B -> 1 / 2
// A -> 2 - (2 - 0 - 1)

func taskSchedulerII(tasks []int, space int) int64 {
	positions := make(map[int]int)
	curr := 0

	for _, v := range tasks {
		if pos, ok := positions[v]; ok && curr-pos <= space {
			curr += space - (curr - pos - 1)
		}
		positions[v] = curr
		curr++
	}

	return int64(curr - 1)
}
