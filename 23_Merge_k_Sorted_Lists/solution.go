package _23_Merge_k_Sorted_Lists

import "fmt"

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

func mergeKLists(lists []*ListNode) *ListNode {
	mp := make(map[int]int)
	maxValue := 0
	minValue := 0

	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			maxValue = max(maxValue, lists[i].Val)
			minValue = min(minValue, lists[i].Val)
			mp[lists[i].Val]++
			lists[i] = lists[i].Next
		}
	}

	var head *ListNode
	curr := head

	for i := minValue; i <= maxValue; i++ {
		if v, ok := mp[i]; ok {
			for j := 0; j < v; j++ {
				if curr == nil {
					head = &ListNode{i, nil}
					curr = head
				} else {
					curr.Next = &ListNode{i, nil}
					curr = curr.Next
				}
			}
		}
	}

	return head
}

func printLL(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}