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
	right.Val = 2
	root.Left = &left
	left.Right = &right
	res := rightSideView(&root)
	fmt.Println(res)
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		res = append(res, list[0].Val)
		for i := 0; i < length; i++ {
			node := list[i]
			if node.Right != nil {
				list = append(list, node.Right)
			}
			if node.Left != nil {
				list = append(list, node.Left)
			}
		}
		list = list[length:]
	}
	return res
}
