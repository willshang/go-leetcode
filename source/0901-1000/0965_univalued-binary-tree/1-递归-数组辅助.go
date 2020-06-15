package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(isUnivalTree(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	arr = make([]int, 0)
	dfs(root)
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode) {
	if root != nil {
		arr = append(arr, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}
