package _145_Binary_Tree_Postorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalRecursive(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, res)
	traversal(root.Right, res)
	*res = append(*res, root.Val)
}
