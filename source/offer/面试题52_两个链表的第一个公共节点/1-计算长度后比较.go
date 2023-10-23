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
	ALength := 0
	A := headA
	for A != nil {
		ALength++
		A = A.Next
	}
	BLength := 0
	B := headB
	for B != nil {
		BLength++
		B = B.Next
	}

	pA := headA
	pB := headB
	if ALength > BLength {
		n := ALength - BLength
		for n > 0 {
			pA = pA.Next
			n--
		}
	} else {
		n := BLength - ALength
		for n > 0 {
			pB = pB.Next
			n--
		}
	}

	for pA != pB {
		pA = pA.Next
		pB = pB.Next
	}
	return pA
}
