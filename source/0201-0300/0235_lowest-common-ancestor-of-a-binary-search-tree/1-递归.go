package main

import "fmt"

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

// leetcode235_二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}
