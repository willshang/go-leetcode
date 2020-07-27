package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(isCousins(&first, 3, 7))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	xNode, xDepth := dfs(root, x, 0)
	yNode, yDepth := dfs(root, y, 0)
	return xDepth == yDepth && xNode != yNode
}

func dfs(root *TreeNode, value int, level int) (*TreeNode, int) {
	if root == nil {
		return nil, -1
	}
	if root.Val == value {
		return nil, 0
	}
	if (root.Left != nil && root.Left.Val == value) ||
		(root.Right != nil && root.Right.Val == value) {
		return root, level + 1
	}
	leftNode, leftLevel := dfs(root.Left, value, level+1)
	if leftNode != nil {
		return leftNode, leftLevel
	}
	return dfs(root.Right, value, level+1)
}
