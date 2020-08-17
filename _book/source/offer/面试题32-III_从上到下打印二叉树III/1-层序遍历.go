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

// 剑指offer_面试题32-III_从上到下打印二叉树III
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	level := 0
	for len(list) > 0 {
		length := len(list)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := list[i]
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		if level%2 == 0 {
			for i := 0; i < length; i++ {
				temp = append(temp, list[i].Val)
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				temp = append(temp, list[i].Val)
			}
		}
		list = list[length:]
		res = append(res, temp)
		level++
	}
	return res
}
