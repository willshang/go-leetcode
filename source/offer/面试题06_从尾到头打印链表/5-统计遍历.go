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
	cur := head
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	res := make([]int, count)
	for cur != nil {
		res[count-1] = cur.Val
		count--
		cur = cur.Next
	}
	return res
}
