package _112_Path_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	if hasPathSum(root.Left, targetSum) {
		return true
	}
	if hasPathSum(root.Right, targetSum) {
		return true
	}
	targetSum += root.Val

	return false
}
