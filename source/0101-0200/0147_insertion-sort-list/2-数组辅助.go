package main

import (
	"fmt"
	"sort"
)

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth

	fmt.Println(insertionSortList(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	res := &ListNode{Next: head}
	cur := res
	arr[len(arr)-1].Next = nil
	for i := 0; i < len(arr); i++ {
		cur.Next = arr[i]
		cur = cur.Next
	}
	return res.Next
}
