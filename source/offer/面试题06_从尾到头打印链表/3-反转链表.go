package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 3}
	third := ListNode{Val: 2}
	first.Next = &second
	second.Next = &third
	fmt.Println(reversePrint(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	if head == nil {
		return res
	}
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res
}
