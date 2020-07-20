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

// leetcode106_从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	last := len(postorder) - 1
	for k := range inorder {
		if inorder[k] == postorder[last] {
			return &TreeNode{
				Val:   postorder[last],
				Left:  buildTree(inorder[0:k], postorder[0:k]),
				Right: buildTree(inorder[k+1:], postorder[k:last]),
			}
		}
	}
	return nil
}
