package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2
	left1 := TreeNode{}
	left1.Val = 22
	left2 := TreeNode{}
	left2.Val = 23

	right := TreeNode{}
	right.Val = 3
	right1 := TreeNode{}
	right1.Val = 33
	right2 := TreeNode{}
	right2.Val = 34

	root.Left = &left
	root.Right = &right

	left.Left = &left1
	left.Right = &left2

	right.Left = &right1
	right.Right = &right2
	res := binaryTreePaths(&root)

	for _, v := range res {
		fmt.Println(v)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	var queue []*TreeNode
	var stringQueue []string
	queue = append(queue, root)
	stringQueue = append(stringQueue, strconv.Itoa(root.Val))
	for len(queue) > 0 {
		node := queue[0]
		path := stringQueue[0]
		queue = queue[1:]
		stringQueue = stringQueue[1:]
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			stringQueue = append(stringQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			stringQueue = append(stringQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return res
}
