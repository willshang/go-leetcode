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

// 程序员面试金典02.02_返回倒数第k个节点
func kthToLast(head *ListNode, k int) int {
	fast := head
	for k > 0 && head != nil {
		fast = fast.Next
		k--
	}
	if k > 0 {
		return -1
	}
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
