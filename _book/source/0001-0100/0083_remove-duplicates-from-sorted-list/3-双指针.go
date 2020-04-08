package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	deleteDuplicates(&first)

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	for p.Next != nil {
		if p.Val == q.Val {
			if q.Next == nil {
				p.Next = nil
			} else {
				p.Next = q.Next
				q = q.Next
			}
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return head
}
