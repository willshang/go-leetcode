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

// leetcode145_二叉树的后序遍历
var res []int

func postorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
}
