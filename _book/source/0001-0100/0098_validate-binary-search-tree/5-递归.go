package main

import (
	"fmt"
	"math"
)

func main() {
	root := TreeNode{Val: 0}
	//rootfirst := TreeNode{Val: 2}
	//rootSecond := TreeNode{Val: 3}
	//root.Left = &rootfirst
	//root.Right = &rootSecond
	fmt.Println(isValidBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int

func isValidBST(root *TreeNode) bool {
	pre = math.MinInt64
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if dfs(root.Left) == false {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return dfs(root.Right)
}
