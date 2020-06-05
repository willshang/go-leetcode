package main

import "fmt"

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 1}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	fmt.Println(searchBST(&first, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val == val {
			return node
		} else if node.Val > val && node.Left != nil {
			stack = append(stack, node.Left)
		} else if node.Val < val && node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return nil
}
