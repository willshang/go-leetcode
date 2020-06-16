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
	res := make([]*ListNode, 0)
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	return res[len(res)/2]
}
