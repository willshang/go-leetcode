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
	fmt.Println(minDiffInBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode783_二叉搜索树节点最小距离
func minDiffInBST(root *TreeNode) int {
	arr := make([]int, 0)
	dfs(root, &arr)
	min := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if min > arr[i]-arr[i-1] {
			min = arr[i] - arr[i-1]
		}
	}
	return min
}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs(root.Right, arr)
}
