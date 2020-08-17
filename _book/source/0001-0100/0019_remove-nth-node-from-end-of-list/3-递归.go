package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = removeNthFromEnd(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var count int

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		count = 0
		return nil
	}
	head.Next = removeNthFromEnd(head.Next, n)
	count = count + 1
	if count == n {
		return head.Next
	}
	return head
}
