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

	fmt.Println(diameterOfBinaryTree(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}

func helper(root *TreeNode) (length,diameter int ) {
	if root == nil{
		return
	}
	leftLen, leftDia := helper(root.Left)
	rightLen, rightDia := helper(root.Right)

	length = max(leftLen,rightLen) + 1
	diameter = max(leftLen+rightLen+1,max(leftDia,rightDia))
	return
}

func max(a, b int) int  {
	if a > b{
		return a
	}
	return b
}