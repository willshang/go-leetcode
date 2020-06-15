package main

import (
	"fmt"
)

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(sumRootToLeaf(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1022_从根到叶的二进制数之和
var res int

func sumRootToLeaf(root *TreeNode) int {
	res = 0
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum*2 + root.Val
	if root.Left == nil && root.Right == nil {
		res = (res + sum) % 1000000007
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}
