package _19_Remove_Nth_Node_From_End_of_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	total := 0

	curr := head
	for curr != nil {
		total++
		curr = curr.Next
	}

	if total == 1 {
		return nil
	}

	delIndex := total - n
	i := 0

	curr = head
	var prev *ListNode
	for curr != nil {

		if i == delIndex {
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
			break
		}

		prev = curr
		curr = curr.Next
		i++
	}

	return head
}
