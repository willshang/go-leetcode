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
	root.Right = &right

	fmt.Println(isSymmetric(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return recur(root.Left,root.Right)
}

func recur(left, right *TreeNode) bool  {
	if left == nil && right == nil{
		return true
	}
	if left == nil || right == nil{
		return false
	}

	return left.Val == right.Val && recur(left.Left,right.Right) && recur(left.Right,right.Left)
}
