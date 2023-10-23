package main

import "fmt"

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

var value int
var res bool

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res = true
	value = root.Val
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		if root.Val != value {
			res = false
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
}
