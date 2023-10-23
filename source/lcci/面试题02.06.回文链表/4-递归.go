package main

import "fmt"

func main() {
	first := new(ListNode)
	first.Val = 1
	second := new(ListNode)
	second.Val = 2
	third := new(ListNode)
	third.Val = 2
	fourth := new(ListNode)
	fourth.Val = 1
	first.Next = second
	second.Next = third
	third.Next = fourth

	fmt.Println(isPalindrome(first))
	for {
		if first == nil {
			break
		}
		fmt.Print(first.Val, " ")
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var p *ListNode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if p == nil {
		p = head
	}
	if isPalindrome(head.Next) && (p.Val == head.Val) {
		p = p.Next
		return true
	}
	p = nil
	return false
}
