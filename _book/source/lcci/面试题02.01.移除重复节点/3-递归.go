package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	removeDuplicateNodes(&first)
	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = *first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var m map[int]bool

func removeDuplicateNodes(head *ListNode) *ListNode {
	m = make(map[int]bool)
	return remove(head)

}

func remove(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if m[head.Val] == true {
		return remove(head.Next)
	}
	m[head.Val] = true
	head.Next = remove(head.Next)
	return head
}
