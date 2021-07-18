package _94_Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, res)
	*res = append(*res, root.Val)
	traversal(root.Right, res)
}