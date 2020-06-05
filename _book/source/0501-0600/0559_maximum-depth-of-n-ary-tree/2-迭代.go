package main

import "fmt"

func main() {
	root := &Node{}
	root.Val = 1

	right := &Node{}
	right.Val = 5

	rightLeft := &Node{}
	rightLeft.Val = 4

	root.Children = append(root.Children, right)
	root.Children = append(root.Children, rightLeft)

	fmt.Println(maxDepth(root))
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queue := make([]*Node, 0)
	depth := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			temp := queue[0]
			for _, node := range temp.Children {
				queue = append(queue, node)
			}
			queue = queue[1:]
		}
		depth++
	}
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
