package main

import "bufio"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1145_二叉树着色游戏
var leftSum, rightSum int

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	leftSum = 0
	rightSum = 0
	total := dfs(root, x)
	return leftSum > n/2 || rightSum > n/2 || (total-1-leftSum-rightSum) > n/2
}

func dfs(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, x)
	right := dfs(root.Right, x)
	if root.Val == x {
		leftSum = left
		rightSum = right
	}
	return 1 + left + right
}
