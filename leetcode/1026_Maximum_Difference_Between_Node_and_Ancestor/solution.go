package _1026_Maximum_Difference_Between_Node_and_Ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, []int{}, 0)
}

func dfs(root *TreeNode, ancestors []int, maxDiff int) int {
	if root == nil {
		return -1
	}

	for i := 0; i < len(ancestors); i++ {
		if abs(ancestors[i]-root.Val) > maxDiff {
			maxDiff = abs(ancestors[i] - root.Val)
		}
	}

	ancestors = append(ancestors, root.Val)
	l := dfs(root.Left, ancestors, maxDiff)
	r := dfs(root.Right, ancestors, maxDiff)
	if l > maxDiff {
		maxDiff = l
	}
	if r > maxDiff {
		maxDiff = r
	}

	return maxDiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
