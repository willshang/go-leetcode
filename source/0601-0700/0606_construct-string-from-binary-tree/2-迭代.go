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
	stack := make([]*TreeNode, 0)
	m := make(map[*TreeNode]bool)
	stack = append(stack, t)
	res := ""
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if _, ok := m[node]; ok {
			stack = stack[:len(stack)-1]
			res = res + ")"
		} else {
			m[node] = true
			res = res + "(" + strconv.Itoa(node.Val)
			if node.Left == nil && node.Right != nil {
				res = res + "()"
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return res[1 : len(res)-1]
}
