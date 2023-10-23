package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third
	fmt.Println(kthToLast(&first, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	if count < k {
		return -1
	}
	for i := 0; i < count-k; i++ {
		head = head.Next
	}
	return head.Val
}
