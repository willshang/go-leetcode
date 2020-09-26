package main

import (
	"fmt"
	"math"
)

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &second

	fmt.Println(detectCycle(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	for head != nil {
		if head.Val == math.MaxInt32 {
			return head
		}
		head.Val = math.MaxInt32
		head = head.Next
	}
	return head
}
