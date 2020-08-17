package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = reverseKGroup(first, 2)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{Next: head}
	prev := res
	left := res
	for head != nil {
		right := left
		for i := 1; i <= k; i++ {
			right = right.Next
			if right == nil {
				return res.Next
			}
		}
		next := right.Next
		right.Next = nil
		left = reverse(left)
		prev.Next = left
		prev = next
		left, right = next, next
		head = next
	}
	return res.Next
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		next := head.Next
		head.Next = result
		result = head
		head = next
	}
	return result
}
