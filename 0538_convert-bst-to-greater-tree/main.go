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
	sum := 0
	travel(root,&sum)
	return root
}

func travel(root *TreeNode, sum *int)  {
	if root == nil{
		return
	}
	travel(root.Right,sum)
	*sum = *sum + root.Val
	root.Val = *sum
	travel(root.Left,sum)
}
