package main

import "fmt"

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
	res := rightSideView(&root)
	fmt.Println(res)
}

// 剑指OfferII046.二叉树的右侧视图
var res []int

func rightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 1)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level > len(res) {
		res = append(res, root.Val)
	}
	dfs(root.Right, level+1)
	dfs(root.Left, level+1)
}
