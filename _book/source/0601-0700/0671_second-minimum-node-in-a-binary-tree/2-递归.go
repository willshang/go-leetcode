package main

import (
	"fmt"
)

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 2}
	first.Left = &third
	first.Right = &second
	fmt.Println(findSecondMinimumValue(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	return dfs(root, root.Val)
}

// leetcode671_二叉树中第二小的节点
func dfs(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}
	if root.Val > val {
		return root.Val
	}
	left := dfs(root.Left, val)
	right := dfs(root.Right, val)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	return min(left, right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
