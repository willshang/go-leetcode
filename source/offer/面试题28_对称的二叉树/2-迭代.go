package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	root.Right = &right

	fmt.Println(isSymmetric(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	leftQ := make([]*TreeNode, 0)
	rightQ := make([]*TreeNode, 0)
	leftQ = append(leftQ, root)
	rightQ = append(rightQ, root)

	for len(leftQ) != 0 && len(rightQ) != 0 {
		leftCur, rightCur := leftQ[0], rightQ[0]
		leftQ, rightQ = leftQ[1:], rightQ[1:]

		if leftCur == nil && rightCur == nil {
			continue
		} else if leftCur != nil && rightCur != nil && leftCur.Val == rightCur.Val {
			leftQ = append(leftQ, leftCur.Left, leftCur.Right)
			rightQ = append(rightQ, rightCur.Right, rightCur.Left)
		} else {
			return false
		}
	}

	if len(leftQ) == 0 && len(rightQ) == 0 {
		return true
	} else {
		return false
	}
}
