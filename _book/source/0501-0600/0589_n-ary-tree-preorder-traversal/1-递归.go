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

// leetcode589_N叉树的前序遍历
var res []int

func preorder(root *Node) []int {
	res = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	for _, value := range root.Children {
		dfs(value)
	}
}
