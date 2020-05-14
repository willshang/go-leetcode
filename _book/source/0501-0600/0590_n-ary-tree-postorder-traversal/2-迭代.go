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
	fmt.Println(postorder(root))
}

type Node struct {
	Val      int
	Children []*Node
}

// 后序：(左右)根
// 前序：根(左右)=>根(右左)=>左右根
func postorder(root *Node) []int {
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
		for i := 0; i < len(temp.Children); i++ {
			stack = append(stack, temp.Children[i])
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
