package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1339_分裂二叉树的最大乘积
var sum int
var res int

func maxProduct(root *TreeNode) int {
	sum = 0
	res = 0
	dfsSum(root)
	dfs(root)
	return res % 1000000007
}

func dfsSum(root *TreeNode) {
	if root == nil {
		return
	}
	sum = sum + root.Val
	dfsSum(root.Left)
	dfsSum(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res = max(res, left*(sum-left))
	res = max(res, right*(sum-right))
	return left + right + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
