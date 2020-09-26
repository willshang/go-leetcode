package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(isCousins(&first, 3, 7))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return true
	}
	fatherM := make(map[int]int)
	levelM := make(map[int]int)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			levelM[node.Val] = level
			if node.Left != nil {
				fatherM[node.Left.Val] = node.Val
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				fatherM[node.Right.Val] = node.Val
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
		level++
	}
	return levelM[x] == levelM[y] && fatherM[x] != fatherM[y]
}
