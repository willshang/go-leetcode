package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1373_二叉搜索子树的最大键值和
var res int

func maxSumBST(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) (bool, int, int, int) {
	if root == nil {
		return true, 0, math.MaxInt32, math.MinInt32
	}

	leftOk, leftValue, leftMin, leftMax := dfs(root.Left)
	rightOk, rightValue, rightMin, rightMax := dfs(root.Right)
	if leftOk == false || rightOk == false || root.Val <= leftMax || root.Val >= rightMin {
		return false, 0, 0, 0
	}
	sum := root.Val + leftValue + rightValue
	res = max(res, sum)
	return true, sum, min(root.Val, leftMin), max(root.Val, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
