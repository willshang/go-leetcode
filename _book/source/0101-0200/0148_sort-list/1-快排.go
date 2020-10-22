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

	fmt.Println(sortList(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode148_重排链表
func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return
	}
	temp := head.Val
	fast, slow := head.Next, head
	for fast != end {
		if fast.Val < temp {
			slow = slow.Next
			slow.Val, fast.Val = fast.Val, slow.Val
		}
		fast = fast.Next
	}
	slow.Val, head.Val = head.Val, slow.Val
	quickSort(head, slow)
	quickSort(slow.Next, end)
}
