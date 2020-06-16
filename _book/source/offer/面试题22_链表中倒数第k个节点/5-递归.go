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

var res *ListNode
var count int

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		count = 0
		return nil
	}
	getKthFromEnd(head.Next, k)
	count++
	if count == k {
		res = head
	}
	return res
}
