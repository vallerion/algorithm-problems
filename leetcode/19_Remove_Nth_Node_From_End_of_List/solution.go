package _19_Remove_Nth_Node_From_End_of_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	nodes := 0

	for curr != nil {
		nodes++
		curr = curr.Next
	}

	i := 0
	var prev *ListNode
	curr = head

	for curr != nil {
		if i == nodes-n {
			if prev != nil {
				prev.Next = curr.Next
			} else {
				head = curr.Next
			}
			break
		}

		i++
		prev = curr
		curr = curr.Next
	}

	return head
}
