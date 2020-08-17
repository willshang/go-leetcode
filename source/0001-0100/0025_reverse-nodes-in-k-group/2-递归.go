package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = reverseKGroup(first, 2)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := 0
	temp := head
	for temp != nil {
		res++
		temp = temp.Next
	}
	if res < k || k <= 1 {
		return head
	}
	pre := &ListNode{}
	cur := head
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = reverseKGroup(cur, k)
	return pre
}
