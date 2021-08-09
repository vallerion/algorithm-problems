package _102_Binary_Tree_Level_Order_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	st := make([]*TreeNode, 0)
	st = append(st, root)

	for len(st) > 0 {
		n := len(st)

		result = append(result, make([]int, 0))

		for i := 0; i < n; i++ {
			result[len(result)-1] = append(result[len(result)-1], st[i].Val)

			if st[i].Left != nil {
				st = append(st, st[i].Left)
			}
			if st[i].Right != nil {
				st = append(st, st[i].Right)
			}
		}

		st = st[n:]
	}

	return result
}
