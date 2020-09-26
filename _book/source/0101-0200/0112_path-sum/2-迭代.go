package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	left.Right = &right
	res := hasPathSum(&root, 5)
	fmt.Println(res)
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	list1 := list.New()
	list2 := list.New()

	list1.PushFront(root)
	list2.PushFront(sum - root.Val)
	for list1.Len() > 0 {
		length := list1.Len()

		for i := 0; i < length; i++ {
			node := list1.Remove(list1.Back()).(*TreeNode)
			currentSum := list2.Remove(list2.Back()).(int)
			if node.Left == nil && node.Right == nil && currentSum == 0 {
				return true
			}
			if node.Left != nil {
				list1.PushFront(node.Left)
				list2.PushFront(currentSum - node.Left.Val)
			}
			if node.Right != nil {
				list1.PushFront(node.Right)
				list2.PushFront(currentSum - node.Right.Val)
			}
		}
	}
	return false
}
