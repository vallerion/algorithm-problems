package _933_Number_of_Recent_Calls

type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)

	for len(this.times) > 0 && this.times[0] < t-3000 {
		this.times = this.times[1:]
	}

	return len(this.times)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
