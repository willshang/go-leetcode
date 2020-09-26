package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 3

	root.Left = &left
	left.Right = &right
	res := minDepth(&root)
	fmt.Println(res)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	list := make([]*TreeNode, 0)
	list = append(list, root)
	depth := 1

	for len(list) > 0 {
		length := len(list)
		for i := 0; i < length; i++ {
			node := list[0]
			list = list[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		depth++
	}
	return depth
}
