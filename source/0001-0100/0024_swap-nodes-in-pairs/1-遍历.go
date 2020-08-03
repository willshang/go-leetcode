package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = swapPairs(first)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode24_两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	temp := &ListNode{Next: head}
	prev := temp
	for head != nil && head.Next != nil {
		first, second := head, head.Next
		prev.Next = second
		first.Next, second.Next = second.Next, first
		prev, head = first, first.Next
	}
	return temp.Next
}
