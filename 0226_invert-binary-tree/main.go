package main

import "fmt"

func main() {
	first := TreeNode{Val:1}
	second := TreeNode{Val:2}
	third := TreeNode{Val:3}
	first.Left = &second
	first.Right = &third

	invertTree(&first)
	fmt.Println(first.Val)
	fmt.Println(first.Left.Val)
	fmt.Println(first.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil){
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
