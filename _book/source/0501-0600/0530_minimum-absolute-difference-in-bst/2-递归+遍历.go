package main

import (
	"fmt"
)

func main() {
	root := TreeNode{}
	root.Val = 1

	right := TreeNode{}
	right.Val = 5

	rightLeft := TreeNode{}
	rightLeft.Val = 4

	root.Right = &right
	right.Left = &rightLeft

	fmt.Println(getMinimumDifference(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode530_二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	arr := make([]int, 0)
	dfs(root, &arr)
	minDiff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if minDiff > arr[i]-arr[i-1] {
			minDiff = arr[i] - arr[i-1]
		}
	}
	return minDiff
}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs(root.Right, arr)
}
