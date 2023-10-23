package main

import "fmt"

func main() {
	root := &Node{}
	root.Val = 1

	right := &Node{}
	right.Val = 5

	rightLeft := &Node{}
	rightLeft.Val = 4

	root.Children = append(root.Children, right)
	root.Children = append(root.Children, rightLeft)

	fmt.Println(maxDepth(root))
}

type Node struct {
	Val      int
	Children []*Node
}

// leetcode559_N叉树的最大深度
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for _, node := range root.Children {
		depth = max(depth, maxDepth(node))
	}
	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
