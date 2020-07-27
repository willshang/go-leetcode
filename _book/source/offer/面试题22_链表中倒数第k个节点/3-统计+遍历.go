package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third
	fmt.Println(getKthFromEnd(&first, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	if count < k {
		return nil
	}
	for i := 0; i < count-k; i++ {
		head = head.Next
	}
	return head
}
