package main

import "fmt"

func main() {
	five := ListNode{Val: 5}
	four := ListNode{Val: 4, Next: &five}
	three := ListNode{Val: 3, Next: &four}
	two := ListNode{Val: 2, Next: &three}
	one := ListNode{Val: 1, Next: &two}

	fmt.Println(middleNode(&one))
	six := ListNode{Val: 6}
	five.Next = &six

	fmt.Println(middleNode(&one))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
