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
	flag := false
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
		flag = true
	}
	head.Next = deleteDuplicates(head.Next)
	if flag {
		return head.Next
	}
	return head
}
