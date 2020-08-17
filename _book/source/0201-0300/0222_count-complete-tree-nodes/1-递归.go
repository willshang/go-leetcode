package main

import "fmt"

func main() {
	first := &TreeNode{Val: 6}
	second := &TreeNode{Val: 2}
	third := &TreeNode{Val: 8}
	first.Left = second
	first.Right = third

	fmt.Println(countNodes(first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode222_完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
