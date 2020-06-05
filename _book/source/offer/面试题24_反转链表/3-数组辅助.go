package main

import "fmt"

func main() {
	first := new(ListNode)
	first.Val = 1
	second := new(ListNode)
	second.Val = 2
	third := new(ListNode)
	third.Val = 6
	fourth := new(ListNode)
	fourth.Val = 3
	fifth := new(ListNode)
	fifth.Val = 4
	sixth := new(ListNode)
	sixth.Val = 5
	seventh := new(ListNode)
	seventh.Val = 6
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = sixth
	sixth.Next = seventh
	first = reverseList(first)
	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	result := &ListNode{}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	temp := result
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i].Next = nil
		temp.Next = arr[i]
		temp = temp.Next
	}
	return result.Next
}
