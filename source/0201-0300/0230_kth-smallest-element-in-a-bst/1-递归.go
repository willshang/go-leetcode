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

var res int
var index int

func kthSmallest(root *TreeNode, k int) int {
	res = 0
	index = k
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		index--
		if index == 0 {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
}
