package main

import "fmt"

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	reorderList(&first)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	second := reverse(slow.Next)
	slow.Next = nil
	cur := head
	count := 0
	for cur != nil && second != nil {
		a := cur.Next
		b := second.Next
		if count%2 == 0 {
			cur.Next = second
			cur = a
		} else {
			second.Next = cur
			second = b
		}
		count++
	}
}

func reverse(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}
