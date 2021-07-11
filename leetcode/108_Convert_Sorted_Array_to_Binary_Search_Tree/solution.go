package _108_Convert_Sorted_Array_to_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return binaryTraversal(nums, 0, len(nums)-1)
}

func binaryTraversal(nums []int, start, end int) *TreeNode {
	if start == end {
		return &TreeNode{nums[start], nil, nil}
	}

	if start > end {
		return nil
	}

	q := start + ((end - start) / 2)
	return &TreeNode{nums[q], binaryTraversal(nums, start, q-1), binaryTraversal(nums, q+1, end)}
}
