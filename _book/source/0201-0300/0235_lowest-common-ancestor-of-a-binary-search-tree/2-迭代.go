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

	fmt.Println(lowestCommonAncestor(first, second, third))
	fmt.Println(first.Val)
	fmt.Println(first.Left.Val)
	fmt.Println(first.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}
