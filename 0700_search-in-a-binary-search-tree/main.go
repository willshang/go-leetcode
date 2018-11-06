package main

import "fmt"

func main() {
	first := TreeNode{Val:2}
	second := TreeNode{Val:1}
	third := TreeNode{Val:3}
	first.Left = &second
	first.Right = &third

	fmt.Println(searchBST(&first,3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return nil
	}
	switch {
	case root.Val < val:
		return searchBST(root.Right,val)
	case root.Val > val:
		return searchBST(root.Left,val)
	default:
		return root
	}
}
