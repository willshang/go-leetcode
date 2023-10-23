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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{Next: head}
	cur := head
	var left, right *ListNode
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	for i := 1; i < length; i = i * 2 {
		cur = res.Next
		tail := res
		for cur != nil {
			left = cur
			right = split(left, i)
			cur = split(right, i)
			tail.Next = mergeTwoLists(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return res.Next
}

func split(head *ListNode, length int) *ListNode {
	cur := head
	var right *ListNode
	length--
	for length > 0 && cur != nil {
		length--
		cur = cur.Next
	}
	if cur == nil {
		return nil
	}
	right = cur.Next
	cur.Next = nil
	return right
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	} else {
		temp.Next = l2
	}
	return res.Next
}
