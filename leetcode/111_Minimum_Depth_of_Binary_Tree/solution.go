package _111_Minimum_Depth_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	st := make([]*TreeNode, 0)
	st = append(st, root)
	level := 0

	for len(st) > 0 {
		n := len(st)

		for i := 0; i < n; i++ {
			node := st[i]

			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				st = append(st, node.Left)
			}
			if node.Right != nil {
				st = append(st, node.Right)
			}
		}

		st = st[n:]
		level++
	}

	return level
}
