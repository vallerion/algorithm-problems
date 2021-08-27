package _94_Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	result := make([]int, 0)

	root = root.Left

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, top.Val)
		root = top.Right
	}

	return result
}

func inorderTraversalRecursive(root *TreeNode) []int {
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