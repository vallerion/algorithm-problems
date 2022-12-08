package _872_Leaf_Similar_Trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n+m)
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1, leafs2 := make([]int, 0), make([]int, 0)
	retrieveLeafs(root1, &leafs1)
	retrieveLeafs(root2, &leafs2)
	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}

func retrieveLeafs(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	}
	retrieveLeafs(root.Left, leafs)
	retrieveLeafs(root.Right, leafs)
}
