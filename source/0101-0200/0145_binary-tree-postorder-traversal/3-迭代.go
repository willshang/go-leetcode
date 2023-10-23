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
	res := postorderTraversal(&root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last != nil {
			stack = append(stack, last)
			stack = append(stack, nil)
			if last.Right != nil {
				stack = append(stack, last.Right)
			}
			if last.Left != nil {
				stack = append(stack, last.Left)
			}
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
		}
	}
	return res
}
