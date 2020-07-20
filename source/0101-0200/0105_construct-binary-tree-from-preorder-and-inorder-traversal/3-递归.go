package main

import "fmt"

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	return helper(preorder, inorder)
}

func helper(preorder []int, inorder []int) *TreeNode {
	var root *TreeNode
	for k := range inorder {
		if inorder[k] == preorder[0] {
			root = &TreeNode{Val: preorder[0]}
			root.Left = helper(preorder[1:k+1], inorder[0:k])
			root.Right = helper(preorder[k+1:], inorder[k+1:])
		}
	}
	return root
}
