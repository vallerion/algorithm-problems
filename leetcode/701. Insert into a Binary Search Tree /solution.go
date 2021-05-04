package _01__Insert_into_a_Binary_Search_Tree_

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	if root.Val == val {
		return root
	}

	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Left, val)
		}
	}

	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		res := new(TreeNode)
		res.Val = val
		return res
	}

	if root.Val == val {
		return root
	}

	if val > root.Val {
			root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}