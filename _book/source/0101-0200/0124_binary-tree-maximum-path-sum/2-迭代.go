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

var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	stack := make([]*TreeNode, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = max(res, node.Val)
		var left, right int
		if node.Left == nil {
			left = 0
		} else {
			left = max(node.Left.Val, 0)
		}
		if node.Right == nil {
			right = 0
		} else {
			right = max(node.Right.Val, 0)
		}
		sum := node.Val + left + right
		res = max(res, sum)
		node.Val = node.Val + max(left, right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
