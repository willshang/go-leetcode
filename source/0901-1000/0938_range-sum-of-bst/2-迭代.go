package main

import "fmt"

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 1}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	fmt.Println(rangeSumBST(&first, 1, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	if root.Val > R && root.Left != nil {
		stack = append(stack, root.Left)
	} else if root.Val < L && root.Right != nil {
		stack = append(stack, root.Right)
	} else {
		stack = append(stack, root)
	}
	res := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val <= R && node.Val >= L {
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			res = res + node.Val
		} else if node.Val > R && node.Left != nil {
			stack = append(stack, node.Left)
		} else if node.Val < L && node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return res
}
