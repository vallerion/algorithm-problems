package _1339_Maximum_Product_of_Splitted_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const mod = 1e9 + 7

var nodesSum = map[*TreeNode]int{}

func maxProduct(root *TreeNode) int {
	total := countTotal(root)
	max := 0
	calcProduct(root, total, &max)
	return max % mod
}

func countTotal(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodesSum[root] = root.Val + countTotal(root.Left) + countTotal(root.Right)

	return nodesSum[root]
}

func calcProduct(root *TreeNode, total int, max *int) {
	if root == nil {
		return
	}

	calcProduct(root.Left, total, max)
	calcProduct(root.Right, total, max)

	subtotal := nodesSum[root]
	calc := (total - subtotal) * subtotal
	if *max < calc {
		*max = calc
	}
}
