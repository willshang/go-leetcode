package main

import "fmt"

func main() {
	root := TreeNode{Val: 3}
	rootfirst := TreeNode{Val: 9}
	rootSecond := TreeNode{Val: 20}
	root.Left = &rootfirst
	root.Right = &rootSecond

	secondLeft := TreeNode{Val: 15}
	secondRight := TreeNode{Val: 7}
	rootSecond.Left = &secondLeft
	rootSecond.Right = &secondRight
	fmt.Println(sumOfLeftLeaves(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum = sum + node.Left.Val
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return sum
}
