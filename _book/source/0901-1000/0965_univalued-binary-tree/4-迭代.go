package main

import (
	"fmt"
)

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(isUnivalTree(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	value := root.Val
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node.Val != value {
			return false
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return true
}
