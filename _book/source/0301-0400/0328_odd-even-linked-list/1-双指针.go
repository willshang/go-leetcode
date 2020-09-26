package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	oddEvenList(&first)

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

func oddEvenList(head *ListNode) *ListNode {
	odd := &ListNode{}
	even := &ListNode{}
	a := odd
	b := even
	count := 1
	for head != nil {
		if count%2 == 1 {
			a.Next = head
			a = head
		} else {
			b.Next = head
			b = head
		}
		count++
		head = head.Next
	}
	b.Next = nil
	a.Next = even.Next
	return odd.Next
}
