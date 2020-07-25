package main

import (
	"fmt"
)

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	reverseBetween(&first, 2, 3)
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

// leetcode92_反转链表II
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil {
		return head
	}
	temp := &ListNode{Next: head}
	prev := temp
	for i := 1; i < m; i++ {
		prev = prev.Next
	}
	head = prev.Next
	for i := m; i < n; i++ {
		next := head.Next
		head.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return temp.Next
}
