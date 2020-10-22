package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(levelOrder(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 程序员面试金典04.03_特定深度节点链表
func listOfDepth(tree *TreeNode) []*ListNode {
	res := make([]*ListNode, 0)
	if tree == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)
	for len(queue) > 0 {
		length := len(queue)
		node := &ListNode{}
		tempNode := node
		for i := 0; i < length; i++ {
			node := queue[i]
			tempNode.Next = &ListNode{
				Val: node.Val,
			}
			tempNode = tempNode.Next
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, node.Next)
		queue = queue[length:]
	}
	return res
}
