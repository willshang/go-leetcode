package main

import (
	"fmt"
)

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(sumNumbers(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		for i := 0; i < length; i++ {
			node := list[i]
			value := node.Val
			if node.Left == nil && node.Right == nil {
				res = res + value
			}
			if node.Left != nil {
				node.Left.Val = node.Left.Val + value*10
				list = append(list, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Right.Val + value*10
				list = append(list, node.Right)
			}
		}
		list = list[length:]
	}
	return res
}
