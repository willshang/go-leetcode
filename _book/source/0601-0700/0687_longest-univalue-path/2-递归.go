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

var maxLen int

func longestUnivaluePath(root *TreeNode) int {
	maxLen = 0
	if root == nil {
		return 0
	}
	dfs(root, root.Val)
	return maxLen
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, root.Val)
	right := dfs(root.Right, root.Val)
	if left+right > maxLen {
		maxLen = left + right
	}
	if root.Val == val {
		return max(left, right) + 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
