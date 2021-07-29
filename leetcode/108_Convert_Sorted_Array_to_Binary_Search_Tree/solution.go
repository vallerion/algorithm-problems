package _108_Convert_Sorted_Array_to_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [-10,-3,0,5,9], len=5 -> 5/2=2
//    0  1 2 3 4

// 0 ->

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return arrToTree(nums, 0, len(nums)-1)
}

func arrToTree(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	q := start + ((end - start) / 2)

	return &TreeNode{nums[q], arrToTree(nums, start, q-1), arrToTree(nums, q+1, end)}
}
