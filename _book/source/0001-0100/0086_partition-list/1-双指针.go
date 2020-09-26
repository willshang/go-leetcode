package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	partition(&first, 3)

	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = *first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode86_分隔链表
func partition(head *ListNode, x int) *ListNode {
	first := &ListNode{}
	second := &ListNode{}
	a := first
	b := second
	for head != nil {
		if head.Val < x {
			a.Next = head
			a = head
		} else {
			b.Next = head
			b = head
		}
		head = head.Next
	}
	b.Next = nil
	a.Next = second.Next
	return first.Next
}
