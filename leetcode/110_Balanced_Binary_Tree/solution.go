package _110_Balanced_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if height(root) == -1 {
		return false
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight, rightHeight := height(root.Left), height(root.Right)

	if leftHeight == -1 || rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}
