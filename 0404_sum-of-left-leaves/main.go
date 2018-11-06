package main

import "fmt"

func main() {
	root := TreeNode{Val:1}
	rootfirst := TreeNode{Val:2}
	rootSecond := TreeNode{Val:3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(sumOfLeftLeaves(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Left == nil{
		return sumOfLeftLeaves(root.Right)
	}
	if root.Left.Left == nil && root.Left.Right == nil{
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
