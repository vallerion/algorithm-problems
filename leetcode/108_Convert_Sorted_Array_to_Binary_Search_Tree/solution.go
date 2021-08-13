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

	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	q := start + ((end - start) / 2)

	return &TreeNode{nums[q], helper(nums, start, q-1), helper(nums, q+1, end)}
}
