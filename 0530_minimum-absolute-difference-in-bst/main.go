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

	fmt.Println(getMinimumDifference(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type state struct {
	minDiff, previous int
}
func getMinimumDifference(root *TreeNode) int {
	st := state{1024,1024}
	search(root, &st)
	return st.minDiff
}

func search(root *TreeNode, st *state)  {
	if root == nil{
		return
	}
	search(root.Left,st)

	newDiff := diff(st.previous,root.Val)
	if st.minDiff > newDiff{
		st.minDiff = newDiff
	}
	st.previous = root.Val
	search(root.Right,st)
}

func diff(a,b int)int  {
	if a > b {
		return a - b
	}
	return b - a
}
