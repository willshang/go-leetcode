package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(sumNumbers(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指OfferII049.从根节点到叶节点的路径数字之和
var res int

func sumNumbers(root *TreeNode) int {
	res = 0
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		res = res + sum
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}
