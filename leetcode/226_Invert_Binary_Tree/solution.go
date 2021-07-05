package _226_Invert_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {

	if root == nil {
		return
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invert(root.Left)
	invert(root.Right)
}
