package _1302_Deepest_Leaves_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	st := make([]*TreeNode, 0)
	st = append(st, root)
	sum := 0

	for len(st) > 0 {
		n := len(st)
		sum = 0

		for i := 0; i < n; i++ {
			sum += st[i].Val
			if st[i].Left != nil {
				st = append(st, st[i].Left)
			}
			if st[i].Right != nil {
				st = append(st, st[i].Right)
			}
		}

		st = st[n:]
	}

	return sum
}
