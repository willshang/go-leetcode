package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1448_统计二叉树中好节点的数目
func goodNodes(root *TreeNode) int {
	maxValue := math.MinInt32
	return dfs(root, maxValue)
}

func dfs(root *TreeNode, maxValue int) int {
	if root == nil {
		return 0
	}
	if root.Val >= maxValue {
		return dfs(root.Left, root.Val) + dfs(root.Right, root.Val) + 1
	}
	return dfs(root.Left, maxValue) + dfs(root.Right, maxValue)
}
