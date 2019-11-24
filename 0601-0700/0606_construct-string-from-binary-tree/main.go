package main

import (
	"fmt"
	"strconv"
)

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 4}
	secondRight := TreeNode{Val: 3}

	first.Left = &second
	first.Right = &secondRight
	second.Left = &third

	fmt.Println(tree2str(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	res := strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		return res
	}

	res += "(" + tree2str(t.Left) + ")"

	if t.Right != nil {
		res += "(" + tree2str(t.Right) + ")"
	}

	return res
}
