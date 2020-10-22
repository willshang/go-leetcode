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
	fourth.Next = &second

	fmt.Println(detectCycle(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode142_环形链表II
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
