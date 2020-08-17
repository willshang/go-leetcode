package main

import "fmt"

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth

	fmt.Println(insertionSortList(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{Next: head}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		prev := res
		next := cur.Next

	}
	return res.Next
}
