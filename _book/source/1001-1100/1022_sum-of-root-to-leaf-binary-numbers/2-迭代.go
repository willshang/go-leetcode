package main

import (
	"fmt"
)

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(sumRootToLeaf(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node *TreeNode
	sum  int
}

func sumRootToLeaf(root *TreeNode) int {
	res := 0
	stack := make([]Node, 0)
	stack = append(stack, Node{
		node: root,
		sum:  0,
	})
	for len(stack) > 0 {
		node, sum := stack[len(stack)-1].node, stack[len(stack)-1].sum
		stack = stack[:len(stack)-1]
		sum = sum*2 + node.Val
		if node.Left == nil && node.Right == nil {
			res = (res + sum) % 1000000007
		}
		if node.Left != nil {
			stack = append(stack, Node{
				node: node.Left,
				sum:  sum,
			})
		}
		if node.Right != nil {
			stack = append(stack, Node{
				node: node.Right,
				sum:  sum,
			})
		}
	}
	return res
}
