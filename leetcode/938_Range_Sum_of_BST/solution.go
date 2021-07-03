package _938_Range_Sum_of_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0

	if root == nil {
		return res
	}

	if root.Val >= low && root.Val <= high {
		res += root.Val
		return rangeSumBST(root.Left, low, high) + res + rangeSumBST(root.Right, low, high)
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high) + res
	} else {
		return res + rangeSumBST(root.Right, low, high)
	}
}
