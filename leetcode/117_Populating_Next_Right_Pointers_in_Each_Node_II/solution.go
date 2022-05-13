package _117_Populating_Next_Right_Pointers_in_Each_Node_II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	st := make([]*Node, 0)
	st = append(st, root)

	for len(st) > 0 {
		n := len(st)
		var prev *Node
		for i := 0; i < n; i++ {
			if prev != nil {
				st[i].Next = prev
			}
			prev = st[i]
			if st[i].Right != nil {
				st = append(st, st[i].Right)
			}
			if st[i].Left != nil {
				st = append(st, st[i].Left)
			}
		}

		st = st[n:]
	}

	return root
}
