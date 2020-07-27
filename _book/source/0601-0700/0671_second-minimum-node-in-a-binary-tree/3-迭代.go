package main

import (
	"fmt"
	"math"
)

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 2}

	first.Left = &third
	first.Right = &second

	fmt.Println(findSecondMinimumValue(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	min, second := root.Val, math.MaxInt32
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	flag := 0
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node.Val < min {
			second = min
			min = node.Val
		} else if min < node.Val && node.Val <= second {
			flag = 1
			second = node.Val
		}
		if node.Left != nil {
			// 有0个或2节点
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	if second == math.MaxInt32 && flag == 0 {
		return -1
	}
	return second
}
