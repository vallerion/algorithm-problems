package _124_Binary_Tree_Maximum_Path_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeMax struct {
	OneSide  int
	BothSide int
}

var maxes map[*TreeNode]nodeMax

func maxPathSum(root *TreeNode) int {
	maxes = make(map[*TreeNode]nodeMax)
	calcMaxes(root)

	res := root.Val
	for _, v := range maxes {
		res = max(res, v.BothSide)
		res = max(res, v.OneSide)
	}

	return res
}

func calcMaxes(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		maxes[root] = nodeMax{root.Val, root.Val}
		return
	}

	calcMaxes(root.Left)
	calcMaxes(root.Right)

	if root.Left != nil && root.Right != nil {
		one := max(root.Val, root.Val+max(maxes[root.Left].OneSide, maxes[root.Right].OneSide))
		both := max(root.Val, root.Val+maxes[root.Left].OneSide+maxes[root.Right].OneSide)
		maxes[root] = nodeMax{one, both}
	} else if root.Left != nil {
		v := max(root.Val, root.Val+maxes[root.Left].OneSide)
		maxes[root] = nodeMax{v, v}
	} else {
		v := max(root.Val, root.Val+maxes[root.Right].OneSide)
		maxes[root] = nodeMax{v, v}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
