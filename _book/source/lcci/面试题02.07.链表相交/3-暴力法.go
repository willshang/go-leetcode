package main

import "fmt"

func main() {
	a := ListNode{Val: 10}
	b := ListNode{Val: 99}
	c := ListNode{Val: 100}
	d := ListNode{Val: 80}

	a.Next = &b
	d.Next = &b
	b.Next = &c
	fmt.Println(getIntersectionNode(&a, &d))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for A != nil {
		for B != nil {
			if A == B {
				return A
			}
			B = B.Next
		}
		A = A.Next
		B = headB
	}
	return nil
}
