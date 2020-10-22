package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	fmt.Println(kthSmallest(&first, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode230_二叉搜索树中第K小的元素
var res []int

func kthSmallest(root *TreeNode, k int) int {
	res = make([]int, 0)
	dfs(root)
	return res[k-1]
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}
