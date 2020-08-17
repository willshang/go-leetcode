package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(levelOrder(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode102_二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := list[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		res = append(res, temp)
		list = list[length:]
	}
	return res
}
