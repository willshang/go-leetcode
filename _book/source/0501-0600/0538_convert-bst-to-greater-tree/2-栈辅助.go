package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	right := TreeNode{}
	right.Val = 5

	rightLeft := TreeNode{}
	rightLeft.Val = 4

	root.Right = &right
	right.Left = &rightLeft

	fmt.Println(convertBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0)
	temp := root
	sum := 0
	for {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.Right
		} else if len(stack) != 0 {
			temp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temp.Val = temp.Val + sum
			sum = temp.Val
			temp = temp.Left
		} else {
			break
		}
	}
	return root
}
