package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third
	fmt.Println(levelOrder(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指offer_面试题32-II_从上到下打印二叉树II
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
		list = list[length:]
		res = append(res, temp)
	}
	return res
}
