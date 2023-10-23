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
	length := getLength(head)
	if length < k || k <= 1 {
		return head
	}
	pre := &ListNode{}
	cur := head
	for i := 0; i < k; i++ {
		temp := cur
		cur = cur.Next
		temp.Next = pre
		pre = temp
	}
	head.Next = reverseKGroup(cur, k)
	return pre
}

func getLength(head *ListNode) int {
	if head == nil {
		return 0
	}
	temp := head
	res := 0
	for temp != nil {
		res++
		temp = temp.Next
	}
	return res
}
