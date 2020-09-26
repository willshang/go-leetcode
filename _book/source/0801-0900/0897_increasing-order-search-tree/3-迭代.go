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

func increasingBST(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	newRoot := &TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
			continue
		}
		stack = stack[:len(stack)-1]
		node.Right = newRoot.Right
		newRoot.Right = node
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}
	return newRoot.Right
}
