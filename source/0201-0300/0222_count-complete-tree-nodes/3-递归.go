package main

import (
	"fmt"
)

func main() {
	first := &TreeNode{Val: 6}
	second := &TreeNode{Val: 2}
	third := &TreeNode{Val: 8}
	first.Left = second
	first.Right = third

	fmt.Println(countNodes(first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getLevel(root.Left)
	right := getLevel(root.Right)
	if left == right {
		return 1<<left + countNodes(root.Right)
	}
	return countNodes(root.Left) + 1<<right
}

func getLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}
