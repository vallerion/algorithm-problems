package _143_Reorder_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	nums := make([]int, 0)

	curr := head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	if len(nums) == 1 {
		return
	}

	curr = head
	for len(nums) > 0 && curr != nil {
		left, right := nums[0], nums[len(nums)-1]
		if len(nums) == 1 {
			nums = []int{}
		} else {
			nums = nums[1 : len(nums)-1]
		}

		if curr.Val != left {
			curr.Val = left
		}

		if curr.Next != nil {
			curr.Next.Val = right
		}

		if curr.Next != nil && curr.Next.Next != nil {
			curr = curr.Next.Next
		} else {
			curr = nil
		}

	}
}
