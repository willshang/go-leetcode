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
	if head == nil {
		return []int{}
	}
	res := reversePrint(head.Next)
	res = append(res, head.Val)
	return res
}
