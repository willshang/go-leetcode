package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(increasingBST(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode897_单调数列
var prev *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	prev = &TreeNode{}
	head := prev
	dfs(root)
	return head.Right
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	node.Left = nil
	prev.Right = node
	prev = node
	dfs(node.Right)
}
