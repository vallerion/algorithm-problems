package basic
//package main

/*
 * Task #2.3
 * @see ./statements.pdf
 */

import "fmt"

type pckg struct {
	time, duration int
}

func main() {
	var size, n, time, duration int

	fmt.Scanf("%d %d\n", &size, &n)

	if n == 0 {
		fmt.Println("")
	}

	queue := make([]int, 0)

	for n > 0 {
		n--
		fmt.Scanf("%d %d\n", &time, &duration)

		if duration == 0 && len(queue) == 0 {
			fmt.Println(time)
			continue
		}

		for len(queue) > 0 && queue[0] <= time {
			queue = queue[1:]
		}

		if len(queue) == size {
			fmt.Println("-1")
			continue
		}

		if len(queue) == 0 {
			fmt.Println(time)
			queue = append(queue, time+duration)
		} else {
			fmt.Println(queue[len(queue)-1])
			queue = append(queue, queue[len(queue)-1]+duration)
		}
	}
}

func mainOLD() {
	var size, n, time, duration int

	fmt.Scanf("%d %d\n", &size, &n)

	if n == 0 {
		fmt.Println("")
	}

	queue := make([]pckg, 0)

	for n > 0 {
		n--
		fmt.Scanf("%d %d\n", &time, &duration)

		if duration == 0 && len(queue) == 0 {
			fmt.Println(time)
			continue
		}

		//fmt.Println(time, duration, len(queue), size)

		if len(queue) == size {
			spentTime := queue[len(queue)-1].time + queue[len(queue)-1].duration

			if time < spentTime {
				fmt.Println("-1")
				continue
			}

			if time >= spentTime {
				queue = make([]pckg, 0)
			}
		} else if len(queue) > 0 {
			spentTime := queue[len(queue)-1].time + queue[len(queue)-1].duration

			if time < spentTime {
				time = spentTime
			}
		}

		queue = append(queue, pckg{time, duration})
		fmt.Println(time)
	}
}
