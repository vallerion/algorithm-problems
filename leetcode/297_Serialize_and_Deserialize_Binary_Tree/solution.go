package _297_Serialize_and_Deserialize_Binary_Tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	queue, stack []*TreeNode
}

func Constructor() Codec {
	return Codec{queue: make([]*TreeNode, 0), stack: make([]*TreeNode, 0)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	this.queue = append(this.queue, root)
	result := ""

	for len(this.queue) > 0 {
		n := len(this.queue)

		for i := 0; i < n; i++ {
			curr := this.queue[i]

			if curr == nil {
				result += "null,"
				continue
			}

			result += strconv.Itoa(curr.Val) + ","

			this.queue = append(this.queue, curr.Left)
			this.queue = append(this.queue, curr.Right)
		}

		this.queue = this.queue[n:]
	}

	return strings.TrimRight(result, ",")
}

// Deserializes your encoded data to tree.
// 1,2,3,null,null,4,5
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || (len(data) == 2 && data == "[]") {
		return nil
	}

	items := strings.Split(data, ",")
	rootValue, _ := strconv.Atoi(items[0])
	root := &TreeNode{rootValue, nil, nil}

	this.stack = append(this.stack, root)

	i := 1
	for len(this.stack) > 0 {
		if this.stack[0].Left == nil {
			if items[i] == "null" {
				i++
			} else {
				nodeValue, _ := strconv.Atoi(items[i])
				this.stack[0].Left = &TreeNode{nodeValue, nil, nil}
				this.stack = append(this.stack, this.stack[0].Left)
				i++
			}
		}
		if this.stack[0].Right == nil {
			if items[i] == "null" {
				i++
			} else {
				nodeValue, _ := strconv.Atoi(items[i])
				this.stack[0].Right = &TreeNode{nodeValue, nil, nil}
				this.stack = append(this.stack, this.stack[0].Right)
				i++
			}
		}

		this.stack = this.stack[1:]
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
