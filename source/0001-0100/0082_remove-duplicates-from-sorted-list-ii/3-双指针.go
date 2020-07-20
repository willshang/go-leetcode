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
	temp := &ListNode{Next: head}
	second := temp
	first := head
	for first != nil && first.Next != nil {
		if first.Next.Val != second.Next.Val {
			first = first.Next
			second = second.Next
		} else {
			for first.Next != nil && first.Next.Val == second.Next.Val {
				first = first.Next
			}
			second.Next = first.Next
			first = first.Next
		}
	}
	return temp.Next
}
