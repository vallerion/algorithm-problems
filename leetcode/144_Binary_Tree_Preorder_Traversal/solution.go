package _144_Binary_Tree_Preorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	result := make([]int, 0)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, top.Val)

		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return result
}

func preorderTraversalRecursive(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	traversal(root.Left, res)
	traversal(root.Right, res)
}
