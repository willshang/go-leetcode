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

	right := TreeNode{}
	right.Val = 3

	root.Left = &left
	root.Right = &right
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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := make([]string, 0, 16)
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
