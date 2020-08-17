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

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	list := make([]*TreeNode, 0)
	total := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		list = append([]*TreeNode{node}, list...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	for i := range list {
		node := list[i]
		left := 0
		right := 0
		if node.Left != nil {
			left = node.Left.Val
		}
		if node.Right != nil {
			right = node.Right.Val
		}
		total = total + abs(left, right)
		node.Val = left + right + node.Val
	}
	return total
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
