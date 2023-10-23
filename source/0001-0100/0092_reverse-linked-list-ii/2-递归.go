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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var next *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		next = head.Next
		return head
	}
	prev := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = next
	return prev
}
