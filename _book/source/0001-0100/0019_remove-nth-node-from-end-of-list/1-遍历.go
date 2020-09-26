package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = removeNthFromEnd(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}
	cur := temp
	total := 0
	for cur.Next != nil {
		cur = cur.Next
		total++
	}
	cur = temp
	count := 0
	for cur.Next != nil {
		if total-n == count {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
		count++
	}
	return temp.Next
}
