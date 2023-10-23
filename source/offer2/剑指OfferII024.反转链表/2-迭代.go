package main

import "fmt"

func main() {
	first := new(ListNode)
	first.Val = 1
	second := new(ListNode)
	second.Val = 2
	third := new(ListNode)
	third.Val = 6
	fourth := new(ListNode)
	fourth.Val = 3
	fifth := new(ListNode)
	fifth.Val = 4
	sixth := new(ListNode)
	sixth.Val = 5
	seventh := new(ListNode)
	seventh.Val = 6
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = sixth
	sixth.Next = seventh
	first = reverseList(first)
	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指OfferII024.反转链表
func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = result
		result = head
		head = temp
	}
	return result
}
