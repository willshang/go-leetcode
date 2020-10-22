package main

import (
	"fmt"
	"math"
)

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(isValidBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if left >= root.Val || right <= root.Val {
		return false
	}
	return dfs(root.Left, left, root.Val) && dfs(root.Right, root.Val, right)
}
