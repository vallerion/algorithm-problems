package _133_Clone_Graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	if len(node.Neighbors) == 0 {
		return &Node{node.Val, []*Node{}}
	}

	return helper(node, make(map[int]*Node))
}

func helper(node *Node, hp map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := hp[node.Val]; ok {
		return hp[node.Val]
	}

	newNode := new(Node)
	newNode.Val = node.Val
	newNode.Neighbors = []*Node{}

	hp[node.Val] = newNode

	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors = append(newNode.Neighbors, helper(node.Neighbors[i], hp))
	}

	return newNode
}
