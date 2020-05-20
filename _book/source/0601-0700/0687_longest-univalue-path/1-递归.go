package main

import "fmt"

func main() {
	root := TreeNode{Val: 5}
	rootLeft := TreeNode{Val: 4}
	rootRight := TreeNode{Val: 5}
	rootLeftLeft := TreeNode{Val: 1}
	rootLeftRight := TreeNode{Val: 1}
	rootRightRight := TreeNode{Val: 5}

	root.Left = &rootLeft
	root.Right = &rootRight
	rootLeft.Left = &rootLeftLeft
	rootLeft.Right = &rootLeftRight
	rootRight.Right = &rootRightRight

	fmt.Println(longestUnivaluePath(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode687_最长同值路径
var maxLen int

func longestUnivaluePath(root *TreeNode) int {
	maxLen = 0
	dfs(root)
	return maxLen
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	l, r := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		l = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		r = right + 1
	}
	if l+r > maxLen {
		maxLen = l + r
	}
	return max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
