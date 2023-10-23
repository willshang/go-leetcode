package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third

	deleteNode(&first, 3)
	fmt.Println(first.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = deleteNode(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
