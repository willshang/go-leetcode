package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func largestValues(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level >= len(res) {
		res = append(res, math.MinInt32)
	}
	res[level] = max(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
