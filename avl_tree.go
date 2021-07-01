package main

import "fmt"

type TreeValue interface {
	// returns:
	// 1  - if current value greater then input value
	// -1 - if less
	// 0  - if equals
	compare(*TreeValue) int
}

type AVLTree struct {
	root *treeNode
}

func (tree *AVLTree) add(value *TreeValue) {
	tree.root = tree.root.add(value)
}

func (tree *AVLTree) remove(value *TreeValue) {
	tree.root = tree.root.remove(value)
}

func (tree *AVLTree) search(value *TreeValue) TreeValue {
	node := tree.root.search(value)
	if node == nil {
		return nil
	}

	return node.value
}

func (tree *AVLTree) print() {
	tree.root.printBFS()
}

type treeNode struct {
	value       TreeValue
	height      int
	left, right *treeNode
}

func (root *treeNode) getHeight() int {
	if root == nil {
		return 0
	}

	return root.height
}

func (root *treeNode) recalculateHeight() {
	if root == nil {
		return
	}

	root.height = 1 + max(root.left.getHeight(), root.right.getHeight())
}

func (root *treeNode) rotateRight() *treeNode {
	temp := root.left
	root.left = temp.right
	temp.right = root

	root.recalculateHeight()
	temp.recalculateHeight()

	return temp
}

func (root *treeNode) rotateLeft() *treeNode {
	temp := root.right
	root.right = temp.left
	temp.left = root

	root.recalculateHeight()
	temp.recalculateHeight()

	return temp
}

func (root *treeNode) rebalance() *treeNode {
	if root == nil {
		return nil
	}

	balance := root.right.getHeight() - root.left.getHeight()

	// balance =  2 - means right sub-tree bigger than left
	// balance = -2 - otherwise
	if balance == 2 {
		if root.right.left.getHeight() > root.right.right.getHeight() {
			root.right = root.right.rotateRight()
		}
		return root.rotateLeft()
	} else if balance == -2 {
		if root.left.right.getHeight() > root.left.left.getHeight() {
			root.left = root.left.rotateLeft()
		}
		return root.rotateRight()
	}

	return root
}

func (root *treeNode) search(value *TreeValue) *treeNode {
	if root == nil || value == nil {
		return nil
	}

	if root.value.compare(value) == 0 {
		return root
	}

	if root.left == nil && root.right == nil {
		return nil
	}

	if root.value.compare(value) == 1 {
		return root.left.search(value)
	} else {
		return root.right.search(value)
	}
}

func (root *treeNode) add(value *TreeValue) *treeNode {
	if root == nil {
		if value == nil {
			return nil
		} else {
			return &treeNode{*value, 1, nil, nil}
		}
	}

	if root.value.compare(value) == 0 {
		root.value = *value
	} else if root.value.compare(value) == 1 {
		root.left = root.left.add(value)
	} else {
		root.right = root.right.add(value)
	}

	root.recalculateHeight()
	return root.rebalance()
}

func (root *treeNode) remove(value *TreeValue) *treeNode {
	if root == nil || value == nil {
		return nil
	}

	if root.value.compare(value) == 1 {
		return root.left.remove(value)
	}

	if root.value.compare(value) == -1 {
		return root.right.remove(value)
	}

	if root.value.compare(value) == 0 {
		if root.left != nil && root.right != nil {
			smallest := root.right.findSmallest()
			root.value = smallest.value
			smallest.remove(value)
		} else if root.left != nil {
			root = root.left
		} else if root.right != nil {
			root = root.right
		} else {
			root = nil
		}
	}

	root.recalculateHeight()
	return root.rebalance()
}

func (root *treeNode) findSmallest() *treeNode {
	if root == nil || root.left == nil {
		return nil
	} else {
		return root.left.findSmallest()
	}
}

func (root *treeNode) printBFS() {
	queue := make([]*treeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		fmt.Print(pop.value)
		fmt.Print(" ")
		if pop.left != nil {
			queue = append(queue, pop.left)
		}
		if pop.right != nil {
			queue = append(queue, pop.right)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
