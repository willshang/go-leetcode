package main

import "fmt"

func main() {
	first := ListNode{Val: 2}
	first1 := ListNode{Val: 4}
	first2 := ListNode{Val: 3}
	first.Next = &first1
	first1.Next = &first2
	fmt.Println(rotateRight(&first, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode61_旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	temp := head
	count := 1
	for temp.Next != nil {
		temp = temp.Next
		count++
	}
	temp.Next = head
	k = k % count
	for i := 0; i < count-k; i++ {
		temp = temp.Next
	}
	head, temp.Next = temp.Next, nil
	return head
}
