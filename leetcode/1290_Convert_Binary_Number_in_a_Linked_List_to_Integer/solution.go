package _1290_Convert_Binary_Number_in_a_Linked_List_to_Integer

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	count := 0
	curr := head
	decimal := 0

	for curr != nil {
		curr = curr.Next
		count++
	}

	curr = head

	for i := count - 1; i >= 0; i-- {
		decimal += curr.Val * int(math.Pow(2, float64(i)))
		curr = curr.Next
	}

	return decimal
}
