package _1382_Balance_a_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]*TreeNode, 0)
	treeToArray(root, &nums)
	return arrayToTree(&nums, 0, len(nums)-1)
}

func treeToArray(root *TreeNode, nums *[]*TreeNode) {
	if root == nil {
		return
	}

	treeToArray(root.Left, nums)
	*nums = append(*nums, root)
	treeToArray(root.Right, nums)
}

func arrayToTree(nums *[]*TreeNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	q := start + ((end - start) / 2)

	node := (*nums)[q]
	node.Left = arrayToTree(nums, start, q-1)
	node.Right = arrayToTree(nums, q+1, end)

	return node
}
