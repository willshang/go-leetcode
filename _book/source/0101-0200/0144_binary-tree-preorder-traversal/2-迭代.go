package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	left.Right = &right
	res := preorderTraversal(&root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		last := len(stack) - 1
		root = stack[last]
		stack = stack[:last]
	}
	return res
}
