package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	left.Right = &right
	res := maxPathSum(&root)
	fmt.Println(res)
}

// 剑指OfferII051.节点之和最大的路径
var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	// 该顶点路径和=root.Val+2边和
	value := left + right + root.Val
	res = max(res, value)
	// 单分支
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
