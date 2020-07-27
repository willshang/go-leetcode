package main

import "fmt"

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	return helper(inorder, postorder)
}

func helper(inorder []int, postorder []int) *TreeNode {
	var root *TreeNode
	last := len(postorder) - 1
	for k := range inorder {
		if inorder[k] == postorder[last] {
			root = &TreeNode{Val: postorder[last]}
			root.Left = helper(inorder[0:k], postorder[0:k])
			root.Right = helper(inorder[k+1:], postorder[k:last])
		}
	}
	return root
}
