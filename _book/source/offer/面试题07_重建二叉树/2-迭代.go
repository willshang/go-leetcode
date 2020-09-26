package main

import "fmt"

func main() {
	//root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	root := buildTree([]int{3, 9, 8, 5, 4, 10, 20, 15, 7}, []int{4, 5, 8, 10, 9, 3, 15, 20, 7})
	fmt.Printf("%#+v", root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1.若上一个元素不等于中序遍历下标指向的元素，则将当前元素作为其上一个元素的左子节点，并将当前元素压入栈内。
2.若上一个元素等于中序遍历下标指向的元素，则从栈内弹出一个元素，同时令中序遍历下标指向下一个元素，
之后继续判断栈顶元素是否等于中序遍历下标指向的元素，若相等则重复该操作，直至栈为空或者元素不相等。
然后令当前元素为最后一个想等元素的右节点。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	length := len(preorder)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	index := 0
	for i := 1; i < length; i++ {
		value := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[index] {
			node.Left = &TreeNode{Val: value}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index++
			}
			node.Right = &TreeNode{Val: value}
			stack = append(stack, node.Right)
		}
	}
	return root
}
