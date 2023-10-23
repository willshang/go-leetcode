package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1372_二叉树中的最长交错路径
var res int

func longestZigZag(root *TreeNode) int {
	res = 0
	dfs(root, 0, 0)
	return res
}

func dfs(root *TreeNode, left, right int) {
	if root != nil {
		res = max(res, left)
		res = max(res, right)
		dfs(root.Left, right+1, 0)
		dfs(root.Right, 0, left+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
