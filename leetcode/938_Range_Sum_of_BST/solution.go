package _938_Range_Sum_of_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [1,2,3,4,5]
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val >= low && root.Val <= high {
		sum = root.Val
	}

	return sum + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
