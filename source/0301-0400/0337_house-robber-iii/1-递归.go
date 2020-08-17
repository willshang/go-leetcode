package main

import (
	"fmt"
)

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(rob(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode337_打家劫舍III
func rob(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftA, leftB := dfs(root.Left)
	rightA, rightB := dfs(root.Right)
	a := root.Val + leftB + rightB               // A=>偷
	b := max(leftA, leftB) + max(rightA, rightB) // B=>不偷
	return a, b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
