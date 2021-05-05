package _203_Remove_Linked_List_Elements

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	curr := head
	var prev *ListNode = nil

	for curr != nil {
		if curr.Val == val && prev == nil {
			head = curr.Next
		} else if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return head
}
