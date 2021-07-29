package _110_Balanced_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left, right := calcHeight(root.Left), calcHeight(root.Right)

	if left == -1 || right == -1 || abs(left-right) > 1 {
		return false
	}

	return true
}

func calcHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := calcHeight(root.Left)
	right := calcHeight(root.Right)

	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
