package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third

	deleteNode(&first)
	fmt.Println(first.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 程序员面试金典02.03_删除中间节点
func deleteNode(node *ListNode) {
	// *node = *node.Next
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
