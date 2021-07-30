package _100_Same_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	leftSame := isSameTree(p.Left, q.Left)
	if leftSame == false {
		return false
	}

	rightSame := isSameTree(p.Right, q.Right)
	if rightSame == false {
		return false
	}

	return p.Val == q.Val
}
