package main

import (
	"fmt"
	"math"
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

var minDiff, previous int

func minDiffInBST(root *TreeNode) int {
	minDiff, previous = math.MaxInt32, math.MaxInt32
	dfs(root)
	return minDiff
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	newDiff := diff(previous, root.Val)
	if minDiff > newDiff {
		minDiff = newDiff
	}
	previous = root.Val
	dfs(root.Right)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
