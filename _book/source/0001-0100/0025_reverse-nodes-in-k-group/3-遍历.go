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
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return res.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return res.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	temp := head
	for prev != tail {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return tail, head
}
