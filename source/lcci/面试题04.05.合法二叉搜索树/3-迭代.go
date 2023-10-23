package main

import (
	"fmt"
	"math"
)

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(isValidBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		res = append(res, stack[last].Val)
		root = stack[last].Right
		stack = stack[:last]
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}
