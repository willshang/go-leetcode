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

	fmt.Println(hasCycle(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode 141 环形链表
func hasCycle(head *ListNode) bool {
	for head != nil {
		if head.Val == math.MaxInt32 {
			return true
		}
		head.Val = math.MaxInt32
		head = head.Next
	}
	return false
}
