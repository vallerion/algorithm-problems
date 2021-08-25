package _230_Kth_Smallest_Element_in_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

func kthSmallestTreeToArray(root *TreeNode, k int) int {
	arr := make([]int, 0)
	treeToArray(root, &arr)
	return arr[k-1]
}

func treeToArray(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	treeToArray(root.Left, arr)
	*arr = append(*arr, root.Val)
	treeToArray(root.Right, arr)
}
