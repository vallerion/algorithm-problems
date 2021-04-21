package main

import "fmt"

type NodeTree struct {
	Key   int
	Value string
	Left  *NodeTree
	Right *NodeTree
}

func Find(root *NodeTree, needle int) *string {
	// find in sorted tree
	if root == nil {
		return nil
	}

	if needle == root.Key {
		return &root.Value
	} else if needle > root.Key {
		return Find(root.Right, needle)
	} else {
		return Find(root.Left, needle)
	}
}

func Add(root, new *NodeTree) {
	// find in sorted tree
	if root == nil {
		root = new
	}

	if new.Key == root.Key {
		root.Value = new.Value
		return
	} else if new.Key > root.Key {
		if root.Right != nil {
			Add(root.Right, new)
		} else {
			root.Right = new
		}
	} else {
		if root.Left != nil {
			Add(root.Left, new)
		} else {
			root.Left = new
		}
	}
}

func PrintTree(root *NodeTree) {
	if root == nil {
		return
	}

	PrintTree(root.Left)
	fmt.Print(root.Value)
	fmt.Print("-")
	PrintTree(root.Right)
}

func main2() {
	//	      10
	//     /     \
	//    5       35
	//   /      /    \
	//  1      20    99
	//   \    /  \
	//    4  17  31
	tree := NodeTree{
		10,
		"root",
		&NodeTree{
			5,
			"root5",
			&NodeTree{1, "root1", nil, &NodeTree{4, "root4", nil, nil}},
			nil,
		},
		&NodeTree{
			35,
			"root35",
			&NodeTree{20, "root20", &NodeTree{17, "root17", nil, nil}, &NodeTree{31, "root31", nil, nil}},
			&NodeTree{99, "root99", nil, nil},
		},
	}

	//fmt.Print(*Find(&tree, 20))
	Add(&tree, &NodeTree{5, "doot5", nil, nil})
	PrintTree(&tree)
}
