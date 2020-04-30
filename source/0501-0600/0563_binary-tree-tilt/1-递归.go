package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	fmt.Println(findTilt(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode563_二叉树的坡度
var total int

func findTilt(root *TreeNode) int {
	total = 0
	dfs(root)
	return total
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	total = total + abs(left, right)
	return left + right + root.Val // 返回节点之和
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
