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

// leetcode590_N叉树的后序遍历
var res []int

func postorder(root *Node) []int {
	res = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	for _, value := range root.Children {
		dfs(value)
	}
	res = append(res, root.Val)
}
