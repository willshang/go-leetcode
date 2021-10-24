package main

import (
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode655_输出二叉树
var res [][]string

func printTree(root *TreeNode) [][]string {
	h := getHeightDFS(root)
	w := (1 << h) - 1
	res = make([][]string, h)
	for i := 0; i < h; i++ {
		res[i] = make([]string, w)
	}
	dfs(root, 0, 0, w-1)
	return res
}

func dfs(root *TreeNode, h, left, right int) {
	if root == nil {
		return
	}
	mid := left + (right-left)/2
	res[h][mid] = strconv.Itoa(root.Val)
	dfs(root.Left, h+1, left, mid-1)
	dfs(root.Right, h+1, mid+1, right)
}

func getHeightDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeightDFS(root.Left), getHeightDFS(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
