package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	removeDuplicateNodes(&first)
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

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	for temp != nil {
		second := temp
		for second.Next != nil {
			if second.Next.Val == temp.Val {
				second.Next = second.Next.Next
			} else {
				second = second.Next
			}
		}
		temp = temp.Next
	}
	return head
}
