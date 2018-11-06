package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	left.Right = &right
	res := hasPathSum(&root,5)
	fmt.Println(res)
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	sum = sum - root.Val

	if root.Left == nil && root.Right == nil{
		return sum == 0
	}

	return hasPathSum(root.Left,sum) || hasPathSum(root.Right,sum)
}