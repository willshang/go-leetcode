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
	res := isBalanced(&root)
	fmt.Println(res)
}

// 程序员面试金典04.04_检查平衡性
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(dfs(root.Left)-dfs(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func dfs(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(dfs(root.Left), dfs(root.Right)) + 1
}
