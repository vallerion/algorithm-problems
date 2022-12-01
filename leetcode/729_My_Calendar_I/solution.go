package _729_My_Calendar_I

// [10,20][5,25][8,9][20,30]

// 1. one of time point inside
// 2. left smaller and right bigger

type MyCalendar struct {
	nums []int
}

func Constructor() MyCalendar {
	return MyCalendar{nums: make([]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < len(this.nums); i += 2 {
		if start > this.nums[i] && start < this.nums[i+1] {
			return false
		}
		if end > this.nums[i] && end < this.nums[i+1] {
			return false
		}
		if start < this.nums[i] && end > this.nums[i+1] {
			return false
		}
		if start == this.nums[i] || end == this.nums[i+1] {
			return false
		}
	}

	this.nums = append(this.nums, start, end)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
