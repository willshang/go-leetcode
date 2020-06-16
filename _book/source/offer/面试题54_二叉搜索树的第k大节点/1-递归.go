package main

import "fmt"

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 1}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	fmt.Println(kthLargest(&first, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指offer_面试题54_二叉搜索树的第k大节点
var count int
var res int

func kthLargest(root *TreeNode, k int) int {
	count = k
	res = 0
	dfs(root)
	return res
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	count--
	if count == 0 {
		res = root.Val
		return
	}
	dfs(root.Left)
}
