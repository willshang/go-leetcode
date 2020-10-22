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
	if postorder == nil || len(postorder) == 0 {
		return nil
	}
	last := len(postorder) - 1
	root := &TreeNode{
		Val: postorder[last],
	}
	length := len(postorder)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	index := last
	for i := length - 2; i >= 0; i-- {
		value := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[index] {
			node.Right = &TreeNode{Val: value}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index--
			}
			node.Left = &TreeNode{Val: value}
			stack = append(stack, node.Left)
		}
	}
	return root
}
