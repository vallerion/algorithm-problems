package _671_Second_Minimum_Node_In_a_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	nextSmall := root.Val

	if root.Left == nil {
		return -1
	}

	if root.Val == root.Left.Val && root.Val == root.Right.Val {
		left := findSecondMinimumValue(root.Left)
		right := findSecondMinimumValue(root.Right)
		if left == -1 && right == -1 {
			return -1
		} else if left == -1 {
			return right
		} else if right == -1 {
			return left
		} else {
			return min(left, right)
		}

	} else if root.Val == root.Left.Val {
		nextSmall = root.Right.Val
		left := findSecondMinimumValue(root.Left)
		if left != -1 {
			nextSmall = min(nextSmall, left)
		}
	} else {
		nextSmall = root.Left.Val
		right := findSecondMinimumValue(root.Right)
		if right != -1 {
			nextSmall = min(nextSmall, right)
		}
	}

	return nextSmall
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
