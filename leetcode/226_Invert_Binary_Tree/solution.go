package _226_Invert_Binary_Tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	return root
}