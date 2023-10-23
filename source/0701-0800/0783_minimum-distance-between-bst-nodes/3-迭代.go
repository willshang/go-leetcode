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

func minDiffInBST(root *TreeNode) int {
	arr := make([]int, 0)
	stack := make([]*TreeNode, 0)
	min := math.MaxInt32
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		arr = append(arr, node.Val)
		if len(arr) > 1 {
			temp := node.Val - arr[len(arr)-2]
			if min > temp {
				min = temp
			}
		}
		root = node.Right
	}
	return min
}
