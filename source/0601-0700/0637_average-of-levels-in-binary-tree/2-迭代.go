package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 4}
	secondRight := TreeNode{Val: 3}

	first.Left = &second
	first.Right = &secondRight
	second.Left = &third
	fmt.Println(averageOfLevels(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode637_二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		sum := 0
		for i := 0; i < length; i++ {
			sum = sum + list[i].Val
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		res = append(res, float64(sum)/float64(length))
		list = list[length:]
	}
	return res
}
