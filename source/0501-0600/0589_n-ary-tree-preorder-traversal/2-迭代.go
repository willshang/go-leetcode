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

	fmt.Println(preorder(root))
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, temp.Val)
		for i := len(temp.Children) - 1; i >= 0; i-- {
			stack = append(stack, temp.Children[i])
		}
	}
	return res
}
