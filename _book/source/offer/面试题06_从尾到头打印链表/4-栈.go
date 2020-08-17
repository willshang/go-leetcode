package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 3}
	third := ListNode{Val: 2}
	first.Next = &second
	second.Next = &third
	fmt.Println(reversePrint(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	if head == nil {
		return res
	}
	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
	}
	return res
}
