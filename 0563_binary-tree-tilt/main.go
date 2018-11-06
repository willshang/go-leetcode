package main

import "fmt"

func main() {
	first := TreeNode{Val:1}
	second := TreeNode{Val:2}
	third := TreeNode{Val:3}
	first.Left = &second
	first.Right = &third

	fmt.Println(findTilt(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var total int
	helper(root,&total)
	return total
}

func helper(root *TreeNode, total *int)int  {
	if root == nil{
		return 0
	}

	l := helper(root.Left,total)
	r := helper(root.Right,total)

	if l > r {
		*total = *total + l - r
	}else {
		*total = *total + r - l
	}

	return l + r + root.Val
}
