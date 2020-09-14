package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode958_二叉树的完全性检验
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	prev := root
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if prev == nil && node != nil {
			return false
		}
		if node != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		prev = node
	}
	return true
}
