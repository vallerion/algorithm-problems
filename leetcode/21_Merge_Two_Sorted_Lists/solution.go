package _21_Merge_Two_Sorted_Lists

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
//
// [1,2,3] [5,6,7]
// [2,3,7] [1,5,8] -> [1,2,3,7] [5,8] -> 2 < 5 -> 3 < 5 -> [1,2,3,5,7] [8] -> 7 < 8 -> [1,2,3,5,7,8] []
// [1,2,4] [1,3,4]
// [5] [1,2,4]

// Time: O(n + k), where n = length of l1, k = length of l2
// Space: O(2(n + k))
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var curr *ListNode

	for l1 != nil || l2 != nil {

		if l1 == nil {
			if res == nil {
				res = l2
				curr = res
			} else {
				curr.Next = l2
			}
			break
		}
		if l2 == nil {
			if res == nil {
				res = l1
				curr = res
			} else {
				curr.Next = l1
			}
			break
		}

		if l1.Val < l2.Val {
			if res == nil {
				res = &ListNode{l1.Val, nil}
				curr = res
			} else {
				curr.Next = &ListNode{l1.Val, nil}
				curr = curr.Next
			}
			l1 = l1.Next
		} else {
			if res == nil {
				res = &ListNode{l2.Val, nil}
				curr = res
			} else {
				curr.Next = &ListNode{l2.Val, nil}
				curr = curr.Next
			}
			l2 = l2.Next
		}
	}

	return res
}
