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

var arr []int

func kthLargest(root *TreeNode, k int) int {
	arr = make([]int, 0)
	dfs(root)
	return arr[k-1]
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	arr = append(arr, root.Val)
	dfs(root.Left)
}
