package _876_Middle_of_the_Linked_List

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	n := 0

	curr := head
	for curr != nil {
		n++
		curr = curr.Next
	}

	middle := int(math.Ceil(float64(n) / 2))
	n = 0
	curr = head
	for curr != nil {
		if n == middle {
			return curr
		}
		n++
		curr = curr.Next
	}

	return nil
}
