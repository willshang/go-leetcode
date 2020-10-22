package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = removeNthFromEnd(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode19_删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}
	fast, slow := temp, temp
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return temp.Next
}
