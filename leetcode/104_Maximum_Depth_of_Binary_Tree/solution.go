package _104_Maximum_Depth_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0

	depth = max(depth, maxDepth(root.Left))
	depth = max(depth, maxDepth(root.Right))

	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
