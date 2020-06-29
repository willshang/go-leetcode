package main

import "fmt"

func main() {
	first := ListNode{Val: 2}
	first1 := ListNode{Val: 4}
	first2 := ListNode{Val: 3}
	first.Next = &first1
	first1.Next = &first2

	second := ListNode{Val: 5}
	second1 := ListNode{Val: 6}
	second2 := ListNode{Val: 4}
	second.Next = &second1
	second1.Next = &second2

	fmt.Println(addTwoNumbers(&first, &second))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	res := &ListNode{Val: sum % 10}
	if sum >= 10 {
		l1.Next = addTwoNumbers(l1.Next, &ListNode{Val: 1})
	}
	res.Next = addTwoNumbers(l1.Next, l2.Next)
	return res
}
