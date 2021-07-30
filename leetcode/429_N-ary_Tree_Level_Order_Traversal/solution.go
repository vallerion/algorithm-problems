package _429_N_ary_Tree_Level_Order_Traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	st := make([]*Node, 0)
	st = append(st, root)

	level := 0

	for len(st) > 0 {
		result = append(result, make([]int, 0))

		n := len(st)

		for i := 0; i < n; i++ {
			result[level] = append(result[level], st[i].Val)

			if len(st[i].Children) > 0 {
				st = append(st, st[i].Children...)
			}
		}

		st = st[n:]
		level++
	}

	return result
}
