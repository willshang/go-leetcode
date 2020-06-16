package main

import (
	"fmt"
)

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 4}
	secondRight := TreeNode{Val: 3}

	first.Left = &second
	first.Right = &secondRight
	second.Left = &third

	first1 := TreeNode{Val: 10}
	second1 := TreeNode{Val: 20}
	third1 := TreeNode{Val: 40}
	secondRight1 := TreeNode{Val: 30}

	first1.Left = &second1
	first1.Right = &secondRight1
	second1.Left = &third1

	fmt.Println(mergeTrees(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	list := make([]*TreeNode, 0)
	list = append(list, t1)
	list = append(list, t2)
	for len(list) > 0 {
		node1 := list[0]
		node2 := list[1]
		node1.Val = node1.Val + node2.Val
		if node1.Left != nil && node2.Left != nil {
			list = append(list, node1.Left)
			list = append(list, node2.Left)
		} else if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		if node1.Right != nil && node2.Right != nil {
			list = append(list, node1.Right)
			list = append(list, node2.Right)
		} else if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
		list = list[2:]
	}
	return t1
}
