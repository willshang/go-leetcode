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
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}
	res = append(res, root.Val)
	return res
}
