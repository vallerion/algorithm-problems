package _328_Odd_Even_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 4 5 6 7

// slow=1, fast=2
// 1 3 2 4 5 6 7

// slow=3, fast=4
// 1 3 5 2 4 6 7

// slow=5, fast=6
// 1 3 5 7 2 4 6

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		temp := fast.Next
		fast.Next = fast.Next.Next

		temp2 := slow.Next
		slow.Next = temp
		temp.Next = temp2

		slow = slow.Next
		fast = fast.Next
	}

	return head
}
