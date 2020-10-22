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

// leetcode965_单值二叉树
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Left.Val != root.Val) ||
		(root.Right != nil && root.Right.Val != root.Val) {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
