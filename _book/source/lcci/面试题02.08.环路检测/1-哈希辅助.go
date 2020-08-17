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
	fourth.Next = &second

	fmt.Println(detectCycle(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 程序员面试金典02.08_环路检测
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}
