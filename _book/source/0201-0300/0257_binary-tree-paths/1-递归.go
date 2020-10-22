package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2
	left1 := TreeNode{}
	left1.Val = 22
	left2 := TreeNode{}
	left2.Val = 23

	right := TreeNode{}
	right.Val = 3
	right1 := TreeNode{}
	right1.Val = 33
	right2 := TreeNode{}
	right2.Val = 34

	root.Left = &left
	root.Right = &right

	left.Left = &left1
	left.Right = &left2

	right.Left = &right1
	right.Right = &right2
	res := binaryTreePaths(&root)

	for _, v := range res {
		fmt.Println(v)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode257_二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := make([]string, 0)
	var dfs func(string, *TreeNode)
	dfs = func(pre string, root *TreeNode) {
		if pre == "" {
			pre = strconv.Itoa(root.Val)
		} else {
			pre += "->" + strconv.Itoa(root.Val)
		}

		if root.Left != nil {
			dfs(pre, root.Left)
		}

		if root.Right != nil {
			dfs(pre, root.Right)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, pre)
		}
	}

	dfs("", root)
	return res
}
