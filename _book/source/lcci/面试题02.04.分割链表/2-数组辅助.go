package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	partition(&first, 3)

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

func partition(head *ListNode, x int) *ListNode {
	a := make([]*ListNode, 0)
	b := make([]*ListNode, 0)

	for head != nil {
		if head.Val < x {
			a = append(a, head)
		} else {
			b = append(b, head)
		}
		head = head.Next
	}
	temp := &ListNode{}
	node := temp
	for i := 0; i < len(a); i++ {
		node.Next = a[i]
		node = node.Next
	}
	for i := 0; i < len(b); i++ {
		node.Next = b[i]
		node = node.Next
	}
	node.Next = nil
	return temp.Next
}
